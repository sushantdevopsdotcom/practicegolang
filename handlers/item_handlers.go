package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "big-web-app/config"
    "big-web-app/models"
)

func GetItems(c *gin.Context) {
    var items []models.Item
    config.DB.Find(&items)
    c.JSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
    var item models.Item
    if err := c.ShouldBindJSON(&item); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&item)
    c.JSON(http.StatusOK, item)
}

