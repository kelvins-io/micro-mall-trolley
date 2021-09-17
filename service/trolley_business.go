package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-trolley/model/mysql"
	"gitee.com/cristiane/micro-mall-trolley/pkg/code"
	"gitee.com/cristiane/micro-mall-trolley/proto/micro_mall_trolley_proto/trolley_business"
	"gitee.com/cristiane/micro-mall-trolley/repository"
	"gitee.com/kelvins-io/common/json"
	"gitee.com/kelvins-io/kelvins"
	"time"
)

func SkuJoinTrolley(ctx context.Context, req *trolley_business.JoinSkuRequest) (retCode int) {
	retCode = code.Success
	query := map[string]interface{}{
		"uid":      req.Uid,
		"sku_code": req.SkuCode,
		"shop_id":  req.ShopId,
		"state":    1,
	}
	record, err := repository.GetSkuUserTrolley("id,count,state,selected,update_time", query)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetSkuUserTrolley err: %v, query: %v", err, json.MarshalToStringNoError(query))
		retCode = code.ErrorServer
		return
	}
	if record.Id > 0 {
		query := map[string]interface{}{
			"uid":      req.Uid,
			"sku_code": req.SkuCode,
			"shop_id":  req.ShopId,
			"state":    1,
			"count":    record.Count,
		}
		record.Count += int(req.Count)
		record.State = 1
		if req.Selected {
			record.Selected = 2
		} else {
			record.Selected = 1
		}
		record.UpdateTime = time.Now()
		rowsAffected, err := repository.UpdateSkuTrolleyStruct(query, record)
		if err != nil || rowsAffected != 1 {
			kelvins.ErrLogger.Errorf(ctx, "UpdateSkuTrolleyStruct rowsAffected: %v, err: %v, query: %v, record: %v",
				rowsAffected, err, json.MarshalToStringNoError(query), json.MarshalToStringNoError(record))
			retCode = code.ErrorServer
			return
		}
	} else {
		skuAdd := mysql.UserTrolley{
			Uid:        req.Uid,
			ShopId:     req.ShopId,
			SkuCode:    req.SkuCode,
			Count:      int(req.Count),
			JoinTime:   time.Now(),
			State:      1,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		if req.Selected {
			skuAdd.Selected = 2
		} else {
			skuAdd.Selected = 1
		}
		err := repository.AddSkuTrolley(&skuAdd)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "AddSkuTrolley err: %v, skuAdd: %v", err, json.MarshalToStringNoError(skuAdd))
			retCode = code.ErrorServer
			return
		}
	}
	return
}

func RemoveSkuTrolley(ctx context.Context, req *trolley_business.RemoveSkuRequest) (retCode int) {
	if req.Count == 0 {
		req.Count = 1
	}
	originalCount := -1
	{
		query := map[string]interface{}{
			"uid":      req.Uid,
			"sku_code": req.SkuCode,
			"shop_id":  req.ShopId,
			"state":    1,
		}
		record, err := repository.GetSkuUserTrolley("id,count", query)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetSkuUserTrolley err: %v, query: %v", err, json.MarshalToStringNoError(query))
			retCode = code.ErrorServer
			return
		}
		if record != nil && record.Id > 0 {
			originalCount = record.Count
		}
	}
	// 从购物车移除商品
	query := map[string]interface{}{
		"uid":      req.Uid,
		"sku_code": req.SkuCode,
		"shop_id":  req.ShopId,
		"state":    1,
	}
	if originalCount > 0 {
		query["count"] = originalCount
	}
	maps := map[string]interface{}{
		"update_time": time.Now(),
	}
	if req.Count > 0 {
		diffCount := originalCount - int(req.Count)
		if diffCount < 0 {
			diffCount = 0
		}
		maps["count"] = diffCount
		if diffCount == 0 {
			maps["state"] = 2
		}
	} else {
		maps["state"] = 2
		maps["count"] = 0
	}
	rowsAffected, err := repository.UpdateSkuTrolley(query, maps)
	if err != nil || rowsAffected != 1 {
		kelvins.ErrLogger.Errorf(ctx, "UpdateSkuTrolley rowsAffected: %v, err: %v, query: %+v, maps: %+v", rowsAffected, err, query, maps)
		retCode = code.ErrorServer
		return
	}
	retCode = code.Success

	return
}

func GetUserTrolleyList(ctx context.Context, uid int64) ([]mysql.UserTrolley, int) {
	var result = make([]mysql.UserTrolley, 0)
	list, err := repository.GetUserTrolleyList(uid)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetUserTrolleyList err: %v, uid: %v", err, uid)
		return result, code.ErrorServer
	}
	return list, code.Success
}
