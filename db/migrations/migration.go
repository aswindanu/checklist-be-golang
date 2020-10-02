package migrations

import (
	"encoding/json"
	"log"
	"golang-starterapi/models"
	"golang-starterbootstrap"
)

// Attention on table relation, seed table that not mutual first.
func Seed() {
	if bootstrap.App.ENV == "dev" || bootstrap.App.ENV == "staging" {

		seeder := bootstrap.App.DB.Begin()

		log.Printf("------------- SEED BEGIN ---------------")

		log.Printf("------------- SEED User ---------------")

		var users []models.User = []models.User{
			models.User{Name: "vian", Username: "vian", Password: "1234"},
			models.User{Name: "teta", Username: "teta", Password: "1234"},
			models.User{Name: "dhana", Username: "dhana", Password: "1234"},
		}

		for _, user := range users {
			if err := seeder.Create(&user).Error; err != nil {
				debugLog, _ := json.Marshal(err)
				log.Printf(string(debugLog))
			} else {
				log.Printf("created")
			}
		}

		if err := seeder.Commit().Error; err != nil {
			debugLog, _ := json.Marshal(err)
			log.Printf(string(debugLog))
		} else {
			log.Printf("------------- SEED Commited ---------------")
		}
		log.Printf("------------- SEED END ---------------")
	}
}

func Truncate() {
	if bootstrap.App.ENV == "dev" || bootstrap.App.ENV == "staging" {
		log.Printf("------------- TRUNCATE TABLE ---------------")
		bootstrap.App.DB.Exec("TRUNCATE TABLE user RESTART IDENTITY CASCADE;")
	}
}
