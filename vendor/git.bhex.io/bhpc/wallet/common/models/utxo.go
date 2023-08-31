/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: utxo.go
 * @Date: 2018/07/20
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

import "github.com/shopspring/decimal"

// UTXO define define the table utxo
type UTXO struct {
	Base
	TokenID   TokenIDEnum     `json:"token_id" gorm:"not null;index"`                      // 对应的币种ID
	Address   string          `json:"address" gorm:"not null; type:varchar(255); index"`   // 地址
	Amount    decimal.Decimal `json:"amount"  sql:"type:decimal(65,18);" gorm:"not null;"` // 该UTXO的值
	TxHash    string          `json:"tx_hash" gorm:"not null;unique_index:idx_utxo_hash"`  // 该UTXO对应的交易hash
	Index     uint32          `json:"index" gorm:"not null;unique_index:idx_utxo_hash"`    // UTXO在交易中的位置
	Confirm   uint32          `json:"confirm" gorm:"default: 0"`                           // 交易确认数
	Status    UtxoStatus      `json:"status"  gorm:"not null"`                             // 该UTXO状态： 0 未花费 1 已花费 2 已占用
	MuxID     string          `json:"mux_id"  gorm:"not null"`                             // 该UTXO状态： 0 未花费 1 已花费 2 已占用
	Extension string          `json:"extension"`                                           // btm 专用的
}

// UtxoStatus utxo status enum
type UtxoStatus uint32

const (
	// UnSpent 未花费
	UnSpent UtxoStatus = 0
	// Spent 已花费
	Spent UtxoStatus = 1
	// Occupy 已占用
	Occupy UtxoStatus = 2
)
