/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: token.go
 * @Date: 2018/06/21
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

import (
	"github.com/shopspring/decimal"
)

// Token define the table token
type Token struct {
	Base
	TokenID                 TokenIDEnum     `json:"token_id" gorm:"not null;unique_index"`                             // 币种symbol名称
	MinDepositAmount        decimal.Decimal `json:"min_deposit_amount"  sql:"type:decimal(65,18);" gorm:"default: 0;"` // 该币的最小充值金额,小于该数值的充值交易，不通知交易所
	CanTransactionConfirms  uint32          `json:"can_transaction_confirms" gorm:"not null;"`                         // 该Token可以交易的所需确认数
	CanWithdrawalConfirms   uint32          `json:"can_withdrawal_confirms" gorm:"not null;"`                          // 该Token可以交易的所需确认数
	EnableNotify            bool            `json:"enable_notify" gorm:"not null;"`                                    // 开启/关闭 到账通知开关
	EnableWithdrawal        bool            `json:"enable_withdrawal" gorm:"not null;"`                                // 开启/关闭 提现开关
	Status                  TokenStatus     `json:"status" gorm:"default: 0"`                                          // 币种状态：0 支持中 1 已上架  2 已下架
	ContractAddress         string          `json:"contract_address,omitempty" gorm:"type:varchar(60)"`                // 若该币种为ERC20代币，或者其他类型的代币，需要提供智能合约地址
	Decimals                uint32          `json:"decimals,omitempty"`                                                // 若该币种为ERC20代币，或者其他类型的代币，需要提供代币精度位数
	ValidBlockHeight        uint64          `json:"valid_block_height,omitempty"`                                      // 若该币种为ERC20代币，或者其他类型的代币，需要提供代币精度位数
	TotalSupply             decimal.Decimal `json:"total_supply,omitempty"  sql:"type:decimal(65,18);"`                // 若该币种为ERC20代币，或者其他类型的代币，需要提供代币总量
	IsNodeWallet            bool            `json:"is_node_wallet" gorm:"not null;"`                                   // 标志是否为节点钱包，用于特殊处理找零交易
	NeedMemo                bool            `json:"need_memo" gorm:"not null;"`                                        // 是否需要memo
	MaxNearWithdrawalCnt    uint32          `json:"max_near_withdrawal_cnt" `                                          // 单位时间内最大提现订单数
	MaxNearWithdrawalAmount decimal.Decimal `json:"max_near_withdrawal_amount"  sql:"type:decimal(65,18);"`            // 单位时间内最大提现金额
	NeedPreGenerateAddress  bool            `json:"need_pre_generate_address"`                                         // 是否需要预生成地址
	CanShareHotAddress      bool            `json:"can_share_hot_address"`                                             // 是否可以和其他币种共用地址
	Extension               string          `json:"extension"`                                                         // 扩展字段
}

// TokenStatus token status enum
type TokenStatus uint32

const (
	// Support 支持中
	Support TokenStatus = 0
	// Online 已上线
	Online TokenStatus = 1
	// Offline 已下线
	Offline TokenStatus = 2
)

// TokenIDEnum token id 的类型
type TokenIDEnum string

const (
	// TokenIDUSDT USDT 1
	TokenIDUSDT TokenIDEnum = "USDT"
	// TokenIDADA ADA 1
	TokenIDADA TokenIDEnum = "ADA"
	//TokenIDALGO  ALGO 1
	TokenIDALGO TokenIDEnum = "ALGO"
	//TokenIDBEAM  BEAM 1
	TokenIDBEAM TokenIDEnum = "BEAM"
	// TokenIDBTC BTC 1
	TokenIDBTC TokenIDEnum = "BTC"
	// TokenIDDOT DOT 1
	TokenIDDOT TokenIDEnum = "DOT"
	// TokenIDEOS EOS 1
	TokenIDEOS TokenIDEnum = "EOS"
	// TokenIDETH eth 1
	TokenIDETH TokenIDEnum = "ETH"
	//TokenIDGRIN  GRIN 1
	TokenIDGRIN TokenIDEnum = "GRIN"
	// TokenIDIOST IOST => use github/official directly
	TokenIDIOST TokenIDEnum = "IOST"
	//TokenIDONG  ONG 1
	TokenIDONG TokenIDEnum = "ONG"
	//TokenIDSERO SERO 1
	TokenIDSERO TokenIDEnum = "SERO"
	// TokenIDTRON TRX 1
	TokenIDTRON TokenIDEnum = "TRX"
	//TokenIDXRP XRP => ripple
	TokenIDXRP TokenIDEnum = "XRP"
)
