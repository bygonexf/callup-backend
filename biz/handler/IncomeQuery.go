package handler

import (
	"github.com/gin-gonic/gin"
)

// @Summary 查询中介收入
// @Tags 其他
// @Produce json
// @Param data body IncomeV1QueryRequest true "request body"
// @Success 200 {object} IncomeV1QueryResponse
// @Failure 400 {object} resp.ErrorStatus400
// @Failure 401 {object} resp.ErrorStatus401
// @Failure 404 {object} resp.ErrorStatus404
// @Failure 500 {object} resp.ErrorStatus500
// @Router /api/v1/income [get]
func IncomeV1Query(c *gin.Context) {

}

type IncomeV1QueryRequest struct {
	BeginTime  int64 `json:"begin_time"`
	EndTime    int64 `json:"end_time"`
	CallupType int32 `json:"callup_type"`
	City       int32 `json:"city"`
}

type IncomeV1QueryResponse struct {
	TotalIncome int32                `json:"total_income"`
	RecordList  []RecordStruct       `json:"record_list"`
	MonthMap    map[string]MonthData `json:"month_map"`
}

type RecordStruct struct {
	Id            int64 `json:"id"`
	CallupId      int64 `json:"callup_id"`
	ApplicationId int64 `json:"application_id"`
	CallerId      int64 `json:"caller_id"`
	CalleeId      int64 `json:"callee_id"`
	SucceedTime   int64 `json:"succeed_time"`
}

type MonthData struct {
	RecordNum int32 `json:"record_num"`
	Income    int32 `json:"income"`
}
