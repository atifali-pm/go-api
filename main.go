package main

import (
    "github.com/gin-gonic/gin"
    "github.com/atifali-pm/go-api/routes"
)

func main() {
    r := gin.Default()

    routes.InitializeRoutes(r)

    r.Run(":8000")
}