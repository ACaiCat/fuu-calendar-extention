package errno

import (
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var (
	Success            = NewErrNo(SuccessCode, consts.StatusOK, "成功")
	ParamErr           = NewErrNo(ParamErrCode, consts.StatusBadRequest, "参数错误")
	InternalServerErr  = NewErrNo(InternalServerErrCode, consts.StatusInternalServerError, "服务器内部错误")
	JwchLoginFailedErr = NewErrNo(JwchLoginFailedErrCode, consts.StatusForbidden, "教务处登录失败")
)
