package modules

import (
	"change_password_service/api"
	changepassword "change_password_service/api/v1/changepassword"
	changepasswordService "change_password_service/business/changepassword"
	jwtService "change_password_service/business/jwt"
	"change_password_service/config"

	changepasswordRepo "change_password_service/repository/changepassword"
	"change_password_service/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	changepasswordRepo := changepasswordRepo.RepositoryFactory(dbCon)
	jwtService := jwtService.NewJwtService()

	changepasswordService := changepasswordService.NewChangepasswordService(changepasswordRepo)
	controller := api.Controller{
		Con: changepassword.NewController(changepasswordService, jwtService),
	}
	return controller
}
