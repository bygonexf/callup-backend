package user

import "github.com/gin-gonic/gin"

// @Summary 新增用户
// @Tags 用户相关
// @Produce json
// @Param data body UserV1CreateRequest true "request body"
// @Success 200 {object} UserV1CreateResponse
// @Failure 400 {object} resp.ErrorStatus400
// @Failure 401 {object} resp.ErrorStatus401
// @Failure 404 {object} resp.ErrorStatus404
// @Failure 500 {object} resp.ErrorStatus500
// @Router /api/v1/user [post]
func UserV1Create(c *gin.Context) {

}

type UserV1CreateRequest struct {
	SsoName          string `json:"sso_name"`
	Password         string `json:"password"`
	Name             string `json:"name"`
	CredentialNumber string `json:"credential_number"`
	Phone            string `json:"phone"`
	Desc             string `json:"desc"`
	City             int32  `json:"city"`
}

type UserV1CreateResponse struct {
	Id int64 `json:"id"`
}
