package roles

import (
	"encoding/json"
	"github.com/darwinOrg/go-common/enums/domains"
	"golang.org/x/exp/slices"
	"strings"
)

type Role struct {
	Id      int              `json:"id,omitempty"`
	Name    string           `json:"name,omitempty"`
	Desc    string           `json:"desc,omitempty"`
	Domains []domains.Domain `json:"domains,omitempty"`
}

func (r *Role) String() string {
	if r == nil {
		return ""
	} else {
		return r.Name
	}
}

func (r *Role) MarshalJSON() ([]byte, error) {
	if r == nil {
		return nil, nil
	} else {
		return json.Marshal(r.Name)
	}
}

func (r *Role) Marshal() (string, error) {
	if r == nil {
		return "", nil
	} else {
		return r.Name, nil
	}
}

var (
	Manger       = role(1, "MANAGER", "管理员", []domains.Domain{domains.Supplier, domains.Platform})
	Staff        = role(2, "STAFF", "员工", []domains.Domain{domains.Platform})
	Salesman     = role(3, "SALESMAN", "业务员", []domains.Domain{domains.Supplier})
	SuperManager = role(4, "SUPER_MANAGER", "超管", []domains.Domain{domains.Supplier, domains.Business})
	GroupManager = role(5, "GROUP_MANAGER", "组管理员", []domains.Domain{domains.Business})
	Finance      = role(6, "FINANCE", "财务", []domains.Domain{domains.Supplier, domains.Business})
	HR           = role(7, "HR", "HR", []domains.Domain{domains.Business})
)

func role(id int, name string, desc string, domains []domains.Domain) *Role {
	return &Role{
		Id:      id,
		Name:    name,
		Desc:    desc,
		Domains: domains,
	}
}

var (
	all = []*Role{
		Manger, Staff, Salesman, SuperManager, GroupManager, Finance, HR,
	}
	idRoleMap   = initIdRoleMap()
	nameRoleMap = initNameRoleMap()
)

func All() []*Role {
	return all
}

func initIdRoleMap() map[int]*Role {
	m := make(map[int]*Role)
	for _, r := range all {
		m[r.Id] = r
	}
	return m
}

func initNameRoleMap() map[string]*Role {
	m := make(map[string]*Role)
	for _, r := range all {
		m[r.Name] = r
	}
	return m
}

func OfId(id int) *Role {
	return idRoleMap[id]
}

func OfName(name string) *Role {
	return nameRoleMap[strings.ToUpper(name)]
}

func IsAny(this *Role, any ...*Role) bool {
	if len(any) == 0 {
		return false
	}
	for _, e := range any {
		if this == e {
			return true
		}
	}
	return false
}

func (r *Role) ContainsIn(domain domains.Domain) bool {
	if r == nil {
		return false
	}
	if len(domain) == 0 {
		return false
	}
	return slices.Contains(r.Domains, domain)
}

func (r *Role) NotContainsIn(domain domains.Domain) bool {
	return !r.ContainsIn(domain)
}

func IsNotSuperManager(r *Role) bool {
	return !IsSuperManager(r)
}
func IsSuperManager(r *Role) bool {
	if r == nil {
		return false
	}
	if r.Id == SuperManager.Id {
		return true
	}
	return false
}
