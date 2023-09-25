package utils

import (
	"fmt"
	"os"

	goshopify "github.com/bold-commerce/go-shopify/v3"
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
