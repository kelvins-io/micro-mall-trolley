package mysql

import (
	"time"
)

const (
	TableUserTrolley = "user_trolley"
)

type UserTrolley struct {
	Id         int64     `xorm:"pk autoincr comment('自增ID') BIGINT"`
	Uid        int64     `xorm:"not null comment('用户ID') unique(shop_id_sku_uid_index) BIGINT"`
	ShopId     int64     `xorm:"not null comment('店铺ID') index(shop_id_sku_index) unique(shop_id_sku_uid_index) BIGINT"`
	SkuCode    string    `xorm:"not null comment('商品sku') index(shop_id_sku_index) unique(shop_id_sku_uid_index) index CHAR(40)"`
	Count      int       `xorm:"not null default 1 comment('商品数量') INT"`
	JoinTime   time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('加入时间') DATETIME"`
	Selected   int       `xorm:"default 0 comment('是否选中，0-未选中，1-选中') TINYINT(1)"`
	State      int       `xorm:"default 0 comment('状态，0-有效，1-移除') TINYINT"`
	CreateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('更新时间') DATETIME"`
}
