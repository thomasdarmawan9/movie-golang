package handler

import (
	"h8-movies/dto"
	"h8-movies/pkg/errs"
	"h8-movies/service"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

func (uh *userHandler) Register(ctx echo.Context) error {
	var newUserRequest dto.NewUserRequest

	if err := ctx.Bind(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		return ctx.JSON(errBindJson.Status(), errBindJson)

	}

	result, err := uh.userService.CreateNewUser(newUserRequest)

	if err != nil {
		return ctx.JSON(err.Status(), err)

	}

	return ctx.JSON(result.StatusCode, result)

}

func (uh *userHandler) Login(ctx echo.Context) error {
	var newUserRequest dto.NewUserRequest

	if err := ctx.Bind(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		return ctx.JSON(errBindJson.Status(), errBindJson)

	}

	result, err := uh.userService.Login(newUserRequest)

	if err != nil {

		return ctx.JSON(err.Status(), err)
	}

	return ctx.JSON(result.StatusCode, result)
}
