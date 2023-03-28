package controller

import "github.com/labstack/echo/v4"

type MovieController interface {
	Create(ctx echo.Context)
	Update(ctx echo.Context)
	Delete(ctx echo.Context)
	FindById(ctx echo.Context)
	FindAll(ctx echo.Context)
}
