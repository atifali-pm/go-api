package database

import(
    "fmt"
    "github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    d, err := gorm.Open("mysql", "root:root@/go-api?charset=utf8&parseTime=True&loc=Local")

    if err != nil {
        fmt.Println("Failed to connect to the database!")
        panic(err)
    }

    DB = db

    // Auto migrate the models
    DB.AutoMigrate(&models.Movie{}, &models.Review{})
}