package service

import (
	"context"
	"todo-api/internal/entity"
)

type TodoService interface {
	CreateTask(ctx context.Context, input entity.Todo) (*entity.Todo, error)
	FetchAllTasks(ctx context.Context) ([]entity.Todo, error)
	UpdateTaskStatus(ctx context.Context, id uint, status string) error
	RemoveTask(ctx context.Context, id uint) error
}
