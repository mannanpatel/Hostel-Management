package services

import (
	"hst_manag/internal/app/interfaces"
	"hst_manag/internal/app/models/users"
	gen "hst_manag/internal/utils/generic"
)

type studentServices struct {
	repository interfaces.IRepositoryStudent
}

func NewServicesStudent(repository interfaces.IRepositoryStudent) interfaces.IServicesStudent {
	return &studentServices{
		repository: repository,
	}
}

// func (userServices *userServices) SignUp(user users.Users) *res.GenericResponse {
// 	return userServices.repository.SignUp(user)
// }

func (studentServices *studentServices) Login(request users.UserLogin) *gen.GenericResponse {
	return studentServices.repository.Login(request)
}
