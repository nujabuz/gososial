// Package params
package params

type UserCreateParam struct {
	Username string
	Password string
}

type UserFindByIDParam struct {
	ID int64
}

type UserFindByUsernameParam struct {
	Username string
}
