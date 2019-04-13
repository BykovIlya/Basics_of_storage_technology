package routes

import (
	"Basics_of_storage_technology/backend/models"
	"Basics_of_storage_technology/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var StudioTitle = "studio"

func CreateStudio(c *gin.Context) {
	var w models.Studio
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetStudioByName(w.Name)
		if existItem != nil {
			c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityExistMessage(StudioTitle)})
		} else {
			uid := models.CreateStudio(w)
			c.JSON(http.StatusOK, ApiStatus{uid, utils.SuccessMessage})
		}
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func GetStudioById(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)
	w := models.GetStudioById(id)
	if w != nil {
		c.JSON(http.StatusOK, w)
	} else {
		c.JSON(http.StatusNotFound, ApiMessage{utils.EntityNotExistMessage(StudioTitle)})
	}
}

func GetStudioByName(c *gin.Context) {
	idParam := c.Param(utils.Name)
	w := models.GetStudioByName(idParam)
	if w != nil {
		c.JSON(http.StatusOK, w)
	} else {
		c.JSON(http.StatusNotFound, ApiMessage{utils.EntityNotExistMessage(StudioTitle)})
	}
}

func UpdateStudio(c *gin.Context) {
	var w models.Studio
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetStudioById(w.Id)
		if existItem == nil {
			c.JSON(http.StatusUnprocessableEntity, ApiStatus{existItem.Id, utils.EntityWithIdNotExistMessage(StudioTitle)})
		} else {
			if models.UpdateStudio(w) {
				c.JSON(http.StatusOK, ApiStatus{w.Id, utils.SuccessMessage})
			} else {
				c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotUpdatedMessage(StudioTitle)})
			}
		}
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func DeleteStudio(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)

	existItem := models.GetStudioById(id)
	if existItem == nil {
		c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityNotExistMessage(StudioTitle)})
	} else {
		if models.DeleteStudio(*existItem) {
			c.JSON(http.StatusOK, ApiMessage{utils.SuccessMessage})
		} else {
			c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotDeletedMessage(StudioTitle)})
		}
	}
}

func GetStudios(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetStudios())
}
