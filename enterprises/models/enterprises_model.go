package models

import "github.com/jinzhu/gorm"

// User structure
type Enterprises struct {
	gorm.Model
	EnterpriseName     string `gorm:"column:enterprisename;size:255;not null" json:"enterprisename"`
	EnterpriseEmail    string `gorm:"column:enterpriseemail;size:100;not null" json:"enterpriseemail"`
	NoTlpn   string `gorm:"column:no_tlpn;default:'nil'" json:"no_tlpn"`
	EnterpriseHash string `gorm:"column:enterprisehash;default:'nil'" json:"enterprisehash"`
	CompanyTag string `gorm:"column:companytag;size:21;not null;unique" json:"companytag"`
	State   string `gorm:"column:state;size:15;default:'nil'" json:"state"`
	LGA   string `gorm:"column:lga;size:15;default:'nil'" json:"lga"`
	Address  string `gorm:"column:address;size:300;default:'nil'" json:"address"`
	Motto  string `gorm:"column:motto;size:300;default:'nil'" json:"motto"`
	Role     string `gorm:"column:role;default:'Primary Enterprise'" json:"role"`
	VerifyID     string `gorm:"column:verifyid;size:255;default:'nil'" json:"verifyid"`
	EnterpriseType     string `gorm:"column:enterprisetype;default:'Private Enterprise'" json:"enterprisetype"`
	Misc   string `gorm:"column:misc;default:'Mixed Enterprise'" json:"misc"`
	Description   string `gorm:"column:description;size:300;default:'Mixed Enterprise'" json:"description"`
	EnterpriseLogo    string `gorm:"column:enterpriselogo;default:'default.jpeg'" json:"enterpriselogo"`
	VerifyInstrument    string `gorm:"column:verifyinstrument;default:'default.jpeg'" json:"verifyinstrument"`
	Status    string `gorm:"column:status;default:'0'" json:"status"`
}

// EnterprisesWrapper use for print enterprises data without password
type EnterprisesWrapper struct {
	gorm.Model
	EnterpriseName     string `gorm:"column:enterprisename;size:255;not null" json:"enterprisename"`
	EnterpriseEmail    string `gorm:"column:enterpriseemail;size:100;not null" json:"enterpriseemail"`
	NoTlpn   string `gorm:"column:no_tlpn;default:'nil'" json:"no_tlpn"`
	EnterpriseHash string `gorm:"column:enterprisehash;default:'nil'" json:"enterprisehash"`
	CompanyTag string `gorm:"column:companytag;default:'nil'" json:"companytag"`
	State   string `gorm:"column:state;size:15;default:'nil'" json:"state"`
	LGA   string `gorm:"column:lga;size:15;default:'nil'" json:"lga"`
	Address  string `gorm:"column:address;size:300;default:'nil'" json:"address"`
	Motto  string `gorm:"column:motto;size:300;default:'nil'" json:"motto"`
	Role     string `gorm:"column:role;default:'Primary Enterprise'" json:"role"`
	VerifyID     string `gorm:"column:verifyid;size:255;default:'nil'" json:"verifyid"`
	EnterpriseType     string `gorm:"column:enterprisetype;default:'Private Enterprise'" json:"enterprisetype"`
	Misc   string `gorm:"column:misc;default:'Mixed Enterprise'" json:"misc"`
	Description   string `gorm:"column:description;size:300;default:'Mixed Enterprise'" json:"description"`
	EnterpriseLogo    string `gorm:"column:enterpriselogo;default:'default.jpeg'" json:"enterpriselogo"`
	VerifyInstrument    string `gorm:"column:verifyinstrument;default:'default.jpeg'" json:"verifyinstrument"`
	Status    string `gorm:"column:status;default:'0'" json:"status"`
}

// type Sessions struct {
// 	gorm.Model
// 	Identity 	string `gorm:"column:identity;size:255;not null" json:"identity"`
// 	Startdate string `gorm:"column:startdate;default:'nil'" json:"startdate"`
// 	Enddate 	string `gorm:"column:enddate;default:'nil'" json:"enddate"`
// 	SessionKey   string `gorm:"column:sessionkey;default:'nil'" json:"sessionkey"`
// 	Status    string `gorm:"column:status;default:'1'" json:"status"`
// }

// type Semester struct {
// 	gorm.Model
// 	Identity 	string `gorm:"column:identity;size:255;not null" json:"identity"`
// 	Startdate string `gorm:"column:startdate;default:'nil'" json:"startdate"`
// 	Enddate 	string `gorm:"column:enddate;default:'nil'" json:"enddate"`
// 	SemesterKey   string `gorm:"column:key;default:'nil'" json:"key"`
// 	Status    string `gorm:"column:status;default:'1'" json:"status"`
// }
// Photo use for update photo enterprises
type EnterpriseLogo struct {
	gorm.Model
	EnterpriseLogo string `gorm:"column:enterpriselogo;default:'default.jpeg'" json:"enterpriselogo"`
}

type VerifyInstrument struct {
	gorm.Model
	VerifyInstrument    string `gorm:"column:verifyinstrument;default:'default.jpeg'" json:"verifyinstrument"`
}
