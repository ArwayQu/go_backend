package routers

import (
	"ginDemo/controllers/resblock"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	api := r.Group("/v1")
	{
		api.GET("/resblock/info.json", resblock.GET)
	}

	return r
}
