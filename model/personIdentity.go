package model

type PersonIdentity struct {
	Id               string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	ResidentId       string `gorm:"type:varchar(100);not null;" json:"residentId" binding:"required"`
	DrivingLicensiId string `gorm:"type:varchar(100);not null;" json:"drivingLicensiId" binding:"required"`
	TaxId            string `gorm:"type:varchar(100);not null;" json:"taxId" binding:"required"`
}
