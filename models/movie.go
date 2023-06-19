package models

import "time"

type Movie struct{
    ID              uint            `gorm:"primary_key" json:"id"`
    Title           string          `gorm:"not null" json:"title"`
    Casts           string          `gorm:"not null" json:"casts"`
    Genre           string          `gorm:"not null" json:"genre"`
    Rating          string          `gorm:"not null" json:"rating"`
    ReleaseDate     time.Time       `gorm:"not null" json:"release_date"`
    CreatedAt       time.Time       `json:"-"`
    UpdatedAt       time.Time       `json:"-"`
}