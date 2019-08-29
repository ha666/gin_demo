package proj

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"net/http"
	"strconv"
)

// 获取【项目表】列表接口

// 请求
type getProjListRequest struct {

	// 当前页码
	pageIndex    string
	pageIndexInt int

	// 每页记录数
	pageSize    string
	pageSizeInt int

	// 部分-项目名称
	likeProjName string

	// 用户编码
	userCode string
}

// 方法
func GetProjList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "proj.GetProjList.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getProjListRequest
	//endregion

	//region 验证【当前页码】参数
	request.pageIndexInt = 1
	request.pageIndex = c.DefaultQuery("pageIndex", "")
	if golibs.Length(request.pageIndex) > 0 {
		if !golibs.IsNumber(request.pageIndex) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "proj.GetProjList.pageIndex is number",
				Message: "pageIndex参数格式不正确",
			})
			return
		}
		request.pageIndexInt, err = strconv.Atoi(request.pageIndex)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "proj.GetProjList.pageIndex parse err",
				Message: "pageIndex参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageIndexInt < 1 || request.pageIndexInt > 100000 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "proj.GetProjList.pageIndex value err",
				Message: "pageIndex参数值错误",
			})
			return
		}
	}
	//endregion

	//region 验证【每页记录数】参数
	request.pageSizeInt = 15
	request.pageSize = c.DefaultQuery("pageSize", "")
	if golibs.Length(request.pageSize) > 0 {
		if !golibs.IsNumber(request.pageSize) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "proj.GetProjList.pageSize is number",
				Message: "pageSize参数格式不正确",
			})
			return
		}
		request.pageSizeInt, err = strconv.Atoi(request.pageSize)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "proj.GetProjList.pageSize parse err",
				Message: "pageSize参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageSizeInt < 1 || request.pageSizeInt > 500 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "proj.GetProjList.pageSize value err",
				Message: "pageSize参数值错误",
			})
			return
		}
	}
	//endregion

	//region 验证部分【项目名称】参数,可选
	request.likeProjName = c.DefaultQuery("likeProjName", "")
	if golibs.Length(request.likeProjName) > 0 {
		if !golibs.IsHanOrLetterOrNumber(request.likeProjName) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "proj.GetProjList.likeProjName format err",
				Message: "部分【项目名称】参数格式不正确",
			})
			return
		}
		if golibs.Length(request.likeProjName) > 32 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "proj.GetProjList.likeProjName length err",
				Message: "部分【项目名称】参数长度不能超过32个字符",
			})
			return
		}
	}
	//endregion

	//region 验证【用户编码】参数,可选
	request.userCode = c.DefaultQuery("userCode", "")
	if golibs.Length(request.userCode) > 0 {
		if !golibs.IsLetterOrNumber1(request.userCode) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "proj.GetProjList.userCode format err",
				Message: "【用户编码】参数格式不正确",
			})
			return
		}
		if golibs.Length(request.userCode) > 32 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "proj.GetProjList.userCode length err",
				Message: "【用户编码】参数长度不能超过32个字符",
			})
			return
		}
	}
	//endregion

	//region 查询【项目表】列表
	list, total, err := projService.GetProjList(request.likeProjName, request.userCode, request.pageIndexInt, request.pageSizeInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "proj.GetProjList.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 组装【项目表】列表
	projsArray := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for i, v := range list {
			projsArray[i] = map[string]interface{}{
				"projId":       v.ProjId,                                   //项目编号
				"projName":     v.ProjName,                                 //项目名称
				"userCode":     v.UserCode,                                 //用户编码
				"deleteStatus": v.DeleteStatus,                             //是否删除
				"endTime":      v.EndTime.Format(golibs.Time_TIMEStandard), //结束时间
			}
		}
	}
	//endregion

	//region 返回【项目表】列表
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"total": total,
			"list":  projsArray,
		},
	})
	//endregion
}
