package migration

import "golang-assignment/app/pet"

func MigrateAll() {
	pet.PetMigrate()
}
