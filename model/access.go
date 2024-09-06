package model

type Access struct {
	// Id   string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	Code string `gorm:"type:varchar(10);not null;unique;" json:"code" binding:"required"`
	Name string `gorm:"type:varchar(200);not null" json:"name" binding:"required"`
	Desc string `gorm:"type:varchar(200);not null" json:"desc" binding:"required"`
}
