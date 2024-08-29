package controller

import (
	"hst_manag/internal/app/interfaces"
	"hst_manag/internal/app/models/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type studentController struct {
	services interfaces.IServicesStudent
}

func NewControllerStudent(services interfaces.IServicesStudent) interfaces.IControllerStudent {
	return &studentController{
		services: services,
	}
}

// func (userController *userController) SignUp(c *gin.Context) {
// 	var users users.Users
// 	err := c.ShouldBindJSON(&users)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		c.Abort()
// 	}
// 	fmt.Println("=======111========>>", users)
// 	response := userController.services.SignUp(users)
// 	if response == nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error: response is nil"})
// 		return
// 	}
// 	c.JSON(response.Status, response.Data)

// }

func (studentController *studentController) Login(c *gin.Context) {
	var request users.UserLogin
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
	}
	response := studentController.services.Login(request)
	c.JSON(response.Status, response.Data)
}

func (studentController *studentController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
