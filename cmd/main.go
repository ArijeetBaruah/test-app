package main

import (
	"net/http"

	"github.com/ArijeetBaruah/test-app/app"
	"github.com/ArijeetBaruah/test-app/app/config"
	"github.com/ArijeetBaruah/test-app/app/models"
	"github.com/ArijeetBaruah/test-app/pkg/database"
	"github.com/ArijeetBaruah/test-app/pkg/logger"
	"github.com/ArijeetBaruah/test-app/pkg/session"
	"github.com/ArijeetBaruah/test-app/pkg/templates"
	"github.com/go-zoo/bone"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
)

func main() {
	logger := &logger.RealLogger{}
	logger.Initialise()

	cfg := &config.AppConfig{
		Logger: logger,
	}
	cfg = cfg.ConstructAppConfig()

	db := &database.DatabaseWrapper{}
	dbConn, dbErr := db.Initialise(&cfg.DB)
	if dbErr != nil {
		logger.Panic(dbErr)
		return
	}

	CSRF := csrf.Protect([]byte(cfg.CSRFAuthkey))

	a := app.App{
		Router:    bone.New(),
		Config:    cfg,
		Logger:    logger,
		DB:        dbConn,
		TplParser: &templates.TemplateParser{},
		FlashService: &session.CookieStoreServiceImpl{
			Store:  sessions.NewCookieStore([]byte(cfg.SessionAuthkey)),
			Secure: false,
		},
		CSRF:              CSRF,
		UserService:       &models.UserServiceImpl{DB: dbConn},
		CustomUserService: &models.CustomUserServiceImpl{DB: dbConn},
	}

	a.InitRoute()

	if err := http.ListenAndServe(cfg.Port, a.Router); err != nil {
		panic(err)
	}
}
