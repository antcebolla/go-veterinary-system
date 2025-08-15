package controllers

import (
	"fmt"

	"github.com/antcebolla/web-server/src/utils"
	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	user, claims, err := utils.GetClerkInfo(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": "something went wrong",
		})
		return
	}


	c.JSON(200, gin.H{
		"message": "hello world",
		"user":    user.ID,
		"claims":  claims.ID,
	})
}
