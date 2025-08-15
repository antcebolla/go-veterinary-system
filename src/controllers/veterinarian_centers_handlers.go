package controllers

import (
	"net/http"

	"github.com/antcebolla/web-server/src/database"
	"github.com/antcebolla/web-server/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllVetCentersHandler(c *gin.Context) {
	var VetCenter []models.VeterinaryCenter
	database.DB.Find(&VetCenter)
	c.JSON(http.StatusOK, VetCenter)
}

func GetVeterinarianCenterByIdHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid veterinary center id",
		})
		return
	}

	center := models.VeterinaryCenter{}
	err := database.DB.First(&center, id).Error
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Veterinary center not found",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get veterinary center, invalid id",
		})
		return
	}

	c.JSON(http.StatusOK, center)
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

func DeleteCenterHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid veterinary center id",
		})
		return
	}

	center := models.VeterinaryCenter{}
	err := database.DB.First(&center, id).Error
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Veterinary center not found",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete veterinary center, invalid id",
		})
		return
	}

	err = database.DB.Delete(&center).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete veterinary center",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Veterinary center deleted successfully",
	})
}
