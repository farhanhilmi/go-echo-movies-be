package helper

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func ReadFromRequestBody(c echo.Context, result interface{}) {
	decoder := json.NewDecoder(c.Request().Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(c echo.Context, response interface{}) {
	c.Response().Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Response().Writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
