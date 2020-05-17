package handlers

import (
	"fmt"
	"net/http"
	"time"

	dbase "../dbase"
	models "../models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterLogin(db *dbase.DataBase, w http.ResponseWriter, r *http.Request) {
	//-------DTO----------------------------------------
	var new models.RegisterUser
	err := ReceiveJSON(r, &new)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//--------ENTITY for Users table----------------------
	user := models.User{
		Nickname: new.Nickname,
		RoleID:   3, // role:"user"
	}
	tx, err := db.DB.Begin()
	ID, err := db.CreateUser(user, tx)
	if err != nil && err.Error()[:6] == "UNIQUE" {
		SendJSON(w, models.Error{
			Status:      "Failed",
			Description: "User with such a nickname already exists, please try another one",
		})
		tx.Rollback()
		return
	}
	//---------ENTITY for Credentials table---------------
	HashedPW, err := bcrypt.GenerateFromPassword([]byte(new.Password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	cred := models.Credentials{
		ID:             ID,
		Email:          new.Email,
		HashedPassword: string(HashedPW),
	}
	err = db.CreateUserCredentials(cred)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	session := models.Session{UserID: ID}
	UUID, err := db.CreateSession(session)
	fmt.Println("Last created session's UUID:", UUID)
	err = SetCookie(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	tx.Commit()
}

func SetCookie(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("logged-in_forum")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name: "logged-in_forum",
			//////
			Value:    "1", //UUID
			Expires:  time.Now().Add(time.Hour * 1),
			Secure:   true,
			HttpOnly: true,
		}
	}
	http.SetCookie(w, cookie)
	return nil
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("logged-in_forum")
	if err != nil {
		fmt.Println("DeleteCookie error:")
		return err
	}
	cookie = &http.Cookie{
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	return nil
}

func Test(w http.ResponseWriter, r *http.Request) {

	for _, a := range r.Cookies() {
		fmt.Println(a.Name)
	}

	cookie, err := r.Cookie("logged-in_forum")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:     "logged-in_forum",
			Value:    "1",
			Expires:  time.Now().Add(time.Minute * 1),
			Secure:   true,
			HttpOnly: true,
		}
	} else {
		fmt.Println("1:", cookie.Value)
	}
	http.SetCookie(w, cookie)
	fmt.Println("AAAAAAA")
}
