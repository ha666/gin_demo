package proj

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"net/http"
	"strconv"
)

// 获取【项目表】信息接口

// 请求
type getProjInfoRequest struct {

	// 项目编号
	projId    string
	projIdInt int
}

// 方法
func GetProjInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "proj.GetProjInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getProjInfoRequest
	//endregion

	//region 验证projId参数
	request.projId = c.Param("id")
	if golibs.Length(request.projId) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.GetProjInfo.projId is null",
			Message: "缺少【项目编号】参数",
		})
		return
	}
	if !golibs.IsNumber(request.projId) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.GetProjInfo.projId is number",
			Message: "【项目编号】参数格式不正确",
		})
		return
	}
	request.projIdInt, err = strconv.Atoi(request.projId)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.GetProjInfo.projId parse err",
			Message: "projId参数解析出错:" + err.Error(),
		})
		return
	}
	if request.projIdInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.GetProjInfo.projId value err",
			Message: "【项目编号】参数值错误",
		})
		return
	}
	//endregion

	//region 查询【项目表】信息
	projInfo, err := projService.GetProjInfo(request.projIdInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.GetProjInfo.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	if projInfo.ProjId <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.GetProjInfo.not found",
			Message: "没有找到【项目表】信息",
		})
		return
	}
	if projInfo.DeleteStatus == 2 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.GetProjInfo.has delete",
			Message: "【项目表】信息已被删除",
		})
		return
	}
	//endregion

	//region 返回【项目表】信息
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"projId":       projInfo.ProjId,                                   //项目编号
			"projName":     projInfo.ProjName,                                 //项目名称
			"userCode":     projInfo.UserCode,                                 //用户编码
			"deleteStatus": projInfo.DeleteStatus,                             //是否删除
			"endTime":      projInfo.EndTime.Format(golibs.Time_TIMEStandard), //结束时间
		},
	})
	//endregion
}
