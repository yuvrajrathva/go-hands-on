package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}
func addAlbum(context *gin.Context) {
	var newAlbum album

	err := context.BindJSON(&newAlbum)
	if err != nil {
		return
	}

	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}
func getAlbumById(id string) (*album, error) {
	for i, a := range albums {
		if a.ID == id {
			return &albums[i], nil
		}
	}
	return nil, errors.New("not found in albums")

}
func getAlbum(context *gin.Context) {
	id := context.Param("id")
	result, err := getAlbumById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, result)
}
func main() {
	router := gin.Default()
	router.GET("/getAlbums", getAlbums)
	router.POST("/addAlbum", addAlbum)
	router.GET("/getAlbum/:id", getAlbum)
	router.Run("localhost:8000")
}
