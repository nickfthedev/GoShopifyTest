package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/nickfthedev/goshopifytest/src/handler"
	"github.com/nickfthedev/goshopifytest/src/middleware"
	"github.com/nickfthedev/goshopifytest/src/model"
	"github.com/nickfthedev/goshopifytest/src/utils"
	"github.com/nickfthedev/goshopifytest/src/utils/db"
)

func main() {

	// Load .env
	godotenv.Load(".env")
	if os.Getenv("APP_SECRET") == "" {
		log.Fatalln("No app secret is set. Set up your .env file")
	}

	// Connect Database
	db.ConnectDB()
	db.DB.AutoMigrate(&model.Shop{}, &model.Session{})

	// Init Shopify App
	utils.InitShopifyApp()

	//Init Echo
	e := echo.New()

	e.Use(middleware.CheckOAuthBegin)
	e.Use(middleware.CheckValidAuth)

	// Redirect to auth if query params exists
	e.GET("/", func(c echo.Context) error {

		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/api/auth/tokens", handler.MyHandler)
	e.GET("/api/auth/callback", handler.MyCallbackHandler)

	// Start Server on 1323
	serverAddress := fmt.Sprintf("127.0.0.1:%s", os.Getenv("APP_PORT"))
	e.Logger.Fatal(e.Start(serverAddress))

}
