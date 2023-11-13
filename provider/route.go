package provider

import (
	"github.com/labstack/echo"
	"golang-assignment/app/pet"
	"golang-assignment/helper"
	"os"
)

func Run() {
	r := echo.New()
	pet.Routes(r)
	err := r.Start(":" + os.Getenv("APP_PORT"))
	helper.CheckError(err)
}
