package router

import (
	"backend/controller/adminController"
	m "backend/middleware"
	"backend/repository/adminRepository"
	"backend/service/adminService"
	"backend/utils"
	"os"
	"github.com/labstack/echo/middleware"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	m.LogMiddleware(e)
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	// TODO REPOSITORY
	adminRepository := adminRepository.NewAdminRepository(db)
	

	// TODO SERVICE
	adminService := adminService.NewAdminService(adminRepository)
	

	// TODO CONTROLLER
	adminController := adminController.NewAdminController(adminService)
	

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "ping")
	})

	// TODO ADMIN ROUTE

	v1 := e.Group("/v1")
	v1.Use(middleware.JWT([]byte(os.Getenv("JWT_KEY"))))

	v2 := e.Group("/v2")
	v2.Use(middleware.JWT([]byte(os.Getenv("JWT_KEY"))))

	
	// // TODO AUTH ADMIN
	e.POST("/auth/login", adminController.LoginAdmin)

	// TODO REGISTER
	e.POST("/register", adminController.RegisterAdmin)

	

	// TODO ROLES

	//v1_roles := v1.Group("/role")
	e.POST("/role", adminController.CreateRoles)

	
}
