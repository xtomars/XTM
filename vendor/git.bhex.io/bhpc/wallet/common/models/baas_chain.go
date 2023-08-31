/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: baas_chain.go
 * @Date: 2020/02/06
 * @Author: Jia Liu
 * @Copyright（C）: 2020 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

// BaasChain bass chain
type BaasChain struct {
	Base
	APIKey          string    `json:"api_key" gorm:"not null;type:varchar(1024);"`                                     // api_key 是接入服务的身份识别
	Ed25519PubicKey string    `json:"ed25519_public_key" gorm:"not null;type:varchar(1024);column:ed25519_publickey;"` // api_key 是接入服务的验签公钥
	Chain           string    `json:"chain" gorm:"not null;type:varchar(255);unique_index"`                            // api_key 是接入服务的验签公钥
	ChainType       TokenType `json:"chain_type" gorm:"not null;"`                                                     // chain_type是表明chain类型用来生成地址用的
	AddressRegexp   string    `json:"address_regexp" gorm:"type:varchar(255)"`                                         // address_regexp 用来校验提现地址的正则表达式
	MemoRegexp      string    `json:"memo_regexp" gorm:"type:varchar(255)"`                                            // memo_regexp 用来校验提现地址memo的正则表达式
}
