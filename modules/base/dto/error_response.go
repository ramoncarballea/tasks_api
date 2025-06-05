package dto

type ErrorResponse struct {
	Status  int
	Code    string
	Message string
}

func BadRequest(message string) ErrorResponse {
	response := ErrorResponse{
		Status:  400,
		Code:    "BadRequest",
		Message: message,
	}

	return response
}

func ServerError(message string) ErrorResponse {
	response := ErrorResponse{
		Status:  500,
		Code:    "InternalServerError",
		Message: message,
	}

	return response
}

func Forbidden(message string) ErrorResponse {
	response := ErrorResponse{
		Status:  403,
		Code:    "Forbidden",
		Message: message,
	}

	return response
}
