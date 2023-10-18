package dgerr

type DgError struct {
	Code    int
	Message string
}

func (e *DgError) Error() string {
	return e.Message
}

func NewDgError(code int, msg string) *DgError {
	return &DgError{
		Code:    code,
		Message: msg,
	}
}

func SimpleDgError(msg string) *DgError {
	return &DgError{
		Code:    -1,
		Message: msg,
	}
}

var (
	SYSTEM_ERROR = &DgError{5001, "系统错误"}
	SYSTEM_BUSY  = &DgError{5002, "系统繁忙"}
	TIME_OUT     = &DgError{5003, "请求超时"}

	ARGUMENT_NOT_VALID     = &DgError{4001, "无效参数"}
	INVALID_TOKEN          = &DgError{4002, "无效token"}
	UPLOAD_FILE_SIZE_LIMIT = &DgError{4003, "文件过大"}
	DUPLICATE_PRIMARY_KEY  = &DgError{4004, "重复主键"}

	LOGIN_ERROR       = &DgError{4005, "登录错误"}
	NOT_LOGIN_IN      = &DgError{4006, "用户未登录"}
	USER_NOT_EXISTS   = &DgError{4007, "用户不存在"}
	WRONG_PASSWORD    = &DgError{4008, "密码错误"}
	NO_PERMISSION     = &DgError{4009, "无权限"}
	ILLEGAL_OPERATION = &DgError{4010, "非法操作"}

	RECORD_NOT_EXISTS     = &DgError{4014, "记录不存在"}
	DISABLED_USER         = &DgError{4015, "用户已被禁用"}
	EMAIL_REGISTERED      = &DgError{4016, "邮箱已被注册"}
	CUSTOMER_HAS_NO_GROUP = &DgError{4031, "该企业暂无国家对应团队"}
)
