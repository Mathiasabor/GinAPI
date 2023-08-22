package models
import "time"
type Recipe struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Tags []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
	Instruction []string `json:"instructions"`
	PublishedAt time.Time `json:"publishedAt"`
}