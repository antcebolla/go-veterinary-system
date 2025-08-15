package controllers

import (
	"github.com/antcebolla/web-server/src/database"
	"github.com/antcebolla/web-server/src/models"
	"github.com/gin-gonic/gin"
)

func GetAllSpeciesHandler(c *gin.Context) {
	var Species []models.Species
	database.DB.Model(&models.Species{}).Preload("Breeds").Find(&Species)
	c.JSON(200, Species)
}
