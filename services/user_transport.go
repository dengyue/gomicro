package services

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	// if r.URL.Query().Get("uid") != "" {
	// 	uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	// 	return UserRequest{
	// 		Uid: uid,
	// 	}, nil
	// }

	vars := mux.Vars(r)
	if uid, ok := vars["uid"]; ok {
		uid, _ := strconv.Atoi(uid)
		return UserRequest{
			Uid:    uid,
			Method: r.Method,
		}, nil
	}
	return nil, errors.New("parameter error")
}

func EncodeUserResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
