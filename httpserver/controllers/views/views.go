package views

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

const (
	M_BAD_REQUEST                   = "BAD_REQUEST"
	M_INVALID_CREDENTIALS           = "M_INVALID_CREDENTIALS"
	M_CREATED                       = "CREATED"
	M_OK                            = "OK"
	M_UNAUTHORIZED_ACTION           = "UNAUTHORIZED_ACTION"
	M_INTERNAL_SERVER_ERROR         = "INTERNAL_SERVER_ERROR"
	M_EMAIL_ALREADY_USED            = "EMAIL_ALREADY_USED"
	M_ACCOUNT_SUCCESSFULLY_DELETED  = "Your account has been successfully deleted"
	M_CATEGORY_SUCCESSFULLY_DELETED = "Category has been successfully deleted"
	M_TASK_SUCCESSFULLY_DELETED     = "Task has been successfully deleted"
)

func SuccessResponse(status int, message string, payload interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Payload: payload,
	}
}

func ErrorResponse(status int, message string, error error) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Error:   error.Error(),
	}
}
