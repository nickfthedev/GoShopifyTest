package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	goshopify "github.com/nickfthedev/go-shopify/v3"
	"github.com/nickfthedev/goshopifytest/src/utils"
)

func GraphQLTest(c echo.Context) error {
	// Get Access Token
	t, err := utils.GetAccessToken(c)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return c.String(500, "Token Error")
	}
	fmt.Printf("Token: %s\n Shopname %s\n\n", t.Token, t.Shopname)
	// Endpoint Test
	client := goshopify.NewClient(utils.ShopifyApp, t.Shopname, t.Token)

	req := `query {
		products(first: 10, reverse: true) {
		  edges {
			node {
			  id
			  title
			  handle
			  resourcePublicationOnCurrentPublication {
				publication {
				  name
				  id
				}
				publishDate
				isPublished
			  }
			}
		  }
		}
	  }`

	// settings, err := json.Marshal(map[string]interface{}{
	// 	"accountID": fmt.Sprintf("%d", 12121324),
	// })

	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err.Error())
	// 	return c.String(500, " Error")
	// }

	// variables := map[string]interface{}{
	// 	"webPixel": map[string]interface{}{
	// 		"settings": string(settings),
	// 	},
	// }

	var foo struct {
		id string
	}

	err = client.GraphQL.Query(req, nil, &foo)
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return c.String(500, " Error")

	}
	return c.Render(http.StatusOK, "hello.html", map[string]interface{}{})
}

func Hello(c echo.Context) error {

	// Get Access Token
	t, err := utils.GetAccessToken(c)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return c.String(500, "Token Error")
	}
	fmt.Printf("Token: %s\n", t.Token)
	// Endpoint Test
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
