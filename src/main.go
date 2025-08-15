package main

import (
	"os"

	"github.com/antcebolla/web-server/src/controllers"
	"github.com/antcebolla/web-server/src/database"
	"github.com/antcebolla/web-server/src/initializers"
	"github.com/antcebolla/web-server/src/middlewares"
	"github.com/antcebolla/web-server/src/migration"
	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnviromentalVariables()
	database.ConnectToDB()
	clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))
}

func main() {
	var action string
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	switch action {
	case "migrate":
		migration.MigrateDB()
	case "seed":
		initializers.SeedDB()
	default:
		r := gin.Default()

		r.GET("/", middlewares.AuthMiddleware(), controllers.HomeHandler)

		// Species routes
		species := r.Group("/species", middlewares.AuthMiddleware())
		{
			species.GET("/", controllers.GetAllSpeciesHandler) // GET ALL SPECIES
		}

		// Veterinarian centers routes
		vetCenters := r.Group("/centers", middlewares.AuthMiddleware())
		{
			vetCenters.POST("/", controllers.CreateVetCenterHandler) // CREATE A VETERINARIAN CENTER
		}

		r.Run()
	}

}
