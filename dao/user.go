package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

// 根据【用户编码，取自钉钉编码】查询【用户信息表】表中是否存在相关记录
func (d *UserDao) Exist(userCode string) (bool, error) {
	rows, err := ha666db.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from user where userCode=?", userCode)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	count := 0
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return false, err
		}
		return count > 0, nil
	}
	return false, nil
}

// 插入单条记录到【用户信息表】表中
func (d *UserDao) Insert(m *model.User) (bool, error) {
	result, err := ha666db.Exec("insert into user(userCode,realName,jobNumber,jobPosition,hiredDate,avatar,gender,userType,deleteStatus) values(?,?,?,?,?,?,?,?,?)", m.UserCode, m.RealName, m.JobNumber, m.JobPosition, m.HiredDate, m.Avatar, m.Gender, m.UserType, m.DeleteStatus)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【用户编码，取自钉钉编码】修改【用户信息表】表的单条记录
func (d *UserDao) Update(m *model.User) (bool, error) {
	result, err := ha666db.Exec("update user set realName=?, jobNumber=?, jobPosition=?, hiredDate=?, avatar=?, gender=?, userType=?, deleteStatus=? where userCode=?", m.RealName, m.JobNumber, m.JobPosition, m.HiredDate, m.Avatar, m.Gender, m.UserType, m.DeleteStatus, m.UserCode)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 插入或修改【用户信息表】表的单条记录
func (d *UserDao) InsertUpdate(m *model.User) (bool, error) {
	result, err := ha666db.Exec("insert into user(userCode,realName,jobNumber,jobPosition,hiredDate,avatar,gender,userType,deleteStatus) values(?,?,?,?,?,?,?,?,?) ON DUPLICATE KEY UPDATE realName=?,jobNumber=?,jobPosition=?,hiredDate=?,avatar=?,gender=?,userType=?,deleteStatus=?", m.UserCode, m.RealName, m.JobNumber, m.JobPosition, m.HiredDate, m.Avatar, m.Gender, m.UserType, m.DeleteStatus, m.RealName, m.JobNumber, m.JobPosition, m.HiredDate, m.Avatar, m.Gender, m.UserType, m.DeleteStatus)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【用户编码，取自钉钉编码】软删除【用户信息表】表中的单条记录
func (d *UserDao) Delete(userCode string) (bool, error) {
	result, err := ha666db.Exec("update user set deleteStatus=2 where userCode=?", userCode)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【用户编码，取自钉钉编码】数组软删除【用户信息表】表中的多条记录
func (d *UserDao) DeleteIn(userCodes []string) (count int64, err error) {
	if len(userCodes) <= 0 {
		return count, errors.New("userCodes is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update user set deleteStatus=2")
	sql_str.WriteString(" where userCode in(")
	question_mark := strings.Repeat("?,", len(userCodes))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	vals := make([]interface{}, 0, len(userCodes))
	for _, v := range userCodes {
		vals = append(vals, v)
	}
	result, err := ha666db.Exec(sql_str.String(), vals...)
	if err != nil {
		return count, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return count, err
	}
	return affected, nil
}

// 根据【用户编码，取自钉钉编码】查询【用户信息表】表中的单条记录
func (d *UserDao) Get(userCode string) (user model.User, err error) {
	rows, err := ha666db.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ userCode, realName, jobNumber, jobPosition, hiredDate, avatar, gender, userType, deleteStatus, createTime, updateTime from user where userCode=?", userCode)
	if err != nil {
		return user, err
	}
	defer rows.Close()
	users, err := d._RowsToArray(rows)
	if err != nil {
		return user, err
	}
	if len(users) <= 0 {
		return user, err
	}
	return users[0], nil
}

// 根据【用户编码，取自钉钉编码】数组查询【用户信息表】表中的多条记录
func (d *UserDao) GetIn(userCodes []string) (users []model.User, err error) {
	if len(userCodes) <= 0 {
		return users, errors.New("userCodes is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ userCode, realName, jobNumber, jobPosition, hiredDate, avatar, gender, userType, deleteStatus, createTime, updateTime from ")
	sql_str.WriteString("user")
	sql_str.WriteString(" where userCode in(")
	param_keys := strings.Repeat("?,", len(userCodes))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(userCodes))
	for _, v := range userCodes {
		vals = append(vals, v)
	}
	rows, err = ha666db.Queryx(sql_str.String(), vals...)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 根据【用户编码，取自钉钉编码】查询【用户信息表】表中的单条记录，使用索引【uk_user_userCode,根据用户编码查询唯一用户信息】
func (d *UserDao) GetByUserCode(userCode string) (user model.User, err error) {
	rows, err := ha666db.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ userCode, realName, jobNumber, jobPosition, hiredDate, avatar, gender, userType, deleteStatus, createTime, updateTime from user force index(uk_user_userCode) where userCode=?", userCode)
	if err != nil {
		return user, err
	}
	defer rows.Close()
	users, err := d._RowsToArray(rows)
	if err != nil {
		return user, err
	}
	if len(users) <= 0 {
		return user, err
	}
	return users[0], nil
}

// 根据【姓名】查询【用户信息表】表总记录数，使用索引【idx_user_realName,根据姓名查询多条记录】
func (d *UserDao) GetRowCountByRealName(realName string) (count int, err error) {
	rows, err := ha666db.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from user force index(idx_user_realName) where realName=? and deleteStatus = 1", realName)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return -1, err
		}
		return count, nil
	}
	return -1, nil
}

// 根据【姓名】分页查询【用户信息表】表的记录，使用索引【idx_user_realName,根据姓名查询多条记录】
func (d *UserDao) GetRowListByRealName(realName string, PageIndex, PageSize int) (users []model.User, err error) {
	rows, err := ha666db.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ userCode, realName, jobNumber, jobPosition, hiredDate, avatar, gender, userType, deleteStatus, createTime, updateTime from user force index(idx_user_realName) where realName=? and deleteStatus = 1 limit ?,?", realName, (PageIndex-1)*PageSize, PageSize)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【用户信息表】表总记录数,使用条件：部分【姓名】,【员工的工号】,【职位信息】,【用户类型】,开始【创建时间】,结束【创建时间】
func (d *UserDao) GetRowCount(likeRealName string, jobNumber string, jobPosition string, userType int8, startCreateTime time.Time, endCreateTime time.Time) (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from user")

	//region 处理likeRealName
	if golibs.Length(likeRealName) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("realName like ?")
		params = append(params, "%"+likeRealName+"%")
		conditions++
	}
	//endregion

	//region 处理jobNumber
	if golibs.Length(jobNumber) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("jobNumber=?")
		params = append(params, jobNumber)
		conditions++
	}
	//endregion

	//region 处理jobPosition
	if golibs.Length(jobPosition) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("jobPosition=?")
		params = append(params, jobPosition)
		conditions++
	}
	//endregion

	//region 处理userType
	if userType > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("userType=?")
		params = append(params, userType)
		conditions++
	}
	//endregion

	//region 处理createTime
	isStart := startCreateTime.After(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local))
	isEnd := endCreateTime.After(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local))
	if isStart {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		if isEnd {
			sqlString.Append("createTime between ? and ?")
			params = append(params, startCreateTime.Format(golibs.Time_TIMEStandard))
			params = append(params, endCreateTime.Format(golibs.Time_TIMEStandard))
		} else {
			sqlString.Append("createTime > ?")
			params = append(params, startCreateTime.Format(golibs.Time_TIMEStandard))
		}
		conditions++
	} else {
		if isEnd {
			if conditions > 0 {
				sqlString.Append(" and ")
			} else {
				sqlString.Append(" where ")
			}
			sqlString.Append("createTime < ?")
			params = append(params, endCreateTime.Format(golibs.Time_TIMEStandard))
			conditions++
		}
	}
	//endregion

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region Query
	rows, err := ha666db.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return -1, err
		}
		return count, nil
	}
	//endregion

	return -1, nil
}

// 查询【用户信息表】列表,使用条件：部分【姓名】,【员工的工号】,【职位信息】,【用户类型】,开始【创建时间】,结束【创建时间】
func (d *UserDao) GetRowList(likeRealName string, jobNumber string, jobPosition string, userType int8, startCreateTime time.Time, endCreateTime time.Time, pageIndex, pageSize int) (users []model.User, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ userCode, realName, jobNumber, jobPosition, hiredDate, avatar, gender, userType, deleteStatus, createTime, updateTime from user")

	//region 处理likeRealName
	if golibs.Length(likeRealName) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("realName like ?")
		params = append(params, "%"+likeRealName+"%")
		conditions++
	}
	//endregion

	//region 处理jobNumber
	if golibs.Length(jobNumber) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("jobNumber=?")
		params = append(params, jobNumber)
		conditions++
	}
	//endregion

	//region 处理jobPosition
	if golibs.Length(jobPosition) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("jobPosition=?")
		params = append(params, jobPosition)
		conditions++
	}
	//endregion

	//region 处理userType
	if userType > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("userType=?")
		params = append(params, userType)
		conditions++
	}
	//endregion

	//region 处理createTime
	isStart := startCreateTime.After(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local))
	isEnd := endCreateTime.After(time.Date(2019, 1, 1, 0, 0, 0, 0, time.Local))
	if isStart {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		if isEnd {
			sqlString.Append("createTime between ? and ?")
			params = append(params, startCreateTime.Format(golibs.Time_TIMEStandard))
			params = append(params, endCreateTime.Format(golibs.Time_TIMEStandard))
		} else {
			sqlString.Append("createTime > ?")
			params = append(params, startCreateTime.Format(golibs.Time_TIMEStandard))
		}
		conditions++
	} else {
		if isEnd {
			if conditions > 0 {
				sqlString.Append(" and ")
			} else {
				sqlString.Append(" where ")
			}
			sqlString.Append("createTime < ?")
			params = append(params, endCreateTime.Format(golibs.Time_TIMEStandard))
			conditions++
		}
	}
	//endregion

	//region 处理deleteStatus
	if conditions > 0 {
		sqlString.Append(" and ")
	} else {
		sqlString.Append(" where ")
	}
	sqlString.Append("deleteStatus = 1")
	//endregion

	//region order by
	sqlString.Append(" order by userCode desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := ha666db.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【用户信息表】表记录
func (d *UserDao) _RowsToArray(rows *sqlx.Rows) (users []model.User, err error) {
	for rows.Next() {
		m := model.User{}
		err = rows.Scan(
			&m.UserCode,     //用户编码，取自钉钉编码
			&m.RealName,     //姓名
			&m.JobNumber,    //员工的工号
			&m.JobPosition,  //职位信息
			&m.HiredDate,    //入职时间
			&m.Avatar,       //头像url
			&m.Gender,       //性别，0未知，1男，2女
			&m.UserType,     //用户类型
			&m.DeleteStatus, //删除状态，0未知，1未删除，2删除
			&m.CreateTime,   //创建时间
			&m.UpdateTime)   //修改时间
		if err != nil {
			return users, err
		}
		users = append(users, m)
	}
	return users, err
}
