package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseNormal200(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}

func ResponseBadRequest400(c *gin.Context, msg ...interface{}) {
	c.AbortWithStatusJSON(http.StatusBadRequest, &ErrorStatus400{Code: 3, Message: defaultParamMsg("参数错误", msg...)})
}

func ResponseUnauthorized401(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, &ErrorStatus401{Code: 16, Message: defaultMsg("未授权", msg...)})
}

func ResponseForbidden403(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusForbidden, &ErrorStatus403{Code: 7, Message: defaultMsg("禁止操作", msg...)})
}

func ResponseNotFound404(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusNotFound, &ErrorStatus404{Code: 5, Message: defaultMsg("资源未找到", msg...)})
}

func ResponseInternalError500(c *gin.Context, err error, msg ...string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorStatus500{Code: 10, Message: defaultMsg("服务器繁忙，请稍后再试", msg...)})
}

func ResponseRaw(c *gin.Context, httpCode int32, msg string, interCode ...int32) {
	var iCode int32 = 10
	if len(interCode) > 0 {
		iCode = interCode[0]
	}
	c.AbortWithStatusJSON(int(httpCode), &ErrorStatus{Code: uint32(iCode), Message: defaultMsg(msg)})
}
