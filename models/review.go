package models

import "time"

type Review struct {
    ID              uint            `gorm:"primary_key" json:"id"`
    MovieID         uint            `gorm:"not null" json:"-"`
    Author          string          `gorm:"not null" json:"author"`
    ReviewLikes     int             `gorm:"not null" json:"review_likes"`
    ReviewDisLikes  int             `gorm:"not null" json:"review_dislikes"`
    CreatedAt       time.Time       `json:"-"`
    UpdatedAt       time.Time       `json:"-"`
}