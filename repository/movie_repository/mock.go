package movie_repository

import (
	"h8-movies/entity"
	"h8-movies/pkg/errs"
)

var (
	CreateMovie     func(moviePayload *entity.Movie) (*entity.Movie, errs.MessageErr)
	GetMovieById    func(movieId int) (*entity.Movie, errs.MessageErr)
	UpdateMovieById func(payload entity.Movie) errs.MessageErr
	GetMovies       func() ([]*entity.Movie, errs.MessageErr)
)

type movieRepoMock struct{}

func NewMovieRepoMock() MovieRepository {
	return &movieRepoMock{}
}

func (m *movieRepoMock) GetMovies() ([]*entity.Movie, errs.MessageErr) {
	return GetMovies()
}

func (m *movieRepoMock) CreateMovie(moviePayload *entity.Movie) (*entity.Movie, errs.MessageErr) {
	return CreateMovie(moviePayload)
}
func (m *movieRepoMock) GetMovieById(movieId int) (*entity.Movie, errs.MessageErr) {
	return GetMovieById(movieId)
}
func (m *movieRepoMock) UpdateMovieById(payload entity.Movie) errs.MessageErr {
	return UpdateMovieById(payload)
}
