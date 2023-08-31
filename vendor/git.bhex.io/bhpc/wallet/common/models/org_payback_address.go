/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: org_payback_address.go
 * @Date: 2019/09/04
 * @Author: Jia Liu
 * @Copyright（C）: 2019 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

//  CREATE TABLE `org_payback_address` (
// 	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
// 	`remark` varchar(100) DEFAULT NULL,
// 	`created_at` timestamp NULL DEFAULT NULL,
// 	`updated_at` timestamp NULL DEFAULT NULL,
// 	`deleted_at` timestamp NULL DEFAULT NULL,
// 	`org_id` bigint(20) unsigned NOT NULL,
// 	`token_type` bigint(20) unsigned NOT NULL,
// 	`address` varchar(255) NOT NULL,
// 	`memo` varchar(60) DEFAULT NULL,
// 	`status` int(10) unsigned DEFAULT '0',
// 	`extension` varchar(255) DEFAULT NULL,
// 	PRIMARY KEY (`id`),
// 	KEY `uidx_address_memo_status` (`address`,`memo`,`status` )
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8

// OrgPaybackAddressStatus org payback address status enum
type OrgPaybackAddressStatus int64

const (
	// OrgPaybackAddressValid 有效
	OrgPaybackAddressValid OrgPaybackAddressStatus = 0
	// OrgPaybackAddressInvalid 无效
	OrgPaybackAddressInvalid OrgPaybackAddressStatus = 101
)

// OrgPaybackAddress save org pay back cold wallet address info
type OrgPaybackAddress struct {
	Base
	OrgID     int64                   `json:"org_id"`
	TokenType uint32                  `json:"token_type" gorm:"not null;"`
	Address   string                  `json:"address" gorm:"not null;type:varchar(255);index:uidx_address_memo_status"` // 地址
	Memo      string                  `json:"memo" gorm:"type:varchar(60);index:uidx_address_memo_status"`
	Status    OrgPaybackAddressStatus `json:"status" gorm:"default: 0;index:uidx_address_memo_status"`
	Extension string                  `json:"extension"`
}
