package service

import (
	"context"
	"database/sql"
	"farhanhilmi/movies_be/src/helper"
	"farhanhilmi/movies_be/src/model/domain"
	"farhanhilmi/movies_be/src/model/web"
	"farhanhilmi/movies_be/src/repository"

	"github.com/go-playground/validator/v10"
)

type MovieServiceImpl struct {
	MovieRepository repository.MovieRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewMovieService(movieRepository repository.MovieRepository, DB *sql.DB, validate *validator.Validate) MovieService {
	return &MovieServiceImpl{movieRepository, DB, validate}
}

func (service *MovieServiceImpl) Create(ctx context.Context, request web.MovieCreateRequest) web.MovieResponse {
	print("request: ", request.Title)
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var movie domain.Movie

	if request.Type == "movie" {
		movie = domain.Movie{Title: request.Title, Description: request.Description, Type: request.Type, Imdb_url: request.Imdb_url, Image_url: request.Image_url, Watched: request.Watched, Season: 0, Episode: 0}
	} else {
		movie = domain.Movie{Title: request.Title, Description: request.Description, Type: request.Type, Imdb_url: request.Imdb_url, Image_url: request.Image_url, Watched: request.Watched, Season: request.Season, Episode: request.Episode}
	}

	movie = service.MovieRepository.Save(ctx, tx, movie)

	return helper.ToMovieResponse(movie)
}

func (service *MovieServiceImpl) Update(ctx context.Context, request web.MovieUpdateRequest) web.MovieResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movie, err := service.MovieRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	if request.Type == "movie" {
		movie.Title = request.Title
		movie.Description = request.Description
		movie.Type = request.Type
		movie.Imdb_url = request.Imdb_url
		movie.Image_url = request.Image_url
		movie.Watched = request.Watched
	} else {
		movie.Title = request.Title
		movie.Description = request.Description
		movie.Type = request.Type
		movie.Imdb_url = request.Imdb_url
		movie.Image_url = request.Image_url
		movie.Watched = request.Watched
		movie.Season = request.Season
		movie.Episode = request.Episode
	}

	movie = service.MovieRepository.Update(ctx, tx, movie)

	return helper.ToMovieResponse(movie)
}

func (service *MovieServiceImpl) Delete(ctx context.Context, movieId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movie, err := service.MovieRepository.FindById(ctx, tx, movieId)
	helper.PanicIfError(err)

	service.MovieRepository.DeleteById(ctx, tx, movie.Id)
}

func (service *MovieServiceImpl) FindById(ctx context.Context, movieId int) web.MovieResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movie, err := service.MovieRepository.FindById(ctx, tx, movieId)
	helper.PanicIfError(err)

	return helper.ToMovieResponse(movie)
}

func (service *MovieServiceImpl) FindAll(ctx context.Context) []web.MovieResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	movies := service.MovieRepository.FindAll(ctx, tx)

	return helper.ToMovieResponses(movies)
}
