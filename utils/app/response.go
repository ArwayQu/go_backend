package app

import (
	"ginDemo/utils/e"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"status":        e.GetStatus(errCode),
		"error_code":    errCode,
		"error_message": e.GetMsg(errCode),
		"data":          data,
	})
}
