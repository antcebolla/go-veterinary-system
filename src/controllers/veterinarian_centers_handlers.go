package controllers

import (
	"net/http"

	"github.com/antcebolla/web-server/src/database"
	"github.com/antcebolla/web-server/src/models"
	"github.com/antcebolla/web-server/src/types"
	"github.com/antcebolla/web-server/src/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllVetCentersHandler(c *gin.Context) {
	offset, limit, current_page, page_size := utils.GetPagination(c)
	var vetCenters []models.VeterinaryCenter
	database.DB.Offset(offset).Limit(limit + 1).Find(&vetCenters)

	has_next_page := len(vetCenters) > page_size
	if has_next_page {
		vetCenters = vetCenters[:page_size]
	}

	c.JSON(http.StatusOK, types.PaginatedResponse[models.VeterinaryCenter]{
		Items:       vetCenters,
		HasNextPage: has_next_page,
		CurrentPage: current_page,
		IsFirstPage: current_page == 1,
	})
}

func GetVeterinarianCenterByIdHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid veterinary center id",
		})
		return
	}

	var center models.VeterinaryCenter
	err := database.DB.First(&center, id).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "center not found",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "error, the id provided is not valid",
			})
		}
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
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "center not found",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "error, the id provided is not valid",
			})
		}
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

func UpdateCenterByIdHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid veterinary center id",
		})
		return
	}

	var centerFromRequest models.VeterinaryCenter
	err := c.ShouldBindJSON(&centerFromRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	var centerFromDB models.VeterinaryCenter
	err = database.DB.First(&centerFromDB, id).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "center not found",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "error, the id provided is not valid",
			})
		}
		return
	}

	centerFromDB.Name = centerFromRequest.Name
	centerFromDB.Address = centerFromRequest.Address
	centerFromDB.Phone = centerFromRequest.Phone

	err = database.DB.Save(&centerFromDB).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update veterinary center, bad request",
		})
		return
	}

	c.JSON(http.StatusOK, centerFromDB)
}
