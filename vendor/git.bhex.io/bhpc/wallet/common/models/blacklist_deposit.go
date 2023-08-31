/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: blacklist_deposit.go
 * @Date: 2020/09/28
 * @Author: Jia Liu
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

import "github.com/shopspring/decimal"

// CREATE TABLE `blacklist_deposit` (
// 	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
// 	`remark` varchar(255) DEFAULT NULL,
// 	`created_at` timestamp NULL DEFAULT NULL,
// 	`updated_at` timestamp NULL DEFAULT NULL,
// 	`deleted_at` timestamp NULL DEFAULT NULL,
// 	`token_id` varchar(255) NOT NULL,
// 	`from` varchar(255) NOT NULL,
// 	`to` varchar(255) NOT NULL,
// 	`tx_hash` varchar(255) NOT NULL,
// 	`index` int(10) unsigned NOT NULL,
// 	`amount` decimal(65,18) NOT NULL,
// 	`confirm` int(10) unsigned DEFAULT '0',
// 	`block_height` bigint(20) unsigned DEFAULT '0',
// 	`block_time` bigint(20) unsigned DEFAULT '0',
// 	`status` int(10) unsigned DEFAULT '0',
// 	`can_account` tinyint(1) DEFAULT '0',
// 	`memo` varchar(1024) DEFAULT NULL,
// 	`extension` varchar(1024) DEFAULT NULL,
// 	PRIMARY KEY (`id`),
// 	UNIQUE KEY `unidx_deposit_tokenid_txhash` (`token_id`,`tx_hash`,`index`),
// 	KEY `idx_deposit_order_txhash` (`tx_hash`)
//   ) ENGINE=InnoDB  DEFAULT CHARSET=utf8

// BlacklistDeposit define the table black deposit
type BlacklistDeposit struct {
	Base
	TokenID     TokenIDEnum        `json:"token_id" gorm:"not null;unique_index:idx_blackdeposit_txhash_idx_tokenid"`                              // 币种
	From        string             `json:"from" gorm:"type:varchar(255);not null"`                                                                 // 从那个地址充进来
	To          string             `json:"to" gorm:"type:varchar(255);not null"`                                                                   // 转给那个地址
	TxHash      string             `json:"tx_hash" gorm:"not null;index:idx_blackdeposit_txhash;unique_index:idx_blackdeposit_txhash_idx_tokenid"` // 充值交易hash
	Index       uint32             `json:"index" gorm:"not null;unique_index:idx_blackdeposit_txhash_idx_tokenid"`                                 // 充值交易hash index
	Amount      decimal.Decimal    `json:"amount"  sql:"type:decimal(65,18);not null"`                                                             // 充值金额
	Confirm     uint32             `json:"confirm" gorm:"default: 0"`                                                                              // 交易确认数
	BlockHeight int64              `json:"block_height" gorm:"default: 0"`                                                                         // 该交易所在区块高度
	BlockTime   int64              `json:"block_time" gorm:"default: 0"`                                                                           // 该交易所在区块时间
	Status      BlackDepositStatus `json:"status" gorm:"default: 0"`                                                                               // 订单状态：
	CanAccount  bool               `json:"can_account"`                                                                                            // 是否可以入账

	Memo      string `json:"memo"`
	Extension string `json:"extension"` // 扩展字段

}

// BlackDepositStatus Order status enum
type BlackDepositStatus uint32

const (
	// BlackDepositCreated 已创建
	BlackDepositCreated BlackDepositStatus = 0
	// BlackDepositAccounted 已入账
	BlackDepositAccounted BlackDepositStatus = 100
	// BlackDepositCanceld 手动取消
	BlackDepositCanceld BlackDepositStatus = 101
)
