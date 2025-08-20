package controllers

import (
	"fmt"
	"net/http"

	"github.com/antcebolla/web-server/src/database"
	"github.com/antcebolla/web-server/src/models"
	"github.com/antcebolla/web-server/src/types"
	"github.com/antcebolla/web-server/src/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllVeterinariansHandler(c *gin.Context) {
	limit, offset, current_page, page_size := utils.GetPagination(c)

	var vets []models.Veterinarian
	database.DB.Offset(offset).Limit(limit + 1).Find(&vets)

	has_next_page := len(vets) > page_size
	if has_next_page {
		vets = vets[:page_size]
	}

	c.JSON(http.StatusOK, types.PaginatedResponse[models.Veterinarian]{
		Items:       vets,
		HasNextPage: has_next_page,
		CurrentPage: current_page,
		IsFirstPage: current_page == 1,
	})
}

func GetVeterinarianByIdHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid veterinarian id",
		})
		return
	}

	var vet models.Veterinarian
	err := database.DB.First(&vet, id).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "veterinarian not found",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "error, the id provided is not valid",
			})
		}
		return
	}

	c.JSON(http.StatusOK, vet)
}

func CreateVeterinarianHandler(c *gin.Context) {
	var vet models.Veterinarian
	err := c.ShouldBindJSON(&vet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data, bad request",
		})
		return
	}

	err = database.DB.Create(&vet).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while creating veterinarian, bad request",
		})
		return
	}

	c.JSON(http.StatusCreated, vet)
}

func DeleteVeterinarianByIdHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid veterinarian id",
		})
		return
	}

	var vet models.Veterinarian
	err := database.DB.First(&vet, id).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "veterinarian not found",
			})
		default:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "error, the id provided is not valid",
			})
		}
		return
	}

	err = database.DB.Delete(&vet).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while deleting veterinarian",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "veterinarian deleted succesfully",
	})
}

func UpdateVeterinarianByIdHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invalid veterinarian id",
		})
		return
	}

	var vetFromRequest models.Veterinarian
	err := c.ShouldBindBodyWithJSON(&vetFromRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	var vetFromDB models.Veterinarian
	err = database.DB.First(&vetFromDB, id).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "veterinarian not found",
			})
		default:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "error, the id provided is not valid",
			})
		}
		return
	}

	vetFromDB.Name = vetFromRequest.Name
	vetFromDB.Phone = vetFromRequest.Phone
	vetFromDB.VeterinaryCenterID = vetFromRequest.VeterinaryCenterID
	vetFromDB.Email = vetFromRequest.Email

	err = database.DB.Save(&vetFromDB).Error
	if err != nil {
		fmt.Printf("%v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while updating the veterinarian",
		})
		return
	}

	c.JSON(http.StatusOK, vetFromDB)
}
