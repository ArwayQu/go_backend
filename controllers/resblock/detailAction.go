package resblock

import (
	"ginDemo/utils/app"
	"ginDemo/utils/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
	userInfoMap := map[string]string{
		"username":     "小明",
		"professional": "电竞选手",
	}
	app.Response(c, http.StatusOK, e.SUCCESS, userInfoMap)
}
