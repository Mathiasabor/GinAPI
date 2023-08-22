package main

import(
	 "github.com/gin-gonic/gin"
	 "ginapi/internal/app/handlers"
	 "ginapi/internal/app/models"
	
	)

	var recipes []models.Recipe
	
	func init()  {
		recipes = make([]models.Recipe, 0)
	}
func main(){
	
	router :=gin.Default()
	router.POST("/recipes",handlers.NewReceipeHandler)
	router.Run()
}

