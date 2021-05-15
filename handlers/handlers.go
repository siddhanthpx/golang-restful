package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func queryChecker(ok bool, c *gin.Context) {

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "invalid query",
		})
		return
	}

}
