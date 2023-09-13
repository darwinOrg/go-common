package feeitem

import (
	"errors"
	"fmt"
)

type FeeItem struct {
	Code string
	Desc string
}

var (
	COST                       = &FeeItem{Code: "COST", Desc: "成本"}
	SERVICE_FEE                = &FeeItem{Code: "SERVICE_FEE", Desc: "服务费"}
	SETUP_FEE                  = &FeeItem{Code: "SETUP_FEE", Desc: "一次性配置费"}
	TRANSFER_FEE               = &FeeItem{Code: "TRANSFER_FEE", Desc: "转账费"}
	SALARY_DEPOSIT             = &FeeItem{Code: "SALARY_DEPOSIT", Desc: "薪资押金"}
	SERVICE_FEE_DEPOSIT        = &FeeItem{Code: "SERVICE_FEE_DEPOSIT", Desc: "服务费押金"}
	FIRST_MONTH_SALARY_IMPREST = &FeeItem{Code: "FIRST_MONTH_SALARY_IMPREST", Desc: "首月薪资预付款"}
	TOTAL_SALARY               = &FeeItem{Code: "TOTAL_SALARY", Desc: "工资总额"}
	VAT                        = &FeeItem{Code: "VAT", Desc: "增值税"}
)

var allFeeItemMap = map[string]*FeeItem{
	"COST":                       COST,
	"SERVICE_FEE":                SERVICE_FEE,
	"SETUP_FEE":                  SETUP_FEE,
	"TRANSFER_FEE":               TRANSFER_FEE,
	"SALARY_DEPOSIT":             SALARY_DEPOSIT,
	"SERVICE_FEE_DEPOSIT":        SERVICE_FEE_DEPOSIT,
	"FIRST_MONTH_SALARY_IMPREST": FIRST_MONTH_SALARY_IMPREST,
	"TOTAL_SALARY":               TOTAL_SALARY,
	"VAT":                        VAT,
}

func GetByCode(code string) (*FeeItem, error) {
	v, ok := allFeeItemMap[code]
	if !ok {
		return nil, errors.New(fmt.Sprintf("illegal argument feeitem code: %s", code))
	}
	return v, nil
}
