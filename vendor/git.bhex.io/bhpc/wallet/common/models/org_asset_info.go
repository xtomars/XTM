/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: org_platform_asset_info.go
 * @Date: 2020/04/01
 * @Author: Jia Liu
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

import "github.com/shopspring/decimal"

//  CREATE TABLE `org_asset_info` (
// 	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
// 	`remark` varchar(255) DEFAULT NULL,
// 	`created_at` timestamp NULL DEFAULT NULL,
// 	`updated_at` timestamp NULL DEFAULT NULL,
// 	`deleted_at` timestamp NULL DEFAULT NULL,
// 	`org_id` bigint(20) unsigned NOT NULL,
// 	`token_id` varchar(255)  NOT NULL,
// 	`latest_platform_balance` decimal(65,18) NOT NULL,
// 	`asset_percent_low` int(10) DEFAULT 0,
// 	`asset_percent_high` int(10) DEFAULT 0,
// 	`rebalance_mark` int(10) DEFAULT 0,
// 	PRIMARY KEY (`id`),
// 	UNIQUE KEY `idx_orgid_tokenid` (`org_id`,`token_id`)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

// OrgAssetInfo org asset info
type OrgAssetInfo struct {
	Base
	OrgID                 uint64          `json:"org_id" gorm:"not null;unique_index:idx_orgid_tokenid"`               // orgID
	TokenID               TokenIDEnum     `json:"token_id" gorm:"not null;unique_index:idx_orgid_tokenid"`             // TokenID
	LatestPlatformBalance decimal.Decimal `json:"latest_platform_balance"  sql:"type:decimal(65,18);" gorm:"not null"` // 最新平台资产余额
	AssetPercentLow       int             `json:"asset_percent_low" gorm:"not null"`                                   // 资产百分比下限
	AssetPercentHigh      int             `json:"asset_percent_high" gorm:"not null"`                                  // 资产百分比上限
	RebalanceMark         int             `json:"rebalance_mark" gorm:"not null"`                                      // 此分数决定是否需要rebalance
}
