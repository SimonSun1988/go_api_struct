package controllers

import (
    "github.com/gin-gonic/gin"
)

func Foo (c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "controller foo",
    })
}