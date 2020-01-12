package main

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

type UserIdentityData struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/registration", Registration)
	http.ListenAndServe(":8000", router)
}

func Registration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data UserIdentityData
	var passToHash []byte
	_ = json.NewDecoder(r.Body).Decode(&data)

	login := &data.Login
	email := &data.Email
	password := &data.Password
	phone := &data.Phone

	if *login == "" || *email == "" || *password == "" {
		fmt.Println(w, errors.New("The field must be filled in"))
	}

	err := isValidEmail(*email)

	if err != nil {
		fmt.Println(w, err)
	}

	passToHash = []byte(*password)
	passwordHash := md5.Sum(passToHash)

	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "tst")
	fmt.Println(w, passwordHash)
	fmt.Println(w, phone)
	json.NewEncoder(w).Encode(data)
}

func isValidEmail(email string) error {
	pattern := `^([a-z0-9_-]+\.)*[a-z0-9_-]+@[a-z0-9_-]+(\.[a-z0-9_-]+)*\.[a-z]{2,6}$`
	res, _ := regexp.MatchString(pattern, email)

	if res != true {
		return errors.New("email valid error")
	}

	return nil
}
