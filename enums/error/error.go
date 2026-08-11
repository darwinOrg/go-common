package dgerr

import "fmt"

type DgError struct {
	Code    int
	Message string
	Cause   error
}

func (e *DgError) WithCause(err error) *DgError {
	if e == nil {
		return nil
	}

	return &DgError{Code: e.Code, Message: e.Message, Cause: err}
}

func (e *DgError) Error() string {
	if e == nil {
		return ""
	}

	if e.Cause == nil {
		return e.Message
	}

	return fmt.Sprintf("%s - %v", e.Message, e.Cause)
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

func SimpleDgErrorWithCause(msg string, err error) *DgError {
	return &DgError{
		Code:    -1,
		Message: msg,
		Cause:   err,
	}
}

var (
	SYSTEM_ERROR     = &DgError{Code: 5001, Message: "系统错误"}
	SYSTEM_BUSY      = &DgError{Code: 5002, Message: "系统繁忙"}
	TIME_OUT         = &DgError{Code: 5003, Message: "请求超时"}
	TOO_MANY_REQUEST = &DgError{Code: 5004, Message: "请求太频繁"}

	ARGUMENT_NOT_VALID     = &DgError{Code: 4001, Message: "无效参数"}
	INVALID_TOKEN          = &DgError{Code: 4002, Message: "无效token"}
	UPLOAD_FILE_SIZE_LIMIT = &DgError{Code: 4003, Message: "文件过大"}
	DUPLICATE_PRIMARY_KEY  = &DgError{Code: 4004, Message: "重复主键"}

	LOGIN_ERROR       = &DgError{Code: 4005, Message: "登录错误"}
	NOT_LOGIN_IN      = &DgError{Code: 4006, Message: "用户未登录"}
	USER_NOT_EXISTS   = &DgError{Code: 4007, Message: "用户不存在"}
	WRONG_PASSWORD    = &DgError{Code: 4008, Message: "密码错误"}
	NO_PERMISSION     = &DgError{Code: 4009, Message: "无权限"}
	ILLEGAL_OPERATION = &DgError{Code: 4010, Message: "非法操作"}

	RECORD_EXISTS     = &DgError{Code: 4013, Message: "记录已存在"}
	RECORD_NOT_EXISTS = &DgError{Code: 4014, Message: "记录不存在"}
	DISABLED_USER     = &DgError{Code: 4015, Message: "用户已被禁用"}
	EMAIL_REGISTERED  = &DgError{Code: 4016, Message: "邮箱已被注册"}
)
