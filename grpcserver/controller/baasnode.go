/*
 * *******************************************************************
 * @项目名称: controller
 * @文件名称: baasnode.go
 * @Date: 2020/02/12
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package controller

import (
	"context"
	"fmt"

	baascommon "git.bhex.io/bhpc/wallet/chainnode/baasnode/common"
	"git.bhex.io/bhpc/wallet/common/logger"
	"git.bhex.io/bhpc/wallet/common/protos/chainnode"
	"git.bhex.io/bhpc/wallet/common/protos/common"
)

var log = logger.New("controller")

// chainnode server implement
type baasnode struct{}

// NewBaasnode return the grpc server implement
func NewBaasnode() chainnode.ChainnodeServer {
	return &baasnode{}
}

// Address updates
func (b *baasnode) Address(ctx context.Context, req *chainnode.AddressRequest) (*chainnode.AddressReply, error) {
	var err error
	var newAddress string

	switch req.Action {
	case chainnode.Action_ADD, chainnode.Action_REMOVE:
		return &chainnode.AddressReply{
			Code: common.ReturnCode_SUCCESS,
			Msg:  "not support",
		}, nil
	case chainnode.Action_CREATE:
		newAddress, err = createAddress(req.TokenId)
		if err != nil {
			return &chainnode.AddressReply{
				Code: common.ReturnCode_CHAIN_ERROR,
				Msg:  err.Error(),
			}, nil
		}
	}

	return &chainnode.AddressReply{
		Code:    common.ReturnCode_SUCCESS,
		Address: newAddress,
	}, nil
}

// Send transaction to network
func (b *baasnode) SendTransaction(ctx context.Context, req *chainnode.SendTransactionRequest) (*chainnode.SendTransactionReply, error) {
	return &chainnode.SendTransactionReply{
		Code: common.ReturnCode_SYSTEM_ERROR,
		Msg:  "does not support",
	}, nil
}

// Send transaction to network
func (b *baasnode) SendOnlineWalletTransaction(ctx context.Context, req *common.SendOnlineWalletTransactionRequest) (*common.SendOnlineWalletTransactionReply, error) {
	return &common.SendOnlineWalletTransactionReply{
		Code: common.ReturnCode_SYSTEM_ERROR,
		Msg:  "does not support",
	}, nil
}

// Query some info
func (b *baasnode) Query(ctx context.Context, req *chainnode.QueryRequest) (*chainnode.QueryReply, error) {
	var result uint64

	errFunc := func(msg string) (*chainnode.QueryReply, error) {
		return &chainnode.QueryReply{
			Code: common.ReturnCode_CHAIN_ERROR,
			Msg:  msg,
		}, nil
	}

	switch req.Cmd {
	case chainnode.QueryCmd_BALANCE:

	case chainnode.QueryCmd_BALANCE_OF:

	case chainnode.QueryCmd_ADDRESS_EXIST:
		// para[0] token_id
		// para[1] address
		// para[2] memo

		if len(req.Paras) < 2 {
			return errFunc("erro para")
		}
		log.Infof("QueryCmd_ADDRESS_EXIST Paras:%v ", req.Paras)

		tokenInfo := baascommon.GetTokenInfoByTokenID(req.Paras[0])
		if tokenInfo == nil {
			return errFunc("invalid token_id")
		}

		exist, err := baascommon.Redis.IsSetMember(fmt.Sprintf(baascommon.UsedAddressSet, tokenInfo.Chain), req.Paras[1])
		if err != nil {
			return errFunc("not found address")
		}

		// default valid
		result = 2

		if exist {
			result = 1
		}

	default:
		return errFunc("don't support")
	}

	return &chainnode.QueryReply{
		Code:   common.ReturnCode_SUCCESS,
		Result: result,
	}, nil
}
