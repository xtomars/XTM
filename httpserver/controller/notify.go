/*
 * *******************************************************************
 * @项目名称: controller
 * @文件名称: notify.go
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

// NotifyDeposit receive deposit notify
func NotifyDeposit(c *gin.Context) {
	requestID := getRequestID(c)

	var para common.NotifyDepositParams
	if err := c.ShouldBind(&para); err != nil {
		log.Errorf("NotifyDeposit request_id:%v, parse params err:%v", requestID, err)
		c.JSON(http.StatusOK, Result{
			Code: INVALID_PARAMS,
			Msg:  "invalid params",
		})
		return
	}
	log.Infof("NotifyDeposit request_id:%v, para: %v", requestID, para)

	apiKey := c.Request.Header.Get("BWAAS-API-KEY")
	signature := c.Request.Header.Get("BWAAS-API-SIGNATURE")
	timestamp := c.Request.Header.Get("BWAAS-API-TIMESTAMP")

	// call grpc notify deposit
	err := rpc.NotifyDeposit(para, apiKey, signature, timestamp)
	if err != nil {
		log.Errorf("NotifyDeposit request_id:%v, rpc notify err:%v", requestID, err)

		code, msg := convertErr(err)
		c.JSON(http.StatusOK, Result{
			Code: code,
			Msg:  msg,
		})
		return
	}

	log.Infof("NotifyDeposit request_id:%v,success", requestID)

	c.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Msg:  "success",
	})
}

// NotifyFailed receive withdrawal failed notify
func NotifyFailed(c *gin.Context) {
	requestID := getRequestID(c)

	var para common.NotifyFailedParams
	if err := c.ShouldBind(&para); err != nil {
		log.Errorf("NotifyFailed request_id:%v, parse params err:%v", requestID, err)
		c.JSON(http.StatusOK, Result{
			Code: INVALID_PARAMS,
			Msg:  "invalid params",
		})
		return
	}
	log.Infof("NotifyFailed request_id:%v, para: %v", requestID, para)

	// get redis global lock
	lock, success := common.TryLock(requestID, 30)
	if !success {
		log.Errorf("NotifyFailed need to wait")

		c.JSON(http.StatusOK, Result{
			Code: NEED_WAIT,
			Msg:  "need to wait",
		})
		return
	}
	defer lock.Unlock()

	// call grpc notify withdrawal
	err := rpc.NotifyFailed(para)
	if err != nil {
		log.Errorf("NotifyFailed request_id:%v,rpc notify err:%v", requestID, err)

		code, msg := convertErr(err)
		c.JSON(http.StatusOK, Result{
			Code: code,
			Msg:  msg,
		})

		return
	}

	log.Infof("NotifyFailed request_id:%v success", requestID)

	c.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Msg:  "success",
	})
}

// NotifyWithdrawal receive withdrawal notify
func NotifyWithdrawal(c *gin.Context) {
	requestID := getRequestID(c)

	var para common.NotifyWithdrawalParams
	if err := c.ShouldBind(&para); err != nil {
		log.Errorf("NotifyWithdrawal request_id:%v, parse params err:%v", requestID, err)
		c.JSON(http.StatusOK, Result{
			Code: INVALID_PARAMS,
			Msg:  "invalid params",
		})
		return
	}
	log.Infof("NotifyWithdrawal request_id:%v, para: %v", requestID, para)

	// call grpc notify withdrawal
	err := rpc.NotifyWithdrawal(para)
	if err != nil {
		log.Errorf("NotifyWithdrawal request_id:%v, rpc notify err:%v", requestID, err)

		code, msg := convertErr(err)
		c.JSON(http.StatusOK, Result{
			Code: code,
			Msg:  msg,
		})
		return
	}

	log.Infof("NotifyWithdrawal request_id:%v success", requestID)

	c.JSON(http.StatusOK, Result{
		Code: SUCCESS,
		Msg:  "success",
	})
}
