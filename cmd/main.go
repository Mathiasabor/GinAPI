package main


	import
	(
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
		
	}
func main(){
	//initialisation du routeur gin 
	router :=gin.Default()

	//implementation d'une route http POST
	router.POST("/recipes",NewReceipeHandler)

	// execution au port par défaut 8080
	router.Run()
}

 
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
