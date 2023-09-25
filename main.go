package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/nickfthedev/goshopifytest/src/handler"
	"github.com/nickfthedev/goshopifytest/src/utils"
	"github.com/nickfthedev/goshopifytest/src/utils/db"
)

func main() {

	// Load .env
	godotenv.Load(".env")

	// Connect Database
	db.ConnectDB()
	db.DB.AutoMigrate()

	// Init Shopify App
	utils.InitShopifyApp()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		hmac := c.QueryParam("hmac")
		host := c.QueryParam("shop")
		shop := c.QueryParam("shop")
		timestamp := c.QueryParam("timestamp")
		if shop != "" && hmac != "" && host != "" && timestamp != "" {
			url := fmt.Sprintf("api/auth?hmac=%s&host=%s&shop=%s&timestamp=%s", hmac, host, shop, timestamp)
			return c.Redirect(301, url)
		}

		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/api/auth", handler.MyHandler)
	e.GET("/api/auth/callback", handler.MyCallbackHandler)

	// Start Server on 1323
	serverAddress := fmt.Sprintf("127.0.0.1:%s", os.Getenv("APP_PORT"))
	e.Logger.Fatal(e.Start(serverAddress))

}
