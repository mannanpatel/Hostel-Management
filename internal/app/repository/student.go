package repository

import (
	"hst_manag/internal/app/interfaces"
	"hst_manag/internal/app/models/users"
	"hst_manag/internal/database"
	"hst_manag/internal/utils"
	gen "hst_manag/internal/utils/generic"
	"net/http"
)

type studentRepository struct {
}

func NewRepositoryStudent() interfaces.IRepositoryStudent {
	return &studentRepository{}
}

// func (*userRepository) SignUp(user users.Users) *res.GenericResponse {
// 	var existingUser users.Users
// 	isUserExist, err := database.IsUserExist(&existingUser, user.Email)
// 	if err != nil {
// 		return res.HandleError(err, "Error in checking email eist")
// 	}

// 	if isUserExist {
// 		return res.HandleError(nil, "Enter email is already exists")
// 	}
// 	hashedPassword, err := utils.HashPassword(user.Password)
// 	if err != nil {
// 		return res.HandleError(err, "Error hashing password")
// 	}
// 	user.Password = hashedPassword
// 	response, err := database.CreateUser(&user)
// 	if err != nil {
// 		return res.HandleError(err, "Error Creating user")
// 	}
// 	return res.GetResponse(response, http.StatusOK)
// }

func (*studentRepository) Login(request users.UserLogin) *gen.GenericResponse {
	student, err := database.FindByEmail(request.Email)
	if err != nil || student == nil {
		return gen.HandleError(err, "Invalid Email ID")
	}
	isPasswordValid := utils.CheckPasswordHash(request.Password, student.Password)
	if !isPasswordValid {
		return gen.HandleError(nil, "Invalid Password")
	}
	token, _ := utils.GenerateJWT(student.Email, student.Phone)
	// token, err := generateJWT(*user)

	return gen.GetResponse(token, http.StatusOK)
}
