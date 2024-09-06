package model

type AccountAccess struct {
	Id         string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	AccessCode string `gorm:"not null" json:"accessCode"`
	AccountId  string `gorm:"not null" json:"accountId"`

	Access  Access  `gorm:"foreignKey:AccessCode;references:Code"`
	Account Account `gorm:"foreignKey:AccountId;references:Id"`
}
