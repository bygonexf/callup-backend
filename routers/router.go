package routers

import (
	"github.com/bygonexf/callup-backend/biz/handler/callup"
	"github.com/bygonexf/callup-backend/pkg/logging"
	"github.com/bygonexf/callup-backend/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiV1 := r.Group("/api/v1")
	{
		// 查询
		apiV1.GET("/callup", callup.CallupV1List)
		// 新建
		apiV1.POST("/callup", callup.CallupV1Create)
		// 更新
		apiV1.PATCH("/callup", callup.CallupV1Update)
		// 删除
		apiV1.DELETE("/callup", callup.CallupV1Delete)
	}
	r.GET("/test", func(c *gin.Context) {
		logging.Infoo("sample log")
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}
