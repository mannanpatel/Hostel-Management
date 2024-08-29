package helper

import (
	"encoding/json"
	"fmt"
	"hst_manag/internal/app/models/users"
	"hst_manag/internal/database"
	"hst_manag/internal/utils"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetErrorJson(message interface{}, statuscode int) map[string]interface{} {
	return map[string]interface{}{"message": message, "success": false, "code": statuscode}
}

func GetCommonErrorResponse(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{"message": message, "success": false, "code": http.StatusForbidden}
}

func GetCommonResponse(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{"message": message, "success": true, "code": http.StatusOK, "data": data}
}

func GetErrorMessage(data map[string]interface{}) string {
	return data["error"].(map[string]interface{})["message"].(string)
}

func JSONMarshal(a any) string {
	jsonData, err := json.Marshal(a)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return ""
	}
	jsonString := string(jsonData)
	return jsonString
}

func CreateAdmin() {

	const adminEmail = "admin@example.com"
	const adminRoleID = 1
	const adminPhone = "9898989898"
	const adminfirstname = "admin"
	// Hash the password using a utility function
	hashedPassword, err := utils.HashPassword(os.Getenv("Password"))
	if err != nil {
		log.Fatal("Error hashing admin password:", err)
	}
	// Check if an admin already exists
	isAdminExist, err := database.IsAdminExist(adminRoleID, adminEmail)
	if err != nil {
		log.Fatal("Error checking admin existence:", err)
	}
	if isAdminExist {
		log.Println("Admin user already exists.")
		return
	}
	// If admin doesn't exist, create it
	admin := users.Admin{
		RoleID:    adminRoleID,
		Email:     adminEmail,
		Phone:     adminPhone,
		Password:  hashedPassword,
		FirstName: adminfirstname,
		Status:    1, // Active status
	}

	adminCreate, err := database.CreateAdmin(&admin)
	if err != nil {
		log.Fatal("Error creating admin user:", err)
	}
	log.Println("Admin user created successfully:", adminCreate)
}
