package main

import (
	_ "fmt"
	"net/http"

	"github.com/Egor-Tihonov/SandMailLogic/internal/handler"
	"github.com/Egor-Tihonov/SandMailLogic/internal/repository"
	"github.com/Egor-Tihonov/SandMailLogic/internal/service"
	_ "github.com/caarlos0/env"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {

	repo, err := repository.New( /*cfg.PostgresDBURL*/ "postgresql://postgres:123@localhost:5432/booknetwork")
	if err != nil {
		logrus.Fatalf("Connection was failed, %e", err)
	}
	defer repo.Pool.Close()

	srv := service.New(repo)
	h := handler.New(srv)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://labstack.net"},
		AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.POST("/change-pas", h.SendEmail)

	err = e.Start(":5050")
	if err != nil {
		repo.Pool.Close()
		logrus.Fatalf("error started server", err)
	}
}
