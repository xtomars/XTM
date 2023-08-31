/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: balance.go
 * @Date: 2018/06/22
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

import "github.com/shopspring/decimal"

// PrepareAsset define the table 还没有被打包的资产
type PrepareAsset struct {
	Base
	Address string             `json:"address" gorm:"type:varchar(255);index:idx_prepareasset_address"` //找零地址，
	TokenID TokenIDEnum        `json:"token_id" gorm:"not null;"`                                       // 币种
	TxHash  string             `json:"tx_hash" gorm:"index:idx_prepareasset_hash"`                      // 提现交易hash
	Amount  decimal.Decimal    `json:"amount"  sql:"type:decimal(65,18);" gorm:"not null"`              // 提现金额
	Status  PrepareAssetStatus `json:"status" gorm:"default: 0;"`                                       //状态 0 创建， 1 已打包
}

// PrepareAssetStatus asset status enum
type PrepareAssetStatus uint32

const (
	// PrepareAssetStatusCreated 创建，
	PrepareAssetStatusCreated PrepareAssetStatus = 0
	// PrepareAssetStatusConfirmed 已打包
	PrepareAssetStatusConfirmed PrepareAssetStatus = 1
)
