package server

import (
	"context"
	"gitee.com/cristiane/micro-mall-trolley/pkg/code"
	"gitee.com/cristiane/micro-mall-trolley/pkg/util"
	"gitee.com/cristiane/micro-mall-trolley/proto/micro_mall_trolley_proto/trolley_business"
	"gitee.com/cristiane/micro-mall-trolley/service"
	"gitee.com/kelvins-io/common/errcode"
)

type TrolleyBusinessServer struct{}

func NewTrolleyBusinessServer() trolley_business.TrolleyBusinessServiceServer {
	return new(TrolleyBusinessServer)
}

func (t *TrolleyBusinessServer) JoinSku(ctx context.Context, req *trolley_business.JoinSkuRequest) (*trolley_business.JoinSkuResponse, error) {
	var result trolley_business.JoinSkuResponse
	result.Common = &trolley_business.CommonResponse{
		Code: trolley_business.RetCode_SUCCESS,
		Msg:  errcode.GetErrMsg(code.Success),
	}
	if req.Uid <= 0 {
		result.Common.Code = trolley_business.RetCode_USER_NOT_EXIST
		result.Common.Msg = errcode.GetErrMsg(code.UserNotExist)
		return &result, nil
	}
	if req.ShopId <= 0 {
		result.Common.Code = trolley_business.RetCode_SHOP_NOT_EXIST
		result.Common.Msg = errcode.GetErrMsg(code.ShopBusinessNotExist)
		return &result, nil
	}
	retCode := service.SkuJoinTrolley(ctx, req)
	if retCode != code.Success {
		if retCode == code.UserNotExist {
			result.Common.Code = trolley_business.RetCode_USER_NOT_EXIST
			result.Common.Msg = errcode.GetErrMsg(code.UserNotExist)
			return &result, nil
		} else if retCode == code.ShopBusinessNotExist {
			result.Common.Code = trolley_business.RetCode_SHOP_NOT_EXIST
			result.Common.Msg = errcode.GetErrMsg(code.ShopBusinessNotExist)
			return &result, nil
		} else {
			result.Common.Code = trolley_business.RetCode_ERROR
			result.Common.Msg = errcode.GetErrMsg(code.ErrorServer)
			return &result, nil
		}
	}

	return &result, nil
}

func (t *TrolleyBusinessServer) RemoveSku(ctx context.Context, req *trolley_business.RemoveSkuRequest) (*trolley_business.RemoveSkuResponse, error) {
	var result trolley_business.RemoveSkuResponse
	result.Common = &trolley_business.CommonResponse{
		Code: trolley_business.RetCode_SUCCESS,
		Msg:  errcode.GetErrMsg(code.Success),
	}
	if req.Uid <= 0 {
		result.Common.Code = trolley_business.RetCode_USER_NOT_EXIST
		result.Common.Msg = errcode.GetErrMsg(code.UserNotExist)
		return &result, nil
	}
	if req.ShopId <= 0 {
		result.Common.Code = trolley_business.RetCode_SHOP_NOT_EXIST
		result.Common.Msg = errcode.GetErrMsg(code.ShopBusinessNotExist)
		return &result, nil
	}
	retCode := service.RemoveSkuTrolley(ctx, req)
	if retCode != code.Success {
		if retCode == code.UserNotExist {
			result.Common.Code = trolley_business.RetCode_USER_NOT_EXIST
			result.Common.Msg = errcode.GetErrMsg(code.UserNotExist)
			return &result, nil
		} else if retCode == code.ShopBusinessNotExist {
			result.Common.Code = trolley_business.RetCode_SHOP_NOT_EXIST
			result.Common.Msg = errcode.GetErrMsg(code.ShopBusinessNotExist)
			return &result, nil
		} else {
			result.Common.Code = trolley_business.RetCode_ERROR
			result.Common.Msg = errcode.GetErrMsg(code.ErrorServer)
			return &result, nil
		}
	}

	return &result, nil
}

func (t *TrolleyBusinessServer) GetUserTrolleyList(ctx context.Context, req *trolley_business.GetUserTrolleyListRequest) (*trolley_business.GetUserTrolleyListResponse, error) {
	var result trolley_business.GetUserTrolleyListResponse
	result.Common = &trolley_business.CommonResponse{}
	result.Records = make([]*trolley_business.UserTrolleyRecord, 0)

	list, retCode := service.GetUserTrolleyList(ctx, req.Uid)
	if retCode != code.Success {
		if retCode == code.UserNotExist {
			result.Common.Code = trolley_business.RetCode_USER_NOT_EXIST
			result.Common.Msg = errcode.GetErrMsg(code.UserNotExist)
			return &result, nil
		}
	}
	result.Records = make([]*trolley_business.UserTrolleyRecord, len(list))
	for i := 0; i < len(list); i++ {
		record := trolley_business.UserTrolleyRecord{
			SkuCode:  list[i].SkuCode,
			ShopId:   list[i].ShopId,
			Time:     util.ParseTimeOfStr(list[i].JoinTime.Unix()),
			Count:    int64(list[i].Count),
			Selected: false,
		}
		if list[i].Selected == 1 {
			record.Selected = false
		} else {
			record.Selected = true
		}
		result.Records[i] = &record
	}

	return &result, nil
}
