package handlers

import (
	"ginapi/internal/app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var recipes []models.Recipe


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