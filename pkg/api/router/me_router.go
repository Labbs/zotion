package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/labbs/zotion/pkg/api/middleware"
	"github.com/labbs/zotion/pkg/api/v1/controller"
	"github.com/labbs/zotion/pkg/repository"
	"github.com/labbs/zotion/pkg/service"
)

func NewMeRouter(config *Config, rbacMiddleware fiber.Handler) {
	// Set up the me routes
	config.Logger.Info().Msg("Setting up me routes")

	// initialize the user repository
	ur := repository.NewUserRepository(config.Db)

	// initialize the space repository
	sr := repository.NewSpaceRepository(config.Db)

	// initialize the favorite repository
	fr := repository.NewFavoriteRepository(config.Db)

	// initialize the user service with the database connection
	us := service.NewUserService(ur)
	ss := service.NewSpaceService(sr)
	fs := service.NewFavoriteService(fr)

	c := controller.MeController{
		UserService:     us,
		SpaceService:    ss,
		FavoriteService: fs,
		Logger:          config.Logger,
	}

	v1Me := config.Fiber.Group(ApiV1Path+"/me", middleware.JwtAuthMiddleware(config.Logger, service.NewSessionService(repository.NewSessionRepository(config.Db))), rbacMiddleware)
	v1Me.Get("/profile", c.GetMyProfile)
	v1Me.Get("/favorites", c.GetMyFavorites)
	v1Me.Get("/spaces", c.GetMySpaces)
	v1Me.Post("/favorites/:documentId", c.AddFavorite)
	v1Me.Delete("/favorites/:documentId", c.UnFavorite)
	v1Me.Get("/preferences", c.GetMyPreferences)
	v1Me.Put("/preferences", c.UpdateMyPreferences)
	v1Me.Put("/change-password", c.ChangeMyPassword)
}
