package pet

import (
	"golang-assignment/database"
)

type Pet struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"uniqueIndex;not null" json:"name"`
	Breed string `gorm:"not null" json:"breed"`
	Age   int    `json:"age"`
}

func PetMigrate() {
	database.Migrate(&Pet{})
}

// Get all pets
func GetPetsModel() ([]Pet, error) {
	var pets []Pet
	result := database.DB.Find(&pets)
	return pets, result.Error
}

// Get a user by ID
func GetUserByID(id uint) (Pet, error) {
	var row Pet
	if err := database.DB.First(&row, id).Error; err != nil {
		return row, err
	}
	return row, nil
}

// Functions to perform CRUD operations on the Pet model
// Create a new pet
func CreatePetModel(name, breed string, age int) (error, Pet) {
	pet := Pet{Name: name, Breed: breed, Age: age}
	result := database.DB.Create(&pet)
	return result.Error, pet
}

// Update a pet
func UpdatePetModel(id uint, name, breed string, age int) (error, Pet) {
	var pet Pet
	result := database.DB.First(&pet, id)
	if result.Error != nil {
		return result.Error, pet
	}
	pet.Name = name
	pet.Breed = breed
	pet.Age = age
	result = database.DB.Save(&pet)
	return result.Error, pet
}

// Delete a pet
func DeletePetModel(id uint) error {
	var pet Pet
	result := database.DB.First(&pet, id)
	if result.Error != nil {
		return result.Error
	}
	result = database.DB.Delete(&pet)
	return result.Error
}
