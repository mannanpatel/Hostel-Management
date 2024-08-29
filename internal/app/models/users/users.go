package users

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleName string `json:"role_name"`
}

type Admin struct {
	gorm.Model
	RoleID    int    `json:"role_id"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	Status    int    `json:"status"`
}

type Users struct {
	gorm.Model
	RoleID       int    `json:"role_id"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	MiddleName   string `json:"middle_name"`
	LastName     string `json:"last_name"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	State        string `json:"state"`
	AadharCard   string `json:"aadhar_card"`
	Status       int    `json:"status"`
}

type CreateUserRequest struct {
	RoleID               int    `json:"role_id"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	Password             string `json:"password"`
	FirstName            string `json:"first_name"`
	MiddleName           string `json:"middle_name"`
	LastName             string `json:"last_name"`
	AddressLine1         string `json:"address_line1"`
	AddressLine2         string `json:"address_line2"`
	City                 string `json:"city"`
	State                string `json:"state"`
	AadharCard           string `json:"aadhar_card"`
	Status               int    `json:"status"`
	DOB                  string `json:"DOB"`
	Gender               string `json:"gender"`
	FatherName           string `json:"father_name"`
	FatherContact        string `json:"father_contact"`
	FatherAadharCard     string `json:"father_aadhar_card"`
	CollegeName          string `json:"college_name"`
	Course               string `json:"course"`
	GuardianName         string `json:"guardian_name"`
	GuardianAadharCard   string `json:"guardian_aadhar_card"`
	GuardianRelationship string `json:"guardian_relationship"`
	GuardianContact      string `json:"guardian_contact"`
	BloodGroup           string `json:"blood_group"`
	AnyChronicDisease    string `json:"any_chronic_disease"`
	Allergies            string `json:"allergies"`
}

type UsersDetails struct {
	gorm.Model
	UserID               int       `json:"user_id"`
	DOB                  time.Time `json:"DOB"`
	Gender               string    `json:"gender"`
	FatherName           string    `json:"father_name"`
	FatherContact        string    `json:"father_contact"`
	FatherAadharCard     string    `json:"father_aadhar_card"`
	CollegeName          string    `json:"college_name"`
	Course               string    `json:"course"`
	GuardianName         string    `json:"guardian_name"`
	GuardianAadharCard   string    `json:"guardian_aadhar_card"`
	GuardianRelationship string    `json:"guardian_relationship"`
	GuardianContact      string    `json:"guardian_contact"`
	BloodGroup           string    `json:"blood_group"`
	AnyChronicDisease    string    `json:"any_chronic_disease"`
	Allergies            string    `json:"allergies"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInfo struct {
	gorm.Model
	Address_1   string
	Address_2   string
	City        string
	CountryCode string
	Email       string
	FirstName   string
	LastName    string
	Phone       string
	PostalCode  string
	State       string
	AddressID   int
}
