/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: org_hot_address.go
 * @Date: 2020/04/07
 * @Author: jia.liu
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

// CREATE TABLE `org_hot_address` (
// 	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
// 	`remark` varchar(255) DEFAULT NULL,
// 	`created_at` timestamp NULL DEFAULT NULL,
// 	`updated_at` timestamp NULL DEFAULT NULL,
// 	`deleted_at` timestamp NULL DEFAULT NULL,
// 	`org_id` bigint(20) unsigned NOT NULL,
// 	`token_type` bigint(20) unsigned NOT NULL,
// 	`address` varchar(255) NOT NULL,
// 	`memo` varchar(255) DEFAULT NULL,
// 	`extension` varchar(255) DEFAULT NULL,
//  `dummy_address` tinyint(1) DEFAULT '0',
// 	PRIMARY KEY (`id`),
// 	KEY `idx_org_type` (`org_id`,`token_type`),
//  KEY `address` (`address`)
//   ) ENGINE=InnoDB  DEFAULT CHARSET=utf8

// OrgHotAddress define the table for org hot address
type OrgHotAddress struct {
	Base
	OrgID        uint64    `json:"org_id" gorm:"not null;unique_index:idx_org_type"`        // 账号ID
	TokenType    TokenType `json:"token_type" gorm:"not null;unique_index:idx_org_type"`    // 地址类型
	Address      string    `json:"address" gorm:"not null;type:varchar(255);index:address"` // 地址
	Memo         string    `json:"memo" gorm:"type:varchar(255)"`
	Extension    string    `json:"extension"  gorm:"type:varchar(255);"` // 扩展字段
	DummyAddress bool      `json:"dummy_address"`                        // 是否逻辑上的虚拟地址，这样的地址同一token type可以根据情况决定是否使用dummy地址
}
