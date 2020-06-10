package seeders

import "github.com/Gogotchuri/GoFast/app/models"

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
		Password:  "admin",
		FirstName: "admin",
		LastName:  "admin",
		Email:     "admin@example.com",
		Role:      0,
	}
	defAdmin.SaveUserIfDoesntExist()

	defUser := &models.User{
		Password:  "user",
		FirstName: "user",
		LastName:  "user",
		Email:     "user@example.com",
		Role:      1,
	}
	defUser.SaveUserIfDoesntExist()
}
