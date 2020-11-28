package callup

import (
	"github.com/gin-gonic/gin"
)

// @Summary 查询召集令
// @Tags 召集令相关
// @Produce json
// @Param page 						query int true "必填 页码"
// @Param page_size 				query int true "必填 页面大小"
// @Param caller_id					query int false "选填 令主id"
// @Param type						query int false "选填 召集令类型"
// @Param fuzzy_name				query string false "选填 召集令名称模糊搜索"
// @Param status					query int false "选填 召集令状态"
// @Param callup_id					query int false "选填 召集令 id"
// @Success 200 {object} CallupV1ListResponse
// @Failure 400 {object} resp.ErrorStatus400
// @Failure 401 {object} resp.ErrorStatus401
// @Failure 404 {object} resp.ErrorStatus404
// @Failure 500 {object} resp.ErrorStatus500
// @Router /api/v1/callup [get]
func CallupV1List(c *gin.Context) {

}

type CallupV1ListRequest struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"page_size"`

	CallerId  int64  `json:"caller_id"`
	Type      int32  `json:"type"`
	FuzzyName string `json:"fuzzy_name"`
	Status    int32  `json:"status"`
	CallupId  int64  `json:"callup_id"`
}

type CallupV1ListResponse struct {
	Total      int32          `json:"total"`
	CallupList []CallupStruct `json:"callup_list"`
}

type CallupStruct struct {
	Id              int64               `json:"id"`
	CallerId        int64               `json:"caller_id"`
	Type            int32               `json:"type"`
	Name            string              `json:"name"`
	Desc            string              `json:"desc"`
	Quota           int32               `json:"quota"`
	EndTime         int64               `json:"end_time"`
	PhotoUrl        string              `json:"photo_url"`
	Status          int32               `json:"status"`
	SuccessNum      int32               `json:"success_num"`
	City            int32               `json:"city"`
	ApplicationList []ApplicationStruct `json:"application_list"`
}

type ApplicationStruct struct {
	Id       int64  `json:"id"`
	CalleeId int64  `json:"callee_id"`
	Desc     string `json:"desc"`
	Status   int32  `json:"status"`
}
