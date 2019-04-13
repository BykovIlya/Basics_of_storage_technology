package models

import (
	"Basics_of_storage_technology/backend/utils"
	"fmt"
)

type Studio struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Year      string `json:"year"`
	All_films int64  `json:"all_films"`
}

func CreateStudio(w Studio) int64 {
	var uid int64
	err := DB.QueryRow("INSERT INTO studios(name,year,all_films) VALUES($1,$2,$3) returning id",
		w.Name, w.Year, w.All_films).Scan(&uid)
	utils.CheckErr(err)
	return uid
}

func UpdateStudio(w Studio) bool {
	stmt, err := DB.Prepare("update studios set (name,year,all_films)=($1,$2,$3) where id=$4")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(w.Name, w.Year, w.All_films, w.Id)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")
	if affect > 0 {
		return true
	}
	return false
}

func DeleteStudio(w Studio) bool {
	stmt, err := DB.Prepare("delete from studios where id=$1")
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

func GetStudios() []Studio {
	rows, err := DB.Query("SELECT id,name,year,all_films FROM studios LIMIT 1000")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Studio
	for rows.Next() {
		w := Studio{}
		err = rows.Scan(&w.Id, &w.Name, &w.Year, &w.All_films)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

func GetStudioById(id int64) *Studio {
	rows, err := DB.Query("SELECT id,name,year,all_films FROM studios WHERE id=$1 LIMIT 1", id)
	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		w := Studio{}
		err = rows.Scan(&w.Id, &w.Name, &w.Year, &w.All_films)
		utils.CheckErr(err)
		return &w
	}
	return nil
}
