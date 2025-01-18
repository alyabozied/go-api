package main

import (
	"api/driver"
	"api/env"
	"api/model/store"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  int
	}
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
	storage  store.Storage
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf("%s", fmt.Sprintf("Starting HTTP server in %s mode on port %d\n", app.config.env, app.config.port))

	return srv.ListenAndServe()
}
func main() {
	var cfg config

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg.db.dsn = os.Getenv("database")

	cfg.db.maxOpenConns = env.GetInt("DB_MAX_OPEN_CONNS", 30)
	cfg.db.maxIdleConns = env.GetInt("DB_MAX_IDEL_CONNS", 30)
	cfg.db.maxIdleTime = env.GetInt("DB_MAX_IDEL_TIME", 30)

	cfg.port = env.GetInt("port", 4000)

	cfg.env = os.Getenv("enviroment")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := driver.OpenDB(cfg.db.dsn, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		errorLog.Fatal(err)

	}
	defer conn.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
		storage:  store.NewStorage(conn),
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}

}
