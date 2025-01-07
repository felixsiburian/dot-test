package handler

import (
	"dot-test/service"
	"dot-test/service/model"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type UserHandler struct {
	e           *echo.Echo
	userUsecase service.IUserUsecase
}

func NewUserHandler(
	e *echo.Echo,
	userUsecase service.IUserUsecase,
) *UserHandler {
	return &UserHandler{
		e:           e,
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) Create(e echo.Context) error {
	var request model.User

	if err := json.NewDecoder(e.Request().Body).Decode(&request); err != nil {
		err = fmt.Errorf("invalid request")
		return e.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := h.userUsecase.Create(request); err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
