package callup

import (
	"github.com/gin-gonic/gin"
)

// @Summary 删除召集令
// @Tags 召集令相关
// @Produce json
// @Param data body CallupV1DeleteRequest true "request body"
// @Success 200 {object} CallupV1DeleteResponse
// @Failure 400 {object} resp.ErrorStatus400
// @Failure 401 {object} resp.ErrorStatus401
// @Failure 404 {object} resp.ErrorStatus404
// @Failure 500 {object} resp.ErrorStatus500
// @Router /api/v1/callup [delete]
func CallupV1Delete(c *gin.Context) {

}

type CallupV1DeleteRequest struct {
	CallupId int64 `json:"callup_id"`
}

type CallupV1DeleteResponse struct {
}
