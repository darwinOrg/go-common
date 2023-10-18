package dgerr

type DgErrorML struct {
	Code        int
	MessageCode string
}

func (e *DgErrorML) Error() string {
	return e.MessageCode
}

func NewDgErrorML(code int, msgCode string) *DgErrorML {
	return &DgErrorML{
		Code:        code,
		MessageCode: msgCode,
	}
}

func SimpleDgErrorML(msgCode string) *DgErrorML {
	return &DgErrorML{
		Code:        -1,
		MessageCode: msgCode,
	}
}

var (
	SYSTEM_ERROR_ML = &DgErrorML{5001, "common.system_error"}
	SYSTEM_BUSY_ML  = &DgErrorML{5002, "common.system_invalid"}
	TIME_OUT_ML     = &DgError{5003, "common.time_out"}

	ARGUMENT_NOT_VALID_ML    = &DgErrorML{4001, "common.argument_not_valid"}
	DUPLICATE_PRIMARY_KEY_ML = &DgErrorML{4004, "common.duplicate_primary_key"}
	NOT_LOGIN_IN_ML          = &DgErrorML{4006, "common.not_login_in"}
	USER_NOT_EXISTS_ML       = &DgErrorML{4007, "common.user_not_exist"}
	NO_PERMISSION_ML         = &DgErrorML{4009, "common.no_permission"}
	ILLEGAL_OPERATION_ML     = &DgErrorML{4010, "common.illegal_operation"}
	RECORD_NOT_EXISTS_ML     = &DgErrorML{4014, "common.record_not_exists"}
)
