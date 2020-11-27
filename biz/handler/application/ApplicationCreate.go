package application

import (
	"github.com/bygonexf/callup-backend/pkg/resp"
	"github.com/gin-gonic/gin"
)

// @Summary 新增请求
// @Tags 请求相关
// @Produce json
// @Param data body ApplicationV1CreateRequest true "request body"
// @Success 200 {object} ApplicationV1CreateResponse
// @Failure 400 {object} resp.ErrorStatus400
// @Failure 401 {object} resp.ErrorStatus401
// @Failure 404 {object} resp.ErrorStatus404
// @Failure 500 {object} resp.ErrorStatus500
// @Router /api/v1/application [post]
func ApplicationV1Create(c *gin.Context) {
	resp.ResponseBadRequest400(c, "")
}

type ApplicationV1CreateRequest struct {
	CallupId int64  `json:"callup_id"`
	CalleeId int64  `json:"callee_id"`
	Desc     string `json:"desc"`
}

type ApplicationV1CreateResponse struct {
	ApplicationId int64 `json:"application_id"`
}
