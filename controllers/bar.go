package controllers

import (
    "github.com/gin-gonic/gin"
)

func Bar (c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "controller bar",
    })
}