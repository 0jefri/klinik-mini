package model

type Address struct {
	Id           string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	Country      string `gorm:"type:varchar(100);not null;" json:"country" binding:"required"`
	CountryCode  string `gorm:"type:varchar(100);not null;" json:"countryCode" binding:"required"`
	Province     string `gorm:"type:varchar(100);not null;" json:"province" binding:"required"`
	ProvinceCode string `gorm:"type:varchar(100);not null;" json:"provinceCode" binding:"required"`
	City         string `gorm:"type:varchar(100);not null;" json:"city" binding:"required"`
	CityCode     string `gorm:"type:varchar(100);not null;" json:"cityCode" binding:"required"`
	Regency      string `gorm:"type:varchar(100);not null;" json:"regency" binding:"required"`
	RegencyCode  string `gorm:"type:varchar(100);not null;" json:"regencyCode" binding:"required"`
	District     string `gorm:"type:varchar(100);not null;" json:"district" binding:"required"`
	DistrictCode string `gorm:"type:varchar(100);not null;" json:"districtCode" binding:"required"`
	Street       string `gorm:"type:varchar(100);not null;" json:"street" binding:"required"`
	Description  string `gorm:"type:varchar(100);not null;" json:"description" binding:"required"`
	PostCode     string `gorm:"type:varchar(100);not null;" json:"postCode" binding:"required"`
}
