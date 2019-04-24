package routes

import (
	"Basics_of_storage_technology/backend/models"
	"Basics_of_storage_technology/backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetReq1(c *gin.Context) {
	param := c.Param("id")
	x, err := strconv.ParseInt(param, 10, 64)
	utils.CheckErr(err)
	w := models.GetReq1(x)
	if w != nil {
		c.JSON(http.StatusOK, w)
	} else {
		c.JSON(http.StatusNotFound, ApiMessage{utils.EntityNotExistMessage("req1")})
	}
}

func GetReq2(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetReq2())
}

func GetReq3(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetReq3())
}

func GetReq4(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetReq4())
}

func GetReq5(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetReq5())
}

func GetReq6(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetReq6())
}
