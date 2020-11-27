package callup

import (
	"github.com/gin-gonic/gin"
)

// @Summary 新增召集令
// @Tags 召集令相关
// @Produce json
// @Param data body CallupV1CreateRequest true "request body"
// @Success 200 {object} CallupV1CreateResponse
// @Failure 400 {object} resp.ErrorStatus400
// @Failure 401 {object} resp.ErrorStatus401
// @Failure 404 {object} resp.ErrorStatus404
// @Failure 500 {object} resp.ErrorStatus500
// @Router /api/v1/callup [post]
func CallupV1Create(c *gin.Context) {

}

type CallupV1CreateRequest struct {
	CallerId int64  `json:"caller_id"`
	Type     int32  `json:"type"`
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	Quota    int32  `json:"quota"`
	EndTime  int64  `json:"end_time"`
	PhotoUrl string `json:"photo_url"`
	City     int32  `json:"city"`
}

type CallupV1CreateResponse struct {
	CallupId int64 `json:"callup_id"`
}
