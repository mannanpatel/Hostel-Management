package repository

import (
	"fmt"
	"hst_manag/internal/app/interfaces"
	"hst_manag/internal/app/models/users"
	"hst_manag/internal/database"
	"hst_manag/internal/helper"
	"hst_manag/internal/utils"
	gen "hst_manag/internal/utils/generic"

	"log"

	"net/http"
	"time"
)

type adminRepository struct {
}

func NewRepositoryAdmin() interfaces.IRepositoryAdmin {
	return &adminRepository{}
}

func (adminRepository *adminRepository) AdminLogin(request users.UserLogin) *gen.GenericResponse {
	user, err := database.FindByEmail(request.Email)
	if err != nil || user == nil {
		return gen.HandleError(err, "Invalid Email ID")
	}
	isPasswordValid := utils.CheckPasswordHash(request.Password, user.Password)
	if !isPasswordValid {
		return gen.HandleError(nil, "Invalid Password")
	}
	token, _ := utils.GenerateJWT(user.Email, user.Phone)
	// token, err := generateJWT(*user)

	return gen.GetResponse(token, http.StatusOK)
}

func (adminRepository *adminRepository) CreateStudent(newUser users.CreateUserRequest) *gen.GenericResponse {
	isUserExist, err := database.IsUserExist(newUser.Email)
	if err != nil {
		return gen.GetResponse("Error checking user existence", http.StatusInternalServerError)
	}
	if isUserExist {
		return gen.GetResponse("User with this email already exists", http.StatusConflict)
	}
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return gen.GetResponse("Error hashing admin password:", http.StatusInternalServerError)
	}
	// Parse the date of birth
	birthDate, err := time.Parse("02-01-2006", newUser.DOB)
	if err != nil {
		return gen.GetResponse("Invalid date format for DOB", http.StatusBadRequest)
	}

	// Create the new user
	user := users.Users{
		RoleID:       newUser.RoleID,
		Email:        newUser.Email,
		Phone:        newUser.Phone,
		Password:     hashedPassword,
		FirstName:    newUser.FirstName,
		MiddleName:   newUser.MiddleName,
		LastName:     newUser.LastName,
		AddressLine1: newUser.AddressLine1,
		AddressLine2: newUser.AddressLine2,
		City:         newUser.City,
		State:        newUser.State,
		AadharCard:   newUser.AadharCard,
		Status:       1,
	}

	// Start a new transaction
	tx := database.DB.Begin()

	addUser, err := database.CreateUser(&user)
	if err != nil {
		tx.Rollback()
		return gen.HandleError(err, "Error Creating user")
	}

	// Create student details associated with the new user
	usersDetails := users.UsersDetails{
		UserID:               int(addUser.ID),
		DOB:                  birthDate,
		Gender:               newUser.Gender,
		FatherName:           newUser.FatherName,
		FatherContact:        newUser.FatherContact,
		FatherAadharCard:     newUser.FatherAadharCard,
		CollegeName:          newUser.CollegeName,
		Course:               newUser.Course,
		GuardianName:         newUser.GuardianName,
		GuardianAadharCard:   newUser.GuardianAadharCard,
		GuardianRelationship: newUser.GuardianRelationship,
		GuardianContact:      newUser.GuardianContact,
		BloodGroup:           newUser.BloodGroup,
		AnyChronicDisease:    newUser.AnyChronicDisease,
		Allergies:            newUser.Allergies,
	}

	addUserDetails, err := database.CreateUserDetails(&usersDetails)
	if err != nil {
		tx.Rollback()
		return gen.HandleError(err, "Error creating student details")
	}
	// Commit the transaction
	tx.Commit()

	subject := "Welcome to Hostel Management System"
	body := fmt.Sprintf(
		"<h1>Welcome, %s!</h1>"+
			"<p>We're excited to have you join our service.</p>"+
			"<p>Your login details are as follows:</p>"+
			"<p><strong>User ID (Email):</strong> %s</p>"+
			"<p><strong>Password:</strong> %s</p>"+
			"<p>Please keep this information secure.</p>"+
			"<p>Best regards,<br>Hostel Management</p>",
		user.FirstName, user.Email, newUser.Password)
	err = utils.SendEmail(user.Email, subject, body)
	if err != nil {
		log.Printf("Failed to send Welcome email to %s: %v", user.Email, err)
	}
	response := map[string]interface{}{
		"user":       addUser,
		"userDetail": addUserDetails,
	}
	return gen.GetResponse(response, http.StatusOK)
}

func (adminRepository *adminRepository) GetStudentsDetails() *gen.GenericResponse {
	studentDetails, err := database.GetStudentsDetails()
	if err != nil {
		return gen.GetResponse(helper.GetErrorJson(err.Error(), http.StatusBadRequest), http.StatusOK)
	} else {
		return gen.GetResponse(helper.GetCommonResponse("Data fetch succesfully", studentDetails), http.StatusOK)
	}
}
