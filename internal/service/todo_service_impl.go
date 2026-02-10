package service

import (
	"context"
	"encoding/json"
	"time"
	"todo-api/internal/entity"
	"todo-api/internal/model/request"
	"todo-api/internal/model/response"
	"todo-api/internal/repository"

	"github.com/redis/go-redis/v9"
)

type TodoServiceImpl struct {
	repo repository.TodoRepository
	rdb *redis.Client
}

func NewTodoService(repo repository.TodoRepository, rdb *redis.Client) TodoService {
	return &TodoServiceImpl{repo: repo, rdb: rdb}
}

func (s *TodoServiceImpl) CreateTask(ctx context.Context, input *request.CreateTodoRequest) (*response.TodoResponse, error) {
    task := &entity.Todo{
        Job:         input.Job,
        Description: input.Description,
        Status:      input.Status,
        Audit: entity.Audit{
            RecType:   entity.NEW,
            CreatedBy: "Ryan",
        },
    }

    if err := s.repo.Create(ctx, task); err != nil {
        return nil, err
    }

	err := s.rdb.Del(ctx, "todo_list").Err()
    if err != nil {
        return nil, err
    }

    res := &response.TodoResponse{
        ID:          task.ID,
        Job:         task.Job,
        Description: task.Description,
        Status:      task.Status,
        CreatedAt:   task.Audit.CreatedAt.Format("2006-01-02 15:04:05"),
    }
    return res, nil
}

func (s *TodoServiceImpl) FetchAllTasks(ctx context.Context) ([]response.TodoResponse, error) {
	const cacheKey = "todo_list"
    var todos []entity.Todo

    val, err := s.rdb.Get(ctx, cacheKey).Result()
    if err == nil {
        _ = json.Unmarshal([]byte(val), &todos)
    } else {
        todos, err = s.repo.GetAll(ctx)
        if err != nil {
            return nil, err
        }
        
        if len(todos) > 0 {
            data, _ := json.Marshal(todos)
            s.rdb.Set(ctx, cacheKey, data, 5*time.Minute)
        }
    }

	var responseList []response.TodoResponse
	for _, item := range todos {
		responseList = append(responseList, response.TodoResponse{
			ID:          item.ID,
			Job:         item.Job,
			Description: item.Description,
			Status:      item.Status,
			CreatedAt:   item.Audit.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return responseList, nil
}

func (s *TodoServiceImpl) UpdateTaskStatus(ctx context.Context, id uint, newStatus string) error {
    task, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return err
    }

    task.Status = newStatus
    task.Audit.RecType = entity.UPDATE
    task.Audit.UpdatedBy = "Ryan"

    if err := s.repo.Update(ctx, task); err != nil {
        return err
    }

    return s.rdb.Del(ctx, "todo_list").Err()
}

func (s *TodoServiceImpl) RemoveTask(ctx context.Context, id uint) error {
    if err := s.repo.Delete(ctx, id); err != nil {
        return err
    }

    return s.rdb.Del(ctx, "todo_list").Err()
}
