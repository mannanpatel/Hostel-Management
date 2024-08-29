package router

import (
	"hst_manag/internal/app/controller"
	"hst_manag/internal/app/interfaces"
	"hst_manag/internal/app/middleware"
	"hst_manag/internal/app/repository"
	"hst_manag/internal/app/services"

	"github.com/gin-gonic/gin"
)

var (
	studentRepository interfaces.IRepositoryStudent = repository.NewRepositoryStudent()
	studentServices   interfaces.IServicesStudent   = services.NewServicesStudent(studentRepository)
	studentController interfaces.IControllerStudent = controller.NewControllerStudent(studentServices)
)

// func Routes() *gin.Engine {
func Routes(router *gin.Engine) {

	// router.POST("/signup", userController.SignUp)
	router.POST("/login", studentController.Login)
	auth := router.Group("/admin").Use(middleware.AuthMiddleware())
	{
		auth.GET("/ping", studentController.Ping)
	}
}
