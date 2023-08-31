/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: address.go
 * @Date: 2018/06/22
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

// Address define the table address
type Address struct {
	Base
	TokenType      TokenType     `json:"token_type" gorm:"not null;"`                            // 地址所属token类型
	Address        string        `json:"address" gorm:"not null;type:varchar(255);unique_index"` // 地址
	Type           AddressType   `json:"type" gorm:"not null;"`                                  // 地址类型,1 单签名地址 2 多签名地址 3 用于合成多签名的地址
	Network        NetWork       `json:"network"`                                                // 主网、测试网啥的
	MultiAddresses string        `json:"multi_addresses,omitempty" gorm:"type:varchar(512);"`    // 多签名关联地址,值（required|addr1|addr2|addr3）
	Owner          OwnerType     `json:"owner" gorm:"not null;"`                                 // 钱包类型（1 个人 2 热钱包 3 冷钱包）
	Status         AddressStatus `json:"status" gorm:"default:0"`                                // 状态，（0未分配，1已分配，2废弃）
	Path           string        `json:"path" gorm:"type:varchar(50);"`                          // 生成的地址路径
	Version        uint32        `json:"version"`                                                // 生成的批量地址批次号
	RedeemScript   []byte        `json:"redeem_script,omitempty"`                                // 多签名提取脚本
	Memo           string        `json:"memo"`
	Extension      string        `json:"extension"` // 扩展字段
}

type AddressRc struct {
	ID             uint64 `json:"id" gorm:"primary_key"`                                  // 记录ID
	Address        string `json:"address" gorm:"not null;type:varchar(255);unique_index"` // 地址
	TokenType      uint32 `json:"token_type" gorm:"not null;"`                            // 地址所属token类型
	Type           int64  `json:"type" gorm:"not null;"`                                  // 地址类型,1 单签名地址 2 多签名地址 3 用于合成多签名的地址
	Network        string `json:"network"`                                                // 主网、测试网啥的
	MultiAddresses string `json:"multi_addresses,omitempty" gorm:"type:varchar(512);"`    // 多签名关联地址,值（required|addr1|addr2|addr3）
	Owner          uint32 `json:"owner" gorm:"not null;"`                                 // 钱包类型（1 个人 2 热钱包 3 冷钱包）
	Path           string `json:"path" gorm:"type:varchar(50);"`                          // 生成的地址路径
	Version        uint32 `json:"version"`                                                // 生成的批量地址批次号
	RedeemScript   []byte `json:"redeem_script,omitempty"`                                // 多签名提取脚本
	CheckSign      string `json:"check_sign"`                                             // 地址正确性的签名
}

// AddressType address type enum
type AddressType int64

const (
	// SingleAddress 单签名地址
	SingleAddress AddressType = 1
	// MultiAddress 多签名地址
	MultiAddress AddressType = 2
	// ForMultiAddress 用来生成多签名地址的单签名地址
	ForMultiAddress AddressType = 3
)

// AddressStatus address status enum
type AddressStatus int64

const (
	// UnUsed 未分配
	UnUsed AddressStatus = 0
	// Used 已分配
	Used AddressStatus = 1
	// Discard 已废弃
	Discard AddressStatus = 2
)

// OwnerType address owner tyoe enum
type OwnerType uint32

const (
	// User 个人
	User OwnerType = 1
	// HotWallet 热钱包
	HotWallet OwnerType = 2
	// ColdWallet 冷钱包
	ColdWallet OwnerType = 3
)

// NetWork internal define
type NetWork string

const (
	// MainNet 主网
	MainNet NetWork = "MAINNET"
	// TestNet 测试网络
	TestNet NetWork = "TESTNET"
)
