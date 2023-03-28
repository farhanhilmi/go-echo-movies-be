package controller

import (
	"farhanhilmi/movies_be/src/helper"
	"farhanhilmi/movies_be/src/model/web"
	"farhanhilmi/movies_be/src/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MovieControllerImpl struct {
	MovieService service.MovieService
}

func NewMovieController(movieService service.MovieService) MovieController {
	return &MovieControllerImpl{MovieService: movieService}
}

func (controller *MovieControllerImpl) Create(ctx echo.Context) {
	movieCreateRequest := web.MovieCreateRequest{}
	helper.ReadFromRequestBody(ctx, &movieCreateRequest)

	movieResponse := controller.MovieService.Create(ctx.Request().Context(), movieCreateRequest)
	webReponse := web.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   movieResponse,
	}

	helper.WriteToResponseBody(ctx, webReponse)
}

func (controller *MovieControllerImpl) Update(ctx echo.Context) {
	movieUpdateRequest := web.MovieUpdateRequest{}
	helper.ReadFromRequestBody(ctx, &movieUpdateRequest)

	movieId := ctx.Param("movieId")
	id, err := strconv.Atoi(movieId)
	helper.PanicIfError(err)

	movieUpdateRequest.Id = id

	movieResponse := controller.MovieService.Update(ctx.Request().Context(), movieUpdateRequest)
	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   movieResponse,
	}

	helper.WriteToResponseBody(ctx, webReponse)
}

func (controller *MovieControllerImpl) Delete(ctx echo.Context) {
	movieId := ctx.Param("movieId")
	id, err := strconv.Atoi(movieId)
	helper.PanicIfError(err)

	controller.MovieService.Delete(ctx.Request().Context(), id)
	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(ctx, webReponse)
}

func (controller *MovieControllerImpl) FindById(ctx echo.Context) {
	movieId := ctx.Param("movieId")
	id, err := strconv.Atoi(movieId)
	helper.PanicIfError(err)

	movieResponse := controller.MovieService.FindById(ctx.Request().Context(), id)
	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   movieResponse,
	}

	helper.WriteToResponseBody(ctx, webReponse)
}

func (controller *MovieControllerImpl) FindAll(ctx echo.Context) {
	movieResponses := controller.MovieService.FindAll(ctx.Request().Context())
	webReponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   movieResponses,
	}

	helper.WriteToResponseBody(ctx, webReponse)
}
