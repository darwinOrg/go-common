package dgerr

type DgErrorML struct {
	Code        int
	MessageCode string
	Cause       error
}

func (e *DgErrorML) WithCause(err error) *DgErrorML {
	if e == nil {
		return nil
	}

	return &DgErrorML{Code: e.Code, MessageCode: e.MessageCode, Cause: err}
}

func (e *DgErrorML) Error() string {
	if e == nil {
		return ""
	}

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

func SimpleDgErrorMLWithCause(msgCode string, err error) *DgErrorML {
	return &DgErrorML{
		Code:        -1,
		MessageCode: msgCode,
		Cause:       err,
	}
}

var (
	SYSTEM_ERROR_ML = &DgErrorML{Code: 5001, MessageCode: "common.system_error"}
	SYSTEM_BUSY_ML  = &DgErrorML{Code: 5002, MessageCode: "common.system_invalid"}
	TIME_OUT_ML     = &DgErrorML{Code: 5003, MessageCode: "common.time_out"}

	ARGUMENT_NOT_VALID_ML    = &DgErrorML{Code: 4001, MessageCode: "common.argument_not_valid"}
	DUPLICATE_PRIMARY_KEY_ML = &DgErrorML{Code: 4004, MessageCode: "common.duplicate_primary_key"}
	NOT_LOGIN_IN_ML          = &DgErrorML{Code: 4006, MessageCode: "common.not_login_in"}
	USER_NOT_EXISTS_ML       = &DgErrorML{Code: 4007, MessageCode: "common.user_not_exist"}
	NO_PERMISSION_ML         = &DgErrorML{Code: 4009, MessageCode: "common.no_permission"}
	ILLEGAL_OPERATION_ML     = &DgErrorML{Code: 4010, MessageCode: "common.illegal_operation"}
	RECORD_EXISTS_ML         = &DgErrorML{Code: 4013, MessageCode: "common.record_exists"}
	RECORD_NOT_EXISTS_ML     = &DgErrorML{Code: 4014, MessageCode: "common.record_not_exists"}
)
