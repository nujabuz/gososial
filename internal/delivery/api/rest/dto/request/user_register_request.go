// Package request
package request

import "github.com/nujabuz/gososial/internal/application/params"

type RequestRegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *RequestRegisterUser) ToParam() *params.UserCreateParam {
	return &params.UserCreateParam{
		Username: r.Username,
		Password: r.Password,
	}
}
