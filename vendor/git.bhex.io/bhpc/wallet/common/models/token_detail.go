/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: token_collect.go
 * @Date: 2018/08/22
 * @Author: qiangjun.chen
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */
package models

import (
	"github.com/shopspring/decimal"
)

// TokenDetail define the table token_collect ‘归集’相关的配置
type TokenDetail struct {
	Base
	TokenID          TokenIDEnum      `json:"token_id" gorm:"not null;unique_index"`          // 币种symbol名称
	LogicType        TokenLogicType   `json:"logic_type" gorm:"not null;"`                    // token 对应的逻辑处理类型
	Type             TokenType        `json:"type" gorm:"not null;"`                          // 币种类型 1 utxo, 2 eth etc, 3 ltc, 4 eos
	GiveFeeThreshold decimal.Decimal  `json:"give_fee_threshold"  sql:"type:decimal(65,18);"` // 用户充了多少币(usdt, erc20)系统才会给母链币（BTC，ETH）作为fee。只有usdt、erc20需要，其他token不读
	GiveFeeValue     decimal.Decimal  `json:"give_fee_value"  sql:"type:decimal(65,18);"`     // 系统给多少（BTC，ETH）。只有usdt、erc20需要，其他token不读
	CollectThreshold decimal.Decimal  `json:"collect_threshold"  sql:"type:decimal(65,18);"`  // 达到多少会被收集
	CollectRemain    decimal.Decimal  `json:"collect_remain"  sql:"type:decimal(65,18);"`     // 保留多少不被收集，只有btc，eth才需要，其他token不读
	CollectInterval  uint32           `json:"collect_interval" gorm:"default 0;"`             // 自动归集时间间隔，单位小时数，0表示不自动归集
	CollectType      TokenCollectType `json:"collect_type" gorm:"default 0;"`                 // 归集类型
	Extension        string           `json:"extension"`                                      // 扩展字段
}

// TokenLogicType token type enum
type TokenLogicType uint32

const (
	// TokenLogicTypeUTXO btc like （ltc， bch，utxo）
	TokenLogicTypeUTXO TokenLogicType = 1
	// TokenLogicTypeETH eth，etc
	TokenLogicTypeETH TokenLogicType = 2
	// TokenLogicTypeERC20 erc20
	TokenLogicTypeERC20 TokenLogicType = 3
	// TokenLogicTypeEosLike eos
	TokenLogicTypeEosLike TokenLogicType = 4
	// TokenLogicTypeXrp XRP
	TokenLogicTypeXrp TokenLogicType = 5
	// TokenLogicTypeTron trx
	TokenLogicTypeTron TokenLogicType = 6
	// TokenLogicTypeBeam BEAM
	TokenLogicTypeBeam TokenLogicType = 7
	// TokenLogicTypeGrin GRIN
	TokenLogicTypeGrin TokenLogicType = 8
	//TokenLogicTypeTRC10 trc10 token
	TokenLogicTypeTRC10 TokenLogicType = 9
	//TokenLogicTypeTRC20 trc20 token
	TokenLogicTypeTRC20 TokenLogicType = 10
	TokenLogicTypeOnt   TokenLogicType = 11
	TokenLogicTypeNEO   TokenLogicType = 12
	TokenLogicTypeNAS   TokenLogicType = 13
	TokenLogicTypeCOBO  TokenLogicType = 14
	// TokenLogicTypeVOLUME vol mainnet token
	TokenLogicTypeVOLUME  TokenLogicType = 15
	TokenLogicTypeTD      TokenLogicType = 16
	TokenLogicTypeCOBOAPI TokenLogicType = 17
	// TokenLogicTypeBAASAPI baas api logic type
	TokenLogicTypeBAASAPI TokenLogicType = 18
	TokenLogicTypeWCG     TokenLogicType = 19
	TokenLogicTypeATOM    TokenLogicType = 20
	TokenLogicTypeBNB     TokenLogicType = 21
	TokenLogicTypeIOST    TokenLogicType = 22
	TokenLogicTypePOC     TokenLogicType = 23
	// TokenLogicTypeASCHLike offline sign fix fee and account type
	TokenLogicTypeASCHLike TokenLogicType = 24
	TokenLogicTypeXTZ      TokenLogicType = 25
	TokenLogicTypeIOTEX    TokenLogicType = 26
	TokenLogicTypeXRC20    TokenLogicType = 27
	TokenLogicTypeNULS     TokenLogicType = 28
	TokenLogicTypeNCC      TokenLogicType = 29 // nuls cross chain
	TokenLogicTypeNULS20   TokenLogicType = 30 // nuls nrc20
	TokenLogicTypeDOT      TokenLogicType = 31 // polka dot
	TokenLogicTypeSOL      TokenLogicType = 32 // solana
	// TokenLogicTypeCount 数量，用来校验，修改这个结构需要保证这条是在最后且最大
	TokenLogicTypeNEAR  TokenLogicType = 33 // solana
	TokenLogicTypeAVAX  TokenLogicType = 34 // avax
	TokenLogicTypeCount TokenLogicType = (TokenLogicTypeAVAX + 1)
)

// TokenType 不同(公链)的类型有不同的地址
type TokenType uint32

const (
	TokenTypeUtxo    TokenType = 1
	TokenTypeETH     TokenType = 2
	TokenTypeLTC     TokenType = 3
	TokenTypeEOS     TokenType = 4
	TokenTypeIPC     TokenType = 5
	TokenTypeTRON    TokenType = 6
	TokenTypeQtum    TokenType = 7
	TokenTypeBTM     TokenType = 8
	TokenTypeXRP     TokenType = 9
	TokenTypeBEAM    TokenType = 10
	TokenTypeGRIN    TokenType = 11
	TokenTypeDOGE    TokenType = 12
	TokenTypeVDS     TokenType = 13
	TokenTypeBHD     TokenType = 14
	TokenTypeONT     TokenType = 15
	TokenTypeALGO    TokenType = 16
	TokenTypeCXC     TokenType = 17
	TokenTypeTSC     TokenType = 18
	TokenTypeZEC     TokenType = 19
	TokenTypeNEO     TokenType = 20
	TokenTypeNAS     TokenType = 21
	TokenTypeCOBO    TokenType = 22
	TokenTypeSERO    TokenType = 23
	TokenTypeVOLUME  TokenType = 24
	TokenTypeCZZ     TokenType = 25
	TokenTypeVAS     TokenType = 26
	TokenTypeYTA     TokenType = 27
	TokenTypeXAG     TokenType = 28
	TokenTypeTD      TokenType = 29
	TokenTypeBR      TokenType = 30
	TokenTypeFN      TokenType = 31
	TokenTypePGN     TokenType = 32
	TokenTypeBCR     TokenType = 33
	TokenTypeFTO     TokenType = 34
	TokenTypeZOL     TokenType = 35
	TokenTypeGN      TokenType = 36
	TokenTypeCOBOAPI TokenType = 37
	TokenTypeWCG     TokenType = 38
	TokenTypeXMR     TokenType = 39
	TokenTypeDASH    TokenType = 40
	TokenTypeADA     TokenType = 41
	TokenTypeATOM    TokenType = 42
	TokenTypeBNB     TokenType = 43
	TokenTypeCKB     TokenType = 44
	TokenTypeIOST    TokenType = 45
	TokenTypePOC     TokenType = 46
	TokenTypeASCH    TokenType = 47
	TokenTypeXTZ     TokenType = 48
	TokenTypeIOTEX   TokenType = 49
	TokenTypeNULS    TokenType = 50
	TokenTypeFIL     TokenType = 51
	TokenTypeDOT     TokenType = 52
	TokenTypeSOL     TokenType = 53
	TokenTypeNEAR    TokenType = 54
	TokenTypeAVAX    TokenType = 55
	TokenTypeOASIS   TokenType = 56
)

// TokenCollectType token auto collect 的类型
type TokenCollectType uint32

const (
	UxtoLikeType  TokenCollectType = 1
	EthLikeType   TokenCollectType = 2
	Erc20LikeType TokenCollectType = 3
	UsdtLikeType  TokenCollectType = 4
	EosLikeType   TokenCollectType = 5
	XrpLikeType   TokenCollectType = 6
	BeamLikeType  TokenCollectType = 7
	GrinLikeType  TokenCollectType = 8
)
