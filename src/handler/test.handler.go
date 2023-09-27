package handler

import (
	"fmt"
	"net/http"

	goshopify "github.com/bold-commerce/go-shopify/v3"
	"github.com/labstack/echo/v4"
	"github.com/nickfthedev/goshopifytest/src/utils"
)

func Hello(c echo.Context) error {

	// Get Access Token
	t, err := utils.GetAccessToken(c)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return c.String(500, "Token Error")
	}
	fmt.Printf("Token: %s\n", t.Token)
	// GraphQL
	client := goshopify.NewClient(utils.ShopifyApp, t.Shopname, t.Token)

	test, cErr := client.Product.List(nil)
	if cErr != nil {
		fmt.Printf("Error: %v\n", cErr.Error())
		return c.String(500, "Fetch Error")
	}
	fmt.Printf("%+v\n\n", test)

	//Random Test Data
	type myStruct2 struct {
		Val2 int
	}
	data2 := myStruct2{42}

	// Render Template
	return c.Render(http.StatusOK, "hello.html", map[string]interface{}{
		"data2": data2,
		"world": "Hello!",
	})
}
