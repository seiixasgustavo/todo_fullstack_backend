package controllers

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/seiixasgustavo/ToDo-fullstack-backend/pkg/controllers/middlewares"
	"github.com/seiixasgustavo/ToDo-fullstack-backend/pkg/helper"
	"github.com/seiixasgustavo/ToDo-fullstack-backend/pkg/models"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	db.AutoMigrate(models.User{})
	return &UserController{db: db}
}

func (u *UserController) Routes(e *echo.Group) {
	router := e.Group("/user")
	router.POST("/update", u.Update)
	router.GET("/delete/:id", u.Delete)
	router.GET("/find/:id", u.FindByPk)
	router.POST("/find/name", u.FindByName)
}

func (u *UserController) RoutesWithoutSecurity(e *echo.Echo) {
	router := e.Group("/security")
	router.POST("/login", u.Login)
	router.POST("/signup", u.Create)
}

func (u *UserController) Login(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "Body Error")
	}
	if user.Email == "" {
		return c.String(http.StatusBadRequest, "Body Error")
	}
	if login := user.Login(u.db); !login {
		return c.String(http.StatusBadRequest, "Failed to Login")
	}

	token, err := middlewares.GenerateToken(user.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Token Generation Error")
	}

	user.Password = ""

	response := helper.TokenUserResponse{
		Status: true,
		Token:  token,
		User:   user,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *UserController) Create(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "Body Error")
	}
	if user.Name == "" && user.Email == "" {
		return c.String(http.StatusBadRequest, "Body Error")
	}
	if err := user.Create(u.db); err != nil {
		return c.String(http.StatusInternalServerError, "Server Error")
	}
	response := helper.Response{Status: true}
	return c.JSON(http.StatusOK, response)
}

func (u *UserController) Update(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "Body Error")
	}
	if user.Name == "" && user.Password == "" {
		return c.String(http.StatusBadRequest, "Body Error")
	}
	if err := user.Delete(u.db, user.ID); err != nil {
		return c.String(http.StatusInternalServerError, "Server Error")
	}
	response := struct{ Status bool }{Status: true}
	return c.JSON(http.StatusOK, response)
}

func (u *UserController) Delete(c echo.Context) error {
	var user models.User
	idString := c.Param("id")
	if idString == "" {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	id, parseErr := strconv.ParseUint(idString, 10, 64)
	if parseErr != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err := user.Delete(u.db, uint(id)); err != nil {
		return c.String(http.StatusInternalServerError, "Server Error")
	}
	response := struct{ Status bool }{Status: true}
	return c.JSON(http.StatusOK, response)
}

func (u *UserController) FindByPk(c echo.Context) error {
	var user models.User
	idString := c.Param("id")
	if idString == "" {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	id, parseErr := strconv.ParseUint(idString, 10, 64)
	if parseErr != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	usr, err := user.FindByPk(u.db, uint(id))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Server Error")
	}
	response := helper.UserResponse{
		Status: true,
		User:   *usr,
	}
	return c.JSON(http.StatusOK, response)
}

func (u *UserController) FindByName(c echo.Context) error {
	var user models.User
	var name string
	if err := c.Bind(name); err != nil {
		return c.String(http.StatusBadRequest, "Body Error")
	}
	if name == "" {
		return c.String(http.StatusBadRequest, "Body Error")
	}
	usr, err := user.FindByName(u.db, name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Server Error")
	}
	response := helper.UserResponse{
		Status: true,
		User:   *usr,
	}
	return c.JSON(http.StatusOK, response)
}
