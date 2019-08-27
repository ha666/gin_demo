package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"net/http"
	"time"
)

// 插入【用户信息表】信息接口

// 请求
type insertUserInfoRequest struct {

	// 用户编码，取自钉钉编码
	UserCode string `form:"userCode"`

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
func InsertUserInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.InsertUserInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request insertUserInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证userCode参数
	if golibs.Length(request.UserCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.userCode is null",
			Message: "缺少【用户编码，取自钉钉编码】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.UserCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.userCode format err",
			Message: "【用户编码，取自钉钉编码】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.UserCode) > 32 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.userCode length err",
			Message: "【用户编码，取自钉钉编码】参数长度不能超过32个字符",
		})
		return
	}
	//endregion

	//region 验证realName参数
	if golibs.Length(request.RealName) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.realName is null",
			Message: "缺少【姓名】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.RealName) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.realName format err",
			Message: "【姓名】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.RealName) > 16 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.realName length err",
			Message: "【姓名】参数长度不能超过16个字符",
		})
		return
	}
	//endregion

	//region 验证jobNumber参数
	if golibs.Length(request.JobNumber) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.jobNumber is null",
			Message: "缺少【员工的工号】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.JobNumber) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.jobNumber format err",
			Message: "【员工的工号】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.JobNumber) > 16 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.jobNumber length err",
			Message: "【员工的工号】参数长度不能超过16个字符",
		})
		return
	}
	//endregion

	//region 验证jobPosition参数
	if golibs.Length(request.JobPosition) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.jobPosition is null",
			Message: "缺少【职位信息】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.JobPosition) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.jobPosition format err",
			Message: "【职位信息】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.JobPosition) > 16 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.jobPosition length err",
			Message: "【职位信息】参数长度不能超过16个字符",
		})
		return
	}
	//endregion

	//region 验证hiredDate参数
	if golibs.Length(request.HiredDate) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.hiredDate length err",
			Message: "缺少【入职时间】参数",
		})
		return
	}
	request.hiredDateTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.HiredDate, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.hiredDate parse err",
			Message: "【入职时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.hiredDateTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.hiredDate value err",
			Message: "【入职时间】参数值错误:" + request.HiredDate,
		})
		return
	}
	//endregion

	//region 验证avatar参数
	if golibs.Length(request.Avatar) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.avatar is null",
			Message: "缺少【头像url】参数",
		})
		return
	}
	if !golibs.IsGeneralString(request.Avatar) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.avatar format err",
			Message: "【头像url】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.Avatar) > 256 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.avatar length err",
			Message: "【头像url】参数长度不能超过256个字符",
		})
		return
	}
	//endregion

	//region 验证gender参数
	if request.Gender <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.gender value err",
			Message: "gender参数值错误",
		})
		return
	}
	//endregion

	//region 验证userType参数
	if request.UserType <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.userType value err",
			Message: "userType参数值错误",
		})
		return
	}
	//endregion

	//region 插入【用户信息表】信息
	isSuccess, err := userService.Insert(request.UserCode, request.RealName, request.JobNumber, request.JobPosition, request.hiredDateTime, request.Avatar, int8(request.Gender), int8(request.UserType))
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.insert err",
			Message: "插入出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.InsertUserInfo.insert failure",
			Message: "插入失败",
		})
		return
	}
	//endregion

	//region 返回插入【用户信息表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}
