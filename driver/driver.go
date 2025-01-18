package driver

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func OpenDB(dsn string, maxConnections int, maxIdleConnections int, idleTime int) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	db.SetMaxOpenConns(maxConnections)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetConnMaxIdleTime(time.Duration(idleTime) * time.Minute)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, err
}
