package model

type RoleAccess struct {
	Id         string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	AccessCode string `gorm:"not null" json:"accessCode"`
	RoleId     string `gorm:"not null" json:"roleId"`

	Access Access `gorm:"foreignKey:AccessCode;references:Code"`
	Role   Role   `gorm:"foreignKey:RoleId;references:Id"`
}
