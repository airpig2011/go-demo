package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"one/go-misc/modules/map/dao"
	"one/go-misc/utils/response"
)

func Test(context *gin.Context) {

	appG := response.Gin{C: context}
	//name := context.Query("name")
	//state := -1
	//if arg := context.Query("state"); arg != "" {
	//	state = 1
	//}

	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
	//	return
	//}
	dao.Insert()


	appG.Response(http.StatusOK, response.SUCCESS, map[string]interface{}{
		"lists": 0,
		"total": 0,
	})
}