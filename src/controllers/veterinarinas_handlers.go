package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllVeterinariansHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "All veterinarians",
	})
}