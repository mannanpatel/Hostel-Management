package interfaces

import (
	"hst_manag/internal/app/models/users"
	gen "hst_manag/internal/utils/generic"
)

type IServicesStudent interface {
	// SignUp(user users.Users) *gen.GenericResponse
	Login(request users.UserLogin) *gen.GenericResponse
}

type IServicesAdmin interface {
	AdminLogin(request users.UserLogin) *gen.GenericResponse
	CreateStudent(newUser users.CreateUserRequest) *gen.GenericResponse
	GetStudentsDetails() *gen.GenericResponse
}
