package controllers

import (
	"net/http"

	"github.com/antcebolla/web-server/src/database"
	"github.com/antcebolla/web-server/src/models"
	"github.com/gin-gonic/gin"
)

func GetAllVetCentersHandler(c *gin.Context) {
	var VetCenter []models.VeterinaryCenter
	database.DB.Find(&VetCenter)
	c.JSON(200, VetCenter)
}

func CreateVetCenterHandler(c *gin.Context) {
	var VetCenter models.VeterinaryCenter
	err := c.ShouldBindJSON(&VetCenter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	err = database.DB.Create(&VetCenter).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create veterinary center, bad request",
		})
		return
	}
	c.JSON(http.StatusCreated, VetCenter)
}
