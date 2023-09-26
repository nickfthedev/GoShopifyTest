package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/nickfthedev/goshopifytest/src/model"
	"github.com/nickfthedev/goshopifytest/src/utils/db"
)

// Checks if a OAuth Process begins
func CheckOAuthBegin(next echo.HandlerFunc) echo.HandlerFunc {
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

// Checks if a OAuth Process begins
func CheckValidAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil {
			return c.String(http.StatusForbidden, "Invalid Cookie!")
		}
		session := new(model.Session)
		db.DB.Where("access_token = ?", cookie).First(&session)
		if session.ID == 0 {
			return c.String(http.StatusForbidden, "Session not found!")
		}
		return next(c)
	}
}
