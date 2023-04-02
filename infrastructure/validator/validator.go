package validator

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func New() *CustomValidator {
	return &CustomValidator{Validator: validator.New()}
}

var messages map[string]string = map[string]string{
	"required": "This field is required",
	"gte":      "Must be greater than or equal to %s",
	"lte":      "Must be less than or equal to %s",
}

type ValidationErrors map[string]string

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)
		res := ValidationErrors{}
		for _, ve := range errs {
			msg, ok := messages[ve.Tag()]
			if ok {
				res[ve.Field()] = msg
				if len(ve.Param()) > 0 {
					res[ve.Field()] = fmt.Sprintf(msg, ve.Param())
				}
			} else {
				res[ve.Field()] = fmt.Sprintf("%s failed for tag '%s'", ve.Field(), ve.Tag())
			}
		}
		return echo.NewHTTPError(http.StatusUnprocessableEntity, res)
	}
	return nil
}
