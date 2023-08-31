/*
 * *******************************************************************
 * @项目名称: privatekeyserver
 * @文件名称: encrypted_key.go
 * @Date: 2018/09/23
 * @Author: Jia Liu
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */
package models

// CREATE TABLE `encrypted_key` (
// 	`id` int(20) unsigned NOT NULL AUTO_INCREMENT,
// 	`pwd_hash` varbinary(20) NOT NULL,
// 	`pwd_version` bigint(20) unsigned NOT NULL,
// 	`key_cipher` varbinary(128) NOT NULL,
// 	`key_code` varbinary(20) DEFAULT NULL,
// 	`version` bigint(20) unsigned NOT NULL,
// 	`created_at` timestamp NULL DEFAULT NULL,
// 	`updated_at` timestamp NULL DEFAULT NULL,
// 	PRIMARY KEY (`id`)
//   )
type EncryptedKey struct {
	Base
	PwdHash      []byte `gorm:"not null"` // password hash
	PwdPart1Hash []byte //part1 pwd hash
	PwdPart2Hash []byte // part2 pwd hash
	PwdVersion   uint64 `gorm:"not null"` // password 版本
	KeyCipher    []byte `gorm:"not null"` // 加密密钥的密文
	KeyCode      []byte `gorm:"not null"` // 加密密钥的hash
	Version      uint64 `gorm:"not null"` // 加密密钥的版本

}
