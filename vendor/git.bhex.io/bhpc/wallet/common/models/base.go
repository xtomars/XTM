/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: base.go
 * @Date: 2018/06/21
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

import (
	"time"
)

// Base define the common fields in table
type Base struct {
	ID        uint64     `json:"id" gorm:"primary_key"`                      // 记录ID
	Remark    string     `json:"remark,omitempty" gorm:"type:varchar(100);"` // 备注字段
	CreatedAt time.Time  `json:"created_at,omitempty"`                       // 记录创建时间
	UpdatedAt time.Time  `json:"updated_at,omitempty"`                       // 记录被更新时间
}
