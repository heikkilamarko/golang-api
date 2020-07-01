package utils

type dataResponse struct {
	Meta interface{} `json:"meta,omitempty"`
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Error *errorResponseError `json:"error"`
}

type errorResponseError struct {
	Code    string            `json:"code"`
	Details map[string]string `json:"details,omitempty"`
}

func newDataResponse(data, meta interface{}) *dataResponse {
	return &dataResponse{meta, data}
}

func newErrorResponse(code string, details map[string]string) *errorResponse {
	return &errorResponse{
		Error: &errorResponseError{code, details},
	}
}

func newBadRequestResponse(details map[string]string) *errorResponse {
	return newErrorResponse(ErrCodeBadRequest, details)
}

func newUnauthorizedResponse(details map[string]string) *errorResponse {
	return newErrorResponse(ErrCodeUnauthorized, details)
}

func newNotFoundResponse(details map[string]string) *errorResponse {
	return newErrorResponse(ErrCodeNotFound, details)
}

func newInternalErrorResponse(details map[string]string) *errorResponse {
	return newErrorResponse(ErrCodeInternalError, details)
}
