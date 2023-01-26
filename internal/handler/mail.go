package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/SandMailLogic/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *Handler) SendEmail(c echo.Context) error {
	var email models.EmailRequest
	err := json.NewDecoder(c.Request().Body).Decode(&email)
	if err != nil {
		logrus.Error("Failed to decode json")
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = h.se.SendEmail(c.Request().Context(), email.Email)
	if err != nil {
		if err == models.ErrorSendMail || err == models.ErrorUserDoesntExist {
			return echo.NewHTTPError(404, err.Error())
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return echo.NewHTTPError(http.StatusOK)
}
