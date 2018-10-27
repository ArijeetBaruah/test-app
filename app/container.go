package app

import (
	"database/sql"
	"net/http"

	"github.com/ArijeetBaruah/test-app/app/config"
	"github.com/ArijeetBaruah/test-app/app/models"
	"github.com/ArijeetBaruah/test-app/pkg/logger"
	"github.com/ArijeetBaruah/test-app/pkg/session"
	"github.com/ArijeetBaruah/test-app/pkg/templates"
	"github.com/go-zoo/bone"
)

// App wrapper for go application
type App struct {
	Router            *bone.Mux
	Config            config.Config
	Logger            logger.ILogger
	TplParser         templates.ITemplateParser
	DB                *sql.DB
	CSRF              func(http.Handler) http.Handler
	FlashService      session.ISessionService
	UserService       models.UserService
	CustomUserService models.CustomUserService
}
