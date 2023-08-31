package models

import "github.com/shopspring/decimal"

type ContributionRecord struct {
    Base
    PtContributionId string          `json:"pt_contribution_id" grom:"not null;type:varchar(100);unique_index"`
    OrgId            uint64          `json:"org_id" grom:"not null"`
    AccountId        uint64          `json:"account_id" grom:"not null"`
    PtAccountId      string          `json:"pt_account_id" grom:"not null;type:varchar(100)"`
    TokenID          TokenIDEnum     `json:"token_id" gorm:"not null"`
    Reference        string          `json:"reference" grom:"not null;type:varchar(100)"`
    ExpectedAmount   decimal.Decimal `json:"expected_amount"  sql:"type:decimal(65,18);not null"`
    Amount           decimal.Decimal `json:"amount" sql:"type:decimal(65,18);"`
    Status           string          `json:"status" grom:"not null;type:varchar(50)"`
}
