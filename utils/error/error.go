package er

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, msg string) error {
	return errors.New(msg)
}

// Write error to the context only
func WithError(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{
		"message": msg,
	})
}

// Write error at the context and return the error
func WithContextError(c *gin.Context, status int, msg string) error {
	c.JSON(status, gin.H{
		"message": msg,
	})
	return errors.New(msg)
}
