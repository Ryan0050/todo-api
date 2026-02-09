package service

import (
	"context"
	"todo-api/internal/entity"
	"todo-api/internal/repository"
)

type TodoServiceImpl struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &TodoServiceImpl{repo: repo}
}

func (s *TodoServiceImpl) CreateTask(ctx context.Context, input entity.Todo) (*entity.Todo, error) {
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

	_ = s.repo.DeleteCache(ctx, "todo_list")
	return task, nil
}

func (s *TodoServiceImpl) FetchAllTasks(ctx context.Context) ([]entity.Todo, error) {
	todos, err := s.repo.GetCache(ctx, "todo_list")
	if err == nil {
		return todos, nil
	}

	todos, err = s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	_ = s.repo.SetCache(ctx, "todo_list", todos)
	return todos, nil
}

func (s *TodoServiceImpl) UpdateTaskStatus(ctx context.Context, id uint, newStatus string) error {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	task.Status = newStatus

	task.Audit.RecType = entity.UPDATE
	task.Audit.UpdatedBy = "ryan"

	if err := s.repo.Update(ctx, task); err != nil {
		return err
	}

	return s.repo.DeleteCache(ctx, "todo_list")
}

func (s *TodoServiceImpl) RemoveTask(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	return s.repo.DeleteCache(ctx, "todo_list")
}
