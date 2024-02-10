package exception

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/arfan21/synapsis_id/pkg/constant"
	"github.com/arfan21/synapsis_id/pkg/pkgutil"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func FiberErrorHandler(ctx *fiber.Ctx, err error) error {

	defaultRes := pkgutil.HTTPResponse{
		Code:    fiber.StatusInternalServerError,
		Message: "Internal Server Error",
	}

	var errValidation *constant.ErrValidation
	if errors.As(err, &errValidation) {
		data := errValidation.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicIfNeeded(errJson)

		defaultRes.Code = fiber.StatusBadRequest
		defaultRes.Message = "Bad Request"
		var errors []interface{}
		for _, message := range messages {
			errors = append(errors, message)
		}
		defaultRes.Errors = errors
	}

	var notFoundError *constant.ErrNotFound
	if errors.As(err, &notFoundError) {
		defaultRes.Code = fiber.StatusNotFound
		defaultRes.Message = "Not Found"
	}

	var unauthorizedError *constant.ErrUnauthorized
	if errors.As(err, &unauthorizedError) {
		defaultRes.Code = fiber.StatusUnauthorized
		defaultRes.Message = "Unauthorized"
	}

	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		defaultRes.Code = fiberError.Code
		defaultRes.Message = fiberError.Message
	}

	if errors.Is(err, pgx.ErrNoRows) {
		defaultRes.Code = fiber.StatusNotFound
		defaultRes.Message = pgx.ErrNoRows.Error()
	}

	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		defaultRes.Code = fiber.StatusUnprocessableEntity
		defaultRes.Message = http.StatusText(fiber.StatusUnprocessableEntity)

		defaultRes.Errors = []interface{}{
			map[string]interface{}{
				"field":   unmarshalTypeError.Field,
				"message": fmt.Sprintf("%s harus %s", unmarshalTypeError.Field, unmarshalTypeError.Type),
			},
		}
	}

	if defaultRes.Code >= 500 {
		defaultRes.Message = http.StatusText(defaultRes.Code)
	}

	return ctx.Status(defaultRes.Code).JSON(defaultRes)
}
