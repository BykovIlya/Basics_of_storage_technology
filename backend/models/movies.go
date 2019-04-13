package models

import (
	"Basics_of_storage_technology/backend/utils"
	"fmt"
)

type Movie struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Director string `json:"movie"`
	Year     int64  `json:"year"`
	Length   int64  `json:"length"`
	Studio   string `json:"studio"`
}

func CreateMovie(w Movie) int64 {
	var uid int64
	err := DB.QueryRow("INSERT INTO movies(title,director,year,length,studio) VALUES($1,$2,$3,$4,$5) returning id",
		w.Title, w.Director, w.Year, w.Length, w.Studio).Scan(&uid)
	utils.CheckErr(err)
	return uid
}

func UpdateMovie(w Movie) bool {
	stmt, err := DB.Prepare("update movies set (title, director, year,length,studio)=($1,$2,$3,$4,$5) where id=$6")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(w.Title, w.Director, w.Year, &w.Length, &w.Studio, w.Id)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")
	if affect > 0 {
		return true
	}
	return false
}

func DeleteMovie(w Movie) bool {
	stmt, err := DB.Prepare("delete from movies where id=$1")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(w.Id)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")
	if affect > 0 {
		return true
	}
	return false
}

func GetMovies() []Movie {
	rows, err := DB.Query("SELECT id,title,director,year,length,studio FROM movies LIMIT 1000")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Movie
	for rows.Next() {
		w := Movie{}
		err = rows.Scan(&w.Id, &w.Title, &w.Director, &w.Year, &w.Length, &w.Studio)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

func GetMovieById(id int64) *Movie {
	rows, err := DB.Query("SELECT id,title,director,year,length,studio FROM movies WHERE id=$1 LIMIT 1", id)
	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		w := Movie{}
		err = rows.Scan(&w.Id, &w.Title, &w.Director, &w.Year, &w.Length, &w.Studio)
		utils.CheckErr(err)
		return &w
	}
	return nil
}

func GetMovieByTitle(title string) *Movie {
	rows, err := DB.Query("SELECT id,title,director,year,length,studio FROM movies WHERE title=$1 LIMIT 1", title)
	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		w := Movie{}
		err = rows.Scan(&w.Id, &w.Title, &w.Director, &w.Year, &w.Length, &w.Studio)
		utils.CheckErr(err)
		return &w
	}
	return nil
}
