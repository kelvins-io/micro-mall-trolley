package repository

import (
	"gitee.com/cristiane/micro-mall-trolley/model/mysql"
	"gitee.com/kelvins-io/kelvins"
)

func AddSkuTrolley(model *mysql.UserTrolley) (err error) {
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableUserTrolley).Insert(model)
	return
}

func UpdateSkuTrolley(query, maps map[string]interface{}) (err error) {
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableUserTrolley).Where(query).Update(maps)
	return
}

func UpdateSkuTrolleyStruct(query map[string]interface{}, model *mysql.UserTrolley) (err error) {
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableUserTrolley).Where(query).Update(model)
	return
}

func GetSkuUserTrolley(query map[string]interface{}) (*mysql.UserTrolley, error) {
	var result mysql.UserTrolley
	var err error
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableUserTrolley).Where(query).Desc("join_time").Get(&result)
	return &result, err
}

func CheckSkuExistUserTrolley(query map[string]interface{}) (exist bool, err error) {
	exist, err = kelvins.XORM_DBEngine.Table(mysql.TableUserTrolley).Where(query).Exist(&mysql.UserTrolley{})
	return
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
