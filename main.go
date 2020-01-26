package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type UserIdentityData struct {
	Id       int
	Login    string `json:"login"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type DbConfig struct {
	Development struct {
		Dialect    string
		Datasource string
	}
}

type Env struct {
	db *dbx.DB
}

var dbConfigFile = "dbconfig.yml"

func main() {
	dbconf := DbConfig{}
	err := dbconf.getDbParamsFromYaml()
	if err != nil {
		panic(err)
	}

	dialect := &dbconf.Development.Dialect
	datasource := &dbconf.Development.Datasource

	var db, _ = dbx.Open(*dialect, *datasource)
	env := Env{db: db}

	router := mux.NewRouter()
	router.HandleFunc("/registration", env.Registration).Methods("POST")
	router.HandleFunc("/login", env.IdentityByLogin).Methods("POST")
	router.HandleFunc("/login/email", env.IdentityByEmail).Methods("POST")
	http.ListenAndServe(":8000", router)
}

func (env *Env) Registration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data UserIdentityData

	_ = json.NewDecoder(r.Body).Decode(&data)

	login := &data.Login
	email := &data.Email
	password := &data.Password

	if *login == "" || *email == "" || *password == "" {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "The field must be filled in",
		})
		return
	}

	err := isValidEmail(*email)

	if err != nil {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Email no valid",
		})
		return
	}

	data.Password = data.passwordToMd5()

	user := UserIdentityData{}

	_ = env.db.Select("*").From("identity").Where(dbx.HashExp{"login": login}).One(&user)

	if user.Login != "" {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Such user is exist",
		})
		return
	}

	_ = env.db.Model(&data).Insert()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (env *Env) IdentityByLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data UserIdentityData
	_ = json.NewDecoder(r.Body).Decode(&data)

	login := &data.Login
	password := data.passwordToMd5()

	user := UserIdentityData{}
	_ = env.db.Select("*").From("identity").Where(dbx.HashExp{"login": login, "password": password}).One(&user)

	if user.Login == "" {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User not found",
		})
		return
	}

	token := tokenGenerator()
	toketExpired := int32(time.Now().Unix()) + 10800

	_, _ = env.db.Update("identity", dbx.Params{"token": token, "token_expired": toketExpired}, dbx.HashExp{"id": user.Id}).Execute()

	type Result struct {
		Id    int
		Token string
	}
	result := Result{user.Id, token}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (env *Env) IdentityByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data UserIdentityData
	_ = json.NewDecoder(r.Body).Decode(&data)

	email := &data.Email
	password := data.passwordToMd5()

	user := UserIdentityData{}
	_ = env.db.Select("*").From("identity").Where(dbx.HashExp{"email": email, "password": password}).One(&user)

	if user.Email == "" {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User not found",
		})
		return
	}

	token := tokenGenerator()
	toketExpired := int32(time.Now().Unix()) + 10800

	_, _ = env.db.Update("identity", dbx.Params{"token": token, "token_expired": toketExpired}, dbx.HashExp{"id": user.Id}).Execute()

	type Result struct {
		Id    int
		Token string
	}
	result := Result{user.Id, token}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func isValidEmail(email string) error {
	pattern := `^([a-z0-9_-]+\.)*[a-z0-9_-]+@[a-z0-9_-]+(\.[a-z0-9_-]+)*\.[a-z]{2,6}$`
	res, _ := regexp.MatchString(pattern, email)

	if res != true {
		return errors.New("email valid error")
	}

	return nil
}

func (dbconf *DbConfig) getDbParamsFromYaml() error {
	fopen, err := ioutil.ReadFile(dbConfigFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fopen, &dbconf)
	if err != nil {
		return err
	}

	return nil
}

func (u UserIdentityData) TableName() string {
	return "identity"
}

func (u UserIdentityData) passwordToMd5() string {
	passByte := []byte(u.Password)
	passwordHash := md5.Sum(passByte)
	passString := hex.EncodeToString(passwordHash[:])

	return passString
}

func tokenGenerator() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
