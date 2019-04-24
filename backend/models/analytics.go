package models

import (
	"Basics_of_storage_technology/backend/utils"
	"fmt"
)

type Request_1 struct {
	Movie       string `json:"movie"`
	Year        int64  `json:"year"`
	Studio      string `json:"studio"`
	Total_sales int64  `json:"total_sales"`
}

func GetReq1(x int64) []Request_1 {
	rows, err := DB.Query("SELECT title,year,studio,(domestic_sales+international_sales) "+
		"FROM movies JOIN boxoffice "+
		"ON movies.title = boxoffice.movie "+
		"WHERE domestic_sales+international_sales > $1", x)
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Request_1
	for rows.Next() {
		w := Request_1{}
		err = rows.Scan(&w.Movie, &w.Year, &w.Studio, &w.Total_sales)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

type Request_2 struct {
	Director string `json:"director"`
	Films    int64  `json:"films"`
}

func GetReq2() []Request_2 {
	rows, err := DB.Query("SELECT director, count(*) from movies group by director having count(*)>1")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Request_2
	for rows.Next() {
		w := Request_2{}
		err = rows.Scan(&w.Director, &w.Films)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

type Request_3 struct {
	Studio string `json:"studio"`
	Total  string `json:"total"`
}

func GetReq3() []Request_3 {
	rows, err := DB.Query("SELECT studio, sum(domestic_sales+international_sales) from movies left join boxoffice on movies.title = boxoffice.movie group by movies.studio order by  sum(domestic_sales+international_sales) DESC ")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Request_3
	for rows.Next() {
		w := Request_3{}
		err = rows.Scan(&w.Studio, &w.Total)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

type Request_4 struct {
	Director string `json:"director"`
	Age      int64  `json:"age"`
}

func GetReq4() []Request_4 {
	rows, err := DB.Query("SELECT DISTINCT director, age from movies inner join directors on movies.director = directors.name and movies.studio = 'Universal Studios' order by age DESC")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Request_4
	for rows.Next() {
		w := Request_4{}
		err = rows.Scan(&w.Director, &w.Age)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

type Request_5 struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Films  int64  `json:"films"`
	Email  string `json:"email"`
}

func GetReq5() []Request_5 {
	rows, err := DB.Query("select director, cnt, gender, email from (select director, count(*) as cnt from movies group by director having count(*)<2) movies join directors on movies.director = directors.name and gender = 'муж.' and email LIKE '%@gmail.com'")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Request_5
	for rows.Next() {
		w := Request_5{}
		err = rows.Scan(&w.Name, &w.Films, &w.Gender, &w.Email)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

type Request_6 struct {
	Studio string `json:"studio"`
	Year   int64  `json:"year"`
	Films  int64  `json:"films"`
}

func GetReq6() []Request_6 {
	rows, err := DB.Query("select studio, year, cnt from (select studio, count(*) as cnt from movies group by studio having count(*)>1) movies join studios on movies.studio = studios.name where year = (SELECT MAX(year) FROM (select studio, count(*) as cnt from movies group by studio having count(*)>1) movies join studios on movies.studio = studios.name)")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Request_6
	for rows.Next() {
		w := Request_6{}
		err = rows.Scan(&w.Studio, &w.Year, &w.Films)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

type Request_7 struct {
	Studio  string `json:"studio"`
	Percent string `json:"percent"`
}

func GetReq7() []Request_7 {
	rows, err := DB.Query("select studio, cast ((cast(cnt as float) / all_films * 100) as char(4))  from (select studio, count(*) as cnt from movies group by studio) movies join studios on movies.studio = studios.name")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Request_7
	for rows.Next() {
		w := Request_7{}
		err = rows.Scan(&w.Studio, &w.Percent)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

type Request_8 struct {
	Studio string `json:"studio"`
	Time   int64  `json:"time"`
}

func GetReq8(x int64) []Request_8 {
	rows, err := DB.Query("select studio, len  from (select studio, sum(length) as len from movies group by studio) movies join studios on movies.studio = studios.name where studios.year > $1", x)
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Request_8
	for rows.Next() {
		w := Request_8{}
		err = rows.Scan(&w.Studio, &w.Time)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

type SimpleRequest struct {
	Answer string `json:"answer"`
}

func GetReq9(x int64) []SimpleRequest {
	rows, err := DB.Query("select cast(avg(year) as char(4)) from movies where length > $1", x)
	utils.CheckErr(err)
	defer rows.Close()
	var ws []SimpleRequest
	for rows.Next() {
		w := SimpleRequest{}
		err = rows.Scan(&w.Answer)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

func GetReq10() []SimpleRequest {
	rows, err := DB.Query("select MAX(age) from directors where age NOT IN (select MAX(age) from directors)")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []SimpleRequest
	for rows.Next() {
		w := SimpleRequest{}
		err = rows.Scan(&w.Answer)
		utils.CheckErr(err)
		ws = append(ws, w)
		fmt.Print(w.Answer)
	}
	return ws
}

type Request_11 struct {
	Movie    string `json:"movie"`
	Director string `json:"director"`
	Age      string `json:"age"`
}

func GetReq11() []Request_11 {
	rows, err := DB.Query("select studio, cast ((cast(cnt as float) / all_films * 100) as char(4))  from (select studio, count(*) as cnt from movies group by studio) movies join studios on movies.studio = studios.name")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Request_11
	for rows.Next() {
		w := Request_11{}
		err = rows.Scan(&w.Movie, &w.Director, &w.Age)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

type Request_12 struct {
	Movie    string `json:"movie"`
	Director string `json:"director"`
	Age      string `json:"age"`
}

func GetReq12() []Request_12 {
	rows, err := DB.Query("select studio, cast ((cast(cnt as float) / all_films * 100) as char(4))  from (select studio, count(*) as cnt from movies group by studio) movies join studios on movies.studio = studios.name")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Request_12
	for rows.Next() {
		w := Request_12{}
		err = rows.Scan(&w.Movie, &w.Director, &w.Age)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

type Request_13 struct {
	Movie    string `json:"movie"`
	Director string `json:"director"`
	Age      string `json:"age"`
}

func GetReq13() []Request_13 {
	rows, err := DB.Query("select studio, cast ((cast(cnt as float) / all_films * 100) as char(4))  from (select studio, count(*) as cnt from movies group by studio) movies join studios on movies.studio = studios.name")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Request_13
	for rows.Next() {
		w := Request_13{}
		err = rows.Scan(&w.Movie, &w.Director, &w.Age)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

type Request_14 struct {
	Movie    string `json:"movie"`
	Director string `json:"director"`
	Age      string `json:"age"`
}

func GetReq14() []Request_14 {
	rows, err := DB.Query("select studio, cast ((cast(cnt as float) / all_films * 100) as char(4))  from (select studio, count(*) as cnt from movies group by studio) movies join studios on movies.studio = studios.name")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Request_14
	for rows.Next() {
		w := Request_14{}
		err = rows.Scan(&w.Movie, &w.Director, &w.Age)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}

type Request_15 struct {
	Movie    string `json:"movie"`
	Director string `json:"director"`
	Age      string `json:"age"`
}

func GetReq15() []Request_15 {
	rows, err := DB.Query("select studio, cast ((cast(cnt as float) / all_films * 100) as char(4))  from (select studio, count(*) as cnt from movies group by studio) movies join studios on movies.studio = studios.name")
	utils.CheckErr(err)
	defer rows.Close()
	var ws []Request_15
	for rows.Next() {
		w := Request_15{}
		err = rows.Scan(&w.Movie, &w.Director, &w.Age)
		utils.CheckErr(err)
		ws = append(ws, w)
	}
	return ws
}
