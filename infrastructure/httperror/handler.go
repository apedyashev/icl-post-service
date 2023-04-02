package httperror

import (
	"icl-posts/infrastructure/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ValidationErrorResponse struct {
	Errors validator.ValidationErrors `json:"errors"`
}

type DefaultErrorResponse struct {
	Message string `json:"message"`
}

func Handler(err error, ctx echo.Context) {
	ctx.Logger().Error(err)

	he, ok := err.(*echo.HTTPError)
	if !ok {
		code := http.StatusInternalServerError
		message := http.StatusText(http.StatusInternalServerError)
		ctx.JSON(code, DefaultErrorResponse{message})

		return
	}

	code := he.Code
	switch code {
	case http.StatusUnprocessableEntity:
		ve := he.Message.(validator.ValidationErrors)
		ctx.JSON(http.StatusUnprocessableEntity, ValidationErrorResponse{
			Errors: ve,
		})
	default:
		message := http.StatusText(code)
		ctx.JSON(code, DefaultErrorResponse{message})
	}
}
