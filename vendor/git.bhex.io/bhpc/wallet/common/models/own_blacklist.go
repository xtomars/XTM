/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: own_blacklist.go
 * @Date: 2020/05/26
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

// OwnBlacklistAddress store oblack address
type OwnBlacklistAddress struct {
	Base
	TokenID        string `json:"token_id" gorm:"token_id"`
	Address        string `json:"address" gorm:"address"`
	IsBlackAddress bool   `json:"is_black_address" gorm:"is_black_address"`
	Level          int    `json:"level" gorm:"level"`
	Score          int    `json:"score" gorm:"score"`
	OrderType      int    `json:"order_type" gorm:"order_type"`
	OrderID        uint64 `json:"order_id" gorm:"order_id"`
}
