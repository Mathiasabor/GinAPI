package models
import "time"
//ici nous allong définir le model de donnée et le format json et leur noms correspondant
type Recipe struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Tags []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
	Instruction []string `json:"instructions"`
	PublishedAt time.Time `json:"publishedAt"`
}