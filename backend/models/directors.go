package models

import (
	"Basics_of_storage_technology/backend/utils"
	"fmt"
)

type Director struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Age    int64  `json:"age"`
	Gender bool   `json:"gender"`
}

func CreateDirector(w Director) int64 {
	var uid int64
	err := DB.QueryRow("INSERT INTO directors(name,age,gender) VALUES($1,$2,$3) returning id",
		w.Name, w.Age, w.Gender).Scan(&uid)
	utils.CheckErr(err)
	return uid
}

func UpdateDirector(w Director) bool {
	stmt, err := DB.Prepare("update directors set (name, age, gender)=($1,$2,$3) where id=$4")
	utils.CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(w.Name, w.Age, w.Gender, w.Id)
	utils.CheckErr(err)

	affect, err := res.RowsAffected()
	utils.CheckErr(err)

	fmt.Println(affect, "rows changed")
	if affect > 0 {
		return true
	}
	return false
}

func DeleteDirector(w Director) bool {
	stmt, err := DB.Prepare("delete from directors where id=$1")
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

func GetDirectors() []Director {
	rows, err := DB.Query("SELECT id,name,age,gender FROM directors LIMIT 1000")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Director
	for rows.Next() {
		w := Director{}
		err = rows.Scan(&w.Id, &w.Name, &w.Age, &w.Gender)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

func GetDirectorById(id int64) *Director {
	rows, err := DB.Query("SELECT id,name,age,gender FROM directors WHERE id=$1 LIMIT 1", id)
	utils.CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		w := Director{}
		err = rows.Scan(&w.Id, &w.Name, &w.Age, &w.Gender)
		utils.CheckErr(err)
		return &w
	}
	return nil
}
