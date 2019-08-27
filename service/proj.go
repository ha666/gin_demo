package service

import (
	"errors"
	"github.com/ha666/gin_demo/dao"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"time"
)

type ProjService struct {
	ProjDao *dao.ProjDao
	UserDao *dao.UserDao
}

func NewProjService() *ProjService {
	return &ProjService{
		ProjDao: dao.NewProjDao(),
		UserDao: dao.NewUserDao(),
	}
}

// 获取【项目表】信息
func (d *ProjService) GetProjInfo(projId int) (proj model.Proj, err error) {

	//region 查询【项目表】信息
	proj, err = d.ProjDao.Get(projId)
	if err != nil {
		return
	}
	//endregion

	//region 获取【用户信息表】信息
	if golibs.Length(proj.UserCode) > 0 {
		var user model.User
		user, err = d.UserDao.Get(proj.UserCode)
		if err != nil {
			return
		}
		proj.RealName = user.RealName
	}
	//endregion

	return
}

// 根据【项目编号】批量获取【项目表】列表
func (d *ProjService) GetInProjList(projIds []int) (projs []model.Proj, count int, err error) {

	//region 查询列表
	projs, err = d.ProjDao.GetIn(projIds)
	if err != nil {
		return
	}
	count = len(projs)
	if len(projs) > 0 && projs[0].ProjId > 0 {

		//region 查询外键表【用户信息表】列表
		userCodes := make([]string, 0, len(projs))
		var users []model.User
		for _, proj := range projs {
			if golibs.Length(proj.UserCode) > 0 {
				userCodes = append(userCodes, proj.UserCode)
			}
		}
		if len(userCodes) > 0 {
			users, err = d.UserDao.GetIn(userCodes)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, proj := range projs {

			//region 处理【用户信息表】列表
			if len(users) > 0 {
				for _, user := range users {
					if proj.UserCode == user.UserCode {
						projs[i].RealName = user.RealName
						break
					}
				}
			}
			//endregion

		}
		//endregion

	}
	//endregion

	return
}

// 根据【项目名称】分页查询【项目表】列表
func (d *ProjService) GetProjListByProjName(projName string, pageIndex, pageSize int) (projs []model.Proj, count int, err error) {

	//region 查询总数
	count, err = d.ProjDao.GetRowCountByProjName(projName)
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	projs, err = d.ProjDao.GetRowListByProjName(projName, pageIndex, pageSize)
	if err != nil {
		return
	}
	if len(projs) > 0 && projs[0].ProjId > 0 {

		//region 查询外键表【用户信息表】列表
		userCodes := make([]string, 0, len(projs))
		var users []model.User
		for _, proj := range projs {
			if golibs.Length(proj.UserCode) > 0 {
				userCodes = append(userCodes, proj.UserCode)
			}
		}
		if len(userCodes) > 0 {
			users, err = d.UserDao.GetIn(userCodes)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, proj := range projs {

			//region 处理【用户信息表】列表
			if len(users) > 0 {
				for _, user := range users {
					if proj.UserCode == user.UserCode {
						projs[i].RealName = user.RealName
						break
					}
				}
			}
			//endregion

		}
		//endregion

	}
	//endregion

	return
}

// 分页查询【项目表】列表
func (d *ProjService) GetProjList(pageIndex, pageSize int) (projs []model.Proj, count int, err error) {

	//region 查询总数
	count, err = d.ProjDao.GetRowCount()
	if err != nil || count <= 0 {
		return
	}
	//endregion

	//region 查询列表
	projs, err = d.ProjDao.GetRowList(pageIndex, pageSize)
	if err != nil {
		return
	}
	if len(projs) > 0 && projs[0].ProjId > 0 {

		//region 查询外键表【用户信息表】列表
		userCodes := make([]string, 0, len(projs))
		var users []model.User
		for _, proj := range projs {
			if golibs.Length(proj.UserCode) > 0 {
				userCodes = append(userCodes, proj.UserCode)
			}
		}
		if len(userCodes) > 0 {
			users, err = d.UserDao.GetIn(userCodes)
			if err != nil {
				return
			}
		}
		//endregion

		//region 处理外键表
		for i, proj := range projs {

			//region 处理【用户信息表】列表
			if len(users) > 0 {
				for _, user := range users {
					if proj.UserCode == user.UserCode {
						projs[i].RealName = user.RealName
						break
					}
				}
			}
			//endregion

		}
		//endregion

	}
	//endregion

	return
}

// 插入【项目表】信息
func (d *ProjService) Insert(projName string, userCode string, endTime time.Time) (int, error) {

	var err error
	isExist := false
	var projInfo model.Proj

	//region 判断userCode值是否存在
	user, err := d.UserDao.Get(userCode)
	if err != nil {
		return -1, err
	}
	if golibs.Length(user.UserCode) <= 0 {
		return -1, errors.New("userCode值不存在")
	}
	if user.DeleteStatus != 1 {
		return -1, errors.New("userCode值已删除")
	}
	//endregion

	//region 构造【项目表】信息
	projInfo.ProjName = projName
	projInfo.UserCode = userCode
	projInfo.EndTime = endTime
	projInfo.DeleteStatus = 1
	//endregion

	if isExist {

		//region 修改【项目表】信息
		isSuccess, err := d.ProjDao.Update(&projInfo)
		if err != nil {
			return -1, errors.New("插入【项目表】出错:" + err.Error())
		}
		if !isSuccess {
			return -1, errors.New("插入【项目表】失败")
		}
		return projInfo.ProjId, nil
		//endregion

	} else {

		//region 插入数据库
		projId, err := d.ProjDao.Insert(&projInfo)
		if err != nil {
			return -1, errors.New("插入【项目表】出错:" + err.Error())
		}
		if projId <= 0 {
			return -1, errors.New("插入【项目表】失败")
		}
		return int(projId), nil
		//endregion

	}

}

//修改【项目表】信息
func (d *ProjService) Update(projId int, projName string, userCode string, endTime time.Time) (bool, error) {

	//region 查询【项目表】信息
	projInfo, err := d.ProjDao.Get(projId)
	if err != nil {
		return false, errors.New("查询【项目表】信息出错:" + err.Error())
	}
	if projInfo.ProjId <= 0 {
		return false, errors.New("【项目表】信息不存在")
	}
	if projInfo.DeleteStatus != 1 {
		return false, errors.New("【项目表】信息已被删除")
	}
	//endregion

	//region 验证是否需要执行修改
	if projInfo.ProjId == projId &&
		projInfo.ProjName == projName &&
		projInfo.UserCode == userCode &&
		projInfo.EndTime == endTime {
		return true, nil
	}
	//endregion

	//region 构造【项目表】信息
	projInfo.ProjName = projName
	projInfo.UserCode = userCode
	projInfo.EndTime = endTime
	projInfo.DeleteStatus = 1
	//endregion

	//region 修改【项目表】信息
	return d.ProjDao.Update(&projInfo)
	//endregion

}
