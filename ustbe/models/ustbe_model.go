package models

import "github.com/jinzhu/gorm"

// User structure
type Ustbe struct {
	gorm.Model
	// Fullname     string `gorm:"column:fullname;size:255;not null" json:"fullname"`
	Useremail    string `gorm:"column:useremail;size:100;not null;unique" json:"useremail"`
	UserID   string `gorm:"column:userid;not null" json:"userid"`
	CompanyTag string `gorm:"column:companytag;size:100;not null;" json:"companytag"`
	CompanyID   string `gorm:"column:companyid;size:15;not null" json:"companyid"`
	StartDate     string `gorm:"column:startdate;default:'nil'" json:"startdate"`
	RecordID     string `gorm:"column:recordid;default:'nil';unique" json:"recordid"`
	EndDate     string `gorm:"column:enddate;default:'nil'" json:"enddate"`
	Status  string `gorm:"column:status;default:'1'" json:"status"`
	Description     string `gorm:"column:description;default:'nil'" json:"description"`
	Photo   string `gorm:"column:photo;default:'default.jpeg'" json:"photo"`
}

//UstbeWrapper

type UstbeWrapper struct {
	gorm.Model
	// Fullname     string `gorm:"column:fullname;size:255;not null" json:"fullname"`
	Useremail    string `gorm:"column:useremail;size:100;not null;unique" json:"useremail"`
	UserID   string `gorm:"column:userid;not null" json:"userid"`
	CompanyTag string `gorm:"column:companytag;size:100;not null;" json:"companytag"`
	CompanyID   string `gorm:"column:companyid;size:15;not null" json:"companyid"`
	StartDate     string `gorm:"column:startdate;default:'nil'" json:"startdate"`
	RecordID     string `gorm:"column:recordid;default:'nil';unique" json:"recordid"`
	EndDate     string `gorm:"column:enddate;default:'nil'" json:"enddate"`
	Status  string `gorm:"column:status;default:'1'" json:"status"`
	Description string `gorm:"column:description;default:'nil'" json:"description"`
	Photo   string `gorm:"column:photo;default:'default.jpeg'" json:"photo"`
}

// UstbeWrapper use for print ustbe data without password
type AllUserUstbe struct {
	gorm.Model
	CompanyTag string `gorm:"column:companytag;size:100;not null;" json:"companytag"`
	CompanyID   string `gorm:"column:companyid;size:15;not null" json:"companyid"`
	StartDate     string `gorm:"column:startdate;default:'nil'" json:"startdate"`
	EndDate     string `gorm:"column:enddate;default:'nil'" json:"enddate"`
	RecordID     string `gorm:"column:recordid;default:'nil';unique" json:"recordid"`
	Description     string `gorm:"column:description;default:'nil'" json:"description"`
}

type Photo struct {
	gorm.Model
	Photo string `gorm:"column:photo;default:'default.jpeg'" json:"photo"`
}
