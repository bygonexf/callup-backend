package application

import (
	"github.com/gin-gonic/gin"
)

// @Summary 删除请求
// @Tags 请求相关
// @Produce json
// @Param data body ApplicationV1DeleteRequest true "request body"
// @Success 200 {object} ApplicationV1DeleteResponse
// @Failure 400 {object} resp.ErrorStatus400
// @Failure 401 {object} resp.ErrorStatus401
// @Failure 404 {object} resp.ErrorStatus404
// @Failure 500 {object} resp.ErrorStatus500
// @Router /api/v1/application [delete]
func ApplicationV1Delete(c *gin.Context) {

}

type ApplicationV1DeleteRequest struct {
	ApplicationId int64 `json:"application_id"`
}

type ApplicationV1DeleteResponse struct {
}
