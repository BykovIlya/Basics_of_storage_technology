package routes

import (
	"Basics_of_storage_technology/backend/models"
	"Basics_of_storage_technology/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var MovieTitle = "movie"

func CreateMovie(c *gin.Context) {
	var w models.Movie
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetMovieByTitle(w.Title)
		if existItem != nil {
			c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityExistMessage(MovieTitle)})
		} else {
			uid := models.CreateMovie(w)
			c.JSON(http.StatusOK, ApiStatus{uid, utils.SuccessMessage})
		}
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func GetMovieById(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)
	w := models.GetMovieById(id)
	if w != nil {
		c.JSON(http.StatusOK, w)
	} else {
		c.JSON(http.StatusNotFound, ApiMessage{utils.EntityNotExistMessage(MovieTitle)})
	}
}

func GetMovieByTitle(c *gin.Context) {
	idParam := c.Param(utils.Title)
	w := models.GetMovieByTitle(idParam)
	if w != nil {
		c.JSON(http.StatusOK, w)
	} else {
		c.JSON(http.StatusNotFound, ApiMessage{utils.EntityNotExistMessage(MovieTitle)})
	}
}

func UpdateMovie(c *gin.Context) {
	var w models.Movie
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetMovieById(w.Id)
		if existItem == nil {
			c.JSON(http.StatusUnprocessableEntity, ApiStatus{existItem.Id, utils.EntityWithIdNotExistMessage(MovieTitle)})
		} else {
			if models.UpdateMovie(w) {
				c.JSON(http.StatusOK, ApiStatus{w.Id, utils.SuccessMessage})
			} else {
				c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotUpdatedMessage(MovieTitle)})
			}
		}
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func DeleteMovie(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)

	existItem := models.GetMovieById(id)
	if existItem == nil {
		c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityNotExistMessage(MovieTitle)})
	} else {
		if models.DeleteMovie(*existItem) {
			c.JSON(http.StatusOK, ApiMessage{utils.SuccessMessage})
		} else {
			c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotDeletedMessage(MovieTitle)})
		}
	}
}

func GetMovies(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetMovies())
}
