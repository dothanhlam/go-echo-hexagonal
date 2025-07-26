package handlers

import (
	"net/http"

	"go-echo-hexagonal/internal/core/domain"
	"go-echo-hexagonal/internal/core/ports"

	"github.com/labstack/echo/v4"
)

// AuthHdl handles authentication-related requests.
type AuthHdl struct {
	service ports.UserService
}

// NewAuthHdl creates a new AuthHandler.
func NewAuthHdl(service ports.UserService) *AuthHdl {
	return &AuthHdl{service: service}
}

// Login handles user login.
func (h *AuthHdl) Login(c echo.Context) error {
	var req domain.AuthRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := h.service.Login(c.Request().Context(), req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(http.StatusOK, domain.AuthResponse{Token: token})
}
