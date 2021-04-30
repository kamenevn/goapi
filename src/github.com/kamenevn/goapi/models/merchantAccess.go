package models

type (
	MerchantAccess struct {
		//gorm.Model
		Id         int `gorm:"AUTO_INCREMENT"`
		MerchantId int
		RouterId   int
		Merchant   Merchant `gorm:"foreignkey:MerchantId"`
	}
)

func (MerchantAccess) TableName() string {
	return "merchant_access"
}