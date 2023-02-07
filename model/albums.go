package model

import (
	"database/sql"
	"fmt"
	"log"
)

type Album struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  float32 `json:"price"`
}

func AllAlbums() (albums []Album, err error) {
	rows, err := Db.Query("SELECT * FROM album")
	if err != nil {
		return nil,err
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil,err
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return albums,nil
}

func AlbumByID(id int) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := Db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

func AlbumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := Db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}


func AddAlbum(alb Album) (int64, error) {
	result, err := Db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func DeleteAlbum(id int)(err error){
	result, err := Db.Exec("DELETE FROM album WHERE id=?;",id)
	if err != nil {
		return fmt.Errorf("addAlbum: %v", err)
	}
	log.Println(result)
	return nil
}

func UpdateAlbum(alb Album)(err error){
	up, err := Db.Prepare("UPDATE album SET title=?, artist=?, price=? WHERE id=?")
	// up, err := Db.Prepare("UPDATE album SET title=? WHERE id=?")
	if err != nil {
		return fmt.Errorf("addAlbum: %v", err)
	}
	
	result, err := up.Exec(alb.Title, alb.Artist, alb.Price, alb.ID)
	if err != nil {
		return fmt.Errorf("addAlbum: %v", err)
	}
	log.Println(result)
	return nil
}

