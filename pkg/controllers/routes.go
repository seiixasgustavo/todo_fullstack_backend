package controllers

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/seiixasgustavo/ToDo-fullstack-backend/pkg/config"
	"github.com/seiixasgustavo/ToDo-fullstack-backend/pkg/controllers/middlewares"
)

func routes(db *gorm.DB, e *echo.Echo, tokenSecret string) {
	// Routes Todo
	router := e.Group("/secure")
	router.Use(middleware.JWT([]byte(tokenSecret)))

	td := NewTodoController(db)
	td.Routes(router)

	// Routes User
	usr := NewUserController(db)
	usr.Routes(router)
	usr.RoutesWithoutSecurity(e)
}

func Router(db *gorm.DB) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	middlewares.SetLogger(e)

	routes(db, e, config.Cfg.Secret)

	return e
}
