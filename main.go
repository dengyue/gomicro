package main

import (
	"dengyue.org/gomicro/util"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"dengyue.org/gomicro/services"
)

func main() {
	user := services.UserService{}
	endpoint := services.GetUserEndpoint(user)
	serverHandler := transport.NewServer(endpoint, services.DecodeUserRequest, services.EncodeUserResponse)

	r := mux.NewRouter()
	{
		// r.Handle(`/user/{uid:\d+}`, serverHandler)
		r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(serverHandler)
		r.Methods("GET").Path("/health").HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Content-type","application/json")
			rw.Write([]byte(`{"status":"ok"}`))
		})
	}

	errChan := make(chan error)
	go func() {
		util.RegService()
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			log.Println(err)
			errChan <- err
		}
	}()

	go func() {
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		errChan<- fmt.Errorf("%s", <-sigChan)
	}()

	result := <-errChan
	util.UnRegService()
	log.Println(result)

}
