package model

type PersonInsurance struct {
	Id          string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	PersonId    string `gorm:"not null" json:"personId"`
	InsuranceId string `gorm:"not null" json:"insuranceId"`

	Person    Person    `gorm:"foreignKey:PersonId;references:Id"`
	Insurance Insurance `gorm:"foreignKey:InsuranceId;references:Id"`
}
