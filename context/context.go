package dgctx

type DgContext struct {
	TraceId   string
	UserId    int64
	OpId      int64
	Role      string
	GroupId   int64
	Platform  string
	UserAgent string
	Lang      string
	GoId      uint64
	Token     string
	RemoteIp  string
	Extra     map[string]any
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
