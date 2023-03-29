package main

import (
	"farhanhilmi/movies_be/src/app"
	"farhanhilmi/movies_be/src/controller"
	"farhanhilmi/movies_be/src/repository"
	"farhanhilmi/movies_be/src/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	movieRepository := repository.NewMovieRepository()
	movieService := service.NewMovieService(movieRepository, db, validate)
	movieController := controller.NewMovieController(movieService)
	router := app.NewRouter(movieController)
	print("MASOK")
	router.GET("/", func(c echo.Context) error {
		return c.JSON(200, "Hello, World!")
	})

	router.Logger.Fatal(router.Start(":1323"))
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(200, "Hello, World!")
	// })

	// e.Logger.Fatal(e.Start(":1323"))
}
