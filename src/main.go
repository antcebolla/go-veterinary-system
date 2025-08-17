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

		// Species routes
		speciesRouter := r.Group("/species", middlewares.AuthMiddleware())
		{
			speciesRouter.GET("/", controllers.GetAllSpeciesHandler) // GET ALL SPECIES
		}

		// Veterinarian centers routes
		vetCentersRouter := r.Group("/centers", middlewares.AuthMiddleware())
		{
			vetCentersRouter.GET("/", controllers.GetAllVetCentersHandler)             // GET ALL VETERINARIAN CENTERS
			vetCentersRouter.GET("/:id", controllers.GetVeterinarianCenterByIdHandler) // GET A VETERINARIAN CENTER
			vetCentersRouter.POST("/", controllers.CreateVetCenterHandler)             // CREATE A VETERINARIAN CENTER
			vetCentersRouter.DELETE("/:id", controllers.DeleteCenterHandler)           // DELETE A VETERINARIAN CENTER
			vetCentersRouter.PUT("/:id", controllers.UpdateCenterByIdHandler)          // UPDATE A VETERINARIAN CENTER
		}

		// veterinarians routes
		veterinariansRouter := r.Group("/veterinarians", middlewares.AuthMiddleware())
		{
			veterinariansRouter.GET("/", controllers.GetAllVeterinariansHandler)
		}

		r.Run()
	}

}
