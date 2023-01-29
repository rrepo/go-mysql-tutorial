package main

import (
	"api-db/model"
	"log"
)

type Album struct {  
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	model.ConnectDb()

	allAlbum,err := model.AllAlbums()
	if(err != nil){
		log.Println("allAlbum",err)
	}
	log.Println(allAlbum)

	albumById,err := model.AlbumByID(1)
	if(err != nil){
		log.Println("AlbumByID",err)
	}
	log.Println(albumById)

	albumsByArtist, err := model.AlbumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Albums found: %v\n", albumsByArtist)

	albID, err := model.AddAlbum(model.Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID of added album: %v\n", albID)
}



// show databases;
// SELECT * FROM album;