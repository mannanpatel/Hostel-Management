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
	adminRepository interfaces.IRepositoryAdmin = repository.NewRepositoryAdmin()
	adminServices   interfaces.IServicesAdmin   = services.NewServicesAdmin(adminRepository)
	adminController interfaces.IControllerAdmin = controller.NewControllerAdmin(adminServices)
)

func AdminRoutes(router *gin.Engine) {
	router.POST("/admin/login", adminController.AdminLogin)
	// Protected routes - require admin authentication
	adminGroup := router.Group("/admin").Use(middleware.AuthMiddleware())
	{
		adminGroup.POST("/createuser", adminController.CreateStudent)
		adminGroup.GET("/student", adminController.GetStudentsDetails)
		// Add other admin routes here that require authentication
	}

}
