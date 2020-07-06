package main

import (
	"fmt"
	"net"

	dbx "github.com/go-ozzo/ozzo-dbx"

	auth "github.com/AlekseiAnisimov/auth/packages/auth"
	register "github.com/AlekseiAnisimov/auth/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":1111")
	if err != nil {
		fmt.Println("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	register.RegisterRegisterServiceServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

func (s *server) Registration(ctn context.Context, request *register.RegisterRequest) (response *register.RegisterResponse, err error) {
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
	envDb := auth.Env.GetEnvDbPointer()

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
