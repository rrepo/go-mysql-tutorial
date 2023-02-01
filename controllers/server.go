package controllers

import (
	"github.com/gin-gonic/gin"
)


func Server(){
	router := gin.Default()
	router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumByID)
	
	router.Run("localhost:8080")
}