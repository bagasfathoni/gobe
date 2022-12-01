package gobe

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ==SUCCESS RESPONSES (2xx)==

// Return error status 200
func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
	})
}

// Return error status 200 with a specific message
func SuccessWithMessage(c *gin.Context, message interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "SUCCESS",
		"message": message,
	})
}

// 	==CLIENT ERROR RESPONSES (4xx)==

// Abort and return error status 400
func BadRequestError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"status": "FAILED",
	})
}

// Abort and return error status 400 with a specific message
func BadRequestErrorWithMessage(c *gin.Context, message interface{}) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"status":  "FAILED",
		"message": message,
	})
}

// Abort and return error status 403
func ForbiddenError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"status": "FAILED",
	})
}

// Abort and return error status 403 with a specific message
func ForbiddenErrorWithMessage(c *gin.Context, message interface{}) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"status":  "FAILED",
		"message": message,
	})
}

// Abort and return error status 401
func UnauthorizedError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"status": "FAILED",
	})
}

// Abort and return error status 401 with a specific message
func UnauthorizedErrorWithMessage(c *gin.Context, message interface{}) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"status":  "FAILED",
		"message": message,
	})
}

// ==SERVER ERROR RESPONSES(5xx)==

// Abort and return error status 500
func InternalServerError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"status": "FAILED",
	})
}

// Abort and return error status 500 with a specific message
func InternalServerErrorWithMessage(c *gin.Context, message interface{}) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"status":  "FAILED",
		"message": message,
	})
}
