package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-trolley/model/args"
	"gitee.com/cristiane/micro-mall-trolley/model/mysql"
	"gitee.com/cristiane/micro-mall-trolley/pkg/code"
	"gitee.com/cristiane/micro-mall-trolley/pkg/util"
	"gitee.com/cristiane/micro-mall-trolley/proto/micro_mall_shop_proto/shop_business"
	"gitee.com/cristiane/micro-mall-trolley/proto/micro_mall_trolley_proto/trolley_business"
	"gitee.com/cristiane/micro-mall-trolley/proto/micro_mall_users_proto/users"
	"gitee.com/cristiane/micro-mall-trolley/repository"
	"gitee.com/kelvins-io/kelvins"
	"time"
)

func SkuJoinTrolley(ctx context.Context, req *trolley_business.JoinSkuRequest) (retCode int) {
	// 检查店铺是否存在
	if req.ShopId > 0 {
		serverName := args.RpcServiceMicroMallShop
		conn, err := util.GetGrpcClient(serverName)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
			retCode = code.ErrorServer
			return
		}
		defer conn.Close()

		client := shop_business.NewShopBusinessServiceClient(conn)
		r := shop_business.GetShopMaterialRequest{
			ShopId: req.ShopId,
		}
		rsp, err := client.GetShopMaterial(ctx, &r)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetShopMaterial %v,err: %v, req: %+v", serverName, err, r)
			retCode = code.ErrorServer
			return
		}
		if rsp == nil || rsp.Material == nil || rsp.Material.ShopId <= 0 {
			retCode = code.ShopBusinessNotExist
			return
		}
	}
	// 检查用户是否存在
	if req.Uid > 0 {
		serverName := args.RpcServiceMicroMallUsers
		conn, err := util.GetGrpcClient(serverName)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
			retCode = code.ErrorServer
			return
		}
		defer conn.Close()
		client := users.NewUsersServiceClient(conn)
		r := users.GetUserInfoRequest{
			Uid: req.Uid,
		}
		rsp, err := client.GetUserInfo(ctx, &r)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetUserInfo %v,err: %v, req: %+v", serverName, err, r)
			retCode = code.ErrorServer
			return
		}
		if rsp.Info.Uid <= 0 {
			retCode = code.UserNotExist
			return
		}
	}
	query := map[string]interface{}{
		"uid":      req.Uid,
		"sku_code": req.SkuCode,
		"shop_id":  req.ShopId,
		"state":    1,
	}

	record, err := repository.GetSkuUserTrolley(query)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "CheckSkuExistUserTrolley %v,err: %v, query: %+v", err, query)
		retCode = code.ErrorServer
		return
	}
	if record.Uid > 0 {
		query := map[string]interface{}{
			"uid":      req.Uid,
			"sku_code": req.SkuCode,
			"shop_id":  req.ShopId,
			"state":    1,
		}
		record.Count += int(req.Count)
		record.State = 1
		if req.Selected {
			record.Selected = 2
		} else {
			record.Selected = 1
		}
		record.JoinTime = time.Now()
		record.UpdateTime = time.Now()
		err := repository.UpdateSkuTrolleyStruct(query, record)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "UpdateSkuTrolleyStruct %v,err: %v, query: %+v, record: %+v", err, query, record)
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
			kelvins.ErrLogger.Errorf(ctx, "AddSkuTrolley %v,err: %v, skuAdd: %+v", err, skuAdd)
			retCode = code.ErrorServer
			return
		}
	}

	retCode = code.Success

	return
}

func RemoveSkuTrolley(ctx context.Context, req *trolley_business.RemoveSkuRequest) (retCode int) {
	// 检查店铺是否存在
	//if req.ShopId > 0 {
	//	serverName := args.RpcServiceMicroMallShop
	//	conn, err := util.GetGrpcClient(serverName)
	//	if err != nil {
	//		kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
	//		retCode = code.ErrorServer
	//		return
	//	}
	//	defer conn.Close()
	//
	//	client := shop_business.NewShopBusinessServiceClient(conn)
	//	r := shop_business.GetShopMajorInfoRequest{
	//		ShopIds: []int64{req.ShopId},
	//	}
	//	rsp, err := client.GetShopMajorInfo(ctx, &r)
	//	if err != nil {
	//		kelvins.ErrLogger.Errorf(ctx, "GetShopMaterial %v,err: %v, req: %+v", serverName, err, r)
	//		retCode = code.ErrorServer
	//		return
	//	}
	//	if rsp.Common.Code != shop_business.RetCode_SUCCESS {
	//		retCode = code.ShopBusinessNotExist
	//		return
	//	}
	//}
	//// 检查用户是否存在
	//if req.Uid > 0 {
	//	serverName := args.RpcServiceMicroMallUsers
	//	conn, err := util.GetGrpcClient(serverName)
	//	if err != nil {
	//		kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
	//		retCode = code.ErrorServer
	//		return
	//	}
	//	defer conn.Close()
	//	client := users.NewUsersServiceClient(conn)
	//	r := users.GetUserAccountIdRequest{
	//		UidList: []int64{req.Uid},
	//	}
	//	rsp, err := client.GetUserAccountId(ctx, &r)
	//	if err != nil {
	//		kelvins.ErrLogger.Errorf(ctx, "GetUserInfo %v,err: %v, req: %+v", serverName, err, r)
	//		retCode = code.ErrorServer
	//		return
	//	}
	//	if rsp.Common.Code != users.RetCode_SUCCESS {
	//		retCode = code.UserNotExist
	//		return
	//	}
	//}
	// 从购物车移除商品
	query := map[string]interface{}{
		"uid":      req.Uid,
		"sku_code": req.SkuCode,
		"shop_id":  req.ShopId,
	}
	maps := map[string]interface{}{
		"state": 2,
	}
	err := repository.UpdateSkuTrolley(query, maps)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "UpdateSkuTrolley %v,err: %v, query: %+v, maps: %+v", err, query, maps)
		retCode = code.ErrorServer
		return
	}
	retCode = code.Success

	return
}

func GetUserTrolleyList(ctx context.Context, uid int64) ([]mysql.UserTrolley, int) {
	var result = make([]mysql.UserTrolley, 0)
	// 检查用户是否存在
	if uid > 0 {
		serverName := args.RpcServiceMicroMallUsers
		conn, err := util.GetGrpcClient(serverName)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
			return result, code.ErrorServer
		}
		defer conn.Close()
		client := users.NewUsersServiceClient(conn)
		r := users.GetUserInfoRequest{
			Uid: uid,
		}
		rsp, err := client.GetUserInfo(ctx, &r)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetUserInfo %v,err: %v, req: %+v", serverName, err, r)
			return result, code.ErrorServer
		}
		if rsp.Info.Uid <= 0 {
			return result, code.UserNotExist
		}
	}

	list, err := repository.GetUserTrolleyList(uid)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetUserTrolleyList %v,err: %v, uid: %+v", err, uid)
		return result, code.ErrorServer
	}
	return list, code.Success
}
