package handlers

import (
	"em/internal/storage"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHanlder(ctx *fiber.Ctx, err error) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	return ctx.Status(getHttpCodeBy(err)).JSON(fiber.Map{"error": err.Error()})
}

func getHttpCodeBy(err error) int {
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return fiberErr.Code
	}

	if _, ok := err.(validator.ValidationErrors); ok {
		return http.StatusBadRequest
	}

	if errors.Is(err, storage.ErrUserNotFound) {
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}
