package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	goshopify "github.com/nickfthedev/go-shopify/v3"
	"github.com/nickfthedev/goshopifytest/src/model"
	"github.com/nickfthedev/goshopifytest/src/utils"
	"github.com/nickfthedev/goshopifytest/src/utils/db"
)

// Create an oauth-authorize url for the app and redirect to it.
// In some request handler, you probably want something like this:
func MyHandler(c echo.Context) error {
	shopName := c.QueryParam("shop")
	state := "nonce"
	authUrl := utils.ShopifyApp.AuthorizeUrl(shopName, state)
	return c.Redirect(301, authUrl)
}

// Fetch a permanent access token in the callback
func MyCallbackHandler(c echo.Context) error {
	// Check that the callback signature is valid
	if ok, _ := utils.ShopifyApp.VerifyAuthorizationURL(c.Request().URL); !ok {
		return c.String(http.StatusInternalServerError, "Invalid Signature!")
	}

	shopName := c.QueryParam("shop")
	code := c.QueryParam("code")
	token, err := utils.ShopifyApp.GetAccessToken(shopName, code)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Create a new API client
	client := goshopify.NewClient(utils.ShopifyApp, shopName, token)
	shop, _ := client.Shop.Get(nil)
	//fmt.Printf("%+v \n \n MyShopifyDomain: %s \n", shop, shop.MyshopifyDomain)

	shopdb := new(model.Shop)
	db.DB.Where("myshopify_domain = ?", shop.MyshopifyDomain).First(&shopdb)

	shopdb.Name = shop.Name
	shopdb.Email = shop.Email
	shopdb.ShopOwner = shop.ShopOwner
	shopdb.Country = shop.Country
	shopdb.Domain = shop.Domain
	shopdb.MyshopifyDomain = shop.MyshopifyDomain

	// Update or Save
	db.DB.Save(&shopdb)

	// Encrypt Token
	encryptedtoken, cryptoError := utils.EncryptSessionToken(token)
	if cryptoError != nil {
		fmt.Println(cryptoError.Error())
		return c.String(500, "Failed to Encrypt Session Token")
	}
	// Create Cookie for Session
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Value = encryptedtoken
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.SameSite = http.SameSiteDefaultMode
	//cookie.Domain = os.Getenv("SHOPIFY_APP_URL")
	c.SetCookie(cookie)

	// Save Session token to DB (encrypted)
	session := new(model.Session)
	db.DB.Where("shop_id = ?", shopdb.ID).First(&session)
	session.Shop = *shopdb

	session.AccessToken = encryptedtoken

	session.IP = c.RealIP()
	db.DB.Save(&session)

	fmt.Println("OAuth completeted. Redirecting...")
	return c.Redirect(301, "/")

}
