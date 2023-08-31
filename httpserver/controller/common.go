/*
 * *******************************************************************
 * @项目名称: controller
 * @文件名称: common.go
 * @Date: 2020/02/07
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package controller

import (
	"git.bhex.io/bhpc/wallet/chainnode/baasnode/common"
	"git.bhex.io/bhpc/wallet/common/logger"
	"github.com/gin-gonic/gin"
)

var log = logger.New("controller")

// Code self code
type Code int

const (
	_ Code = 9999 + iota
	SUCCESS
	INVALID_SIGN
	INVALID_APIKEY
	INVALID_CHAIN
	INVALID_TOKEN_ID
	INVALID_PARAMS
	INVALID_TO_ADDRESS
	INVALID_ORDER_ID
	INVALID_AMOUNT
	INVALID_DECIMALS
	INVALID_BLOCK_HEIGHT
	INVALID_BLOCK_TIME
	INVALID_TXHASH
	INVALID_INDEX
	NETWORK_ERROR
	REPEAT_DEPOSIT
	ASSET_VERIFY_FAILED
	DEPOSIT_SUSPENDED
	WITHDRAWAL_SUSPENDED
	TIMESTAMP_EXPIRED
	MEMO_REQUIRED
	NEED_WAIT
	INVALID_FROM_ADDRESS
	ADDRESS_ENOUGH
	NEED_RETRY
	INVALID_MEMO
)

// Result common resp struct
type Result struct {
	Code Code        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func convertErr(err error) (Code, string) {
	var code Code
	var msg string

	switch err {
	case common.ErrInvalidChain:
		code = INVALID_CHAIN
		msg = "invalid chain"
	case common.ErrInvalidAPIKey:
		code = INVALID_APIKEY
		msg = "invalid apikey"
	case common.ErrInvalidToAddress:
		code = INVALID_TO_ADDRESS
		msg = "invalid to address"
	case common.ErrInvalidOrderID:
		code = INVALID_ORDER_ID
		msg = "invalid order id"
	case common.ErrInvalidTokenID:
		code = INVALID_TOKEN_ID
		msg = "invalid token id"
	case common.ErrVerify:
		code = ASSET_VERIFY_FAILED
		msg = "asset verify failed"
	case common.ErrInvalidAmount:
		code = INVALID_AMOUNT
		msg = "invalid amount"
	case common.ErrInvalidDecimals:
		code = INVALID_DECIMALS
		msg = "invalid decimals"
	case common.ErrInvalidBlockHeight:
		code = INVALID_BLOCK_HEIGHT
		msg = "invalid block height"
	case common.ErrinvalidBlockTime:
		code = INVALID_BLOCK_TIME
		msg = "invalid block time"
	case common.ErrInvalidIndex:
		code = INVALID_INDEX
		msg = "invalid index"
	case common.ErrInvalidTxHash:
		code = INVALID_TXHASH
		msg = "invalid tx_hash"
	// case common.ErrNetwork:
	// 	code = NETWORK_ERROR
	// 	msg = "network error"
	case common.ErrInvalidParams:
		code = INVALID_PARAMS
		msg = "invalid params"
	case common.ErrPauseDeposit:
		code = DEPOSIT_SUSPENDED
		msg = "deposit suspended"
	case common.ErrPauseWithdrawal:
		code = WITHDRAWAL_SUSPENDED
		msg = "withdrawal suspended"
	case common.ErrNeedMemo:
		code = MEMO_REQUIRED
		msg = "memo required"
	case common.ErrInvalidFromAddress:
		code = INVALID_FROM_ADDRESS
		msg = "invalid from address"
	case common.ErrNeedRetry:
		code = NEED_RETRY
		msg = "need retry"
	case common.ErrRepeatDeposit:
		code = REPEAT_DEPOSIT
		msg = "repeat deposit"
	case common.ErrInvalidMemo:
		code = INVALID_MEMO
		msg = "invalid memo"
	default:
		code = SUCCESS
		msg = "success"
	}
	return code, msg
}

func getRequestID(c *gin.Context) string {
	requestID := c.Request.Header.Get("X-Request-Id")
	if requestID == "" {
		requestID = "unknown"
	}
	return requestID
}
