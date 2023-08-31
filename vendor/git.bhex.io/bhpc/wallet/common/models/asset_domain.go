/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: asset_domain.go
 * @Date: 2019/10/27
 * @Author: jia.liu
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

// CREATE TABLE `asset_domain` (
// 	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
// 	`remark` varchar(255) DEFAULT NULL,
// 	`created_at` timestamp NULL DEFAULT NULL,
// 	`updated_at` timestamp NULL DEFAULT NULL,
// 	`deleted_at` timestamp NULL DEFAULT NULL,
// 	`deposit` varchar(255) DEFAULT NULL,
// 	`withdrawal` varchar(255) DEFAULT NULL,
// 	`notify` varchar(255) DEFAULT NULL,
// 	`givefee` varchar(255) DEFAULT NULL,
// 	`autocollect` varchar(255) DEFAULT NULL,
// 	`rule1` varchar(255) DEFAULT NULL,
// 	`rule2` varchar(255) DEFAULT NULL,
// 	PRIMARY KEY (`id`)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8

// AssetDomain define the table of asset domain
type AssetDomain struct {
	Base
	Deposit     string `json:"deposit" gorm:"type:varchar(255)"`     // 充值开关设置
	Withdrawal  string `json:"withdrawal" gorm:"type:varchar(255)"`  // 提币开关设置
	Notify      string `json:"notify" gorm:"type:varchar(255)"`      // 通知开关设置
	GiveFee     string `json:"givefee" gorm:"type:varchar(255)"`     // 自动给矿工费开关
	AutoCollect string `json:"autocollect" gorm:"type:varchar(255)"` // 自动归集开关
}
