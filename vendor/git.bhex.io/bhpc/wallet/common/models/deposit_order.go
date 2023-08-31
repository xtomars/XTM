/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: deposit_order.go
 * @Date: 2018/07/20
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

import "github.com/shopspring/decimal"

// DepositOrder define the table deposit_order
type DepositOrder struct {
	Base
	TokenID        TokenIDEnum        `json:"token_id" gorm:"not null;unique_index:idx_depositorder_txhash_idx_tokenid"`                              // 币种
	From           string             `json:"from" gorm:"type:varchar(255);not null"`                                                                 // 从那个地址充进来
	To             string             `json:"to" gorm:"type:varchar(255);not null"`                                                                   // 转给那个地址
	TxHash         string             `json:"tx_hash" gorm:"not null;index:idx_depositorder_txhash;unique_index:idx_depositorder_txhash_idx_tokenid"` // 充值交易hash
	Index          uint32             `json:"index" gorm:"not null;unique_index:idx_depositorder_txhash_idx_tokenid"`                                 // 充值交易hash index
	Amount         decimal.Decimal    `json:"amount"  sql:"type:decimal(65,18);not null"`                                                             // 充值金额
	Confirm        uint32             `json:"confirm" gorm:"default: 0"`                                                                              // 交易确认数
	BlockHeight    int64              `json:"block_height" gorm:"default: 0"`                                                                         // 该交易所在区块高度
	BlockTime      int64              `json:"block_time" gorm:"default: 0"`                                                                           // 该交易所在区块时间
	Status         DepositOrderStatus `json:"status" gorm:"default: 0"`                                                                               // 订单状态：0 订单已创建 1 已过风控 10 已完成
	HadAccounted   bool               `json:"had_accounted"`                                                                                          // 是否已经入账
	HadPassRisk    bool               `json:"had_pass_risk"`                                                                                          // 是否已过风控
	Memo           string             `json:"memo"`
	Extension      string             `json:"extension"`                    // 扩展字段
	AssetDomain    uint64             `json:"asset_domain" gorm:"not null"` // 资产域
	OrgID          uint64             `json:"org_id" gorm:"not null"`       // Org ID
	Type           DepositType        `json:"type" gorm:"not null"`         // 充值类型
	ReservedStatus int64              `json:"reserved_status"`              // 保留状态字段，用来标记订单是否被正确刷新
}

// DepositOrderStatus Order status enum
type DepositOrderStatus uint32

const (
	// DepositOrderCreated 已创建
	DepositOrderCreated DepositOrderStatus = 0
	// DepositOrderPassRisk 已过风控
	DepositOrderPassRisk DepositOrderStatus = 1
	// DepositOrderFailRisk 没过风控
	DepositOrderFailRisk DepositOrderStatus = 11
	// DepositOrderCompleted 已完成
	DepositOrderCompleted DepositOrderStatus = 100
	// DepositOrderCanceld 手动取消
	DepositOrderCanceld DepositOrderStatus = 101
)

// DepositType deposit type
type DepositType uint32

const (
	// DepositOldStatus 未处理单子
	DepositOldStatus DepositType = 0
	// DepositToAllocatedUser 充值到已分配用户地址
	DepositToAllocatedUser DepositType = 1
	// DepositToUnallocatedUser 充值到未分配用户地址
	DepositToUnallocatedUser DepositType = 2
	// DepositToUserWrongTag 充值到已分配用户地址但tag不存在
	DepositToUserWrongTag DepositType = 3
	// DepositToHotAddr 充值到运营地址
	DepositToHotAddr DepositType = 4
	// CollectDeposit 归集
	CollectDeposit DepositType = 11
	// GiveFeeDeposit 打矿⼯费
	GiveFeeDeposit DepositType = 12
	// ReedemDeposit 找零
	ReedemDeposit DepositType = 13
	// NoNeedHandleDeposit 无需处理
	NoNeedHandleDeposit DepositType = 101
)
