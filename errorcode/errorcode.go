package errorcode

const (
	// Success response ok
	Success = 0

	// params error

	// ErrParams general params error
	ErrParams = 40000

	// ErrAuthenticationRequired un authenticated
	ErrAuthenticationRequired = 40001

	// ErrRolePermissionRequired role no permission
	ErrRolePermissionRequired = 40101

	// ErrParamRequired param required
	ErrParamRequired = 40002

	// ErrParamRangeErr param format error
	ErrParamRangeErr = 40003

	// ErrDataSourceErr param format error
	ErrDataSourceErr = 40004

	// ErrAuthDenyErr param format error
	ErrAuthDenyErr = 40105

	// internal logic error

	// ErrBizLogic internal logic error
	ErrBizLogic = 50000

	//ErrBizRecordNotFound record not found
	ErrBizRecordNotFound = 50001

	//ErrBizRecordDuplicate record duplicate
	ErrBizRecordDuplicate = 50002

	// service error

	// ErrService service general errors
	ErrService = 60000

	// ErrServiceNotAvailable service not available
	ErrServiceNotAvailable = 60001

	// ErrServiceRespCodeErr service response code error
	ErrServiceRespCodeErr = 60002

	// ErrServiceRespDataErr service response json error code error
	ErrServiceRespDataErr = 60003

	// ErrServiceRespTimeout service timeout
	ErrServiceRespTimeout = 60004

	// ErrServiceRespNoData no data
	ErrServiceRespNoData = 60005

	// infrastructure error

	// ErrInfrastructure general errors
	ErrInfrastructure = 70000

	// ErrInfraMySQL mysql error
	ErrInfraMySQL = 70001

	// ErrInfraRedis redis error
	ErrInfraRedis = 70002

	// ErrInfraEmail email client error
	ErrInfraEmail = 70003

	// ErrCms send cms error
	ErrCms = 80000
)

const (
	// SuccessMsg success msg
	SuccessMsg = "success"

	// DefaultErrorMsg general error msg
	DefaultErrorMsg = "服务器内部错误"

	//ErrParamsMsg params err
	ErrParamsMsg = "参数错误"

	//RolePermissionRequiredMsg 角色无权限
	RolePermissionRequiredMsg = "角色无权限"

	//RoleUpdateMsg 角色修改错误
	RoleUpdateMsg = "角色修改错误，请重试"

	//RecordNotFoundMsg record not found
	RecordNotFoundMsg = "记录不存在"

	//RecordDuplicateMsg record duplicate
	RecordDuplicateMsg = "记录已存在"

	//RetryCommitMsg retry commit
	RetryCommitMsg = "请重新尝试或联系管理员"

	//SystemSettingMsg system setting msg
	SystemSettingMsg = "系统设置信息，无法删除"
)
