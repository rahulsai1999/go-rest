package misc

import (
	"github.com/gin-gonic/gin"
)

// External is an example of refactoring
func External(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "external",
	})
}
