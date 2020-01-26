package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-ozzo/ozzo-dbx"
)

func (env *Env) checkToken(w http.ResponseWriter, r *http.Request) {
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
