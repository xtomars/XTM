/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: order.go
 * @Date: 2020/02/12
 * @Author: zhiming.sun
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package common

// WithdrawalOrder define
type WithdrawalOrder struct {
	OrderID   string `json:"order_id"`
	TokenID   string `json:"token_id"`
	To        string `json:"to"`
	Memo      string `json:"memo"`
	Amount    string `json:"amount"`
	TimeStamp string `json:"time_stamp"`
	Signature string `json:"signature"`
}
