package model

type Contact struct {
	Id        string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	Telp1     string `gorm:"type:varchar(100);not null;" json:"telp1" binding:"required"`
	Telp2     string `gorm:"type:varchar(100);not null;" json:"telp2" binding:"required"`
	Phone1    string `gorm:"type:varchar(100);not null;" json:"phone1" binding:"required"`
	Phone2    string `gorm:"type:varchar(100);not null;" json:"phone2" binding:"required"`
	Instagram string `gorm:"type:varchar(100);not null;" json:"instagram" binding:"required"`
	Facebook  string `gorm:"type:varchar(100);not null;" json:"facebook" binding:"required"`
	Twitter   string `gorm:"type:varchar(100);not null;" json:"twitter" binding:"required"`
	LinkedIn  string `gorm:"type:varchar(100);not null;" json:"linkedIn" binding:"required"`
	Website   string `gorm:"type:varchar(100);not null;" json:"website" binding:"required"`
}
