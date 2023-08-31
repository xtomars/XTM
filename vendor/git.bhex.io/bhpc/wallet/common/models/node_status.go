/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: node_status.go
 * @Date: 2018/06/21
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

import "github.com/shopspring/decimal"

// NodeStatus define the table node_status
type NodeStatus struct {
	Base
	Name        string          `json:"name" gorm:"type:varchar(20);"`   // 该节点名字
	IsOnline    bool            `json:"is_online"`                       // 是否正在运行
	CPUUse      float64         `json:"cpu_use"`                         // 该节点CPU使用率
	MemUse      float64         `json:"mem_use"`                         // 该节点内存使用率
	DiskUse     float64         `json:"disk_use"`                        // 该节点硬盘使用率
	Ping        uint32          `json:"ping"`                            // 该节点对外部节点的ping值
	TTL         uint32          `json:"ttl"`                             // 该节点对外部节点的延时
	BlockHeight int64           `json:"block_height"`                    // 该节点的当前区块高度
	Fee         decimal.Decimal `json:"fee"  sql:"type:decimal(65,18);"` // 最近的网络交易费用
}
