package myWallet

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_example/util"
	"time"
)

var priKey []byte
var priBaseUrl = ""
var merchantCode = ""
var platformPubKey []byte

func Init(priKey_, priBaseUrl_, merchantCode_, platformPubKey_ string) error {
	pk, err := util.DecodeBase64(priKey_)
	if nil != err {
		return err
	}

	platformPub, err := util.DecodeBase64(platformPubKey_)
	if nil != err {
		return err
	}

	priKey = pk
	priBaseUrl = priBaseUrl_
	merchantCode = merchantCode_
	platformPubKey = platformPub
	return nil
}

func GetPlatformPubKey() []byte {
	return platformPubKey
}

func GetTokenList() ([]GetTokenListRes, error) {
	in := GetTokenListIn{
		CallTimestamp: time.Now().Unix(),
	}

	inBuf, err := json.Marshal(in)
	if nil != err {
		return nil, err
	}

	sign, err := util.NewRSAPemPKCS8().RsaSignWithSha256(inBuf, priKey)
	if nil != err {
		return nil, err
	}

	sign64 := util.EncodeBase64(sign)

	url := fmt.Sprintf("%s/GetTokenList", priBaseUrl)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header[ReqHeaderMerchantCode] = merchantCode
	header[ReqSign] = sign64

	body, err := util.GetInstanceByHttpUtil().PostClient(url, header, inBuf, 0, nil)
	if nil != err {
		return nil, err
	}

	res := new(Result)
	o := make([]GetTokenListRes, 0)
	res.Data = &o
	if e := json.Unmarshal(body, res); nil != e {
		return nil, e
	}

	if 2000 != res.Code {
		return nil, errors.New(fmt.Sprintf("code=%d,tip=%s", res.Code, res.Tip))
	}

	return o, nil
}

// PostCreateAddress 生成地址
// 地址根据不同 netWork 区分，可以用一个 netWork 的地址表示该网络下所有 token
func PostCreateAddress(netWork, callBackUrl string) (*PostCreateAddressOut, error) {
	in := PostCreateAddressIn{
		NetWork:       netWork,
		CallBackUrl:   callBackUrl,
		CallTimestamp: time.Now().Unix(),
	}

	inBuf, err := json.Marshal(in)
	if nil != err {
		return nil, err
	}

	sign, err := util.NewRSAPemPKCS8().RsaSignWithSha256(inBuf, priKey)
	if nil != err {
		return nil, err
	}

	sign64 := util.EncodeBase64(sign)

	url := fmt.Sprintf("%s/PostCreateAddress", priBaseUrl)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header[ReqHeaderMerchantCode] = merchantCode
	header[ReqSign] = sign64

	body, err := util.GetInstanceByHttpUtil().PostClient(url, header, inBuf, 0, nil)
	if nil != err {
		return nil, err
	}

	res := new(Result)
	o := new(PostCreateAddressOut)
	res.Data = o
	if e := json.Unmarshal(body, res); nil != e {
		return nil, e
	}

	if 2000 != res.Code {
		return nil, errors.New(fmt.Sprintf("code=%d,tip=%s", res.Code, res.Tip))
	}

	return o, nil
}

// PostChangeAddress 更新地址
// 修改地址的 callBackUrl
func PostChangeAddress(address, callBackUrl string) (*PostCreateAddressOut, error) {

	in := PostChangeAddressIn{
		Address:       address,
		CallBackUrl:   callBackUrl,
		CallTimestamp: time.Now().Unix(),
	}

	inBuf, err := json.Marshal(in)
	if nil != err {
		return nil, err
	}

	sign, err := util.NewRSAPemPKCS8().RsaSignWithSha256(inBuf, priKey)
	if nil != err {
		return nil, err
	}

	sign64 := util.EncodeBase64(sign)

	url := fmt.Sprintf("%s/PostChangeAddress", priBaseUrl)

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header[ReqHeaderMerchantCode] = merchantCode
	header[ReqSign] = sign64

	body, err := util.GetInstanceByHttpUtil().PostClient(url, header, inBuf, 0, nil)
	if nil != err {
		return nil, err
	}

	res := new(Result)
	o := new(PostCreateAddressOut)
	res.Data = o
	if e := json.Unmarshal(body, res); nil != e {
		return nil, e
	}

	if 2000 != res.Code {
		return nil, errors.New(fmt.Sprintf("code=%d,tip=%s", res.Code, res.Tip))
	}

	return o, nil
}

// GetTradeConfirm 交易链上确认数
// 2001 -> 交易处理中
// 2002 -> 未在链上找到交易
func GetTradeConfirm(tradeId string) (int, *GetTradeConfirmOut, error) {
	in := GetTradeConfirmIn{
		TradeId:       tradeId,
		CallTimestamp: time.Now().Unix(),
	}

	inBuf, err := json.Marshal(in)
	if nil != err {
		return 0, nil, err
	}

	sign, err := util.NewRSAPemPKCS8().RsaSignWithSha256(inBuf, priKey)
	if nil != err {
		return 0, nil, err
	}

	sign64 := util.EncodeBase64(sign)

	url := fmt.Sprintf("%s/GetTradeConfirm", priBaseUrl)
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header[ReqHeaderMerchantCode] = merchantCode
	header[ReqSign] = sign64

	body, err := util.GetInstanceByHttpUtil().PostClient(url, header, inBuf, 0, nil)
	if nil != err {
		return 0, nil, err
	}

	res := new(Result)
	o := new(GetTradeConfirmOut)
	res.Data = o
	if e := json.Unmarshal(body, res); nil != e {
		return 0, nil, e
	}

	if 2000 != res.Code {
		return res.Code, nil, errors.New(fmt.Sprintf("code=%d,tip=%s", res.Code, res.Tip))
	}

	return res.Code, o, nil
}

// GetTrade 查询/补单
// 查询交易，交易存在响应交易。交易不存在补充交易（从区块获得交易并回调）
// return code 2000 -> 交易存在响应交易 []
// return code 2001 -> 交易补充，并进行回调
// return code 2003 -> 未在链上找到交易
func GetTrade(netWork, transactionHash, address string) (int, []GetTradeOut, error) {
	in := GetTradeIn{
		TransactionHash: transactionHash,
		NetWork:         netWork,
		Address:         address,
		CallTimestamp:   time.Now().Unix(),
	}

	inBuf, err := json.Marshal(in)
	if nil != err {
		return 0, nil, err
	}

	sign, err := util.NewRSAPemPKCS8().RsaSignWithSha256(inBuf, priKey)
	if nil != err {
		return 0, nil, err
	}

	sign64 := util.EncodeBase64(sign)

	url := fmt.Sprintf("%s/GetTrade", priBaseUrl)
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header[ReqHeaderMerchantCode] = merchantCode
	header[ReqSign] = sign64

	body, err := util.GetInstanceByHttpUtil().PostClient(url, header, inBuf, 0, nil)
	if nil != err {
		return 0, nil, err
	}

	res := new(Result)
	o := make([]GetTradeOut, 0)
	res.Data = &o
	if e := json.Unmarshal(body, res); nil != e {
		return 0, nil, e
	}

	if 2000 != res.Code {
		return res.Code, nil, errors.New(fmt.Sprintf("code=%d,tip=%s", res.Code, res.Tip))
	}

	return res.Code, o, nil
}

// PostWithdraw 申请提币
func PostWithdraw(tradeId, addressTo, tokenFullName, amount, memo string) (*Result, error) {
	in := PostWithdrawIn{
		TradeId:       tradeId,
		AddressTo:     addressTo,
		TokenFullName: tokenFullName,
		Memo:          memo,
		Amount:        amount,
		CallTimestamp: time.Now().Unix(),
	}

	inBuf, err := json.Marshal(in)
	if nil != err {
		return nil, err
	}

	sign, err := util.NewRSAPemPKCS8().RsaSignWithSha256(inBuf, priKey)
	if nil != err {
		return nil, err
	}

	sign64 := util.EncodeBase64(sign)

	url := fmt.Sprintf("%s/PostWithdraw", priBaseUrl)
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header[ReqHeaderMerchantCode] = merchantCode
	header[ReqSign] = sign64

	body, err := util.GetInstanceByHttpUtil().PostClient(url, header, inBuf, 0, nil)
	if nil != err {
		return nil, err
	}

	res := new(Result)
	if e := json.Unmarshal(body, res); nil != e {
		return nil, e
	}

	if 2000 != res.Code {
		return nil, errors.New(fmt.Sprintf("code=%d,tip=%s", res.Code, res.Tip))
	}

	return res, nil
}

// GetWithdrawInfo 查询提币状态
// 2001 -> 未找到订单
func GetWithdrawInfo(tradeId string) (*GetWithdrawInfoRes, error) {
	in := GetWithdrawInfoIn{
		TradeId:       tradeId,
		CallTimestamp: time.Now().Unix(),
	}

	inBuf, err := json.Marshal(in)
	if nil != err {
		return nil, err
	}

	sign, err := util.NewRSAPemPKCS8().RsaSignWithSha256(inBuf, priKey)
	if nil != err {
		return nil, err
	}

	sign64 := util.EncodeBase64(sign)

	url := fmt.Sprintf("%s/GetWithdrawInfo", priBaseUrl)
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header[ReqHeaderMerchantCode] = merchantCode
	header[ReqSign] = sign64

	body, err := util.GetInstanceByHttpUtil().PostClient(url, header, inBuf, 0, nil)
	if nil != err {
		return nil, err
	}

	res := new(Result)
	o := new(GetWithdrawInfoRes)
	res.Data = o
	if e := json.Unmarshal(body, res); nil != e {
		return nil, e
	}

	if 2000 != res.Code {
		if 2001 == res.Code {
			return nil, nil
		}
		return nil, errors.New(fmt.Sprintf("code=%d,tip=%s", res.Code, res.Tip))
	}

	return o, nil

}

func GetCoinPrice() (map[string]map[string]string, error) {
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header[ReqHeaderMerchantCode] = merchantCode
	header[ReqSign] = merchantCode

	url := fmt.Sprintf("%s/GetCoinPrice", priBaseUrl)
	body, err := util.GetInstanceByHttpUtil().PostClient(url, header, nil, 0, nil)
	if nil != err {
		return nil, err
	}

	fmt.Println(string(body))

	res := new(Result)
	data := make(map[string]map[string]string)
	res.Data = &data

	if e := json.Unmarshal(body, &res); nil != e {
		return nil, e
	}

	return data, nil
}
