package user

import (
	"context"
	"em/internal/model"
	"em/internal/validator"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	EnrichAndSaveUser(context.Context, *model.User) error
}

type Storage interface {
	GetUserById(context.Context, int64) (*model.User, error)
	DeleteUserById(context.Context, int64) error
	UpdateUser(context.Context, *model.User) error
}

type UserHandler struct {
	userService UserService
	// Чтобы не писать:
	// func (this *UserService) SaveUser(ctx context.Context, user *model.User) error {
	//	return this.userStorage.SaveUser(ctx, user)
	//}
	storage Storage
}

func Register(router fiber.Router, userService UserService, storage Storage) {
	handler := UserHandler{
		userService: userService,
		storage:     storage,
	}

	router.Get("/user/:id", handler.get)
	router.Post("/user", handler.create)
	router.Put("/user", handler.update())
	router.Delete("/user/:id", handler.delete)
}

type HttpError struct {
	Error string `json:"error"`
}

// @Summary	Create user
// @Tags		user
// @Accept		json
// @Produce	json
// @Param		data	body		createReq	true	"User"
// @Success	201		{object}	model.User
// @Failure	400		{object}	HttpError
// @Failure	500		{object}	HttpError
// @Router		/user [post]
func (this *UserHandler) create(ctx *fiber.Ctx) error {
	var body createReq
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	if err := validator.Validate(body); err != nil {
		return err
	}

	user := model.User{
		Name:       body.Name,
		Surname:    body.Surname,
		Patronymic: body.Patronymic,
	}

	if err := this.userService.EnrichAndSaveUser(ctx.Context(), &user); err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).JSON(user)
}

// @Summary	Update user
// @Tags		user
// @Accept		json
// @Produce	json
// @Param		data	body		model.User	true	"User"
// @Success	200		{object}	model.User
// @Failure	400		{object}	HttpError
// @Failure	500		{object}	HttpError
// @Router		/user [put]
func (this *UserHandler) update() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var body model.User
		if err := ctx.BodyParser(&body); err != nil {
			return err
		}

		if err := this.storage.UpdateUser(ctx.Context(), &body); err != nil {
			return err
		}

		return ctx.Status(http.StatusOK).JSON(body)
	}
}

// @Summary	Get user by id
// @Tags		user
// @Accept		json
// @Produce	json
// @Param		id	path		int	true	"User id"
// @Success	200	{object}	model.User
// @Failure	400	{object}	HttpError
// @Failure	500	{object}	HttpError
// @Router		/user/{id} [get]
func (this *UserHandler) get(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	user, err := this.storage.GetUserById(ctx.Context(), int64(id))
	if err != nil {
		return err
	}

	return ctx.Status(http.StatusOK).JSON(user)
}

// @Summary	Delete user by id
// @Tags		user
// @Accept		json
// @Produce	json
// @Param		id	path	int	true	"User id"
// @Success	200
// @Failure	400	{object}	HttpError
// @Failure	500	{object}	HttpError
// @Router		/user/{id} [delete]
func (this *UserHandler) delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	return this.storage.DeleteUserById(ctx.Context(), int64(id))
}
