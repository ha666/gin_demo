package proj

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"net/http"
	"strconv"
)

// 删除【项目表】信息接口

// 请求
type deleteProjInfoRequest struct {

	// 项目编号
	projId    string
	projIdInt int
}

// 方法
func DeleteProjInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "proj.DeleteProjInfo.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request deleteProjInfoRequest
	//endregion

	//region 验证projId参数
	request.projId = c.Param("id")
	if golibs.Length(request.projId) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.DeleteProjInfo.projId is null",
			Message: "缺少【项目编号】参数",
		})
		return
	}
	if !golibs.IsNumber(request.projId) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.DeleteProjInfo.projId is number",
			Message: "【项目编号】参数格式不正确",
		})
		return
	}
	request.projIdInt, err = strconv.Atoi(request.projId)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.DeleteProjInfo.projId parse err",
			Message: "【项目编号】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.projIdInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.DeleteProjInfo.projId value err",
			Message: "【项目编号】参数值错误",
		})
		return
	}
	//endregion

	//region 删除【项目表】信息
	isSuccess, err := projService.Delete(request.projIdInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.DeleteProjInfo.delete err",
			Message: "删除出错:" + err.Error(),
		})
		return
	}
	if !isSuccess {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.DeleteProjInfo.delete failure",
			Message: "删除失败",
		})
		return
	}
	//endregion

	//region 返回删除【项目表】结果
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"result": true,
		},
	})
	//endregion
}
