package request

type CreateTodoRequest struct {
    Job         string `json:"job" validate:"required,min=3,max=100"`
    Description string `json:"description"`
    Status      string `json:"status"`
}

type UpdateStatusRequest struct {
    Status string `json:"status" validate:"required"`
}