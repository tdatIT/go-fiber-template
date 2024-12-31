package errors

import (
	"go-service-template/pkgs/utils/common/enum"
)

var (
	ErrInternalServer = &ServError{
		Status:  500,
		Code:    enum.Internal,
		Message: "Internal server error",
	}

	ErrBadRequest = &ServError{
		Status:  400,
		Code:    enum.InvalidArgument,
		Message: "Bad request",
	}

	ErrNotChange = &ServError{
		Status:  200,
		Code:    enum.Ok,
		Message: "Not change",
	}

	ErrPermissionDenied = &ServError{
		Status:  403,
		Code:    enum.PermissionDenied,
		Message: "Permission denied",
	}

	ErrNotFound = &ServError{
		Status:  404,
		Code:    enum.NotFound,
		Message: "Not found",
	}

	ErrAlreadyExists = &ServError{
		Status:  409,
		Code:    enum.AlreadyExists,
		Message: "Already exists",
	}

	ErrUnauthenticated = &ServError{
		Status:  401,
		Code:    enum.Unauthenticated,
		Message: "Unauthorized",
	}

	ErrNotFoundRecord = &ServError{
		Status:  404,
		Code:    enum.NotFoundEntity,
		Message: "Record does not exist",
	}

	ErrInvalidParameters = &ServError{
		Status:  400,
		Code:    enum.InvalidArgument,
		Message: "Invalid parameters",
	}

	ErrTooManyRequest = &ServError{
		Status:  429,
		Code:    enum.ResourceExhausted,
		Message: "Too Many Requests",
	}
)
