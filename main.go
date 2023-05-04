package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{ 
	{ID: "1", Title: "Cracker Island", Artist: "Gorillaz", Price: 12.99},
	{ID: "2", Title: "Apart Together", Artist: "Tim Minchin", Price: 16.99},
	{ID: "3", Title: "After Hours", Artist: "The Weekend", Price: 9.99},
}
// albums end points
func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:3000")
}

// getAlbums returns a list of albums as json
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}