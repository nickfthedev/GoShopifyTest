package model

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Name            string
	Email           string
	ShopOwner       string
	Country         string
	Domain          string
	MyshopifyDomain string `gorm:"unique"`
}

type Session struct {
	gorm.Model
	AccessToken string
	Shop        Shop
	ShopID      uint
	IP          string
}
