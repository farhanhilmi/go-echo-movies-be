package repository

import (
	"context"
	"database/sql"
	"farhanhilmi/movies_be/src/model/domain"
)

type MovieRepository interface {
	Save(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie
	Update(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Movie
	FindById(ctx context.Context, tx *sql.Tx, movieId int) (domain.Movie, error)
	DeleteById(ctx context.Context, tx *sql.Tx, movieId int)
}
