package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type (
	Router struct {
		//gorm.Model
		Id             int `gorm:"AUTO_INCREMENT"`
		Name           string
		HttpType       string
		Input     	   string
		Output    	   string
		CustomHandler  string `gorm:"nullable"`
		IsActive	   bool
		CreatedAt	   time.Time
		UpdatedAt	   time.Time
		CheckAccess	   bool
		Scheme		   string
		Domain		   string
	}
)

func (Router) TableName() string {
	return "router"
}

func GetRoutes(DB *gorm.DB) ([]Router, error) {
	var selectRoutes []Router
	err := DB.Where("is_active = (?)", true).Find(&selectRoutes).Error
	return selectRoutes, err
}