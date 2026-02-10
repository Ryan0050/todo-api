package repository

import (
	"context"
	"todo-api/internal/entity"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *entity.Todo) error
	GetAll(ctx context.Context) ([]entity.Todo, error)
	GetByID(ctx context.Context, id uint) (*entity.Todo, error)
	Update(ctx context.Context, todo *entity.Todo) error
	Delete(ctx context.Context, id uint) error
}
