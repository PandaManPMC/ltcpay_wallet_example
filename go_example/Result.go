package myWallet

import (
	"encoding/json"
)

type Result struct {
	Code int         `json:"code"`
	Tip  string      `json:"tip"`
	Data interface{} `json:"data"`
}

const (
	ResultCodeSuccess       = 2000 // 目标操作成功
	ResultCodeRedirect      = 2001 // 备选-特殊业务情况-重定向
	ResultCodeFailAuthority = 2002 // 没有权限
	ResultCodeRedirect2     = 2003 // 备选-特殊业务情况-重定向
	ResultCodeError         = 2004 // 系统错误
	ResultCodeFailParams    = 2010 // 缺少必要参数
	ResultCodeFail          = 2014 // 操作失败
	ResultCodeWarn          = 2015 // 操作出现一些意外情况，警告

	ResultCodeFrequent      = 2024 // 操作频繁，限控
	ResultCodeMaintain      = 2025 // 系统维护中
	ResultCodeSignatureFail = 2077 // 签名校验错误

)

// ToJsonString 将 Result 实例转为 json 字符串
// string 失败将为""
func (r *Result) ToJsonString() string {
	json, err := json.Marshal(r)
	if nil != err {
		return ""
	}
	return string(json)
}

// JSONMarshal 使用 json.Marshal()
func (r *Result) JSONMarshal() []byte {
	data, err := json.Marshal(r)
	if nil != err {
		return nil
	}
	return data
}

// ResultNew 创建一个Result实例
// code int	状态码
// tip string	提示
// data interface{}	数据
func ResultNew(code int, tip string, data interface{}) Result {
	return Result{code, tip, data}
}

// ResultSuccess 创建一个Result实例 - 成功的 ResultCodeSuccess
// data interface{}	数据
func ResultSuccess(data interface{}) Result {
	return Result{ResultCodeSuccess, "Successful", data}
}

func ResultSuccessP(data interface{}) *Result {
	return &Result{ResultCodeSuccess, "Successful", data}
}

func ResultSuccessTip(tip string) Result {
	return Result{ResultCodeSuccess, tip, nil}
}

func ResultCodeTip(code int, tip string) Result {
	return Result{code, tip, nil}
}

func ResultSuccessNilData() Result {
	return Result{ResultCodeSuccess, "Successful", nil}
}

func ResultNewSuccessNilData() Result {
	return Result{ResultCodeSuccess, "Successful", nil}
}

func ResultNewSuccessKV(k string, v any) Result {
	r := make(map[string]any)
	r[k] = v
	return Result{ResultCodeSuccess, "Successful", r}
}

func ResultNewSuccessKVAndTip(tip, k string, v any) Result {
	r := make(map[string]any)
	r[k] = v
	return Result{ResultCodeSuccess, tip, r}
}

// ResultNewSuccess 创建一个Result实例 - 成功的 ResultCodeSuccess
// tip string	提示
// data interface{}	数据
func ResultNewSuccess(tip string, data interface{}) Result {
	if "" == tip {
		tip = "Successful"
	}
	return Result{ResultCodeSuccess, tip, data}
}

type ListData struct {
	List  any `json:"list"`
	Count int `json:"count"`
}

type ListStatData struct {
	List  any `json:"list"`
	Count int `json:"count"`
	Stat  any `json:"stat"`
}

func ResultNewSuccessList(tip string, data interface{}, count int) Result {
	if "" == tip {
		tip = "Successful"
	}
	return Result{ResultCodeSuccess, tip, ListData{
		List:  data,
		Count: count,
	}}
}

// ResultNewSuccessStatList 创建一个Result实例 - 成功的 ResultCodeSuccess
// tip string	提示
// data interface{}	数据
// stat  interface{} 总计
// count int 总条数
func ResultNewSuccessStatList(tip string, data interface{}, count int, stat interface{}) Result {
	if "" == tip {
		tip = "Successful"
	}
	return Result{ResultCodeSuccess, tip, ListStatData{
		List:  data,
		Count: count,
		Stat:  stat,
	}}
}

// ResultNewRedirect 创建一个 Result 实例 - 重定向业务 ResultCodeRedirect
func ResultNewRedirect(tip string, data interface{}) Result {
	return Result{ResultCodeRedirect, tip, data}
}

// ResultNewRedirect2 创建一个 Result 实例 - 重定向业务 ResultCodeRedirect
func ResultNewRedirect2(tip string, data interface{}) Result {
	return Result{ResultCodeRedirect2, tip, data}
}

// ResultNewFail 创建一个Result实例 - 失败的 ResultNewFail
// tip string	提示
// data interface{}	数据
func ResultNewFail(tip string, data interface{}) Result {
	if "" == tip {
		tip = "Fail"
	}
	return Result{ResultCodeFail, tip, data}
}

// ResultNewFailTip 返回错误 2014，提示错误，data 为 nil
func ResultNewFailTip(tip string) Result {
	if "" == tip {
		tip = "Fail"
	}
	return Result{ResultCodeFail, tip, nil}
}

// ResultNewWarnTip 返回警告 2015，提示错误，data 为 nil
func ResultNewWarnTip(tip string) Result {
	if "" == tip {
		tip = "Warn"
	}
	return Result{ResultCodeWarn, tip, nil}
}

// ResultNewFailParams 返回 code  ResultCodeFailParams ，参数错误
func ResultNewFailParams(tip string) Result {
	if "" == tip {
		tip = "Fail"
	}
	return Result{ResultCodeFailParams, tip, nil}
}

// ResultNewFailByNotFound 创建一个Result实例 - 失败的 未找到
// data interface{}	数据
func ResultNewFailByNotFound(data interface{}) Result {
	return Result{ResultCodeFail, "Not Fount", data}
}

// ResultNewFailRequiredParameter 创建一个Result实例 - 失败的 ResultNewFail，缺少必要参数的通用提示
// data interface{}	数据
func ResultNewFailRequiredParameter(data interface{}) Result {
	return Result{ResultCodeFail, "Missing required parameter", data}
}

// ResultNewError 创建一个Result实例 - 系统错误的 ResultNewError
// tip string	提示
// data interface{}	数据
func ResultNewError(tip string, data interface{}) Result {
	if "" == tip {
		tip = "Error"
	}
	return Result{ResultCodeError, tip, data}
}

// ResultNewSysError 创建一个Result实例 - 系统错误的 ResultNewError
// data interface{}	数据
func ResultNewSysError(data interface{}) Result {
	return Result{ResultCodeError, "The system encountered an unexpected error", data}
}

func ResultNewSysErrorNilData() Result {
	return Result{ResultCodeError, "The system encountered an unexpected error", nil}
}

func ResultNewSysErrorNilDataP() *Result {
	return &Result{ResultCodeError, "The system encountered an unexpected error", nil}
}

// ResultNewFailPermissionDenied 创建一个Result实例 - 没有权限的 ResultNewFailAuthority
// tip string	提示
// data interface{}	数据
func ResultNewFailPermissionDenied(tip string, data interface{}) Result {
	if "" == tip {
		tip = "Permission denied"
	}
	return Result{Code: ResultCodeFailAuthority, Tip: tip, Data: data}
}

// ResultNewSignatureFail 签名错误
func ResultNewSignatureFail() Result {
	return Result{ResultCodeSignatureFail, "Signature wrong", nil}
}

func ResultNewFrequent() Result {
	return Result{ResultCodeFrequent, "操作频繁，请稍后再试。", nil}
}
