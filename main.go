package main

import (
	"net/http"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"dengyue.org/gomicro/services"
)

func main() {
	user := services.UserService{}
	endpoint := services.GetUserEndpoint(user)
	serverHandler := transport.NewServer(endpoint, services.DecodeUserRequest, services.EncodeUserResponse)

	r := mux.NewRouter()
	// r.Handle(`/user/{uid:\d+}`, serverHandler)
	r.Methods("GET","DELETE").Path(`/user/{uid:\d+}`).Handler(serverHandler)

	http.ListenAndServe(":8080", r)
}
