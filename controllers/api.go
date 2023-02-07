package controllers

import (
	"api-db/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	// "encoding/json"
)

func Add(a, b int) int {
	return a + b
}

type AlbumJson struct {
	ID     int  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// var albums = []AlbumJson{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func GetAlbums(c *gin.Context) {
	allAlbum, err := model.AllAlbums()
	if err != nil {
		log.Println("allAlbum", err)
	}
	log.Println(allAlbum)
	c.IndentedJSON(http.StatusOK, allAlbum)
}

func GetAlbumByID(c *gin.Context) {
	var id string = c.Param("id")
	i, _ := strconv.Atoi(id)
	albumById, err := model.AlbumByID(i)
	if err != nil {
		log.Println("AlbumByID", err)
	}
	c.IndentedJSON(http.StatusOK, albumById)
}

func DeleteAlbumByID(c *gin.Context) {
	var id string = c.Param("id")
	i, _ := strconv.Atoi(id)
	log.Println(i)
	err := model.DeleteAlbum(i)
	if err != nil {
		log.Println("AlbumByID", err)
	}
	c.IndentedJSON(http.StatusOK, "success")
}

func PostAlbums(c *gin.Context) {
    var newAlbum model.Album

    if err := c.BindJSON(&newAlbum); err != nil {
		log.Println(err)
        return
    }

	albID, err := model.AddAlbum(newAlbum)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID of added album: %v\n", albID)
    c.JSON(http.StatusCreated,newAlbum)
}

func UpdateAlbums(c *gin.Context) {
    var album model.Album

    if err := c.BindJSON(&album); err != nil {
		log.Println(err)
        return
    }

	UpdateAlbum := model.UpdateAlbum(album)
	if UpdateAlbum != nil {
		log.Println("AlbumByID", UpdateAlbum)
	}
	log.Println(UpdateAlbum)

    c.JSON(http.StatusCreated,album)
}

