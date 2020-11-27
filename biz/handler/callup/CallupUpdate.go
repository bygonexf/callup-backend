package callup

import (
	"github.com/gin-gonic/gin"
)

// @Summary 更新召集令
// @Tags 召集令相关
// @Produce json
// @Param callup_id path int true "id"
// @Param data body CallupV1UpdateBodyRequest true "request body"
// @Success 200 {object} CallupV1UpdateResponse
// @Failure 400 {object} resp.ErrorStatus400
// @Failure 401 {object} resp.ErrorStatus401
// @Failure 404 {object} resp.ErrorStatus404
// @Failure 500 {object} resp.ErrorStatus500
// @Router /api/v1/callup/{callup_id} [patch]
func CallupV1Update(c *gin.Context) {

}

type CallupV1UpdateRequest struct {
	CallupId string `uri:"callup_id" binding:"gt=0"` // 必填
}

type CallupV1UpdateBodyRequest struct {
	Type       int32  `json:"type"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Quota      int32  `json:"quota"`
	EndTime    int64  `json:"end_time"`
	PhotoUrl   string `json:"photo_url"`
	IsCanceled int32  `json:"is_canceled"`
	City       int32  `json:"city"`
}

type CallupV1UpdateResponse struct {
}
