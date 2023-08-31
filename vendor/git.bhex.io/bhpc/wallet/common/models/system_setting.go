/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: system_setting.go
 * @Date: 2018/06/21
 * @Author: zhiming.sun
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */

package models

// SystemSetting define the table system_setting
type SystemSetting struct {
	Base
	Key   string `json:"key" gorm:"type:varchar(50)" gorm:"unique_index"` // 系统总配置 key
	Value string `json:"value" gorm:"type:varchar(200)"`                  // 系统总配置 value
}
