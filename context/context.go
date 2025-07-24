package dgctx

import (
	"context"
	"github.com/darwinOrg/go-common/utils"
	"sync"
	"time"
)

type DgContext struct {
	TraceId       string  `json:"traceId,omitempty"`
	UserId        int64   `json:"userId,omitempty"`
	OpId          int64   `json:"opId,omitempty"`
	RunAs         int64   `json:"runAs,omitempty"`
	Roles         string  `json:"roles,omitempty"`
	BizTypes      int     `json:"bizTypes,omitempty"`
	GroupId       int64   `json:"groupId,omitempty"`
	Platform      string  `json:"platform,omitempty"`
	UserAgent     string  `json:"userAgent,omitempty"`
	Lang          string  `json:"lang,omitempty"`
	GoId          uint64  `json:"goId,omitempty"`
	Token         string  `json:"token,omitempty"`
	ShareToken    string  `json:"shareToken,omitempty"`
	RemoteIp      string  `json:"remoteIp,omitempty"`
	CompanyId     int64   `json:"companyId,omitempty"`
	Product       int     `json:"product,omitempty"`
	Products      []int   `json:"products,omitempty"`
	DepartmentIds []int64 `json:"departmentIds,omitempty"`
	NotPrintLog   bool    `json:"-"`
	NotLogSQL     bool    `json:"-"`
	superRight    bool
	inner         context.Context
	safeExtra     sync.Map
	unsafeExtra   map[string]any
}

func SimpleDgContext() *DgContext {
	traceId, _ := utils.RandomLetter(32)
	return &DgContext{TraceId: traceId}
}

func WithTimeout(parent context.Context, timeout time.Duration) (*DgContext, context.CancelFunc) {
	ctxWT, cancel := context.WithTimeout(parent, timeout)
	ctx := SimpleDgContext()
	ctx.inner = ctxWT

	return ctx, cancel
}

func WithCancel(parent context.Context) (*DgContext, context.CancelFunc) {
	ctxWT, cancel := context.WithCancel(parent)
	ctx := SimpleDgContext()
	ctx.inner = ctxWT

	return ctx, cancel
}

func WithValue(parent context.Context, key, val any) *DgContext {
	ctxWV := context.WithValue(parent, key, val)
	ctx := SimpleDgContext()
	ctx.inner = ctxWV

	return ctx
}

func (ctx *DgContext) WithTimeout(parent context.Context, timeout time.Duration) context.CancelFunc {
	ctxWT, cancel := context.WithTimeout(parent, timeout)
	ctx.inner = ctxWT

	return cancel
}

func (ctx *DgContext) WithCancel(parent context.Context) context.CancelFunc {
	ctxWT, cancel := context.WithCancel(parent)
	ctx.inner = ctxWT

	return cancel
}

func (ctx *DgContext) WithValue(parent context.Context, key, val any) {
	ctxWV := context.WithValue(parent, key, val)
	ctx.inner = ctxWV
}

func (ctx *DgContext) SetInnerContext(inner context.Context) {
	ctx.inner = inner
}

func (ctx *DgContext) GetInnerContext() context.Context {
	return ctx.inner
}

func (ctx *DgContext) SetExtraKeyValue(key string, val any) {
	ctx.safeExtra.Store(key, val)
}

func (ctx *DgContext) GetExtraValue(key string) any {
	val, _ := ctx.safeExtra.Load(key)
	return val
}

func (ctx *DgContext) SetUnsafeExtraKeyValue(key string, val any) {
	if ctx.unsafeExtra == nil {
		ctx.unsafeExtra = make(map[string]any)
	}

	ctx.unsafeExtra[key] = val
}

func (ctx *DgContext) GetUnsafeExtraValue(key string) any {
	if ctx.unsafeExtra == nil {
		return nil
	}

	return ctx.unsafeExtra[key]
}

func (ctx *DgContext) SetSuperRight(superRight bool) {
	ctx.superRight = superRight
}

func (ctx *DgContext) HasSuperRight() bool {
	return ctx.superRight
}

func (ctx *DgContext) Clone() *DgContext {
	clone := utils.MustConvertToNewBeanByJson[DgContext](ctx)
	clone.superRight = ctx.superRight
	clone.unsafeExtra = ctx.unsafeExtra
	ctx.safeExtra.Range(func(key, val any) bool {
		clone.safeExtra.Store(key, val)
		return true
	})

	return clone
}
