package RoleControllers

import (
    "../../models"
    "github.com/gin-gonic/gin"
)

func List (c *gin.Context) {
    db := models.Database();
    var roles = []models.Role{}
    db.Table("Role").Find(&roles)
    c.JSON(200, roles)
}