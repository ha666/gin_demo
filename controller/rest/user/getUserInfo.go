package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"net/http"
)

// 获取【用户信息表】信息接口

// 请求
type getUserInfoRequest struct {

	// 用户编码，取自钉钉编码
	userCode string
}

// 方法
func GetUserInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getUserInfoRequest
	//endregion

	//region 验证userCode参数
	request.userCode = c.Param("id")
	if golibs.Length(request.userCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.GetUserInfo.userCode is null",
			Message: "缺少【用户编码，取自钉钉编码】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.userCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.GetUserInfo.userCode format err",
			Message: "【用户编码，取自钉钉编码】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.userCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.GetUserInfo.userCode value err",
			Message: "【用户编码，取自钉钉编码】参数值超长了",
		})
		return
	}
	//endregion

	//region 查询【用户信息表】信息
	userInfo, err := userService.GetUserInfo(request.userCode)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.GetUserInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if golibs.Length(userInfo.UserCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.GetUserInfo.not found",
			Message: "没有找到【用户信息表】信息",
		})
		return
	}
	if userInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.GetUserInfo.has delete",
			Message: "【用户信息表】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【用户信息表】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"userCode":     userInfo.UserCode,                                    //用户编码，取自钉钉编码
			"realName":     userInfo.RealName,                                    //姓名
			"jobNumber":    userInfo.JobNumber,                                   //员工的工号
			"jobPosition":  userInfo.JobPosition,                                 //职位信息
			"hiredDate":    userInfo.HiredDate.Format(golibs.Time_TIMEStandard),  //入职时间
			"avatar":       userInfo.Avatar,                                      //头像url
			"gender":       userInfo.Gender,                                      //性别，0未知，1男，2女
			"userType":     userInfo.UserType,                                    //用户类型
			"deleteStatus": userInfo.DeleteStatus,                                //删除状态，0未知，1未删除，2删除
			"createTime":   userInfo.CreateTime.Format(golibs.Time_TIMEStandard), //创建时间
			"updateTime":   userInfo.UpdateTime.Format(golibs.Time_TIMEStandard), //修改时间
		},
	})
	//endregion
}
