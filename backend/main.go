package main

import (
	"Basics_of_storage_technology/backend/models"
	"Basics_of_storage_technology/backend/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

var DBport = "5432"

func main() {

	if models.InitDB(DBport) {
		fmt.Println("Database successfull initialisation")
		models.Migrate(models.DB)
		defer models.DB.Close()
	} else {
		log.Panic("Database not initialisation")
	}
	fmt.Print(models.GetDirectors())
	router := gin.Default()
	router.Use(CORSMiddleware())
	api := router.Group("/api")
	api.Use()
	{
		directors := api.Group("/directors")
		{
			directors.POST("", routes.CreateDirector)
			directors.GET("/id/:id", routes.GetDirectorById)
			directors.PUT("/:id", routes.UpdateDirector)
			directors.DELETE("/:id", routes.DeleteDirector)
			directors.GET("", routes.GetDirectors)
		}
		movies := api.Group("/movies")
		{
			movies.POST("", routes.CreateMovie)
			movies.GET("/id/:id", routes.GetMovieById)
			movies.PUT("/:id", routes.UpdateMovie)
			movies.DELETE("/:id", routes.DeleteMovie)
			movies.GET("", routes.GetMovies)
		}
		boxoffice := api.Group("/boxoffice")
		{
			boxoffice.POST("", routes.CreateBoxoffice)
			boxoffice.GET("/id/:id", routes.GetBoxofficeById)
			boxoffice.PUT("/:id", routes.UpdateBoxoffice)
			boxoffice.DELETE("/:id", routes.DeleteBoxoffice)
			boxoffice.GET("", routes.GetBoxoffices)
		}
		studios := api.Group("/studios")
		{
			studios.POST("", routes.CreateStudio)
			studios.GET("/id/:id", routes.GetStudioById)
			studios.PUT("/:id", routes.UpdateStudio)
			studios.DELETE("/:id", routes.DeleteStudio)
			studios.GET("", routes.GetStudios)
		}
		analytics := api.Group("/analytics")
		{
			analytics.GET("/req1/:id", routes.GetReq1)
			analytics.GET("/req2", routes.GetReq2)
			analytics.GET("/req3", routes.GetReq3)
			analytics.GET("/req4", routes.GetReq4)
			analytics.GET("/req5", routes.GetReq5)
			analytics.GET("/req6", routes.GetReq6)
			analytics.GET("/req7", routes.GetReq7)
			analytics.GET("/req8/:id", routes.GetReq8)
			analytics.GET("/req9/:id", routes.GetReq9)
			analytics.GET("/req10", routes.GetReq10)
			analytics.GET("/req11", routes.GetReq11)
			analytics.GET("/req12", routes.GetReq12)
			analytics.GET("/req13", routes.GetReq13)
			analytics.GET("/req14", routes.GetReq14)
			analytics.GET("/req15", routes.GetReq15)

		}
	}
	router.Run(":5000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length,X-Requested-With, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
