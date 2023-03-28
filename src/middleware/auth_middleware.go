package middleware

import "github.com/labstack/echo/v4"

type AuthMiddleware struct {
	Handler echo.HandlerFunc
}

func NewAuthMiddleware(handler echo.HandlerFunc) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

// func (middleware *AuthMiddleware) ServeHTTP(c echo.Response) error {

// }
