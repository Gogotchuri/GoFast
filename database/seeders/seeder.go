package seeders

import (
	"github.com/Gogotchuri/GoFast/app/models"
	"github.com/Gogotchuri/GoFast/app/services/hash"
)

/**
 * DefaultSeed seeds default seeders
 */
func DefaultSeed() {
	//TODO add other seeders
	Seed(SeedUsers)
}

/*Seed Calls passed seeder functions*/
func Seed(seeders ...func()) {
	for _, s := range seeders {
		s() //Call seeder
	}
}

/**
 * SeedUsers add first users to database
 */
func SeedUsers() {
	defAdmin := &models.User{
		Password:  "k58h4ZogdYZT1pb5YEziooRrzSy2VCs9RQ9QAL3xTKM=", // Hash for "admin"
		FirstName: "admin",
		LastName:  "admin",
		Email:     "admin@example.com",
		Role:      0,
	}
	defAdmin.SaveUserIfDoesntExist()

	defUser := &models.User{
		Password:  "pYxYzVLmGCit32aYrak6AmU1WTsq4Mvm75vUxrtyE6A=", // Hash for "user"
		FirstName: "user",
		LastName:  "user",
		Email:     "user@example.com",
		Role:      1,
	}
	defUser.SaveUserIfDoesntExist()

	defUser1 := &models.User{
		Password:  hash.GetPasswordHash("user"), // Hash for "user"
		FirstName: "user",
		LastName:  "user",
		Email:     "user@example.com",
		Role:      1,
	}

	defUser1.SaveUserIfDoesntExist()
}
