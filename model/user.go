package model

import (
	"time"
)

// 用户信息表
type User struct {
	UserCode     string    //用户编码，取自钉钉编码
	RealName     string    //姓名
	JobNumber    string    //员工的工号
	JobPosition  string    //职位信息
	HiredDate    time.Time //入职时间
	Avatar       string    //头像url
	Gender       int8      //性别，0未知，1男，2女
	UserType     int8      //用户类型
	DeleteStatus int8      //删除状态，0未知，1未删除，2删除
	CreateTime   time.Time //创建时间
	UpdateTime   time.Time //修改时间
}
