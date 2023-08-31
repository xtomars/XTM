/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: error.go
 * @Date: 2020/02/12
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package common

import "errors"

var (
	ErrInvalidChain       = errors.New("invalid chain")
	ErrInvalidAPIKey      = errors.New("invalid apikey")
	ErrInvalidFromAddress = errors.New("invalid from address")
	ErrInvalidMemo        = errors.New("invalid memo")
	ErrInvalidToAddress   = errors.New("invalid to address")
	ErrInvalidOrderID     = errors.New("invalid order_id")
	ErrInvalidTokenID     = errors.New("invalid token_id")
	ErrVerify             = errors.New("verify falied")
	ErrNeedMemo           = errors.New("memo required")
	ErrInvalidAmount      = errors.New("invalid amount")
	ErrInvalidFee         = errors.New("invalid fee")
	ErrInvalidBlockHeight = errors.New("invalid block height")
	ErrinvalidBlockTime   = errors.New("invalid block time")
	ErrInvalidDecimals    = errors.New("invalid decimals")
	ErrInvalidIndex       = errors.New("invalid index")
	ErrInvalidTxHash      = errors.New("invalid tx_hash")
	ErrInvalidParams      = errors.New("invalid params")
	ErrPauseDeposit       = errors.New("deposit suspended")
	ErrPauseWithdrawal    = errors.New("withdrawal suspended")
	ErrNeedRetry          = errors.New("need retry")
	ErrRepeatDeposit      = errors.New("repeat deposit")
	// ErrNetwork            = errors.New("network error")
)
