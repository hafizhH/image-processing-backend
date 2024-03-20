package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		log.Println(err.Err.Error())
	}

	c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error Occured"})
}
