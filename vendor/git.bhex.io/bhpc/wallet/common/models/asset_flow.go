/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: balance_flow.go
 * @Date: 2018/07/20
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

import "github.com/shopspring/decimal"

// AssetFlow define the table asset_flow
type AssetFlow struct {
	Base
	AssetID   uint64          `json:"asset_id" gorm:"not null;"`                            // 账户ID
	Type      AssetFlowType   `json:"type" gorm:"not null;"`                                // 流水类型：1 充值流水 2 提现流水 3 归集流水 4 商户互转
	OrderID   uint64          `json:"order_id" gorm:"not null;"`                            // 流水对应的订单ID 若为商户互转，订单ID为空
	Changed   decimal.Decimal `json:"changed"  sql:"type:decimal(65,18);" gorm:"not null;"` // 变动金额
	Total     decimal.Decimal `json:"total"  sql:"type:decimal(65,18);" gorm:"not null;"`   // 变动之后账户总余额
	OrderIDs  string          `json:"order_ids"  sql:"type:varchar(255);" gorm:"null;"`     // 批量提现的所有orderID
	UTXOID    uint64          `json:"utxo_id" `                                             // UTXO ID(自己管理utxo类型的币种适用)
	HadRefund bool            `json:"had_refund"`                                           // 是否已经做过退款处理
}

// AssetFlowType asset flow type enum
type AssetFlowType int32

const (
	// Deposit 充值流水
	Deposit AssetFlowType = 1
	// DepositCancel 取消充值流水
	DepositCancel AssetFlowType = -1
	// Withdrawal 提现流水
	Withdrawal AssetFlowType = 2
	// WithdrawalCancel 取消提现流水
	WithdrawalCancel AssetFlowType = -2
	// Collect 资金归集流水
	Collect AssetFlowType = 3
	// InternalTransfer 商户互转
	InternalTransfer AssetFlowType = 4
	// ChangeContractClearAsset 更换合约冲账清零
	ChangeContractClearAsset AssetFlowType = 5
)
