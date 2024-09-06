package model

import "time"

type Person struct {
	Id               string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	AddressId        string `gorm:"not null" json:"addressId"`
	ContactId        string `gorm:"not null" json:"contactId"`
	PersonIdentityId string `gorm:"not null" json:"personIdentityId"`

	Address        Address        `gorm:"foreignKey:AddressId;references:Id"`
	Contact        Contact        `gorm:"foreignKey:ContactId;references:Id"`
	PersonIdentity PersonIdentity `gorm:"foreignKey:PersonIdentityId;references:Id"`

	Name       string    `gorm:"type:varchar(100);not null;" json:"name" binding:"required"`
	Gender     int       `gorm:"type:int;not null;" json:"gender" binding:"required"`
	BirthDate  time.Time `gorm:"type:date;not null" json:"birthDate"`
	BirthPlace string    `gorm:"type:varchar(100);not null;" json:"birthPlace" binding:"required"`
}
