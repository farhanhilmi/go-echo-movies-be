package controller

import "github.com/labstack/echo/v4"

type MovieController interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
	FindById(ctx echo.Context) error
	FindAll(ctx echo.Context) error
}
