package serviceitem

import (
	"errors"
	"fmt"
)

const (
	// EOR_EMPLOYEE_ENTRY_TYPE EOR_EMPLOYEE_ENTRY 海外雇佣(员工入职)
	EOR_EMPLOYEE_ENTRY_TYPE = "EOR_EMPLOYEE_ENTRY"
	// EOR_WORK_PERMIT_TYPE EOR_WORK_PERMIT 海外雇佣(工作签)
	EOR_WORK_PERMIT_TYPE = "EOR_WORK_PERMIT"
	// EOR_ID06_TYPE EOR_ID06 海外雇佣(ID06)
	EOR_ID06_TYPE = "EOR_ID06"
	// CONTRACTOR_TYPE CONTRACTOR 灵活用工
	CONTRACTOR_TYPE = "CONTRACTOR"
	// PAYROLL_TYPE PAYROLL 薪酬服务
	PAYROLL_TYPE = "PAYROLL"
	// HEAD_HUNTER_TYPE HEAD_HUNTER 全球猎头
	HEAD_HUNTER_TYPE = "HEAD_HUNTER"
	// CONSULT_TYPE CONSULT 合规咨询
	CONSULT_TYPE = "CONSULT"
	// COMPANY_LANDING_TYPE COMPANY_LANDING 企业落地
	COMPANY_LANDING_TYPE = "COMPANY_LANDING"
	// COUNTRY_SPEC_WORK_PERMIT_TYPE 国别服务工作签
	COUNTRY_SPEC_WORK_PERMIT_TYPE = "COUNTRY_SPEC_WORK_PERMIT"
	// COUNTRY_SPEC_PERSONNEL_TAX_TYPE 国别服务个税
	COUNTRY_SPEC_PERSONNEL_TAX_TYPE = "COUNTRY_SPEC_PERSONNEL_TAX"
	// COUNTRY_SPEC_ID06_TYPE 国别服务ID06
	COUNTRY_SPEC_ID06_TYPE = "COUNTRY_SPEC_ID06"
)

type ServiceItemType struct {
	ServiceProductId int64
	Name             string
	Type             string
	DeliverType      int
	SchemaId         int64
}

var (
	EOR_EMPLOYEE_ENTRY = &ServiceItemType{
		ServiceProductId: 1,
		Name:             "海外雇佣(员工入职)",
		Type:             EOR_EMPLOYEE_ENTRY_TYPE,
		DeliverType:      1,
		SchemaId:         0,
	}
	EOR_WORK_PERMIT = &ServiceItemType{
		ServiceProductId: 1,
		Name:             "海外雇佣(工作签)",
		Type:             EOR_WORK_PERMIT_TYPE,
		DeliverType:      4,
		SchemaId:         0,
	}
	EOR_ID06 = &ServiceItemType{
		ServiceProductId: 1,
		Name:             "海外雇佣(ID06)",
		Type:             EOR_ID06_TYPE,
		DeliverType:      3,
		SchemaId:         0,
	}
	CONTRACTOR = &ServiceItemType{
		ServiceProductId: 2,
		Name:             "灵活用工",
		Type:             CONTRACTOR_TYPE,
		DeliverType:      0,
		SchemaId:         0,
	}
	PAYROLL = &ServiceItemType{
		ServiceProductId: 3,
		Name:             "薪酬服务",
		Type:             PAYROLL_TYPE,
		DeliverType:      2,
		SchemaId:         0,
	}
	HEAD_HUNTER = &ServiceItemType{
		ServiceProductId: 4,
		Name:             "全球猎头",
		Type:             HEAD_HUNTER_TYPE,
		DeliverType:      0,
		SchemaId:         0,
	}
	CONSULT = &ServiceItemType{
		ServiceProductId: 5,
		Name:             "合规咨询",
		Type:             CONSULT_TYPE,
		DeliverType:      0,
		SchemaId:         0,
	}
	COMPANY_LANDING = &ServiceItemType{
		ServiceProductId: 6,
		Name:             "企业落地",
		Type:             COMPANY_LANDING_TYPE,
		DeliverType:      0,
		SchemaId:         0,
	}
	COUNTRY_SPEC_WORK_PERMIT = &ServiceItemType{
		ServiceProductId: 7,
		Name:             "国别服务(工作签)",
		Type:             COUNTRY_SPEC_WORK_PERMIT_TYPE,
		DeliverType:      4,
		SchemaId:         51,
	}
	COUNTRY_SPEC_PERSONNEL_TAX = &ServiceItemType{
		ServiceProductId: 7,
		Name:             "国别服务(个人税务登记)",
		Type:             COUNTRY_SPEC_PERSONNEL_TAX_TYPE,
		DeliverType:      0,
		SchemaId:         50,
	}
	COUNTRY_SPEC_ID06 = &ServiceItemType{
		ServiceProductId: 7,
		Name:             "国别服务(ID06)",
		Type:             COUNTRY_SPEC_ID06_TYPE,
		DeliverType:      3,
		SchemaId:         48,
	}
)

var alItemTypeMap = map[string]*ServiceItemType{
	EOR_EMPLOYEE_ENTRY_TYPE:         EOR_EMPLOYEE_ENTRY,
	EOR_WORK_PERMIT_TYPE:            EOR_WORK_PERMIT,
	EOR_ID06_TYPE:                   EOR_ID06,
	CONTRACTOR_TYPE:                 CONTRACTOR,
	PAYROLL_TYPE:                    PAYROLL,
	HEAD_HUNTER_TYPE:                HEAD_HUNTER,
	CONSULT_TYPE:                    CONSULT,
	COMPANY_LANDING_TYPE:            COMPANY_LANDING,
	COUNTRY_SPEC_WORK_PERMIT_TYPE:   COUNTRY_SPEC_WORK_PERMIT,
	COUNTRY_SPEC_PERSONNEL_TAX_TYPE: COUNTRY_SPEC_PERSONNEL_TAX,
	COUNTRY_SPEC_ID06_TYPE:          COUNTRY_SPEC_ID06,
}

func GetByType(typ string) (*ServiceItemType, error) {
	v, ok := alItemTypeMap[typ]
	if !ok {
		return nil, errors.New(fmt.Sprintf("illegal argument serviceitem type: %s", typ))
	}
	return v, nil
}
