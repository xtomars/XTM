/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: withdrawal_order.go
 * @Date: 2018/07/20
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// CREATE TABLE `withdrawal_order` (
// 	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
// 	`remark` varchar(100) DEFAULT NULL,
// 	`created_at` timestamp NULL DEFAULT NULL,
// 	`updated_at` timestamp NULL DEFAULT NULL,
// 	`deleted_at` timestamp NULL DEFAULT NULL,
// 	`broker_order_id` bigint(20) unsigned NOT NULL,
// 	`token_id` varchar(255) NOT NULL,
// 	`gased_for_token_id` varchar(255) DEFAULT NULL,
// 	`from` varchar(255) DEFAULT NULL,
// 	`to` varchar(255) NOT NULL,
// 	`amount` decimal(65,18) unsigned NOT NULL,
// 	`tx_hash` varchar(255) DEFAULT NULL,
// 	`confirm` int(10) unsigned DEFAULT NULL,
// 	`block_height` bigint(20) unsigned DEFAULT NULL,
// 	`block_time` bigint(20) unsigned DEFAULT NULL,
// 	`status` bigint(20) unsigned DEFAULT '0',
// 	`send_chain_at` timestamp NULL DEFAULT NULL,
// 	`cost_fee` decimal(65,18) unsigned DEFAULT NULL,
// 	`utxo_ids` varchar(255) DEFAULT NULL,
// 	`utxo_redeem` varchar(255) DEFAULT NULL,
// 	`amount_redeem` decimal(65,18) unsigned DEFAULT NULL,
// 	`memo` varchar(60) DEFAULT NULL,
// 	`extension` varchar(255) DEFAULT NULL,
// 	`asset_domain` bigint(20) unsigned DEFAULT NULL,
// 	`org_id` bigint(20) unsigned DEFAULT NULL,
// 	`account_id` bigint(21) DEFAULT '0',
// 	`type` int(10) DEFAULT NULL,
// 	`gas_price` decimal(65,18) unsigned DEFAULT '0.000000000000000000',
// 	`gas_used` varchar(45) DEFAULT '0',
// 	PRIMARY KEY (`id`),
// 	UNIQUE KEY `uix_broker_order_id` (`broker_order_id`),
// 	KEY `idx_withdrawal_order_txhash` (`tx_hash`),
// 	KEY `idx_tokenid_createdat` (`token_id`,`created_at`),
// 	KEY `idx_createdat` (`created_at`),
// 	KEY `idx_status_id` (`status`,`id`)
//   ) ENGINE=InnoDB  DEFAULT CHARSET=utf8

// WithdrawalOrder define define the table withdrawal_order
type WithdrawalOrder struct {
	Base
	BrokerOrderID   int64                 `json:"broker_order_id"  gorm:"not null; unique_index:uidx_withdrawalorder_broorderid"` // 发起提现请求的broker订单ID
	AssetDomain     uint64                `json:"asset_domain"`
	OrgID           uint64                `json:"org_id"`
	AccountID       int64                 `json:"account_id"`
	TokenID         TokenIDEnum           `json:"token_id"  gorm:"not null"`                          // 币种
	GasedForTokenID TokenIDEnum           `json:"gased_for_token_id"  gorm:"not null"`                // 给gas到的币种
	CostFeeTokenID  TokenIDEnum           `json:"cost_fee_token_id"`                                  // 矿工费币种
	From            string                `json:"from" gorm:"type:varchar(255)"`                      // 从那个地址转出去
	To              string                `json:"to" gorm:"type:varchar(255); not null"`              // 转给那个地址
	Amount          decimal.Decimal       `json:"amount"  sql:"type:decimal(65,18);" gorm:"not null"` // 提现金额
	TxHash          string                `json:"tx_hash" gorm:"index:idx_withdrawalorder_hash"`      // 提现交易hash
	Confirm         uint32                `json:"confirm"`                                            // 交易确认数
	BlockHeight     int64                 `json:"block_height"`                                       // 该交易所在区块高度
	BlockTime       int64                 `json:"block_time"`                                         // 该交易所在区块时间
	Status          WithdrawalOrderStatus `json:"status" gorm:"default: 0"`                           // 订单状态：0 订单已创建 1 已过风控 2 成功发送ChainNode 3 已上链 10 已完成
	SendChainAt     time.Time             `json:"send_chain_at,omitempty"`                            // 成功发到链上的时间，用来计算耗时
	CostFee         decimal.Decimal       `json:"cost_fee" sql:"type:decimal(65,18);"`                // 提现手续费
	GasPrice        decimal.Decimal       `json:"gas_price" sql:"type:decimal(65,18);"`               // 提现手续费单价
	GasUsed         int64                 `json:"gas_used"`                                           // 提现消耗的gas
	UtxoIDs         string                `json:"utxo_ids"`                                           // 若提现类型为UTXO: 使用的utxo对应的ID
	UtxoRedeem      string                `json:"utxo_redeem"`                                        // 若提现类型为UTXO: 指定的找零地址
	AmountRedeem    decimal.Decimal       `json:"amount_redeem" sql:"type:decimal(65,18);"`           // 找零金额
	Memo            string                `json:"memo"`
	Extension       string                `json:"extension"`       // 扩展字段
	Type            WithdrawalType        `json:"type"`            // 提现类型
	ReservedStatus  int64                 `json:"reserved_status"` // 保留状态字段，用来标记订单是否被正确刷新
}

// WithdrawalOrderStatus Order status enum
type WithdrawalOrderStatus int64

const (
	// WithdrawalOrderCreated 已创建
	WithdrawalOrderCreated WithdrawalOrderStatus = 0
	// WithdrawalOrderPassRisk 已过风控
	WithdrawalOrderPassRisk WithdrawalOrderStatus = 1
	// WithdrawalOrderFailRisk 没过风控
	WithdrawalOrderFailRisk WithdrawalOrderStatus = 11
	// WithdrawalOrderSendChainnode 已发送给 ChainNode
	WithdrawalOrderSendChainnode WithdrawalOrderStatus = 2
	// WithdrawalOrderFailChainnode ChainNode返回失败
	WithdrawalOrderFailChainnode WithdrawalOrderStatus = 12
	// WithdrawalOrderMined 已上链
	WithdrawalOrderMined WithdrawalOrderStatus = 4
	// WithdrawalOrderFailOnChain 上链失败
	WithdrawalOrderFailOnChain WithdrawalOrderStatus = 14
	// WithdrawalOrderNeedBaasHandle 等待baas拉取处理该订单
	WithdrawalOrderNeedBaasHandle WithdrawalOrderStatus = 8
	// WithdrawalOrderBaasHandleFail baas 处理订单失败
	WithdrawalOrderBaasHandleFail WithdrawalOrderStatus = 18
	// WithdrawalOrderCompleted 已完成
	WithdrawalOrderCompleted WithdrawalOrderStatus = 100
	// WithdrawalOrderCanceled 手动取消
	WithdrawalOrderCanceled WithdrawalOrderStatus = 101
	// WithdrawalOrderInvalid 收到平台的订单，初始状态为WithdrawalOrderInvalid
	WithdrawalOrderInvalid WithdrawalOrderStatus = 102
	// WithdrawalOrderExecFailedOnChain 链上执行失败，此时扣了手续费
	WithdrawalOrderExecFailedOnChain WithdrawalOrderStatus = 103
	// WithdrawalOrderMaxFinalStatus max final status
	WithdrawalOrderMaxFinalStatus WithdrawalOrderStatus = WithdrawalOrderExecFailedOnChain
)

// WithdrawalType 归集类型
type WithdrawalType int32

const (
	// WithdrawalTypeUser broker过来的用户提现
	WithdrawalTypeUser WithdrawalType = 1
	// WithdrawalTypeCollect 普通归集
	WithdrawalTypeCollect WithdrawalType = 2
	// WithdrawalTypeDustCollect dust utxo 归集
	WithdrawalTypeDustCollect WithdrawalType = 3
	// WithdrawalTypeOnlyCostFeeCollect 钱包节点自动归集，需要扣掉手续费
	WithdrawalTypeOnlyCostFeeCollect WithdrawalType = 4
	// WithdrawalTypeGiveGas give gas
	WithdrawalTypeGiveGas WithdrawalType = 10
)
