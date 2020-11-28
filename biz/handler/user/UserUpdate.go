package user

import "github.com/gin-gonic/gin"

// @Summary 更新用户信息
// @Tags 用户相关
// @Produce json
// @Param user_id path int true "id"
// @Param data body UserV1UpdateBodyRequest true "request body"
// @Success 200 {object} UserV1UpdateResponse
// @Failure 400 {object} resp.ErrorStatus400
// @Failure 401 {object} resp.ErrorStatus401
// @Failure 404 {object} resp.ErrorStatus404
// @Failure 500 {object} resp.ErrorStatus500
// @Router /api/v1/user/{user_id} [patch]
func UserV1Update(c *gin.Context) {

}

type UserV1UpdateRequest struct {
	UserId int64 `uri:"user_id" binding:"gt=0"` // 必填
}

type UserV1UpdateBodyRequest struct {
	SsoName          string `json:"sso_name"`
	Password         string `json:"password"`
	Name             string `json:"name"`
	CredentialNumber string `json:"credential_number"`
	Phone            string `json:"phone"`
	Desc             string `json:"desc"`
	City             int32  `json:"city"`
}

type UserV1UpdateResponse struct {
}
