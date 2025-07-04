package rbac

import (
	"github.com/labbs/zotion/pkg/models"
	"github.com/rs/zerolog"
)

type Config struct {
	Logger          zerolog.Logger
	UserService     models.UserService
	GroupService    models.GroupService
	SpaceService    models.SpaceService
	DocumentService models.DocumentService
}
