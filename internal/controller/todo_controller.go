package controller

import (
	"net/http"
	"strconv"
	"todo-api/internal/entity"
	"todo-api/internal/service"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	svc service.TodoService
}

func NewTodoController(svc service.TodoService) *TodoController {
	return &TodoController{svc: svc}
}

// POST /todos
func (ctrl *TodoController) Create(c echo.Context) error {
    var input entity.Todo
    if err := c.Bind(&input); err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid JSON")
    }

	if input.Job == "" {
        return c.JSON(http.StatusBadRequest, "error The 'job' field is required")
    }

	if input.Status == "" {
        input.Status = "on progress"
    }

    result, err := ctrl.svc.CreateTask(c.Request().Context(), input) 
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusCreated, result)
}

// GET /todos
func (ctrl *TodoController) GetAll(c echo.Context) error {
	todos, err := ctrl.svc.FetchAllTasks(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, todos)
}

// PUT /todos/:id
func (ctrl *TodoController) Update(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))

	type StatusUpdate struct {
        Status string `json:"status"`
    }

    var input StatusUpdate

    if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON")
    }

    if input.Status == "" {
        return c.JSON(http.StatusBadRequest, "error The 'status' field is required")
    }

    err := ctrl.svc.UpdateTaskStatus(c.Request().Context(), uint(id), input.Status)
    if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, "updated")
}

// Delete /todos/:id
func (ctrl *TodoController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var input entity.Todo
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid Request")
	}

	err := ctrl.svc.RemoveTask(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted")
}