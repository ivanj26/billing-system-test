package response

type (
	IBaseResponse interface {
		Success() *BaseResponse
		Error() *BaseResponse
		SetMessage(message string) *BaseResponse
		SetData(data interface{}) *BaseResponse
		SetMetadata(data interface{}) *BaseResponse
		SetStatusCode(statusCode int) *BaseResponse
	}

	BaseResponse struct {
		Status     string      `json:"status"`
		StatusCode int         `json:"status_code"`
		Message    string      `json:"message"`
		Metadata   interface{} `json:"metadata"`
		Data       interface{} `json:"data"`
	}
)
