package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewResponse(ctx echo.Context) IBaseResponse {
	return &BaseResponse{
		Message: fmt.Sprintf("success call api %s", ctx.Path()),
	}
}

func (resp *BaseResponse) Error() *BaseResponse {
	resp.Status = "error"
	return resp
}

func (resp *BaseResponse) Success() *BaseResponse {
	resp.Status = "success"
	return resp
}

func (resp *BaseResponse) SetMessage(message string) *BaseResponse {
	resp.Message = message
	return resp
}

func (resp *BaseResponse) SetData(data interface{}) *BaseResponse {
	resp.Data = data
	return resp
}

func (resp *BaseResponse) SetMetadata(data interface{}) *BaseResponse {
	resp.Metadata = data
	return resp
}

func (resp *BaseResponse) SetStatusCode(statusCode int) *BaseResponse {
	resp.StatusCode = statusCode
	return resp
}

func (resp *BaseResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"errors":["Internal Server Error"]}`))
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(b)
}
