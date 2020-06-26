package main

import (
	"encoding/json"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"net"
	"fmt"
	register "../proto"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	auth "github.com/AlekseiAnisimov/auth"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":1111")
	if err != nil {
		fmt.Println("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	register.RegisterServiceServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

func (s *server) Registration(ctn *context.Context, request *register.RegisterRequest) (response *register.RegisterResponse, err error) {
	login := request.Login
	email := request.Email
	password := request.Password

	if login == "" || email == "" || password == "" {
		response = &register.RegisterResponse{
			Message: "The field must be filled in"
		}
		return response, nil
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
}