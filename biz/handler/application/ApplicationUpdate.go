package application

import (
	"github.com/gin-gonic/gin"
)

// @Summary 更新请求
// @Tags 请求相关
// @Produce json
// @Param application_id path int true "id"
// @Param data body ApplicationV1UpdateBodyRequest true "request body"
// @Success 200 {object} ApplicationV1UpdateResponse
// @Failure 400 {object} resp.ErrorStatus400
// @Failure 401 {object} resp.ErrorStatus401
// @Failure 404 {object} resp.ErrorStatus404
// @Failure 500 {object} resp.ErrorStatus500
// @Router /api/v1/application/{application_id} [patch]
func ApplicationV1Update(c *gin.Context) {

}

type ApplicationV1UpdateRequest struct {
	ApplicationId string `uri:"application_id" binding:"gt=0"` // 必填
}

type ApplicationV1UpdateBodyRequest struct {
	Desc       string `json:"desc"`
	IsCanceled int32  `json:"is_canceled"`
}

type ApplicationV1UpdateResponse struct {
}
