package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Album struct
type album struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Albums data for test
var albums = []album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	//Handler route for (GET)/albums
	router.GET("/albums", getAlbums)

	//handler route for (POST)/albums
	router.POST("albums", postAlbums)

	//handler rpute for (get)/algums/:id
	router.GET("albums/:id", GetAlbumById)

	//server listening
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	//CallBind json to bind the recieved json to object
	if error := c.BindJSON(&newAlbum); error != nil {
		return
	}

	//Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetalbumById returns an album that matches the id field
func GetAlbumById(c *gin.Context) {
	//getting id param
	/* id := c.Param("id") */
	id := c.Params.ByName("id")
	fmt.Println(id)

	//look for album
	for _, a := range albums {
		if a.Id == id {
			c.IndentedJSON(http.StatusFound, a)
			return
		}
	}

	//if album not found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
