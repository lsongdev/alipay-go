package alipay

type Trade struct {
	OutTradeNo  string `json:"out_trade_no"`
	TotalAmount string `json:"total_amount"`
	Subject     string `json:"subject"`
	Body        string `json:"body,omitempty"`
}

type TradePay struct {
	Trade
}

// alipay.trade.pay(统一收单交易支付接口)
// @docs https://opendocs.alipay.com/open/1f1fe18c_alipay.trade.pay?scene=32&pathHash=29c9a9ba
func (c *Client) TradePay(params TradePay) (Response, error) {
	return c.Execute("alipay.trade.pay", params)
}

type TradePreCreate struct {
	Trade
}

// alipay.trade.precreate(统一收单线下交易预创建)
// @docs https://opendocs.alipay.com/open/f540afd8_alipay.trade.precreate?pathHash=d3c84596
func (c *Client) TradePreCreate(params TradePreCreate) (Response, error) {
	return c.Execute("alipay.trade.precreate", params)
}

type TradeCreate struct {
	Trade

	BuyerOpenId  string `json:"buyer_open_id"`
	BuyerLogonId string `json:"buyer_logon_id"`
}

// alipay.trade.create(统一收单交易创建接口)
// @docs https://opendocs.alipay.com/open/f72f0792_alipay.trade.create?scene=2d8d65b1350f44bfa394347f06700c4f&pathHash=8919111c
func (c *Client) TradeCreate(params TradeCreate) (Response, error) {
	return c.Execute("alipay.trade.create", params)
}

type TradeQuery struct {
	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
}

// alipay.trade.query(统一收单交易查询)
// @docs https://opendocs.alipay.com/open/6f534d7f_alipay.trade.query?scene=23&pathHash=925e7dfc
func (c *Client) TradeQuery(params TradeQuery) (Response, error) {
	return c.Execute("alipay.trade.query", params)
}

type TradeRefund struct {
	RefundAmount string `json:"refund_amount"`
	OutTradeNo   string `json:"out_trade_no,omitempty"`
	TradeNo      string `json:"trade_no,omitempty"`
	RefundReason string `json:"refund_reason,omitempty"`
}

// alipay.trade.refund(统一收单交易退款接口)
// @docs https://opendocs.alipay.com/open/3aea9b48_alipay.trade.refund?scene=common&pathHash=b45b14f7
func (c *Client) TradeRefund(params TradeRefund) (Response, error) {
	return c.Execute("alipay.trade.refund", params)
}

type TradeClose struct {
	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
	OperatorId string `json:"operator_id,omitempty"`
}

// alipay.trade.close(统一收单交易关闭接口)
// @docs https://opendocs.alipay.com/open/e84f0d79_alipay.trade.close?scene=common&pathHash=7a011fc5
func (c *Client) TradeClose(params TradeClose) (Response, error) {
	return c.Execute("alipay.trade.close", params)
}
