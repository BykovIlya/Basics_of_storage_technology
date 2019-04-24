package models

import (
	"Basics_of_storage_technology/backend/utils"
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
