package middleware

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

// / Checks the current session or query params for renewing token if available
func CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		hmac := c.QueryParam("hmac")
		host := c.QueryParam("host")
		shop := c.QueryParam("shop")
		timestamp := c.QueryParam("timestamp")

		if shop != "" && hmac != "" && host != "" && timestamp != "" && !strings.Contains(c.Request().URL.Path, "/api/auth/tokens") && !strings.Contains(c.Request().URL.Path, "/api/auth/callback") {
			url := fmt.Sprintf("/api/auth/tokens?hmac=%s&host=%s&shop=%s&timestamp=%s", hmac, host, shop, timestamp)
			return c.Redirect(301, url) // Redirect to auth route
		}
		return next(c)
	}
}
