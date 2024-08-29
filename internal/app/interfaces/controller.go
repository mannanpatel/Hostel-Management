package interfaces

import (
	"github.com/gin-gonic/gin"
)

type IControllerStudent interface {
	// SignUp(c *gin.Context)
	Login(c *gin.Context)
	Ping(c *gin.Context)
}

type IControllerAdmin interface {
	AdminLogin(c *gin.Context)
	CreateStudent(c *gin.Context)
	GetStudentsDetails(c *gin.Context)
}
