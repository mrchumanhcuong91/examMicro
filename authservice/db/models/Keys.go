package models

import(	"github.com/jinzhu/gorm")

type Key struct{
	gorm.Model 
	Pubkey string
	Encryptkey string
}