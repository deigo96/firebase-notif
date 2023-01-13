package changepassword

import (
	"change_password_service/business/changepassword"
	"change_password_service/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) changepassword.ChangepasswordRepo {
	var Repository changepassword.ChangepasswordRepo

	if dbCon.Driver == util.PostgreSQL {
		Repository = NewchangePassword(dbCon.PostgreSQL)

	} else {
		panic("Database driver not supported")
	}

	return Repository
}
