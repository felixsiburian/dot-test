package router

import (
	"dot-test/service"
	"dot-test/service/delivery/handler"
	"github.com/labstack/echo"
)

func NewRouter(
	e *echo.Echo,
	userUsecase service.IUserUsecase,
) {
	u := handler.NewUserHandler(e, userUsecase)

	r := e.Group("/v1")

	r.POST("/user", u.Create)
	r.GET("/user/:id", u.RetrieveById)
	r.PATCH("/user/:id", u.UpdatePassword)
	r.PUT("/user/edit", u.Update)
	r.DELETE("/user/:id", u.Delete)
}
