package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Album sliced to seed record album data
var albums = []album{
	{ID: "1", Title: "After Hours", Artist: "The Weeknd", Price: 42.99},
	{ID: "2", Title: "House of Balloons", Artist: "The Weeknd", Price: 32.99},
	{ID: "3", Title: "Heros and Villains", Artist: "Metro Boomin", Price: 25.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Lets you POST movies in the form of a JSON request
func postAlbums(c *gin.Context) {
	var newAlbum album
	// Call BindJSON to bind the received JSON to newAlbum

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/abums/:id", getAlbumByID)

	router.Run()
}
