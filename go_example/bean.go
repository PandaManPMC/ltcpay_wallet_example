package myWallet

type GetTokenListIn struct {
	CallTimestamp int64 `json:"callTimestamp" required:"true"`
}

type GetTokenListRes struct {
	TokenFullName     string `json:"tokenFullName"`     // 代币全称
	NetWork           string `json:"netWork"`           // 主网
	ChainName         string `json:"chainName"`         // 链名
	TokenName         string `json:"tokenName"`         // 代币名称
	TokenSymbol       string `json:"tokenSymbol"`       // 代币符号
	ContractAddress   string `json:"contractAddress"`   // 合约地址
	AmountDecimals    uint32 `json:"amountDecimals"`    // 精度
	MainTokenFullName string `json:"mainTokenFullName"` // 主链代币
	MinWithdrawAmount string `json:"minWithdrawAmount"` // 最小提款金额
}

type PostCreateAddressIn struct {
	NetWork       string `json:"netWork" required:"true"`       // 主网
	CallBackUrl   string `json:"callBackUrl" required:"true"`   // 回调地址
	CallTimestamp int64  `json:"callTimestamp" required:"true"` // 调用时间
}

type PostCreateAddressOut struct {
	Address string `json:"address"` // 地址
}

type PostChangeAddressIn struct {
	Address       string `json:"address" required:"true"`       // 主网
	CallBackUrl   string `json:"callBackUrl" required:"true"`   // 回调地址
	CallTimestamp int64  `json:"callTimestamp" required:"true"` // 调用时间
}

type GetTradeConfirmIn struct {
	TradeId       string `json:"tradeId" required:"true"`       // 业务流水号
	CallTimestamp int64  `json:"callTimestamp" required:"true"` // 调用时间
}

type GetTradeIn struct {
	TransactionHash string `json:"transactionHash" required:"true"` // 交易 hash
	NetWork         string `json:"netWork" required:"true"`         // 主网
	Address         string `json:"address" required:"true"`         // 地址
	CallTimestamp   int64  `json:"callTimestamp" required:"true"`   // 调用时间
}

type GetTradeOut struct {
	TradeId         string `json:"tradeId"`         // search业务流水号
	TransactionHash string `json:"transactionHash"` // 交易 hash
	AddressFrom     string `json:"addressFrom"`
	AddressTo       string `json:"addressTo"`
	NetWork         string `json:"netWork"`
	TokenFullName   string `json:"tokenFullName"`
	Amount          string `json:"amount"`
	Block           uint64 `json:"block"`
	Timestamp       uint64 `json:"timestamp"`
	CallTimestamp   int64  `json:"callTimestamp"`
}

type GetTradeConfirmOut struct {
	TradeId       string `json:"tradeId"`      // search业务流水号
	ConfirmBlock  uint64 `json:"confirmBlock"` // 区块确认数
	Height        uint64 `json:"height"`       // 当前区块高度
	Status        string `json:"status"`       // 状态,SUCCESS 为成功，其它为失败
	CallTimestamp int64  `json:"callTimestamp"`
}

type PostWithdrawIn struct {
	TradeId       string `json:"tradeId" required:"true" min:"10" max:"32"` // 业务流水号
	AddressTo     string `json:"addressTo" required:"true"`                 // 收款地址
	TokenFullName string `json:"tokenFullName" required:"true"`             // 代币全称
	Memo          string `json:"memo" required:"true"`                      // 备忘码
	Amount        string `json:"amount" required:"true"`                    // 金额(链上整数金额)
	CallTimestamp int64  `json:"callTimestamp" required:"true"`             // 调用时间
}

type GetWithdrawInfoIn struct {
	TradeId       string `json:"tradeId" required:"true" max:"32"` // 业务流水号
	CallTimestamp int64  `json:"callTimestamp" required:"true"`    // 调用时间
}

type GetWithdrawInfoRes struct {
	TradeId         string `json:"tradeId"`         // 业务流水号
	NetWork         string `json:"netWork"`         // 主网
	TokenFullName   string `json:"tokenFullName"`   // 代币全称
	TransactionHash string `json:"transactionHash"` // 交易哈希
	AddressFrom     string `json:"addressFrom"`     // 发起地址
	AddressTo       string `json:"addressTo"`       // 到达地址
	Amount          string `json:"amount"`          // 金额
	StateTransfer   uint8  `json:"stateTransfer"`   // 转账状态:1@待提交;2@区块处理中;3@交易失败;4@处理完成;5@取消订单;6@待审核;7@审核拒绝
	Status          string `json:"status"`          // 状态
}

// CallMerchantTransactionData 交易回调
type CallMerchantTransactionData struct {
	TradeId         string `json:"tradeId"`         // 业务流水号
	TransactionHash string `json:"transactionHash"` // 交易 hash
	AddressFrom     string `json:"addressFrom"`     // 转账地址
	AddressTo       string `json:"addressTo"`       // 收款地址
	NetWork         string `json:"netWork"`         // 主网
	TokenFullName   string `json:"tokenFullName"`   // 代币全称
	Amount          string `json:"amount"`          // 金额
	Block           uint64 `json:"block"`           // 区块
	ConfirmBlock    uint64 `json:"confirmBlock"`    // 区块确认数
	Timestamp       uint64 `json:"timestamp"`       // 时间龊（无意义字段）
	Status          string `json:"status"`          // 状态，SUCCESS 为成功,为空字符串则还没有状态，其它均为失败
	Memo            string `json:"memo"`            // memo
	CallTimestamp   int64  `json:"callTimestamp"`   // 请求时间
}
