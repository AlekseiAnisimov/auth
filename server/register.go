package main

import (
	"log"
	"net"

	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"

	auth "github.com/AlekseiAnisimov/auth/packages/auth"
	register "github.com/AlekseiAnisimov/auth/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	envDb auth.Env
}

func main() {
	listener, err := net.Listen("tcp", ":1111")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	conn, _ := dbConnection()

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	register.RegisterRegisterServiceServer(grpcServer, &server{envDb: conn})
	grpcServer.Serve(listener)
}

func (s *server) Registration(ctn context.Context, request *register.RegisterRequest) (response *register.RegisterResponse, err error) {
	env := s.envDb
	login := request.Login
	email := request.Email
	password := request.Password

	if login == "" || email == "" || password == "" {
		response = &register.RegisterResponse{
			Message: "The field must be filled in",
		}
		return response, nil
	}

	err = auth.IsValidEmail(email)

	if err != nil {
		response = &register.RegisterResponse{
			Message: "Email no valid",
		}
		return response, nil
	}

	data := auth.UserIdentityData{Login: login, Email: email, Password: password}
	data.Password = data.PasswordToMd5()

	user := auth.UserIdentityData{}

	envDb := env.GetEnvDbPointer()

	_ = envDb.Select("*").From("identity").Where(dbx.HashExp{"login": login}).One(&user)

	if user.Login != "" {
		response = &register.RegisterResponse{
			Message: "Such user is exist",
		}

		return response, nil
	}

	_ = envDb.Model(&data).Insert()

	response = &register.RegisterResponse{
		Message: "Success",
		//UserData: &data,
	}

	return response, nil
}

//TODO:
//  1) переписать эту шляпу
// 2) вынести куда-нибудь
func dbConnection() (envDb auth.Env, err error) {
	dbconf := auth.DbConfig{}
	err = dbconf.GetDbParamsFromYaml()
	if err != nil {
		log.Fatalf(" Получение данных для БД %v", err)
	}

	dialect := &dbconf.Development.Dialect
	datasource := &dbconf.Development.Datasource

	db, err := dbx.Open(*dialect, *datasource)

	env := auth.Env{}
	env.SetEnvDbPointer(db)

	return env, nil
}
