package database

import (
	"hst_manag/internal/app/models/users"

	"gorm.io/gorm"
)

func IsUserExist(email string) (bool, error) {
	var existingUser users.Users
	err := DB.Table("users").Where("email = ?", email).Find(&existingUser).Error
	return (email == existingUser.Email), err
}

func CreateAdmin(users *users.Admin) (*users.Admin, error) {
	err := DB.Table("users").Create(&users).Error
	return users, err
}

func CreateUser(request *users.Users) (*users.Users, error) {
	err := DB.Table("users").Create(&request).Error
	return request, err
}

func CreateUserDetails(userDetail *users.UsersDetails) (*users.UsersDetails, error) {
	err := DB.Table("user_details").Create(&userDetail).Error
	return userDetail, err
}

func FindByEmail(email string) (*users.Users, error) {
	var user users.Users
	err := DB.Table("users").Where("email = ?", email).Find(&user).Error
	return &user, err
}

func IsAdminExist(roleId int, email string) (bool, error) {
	var admin users.Users
	err := DB.Where("role_id = ? and email = ?", 1, email).First(&admin).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func GetStudentsDetails() ([]users.CreateUserRequest, error) {
	var students []users.CreateUserRequest
	err := DB.Table("users").
		Select(`users.role_id, users.email, users.phone, users.password, users.first_name, users.middle_name, users.last_name, users.address_line1, users.address_line2, users.city, users.state, users.aadhar_card, users.status,DATE_FORMAT(user_details.dob, '%d-%m-%Y') as dob,user_details.gender,user_details.father_name,user_details.father_contact,user_details.father_aadhar_card,user_details.college_name,user_details.course,user_details.guardian_name,user_details.guardian_aadhar_card,user_details.guardian_relationship,user_details.guardian_contact,user_details.blood_group,user_details.any_chronic_disease,user_details.allergies`).
		Joins("left join user_details on user_details.user_id = users.id").
		Where("users.role_id = ?", 2).
		Find(&students).Error
	return students, err
}
