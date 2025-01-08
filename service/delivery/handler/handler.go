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

func (h *UserHandler) RetrieveById(e echo.Context) error {
	var id = e.Param("id")

	userData, err := h.userUsecase.RetrieveById(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    userData,
	})
}

func (h *UserHandler) UpdatePassword(e echo.Context) error {
	var request struct {
		Password string `json:"password"`
	}

	var id = e.Param("id")

	if err := json.NewDecoder(e.Request().Body).Decode(&request); err != nil {
		err = fmt.Errorf("invalid request")
		return e.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := h.userUsecase.UpdatePassword(request.Password, id); err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func (h *UserHandler) Update(e echo.Context) error {
	var request model.User

	if err := json.NewDecoder(e.Request().Body).Decode(&request); err != nil {
		err = fmt.Errorf("invalid request")
		return e.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := h.userUsecase.Update(request); err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
