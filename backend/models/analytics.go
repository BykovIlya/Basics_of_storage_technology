package models

import "Basics_of_storage_technology/backend/utils"

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
