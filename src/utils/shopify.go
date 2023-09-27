package utils

import (
	"errors"
	"fmt"
	"os"

	goshopify "github.com/bold-commerce/go-shopify/v3"
	"github.com/labstack/echo/v4"
	"github.com/nickfthedev/goshopifytest/src/model"
	"github.com/nickfthedev/goshopifytest/src/utils/db"
)

var ShopifyApp goshopify.App

func InitShopifyApp() {
	// Create an app somewhere.
	ShopifyApp = goshopify.App{
		ApiKey:      os.Getenv("SHOPIFY_API_KEY"),
		ApiSecret:   os.Getenv("SHOPIFY_API_SECRET"),
		RedirectUrl: fmt.Sprintf("%s/api/auth/callback", os.Getenv("SHOPIFY_APP_URL")),
		Scope:       os.Getenv("SHOPIFY_API_SCOPES"),
	}
}

func GetAccessToken(c echo.Context) (*model.AccessToken, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return nil, errors.New("No Cookie Token")
	}
	session := new(model.Session)
	db.DB.Where("access_token = ?", cookie.Value).Preload("Shop").First(&session)
	if session.ID == 0 {
		return nil, errors.New("Session not found")
	}
	encryptedToken, dErr := DecryptSessionToken(session.AccessToken)
	if dErr != nil {
		return nil, errors.New("Error decrypting token")
	}
	t := &model.AccessToken{}
	t.Token = encryptedToken
	t.Shopname = session.Shop.Name
	return t, nil
}
