/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: address_report.go
 * @Date: 2018/06/21
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

// AddressReport define the table address_report
type AddressReport struct {
	Base
	TokenType     uint64 `json:"token_type"`                    // 地址所属币种类型
	Date          string `json:"date" gorm:"type:varchar(20);"` // 统计日期当前日期（按天统计）
	AdddressUsed  uint64 `json:"address_used"`                  // 当天的地址派发数量
	AddressRemain uint64 `json:"address_remain"`                // 当前可用地址剩余量
}
