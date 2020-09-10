package repository

import (
	"gitee.com/cristiane/micro-mall-trolley/model/mysql"
	"gitee.com/kelvins-io/kelvins"
)

func AddSkuTrolley(model *mysql.UserTrolley) (err error) {
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableUserTrolley).Insert(model)
	return
}

func RemoveSkuTrolley(query, maps map[string]interface{}) (err error) {
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableUserTrolley).Where(query).Update(maps)
	return
}
