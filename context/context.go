package dgctx

import "sync"

var initUnsafeExtraOnce sync.Once

type DgContext struct {
	TraceId     string `json:"traceId,omitempty"`
	UserId      int64  `json:"userId,omitempty"`
	OpId        int64  `json:"opId,omitempty"`
	RunAs       int64  `json:"runAs,omitempty"`
	Roles       string `json:"roles,omitempty"`
	BizTypes    int    `json:"bizTypes,omitempty"`
	GroupId     int64  `json:"groupId,omitempty"`
	Platform    string `json:"platform,omitempty"`
	UserAgent   string `json:"userAgent,omitempty"`
	Lang        string `json:"lang,omitempty"`
	GoId        uint64 `json:"goId,omitempty"`
	Token       string `json:"token,omitempty"`
	ShareToken  string `json:"shareToken,omitempty"`
	RemoteIp    string `json:"remoteIp,omitempty"`
	CompanyId   int64  `json:"companyId,omitempty"`
	Product     int    `json:"product,omitempty"`
	safeExtra   sync.Map
	unsafeExtra map[string]any
}

func (ctx *DgContext) SetExtraKeyValue(key string, val any) {
	ctx.safeExtra.Store(key, val)
}

func (ctx *DgContext) GetExtraValue(key string) any {
	val, _ := ctx.safeExtra.Load(key)
	return val
}

func (ctx *DgContext) SetUnsafeExtraKeyValue(key string, val any) {
	initUnsafeExtraOnce.Do(func() {
		ctx.unsafeExtra = make(map[string]any)
	})
	ctx.unsafeExtra[key] = val
}

func (ctx *DgContext) GetUnsafeExtraValue(key string) any {
	if ctx.unsafeExtra == nil {
		return nil
	}

	return ctx.unsafeExtra[key]
}
