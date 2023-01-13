package util

import (
	"fmt"

	"change_password_service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseDriver string

const (
	PostgreSQL DatabaseDriver = "postgres"

	Static DatabaseDriver = "static"
)

type DatabaseConnection struct {
	Driver DatabaseDriver

	PostgreSQL *gorm.DB

	SqlServer *gorm.DB
}

func NewConnectionDatabase(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection

	switch config.Driver {
	case "PostgreSQL":
		db.Driver = PostgreSQL
		db.PostgreSQL = NewPostgreSQL(config)
	default:
		panic("Database driver not supported")
	}
	return &db
}
func NewPostgreSQL(config *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DB_Host,
		config.DB_User,
		config.DB_Pass,
		config.DB_Name,
		config.DB_Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	return db
}

func (db *DatabaseConnection) CloseConnection() {
	if db.PostgreSQL != nil {
		db, _ := db.PostgreSQL.DB()
		db.Close()
	}
}
