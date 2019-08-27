package model

import (
	"time"
)

// 项目表
type Proj struct {
	ProjId       int       //项目编号
	ProjName     string    //项目名称
	UserCode     string    //用户编码
	RealName     string    //姓名【外键表:用户信息表】
	DeleteStatus int8      //是否删除
	EndTime      time.Time //结束时间
}
