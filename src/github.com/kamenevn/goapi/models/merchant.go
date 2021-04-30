package models

import (
	"github.com/kamenevn/goapi/config"
	"time"
)

type (
	Merchant struct {
		//gorm.Model
		Id             int `gorm:"AUTO_INCREMENT"`
		PublicId       string
		Ip       	   string
		PrivateKey     string
		UserId    	   int
		IsActive	   bool
		CreatedAt	   time.Time
		UpdatedAt	   time.Time
		AccessRoutes   []MerchantAccess `gorm:"foreignkey:MerchantId"`
	}
)

func (Merchant) TableName() string {
	return "merchants"
}

/*
func GetMerchant(publicId string, ip string, privateKey string) (*Merchant, error) {
	var selectMerchant *Merchant
	err := config.DB.Where("is_active = (?)", true).First(&selectMerchant).Error
	return selectMerchant, err
}
*/

func CheckAccess(publicId string, ip string, routeId int) (Merchant, error) {
	var selectMerchant Merchant

	err := config.DB.Joins("INNER JOIN merchant_access on merchant_access.merchant_id = merchants.id").
		Where("merchants.is_active = ? AND merchants.public_id = ? AND merchants.ip = ? AND merchant_access.router_id = ?", true, publicId, ip, routeId).
		First(&selectMerchant).Error

	return selectMerchant, err
}