package entity

import (
	"h8-movies/dto"
	"time"
)

type Movie struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	ImageUrl  string    `json:"image_url"`
	Price     int       `json:"price"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Movie) EntityToMovieResponseDto() dto.MovieResponse {
	return dto.MovieResponse{
		Id:        m.Id,
		Title:     m.Title,
		ImageUrl:  m.ImageUrl,
		Price:     m.Price,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
