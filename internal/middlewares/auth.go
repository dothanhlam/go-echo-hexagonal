package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Auth is a middleware for authentication and authorization.
func Auth(jwtSecret string, roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
				}
				return []byte(jwtSecret), nil
			})

			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
			}

			// RBAC check
			if len(roles) > 0 {
				userRole, ok := claims["role"].(string)
				if !ok {
					return echo.NewHTTPError(http.StatusForbidden, "Invalid role in token")
				}

				authorized := false
				for _, role := range roles {
					if userRole == role {
						authorized = true
						break
					}
				}

				if !authorized {
					return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to access this resource")
				}
			}

			c.Set("user", claims)

			return next(c)
		}
	}
}
