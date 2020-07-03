package auth

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Registration функция регистрации
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

	err := IsValidEmail(*email)

	if err != nil {
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Email no valid",
		})
		return
	}

	data.Password = data.PasswordToMd5()

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

// IdentityByLogin аутентификация по логину
func (env *Env) IdentityByLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data UserIdentityData
	_ = json.NewDecoder(r.Body).Decode(&data)

	login := &data.Login
	password := data.PasswordToMd5()

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
	tokenExpired := int32(time.Now().Unix() + 20800)

	_, _ = env.db.Update("identity", dbx.Params{"token": token, "token_expired": tokenExpired}, dbx.HashExp{"id": user.ID}).Execute()

	type Result struct {
		ID    int    `json:"id"`
		Token string `json:"token"`
	}
	result := Result{user.ID, token}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// IdentityByEmail аутентификация по email
func (env *Env) IdentityByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data UserIdentityData
	_ = json.NewDecoder(r.Body).Decode(&data)

	email := &data.Email
	password := data.PasswordToMd5()

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

	_, _ = env.db.Update("identity", dbx.Params{"token": token, "token_expired": toketExpired}, dbx.HashExp{"id": user.ID}).Execute()

	type Result struct {
		ID    int    `json:"id"`
		Token string `json:"token"`
	}
	result := Result{user.ID, token}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// IsValidEmail проверка на валидность email
func IsValidEmail(email string) error {
	pattern := `^([a-z0-9_-]+\.)*[a-z0-9_-]+@[a-z0-9_-]+(\.[a-z0-9_-]+)*\.[a-z]{2,6}$`
	res, _ := regexp.MatchString(pattern, email)

	if res != true {
		return errors.New("email valid error")
	}

	return nil
}

func tokenGenerator() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// CheckToken проверка на действительность токена
func (env *Env) CheckToken(w http.ResponseWriter, r *http.Request) {
	bearer := r.Header.Get("Authorization")
	splitToken := strings.Split(bearer, "Bearer ")
	token := splitToken[1]

	type Result struct {
		Cnt int `db:"cnt"`
	}

	result := Result{}

	err := env.db.Select("count(*) as cnt").
		From("identity").
		Where(dbx.HashExp{"token": token}).
		AndWhere(dbx.NewExp("token_expired > {:time_now}", dbx.Params{"time_now": int32(time.Now().Unix())})).
		One(&result)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	w.WriteHeader(200)
	if result.Cnt == 0 {
		json.NewEncoder(w).Encode(map[string]bool{
			"active": false,
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{
		"active": true,
	})
	return

}
