package service

import (
	"context"
	"packageServe/model"
)

type UserServer struct {
}

func (u UserServer) Login(ctx context.Context, in *LoginRequest) (*LoginResponse, error) {
	var err error
	var code int32
	if in.Typ == 0 {
		err = model.Login(in.Username,in.Password)
	} else {
		err = model.Register(in.Username,in.Password)
	}
	if err == nil {
		code = 0
	} else {
		code = 1001
	}
	return &LoginResponse{Code:code},nil
}