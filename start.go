package main

import (
    "github.com/gin-gonic/gin"
    "./controllers"
    "./controllers/Role"
    "./models"
)

func main() {
    models.InitDatabase("mysql",  "root:0000@tcp(localhost:3306)/dev");

    router := gin.Default()

    v1 := router.Group("/api/v1/")
    v1.Use()
    {
        v1.GET("/foo", controllers.Foo)
        v1.GET("/bar", controllers.Bar)
        v1.GET("/roles", RoleControllers.List)
    }

    router.Run(":8888")
}