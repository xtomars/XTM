/*
 * *******************************************************************
 * @项目名称: common
 * @文件名称: Certificate.go
 * @Date: 2018/10/14
 * @Author: chunhua.guo
 * @Copyright（C）: 2018 BlueHelix Inc.   All rights reserved.
 * 注意：本内容仅限于内部传阅，禁止外泄以及用于其他的商业目的.
 * *******************************************************************
 */
package models

// Asset define the table asset
type Certificate struct {
	Base
	Name    string          `json:"name" gorm:"not null;type:varchar(60);unique_index:idx_name"` // 证书名
	Type    CertificateType `json:"type" gorm:"not null"`                                        // 类型
	Version int             `json:"version" gorm:"not null"`                                     // 版本
	Content string          `json:"content" gorm:"not null;type:varchar(1024)"`                  // 内容
}

// CertificateType certificateType type enum
type CertificateType uint32

const (
	//private certificate
	PriCert CertificateType = 1
	//public certificate
	PubCert CertificateType = 2
)
