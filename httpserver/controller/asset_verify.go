/*
 * *******************************************************************
 * @项目名称: controller
 * @文件名称: asset_verify.go
 * @Date: 2020/02/07
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package controller

import (
	"net/http"

	"git.bhex.io/bhpc/wallet/chainnode/baasnode/common"
	rpc "git.bhex.io/bhpc/wallet/chainnode/baasnode/grpcserver/controller"

	"github.com/gin-gonic/gin"
)

// AssetVerify verify the request asset
func AssetVerify(c *gin.Context) {
	requestID := getRequestID(c)

	var para common.AssetInfo
	if err := c.ShouldBind(&para); err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, Result{
			Code: INVALID_PARAMS,
			Msg:  "invalid params",
		})
		return
	}
	log.Infof("AssetVerify request_id:%v, para: %v", requestID, para)

	if para.TokenID == "" {
		log.Errorf("AssetVerify request_id:%v, invalid token_id:%v", requestID, para.TokenID)

		c.JSON(http.StatusOK, Result{
			Code: INVALID_TOKEN_ID,
			Msg:  "invalid token_id",
		})
		return
	}

	if para.LastBlockHeight == "" || para.TotalWithdrawalAmount == "" || para.TotalDepositAmount == "" {
		c.JSON(http.StatusOK, Result{
			Code: INVALID_PARAMS,
			Msg:  "invalid params",
		})
		return
	}

	// call grpc verify asset
	result, err := rpc.AssetVerify(&para)
	if err != nil {
		log.Errorf("AssetVerify request_id:%v, rpc notify err:%v", requestID, err)

		code, msg := convertErr(err)
		c.JSON(http.StatusOK, Result{
			Code: code,
			Msg:  msg,
		})
		return
	}
	log.Infof("AssetVerify request_id:%v, asset verify result: %v", requestID, result)

	if !result.Success {
		c.JSON(http.StatusOK, Result{
			Code: ASSET_VERIFY_FAILED,
			Msg:  "asset verify failed",
			Data: result,
		})
		return
	}

	c.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Msg:  "success",
	})
}
