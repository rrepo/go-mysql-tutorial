package main

import (
	"api-db/controllers"
	"api-db/model"
	// "log"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	model.ConnectDb()

	// albID, err := model.AddAlbum(model.Album{
	// 	Title:  "3433221",
	// 	Artist: "Betty Carter",
	// 	Price:  49.99,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("ID of added album: %v\n", albID)

	// allAlbum,err := model.AllAlbums()
	// if(err != nil){
	// 	log.Println("allAlbum",err)
	// }
	// log.Println(allAlbum)

	// UpdateAlbum := model.UpdateAlbum(
	// 	2,
	// 	model.Album{
	// 		Title:  "Test butt",
	// 		Artist: "Betty buttcheek",
	// 		Price:  243,
	// 	})
	// if UpdateAlbum != nil {
	// 	log.Println("AlbumByID", UpdateAlbum)
	// }
	// log.Println(UpdateAlbum)

	controllers.Server()

	// albumById,err := model.AlbumByID(1)
	// if(err != nil){
	// 	log.Println("AlbumByID",err)
	// }
	// log.Println(albumById)

	// albumsByArtist, err := model.AlbumsByArtist("John Coltrane")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("Albums found: %v\n", albumsByArtist)

	// model.DeleteAlbum(14)

	// UpdateAlbum := model.UpdateAlbum(1)
	// if(err != nil){
	// 	log.Println("AlbumByID",err)
	// }
	// log.Println(UpdateAlbum)
}

// show databases;
// USE recordings;
// SELECT * FROM album;
