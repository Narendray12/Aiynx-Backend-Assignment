package handler

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"aiynx/internal/logger"
	"aiynx/internal/service"
)

type UserHandler struct {
	svc *service.UserService
	val *validator.Validate
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
		val: validator.New(),
	}
}

/*
========================
GET /users/:id
========================
*/
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}

	user, err := h.svc.GetUser(c.Context(), int32(id))
	if err != nil {
		logger.Log.Error("get user failed", zap.Error(err))
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}

/*
========================
POST /users
========================
*/
type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if err := h.val.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"invalid dob format, expected YYYY-MM-DD",
		)
	}

	user, err := h.svc.CreateUser(
		c.Context(),
		req.Name,
		dob,
	)
	if err != nil {
		logger.Log.Error("create user failed", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

/*
========================
GET /users
========================
*/
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	// basic pagination (optional)
	limit := int32(20)
	offset := int32(0)

	users, err := h.svc.ListUsers(c.Context(), limit, offset)
	if err != nil {
		logger.Log.Error("list users failed", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.JSON(users)
}

/*
========================
PUT /users/:id
========================
*/
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}

	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if err := h.val.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"invalid dob format, expected YYYY-MM-DD",
		)
	}

	user, err := h.svc.UpdateUser(
		c.Context(),
		int32(id),
		req.Name,
		dob,
	)
	if err != nil {
		logger.Log.Error("update user failed", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.JSON(user)
}

/*
========================
DELETE /users/:id
========================
*/
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}

	if err := h.svc.DeleteUser(c.Context(), int32(id)); err != nil {
		logger.Log.Error("delete user failed", zap.Error(err))
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusNoContent)
}
