package helper

import "github.com/seiixasgustavo/ToDo-fullstack-backend/pkg/models"

type Response struct {
	Status bool
}

type TodoResponse struct {
	Status bool
	Todo   models.Todo
}

type TodoIdResponse struct {
	Status bool
	ID     uint
}

type TodosResponse struct {
	Status bool
	Todos  []models.Todo
}

type TokenResponse struct {
	Status bool
	Token  string
}
type TokenUserResponse struct {
	Status bool
	Token  string
	User   models.User
}

type UserResponse struct {
	Status bool
	User   models.User
}
