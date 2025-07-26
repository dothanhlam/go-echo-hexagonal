package handlers

import (
	"net/http"
	"strconv"

	"go-echo-hexagonal/internal/core/ports"

	"github.com/labstack/echo/v4"
)

// UserHdl implements the http.Handler interface.
type UserHdl struct {
	service ports.UserService
}

// NewUserHdl creates a new UserHandler.
func NewUserHdl(service ports.UserService) *UserHdl {
	return &UserHdl{service: service}
}

// CreateUser handles the creation of a new user.
func (h *UserHdl) CreateUser(c echo.Context) error {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req request
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.service.CreateUser(c.Request().Context(), req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

// GetUser handles the retrieval of a user by ID.
func (h *UserHdl) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := h.service.GetUser(c.Request().Context(), uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

// ListUsers handles the retrieval of a paginated list of users.
func (h *UserHdl) ListUsers(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	users, err := h.service.ListUsers(c.Request().Context(), page, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}
