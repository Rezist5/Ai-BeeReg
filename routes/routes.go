package routes

import (
	"github.com/Rezist5/Ai-BeeReg/controllers"
	"github.com/Rezist5/Ai-BeeReg/middleware"
	"github.com/Rezist5/Ai-BeeReg/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.Engine, db *gorm.DB) {
	userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", middleware.CheckRoleMiddleware("ADMIN", "ROOT"), userController.CreateUser)
		userRoutes.POST("/admin", middleware.CheckRoleMiddleware("ROOT"), userController.CreateAdmin)
		userRoutes.GET("/", middleware.CheckRoleMiddleware("ADMIN", "ROOT"), userController.GetAllUsers)
		userRoutes.GET("/:id", middleware.CheckRoleMiddleware("ADMIN", "ROOT"), userController.GetUserByID)
		userRoutes.PUT("/:id", middleware.CheckRoleMiddleware("ADMIN", "ROOT"), userController.UpdateUser)
		userRoutes.DELETE("/:id", middleware.CheckRoleMiddleware("ADMIN", "ROOT"), userController.DeleteUser)
	}
}

func SetupAuthRoutes(router *gin.Engine, db *gorm.DB) {
	authService := services.NewAuthService(db)
	authController := controllers.NewAuthController(authService)

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/logout", authController.Logout)
	}
}

func SetupCompanyRoutes(router *gin.Engine, db *gorm.DB) {
	companyService := services.NewCompanyService(db)
	userService := services.NewUserService(db)

	companyController := &controllers.CompanyController{
		CompanyService: companyService,
		UserService:    userService,
	}

	companyRoutes := router.Group("/companies")
	{
		companyRoutes.POST("/", middleware.CheckRoleMiddleware("ROOT"), companyController.CreateCompany)
		companyRoutes.GET("/", companyController.GetAllCompanies)
		companyRoutes.GET("/:id", companyController.GetCompanyByID)
		companyRoutes.PUT("/:id", middleware.CheckRoleMiddleware("ADMIN"), companyController.UpdateCompany)
		companyRoutes.DELETE("/:id", middleware.CheckRoleMiddleware("ADMIN"), companyController.DeleteCompany)
		companyRoutes.POST("/:id/admin", middleware.CheckRoleMiddleware("ROOT"), companyController.AssignAdmin)
		companyRoutes.DELETE("/:id/admin", middleware.CheckRoleMiddleware("ROOT"), companyController.RemoveAdmin)
	}
}

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	SetupUserRoutes(router, db)
	SetupAuthRoutes(router, db)
	SetupCompanyRoutes(router, db)
}
