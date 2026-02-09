package repository

import (
	"context"
	"encoding/json"
	"time"
	"todo-api/internal/entity"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewTodoRepository(db *gorm.DB, rdb *redis.Client) TodoRepository {
	return &TodoRepositoryImpl{db: db, rdb: rdb}
}

// Postgres Methods
func (r *TodoRepositoryImpl) Create(ctx context.Context, todo *entity.Todo) error {
	return r.db.WithContext(ctx).Create(todo).Error
}

func (r *TodoRepositoryImpl) GetAll(ctx context.Context) ([]entity.Todo, error) {
	var todos []entity.Todo
	err := r.db.WithContext(ctx).Find(&todos).Error
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

// Redis Cache Methods
func (r *TodoRepositoryImpl) SetCache(ctx context.Context, key string, todos []entity.Todo) error {
	data, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	return r.rdb.Set(ctx, key, data, 5*time.Minute).Err()
}

func (r *TodoRepositoryImpl) GetCache(ctx context.Context, key string) ([]entity.Todo, error) {
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var todos []entity.Todo
	err = json.Unmarshal([]byte(val), &todos)
	return todos, err
}

func (r *TodoRepositoryImpl) DeleteCache(ctx context.Context, key string) error {
	return r.rdb.Del(ctx, key).Err()
}