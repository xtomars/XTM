/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: miner_fee.go
 * @Date: 2018/08/05
 * @Author: qiangjun.chen
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */
package models

import "github.com/shopspring/decimal"

// MinerFee define the table miner_fee
type MinerFee struct {
	Base
	TokenID        TokenIDEnum     `json:"token_id" gorm:"not null;unique_index"`                        // 币种
	FeeTokenID     TokenIDEnum     `json:"fee_token_id" gorm:"not null;"`                                // 消耗币种
	Min            decimal.Decimal `json:"min"  sql:"type:decimal(65,18);" gorm:"not null;"`             // 最小
	Max            decimal.Decimal `json:"max"  sql:"type:decimal(65,18);" gorm:"not null;"`             // 最大
	Suggest        decimal.Decimal `json:"suggest"  sql:"type:decimal(65,18);" gorm:"not null;"`         // 推荐
	MaxAutocollect decimal.Decimal `json:"max_autocollect"  sql:"type:decimal(65,18);" gorm:"not null;"` // 归集的时候最大矿工费
	GasLimit       uint64          `json:"gas_limit" gorm:"default: 0"`                                  // gas限制
	Extension      string          `json:"extension"`                                                    // 扩展字段
	AutoUpdateFee  bool            `json:"auto_update_fee"`                                              // 是否自动更新费率
	UseZeroFee     bool            `json:"use_zero_fee"`                                                 // 是否返回0费率
}
