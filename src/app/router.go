package app

import (
	"farhanhilmi/movies_be/src/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(movieController controller.MovieController) *echo.Echo {

	e := echo.New()
	// e.Group("/api")
	e.POST("/movies", movieController.Create)
	e.PUT("/movies/:movieId", movieController.Update)
	e.DELETE("/movies/:movieId", movieController.Delete)
	e.GET("/movies", movieController.FindAll)
	e.GET("/movies/:movieId", movieController.FindById)

	return e
}
