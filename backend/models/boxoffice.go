package models

import (
	"Basics_of_storage_technology/backend/utils"
	"fmt"
)

type Boxoffice struct {
	Id                  int64  `json:"id"`
	Movie               string `json:"movie"`
	Domestic_sales      int64  `json:"domestic_sales"`
	International_sales int64  `json:"international_sales"`
}

func CreateBoxoffice(w Boxoffice) int64 {
	var uid int64
	err := DB.QueryRow("INSERT INTO boxoffice(movie,domestic_sales,international_sales) VALUES($1,$2,$3) returning id",
		w.Movie, w.Domestic_sales, w.International_sales).Scan(&uid)
	utils.CheckErr(err)
	return uid
}

func UpdateBoxoffice(w Boxoffice) bool {
	stmt, err := DB.Prepare("update boxoffice set (movie, domestic_sales, international_sales)=($1,$2,$3) where id=$4")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(w.Movie, w.Domestic_sales, w.International_sales, w.Id)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")
	if affect > 0 {
		return true
	}
	return false
}

func DeleteBoxoffice(w Boxoffice) bool {
	stmt, err := DB.Prepare("delete from boxoffice where id=$1")
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

func GetBoxoffices() []Boxoffice {
	rows, err := DB.Query("SELECT id,movie,domestic_sales,international_sales FROM boxoffice LIMIT 1000")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Boxoffice
	for rows.Next() {
		w := Boxoffice{}
		err = rows.Scan(&w.Id, &w.Movie, &w.Domestic_sales, &w.International_sales)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

func GetBoxofficeById(id int64) *Boxoffice {
	rows, err := DB.Query("SELECT id,movie,domestic_sales,international_sales FROM boxoffice WHERE id=$1 LIMIT 1", id)
	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		w := Boxoffice{}
		err = rows.Scan(&w.Id, &w.Movie, &w.Domestic_sales, &w.International_sales)
		utils.CheckErr(err)
		return &w
	}
	return nil
}

func GetBoxofficeByMovie(movie string) *Boxoffice {
	rows, err := DB.Query("SELECT id,movie,domestic_sales,international_sales FROM boxoffice WHERE movie=$1 LIMIT 1", movie)
	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		w := Boxoffice{}
		err = rows.Scan(&w.Id, &w.Movie, &w.Domestic_sales, &w.International_sales)
		utils.CheckErr(err)
		return &w
	}
	return nil
}
