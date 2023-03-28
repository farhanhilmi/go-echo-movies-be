package service

import (
	"context"
	"farhanhilmi/movies_be/src/model/web"
)

type MovieService interface {
	Create(ctx context.Context, request web.MovieCreateRequest) web.MovieResponse
	Update(ctx context.Context, request web.MovieUpdateRequest) web.MovieResponse
	Delete(ctx context.Context, movieId int)
	FindById(ctx context.Context, movieId int) web.MovieResponse
	FindAll(ctx context.Context) []web.MovieResponse
}
