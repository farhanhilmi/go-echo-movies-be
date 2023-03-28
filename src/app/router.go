package app

import (
	"farhanhilmi/movies_be/src/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(movieController controller.MovieController) *echo.Echo {
	e := echo.New()
	e.POST("/movies", movieController.Create)
	e.PUT("/movies/:movieId", movieController.Update)
	e.DELETE("/movies/:movieId", movieController.Delete)
	e.GET("/movies/:movieId", movieController.FindById)
	e.GET("/movies", movieController.FindAll)

	return e
}
