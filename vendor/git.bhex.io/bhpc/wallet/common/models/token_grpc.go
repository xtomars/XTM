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

// TokenGrpc define the table token_grpc 币相关grpc的配置
type TokenGrpc struct {
	Base
	TokenID    TokenIDEnum `json:"token_id" gorm:"not null;unique_index"` // 币种symbol名称
	SignServer string      `json:"sign_server" gorm:"not null;"`          // 签名服务器
	RiskServer string      `json:"risk_server" gorm:"not null;"`          // 风控服务器
	ChainNode  string      `json:"chain_node"  sql:"not null;"`           // chainnode
	Extension  string      `json:"extension"`                             // 扩展字段
}
