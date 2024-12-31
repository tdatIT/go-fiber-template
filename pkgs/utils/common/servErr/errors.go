package errors

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-service-template/pkgs/utils/common/enum"
	responses "go-service-template/pkgs/utils/common/response"
	"net/http"
	"os"
)

type ServError struct {
	Status               int    `json:"-"`
	InternalErrorMessage string `json:"-"`
	Code                 string `json:"code"`
	Message              string `json:"message"`
}

func (e *ServError) Error() string {
	if e.InternalErrorMessage != "" {
		return e.InternalErrorMessage
	}
	return e.Message
}

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	msg := responses.DefaultError

	var (
		customErr *ServError
		fiberErr  *fiber.Error
	)

	switch {
	//catch fiber error and return it in lower
	case errors.As(err, &fiberErr):
		msg.Status = fiberErr.Code
		msg.Code = fmt.Sprintf("%d", fiberErr.Code)
		msg.Message = fiberErr.Message
	//catch custom error and return it in lower
	case errors.As(err, &customErr):
		msg.Status = customErr.Status
		msg.Code = customErr.Code
		msg.Message = customErr.Message
	//catch validation error and return it in lower
	case errors.As(err, &validator.ValidationErrors{}):
		msg.Status = http.StatusBadRequest
		msg.Code = enum.InvalidArgument
		if os.Getenv("SERVER_MODE") == "prod" {
			msg.Message = "Invalid parameter"
		} else {
			msg.Message = err.Error()
		}
	default:
		msg.Status = http.StatusInternalServerError
		msg.Code = enum.Internal
		msg.Message = "Internal server error"
	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return msg.JSON(ctx)
}
