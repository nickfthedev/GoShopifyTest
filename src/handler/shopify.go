package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	lib "github.com/nickfthedev/goshopifytest/src/utils"
)

// Create an oauth-authorize url for the app and redirect to it.
// In some request handler, you probably want something like this:
func MyHandler(c echo.Context) error {
	shopName := c.QueryParam("shop")
	fmt.Printf("param: %s \n", shopName)
	state := "nonce"
	authUrl := lib.ShopifyApp.AuthorizeUrl(shopName, state)
	fmt.Printf("AUTHURL: %s \n", authUrl)
	return c.Redirect(301, authUrl)
}

// Fetch a permanent access token in the callback
func MyCallbackHandler(c echo.Context) error {
	// Check that the callback signature is valid
	if ok, _ := lib.ShopifyApp.VerifyAuthorizationURL(c.Request().URL); !ok {
		return c.String(http.StatusInternalServerError, "Invalid Signature!")
	}

	shopName := c.QueryParam("shop")
	code := c.QueryParam("code")
	token, err := lib.ShopifyApp.GetAccessToken(shopName, code)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(token)
	return c.String(http.StatusOK, "OK") // TODO  shpua_f4674ec0357945326dcf7b6d829e3107   Test only remove asap

	// Do something with the token, like store it in a DB.
}
