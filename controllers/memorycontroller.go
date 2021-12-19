package controllers

import (
	CONSTANTS "github.com/alicobanserver/constants"
	"github.com/alicobanserver/models"
	store "github.com/alicobanserver/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = store.DB(models.Database{})

/*
 - Set Endpoint
 - Set Key and Value
*/
func Set(c *gin.Context) {
	var memory = models.Memory{}
	c.Bind(&memory)

	value := memory.Value
	key, err := db.Set(value)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": CONSTANTS.SETDATA_ERR_MSG,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"key": key,
	})
}

/*
 - Get Endpoint(GET)
 - Get Value by key
*/
func Get(c *gin.Context) {

	id := c.Param("key")
	getValue, err := db.Get(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": CONSTANTS.GETDATA_ERR_MSG,
		})
		return
	}
	if getValue == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": CONSTANTS.NOT_FOUND,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"value": getValue,
	})

}

/*
 - Delete Endpoint(POST)
 - Delete Value by key
*/
func DeleteSingle(c *gin.Context) {
	key := c.Param("key")
	err := db.DeleteSingle(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": CONSTANTS.DELETESINGLE_ERR_MSG,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success Delete Single Data",
	})
}

/*
 - Flush Endpoint(POST)
 - Remove All data
*/
func FlushData(c *gin.Context) {
	err := db.FlushData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": CONSTANTS.FLUSHDATA_ERR_MSG,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success Flush Data",
	})
	return
}

/*
 - GetAll Endpoint(POST)
 - Get all value
*/
func GetAll(c *gin.Context) {
	arr, err := db.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": arr,
	})
}
