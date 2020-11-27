package routers

import (
	"github.com/bygonexf/callup-backend/biz/handler"
	"github.com/bygonexf/callup-backend/biz/handler/application"
	"github.com/bygonexf/callup-backend/biz/handler/callup"
	"github.com/bygonexf/callup-backend/biz/handler/user"
	_ "github.com/bygonexf/callup-backend/docs"
	"github.com/bygonexf/callup-backend/pkg/logging"
	"github.com/bygonexf/callup-backend/pkg/setting"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
		apiV1.PATCH("/callup/:callup_id", callup.CallupV1Update)
		// 删除
		apiV1.DELETE("/callup", callup.CallupV1Delete)

		// 查询
		apiV1.GET("/application", application.ApplicationV1List)
		// 新建
		apiV1.POST("/application", application.ApplicationV1Create)
		// 更新
		apiV1.PATCH("/application/:application_id", application.ApplicationV1Update)
		// 删除
		apiV1.DELETE("/application", application.ApplicationV1Delete)

		// 查询
		apiV1.GET("/user", user.UserV1List)
		// 新建
		apiV1.POST("/user", user.UserV1Create)
		// 更新
		apiV1.PATCH("/user/:user_id", user.UserV1Update)

		// 查询中介费收入
		apiV1.GET("/income", handler.IncomeV1Query)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/test", func(c *gin.Context) {
		logging.Infoo("sample log")
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}
