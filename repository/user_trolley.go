package repository

import (
	"gitee.com/cristiane/micro-mall-trolley/model/mysql"
	"gitee.com/kelvins-io/kelvins"
)

func AddSkuTrolley(model *mysql.UserTrolley) (err error) {
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableUserTrolley).Insert(model)
	return
}

func UpdateSkuTrolley(query, maps map[string]interface{}) (rowsAffected int64, err error) {
	rowsAffected, err = kelvins.XORM_DBEngine.Table(mysql.TableUserTrolley).Where(query).Update(maps)
	return
}

func UpdateSkuTrolleyStruct(query map[string]interface{}, model *mysql.UserTrolley) (rowsAffected int64, err error) {
	rowsAffected, err = kelvins.XORM_DBEngine.Table(mysql.TableUserTrolley).Where(query).Update(model)
	return
}

func GetSkuUserTrolley(selectSql string, query map[string]interface{}) (*mysql.UserTrolley, error) {
	var result mysql.UserTrolley
	var err error
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableUserTrolley).Select(selectSql).Where(query).Desc("join_time").Get(&result)
	return &result, err
}

func GetUserTrolleyList(uid int64) ([]mysql.UserTrolley, error) {
	var list = make([]mysql.UserTrolley, 0)
	var err error
	err = kelvins.XORM_DBEngine.Table(mysql.TableUserTrolley).
		Where("uid = ?", uid).
		Where("state = 1").
		Desc("join_time").Find(&list)
	return list, err
}
