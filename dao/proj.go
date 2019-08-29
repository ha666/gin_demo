package dao

import (
	"bytes"
	"errors"
	"github.com/ha666/gin_demo/model"
	"github.com/ha666/golibs"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ProjDao struct {
}

func NewProjDao() *ProjDao {
	return &ProjDao{}
}

// 根据【项目编号】查询【项目表】表中是否存在相关记录
func (d *ProjDao) Exist(projId int) (bool, error) {
	rows, err := ha666db.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from proj where projId=?", projId)
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

// 插入单条记录到【项目表】表中
func (d *ProjDao) Insert(m *model.Proj) (int64, error) {
	result, err := ha666db.Exec("insert into proj(projName,userCode,deleteStatus,endTime) values(?,?,?,?)", m.ProjName, m.UserCode, m.DeleteStatus, m.EndTime)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// 根据【项目编号】修改【项目表】表的单条记录
func (d *ProjDao) Update(m *model.Proj) (bool, error) {
	result, err := ha666db.Exec("update proj set projName=?, userCode=?, deleteStatus=?, endTime=? where projId=?", m.ProjName, m.UserCode, m.DeleteStatus, m.EndTime, m.ProjId)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【项目编号】软删除【项目表】表中的单条记录
func (d *ProjDao) Delete(projId int) (bool, error) {
	result, err := ha666db.Exec("update proj set deleteStatus=2 where projId=?", projId)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return affected > 0, nil
}

// 根据【项目编号】数组软删除【项目表】表中的多条记录
func (d *ProjDao) DeleteIn(projIds []int) (count int64, err error) {
	if len(projIds) <= 0 {
		return count, errors.New("projIds is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("update proj set deleteStatus=2")
	sql_str.WriteString(" where projId in(")
	question_mark := strings.Repeat("?,", len(projIds))
	sql_str.WriteString(question_mark[:len(question_mark)-1])
	sql_str.WriteString(")")
	vals := make([]interface{}, 0, len(projIds))
	for _, v := range projIds {
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

// 根据【项目编号】查询【项目表】表中的单条记录
func (d *ProjDao) Get(projId int) (proj model.Proj, err error) {
	rows, err := ha666db.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ projId, projName, userCode, deleteStatus, endTime from proj where projId=?", projId)
	if err != nil {
		return proj, err
	}
	defer rows.Close()
	projs, err := d._RowsToArray(rows)
	if err != nil {
		return proj, err
	}
	if len(projs) <= 0 {
		return proj, err
	}
	return projs[0], nil
}

// 根据【项目编号】数组查询【项目表】表中的多条记录
func (d *ProjDao) GetIn(projIds []int) (projs []model.Proj, err error) {
	if len(projIds) <= 0 {
		return projs, errors.New("projIds is empty")
	}
	sql_str := bytes.Buffer{}
	sql_str.WriteString("select /*+ MAX_EXECUTION_TIME(5000) */ projId, projName, userCode, deleteStatus, endTime from ")
	sql_str.WriteString("proj")
	sql_str.WriteString(" where projId in(")
	param_keys := strings.Repeat("?,", len(projIds))
	sql_str.WriteString(param_keys[:len(param_keys)-1])
	sql_str.WriteString(")")
	var rows *sqlx.Rows
	vals := make([]interface{}, 0, len(projIds))
	for _, v := range projIds {
		vals = append(vals, v)
	}
	rows, err = ha666db.Queryx(sql_str.String(), vals...)
	if err != nil {
		return projs, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 根据【项目名称】查询【项目表】表总记录数，使用索引【idx_proj_projName,根据项目名称查询列表】
func (d *ProjDao) GetRowCountByProjName(projName string) (count int, err error) {
	rows, err := ha666db.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from proj force index(idx_proj_projName) where projName=? and deleteStatus = 1", projName)
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

// 根据【项目名称】分页查询【项目表】表的记录，使用索引【idx_proj_projName,根据项目名称查询列表】
func (d *ProjDao) GetRowListByProjName(projName string, PageIndex, PageSize int) (projs []model.Proj, err error) {
	rows, err := ha666db.Queryx("select /*+ MAX_EXECUTION_TIME(5000) */ projId, projName, userCode, deleteStatus, endTime from proj force index(idx_proj_projName) where projName=? and deleteStatus = 1 limit ?,?", projName, (PageIndex-1)*PageSize, PageSize)
	if err != nil {
		return projs, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
}

// 查询【项目表】表总记录数,使用条件：部分【项目名称】,【用户编码】
func (d *ProjDao) GetRowCount(likeProjName string, userCode string) (count int, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ count(0) Count from proj")

	//region 处理likeProjName
	if golibs.Length(likeProjName) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("projName like ?")
		params = append(params, "%"+likeProjName+"%")
		conditions++
	}
	//endregion

	//region 处理userCode
	if golibs.Length(userCode) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("userCode=?")
		params = append(params, userCode)
		conditions++
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

// 根据【用户编码】查询【项目表】总记录数
func (d *ProjDao) GetRowCountByUserCode(userCode string) (count int, err error) {
	rows, err := ha666db.Queryx("select count(0) Count from proj where userCode=? and deleteStatus = 1", userCode)
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
	return -1, err
}

// 查询【项目表】列表,使用条件：部分【项目名称】,【用户编码】
func (d *ProjDao) GetRowList(likeProjName string, userCode string, pageIndex, pageSize int) (projs []model.Proj, err error) {
	sqlString := golibs.NewStringBuilder()
	params := make([]interface{}, 0)
	conditions := 0

	sqlString.Append("select /*+ MAX_EXECUTION_TIME(5000) */ projId, projName, userCode, deleteStatus, endTime from proj")

	//region 处理likeProjName
	if golibs.Length(likeProjName) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("projName like ?")
		params = append(params, "%"+likeProjName+"%")
		conditions++
	}
	//endregion

	//region 处理userCode
	if golibs.Length(userCode) > 0 {
		if conditions > 0 {
			sqlString.Append(" and ")
		} else {
			sqlString.Append(" where ")
		}
		sqlString.Append("userCode=?")
		params = append(params, userCode)
		conditions++
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
	sqlString.Append(" order by projId desc")
	//endregion

	//region limit
	sqlString.Append(" limit ?,?")
	params = append(params, (pageIndex-1)*pageSize)
	params = append(params, pageSize)
	//endregion

	//region Query
	rows, err := ha666db.Queryx(sqlString.ToString(), params...)
	if err != nil {
		return projs, err
	}
	defer rows.Close()
	return d._RowsToArray(rows)
	//endregion
}

// 解析【项目表】表记录
func (d *ProjDao) _RowsToArray(rows *sqlx.Rows) (projs []model.Proj, err error) {
	for rows.Next() {
		m := model.Proj{}
		err = rows.Scan(
			&m.ProjId,       //项目编号
			&m.ProjName,     //项目名称
			&m.UserCode,     //用户编码
			&m.DeleteStatus, //是否删除
			&m.EndTime)      //结束时间
		if err != nil {
			return projs, err
		}
		projs = append(projs, m)
	}
	return projs, err
}
