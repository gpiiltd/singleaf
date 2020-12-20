package models

import "github.com/jinzhu/gorm"

// User structure
type Besta struct {
	gorm.Model
	Fullname     string `gorm:"column:fullname;size:255;not null" json:"fullname"`
	Useremail    string `gorm:"column:useremail;size:100;not null;unique" json:"useremail"`
	UserID   string `gorm:"column:userid;not null" json:"userid"`
	Service string `gorm:"column:service;size:100;not null;" json:"service"`
	ServiceID   string `gorm:"column:serviceid;size:15;not null" json:"serviceid"`
	StartDate     string `gorm:"column:startdate;default:'nil'" json:"startdate"`
	EndDate     string `gorm:"column:enddate;default:'nil'" json:"enddate"`
	Status  string `gorm:"column:status;default:'1'" json:"status"`
	Description     string `gorm:"column:description;default:'nil'" json:"description"`
	Photo   string `gorm:"column:photo;default:'default.jpeg'" json:"photo"`
}

//BestaWrapper

type BestaWrapper struct {
	gorm.Model
	Fullname     string `gorm:"column:fullname;size:255;not null" json:"fullname"`
	Useremail    string `gorm:"column:useremail;size:100;not null;unique" json:"useremail"`
	UserID   string `gorm:"column:userid;not null" json:"userid"`
	Service string `gorm:"column:service;size:100;not null;" json:"service"`
	ServiceID   string `gorm:"column:serviceid;size:15;not null" json:"serviceid"`
	StartDate     string `gorm:"column:startdate;default:'nil'" json:"startdate"`
	EndDate     string `gorm:"column:enddate;default:'nil'" json:"enddate"`
	Status  string `gorm:"column:status;default:'1'" json:"status"`
	Description     string `gorm:"column:description;default:'nil'" json:"description"`
	Photo   string `gorm:"column:photo;default:'default.jpeg'" json:"photo"`
}

// BestaWrapper use for print besta data without password
type AllUserBesta struct {
	gorm.Model
	Service string `gorm:"column:service;size:100;not null;" json:"service"`
	ServiceID   string `gorm:"column:serviceid;size:15;not null" json:"serviceid"`
	StartDate     string `gorm:"column:startdate;default:'nil'" json:"startdate"`
	EndDate     string `gorm:"column:enddate;default:'nil'" json:"enddate"`
	Description     string `gorm:"column:description;default:'nil'" json:"description"`
}

type Photo struct {
	gorm.Model
	Photo string `gorm:"column:photo;default:'default.jpeg'" json:"photo"`
}
