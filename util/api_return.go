package util

import (
	"time"
)

// 普通json格式
func ReturnJson(code int, msg string, data interface{}) (jsonData map[string]interface{}) {
	jsonData = make(map[string]interface{}, 3)
	jsonData["time_stamp"] = time.Now()
	jsonData["code"] = code
	jsonData["msg"] = msg
	if data != nil {
		jsonData["data"] = data
	}
	return
}

func ReturnSuccessJson(data interface{}) map[string]interface{} {
	return ReturnJson(SUCCESS, ResponseMap[SUCCESS], data)
}
func ReturnServerErrJson(data interface{}) map[string]interface{} {
	return ReturnJson(ERROR, ResponseMap[ERROR], data)
}
func ReturnParamErrJson(data interface{}) map[string]interface{} {
	return ReturnJson(INVALID_PARAMS, ResponseMap[INVALID_PARAMS], data)
}

// layui 后台返回需要的json格式
func LayuiJson(code int, msg string, data, count, pageNo, pageSize interface{}) (jsonData map[string]interface{}) {
	jsonData = make(map[string]interface{}, 3)
	jsonData["code"] = code
	jsonData["msg"] = msg
	if count != false {
		jsonData["data"] = PageUtil(count.(int64), pageNo.(int64), pageSize.(int64), data)
	} else {
		jsonData["data"] = data
	}

	jsonData["time_stamp"] = time.Now()
	return
}

// bootstrap table 返回json
func TableJson(data, offset, limit, total interface{}) (jsonData map[string]interface{}) {
	jsonData = make(map[string]interface{}, 3)
	jsonData["rows"] = data
	jsonData["offset"] = offset
	jsonData["limit"] = limit
	jsonData["total"] = total
	jsonData["time_stamp"] = time.Now()
	return
}
