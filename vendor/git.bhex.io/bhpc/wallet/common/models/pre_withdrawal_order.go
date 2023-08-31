/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: pre_withdrawal_order.go
 * @Date: 2019/07/11
 * @Author: Jia Liu
 * @Copyright（C）: 2019 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

// CREATE TABLE `pre_withdrawal_order` (
//   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
//   `created_at` timestamp NULL DEFAULT NULL,
//   `updated_at` timestamp NULL DEFAULT NULL,
//   `deleted_at` timestamp NULL DEFAULT NULL,
//   `org_id` bigint(21) unsigned DEFAULT NULL,
//   `account_id` bigint(21) unsigned DEFAULT '0',
//   `broker_order_id` bigint(21) unsigned NOT NULL,
//   `token_id` varchar(255) NOT NULL,
//   `to` varchar(255) NOT NULL,
//   `memo` varchar(255) DEFAULT NULL,
//   `amount` decimal(65,18) DEFAULT NULL,
//   `cost_fee` decimal(65,18) DEFAULT NULL,
//   `sign_time` bigint(21) DEFAULT NULL,
//   `sign_nonce` varchar(2048) DEFAULT NULL,
//   `sign_bh` varchar(2048) DEFAULT NULL,
//   `sign_rc` varchar(2048) DEFAULT NULL,
//   `extension` varchar(2048) DEFAULT NULL,
//   `status` bigint(20) DEFAULT NULL,
//   `` bigint(20) DEFAULT NULL,
//   `remark` varchar(2048) DEFAULT NULL,
//   PRIMARY KEY (`id`),
//   UNIQUE KEY `broker_order_id_UNIQUE` (`broker_order_id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8

// PreWithdrawalOrderStatus pre withdrawal Order status enum
type PreWithdrawalOrderStatus int64

const (
	// PreWithdrawalOrderCreated 已创建
	PreWithdrawalOrderCreated PreWithdrawalOrderStatus = 0
	// PreWithdrawalOrderHandling 处理中
	PreWithdrawalOrderHandling PreWithdrawalOrderStatus = 10
	// PreWithdrawalOrderRetrying 重试中
	PreWithdrawalOrderRetrying PreWithdrawalOrderStatus = 20
	// PreWithdrawalOrderSuccess 处理成功
	PreWithdrawalOrderSuccess PreWithdrawalOrderStatus = 100
	// PreWithdrawalOrderInvalid 订单无效，无效原因可以查看remark字段
	PreWithdrawalOrderInvalid PreWithdrawalOrderStatus = 102
	// PreWithdrawalOrderRetryFailed retry failed
	PreWithdrawalOrderRetryFailed PreWithdrawalOrderStatus = 103
)

// PreWithdrawalOrder save order info received from platform
type PreWithdrawalOrder struct {
	Base
	OrgID         int64                    `json:"org_id"`
	AccountID     int64                    `json:"account_id"`
	BrokerOrderID int64                    `json:"broker_order_id"  gorm:"not null; unique_index:broker_order_id_UNIQUE"` // 发起提现请求的broker订单ID
	TokenID       TokenIDEnum              `json:"token_id"  gorm:"not null"`                                             // 币种
	To            string                   `json:"to" gorm:"type:varchar(255); not null"`                                 // 转给那个地址
	Memo          string                   `json:"memo"`
	Amount        string                   `json:"amount"` // 提现金额
	CostFee       string                   `json:"cost_fee"`
	SignTime      int64                    `json:"sign_time"`
	SignNonce     string                   `json:"sign_nonce"`
	SignBh        string                   `json:"sign_bh"`
	SignRc        string                   `json:"sign_rc"`
	Extension     string                   `json:"extension"`                // 扩展字段
	Status        PreWithdrawalOrderStatus `json:"status" gorm:"default: 0"` // 订单状态：0 订单已创建
	RetryCount    int64                    `json:"retry_count"`              // 重试次数
}
