package main 
import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "The Pointer Sisters", Price: 12.99},
	{ID: "2", Title: "Jeru", Artist: "The Gentlemen", Price: 10.99},
	{ID: "3", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 8.99},
}
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run(":8080")
	
}