package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"

	"ginapi/internal/app/models"

	"net/http"
	"time"

	"github.com/rs/xid"
)

	var recipes []models.Recipe
	//la fonction init permet d'initialiser la variable recipes
	func init()  {
		recipes = make([]models.Recipe, 0)
		file, _ := ioutil.ReadFile("recipes.json")
		_ = json.Unmarshal([]byte(file), &recipes)
	}
//manipulateur d'ajout 	 
func NewReceipeHandler(c *gin.Context){
	var recipe models.Recipe
	if err:=c.ShouldBindJSON(&recipe); err !=nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"error" : err.Error()})
				return
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)


}
func main(){
	//initialisation du routeur gin 
	router :=gin.Default()
	//implementation d'une route http POST
	router.POST("/recipes",NewReceipeHandler)
	//implementation d'une route http GET
	router.GET("/recipes",ListRecipesHandler)

	// execution au port par défaut 8080
	router.Run()
}

//manipulateur de lecture 
func ListRecipesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}