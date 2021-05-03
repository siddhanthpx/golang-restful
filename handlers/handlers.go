package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategory(c *gin.Context) {
	category, ok := c.Params.Get("category")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "invalid query",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": category,
	})
}
