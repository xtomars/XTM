/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: org.go
 * @Date: 2019/10/25
 * @Author: Jia Liu
 * @Copyright（C）: 2019 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

import "time"

//  CREATE TABLE `org` (
// 	`org_id` bigint(20) unsigned NOT NULL,
// 	`remark` varchar(255) DEFAULT NULL,
// 	`created_at` timestamp NULL DEFAULT NULL,
// 	`updated_at` timestamp NULL DEFAULT NULL,
// 	`deleted_at` timestamp NULL DEFAULT NULL,
// 	`asset_domain` bigint(20) unsigned NOT NULL,
// 	`asset_percent` int(10) DEFAULT NULL,
// 	`deposit` varchar(255) DEFAULT NULL,
// 	`withdrawal` varchar(255) DEFAULT NULL,
// 	`notify` varchar(255) DEFAULT NULL,
// 	`givefee` varchar(255) DEFAULT NULL,
// 	`autocollect` varchar(255) DEFAULT NULL,
// 	PRIMARY KEY (`org_id`)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8

// Org org info
type Org struct {
	OrgID        uint64    `json:"org_id" gorm:"primary_key"`                  // 记录ID
	Remark       string    `json:"remark,omitempty" gorm:"type:varchar(255);"` // 备注字段
	CreatedAt    time.Time `json:"created_at,omitempty"`                       // 记录创建时间
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	AssetDomain  uint64    `json:"asset_domain" gorm:"not null"`         // 资产域ID
	AssetPercent int       `json:"asset_percent" gorm:"not null"`        // 资产百分比
	Deposit      string    `json:"deposit" gorm:"type:varchar(255)"`     // 充值开关设置
	Withdrawal   string    `json:"withdrawal" gorm:"type:varchar(255)"`  // 提币开关设置
	Notify       string    `json:"notify" gorm:"type:varchar(255)"`      // 通知开关设置
	GiveFee      string    `json:"givefee" gorm:"type:varchar(255)"`     // 自动给矿工费开关
	AutoCollect  string    `json:"autocollect" gorm:"type:varchar(255)"` // 自动归集开关
}
