package api

import (
	"change_password_service/api/v1/changepassword"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Con *changepassword.Controller
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {
	e.POST("", controller.Con.Changepassword)
	e.GET("/forgot-password", controller.Con.ForgotPassword)
	e.POST("/check-email", controller.Con.CheckEmail)
	e.GET("/reset-password/:token", controller.Con.TemplateReset)
	e.POST("/update-password/:token", controller.Con.ResetPassword)
}
