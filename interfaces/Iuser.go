package interfaces

import "jayalsa/project_golang/entities"

type IUser interface {
	Register(user *entities.User) (*entities.SignupResponse, error)
	Login(user *entities.Login) (*entities.LoginResponse, error)
	Logout() error
}
