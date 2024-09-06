package model

type Insurance struct {
	Id        string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	AddressId string `gorm:"not null" json:"addressId"`
	ContactId string `gorm:"not null" json:"contactId"`

	Address Address `gorm:"foreignKey:AddressId;references:Id"`
	Contact Contact `gorm:"foreignKey:ContactId;references:Id"`

	Code  string `gorm:"type:varchar(100);not null;" json:"code" binding:"required"`
	Code2 string `gorm:"type:varchar(100);not null;" json:"code2" binding:"required"`
	Code3 string `gorm:"type:varchar(100);not null;" json:"code3" binding:"required"`
	Name  string `gorm:"type:varchar(100);not null;" json:"name" binding:"required"`
	Desc  string `gorm:"type:varchar(100);not null;" json:"desc" binding:"required"`
}
