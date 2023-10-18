package dgctx

type DgContext struct {
	TraceId    string
	UserId     int64
	OpId       int64
	RunAs      int64
	Roles      string
	BizTypes   int
	GroupId    int64
	Platform   string
	UserAgent  string
	Lang       string
	GoId       uint64
	Token      string
	ShareToken string
	RemoteIp   string
	CompanyId  int64
	Product    int
	Extra      map[string]any
}

func (ctx *DgContext) SetExtraKeyValue(key string, val any) {
	if ctx.Extra == nil {
		ctx.Extra = map[string]any{key: val}
	} else {
		ctx.Extra[key] = val
	}
}

func (ctx *DgContext) GetExtraValue(key string) any {
	if ctx.Extra == nil {
		return nil
	} else {
		return ctx.Extra[key]
	}
}
