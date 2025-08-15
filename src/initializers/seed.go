package initializers

import (
	"fmt"

	"github.com/antcebolla/web-server/src/database"
	"github.com/antcebolla/web-server/src/models"
)

var speciesData = []map[string]interface{}{
	{
		"species": "Canine",
		"breeds":  []string{"Labrador", "German Shepherd", "Bulldog"},
	},
	{
		"species": "Feline",
		"breeds":  []string{"Siamese", "Persian", "Maine Coon"},
	},
	{
		"species": "Porcine",
		"breeds":  []string{"Duroc", "Hampshire", "Yorkshire"},
	},
	{
		"species": "Equine",
		"breeds":  []string{"Thoroughbred", "Quarter Horse", "Arabian"},
	},
	{
		"species": "Bovine",
		"breeds":  []string{"Angus", "Holstein", "Charolais"},
	},
	{
		"species": "Ovine",
		"breeds":  []string{"Merino", "Rambouillet", "Targhee"},
	},
	{
		"species": "Caprine",
		"breeds":  []string{"Angora", "Boer", "Saanen"},
	},
	{
		"species": "Avian",
		"breeds":  []string{"Leghorn", "Silkie", "Rhode Island Red"},
	},
	{
		"species": "Reptilian",
		"breeds":  []string{"Ball Python", "Corn Snake", "Green Iguana"},
	},
	{
		"species": "Amphibian",
		"breeds":  []string{"American Bullfrog", "African Clawed Frog", "Red-Eyed Tree Frog"},
	},
	{
		"species": "Fish",
		"breeds":  []string{"Goldfish", "Guppy", "Corydoras"},
	},
	{
		"species": "Rodent",
		"breeds":  []string{"Rat", "Mouse", "Hamster", "Guinea Pig"},
	},
}

func SeedDB() {
	for _, s := range speciesData {
		name := s["species"].(string)
		var specie models.Species
		result := database.DB.FirstOrCreate(&specie, models.Species{Name: name})
		err := result.Error
		if err != nil {
			fmt.Println("Error while creating species:", name)
		} else {
			if result.RowsAffected == 0 {
				fmt.Println("Species already exists:", specie.Name)
			} else {
				fmt.Println("Species created:", specie.Name)
			}
		}

		fmt.Printf("\n")

		for _, b := range s["breeds"].([]string) {
			var breed models.Breed
			result := database.DB.FirstOrCreate(&breed, models.Breed{Name: b, SpeciesID: specie.ID})
			err := result.Error
			if err != nil {
				fmt.Println("Error while creating breed:", b)
			} else {
				if result.RowsAffected == 0 {
					fmt.Println("Breed already exists:", breed.Name)
				} else {
					fmt.Println("Breed created:", breed.Name)
				}
			}
		}

		fmt.Printf("\n")

	}
}
