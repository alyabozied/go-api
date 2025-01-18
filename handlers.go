package main

import (
	"api/model/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
	Content string `json:"content"`
	ID      int    `json:"id"`
}

func (app *application) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserPayload CreateUserPayload
	err := readJson(w, r, &createUserPayload)
	if err != nil {
		app.errorLog.Println(err.Error())
		writeJsonError(w, http.StatusBadRequest, "bad request")
		return
	}
	encyptedPasswod, err := bcrypt.GenerateFromPassword([]byte(createUserPayload.Password), bcrypt.DefaultCost)
	if err != nil {
		app.errorLog.Println(err.Error())
		writeJsonError(w, http.StatusInternalServerError, "internal server eror")
		return
	}

	user := store.User{
		Username: createUserPayload.Username,
		Email:    createUserPayload.Email,
		Password: encyptedPasswod,
	}

	id, err := app.storage.Users.Create(r.Context(), &user)
	type createdUser struct {
		ID      int    `json:"id"`
		Message string `json:"message"`
	}
	if err != nil {
		writeJsonError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJson(w, http.StatusCreated, createdUser{ID: id, Message: "user created"})

}

func (app *application) GetUserByID(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	userID, _ := strconv.Atoi(id)

	user, err := app.storage.Users.GetByID(r.Context(), userID)
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "eror")
		app.errorLog.Println(err)
		return
	}

	if err := writeJson(w, http.StatusOK, user); err != nil {
		app.errorLog.Println(err)
		writeJsonError(w, http.StatusInternalServerError, "eror")

	}
}
func (app *application) SignIn(w http.ResponseWriter, r *http.Request) {
	type Payload struct {
		Email    string
		Password string
	}
	var payload Payload
	readJson(w, r, &payload)
	user, err := app.storage.Users.GetByEmail(r.Context(), payload.Email)
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, "eror")
		app.errorLog.Println(err)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		writeJsonError(w, http.StatusUnauthorized, "Wrong credentials")
		return
	}
	type Response struct {
		Token string
	}
	if err := writeJson(w, http.StatusOK, Response{Token: "my token"}); err != nil {
		app.errorLog.Println(err)
		writeJsonError(w, http.StatusInternalServerError, "error")

	}
}
