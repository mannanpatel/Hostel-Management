package services

import (
	"hst_manag/internal/app/interfaces"
	"hst_manag/internal/app/models/users"
	gen "hst_manag/internal/utils/generic"
)

type adminServices struct {
	repository interfaces.IRepositoryAdmin
}

// NewServicesAdmin creates a new instance of adminServices and returns it as IServicesAdmin.
func NewServicesAdmin(repository interfaces.IRepositoryAdmin) interfaces.IServicesAdmin {
	return &adminServices{
		repository: repository,
	}
}

func (adminServices *adminServices) AdminLogin(request users.UserLogin) *gen.GenericResponse {
	return adminServices.repository.AdminLogin(request)
}

func (adminServices *adminServices) CreateStudent(newUser users.CreateUserRequest) *gen.GenericResponse {
	return adminServices.repository.CreateStudent(newUser)
}

func (adminServices *adminServices) GetStudentsDetails() *gen.GenericResponse {
	return adminServices.repository.GetStudentsDetails()
}
