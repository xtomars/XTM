/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: asset.go
 * @Date: 2020/02/14
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package common

// AssetInfo define
type AssetInfo struct {
	TokenID               string `json:"token_id"`
	TotalDepositAmount    string `json:"total_deposit_amount"`
	TotalWithdrawalAmount string `json:"total_withdrawal_amount"`
	LastBlockHeight       string `json:"last_block_height"`
}

// AssetVerifyResult define
type AssetVerifyResult struct {
	Success             bool   `json:"success"`
	DBDepositAmount     string `json:"db_deposit_amount"`
	ReqDepositAmount    string `json:"req_deposit_amount"`
	DBWithdrawAmount    string `json:"db_withdrawal_amount"`
	ReqWithdrawalAmount string `json:"req_withdrawal_amount"`
}
