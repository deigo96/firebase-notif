package changepassword

import (
	"change_password_service/api/common/obj"
	_response "change_password_service/api/common/response"
	"change_password_service/api/v1/changepassword/request"
	"log"
	"os"
	"strings"

	service "change_password_service/business/changepassword"
	jwtService "change_password_service/business/jwt"
	"fmt"
	"strconv"

	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var Serverurl = os.Getenv("URLSERVER")

type Controller struct {
	Service    service.ChangepasswordService
	jwtService jwtService.JwtService
}

func NewController(
	Service service.ChangepasswordService,
	jwtService jwtService.JwtService,

) *Controller {
	return &Controller{
		Service:    Service,
		jwtService: jwtService,
	}
}
func (controller *Controller) Changepassword(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	validate := controller.jwtService.ValidateToken(header)
	if header == "" {
		response := _response.BuildErrorResponse("Failed to validate token", obj.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, response)
	}
	if validate != true {
		response := _response.BuildErrorResponse("Token expired or invalid", obj.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, response)
	}
	tokenString := strings.Replace(header, "Bearer ", "", -1)
	token, _ := jwt.Parse(tokenString, nil)
	if token == nil {
		return nil
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	intID, _ := strconv.Atoi(id)

	payload := new(request.Req)
	err := c.Bind(payload)
	if err != nil {
		response := _response.BuildErrorResponse("Invalid request body", obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	res, err := controller.Service.ChangePassword(intID, payload.OldPass, payload.NewPass)
	fmt.Println(err, "err handler")
	if err != nil {
		response := _response.BuildErrorResponse("There was a problem with your update.", obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	_ = res
	// data := resp.FromService(res)

	_response := _response.BuildSuccsessResponse("Succsess change password", true, obj.EmptyObj{})
	return c.JSON(http.StatusOK, _response)
}

func (controller *Controller) ForgotPassword(c echo.Context) error {
	data := map[string]string{
		"url": GetUrl(),
	}
	return c.Render(http.StatusOK, "forgot_password.html", data)
}

func (controller *Controller) CheckEmail(c echo.Context) error {
	req := new(request.ReqEmail)
	if err := c.Bind(&req); err != nil {
		response := _response.BuildErrorResponse("Invalid request body", obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	if req.Email == "" {
		response := _response.BuildErrorResponse("Email is required", obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	_, err := controller.Service.CheckEmailService(req.Email)
	if err != nil {
		response := _response.BuildErrorResponse(err.Error(), obj.EmptyObj{})
		return c.JSON(http.StatusOK, response)
	}

	_response := _response.BuildSuccsessResponse("Success send email reset password", true, obj.EmptyObj{})
	return c.JSON(http.StatusOK, _response)
}

func (controller *Controller) TemplateReset(c echo.Context) error {
	param := c.Param("token")
	if param == "" {
		data := map[string]string{
			"message": "Sorry, the url you are looking for doesn't exist",
		}
		return c.Render(http.StatusOK, "invalid_url.html", data)
	}

	_, err := controller.Service.CheckVerification(param)
	if err != nil {
		data := map[string]string{
			"message": "Sorry, the url you are looking for doesn't exist",
		}
		return c.Render(http.StatusOK, "invalid_url.html", data)
	}

	data := map[string]string{
		"token": param,
		"url":   GetUrl(),
	}
	return c.Render(http.StatusOK, "reset_password.html", data)
}

func (controller *Controller) ResetPassword(c echo.Context) error {
	param := c.Param("token")
	if param == "" {
		data := map[string]string{
			"message": "Sorry, the url you are looking for doesn't exist",
		}
		return c.Render(http.StatusOK, "invalid_url.html", data)
	}

	check, err := controller.Service.CheckVerification(param)
	if err != nil {
		data := map[string]string{
			"message": "Sorry, the url you are looking for doesn't exist",
		}
		return c.Render(http.StatusOK, "invalid_url.html", data)
	}

	req := new(request.ReqResetPassword)
	if err := c.Bind(&req); err != nil {
		data := map[string]string{
			"message": "Sorry, the url you are looking for doesn't exist",
		}
		return c.Render(http.StatusOK, "invalid_url.html", data)
	}

	if req.Password == "" || req.PasswordConfirm == "" {
		response := _response.BuildErrorResponse("Password and password confirm are required", obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	if req.Password != req.PasswordConfirm {
		response := _response.BuildErrorResponse("The password and confirm password fields do not match", obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	errReset := controller.Service.ResetPasswordService(check.Id, req.Password)
	if errReset != nil {
		response := _response.BuildErrorResponse(errReset.Error(), obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	_response := _response.BuildSuccsessResponse("Password reset complete. You can now log in with your new password", true, obj.EmptyObj{})
	return c.JSON(http.StatusOK, _response)
}

func GetUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("URLSERVER")
	return url
}
