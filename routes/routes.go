package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/atifali-pm/go-api/controllers"
)

func InitializeRoutes(router *gin.Engine) {
    api := router.Group("/api")

    movieRoutes := api.Group("/movies")
    {
        movieRoutes.POST("/", controllers.CreateMovie)
    }
}