/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: account.go
 * @Date: 2018/08/24
 * @Author: qiangjun.chen
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */
package models

// AccountAddress define the table account
type AccountAddress struct {
	Base
	AccountID uint64    `json:"account_id" gorm:"not null;unique_index:idx_account"`                      // 账号ID
	TokenType TokenType `json:"token_type" gorm:"not null;unique_index:idx_account"`                      // 地址类型
	Address   string    `json:"address" gorm:"not null;type:varchar(255);unique_index:uidx_address_memo"` // 地址
	Memo      string    `json:"memo" gorm:"type:varchar(60);unique_index:uidx_address_memo"`
	Extension string    `json:"extension"  gorm:"type:varchar(6000);"` // 扩展字段
}
