package service

import (
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/pkg/helpers"
	"h8-movies/repository/movie_repository"
	"h8-movies/repository/user_repository"

	"github.com/labstack/echo/v4"
)

type AuthService interface {
	Authentication(next echo.HandlerFunc) echo.HandlerFunc
	Authorization(next echo.HandlerFunc) echo.HandlerFunc
}

type authService struct {
	userRepo  user_repository.UserRepository
	movieRepo movie_repository.MovieRepository
}

func NewAuthService(userRepo user_repository.UserRepository, movieRepo movie_repository.MovieRepository) AuthService {
	return &authService{
		userRepo:  userRepo,
		movieRepo: movieRepo,
	}
}

func (a *authService) Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		user := ctx.Get("userData").(entity.User)

		movieId, err := helpers.GetParamId(ctx, "movieId")

		if err != nil {
			return ctx.JSON(err.Status(), err)

		}

		movie, err := a.movieRepo.GetMovieById(movieId)

		if err != nil {
			return ctx.JSON(err.Status(), err)

		}

		if user.Level == entity.Admin {
			return next(ctx)

		}

		if movie.UserId != user.Id {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to modify the movie data")
			return ctx.JSON(unauthorizedErr.Status(), unauthorizedErr)

		}

		return next(ctx)
	}
}

func (a *authService) Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.Request().Header["Authorization"]

		if len(bearerToken) == 0 {
			unauthorizedError := errs.NewUnauthenticatedError("token required")
			return ctx.JSON(unauthorizedError.Status(), unauthorizedError)
		}

		var user entity.User // User{Id:0, Email: ""}

		err := user.ValidateToken(bearerToken[0])

		if err != nil {
			return ctx.JSON(err.Status(), err)

		}

		result, err := a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			return ctx.JSON(invalidTokenErr.Status(), invalidTokenErr)

		}

		_ = result

		ctx.Set("userData", user)

		return next(ctx)
	}
}
