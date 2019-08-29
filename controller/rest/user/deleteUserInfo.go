package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"net/http"
)

// 删除【用户信息表】信息接口

// 请求
type deleteUserInfoRequest struct {

	// 用户编码，取自钉钉编码
	userCode string
}

// 方法
func DeleteUserInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.DeleteUserInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteUserInfoRequest
	//endregion

	//region 验证userCode参数
	request.userCode = c.Param("id")
	if golibs.Length(request.userCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.userCode is null",
			Message: "缺少【用户编码，取自钉钉编码】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.userCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.userCode format err",
			Message: "【用户编码，取自钉钉编码】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.userCode) > 32 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.userCode length err",
			Message: "【用户编码，取自钉钉编码】参数长度不能超过32个字符",
		})
		return
	}
	//endregion

	//region 删除【用户信息表】信息
	isSuccess, err := userService.Delete(request.userCode)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.DeleteUserInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.DeleteUserInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【用户信息表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}
