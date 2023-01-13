package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"change_password_service/app/modules"
	commonservice "change_password_service/business/common-service"
	"change_password_service/config"

	"change_password_service/api"

	"change_password_service/util"
)

func main() {

	conf := config.GetConfig()
	dbCon := util.NewConnectionDatabase(conf)

	controllers := modules.RegisterModules(dbCon, conf)
	e := echo.New()
	e.Renderer = commonservice.NewRenderer("templates/*.html", true)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	api.RegisterRoutes(e, &controllers)

	server := config.GetServer()
	e.Logger.Fatal(e.Start(server.Port))
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	defer dbCon.CloseConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
