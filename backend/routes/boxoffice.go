package routes

import (
	"Basics_of_storage_technology/backend/models"
	"Basics_of_storage_technology/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var BoxofficeTitle = "boxoffice"

func CreateBoxoffice(c *gin.Context) {
	var w models.Boxoffice
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetBoxofficeByMovie(w.Movie)
		if existItem != nil {
			c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityExistMessage(BoxofficeTitle)})
		} else {
			uid := models.CreateBoxoffice(w)
			c.JSON(http.StatusOK, ApiStatus{uid, utils.SuccessMessage})
		}
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func GetBoxofficeById(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)
	w := models.GetBoxofficeById(id)
	if w != nil {
		c.JSON(http.StatusOK, w)
	} else {
		c.JSON(http.StatusNotFound, ApiMessage{utils.EntityNotExistMessage(BoxofficeTitle)})
	}
}

func GetBoxofficeByMovie(c *gin.Context) {
	idParam := c.Param(utils.Movie)
	w := models.GetBoxofficeByMovie(idParam)
	if w != nil {
		c.JSON(http.StatusOK, w)
	} else {
		c.JSON(http.StatusNotFound, ApiMessage{utils.EntityNotExistMessage(BoxofficeTitle)})
	}
}

func UpdateBoxoffice(c *gin.Context) {
	var w models.Boxoffice
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetBoxofficeById(w.Id)
		if existItem == nil {
			c.JSON(http.StatusUnprocessableEntity, ApiStatus{existItem.Id, utils.EntityWithIdNotExistMessage(BoxofficeTitle)})
		} else {
			if models.UpdateBoxoffice(w) {
				c.JSON(http.StatusOK, ApiStatus{w.Id, utils.SuccessMessage})
			} else {
				c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotUpdatedMessage(BoxofficeTitle)})
			}
		}
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func DeleteBoxoffice(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)

	existItem := models.GetBoxofficeById(id)
	if existItem == nil {
		c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityNotExistMessage(BoxofficeTitle)})
	} else {
		if models.DeleteBoxoffice(*existItem) {
			c.JSON(http.StatusOK, ApiMessage{utils.SuccessMessage})
		} else {
			c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotDeletedMessage(BoxofficeTitle)})
		}
	}
}

func GetBoxoffices(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetBoxoffices())
}
