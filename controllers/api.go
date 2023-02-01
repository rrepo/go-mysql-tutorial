package controllers

import (
	"api-db/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Add(a, b int) int {
	return a + b
}

type AlbumJson struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []AlbumJson{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	allAlbum, err := model.AllAlbums()
	if err != nil {
		log.Println("allAlbum", err)
	}
	log.Println(allAlbum)

	// out, _ := json.Marshal(allAlbum)
	// log.Println("out",string(out))

	c.IndentedJSON(http.StatusOK, allAlbum)
}

func GetAlbumByID(c *gin.Context) {
	var id string = c.Param("id")
	i, _ := strconv.Atoi(id)
	log.Println(i)
	albumById, err := model.AlbumByID(i)
	if err != nil {
		log.Println("AlbumByID", err)
	}
	c.IndentedJSON(http.StatusOK, albumById)
}
