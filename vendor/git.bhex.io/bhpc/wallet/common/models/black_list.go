/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: black_list.go
 * @Date: 2020/09/28
 * @Author: Jia Liu
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

// Blacklist define the table black list address
type Blacklist struct {
	Base
	AddressMemo string `json:"address_memo" gorm:"not null;type:varchar(1024);unique_index"` // 地址
	IsMemo      bool   `json:"is_memo" gorm:"not null"`                                      // 是否是memo
	Extension   string `json:"extension"`                                                    // 扩展字段
}
