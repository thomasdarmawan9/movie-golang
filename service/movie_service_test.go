package service

import (
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/repository/movie_repository"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMovieService_GetMovieById_Success(t *testing.T) {
	movieRepo := movie_repository.NewMovieRepoMock()

	movieService := NewMovieService(movieRepo)

	currentTime := time.Now()

	movie := entity.Movie{
		Id:        1,
		Title:     "Test Movie",
		ImageUrl:  "https://test-movie.png",
		Price:     2000,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	movie_repository.GetMovieById = func(movieId int) (*entity.Movie, errs.MessageErr) {
		return &movie, nil
	}

	response, err := movieService.GetMovieById(1)

	assert.Nil(t, err)

	assert.NotNil(t, response)

	assert.Equal(t, "Test Movie", response.Title)

	assert.Equal(t, 1, response.Id)
}

func TestMovieService_GetMovieById_NotFoundError(t *testing.T) {
	movieRepo := movie_repository.NewMovieRepoMock()

	movieService := NewMovieService(movieRepo)

	movie_repository.GetMovieById = func(movieId int) (*entity.Movie, errs.MessageErr) {
		return nil, errs.NewNotFoundError("movie data not found")
	}

	response, err := movieService.GetMovieById(1)

	assert.Nil(t, response)

	assert.NotNil(t, err)

	assert.Equal(t, http.StatusNotFound, err.Status())

	assert.Equal(t, "movie data not found", err.Message())

	assert.Equal(t, "NOT_FOUND", err.Error())
}

func TestMovieService_GetMovies_Success(t *testing.T) {
	movieRepo := movie_repository.NewMovieRepoMock()

	movieService := NewMovieService(movieRepo)

	currentTime := time.Now()

	movies := []*entity.Movie{
		{
			Id:        1,
			Title:     "Test Movie",
			ImageUrl:  "http://test-movie.png",
			Price:     3000,
			CreatedAt: currentTime,
			UpdatedAt: currentTime,
		},
	}

	movie_repository.GetMovies = func() ([]*entity.Movie, errs.MessageErr) {
		return movies, nil
	}

	response, err := movieService.GetMovies()

	assert.Nil(t, err)

	assert.NotNil(t, response)

	assert.Equal(t, 1, len(response.Data))

	assert.Equal(t, "Test Movie", response.Data[0].Title)
}

func TestMovieService_GetMovies_NotFound(t *testing.T) {
	movieRepo := movie_repository.NewMovieRepoMock()

	movieService := NewMovieService(movieRepo)

	movie_repository.GetMovies = func() ([]*entity.Movie, errs.MessageErr) {
		return []*entity.Movie{}, nil
	}

	response, err := movieService.GetMovies()

	assert.Nil(t, err)

	assert.NotNil(t, response)

	assert.Equal(t, 0, len(response.Data))

}
