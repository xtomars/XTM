/*
 * *******************************************************************
 * @项目名称: controller
 * @文件名称: withdrawal.go
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

// GetWithdrawalOrders return batch unhandled withdrawal orders
func GetWithdrawalOrders(c *gin.Context) {
	chain := c.Query("chain")

	tokenInfo := common.GetTokenInfoByChain(chain)
	if tokenInfo == nil {
		c.JSON(http.StatusOK, Result{
			Code: INVALID_CHAIN,
			Msg:  "invali chain",
		})
		return
	}

	apiKey := c.Request.Header.Get("BWAAS-API-KEY")
	// call grpc get orders
	orders := rpc.GetWithdrawalOrders(chain, apiKey)

	c.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Msg:  "success",
		Data: orders,
	})
}
