/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: collect_order.go
 * @Date: 2018/07/24
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

// CollectOrder define the table collect_order
type CollectOrder struct {
	Base
	TokenID       TokenIDEnum       `json:"token_id" `                              // token id
	Broker        string            `json:"broker"`                                 // 归集发起者
	BrokerOrderID int64             `json:"broker_order_id" gorm:"not null;unique"` // 归集唯一订单ID
	Reason        CollectReasonEnum `json:"reason"`                                 // 归集原因
	ReasonType    CollectReasonType `json:"reason_type"`                            // 归集原因类型
	Extension     string            `json:"extension"`                              // 扩展字段
}

// CollectReasonEnum token id 的类型
type CollectReasonEnum string

const (
	// User2service 用户钱包到营运钱包
	User2service CollectReasonEnum = "User2service"
	// Service2cold 营运钱包到冷钱包
	Service2cold CollectReasonEnum = "Service2cold"
	// GiveUSDTFee 给USDT矿工费
	GiveUSDTFee CollectReasonEnum = "GiveUSDTFee"
	// GiveErc20Fee 给erc20矿工费
	GiveErc20Fee CollectReasonEnum = "GiveErc20Fee"
	// GiveOntOngFee 给Ont矿工费
	GiveOntOngFee CollectReasonEnum = "GiveOntOngFee"
	// PresureTest 回环测试
	PresureTest CollectReasonEnum = "PresureTest"

	// AutoUser2service 自动归集用户钱包到营运钱包
	AutoUser2service CollectReasonEnum = "AutoUser2service"
	// AutoService2cold 自动归集营运钱包到冷钱包
	AutoService2cold CollectReasonEnum = "AutoService2cold"
	// AutoGiveErc20Fee 自动归集给erc20矿工费
	AutoGiveErc20Fee CollectReasonEnum = "AutoGiveErc20Fee"
	// AutoCollectDustUTXO 自动归集Dust UTXO
	AutoCollectDustUTXO CollectReasonEnum = "AutoCollectDustUTXO"
)

// CollectReasonType 类型
type CollectReasonType uint32

const (
	// InvalidType 无效类型
	InvalidType CollectReasonType = 0
	// User2serviceType 用户钱包到营运钱包
	User2serviceType CollectReasonType = 1
	// Service2coldType 营运钱包到冷钱包
	Service2coldType CollectReasonType = 2
	// GiveUSDTFeeType 给USDT矿工费
	GiveUSDTFeeType CollectReasonType = 3
	// GiveErc20FeeType 给erc20矿工费
	GiveErc20FeeType CollectReasonType = 4
	// PresureTestType 回环测试
	PresureTestType CollectReasonType = 5

	// AutoUser2serviceType 自动归集用户钱包到营运钱包
	AutoUser2serviceType CollectReasonType = 6
	// AutoService2coldType 自动归集营运钱包到冷钱包
	AutoService2coldType CollectReasonType = 7
	// AutoGiveErc20FeeType 自动归集给erc20矿工费
	AutoGiveErc20FeeType CollectReasonType = 8
	// AutoCollectDustUTXOType 自动归集Dust UTXO
	AutoCollectDustUTXOType CollectReasonType = 9
	// GiveOntOngFeeType 给erc20矿工费
	GiveOntOngFeeType CollectReasonType = 10
)
