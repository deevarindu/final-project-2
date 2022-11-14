package views

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func BadRequestResponse(err error) *Response {
	return &Response{
		Status:  http.StatusBadRequest,
		Message: "Bad Request",
		Error:   err.Error(),
	}
}

func InternalServerErrorResponse(err error) *Response {
	return &Response{
		Status:  http.StatusInternalServerError,
		Message: "Internal Server Error",
		Error:   err.Error(),
	}
}

func DataNotFoundResponse(err error) *Response {
	return &Response{
		Status:  http.StatusNotFound,
		Message: "Data not found",
		Error:   err.Error(),
	}
}

func SuccessGetResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Payload: payload,
	}
}

func SuccessCreateResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusCreated,
		Message: message,
		Payload: payload,
	}
}

func SuccessUpdateResponse(payload interface{}, message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
		Payload: payload,
	}
}

func SuccessDeleteResponse(message string) *Response {
	return &Response{
		Status:  http.StatusOK,
		Message: message,
	}
}

func DataConflictResponse(err error) *Response {
	return &Response{
		Status:  http.StatusConflict,
		Message: "Duplicate data",
		Error:   err.Error(),
	}
}
