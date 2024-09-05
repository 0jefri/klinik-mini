package model

type Role struct {
	Id   string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	Name string `gorm:"type:varchar(100);not null;" json:"name" binding:"required"`
	Desc string `gorm:"type:varchar(100);not null;" json:"desc" binding:"required"`
}
