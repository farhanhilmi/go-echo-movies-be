package repository

import (
	"context"
	"database/sql"
	"errors"
	"farhanhilmi/movies_be/src/helper"
	"farhanhilmi/movies_be/src/model/domain"
)

type MovieRepositoryImpl struct{}

func NewMovieRepository() MovieRepository {
	return &MovieRepositoryImpl{}
}

func (repository *MovieRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie {
	SQL := "insert into movies(title, description, type, imdb_url, image_url, watched, season, episode) values (?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, SQL, movie.Title, movie.Description, movie.Type, movie.Imdb_url, movie.Image_url, movie.Watched, movie.Season, movie.Episode)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	movie.Id = int(id)
	return movie
}

func (repository *MovieRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Movie {
	SQL := "select * from movies"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var movies []domain.Movie
	for rows.Next() {
		movie := domain.Movie{}
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.Type, &movie.Imdb_url, &movie.Image_url, &movie.Watched, &movie.Season, &movie.Episode)
		helper.PanicIfError(err)
		movies = append(movies, movie)
	}

	return movies
}

func (repository *MovieRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, movieId int) (domain.Movie, error) {
	SQL := "select * from movies where id = ?"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	movie := domain.Movie{}
	if rows.Next() {
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Description, &movie.Type, &movie.Imdb_url, &movie.Image_url, &movie.Watched, &movie.Season, &movie.Episode)
		helper.PanicIfError(err)
		return movie, nil
	}

	return movie, errors.New("movie not found")
}

func (repository *MovieRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, movie domain.Movie) domain.Movie {
	SQL := "update movies set title = ?, description = ?, type = ?, imdb_url = ?, image_url = ?, watched = ?, season = ?, episode = ? where id = ?"

	_, err := tx.ExecContext(ctx, SQL, movie.Title, movie.Description, movie.Type, movie.Imdb_url, movie.Image_url, movie.Watched, movie.Season, movie.Episode, movie.Id)
	helper.PanicIfError(err)

	return movie
}

func (repository *MovieRepositoryImpl) DeleteById(ctx context.Context, tx *sql.Tx, movieId int) {
	SQL := "delete from movies where id = ?"

	_, err := tx.ExecContext(ctx, SQL, movieId)
	helper.PanicIfError(err)
}
