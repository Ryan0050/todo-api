package response

type TodoResponse struct {
    ID          uint   `json:"id"`
    Job         string `json:"job"`
    Description string `json:"description"`
    Status      string `json:"status"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}