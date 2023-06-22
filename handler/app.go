package handler

import (
	"h8-movies/database"
	"h8-movies/repository/movie_repository/movie_pg"
	"h8-movies/repository/user_repository/user_pg"
	"h8-movies/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartApp() {
	var port = "8080"
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	movieRepo := movie_pg.NewMoviePG(db)

	movieService := service.NewMovieService(movieRepo)

	movieHandler := NewMovieHandler(movieService)

	userRepo := user_pg.NewUserPG(db)

	userService := service.NewUserService(userRepo)

	userHandler := NewUserHandler(userService)

	authService := service.NewAuthService(userRepo, movieRepo)

	route := echo.New()

	route.Use(middleware.Logger())

	userRoute := route.Group("/users")
	{
		userRoute.POST("/login", userHandler.Login)
		userRoute.POST("/register", userHandler.Register)
	}

	movieRoute := route.Group("/movies")
	{
		movieRoute.POST("/", movieHandler.CreateMovie, authService.Authentication)

		movieRoute.PUT("/:movieId", movieHandler.UpdateMovieById, authService.Authentication, authService.Authorization)
	}

	route.Start(":" + port)
}
