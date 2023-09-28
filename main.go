package main

import (
	"fmt"
	"log"
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

	// Templates
	t := utils.InitTemplate([]string{"./views", "./build/views"})
	utils.TData = utils.NewTemplateData()

	// Connect Database
	db.ConnectDB()
	db.DB.AutoMigrate(&model.Shop{}, &model.Session{})

	// Init Shopify App
	utils.InitShopifyApp()

	//Init Echo
	e := echo.New()

	e.Static("/", "public")
	e.Static("/", "build/public")

	e.Renderer = t
	e.Use(middleware.CheckOAuthBegin)

	auth := e.Group("/api/auth")
	auth.GET("/tokens", handler.MyHandler)
	auth.GET("/callback", handler.MyCallbackHandler)

	app := e.Group("")
	app.Use(middleware.CheckValidAuth)
	app.GET("/", handler.Hello)
	// Template Test
	app.GET("/hello", handler.Hello)
	app.GET("/graph", handler.GraphQLTest)

	// Redirect to auth if query params exists
	// app.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	// Start Server on 1323
	serverAddress := fmt.Sprintf("127.0.0.1:%s", os.Getenv("APP_PORT"))
	e.Logger.Fatal(e.Start(serverAddress))

}
