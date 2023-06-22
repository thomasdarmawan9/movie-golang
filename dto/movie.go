package dto

import "time"

type NewMovieRequest struct {
	Title    string `json:"title" valid:"required~title cannot be empty" example:"Jelangkung"`
	ImageUrl string `json:"imageUrl" valid:"required~image url cannot be empty" example:"http://imageurl.com"`
	Price    int    `json:"price" valid:"required~price cannot be empty" example:"20000"`
}

type NewMovieResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type MovieResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	ImageUrl  string    `json:"imageUrl"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetMoviesResponse struct {
	Result     string          `json:"result"`
	Message    string          `json:"message"`
	StatusCode int             `json:"statusCode"`
	Data       []MovieResponse `json:"data"`
}
