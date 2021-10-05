package services

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	Uid    int    `json:"uid"`
	Method string `json:"method"`
}
type UserResponse struct {
	Result string `json:"result"`
}

func GetUserEndpoint(userService IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		result := "not support"
		if r.Method == "GET" {
			result = userService.GetName(r.Uid)
			return UserResponse{Result: result}, nil
		} else if r.Method == "DELETE" {
			err := userService.DeleteUser(r.Uid)
			if err != nil {
				result = err.Error()
			} else {
				result = fmt.Sprintf("userid %d was deleted.\n", r.Uid)
			}
		}
		return UserResponse{Result: result}, nil
	}
}
