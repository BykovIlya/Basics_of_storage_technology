package routes

import (
	"Basics_of_storage_technology/backend/models"
	"Basics_of_storage_technology/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var DirectorTitle = "director"

func CreateDirector(c *gin.Context) {
	var w models.Director
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetDirectorByName(w.Name)
		if existItem != nil {
			c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityExistMessage(DirectorTitle)})
		} else {
			uid := models.CreateDirector(w)
			c.JSON(http.StatusOK, ApiStatus{uid, utils.SuccessMessage})
		}
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func GetDirectorById(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)
	w := models.GetDirectorById(id)
	if w != nil {
		c.JSON(http.StatusOK, w)
	} else {
		c.JSON(http.StatusNotFound, ApiMessage{utils.EntityNotExistMessage(DirectorTitle)})
	}
}

func UpdateDirector(c *gin.Context) {
	var w models.Director
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetDirectorById(w.Id)
		if existItem == nil {
			c.JSON(http.StatusUnprocessableEntity, ApiStatus{existItem.Id, utils.EntityWithIdNotExistMessage(DirectorTitle)})
		} else {
			if models.UpdateDirector(w) {
				c.JSON(http.StatusOK, ApiStatus{w.Id, utils.SuccessMessage})
			} else {
				c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotUpdatedMessage(DirectorTitle)})
			}
		}
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func DeleteDirector(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)

	existItem := models.GetDirectorById(id)
	if existItem == nil {
		c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityNotExistMessage(DirectorTitle)})
	} else {
		if models.DeleteDirector(*existItem) {
			c.JSON(http.StatusOK, ApiMessage{utils.SuccessMessage})
		} else {
			c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotDeletedMessage(DirectorTitle)})
		}
	}
}

func GetDirectors(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetDirectors)
}
