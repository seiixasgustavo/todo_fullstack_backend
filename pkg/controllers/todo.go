package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/seiixasgustavo/ToDo-fullstack-backend/pkg/helper"
	"github.com/seiixasgustavo/ToDo-fullstack-backend/pkg/models"
)

type TodoController struct {
	db *gorm.DB
}

func NewTodoController(db *gorm.DB) *TodoController {
	db.AutoMigrate(models.Todo{})
	return &TodoController{db: db}
}

func (t *TodoController) Routes(e *echo.Group) {
	router := e.Group("/todo")
	router.POST("/create", t.Create)
	router.POST("/update/:id", t.Update)
	router.GET("/delete/:id", t.Delete)
	router.GET("/find/:id", t.FindByPk)
	router.GET("/find/user/:id", t.FindByUserId)
	router.GET("/find/user/timestamp/:id", t.FindByUserIdWithTimestamp)
}

func (t *TodoController) Create(c echo.Context) error {
	todo := &models.Todo{}
	if err := c.Bind(&todo); err != nil {
		return helper.WrongBody(c)
	}
	if todo.Text == "" || todo.UserID == 0 {
		return helper.WrongBody(c)
	}
	if err := todo.Create(t.db); err != nil {
		return helper.WrongBody(c)
	}
	return c.JSON(http.StatusOK, helper.TodoIdResponse{Status: true, ID: todo.ID})
}

func (t *TodoController) Delete(c echo.Context) error {
	var tD models.Todo
	todo := c.Param("id")
	if todo == "" {
		return helper.WrongParameters(c)
	}
	todoId, err := strconv.ParseUint(todo, 10, 64)
	if err != nil {
		return helper.WrongParameters(c)
	}
	if tDErr := tD.Delete(t.db, uint(todoId)); tDErr != nil {
		return helper.WrongParameters(c)
	}
	return c.JSON(http.StatusOK, helper.Response{Status: true})
}

func (t *TodoController) FindByPk(c echo.Context) error {
	var todo models.Todo
	todoIdStr := c.Param("id")
	if todoIdStr == "" {
		return c.String(http.StatusBadRequest, "Bad Error")
	}
	todoId, idErr := strconv.ParseUint(todoIdStr, 10, 64)
	if idErr != nil {
		return c.String(http.StatusBadRequest, "Bad Error")
	}
	responseTodo, err := todo.GetByPk(t.db, uint(todoId))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Server Error")
	}
	return c.JSON(http.StatusOK, helper.TodoResponse{
		Status: true,
		Todo:   *responseTodo,
	})
}

func (t *TodoController) Update(c echo.Context) error {
	var todo models.Todo
	id := c.Param("id")
	todoId, todoErr := strconv.ParseUint(id, 10, 64)

	if todoErr != nil {
		return c.String(http.StatusBadRequest, "Id Error")
	}
	if err := todo.Update(t.db, uint(todoId)); err != nil {
		return c.String(http.StatusInternalServerError, "Internal Error")
	}
	return c.JSON(http.StatusOK, helper.Response{Status: true})
}

func (t *TodoController) FindByUserId(c echo.Context) error {
	var todo models.Todo
	id := c.Param("id")
	userId, todoErr := strconv.ParseUint(id, 10, 64)

	if todoErr != nil {
		return c.String(http.StatusBadRequest, "Id Error")
	}
	responseTodo, err := todo.GetByUser(t.db, uint(userId))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Error")
	}
	return c.JSON(http.StatusOK, helper.TodosResponse{Status: true, Todos: *responseTodo})
}

func (t *TodoController) FindByUserIdWithTimestamp(c echo.Context) error {
	var todo models.Todo
	id := c.Param("id")
	userId, todoErr := strconv.ParseUint(id, 10, 64)
	if todoErr != nil {
		return c.String(http.StatusBadRequest, "Id Error")
	}

	responseTodo, err := todo.GetByUser(t.db, uint(userId))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Error")
	}
	res := make(map[string][]models.Todo)

	for _, e := range *responseTodo {
		tInt, _ := strconv.ParseInt(e.Due, 10, 64)
		t := time.Unix(0, tInt*int64(time.Millisecond))
		timestampString := fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
		res[timestampString] = append(res[timestampString], e)
	}
	return c.JSON(http.StatusOK, helper.TodoTimestampResponse{Status: true, Todos: res})
}
