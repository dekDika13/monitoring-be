package adminController

import (
	"backend/dto/adminDto"
	"backend/service/adminService"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo"
)

type AdminController interface{}

type adminController struct {
	adminServ adminService.AdminService
}

func NewAdminController(adminService adminService.AdminService) *adminController {
	return &adminController{
		adminServ: adminService,
	}
}

// TODO ADMIN CONTROLLER HERE

// TODO REGISTER ADMIN
func (u *adminController) RegisterAdmin(c echo.Context) error {
	var payloads adminDto.RegisterAdminDto

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	res, err := u.adminServ.RegisterAdmin(payloads)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "success",
		Code:    http.StatusOK,
		Data:    res,
	})

}

// TODO LOGIN ADMIN
func (u *adminController) LoginAdmin(c echo.Context) error {
	var payloads adminDto.LoginDTO

	if err := c.Bind(&payloads); err != nil {
		return err
	}

	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	res, err := u.adminServ.LoginAdmin(payloads)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.Response{
			Message: err.Error(),
			Code:    http.StatusUnauthorized,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "login success",
		Code:    http.StatusOK,
		Data:    res,
	})
}