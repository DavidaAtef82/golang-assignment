package pet

import (
	"github.com/labstack/echo"
	"golang-assignment/helper"
	"net/http"
	"strconv"
)

func GetPets(c echo.Context) error {
	pets, _ := GetPetsModel()
	return c.JSON(http.StatusOK, pets)
}

func GetPet(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		return c.JSON(http.StatusBadRequest, Pet{})
	}
	row, _ := GetUserByID(uint(id))
	if row.ID == 0 {
		return c.JSON(http.StatusNotFound, "item not found")
	}
	return c.JSON(http.StatusOK, row)
}

func CreatePet(c echo.Context) error {
	pet := Pet{}
	if err := c.Bind(&pet); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}
	if pet.Name == "" {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}
	err, pet := CreatePetModel(pet.Name, pet.Breed, pet.Age)
	if helper.CheckError(err) {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, pet)

}

func UpdatePet(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		return c.JSON(http.StatusBadRequest, Pet{})
	}
	var pet Pet
	if err := c.Bind(&pet); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}
	err, pet := UpdatePetModel(uint(id), pet.Name, pet.Breed, pet.Age)
	if helper.CheckError(err) {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, pet)
}

func DeletePet(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if id == 0 {
		return c.JSON(http.StatusBadRequest, Pet{})
	}
	row, _ := GetUserByID(uint(id))
	if row.ID == 0 {
		return c.JSON(http.StatusNotFound, Pet{})
	}
	err := DeletePetModel(uint(id))
	if helper.CheckError(err) {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Pet deleted")

}
