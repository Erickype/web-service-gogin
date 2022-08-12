package main

import (
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

	//server listening
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
