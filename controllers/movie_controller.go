package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/atifali-pm/go-api/database"
    "github.com/atifali-pm/go-api/models"
)

func CreateMovie(c *gin.Context) {
    var movie models.Movie

    if err := c.ShouldBindJSON(&movie); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    database.DB.Create(&movie)
    c.JSON(201, movie)
}