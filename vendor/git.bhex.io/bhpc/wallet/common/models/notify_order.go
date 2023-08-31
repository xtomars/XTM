/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: notify_order.go
 * @Date: 2018/07/20
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

// NotifyOrder define define the table notify_order
type NotifyOrder struct {
	Base
	Type       NotifyType   `json:"type" gorm:"not null;unique_index:idx_notifyorder_orderid"`     // 通知类型 1 充值类型 2 提现类型 3 归集类型
	OrderID    uint64       `json:"order_id" gorm:"not null;unique_index:idx_notifyorder_orderid"` // 充值/提现/归集 订单ID
	RetryCount uint32       `json:"retyr_count" gorm:"default: 0"`                                 // 重试次数
	Status     NotifyStatus `json:"status" gorm:"default: 0"`                                      // 通知订单状态：1 已创建 2 充值（可交易）通知成功 3 充值（可交易）重试 4 充值（可提现）通知成功 5 充值（可提现）重试 6 提现成功 7 提现重试 8 失败
	Extension  string       `json:"extension"`                                                     // 扩展字段
}

// NotifyType token status enum
type NotifyType uint32

const (
	// NotifyTypeDeposit 充值
	NotifyTypeDeposit NotifyType = 1
	// NotifyTypeWithdrawal 提现
	NotifyTypeWithdrawal NotifyType = 2
	// NotifyTypeCollect 归集
	NotifyTypeCollect NotifyType = 3
)

// NotifyStatus token status enum
type NotifyStatus uint32

const (
	// NSCreated 已创建
	NSCreated NotifyStatus = 0
	// NSDeposit1Success 充值订单可交易，通知成功
	NSDeposit1Success NotifyStatus = 1
	// NSDeposit1Retrying 充值订单可交易，重试中
	NSDeposit1Retrying NotifyStatus = 11
	// NSDeposit2Success 充值订单可提现，通知成功
	NSDeposit2Success NotifyStatus = 2
	// NSDeposit2Retrying 充值订单可提现，重试中
	NSDeposit2Retrying NotifyStatus = 12
	// NSWithdrawalSuccess 提现订单状态，通知成功
	NSWithdrawalSuccess NotifyStatus = 3
	// NSWithdrawalRetrying 提现订单状态，重试中
	NSWithdrawalRetrying NotifyStatus = 13

	//NSOnChainSuccess 充值订单已上链确认，通知成功
	NSOnChainSuccess NotifyStatus = 4
	// NSOnChainRetrying 充值订单已上链确认，重试中
	NSOnChainRetrying NotifyStatus = 14

	// NSFail 失败
	NSFail NotifyStatus = 100
)
