/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: collect_rule.go
 * @Date: 2018/08/27
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

// CollectRule define the table collect_rule
type CollectRule struct {
	Base
	TokenID   string     `json:"token_id"`  // 该地址绑定的币种
	TacticsID uint64     `json:"tatics_id"` // 策略ID
	Status    RuleStatus `json:"status"`    // 状态，（0未启用，1启用）
}

// RuleStatus rulle status enum
type RuleStatus uint32

const (
	// Disable 未启用
	Disable RuleStatus = 0
	// Enable 启用
	Enable RuleStatus = 1
)
