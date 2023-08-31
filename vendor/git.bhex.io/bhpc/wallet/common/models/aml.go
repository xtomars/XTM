/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: aml.go
 * @Date: 2020/05/26
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

import "time"

// SlowMistBlackAddress store slow mist black address
type SlowMistBlackAddress struct {
	Base
	Coin    string    `json:"coin" gorm:"coin"`
	Address string    `json:"address" gorm:"address"`
	Tag     string    `json:"tag" gorm:"tag"`
	Detail  string    `json:"detail" gorm:"detail"`
	AddTime time.Time `json:"add_time" gorm:"add_time"`
}
