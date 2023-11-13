package pet

import "github.com/labstack/echo"

func Routes(r *echo.Echo) *echo.Echo {
	routesGroup := r.Group("/pets")
	{
		// Pet endpoints
		routesGroup.GET("/", GetPets)
		routesGroup.GET("/:id", GetPet)
		routesGroup.POST("/", CreatePet)
		routesGroup.PUT("/:id", UpdatePet)
		routesGroup.DELETE("/:id", DeletePet)
	}
	return r
}
