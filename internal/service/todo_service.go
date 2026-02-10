package service

import (
	"context"
	"todo-api/internal/model/request"
    "todo-api/internal/model/response"
)

type TodoService interface {
	CreateTask(ctx context.Context, input *request.CreateTodoRequest) (*response.TodoResponse, error)
    FetchAllTasks(ctx context.Context) ([]response.TodoResponse, error)
    UpdateTaskStatus(ctx context.Context, id uint, status string) error
	RemoveTask(ctx context.Context, id uint) error
}
