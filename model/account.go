package model

type Account struct {
	Id       string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	PersonId string `gorm:"not null" json:"personId"`
	RoleId   string `gorm:"not null" json:"roleId"`

	Person Person `gorm:"foreignKey:PersonId;references:Id"`
	Role   Role   `gorm:"foreignKey:RoleId;references:Id"`

	Username string `gorm:"type:varchar(100);not null;" json:"username" binding:"required"`
	Email    string `gorm:"type:varchar(100);not null;" json:"email" binding:"required"`
	Pw       string `gorm:"type:varchar(100);not null;" json:"pw" binding:"required"`
	Pin      string `gorm:"type:varchar(100);not null;" json:"pin" binding:"required"`
	Type     int    `gorm:"type:int;not null;" json:"type" binding:"required"`
	Status   int    `gorm:"type:int;not null;" json:"status" binding:"required"`
	Avatar   string `gorm:"type:varchar(100);not null;" json:"avatar" binding:"required"`
	Themes   string `gorm:"type:varchar(100);not null;" json:"themes" binding:"required"`
}
