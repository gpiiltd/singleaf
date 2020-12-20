package models

import "github.com/jinzhu/gorm"

// Apps structure
type Apps struct {
	gorm.Model
	Name     	string `gorm:"column:name;size:255;not null" json:"name"`
	Email		string `gorm:"column:email;size:100;not null;unique" json:"email"`
	Description	string `gorm:"column:description;not null" json:"description"`
	Status   	string `gorm:"column:status;default:'1'" json:"status"`
	Role   		string `gorm:"column:role;default:'0'" json:"role"`
	Photo    	string `gorm:"column:photo;default:'default.jpeg'" json:"photo"`
}

// AppsWrapper use for print data
type AppsWrapper struct {
	gorm.Model
	Name    	string `gorm:"column:name;size:255;not null" json:"name"`
	Email  		string `gorm:"column:email;size:100;not null;unique" json:"email"`
	Description	string `gorm:"column:description;not null" json:"description"`
	Status    	string `gorm:"column:status;default:'1'" json:"status"`
	Role   		string `gorm:"column:role;default:'0'" json:"role"`
	Photo   	string `gorm:"column:photo;default:'default.jpeg'" json:"photo"`
}

// Photo use for update photo 
type Photo struct {
	gorm.Model
	Photo string `gorm:"column:photo;default:'default.jpeg'" json:"photo"`
}
