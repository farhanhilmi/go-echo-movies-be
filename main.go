package main

import (
	"farhanhilmi/movies_be/src/app"
	"farhanhilmi/movies_be/src/controller"
	"farhanhilmi/movies_be/src/repository"
	"farhanhilmi/movies_be/src/service"

	"github.com/go-playground/validator"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	movieRepository := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepository, validate)
	movieController := controller.NewMovieController(movieService)
	router := app.NewRouter(movieController)

	router.Start(":8080")
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(200, "Hello, World!")
	// })

	// e.Logger.Fatal(e.Start(":1323"))
}
