/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: change_contract_record.go
 * @Date: 2019/07/31
 * @Author: zhiming.sun
 * @Copyright（C）: 2019 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

import (
	"github.com/shopspring/decimal"
)

// ChangeContractRecord define the table change_contract_record
type ChangeContractRecord struct {
	Base
	TokenID     TokenIDEnum     `json:"token_id" gorm:"not null;unique_index:idx_depositorder_txhash_idx_tokenid"` // 币种
	OldContract string          `json:"old_contract" gorm:"type:varchar(255);not null"`                            // 旧合约地址
	NewContract string          `json:"new_contract" gorm:"type:varchar(255);not null"`
	OldBalance  decimal.Decimal `json:"old_balance"  sql:"type:decimal(65,18);not null"` // 旧合约资产余额
	NewBalance  decimal.Decimal `json:"new_balance"  sql:"type:decimal(65,18);not null"` // 新合约资产余额
	BlockHeight int64           `json:"block_height" gorm:"default: 0"`                  // 有效地址
	Memo        string          `json:"memo"`
	Extension   string          `json:"extension"` // 扩展字段
}
