package errortype

import (
	"net/http"
)

type ErrorMessage struct {
	StatusCode int
	Message    string
}

func (e *ErrorMessage) Error() string {
	return e.Message
}

var ErrResourceNotFound = ErrorMessage{
	StatusCode: http.StatusNotFound,
	Message:    "resource not found",
}

var ErrRoomDoesNotBelongToDorm = ErrorMessage{
	StatusCode: http.StatusBadRequest,
	Message:    "room does not belong to dorm",
}

var ErrMismatchRoommateRequestType = ErrorMessage{
	StatusCode: http.StatusBadRequest,
	Message:    "student does not have this type of roommate request opened",
}

var ErrOpenRoommateRequest = ErrorMessage{
	StatusCode: http.StatusBadRequest,
	Message:    "student already has an open roommate request",
}

var ErrInvalidInput = ErrorMessage{
	StatusCode: http.StatusNotFound,
	Message:    "invalid input",
}

var ErrUnauthorized = ErrorMessage{
	StatusCode: http.StatusUnauthorized,
	Message:    "unauthorized credentials. wrong username or password",
}

var ErrUserNotFound = ErrorMessage{
	StatusCode: http.StatusNotFound,
	Message:    "user not found",
}

var ErrStorageWrite = ErrorMessage{
	StatusCode: http.StatusBadRequest,
	Message:    "fail to upload file",
}

var ErrNoRoommateRequest = ErrorMessage{
	StatusCode: http.StatusBadRequest,
	Message:    "student has no  opened roommate request",
}
