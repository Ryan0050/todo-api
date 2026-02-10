package response

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

type success struct {
	Status     bool        `json:"status"`
	Code       int         `json:"code"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"-"`
}

type listSuccess struct {
	success
	Total int64 `json:"total"`
}

type ErrorResponse struct {
	Status       bool   `json:"status"`
	Code         int    `json:"code"`
	ErrorMessage string `json:"errMessage"`
	StatusCode   int    `json:"-"`
}

func Success(code int) *success {
	return &success{Status: true, Code: code}
}

func HttpSuccess(status int, code int) *success {
	return &success{Status: true, StatusCode: status, Code: code}
}

func (s *success) SetData(data interface{}) *success {
	s.Data = data
	return s
}

func List(code int) *listSuccess {
	var success = &listSuccess{}
	success.Code = code
	success.Status = true
	return success
}

func HttpListSuccess(status int, code int) *listSuccess {
	list := listSuccess{}
	list.Code = code
	list.Status = true
	list.StatusCode = status
	return &list
}

func (s *success) EchoResponse(c echo.Context) error {
	if s.StatusCode > 0 {
		return c.JSON(s.StatusCode, s)
	}
	return c.JSON(200, s)
}

func (s *listSuccess) SetData(data interface{}) *listSuccess {
	s.Data = data
	return s
}

func (s *listSuccess) SetTotal(total int64) *listSuccess {
	s.Total = total
	return s
}

func (s *listSuccess) EchoResponse(c echo.Context) error {
	if s.StatusCode > 0 {
		return c.JSON(s.StatusCode, s)
	}

	return c.JSON(200, s)
}

func NewError(code int, errorMessage string) *ErrorResponse {
	return &ErrorResponse{Status: false, Code: code, ErrorMessage: errorMessage}
}

func HttpError(status int, code int, errorMessage string) *ErrorResponse {
	return &ErrorResponse{Status: false, Code: code, ErrorMessage: errorMessage, StatusCode: status}
}

func (e *ErrorResponse) Error() string {
	jsonBytes, _ := json.Marshal(e)
	return string(jsonBytes)
}

func (e *ErrorResponse) EchoError(c echo.Context) error {
	if e.StatusCode > 0 {
		return c.JSON(e.StatusCode, e)
	}
	return c.JSON(200, e)
}
