/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: asset_report.go
 * @Date: 2018/06/21
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

import "github.com/shopspring/decimal"

// AssetReport define the table asset_report
type AssetReport struct {
	Base
	TokenID       string          `json:"token_id" gorm:"type:varchar(10);"`           // 当前统计币种
	Date          string          `json:"date" gorm:"type:varchar(20);"`               // 当前日期（按天统计）
	TotalIn       decimal.Decimal `json:"total_in"  sql:"type:decimal(65,18);"`        // 该币种的总入账数量
	TotalOut      decimal.Decimal `json:"total_out"  sql:"type:decimal(65,18);"`       // 该币种的总出账数量
	HotWalletIn   decimal.Decimal `json:"hot_wallet_in"  sql:"type:decimal(65,18);"`   // 热钱包总入账数量
	HotWalletOut  decimal.Decimal `json:"hot_wallet_out"  sql:"type:decimal(65,18);"`  // 热钱包总出账数量
	ColdWalletIn  decimal.Decimal `json:"cold_wallet_in"  sql:"type:decimal(65,18);"`  // 冷钱包总入账数量
	ColdWalletOut decimal.Decimal `json:"cold_walelt_out"  sql:"type:decimal(65,18);"` // 冷钱包总出账数量
}
