/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: notify.go
 * @Date: 2020/02/12
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package common

// NotifyDepositParams define
type NotifyDepositParams struct {
	TokenID     string `json:"token_id"`
	From        string `json:"from"`
	To          string `json:"to"`
	Memo        string `json:"memo"`
	Amount      string `json:"amount"`
	TxHash      string `json:"tx_hash"`
	Index       string `json:"index"`
	BlockHeight string `json:"block_height"`
	BlockTime   string `json:"block_time"`
}

// NotifyWithdrawalParams define
type NotifyWithdrawalParams struct {
	OrderID     string `json:"order_id"`
	TokenID     string `json:"token_id"`
	To          string `json:"to"`
	Memo        string `json:"memo"`
	Amount      string `json:"amount"`
	Fee         string `json:"fee"`
	TxHash      string `json:"tx_hash"`
	BlockHeight string `json:"block_height"`
	BlockTime   string `json:"block_time"`
}

// NotifyFailedParams define
type NotifyFailedParams struct {
	OrderID string `json:"order_id"`
	TokenID string `json:"token_id"`
	Reason  string `json:"reason"`
}
