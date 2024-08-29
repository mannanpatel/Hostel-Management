package controller

import (
	"hst_manag/internal/app/interfaces"
	"hst_manag/internal/app/models/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type adminController struct {
	services interfaces.IServicesAdmin
}

func NewControllerAdmin(services interfaces.IServicesAdmin) interfaces.IControllerAdmin {
	return &adminController{
		services: services,
	}
}

func (adminController *adminController) AdminLogin(c *gin.Context) {
	var request users.UserLogin
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
	}
	response := adminController.services.AdminLogin(request)
	c.JSON(response.Status, response.Data)
}

func (adminController *adminController) CreateStudent(c *gin.Context) {
	var newUser users.CreateUserRequest
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
	}
	response := adminController.services.CreateStudent(newUser)
	c.JSON(response.Status, response.Data)
}

func (adminController *adminController) GetStudentsDetails(c *gin.Context) {
	response := adminController.services.GetStudentsDetails()
	c.JSON(response.Status, response.Data)
}
