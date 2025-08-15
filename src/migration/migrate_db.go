package migration

import (
	"log"

	"github.com/antcebolla/web-server/src/database"
	"github.com/antcebolla/web-server/src/models"
)

func MigrateDB() {
	err := database.DB.AutoMigrate(models.VeterinaryCenter{}, models.Owner{}, models.Animal{}, models.Species{}, models.Breed{}, models.Veterinarian{}, models.Appointment{}, models.BehaviorNote{}, models.HealthNote{})
	if err != nil {
		log.Fatal("Error migrating database:\n", err)
	} else {
		log.Println("Database migrated successfully")
	}
}
