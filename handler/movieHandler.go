package handler

import (
	"h8-movies/dto"
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/pkg/helpers"
	"h8-movies/service"
	"net/http"

	_ "h8-movies/entity"

	"github.com/labstack/echo/v4"
)

type movieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) movieHandler {
	return movieHandler{
		movieService: movieService,
	}
}

// CreateNewMovie godoc
// @Tags movies
// @Description Create New Movie Data
// @ID create-new-movie
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewMovieRequest true "request body json"
// @Success 201 {object} dto.NewMovieRequest
// @Router /movies [post]
func (m movieHandler) CreateMovie(c echo.Context) error {
	var movieRequest dto.NewMovieRequest

	if err := c.Bind(&movieRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		return c.JSON(errBindJson.Status(), errBindJson)

	}

	user := c.Get("userData").(entity.User)

	newMovie, err := m.movieService.CreateMovie(user.Id, movieRequest)

	if err != nil {
		return c.JSON(err.Status(), err)

	}

	return c.JSON(http.StatusCreated, newMovie)
}

func (m movieHandler) UpdateMovieById(c echo.Context) error {
	var movieRequest dto.NewMovieRequest

	if err := c.Bind(&movieRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		return c.JSON(errBindJson.Status(), errBindJson)

	}

	movieId, err := helpers.GetParamId(c, "movieId")

	if err != nil {
		return c.JSON(err.Status(), err)

	}

	response, err := m.movieService.UpdateMovieById(movieId, movieRequest)

	if err != nil {
		return c.JSON(err.Status(), err)

	}

	return c.JSON(response.StatusCode, response)
}
