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

// CREATE TABLE `asset` (
// 	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
// 	`remark` varchar(100) DEFAULT NULL,
// 	`created_at` timestamp NULL DEFAULT NULL,
// 	`updated_at` timestamp NULL DEFAULT NULL,
// 	`deleted_at` timestamp NULL DEFAULT NULL,
// 	`address` varchar(255) NOT NULL,
// 	`token_id` varchar(255) NOT NULL,
// 	`balance` decimal(65,18) DEFAULT NULL,
// 	`pending_balance` decimal(65,18) DEFAULT NULL,
// 	`asset_domain` bigint(20) unsigned DEFAULT NULL,
// 	`org_id` bigint(20) unsigned DEFAULT NULL,
// 	`status` int(10) unsigned DEFAULT '0',
// 	`owner` int(10) DEFAULT '0',
// 	`extension` varchar(255) DEFAULT NULL,
// 	PRIMARY KEY (`id`),
// 	UNIQUE KEY `idx_address_tokenid` (`address`,`token_id`)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8

// Asset define the table asset
type Asset struct {
	Base
	Address        string          `json:"address" gorm:"not null;type:varchar(255);unique_index:idx_address_tokenid"` // 链上地址
	TokenID        TokenIDEnum     `json:"token_id" gorm:"not null;unique_index:idx_address_tokenid"`                  // 该地址绑定的币种
	Balance        decimal.Decimal `json:"balance"  sql:"type:decimal(65,18);" gorm:"not null;"`                       // 余额
	PendingBalance decimal.Decimal `json:"pending_balance"  sql:"type:decimal(65,18);" gorm:"not null;"`               // 在途资金（只计算归集产生的，不是归集产生的，不计算在内）
	AssetDomain    uint64          `json:"asset_domain" gorm:"not null"`                                               // 资产域
	OrgID          uint64          `json:"org_id" gorm:"not null"`                                                     // Org ID
	Status         AssetStatus     `json:"status" gorm:"default: 0"`                                                   // 状态，（0正常，1冻结）
	Owner          AssetOwnerType  `json:"owner" gorm:"not null;default: 1"`                                           // 钱包类型（1 个人 2 热钱包 3 冷钱包）
	Extension      string          `json:"extension"`                                                                  // 扩展字段
}

// AssetStatus asset status enum
type AssetStatus uint32

const (
	// Normal 正常
	Normal AssetStatus = 0
	// Frozen 冻结
	Frozen AssetStatus = 1
	// NeedTransfer 待转移
	NeedTransfer AssetStatus = 10
	// Transfering 转移中
	Transfering AssetStatus = 11
	// AssetDiscard 废弃地址
	AssetDiscard AssetStatus = 101
)

// AssetOwnerType asset owner status enum
type AssetOwnerType uint32

const (
	// UserAsset 用户地址资产
	UserAsset AssetOwnerType = 1
	// DomainHotAsset 资产域运营地址资产
	DomainHotAsset AssetOwnerType = 2
	// OrgHotAsset Org运营地址资产
	OrgHotAsset AssetOwnerType = 3
)
