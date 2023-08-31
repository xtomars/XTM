/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: private_key.go
 * @Date: 2018/07/20
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

// PrivateKey define the table private_key
type PrivateKey struct {
	Base
	Address   string `json:"address" gorm:"type:varchar(255);unique_index:idx_address_seqno"` // 地址
	Seqno     uint32 `json:"seqno" gorm:"unique_index:idx_address_seqno"`                     // 该地址的私钥片段号
	PartedKey []byte `json:"parted_key"`                                                      // 该地址的私钥片段加密内容
	CheckSign string `json:"check_sign"`                                                      // 私钥正确性的签名
	Extension string `json:"extension"`                                                       // 扩展字段
	Version   uint64 `gorm:"not null"`
}
