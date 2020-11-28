package application

import (
	handler2 "github.com/bygonexf/callup-backend/biz/handler/callup"
	"github.com/gin-gonic/gin"
)

// @Summary 查询请求
// @Tags 请求相关
// @Produce json
// @Param page 						query int true "必填 页码"
// @Param page_size 				query int true "必填 页面大小"
// @Param callup_id					query int false "选填 召集令 id"
// @Param callee_id					query int false "选填 请求者 id"
// @Param status					query int false "选填 请求状态"
// @Param application_id 			query int false "选填 申请 Id"
// @Success 200 {object} ApplicationV1ListResponse
// @Failure 400 {object} resp.ErrorStatus400
// @Failure 401 {object} resp.ErrorStatus401
// @Failure 404 {object} resp.ErrorStatus404
// @Failure 500 {object} resp.ErrorStatus500
// @Router /api/v1/application [get]
func ApplicationV1List(c *gin.Context) {

}

type ApplicationV1ListRequest struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"page_size"`

	CallupId      int64 `json:"callup_id"`
	CalleeId      int64 `json:"callee_id"`
	Status        int32 `json:"status"`
	ApplicationId int64 `json:"application_id"`
}

type ApplicationV1ListResponse struct {
	Total           int32                        `json:"total"`
	ApplicationList []handler2.ApplicationStruct `json:"application_list"`
}
