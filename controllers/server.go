package controllers

import (
	"github.com/gin-gonic/gin"
)

func Server() {
	router := gin.Default()
	router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumByID)
	router.GET("/albums/delete/:id", DeleteAlbumByID)
	router.POST("/albums", PostAlbums)
	router.POST("/albums/update", UpdateAlbums)

	router.Run("localhost:8080")
}
