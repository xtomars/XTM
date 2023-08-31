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

// Account define the table account
type Account struct {
	Base
	OrgID uint64 `json:"org_id" gorm:"not null"` // 组织ID
}
