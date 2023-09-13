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
	SYSTEM_ERROR = &DgError{5001, "System error"}
	SYSTEM_BUSY  = &DgError{5002, "System invalid"}

	ARGUMENT_NOT_VALID     = &DgError{4001, "Invalid arguments"}
	INVALID_TOKEN          = &DgError{4002, "Invalid token"}
	UPLOAD_FILE_SIZE_LIMIT = &DgError{4003, "File size overflow"}
	DUPLICATE_PRIMARY_KEY  = &DgError{4004, "Multiple record"}

	LOGIN_ERROR       = &DgError{4005, "Login error"}
	NOT_LOGIN_IN      = &DgError{4006, "User not login"}
	USER_NOT_EXISTS   = &DgError{4007, "User not exist"}
	WRONG_PASSWORD    = &DgError{4008, "Wrong password"}
	NO_PERMISSION     = &DgError{4009, "No permission"}
	ILLEGAL_OPERATION = &DgError{4010, "Illegal operation"}

	RECORD_NOT_EXISTS     = &DgError{4014, "Record not exist"}
	DISABLED_USER         = &DgError{4015, "Disabled user"}
	EMAIL_REGISTERED      = &DgError{4016, "Email has been registered"}
	CUSTOMER_HAS_NO_GROUP = &DgError{4031, "该企业暂无国家对应团队"}
)
