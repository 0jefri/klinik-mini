package database

import (
	"fmt"
	"klinik-mini/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := "host=localhost user=postgres password=02061996 dbname=klinik_mini port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed Connect to Database:", err)
	}

	fmt.Println("Success Connect to Database!!!")

	db.AutoMigrate(&model.Access{}, &model.Role{}, &model.RoleAccess{}, &model.Address{}, &model.Contact{}, &model.Insurance{}, &model.PersonIdentity{}, &model.Person{})

	// access := []model.Access{
	// 	{Code: "001", Name: "Add", Desc: "Add new patien"},
	// 	{Code: "002", Name: "Show", Desc: "Show list patien"},
	// 	{Code: "003", Name: "Update", Desc: "Update Data patien"},
	// 	{Code: "004", Name: "Delete", Desc: "Delete patien"},
	// }

	// result := db.Create(&access)
	// if result.Error != nil {
	// 	fmt.Println("Error Create Data Access:", result.Error)
	// } else {
	// 	fmt.Println("Success Create Data Access!!!")

	// 	for _, akses := range access {
	// 		fmt.Printf("Code: %s, Name: %s, Desc: %s", akses.Code, akses.Name, akses.Desc)
	// 	}
	// }
}
