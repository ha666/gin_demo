package proj

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"net/http"
	"time"
)

// 插入【项目表】信息接口

// 请求
type insertProjInfoRequest struct {

	// 项目名称
	ProjName string `form:"projName"`

	// 用户编码
	UserCode string `form:"userCode"`

	// 结束时间
	EndTime     string `form:"endTime"`
	endTimeTime time.Time
}

// 方法
func InsertProjInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "proj.InsertProjInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var request insertProjInfoRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.InsertProjInfo.request err",
			Message: "请求出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 验证projName参数
	if golibs.Length(request.ProjName) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.InsertProjInfo.projName is null",
			Message: "缺少【项目名称】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.ProjName) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.InsertProjInfo.projName format err",
			Message: "【项目名称】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.ProjName) > 32 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.InsertProjInfo.projName length err",
			Message: "【项目名称】参数长度不能超过32个字符",
		})
		return
	}
	//endregion

	//region 验证userCode参数
	if golibs.Length(request.UserCode) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.InsertProjInfo.userCode is null",
			Message: "缺少【用户编码】参数",
		})
		return
	}
	if !golibs.IsLetterOrNumber1(request.UserCode) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.InsertProjInfo.userCode format err",
			Message: "【用户编码】参数格式不正确",
		})
		return
	}
	if golibs.Length(request.UserCode) > 32 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.InsertProjInfo.userCode length err",
			Message: "【用户编码】参数长度不能超过32个字符",
		})
		return
	}
	//endregion

	//region 验证endTime参数
	if golibs.Length(request.EndTime) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.InsertProjInfo.endTime length err",
			Message: "缺少【结束时间】参数",
		})
		return
	}
	request.endTimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.EndTime, time.Local)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.InsertProjInfo.endTime parse err",
			Message: "【结束时间】参数解析错误:" + err.Error(),
		})
		return
	}
	if request.endTimeTime.Before(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local)) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.InsertProjInfo.endTime value err",
			Message: "【结束时间】参数值错误:" + request.EndTime,
		})
		return
	}
	//endregion

	//region 插入【项目表】信息
	projId, err := projService.Insert(request.ProjName, request.UserCode, request.endTimeTime)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.InsertProjInfo.insert err",
			Message: "插入出错:" + err.Error(),
		})
		return
	}
	if projId <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.InsertProjInfo.insert failure",
			Message: "插入失败",
		})
		return
	}
	//endregion

	//region 返回插入【项目表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
			"projId": projId,
		},
	})
	//endregion
}
