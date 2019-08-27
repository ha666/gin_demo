package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"net/http"
	"strconv"
	"time"
)

// 获取【用户信息表】列表接口

// 请求
type getUserListRequest struct {

	// 当前页码
	pageIndex    string
	pageIndexInt int

	// 每页记录数
	pageSize    string
	pageSizeInt int

	// 部分-姓名
	likeRealName string

	// 员工的工号
	jobNumber string

	// 职位信息
	jobPosition string

	// 用户类型
	userType    string
	userTypeInt int

	// 开始-创建时间
	startCreateTime     string
	startCreateTimeTime time.Time

	// 结束-创建时间
	endCreateTime     string
	endCreateTimeTime time.Time
}

// 方法
func GetUserList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.ex",
				Message: fmt.Sprintf("系统错误:%v", err),
			})
			return
		}
	}()

	//region 解析请求参数
	var err error
	var request getUserListRequest
	//endregion

	//region 验证【当前页码】参数
	request.pageIndexInt = 1
	request.pageIndex = c.DefaultQuery("pageIndex", "")
	if golibs.Length(request.pageIndex) > 0 {
		if !golibs.IsNumber(request.pageIndex) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.pageIndex is number",
				Message: "pageIndex参数格式不正确",
			})
			return
		}
		request.pageIndexInt, err = strconv.Atoi(request.pageIndex)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.pageIndex parse err",
				Message: "pageIndex参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageIndexInt < 1 || request.pageIndexInt > 100000 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.pageIndex value err",
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
				Code:    "user.GetUserList.pageSize is number",
				Message: "pageSize参数格式不正确",
			})
			return
		}
		request.pageSizeInt, err = strconv.Atoi(request.pageSize)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.pageSize parse err",
				Message: "pageSize参数解析出错:" + err.Error(),
			})
			return
		}
		if request.pageSizeInt < 1 || request.pageSizeInt > 500 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.pageSize value err",
				Message: "pageSize参数值错误",
			})
			return
		}
	}
	//endregion

	//region 验证部分【姓名】参数,可选
	request.likeRealName = c.DefaultQuery("likeRealName", "")
	if golibs.Length(request.likeRealName) > 0 {
		if !golibs.IsHanOrLetterOrNumber(request.likeRealName) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.likeRealName format err",
				Message: "部分【姓名】参数格式不正确",
			})
			return
		}
		if golibs.Length(request.likeRealName) > 16 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.likeRealName length err",
				Message: "部分【姓名】参数长度不能超过16个字符",
			})
			return
		}
	}
	//endregion

	//region 验证【员工的工号】参数,可选
	request.jobNumber = c.DefaultQuery("jobNumber", "")
	if golibs.Length(request.jobNumber) > 0 {
		if !golibs.IsLetterOrNumber1(request.jobNumber) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.jobNumber format err",
				Message: "【员工的工号】参数格式不正确",
			})
			return
		}
		if golibs.Length(request.jobNumber) > 16 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.jobNumber length err",
				Message: "【员工的工号】参数长度不能超过16个字符",
			})
			return
		}
	}
	//endregion

	//region 验证【职位信息】参数,可选
	request.jobPosition = c.DefaultQuery("jobPosition", "")
	if golibs.Length(request.jobPosition) > 0 {
		if !golibs.IsHanOrLetterOrNumber(request.jobPosition) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.jobPosition format err",
				Message: "【职位信息】参数格式不正确",
			})
			return
		}
		if golibs.Length(request.jobPosition) > 16 {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.jobPosition length err",
				Message: "【职位信息】参数长度不能超过16个字符",
			})
			return
		}
	}
	//endregion

	//region 验证【用户类型】参数,必选
	request.userType = c.DefaultQuery("userType", "")
	if golibs.Length(request.userType) <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.GetUserList.userType is not null",
			Message: "【用户类型】参数不能为空",
		})
		return
	}
	if !golibs.IsNumber(request.userType) {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.GetUserList.userType is not a number",
			Message: "【用户类型】参数格式不正确",
		})
		return
	}
	request.userTypeInt, err = strconv.Atoi(request.userType)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.GetUserList.userType parse err",
			Message: "【用户类型】参数解析出错:" + err.Error(),
		})
		return
	}
	if request.userTypeInt <= 0 {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.GetUserList.userType value err",
			Message: "【用户类型】参数值错误",
		})
		return
	}
	//endregion

	//region 验证开始【创建时间】参数,可选
	request.startCreateTime = c.DefaultQuery("startCreateTime", "")
	if golibs.Length(request.startCreateTime) > 0 {
		if !golibs.IsStandardTime(request.startCreateTime) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.startCreateTime format err",
				Message: "开始【startCreateTime】参数格式不正确",
			})
			return
		}
		request.startCreateTimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.startCreateTime, time.Local)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.startCreateTime parse err",
				Message: "开始【创建时间】参数解析错误:" + err.Error(),
			})
			return
		}
	}
	//endregion

	//region 验证结束【创建时间】参数,可选
	request.endCreateTime = c.DefaultQuery("endCreateTime", "")
	if golibs.Length(request.endCreateTime) > 0 {
		if !golibs.IsStandardTime(request.endCreateTime) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.endCreateTime format err",
				Message: "结束【endCreateTime】参数格式不正确",
			})
			return
		}
		request.endCreateTimeTime, err = time.ParseInLocation(golibs.Time_TIMEStandard, request.endCreateTime, time.Local)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.endCreateTime parse err",
				Message: "结束【创建时间】参数解析错误:" + err.Error(),
			})
			return
		}
		if request.startCreateTimeTime.Before(request.endCreateTimeTime) {
			c.JSON(http.StatusOK, model.Response{
				Code:    "user.GetUserList.endCreateTime value err",
				Message: "【创建时间】开始时间必须在结束时间以后",
			})
			return
		}
	}
	//endregion

	//region 查询【用户信息表】列表
	list, total, err := userService.GetUserList(request.likeRealName, request.jobNumber, request.jobPosition, int8(request.userTypeInt), request.startCreateTimeTime, request.endCreateTimeTime, request.pageIndexInt, request.pageSizeInt)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code:    "user.GetUserList.query err",
			Message: "查询出错:" + err.Error(),
		})
		return
	}
	//endregion

	//region 组装【用户信息表】列表
	usersArray := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for i, v := range list {
			usersArray[i] = map[string]interface{}{
				"userCode":     v.UserCode,                                    //用户编码，取自钉钉编码
				"realName":     v.RealName,                                    //姓名
				"jobNumber":    v.JobNumber,                                   //员工的工号
				"jobPosition":  v.JobPosition,                                 //职位信息
				"hiredDate":    v.HiredDate.Format(golibs.Time_TIMEStandard),  //入职时间
				"avatar":       v.Avatar,                                      //头像url
				"gender":       v.Gender,                                      //性别，0未知，1男，2女
				"userType":     v.UserType,                                    //用户类型
				"deleteStatus": v.DeleteStatus,                                //删除状态，0未知，1未删除，2删除
				"createTime":   v.CreateTime.Format(golibs.Time_TIMEStandard), //创建时间
				"updateTime":   v.UpdateTime.Format(golibs.Time_TIMEStandard), //修改时间
			}
		}
	}
	//endregion

	//region 返回【用户信息表】列表
	c.JSON(http.StatusOK, model.Response{
		Code:    "ok",
		Message: "",
		Data: map[string]interface{}{
			"total": total,
			"list":  usersArray,
		},
	})
	//endregion
}
