package repository

import (
	"context"
	"todo-api/internal/entity"
	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	db  *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{db: db}
}

func (r *TodoRepositoryImpl) Create(ctx context.Context, todo *entity.Todo) error {
	return r.db.WithContext(ctx).Create(todo).Error
}

func (r *TodoRepositoryImpl) GetAll(ctx context.Context) ([]entity.Todo, error) {
    var todos []entity.Todo
    err := r.db.WithContext(ctx).Order("id ASC").Find(&todos).Error 
    return todos, err
}

func (r *TodoRepositoryImpl) GetByID(ctx context.Context, id uint) (*entity.Todo, error) {
	var todo entity.Todo
	err := r.db.WithContext(ctx).First(&todo, id).Error
	return &todo, err
}

func (r *TodoRepositoryImpl) Update(ctx context.Context, todo *entity.Todo) error {
	return r.db.WithContext(ctx).Save(todo).Error
}

func (r *TodoRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.Todo{}, id).Error
}