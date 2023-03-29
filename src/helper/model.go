package helper

import (
	"farhanhilmi/movies_be/src/model/domain"
	"farhanhilmi/movies_be/src/model/web"
)

func ToMovieResponse(movie domain.Movie) web.MovieResponse {
	return web.MovieResponse{
		Id:          movie.Id,
		Title:       movie.Title,
		Description: movie.Description,
		Type:        movie.Type,
		Imdb_url:    movie.Imdb_url,
		Image_url:   movie.Image_url,
		Watched:     movie.Watched,
		Season:      movie.Season,
		Episode:     movie.Episode,
	}
}

func ToMovieResponses(movies []domain.Movie) []web.MovieResponse {
	var movieResponses []web.MovieResponse
	for _, movie := range movies {
		movieResponses = append(movieResponses, ToMovieResponse(movie))
	}
	return movieResponses
}
