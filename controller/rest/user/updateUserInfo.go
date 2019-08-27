package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"net/http"
	"time"
)

// 修改【用户信息表】信息接口

// 请求
type updateUserInfoRequest struct {

	// 用户编码，取自钉钉编码
	userCode string

	// 姓名
	RealName string `form:"realName"`

	// 员工的工号
	JobNumber string `form:"jobNumber"`

	// 职位信息
	JobPosition string `form:"jobPosition"`

	// 入职时间
	HiredDate     string `form:"hiredDate"`
	hiredDateTime time.Time

	// 头像url
	Avatar string `form:"avatar"`

	// 性别，0未知，1男，2女
	Gender int `form:"gender"`

	// 用户类型
	UserType int `form:"userType"`
}

// 方法
func UpdateUserInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.UpdateUserInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request updateUserInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证realName参数
	if golibs.Length(request.RealName) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.realName is null",
			Message: "缺少【姓名】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.RealName) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.realName format err",
			Message: "【姓名】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.RealName) > 16 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.realName length err",
			Message: "【姓名】参数长度不能超过16个字符",
		})
		return
	}
	//endregion

	//region 验证jobNumber参数
	if golibs.Length(request.JobNumber) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.jobNumber is null",
			Message: "缺少【员工的工号】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.JobNumber) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.jobNumber format err",
			Message: "【员工的工号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.JobNumber) > 16 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.jobNumber length err",
			Message: "【员工的工号】参数长度不能超过16个字符",
		})
		return
	}
	//endregion

	//region 验证jobPosition参数
	if golibs.Length(request.JobPosition) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.jobPosition is null",
			Message: "缺少【职位信息】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.JobPosition) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.jobPosition format err",
			Message: "【职位信息】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.JobPosition) > 16 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.jobPosition length err",
			Message: "【职位信息】参数长度不能超过16个字符",
		})
		return
	}
	//endregion

	//region 验证hiredDate参数
	if golibs.Length(request.HiredDate) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.hiredDate length err",
			Message: "缺少【入职时间】参数",
		})
		return
	}
	request.hiredDateTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.HiredDate, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.hiredDate parse err",
			Message: "【入职时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.hiredDateTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.hiredDate value err",
			Message: "【入职时间】参数值错误:" + request.HiredDate,
		})
		return
	}
	//endregion

	//region 验证avatar参数
	if golibs.Length(request.Avatar) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.avatar is null",
			Message: "缺少【头像url】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Avatar) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.avatar format err",
			Message: "【头像url】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Avatar) > 256 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.avatar length err",
			Message: "【头像url】参数长度不能超过256个字符",
		})
		return
	}
	//endregion

	//region 验证gender参数
	if request.Gender <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.gender value err",
			Message: "【性别，0未知，1男，2女】参数值错误",
		})
		return
	}
	//endregion

	//region 验证userType参数
	if request.UserType <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.userType value err",
			Message: "【用户类型】参数值错误",
		})
		return
	}
	//endregion

	//region 修改【用户信息表】信息
	isSuccess, err := userService.Update(request.userCode, request.RealName, request.JobNumber, request.JobPosition, request.hiredDateTime, request.Avatar, int8(request.Gender), int8(request.UserType))
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.update err",
			Message: "修改出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.UpdateUserInfo.update failure",
			Message: "修改失败",
		})
		return
	}
	//endregion

	//region 返回修改【用户信息表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}
