package util

const (
	SUCCESS        = 0
	ERROR          = 1
	INVALID_PARAMS = 400

	ERROR_EXIST_TAG         = 10001
	ERROR_NOT_EXIST_TAG     = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	STATUS_VALID    = 10  // 账号有效或者已启用
	STATUS_INVALID  = -10 // 临时
	STATUS_INACTIVE = -10 // 账号未启用
	STATUS_DELETED  = -20 // 账号删除
	STATUS_BLOCKED  = -30 // 账号冻结

	// database
	ERROR_CODE__DB__INSERT                      = 1001
	ERROR_CODE__DB__READ                        = 1002
	ERROR_CODE__DB__UPDATE                      = 1003
	ERROR_CODE__DB__NO_ROW                      = 1007
	ERROR_CODE__DB__UNEXPECTED                  = 1008
	ERROR_CODE__DB__TRANSACTION_BEGIN_FAILED    = 1009
	ERROR_CODE__DB__TRANSACTION_COMMIT_FAILED   = 1010
	ERROR_CODE__DB__UPDATE_UNEXPECTED           = 1011
	ERROR_CODE__DB__TRANSACTION_ROLLBACK_FAILED = 1010

	// rbac
	ERROR_CODE__ROLE__NOT_EXIST      = 5101
	ERROR_CODE__USER_ROLE__NOT_EXIST = 5102

	// parameter
	ERROR_CODE__PARAM__ILLEGAL = 1101

	// address
	ERROR_CODE__PROVINCE__NOT_EXIST = 4801
	ERROR_CODE__CITY__NOT_EXIST     = 4802
	ERROR_CODE__DISTRICT__NOT_EXIST = 4803
	ERROR_CODE__STREET__NOT_EXIST   = 4803

	// json
	ERROR_CODE__JSON__UNMARSHAL_FAILED = 1105

	// type
	ERROR_CODE__TYPE__INVALID     = 4101
	ERROR_CODE__ORG_TYPE__ILLEGAL = 4102

	// rpc
	ERROR_CODE__RPC__CALL = 4401

	// verification code
	ERROR_CODE__VERIFICATION_CODE__NOT_MATCH = 4501

	// http
	ERROR_CODE__SOURCE_DATA__ILLEGAL      = 1101 // 外部传入参数错误
	ERROR_CODE__GRPC__FAILED              = 1102 // grpc 调用失败
	ERROR_CODE__HTTP__CALL_FAILD_EXTERNAL = 1103 // http外部调用失败
	ERROR_CODE__HTTP__CALL_FAILD_INTERNAL = 1104 // http内部调用失败
	ERROR_CODE__JSON__PARSE_FAILED        = 1105 // JSON解析失败
	ERROR_CODE__HTTP__INPUT_EMPTY         = 1106

	FIELD_IS_EMPTY  = 1201 // 数据字段为空
	FIELD_DUPLICATE = 1202 // 数据字段重复

	// redis 调用失败
	REDIS_SET_FAILED = 1301
	REDIS_GET_FAILED = 1302

	// session
	ERROR_CODE__SESSION__START_FAILED  = 1404
	ERROR_CODE__SESSION__EMPTY_SESSION = 1410
	ERROR_CODE__SESSION__SET_FAILED    = 1411

	// session 相关错误码
	SESSION_ERROR_NO_USER_ID         = 1401
	SESSION_ERROR_NO_COMPANY_ID      = 1402
	SESSION_ERROR_EMPTY_INPUT        = 1403
	SESSION_ERROR_START_FAIL         = 1404
	SESSION_ERROR_NO_USER_NAME       = 1405
	SESSION_ERROR_NO_COMPANY_NAME    = 1406
	SESSION_ERROR_EMPTY_SESSION      = 1407
	SESSION_ERROR_INVALID_USER_ID    = 1408
	SESSION_ERROR_INVALID_COMPANY_ID = 1409
	SESSION_ERROR_NO_SESSION         = 1410

	// header 相关错误码
	ERROR_CODE__HEADER__NO_USER_ID               = 1501
	ERROR_CODE__HEADER__NO_COMPANY_ID            = 1502
	ERROR_CODE__HEADER__NO_USER_NAME             = 1504
	ERROR_CODE__HEADER__NO_COMPANY_NAME          = 1505
	ERROR_CODE__HEADER__NO_POSITIONS             = 1506
	ERROR_CODE__HEADER__POSITIONS_FORMAT_ILLEGAL = 1507
	ERROR_CODE__HEADER__NO_USER_MOBILE           = 1509
	ERROR_CODE__HEADER__NO_USER_CODE             = 1510

	HEADER_ERROR_NO_USER_ID      = 1501
	HEADER_ERROR_NO_COMPANY_ID   = 1502
	HEADER_ERROR_EMPTY_INPUT     = 1503
	HEADER_ERROR_NO_USER_NAME    = 1504
	HEADER_ERROR_NO_COMPANY_NAME = 1505

	// 用户状态
	USER_NOT_EXIST  = 2001
	USER_EXISTED    = 2002
	USER_PW_ERROR   = 2003
	USER_LOGGED_IN  = 2004
	USER_LOGGED_OUT = 2005
	USER_ACTIVE     = 2006 // 账号已启用
	USER_INACTIVE   = 2006 // 账号未启用
	USER_UNBLOCKED  = 2007 // 账号未冻结

	// etcd相关操作错误码
	ETCD_CREATE_DIR_ERROR = 3001
	ETCD_CREATE_KEY_ERROR = 3002
	ETCD_READ_KEY_ERROR   = 3003

	// IO文件解析错误码
	FILE_OPEN_FAILED = 4001
	FILE_READ_FAILED = 4002

	// GRBAC权限管理错误码
	GRBAC_PERMISSION_REJECT_ACCESS = 4021 // 拒绝访问，没有权限
	// header
	KEY__HEADER__USER_ID      = "Zhibu-Account-User-Id"
	KEY__HEADER__USER_CODE    = "Zhibu-Account-User-Code"
	KEY__HEADER__USER_NAME    = "Zhibu-Account-User-Name"
	KEY__HEADER__USER_MOBILE  = "Zhibu-Account-User-Mobile"
	KEY__HEADER__COMPANY_ID   = "Zhibu-Account-Company-Id"
	KEY__HEADER__COMPANY_NAME = "Zhibu-Account-Company-Name"
	KEY__HEADER__POSITIONS    = "Zhibu-Account-Positions"

	// session
	KEY__SESSION__USER_ID      = "Zhibu-Account-User-Id"
	KEY__SESSION__USER_CODE    = "Zhibu-Account-User-Code"
	KEY__SESSION__USER_NAME    = "Zhibu-Account-User-Name"
	KEY__SESSION__USER_MOBILE  = "Zhibu-Account-User-Mobile"
	KEY__SESSION__COMPANY_ID   = "Zhibu-Account-Company-Id"
	KEY__SESSION__COMPANY_NAME = "Zhibu-Account-Company-Name"
	KEY__SESSION__POSITIONS    = "Zhibu-Account-Positions"

	// auto complete
	KEY__AUTO_COMPLETE__SIZE = 8
)

// 常用的状态code,写入里面
var ResponseMap = map[int]string{
	SUCCESS:                        "OK",
	ERROR:                          "server error",
	INVALID_PARAMS:                 "invalid params",
	ERROR_CODE__JSON__PARSE_FAILED: "request parse params fail.",
}
