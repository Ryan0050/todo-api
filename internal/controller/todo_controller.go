package controller

import (
	"net/http"
	"strconv"
	"todo-api/internal/model/request"
	"todo-api/internal/model/response"
	"todo-api/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type TodoController struct {
	svc service.TodoService
}

func NewTodoController(svc service.TodoService) *TodoController {
	return &TodoController{svc: svc}
}

func (ctrl *TodoController) Routes(e *echo.Echo) {
	todoRoutes := e.Group("/todos")

	todoRoutes.POST("/create", ctrl.Create)
	todoRoutes.GET("/list", ctrl.GetAll)
	todoRoutes.PUT("/update/:id", ctrl.Update)
	todoRoutes.DELETE("/delete/:id", ctrl.Delete)
}

func (ctrl *TodoController) Create(c echo.Context) error {
	var body = new(request.CreateTodoRequest)

	if err := c.Bind(body); err != nil {
		log.Error().Err(err).Msg("")
		return response.ErrBadRequest.EchoError(c)
	} else if err := c.Validate(body); err != nil {
		log.Warn().Err(err).Msg("")
		return response.ErrBadRequest.EchoError(c)
	}

	res, err := ctrl.svc.CreateTask(c.Request().Context(), body)
	if err != nil {
		log.Error().Err(err).Msg("")
		return response.ErrInternalServer.EchoError(c)
	}

	return c.JSON(http.StatusCreated, res)
}
	
func (ctrl *TodoController) GetAll(c echo.Context) error {
	todos, err := ctrl.svc.FetchAllTasks(c.Request().Context())
	if err != nil {
		
		return response.ErrInternalServer.EchoError(c)
	}
	return c.JSON(http.StatusOK, todos)
}

func (ctrl *TodoController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var body = new(request.UpdateStatusRequest)

	if err := c.Bind(body); err != nil {
		log.Error().Err(err).Uint("id", uint(id)).Msg("")
		return response.ErrBadRequest.EchoError(c)
	} else if err := c.Validate(body); err != nil {
		log.Warn().Err(err).Uint("id", uint(id)).Msg("")
		return response.ErrBadRequest.EchoError(c)
	}

	err := ctrl.svc.UpdateTaskStatus(c.Request().Context(), uint(id), body.Status)
	if err != nil {
		log.Error().Err(err).Uint("id", uint(id)).Msg("")
		return response.ErrInternalServer.EchoError(c)
	}

	return c.JSON(http.StatusOK, "updated")
}

func (ctrl *TodoController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := ctrl.svc.RemoveTask(c.Request().Context(), uint(id))
	if err != nil {
		log.Error().Err(err).Uint("id", uint(id)).Msg("")
		return response.ErrInternalServer.EchoError(c)
	}

	return c.JSON(http.StatusOK, "Deleted")
}