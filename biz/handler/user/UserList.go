package user

import (
	"github.com/gin-gonic/gin"
)

// @Summary 查询用户
// @Tags 用户相关
// @Produce json
// @Param page 						query int true "必填 页码"
// @Param page_size 				query int true "必填 页面大小"
// @Param admin_type				query int false "选填 用户类型"
// @Param level						query int false "选填 用户级别"
// @Param city						query int false "选填 城市"
// @Param user_id					query int false "选填 用户 id"
// @Success 200 {object} UserV1ListResponse
// @Failure 400 {object} resp.ErrorStatus400
// @Failure 401 {object} resp.ErrorStatus401
// @Failure 404 {object} resp.ErrorStatus404
// @Failure 500 {object} resp.ErrorStatus500
// @Router /api/v1/user [get]
func UserV1List(c *gin.Context) {

}

type UserV1ListRequest struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"page_size"`

	AdminType int32 `json:"admin_type"`
	Level     int32 `json:"level"`
	City      int32 `json:"city"`
	UserId    int64 `json:"user_id"`
}

type UserV1ListResponse struct {
	Total int32 `json:"total"`

	UserList []UserStruct `json:"user_list"`
}

type UserStruct struct {
	Id               int64  `json:"id"`
	SsoName          string `json:"sso_name"`
	AdminType        int32  `json:"admin_type"`
	Name             string `json:"name"`
	Level            int32  `json:"level"`
	CredentialNumber string `json:"credential_number"`
	Phone            string `json:"phone"`
	Desc             string `json:"desc"`
	City             int32  `json:"city"`
}
