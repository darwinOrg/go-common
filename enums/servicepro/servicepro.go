package servicepro

import (
	"errors"
	"fmt"
)

type ServiceProduct struct {
	Code     int64
	Name     string
	SchemaId int64
}

var (
	EOR = &ServiceProduct{
		Code:     1,
		Name:     "海外雇佣",
		SchemaId: 44,
	}
	CONTRACTOR = &ServiceProduct{
		Code:     2,
		Name:     "灵活用工",
		SchemaId: 0,
	}
	PAYROLL = &ServiceProduct{
		Code:     3,
		Name:     "薪酬服务",
		SchemaId: 49,
	}
	HEAD_HUNTER = &ServiceProduct{
		Code:     4,
		Name:     "全球猎头",
		SchemaId: 47,
	}
	CONSULT = &ServiceProduct{
		Code:     5,
		Name:     "合规咨询",
		SchemaId: 46,
	}
	COMPANY_LANDING = &ServiceProduct{
		Code:     6,
		Name:     "企业落地",
		SchemaId: 45,
	}
	COUNTRY_SPEC = &ServiceProduct{
		Code:     7,
		Name:     "国别服务",
		SchemaId: 0,
	}
)

var allProductMap = map[string]*ServiceProduct{
	"EOR":             EOR,
	"CONTRACTOR":      CONTRACTOR,
	"PAYROLL":         PAYROLL,
	"HEAD_HUNTER":     HEAD_HUNTER,
	"CONSULT":         CONSULT,
	"COMPANY_LANDING": COMPANY_LANDING,
	"COUNTRY_SPEC":    COUNTRY_SPEC,
}

func GetByCode(code string) (*ServiceProduct, error) {
	v, ok := allProductMap[code]
	if !ok {
		return nil, errors.New(fmt.Sprintf("illegal argument serviceproduct code: %s", code))
	}
	return v, nil
}
