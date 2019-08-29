package service

import (
	"errors"
	"github.com/ha666/gin_demo/dao"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"time"
)

type UserService struct {
	UserDao *dao.UserDao
	ProjDao *dao.ProjDao
}

func NewUserService() *UserService {
	return &UserService{
		UserDao: dao.NewUserDao(),
		ProjDao: dao.NewProjDao(),
	}
}

// 获取【用户信息表】信息
func (d *UserService) GetUserInfo(userCode string) (user model.User, err error) {

	//region 查询【用户信息表】信息
	user, err = d.UserDao.Get(userCode)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【用户编码，取自钉钉编码】查询【用户信息表】信息
func (d *UserService) GetUserInfoByUserCode(userCode string) (user model.User, err error) {

	//region 查询【用户信息表】信息
	user, err = d.UserDao.GetByUserCode(userCode)
	if err != nil {
		return
	}
	//endregion

	return
}

// 根据【用户编码，取自钉钉编码】批量获取【用户信息表】列表
func (d *UserService) GetInUserList(userCodes []string) (users []model.User, count int, err error) {

	//region 查询列表
	users, err = d.UserDao.GetIn(userCodes)
	if err != nil {
		return
	}
	count = len(users)
	//endregion

	return
}

// 根据【姓名】分页查询【用户信息表】列表
func (d *UserService) GetUserListByRealName(realName string, pageIndex, pageSize int) (users []model.User, count int, err error) {

	//region 查询总数
	count, err = d.UserDao.GetRowCountByRealName(realName)
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	users, err = d.UserDao.GetRowListByRealName(realName, pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 分页查询【用户信息表】列表
func (d *UserService) GetUserList(likeRealName string, jobNumber string, jobPosition string, userType int8, startCreateTime time.Time, endCreateTime time.Time, pageIndex, pageSize int) (users []model.User, count int, err error) {

	//region 查询总数
	count, err = d.UserDao.GetRowCount(likeRealName, jobNumber, jobPosition, userType, startCreateTime, endCreateTime)
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	users, err = d.UserDao.GetRowList(likeRealName, jobNumber, jobPosition, userType, startCreateTime, endCreateTime, pageIndex, pageSize)
	if err != nil {
		return
	}
	//endregion

	return
}

// 插入【用户信息表】信息
func (d *UserService) Insert(userCode string, realName string, jobNumber string, jobPosition string, hiredDate time.Time, avatar string, gender int8, userType int8) (bool, error) {

	var err error
	isExist := false
	var userInfo model.User

	//region 根据【用户编码，取自钉钉编码】查询【用户信息表】信息
	{
		userInfo, err = d.UserDao.Get(userCode)
		if err != nil {
			return false, err
		}
		if golibs.Length(userInfo.UserCode) > 0 {
			if userInfo.DeleteStatus == 1 {
				return false, errors.New("【用户编码，取自钉钉编码】已存在,不能重复添加")
			}
			isExist = true
		}
	}
	//endregion

	//region 根据【用户编码，取自钉钉编码】查询【用户信息表】信息
	{
		userInfo, err = d.UserDao.GetByUserCode(userCode)
		if err != nil {
			return false, errors.New("查询【用户信息表】出错:" + err.Error())
		}
		if golibs.Length(userInfo.UserCode) > 0 {
			if userInfo.DeleteStatus == 1 {
				return false, errors.New("【用户编码，取自钉钉编码】已存在,不能重复添加")
			}
			isExist = true
		}
	}
	//endregion

	//region 构造【用户信息表】信息
	userInfo.UserCode = userCode
	userInfo.RealName = realName
	userInfo.JobNumber = jobNumber
	userInfo.JobPosition = jobPosition
	userInfo.HiredDate = hiredDate
	userInfo.Avatar = avatar
	userInfo.Gender = gender
	userInfo.UserType = userType
	userInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【用户信息表】信息
		isSuccess, err := d.UserDao.Update(&userInfo)
		if err != nil {
			return false, errors.New("插入【用户信息表】出错:" + err.Error())
		}
		if !isSuccess {
			return false, errors.New("插入【用户信息表】失败")
		}
		return isSuccess, nil
		//endregion

	} else {

		//region 插入数据库
		isSuccess, err := d.UserDao.Insert(&userInfo)
		if err != nil {
			return false, errors.New("插入【用户信息表】出错:" + err.Error())
		}
		if !isSuccess {
			return false, errors.New("插入【用户信息表】失败")
		}
		return isSuccess, nil
		//endregion

	}

}

//修改【用户信息表】信息
func (d *UserService) Update(userCode string, realName string, jobNumber string, jobPosition string, hiredDate time.Time, avatar string, gender int8, userType int8) (bool, error) {

	//region 查询【用户信息表】信息
	userInfo, err := d.UserDao.Get(userCode)
	if err != nil {
		return false, errors.New("查询【用户信息表】信息出错:" + err.Error())
	}
	if golibs.Length(userInfo.UserCode) <= 0 {
		return false, errors.New("【用户信息表】信息不存在")
	}
	if userInfo.DeleteStatus != 1 {
		return false, errors.New("【用户信息表】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if userInfo.UserCode == userCode &&
		userInfo.RealName == realName &&
		userInfo.JobNumber == jobNumber &&
		userInfo.JobPosition == jobPosition &&
		userInfo.HiredDate == hiredDate &&
		userInfo.Avatar == avatar &&
		userInfo.Gender == gender &&
		userInfo.UserType == userType {
		return true, nil
	}
	//endregion

	//region 根据【用户编码，取自钉钉编码】查询【用户信息表】信息
	{
		_userInfo, err := d.UserDao.GetByUserCode(userCode)
		if err != nil {
			return false, errors.New("查询【用户信息表】出错:" + err.Error())
		}
		if golibs.Length(_userInfo.UserCode) > 0 && _userInfo.UserCode != userInfo.UserCode {
			if _userInfo.DeleteStatus == 1 {
				return false, errors.New("【用户编码，取自钉钉编码】已存在,不能重复添加")
			}
		}
	}
	//endregion

	//region 构造【用户信息表】信息
	userInfo.RealName = realName
	userInfo.JobNumber = jobNumber
	userInfo.JobPosition = jobPosition
	userInfo.HiredDate = hiredDate
	userInfo.Avatar = avatar
	userInfo.Gender = gender
	userInfo.UserType = userType
	userInfo.DeleteStatus = 1
	//endregion

	//region 修改【用户信息表】信息
	return d.UserDao.Update(&userInfo)
	//endregion

}

//删除【用户信息表】信息
func (d *UserService) Delete(userCode string) (isSuccess bool, err error) {

	//region 根据【用户编码】查询【项目表】总数
	{
		count, err := d.ProjDao.GetRowCountByUserCode(userCode)
		if err != nil {
			return false, errors.New("查询【项目表】总数出错:" + err.Error())
		}
		if count > 0 {
			return false, errors.New("【项目表】存在相关记录")
		}
	}
	//endregion

	//region 查询【用户信息表】信息
	{
		userInfo, err := d.UserDao.Get(userCode)
		if err != nil {
			return false, errors.New("查询【用户信息表】信息出错:" + err.Error())
		}
		if golibs.Length(userInfo.UserCode) <= 0 {
			return false, errors.New("没有找到【用户信息表】信息")
		}
		if userInfo.DeleteStatus != 1 {
			return false, errors.New("【用户信息表】信息已被删除")
		}
	}
	//endregion

	//region 删除【用户信息表】信息
	return d.UserDao.Delete(userCode)
	//endregion

}
