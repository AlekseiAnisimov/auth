package main

import (
	"auth/packages/auth"
	"net/http"

	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	dbconf := auth.DbConfig{}
	err := dbconf.GetDbParamsFromYaml()
	if err != nil {
		panic(err)
	}

	dialect := &dbconf.Development.Dialect
	datasource := &dbconf.Development.Datasource

	var db, _ = dbx.Open(*dialect, *datasource)

	env := auth.Env{}
	env.SetEnvDbPointer(db)

	router := mux.NewRouter()
	router.HandleFunc("/registration", env.Registration).Methods("POST")
	router.HandleFunc("/login", env.IdentityByLogin).Methods("POST")
	router.HandleFunc("/login/email", env.IdentityByEmail).Methods("POST")
	router.HandleFunc("/token/check", env.CheckToken).Methods("GET", "POST")
	http.ListenAndServe(":8000", router)
}
