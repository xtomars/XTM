/*
 * *******************************************************************
 * @项目名称: controller
 * @文件名称: asset_verify.go
 * @Date: 2020/02/14
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package controller

import (
	"context"
	"encoding/json"

	"git.bhex.io/bhpc/wallet/chainnode/baasnode/common"
	pbcommon "git.bhex.io/bhpc/wallet/common/protos/common"
	"git.bhex.io/bhpc/wallet/common/protos/wallet"
)

// AssetVerify call by httpserver
func AssetVerify(assetInfo *common.AssetInfo) (*common.AssetVerifyResult, error) {
	log.Infof("AssetVerify asset info:%v", *assetInfo)

	// query wallet db
	reqData := map[string]interface{}{
		"token_id":                assetInfo.TokenID,
		"total_deposit_amount":    assetInfo.TotalDepositAmount,
		"total_withdrawal_amount": assetInfo.TotalWithdrawalAmount,
		"last_block_height":       assetInfo.LastBlockHeight,
	}

	data, err := json.Marshal(reqData)
	if err != nil {
		log.Error("AssetVerify json marshal err:%v", err)
		return nil, common.ErrInvalidParams
	}

	reply := &wallet.QueryWalletDBReply{}
	common.FuncWithRetry(5, func() error {
		reply, err = walletClient.QueryWalletDB(context.Background(), &wallet.QueryWalletDBRequest{
			Method:  "BaasAssetVerify",
			ReqData: data,
		})
		return err
	})
	if err != nil {
		log.Error("AssetVerify query wallet db err:%v", err)
		return nil, common.ErrNeedRetry
	}

	if reply.Code != pbcommon.ReturnCode_SUCCESS {
		log.Error("AssetVerify query wallet db err:%v", reply.Msg)
		return nil, common.ErrInvalidParams
	}

	var result common.AssetVerifyResult
	json.Unmarshal(reply.ResData, &result)

	return &result, nil
}
