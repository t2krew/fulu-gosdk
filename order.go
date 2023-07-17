package fulu_gosdk

import (
	"context"
	jsoniter "github.com/json-iterator/go"
)

type OrderState string

const (
	OrderStateSuccess    = "success"    // 成功
	OrderStateProcessing = "processing" // 处理中
	OrderStateFailed     = "failed"     // 失败
	OrderStateUntreated  = "untreated"  // 未处理
)

type CreateDirectOrderBizContent struct {
	ProductID        int64   `json:"product_id"`
	CustomerOrder    string  `json:"customer_order"`
	ChargeAccount    string  `json:"charge_account"`
	BuyNum           int     `json:"buy_num"`
	ChargeGameName   string  `json:"charge_game_name"`
	ChargeGameRegion string  `json:"charge_game_region"`
	ChargeType       string  `json:"charge_type"`
	ChargePassword   string  `json:"charge_password"`
	ChargeIp         string  `json:"charge_ip"`
	ContactQQ        string  `json:"contact_qq"`
	ContactTel       string  `json:"contact_tel"`
	RemainingNumber  string  `json:"remaining_number"`
	ChargeGameRole   string  `json:"charge_game_role"`
	CustomerPrice    float64 `json:"customer_price"`
	ShopType         string  `json:"shop_type"`
	ExternalBizId    string  `json:"external_biz_id"`
}

type DirectOrderResult struct {
	OrderID              string  `json:"order_id"`
	CustomerOrderNO      string  `json:"customer_order_no"`
	ProductID            int64   `json:"product_id"`
	ProductName          string  `json:"product_name"`
	ChargeAccount        string  `json:"charge_account"`
	BuyNum               int     `json:"buy_num"`
	OrderType            int     `json:"order_type"`
	OrderPrice           float64 `json:"order_price"`
	PrderState           string  `json:"prder_state"`
	CreateTime           string  `json:"create_time"`
	FinishTime           string  `json:"finish_time"`
	Area                 string  `json:"area"`
	Server               string  `json:"server"`
	Type                 string  `json:"type"`
	OperatorSerialNumber string  `json:"operator_serial_number"`
}

// CreateDirectOrder 创建直充订单
func (c *Client) CreateDirectOrder(ctx context.Context, params CreateDirectOrderBizContent) (*DirectOrderResult, error) {
	var result DirectOrderResult
	err := c.Request(ctx, MethodCreateDirectOrder, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type CreateCardOrderBizContent struct {
	ProductID       int64   `json:"product_id"`
	BuyNum          int     `json:"buy_num"`
	CustomerOrderNO string  `json:"customer_order_no"`
	CustomerPrice   float64 `json:"customer_price"`
	ShopType        string  `json:"shop_type"`
	ExternalBizID   string  `json:"external_biz_id"`
}

type CardOrderResult struct {
	OrderID              string  `json:"order_id"`
	CustomerOrderNO      string  `json:"customer_order_no"`
	ProductID            int64   `json:"product_id"`
	ProductName          string  `json:"product_name"`
	BuyNum               int     `json:"buy_num"`
	OrderType            int     `json:"order_type"`
	OrderPrice           float64 `json:"order_price"`
	OrderState           string  `json:"order_state"`
	CreateTime           string  `json:"create_time"`
	FinishTime           string  `json:"finish_time"`
	OperatorSerialNumber string  `json:"operator_serial_number"`
}

// CreateCardOrder 创建卡密订单
func (c *Client) CreateCardOrder(ctx context.Context, params CreateCardOrderBizContent) (*CardOrderResult, error) {
	var result CardOrderResult
	err := c.Request(ctx, MethodCreateCardOrder, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type CreateMobileOrderBizContent struct {
	ChargePhone     string  `json:"charge_phone"`
	ChargeValue     float64 `json:"charge_value"`
	CustomerOrderNO string  `json:"customer_order_no"`
	CustomerPrice   float64 `json:"customer_price"`
	ShopType        string  `json:"shop_type"`
	ExternalBizID   string  `json:"external_biz_id"`
}

type MobileOrderResult struct {
	OrderID              string  `json:"order_id"`
	CustomerOrderNO      string  `json:"customer_order_no"`
	ProductID            int64   `json:"product_id"`
	ProductName          string  `json:"product_name"`
	ChargeAccount        string  `json:"charge_account"`
	BuyNum               int     `json:"buy_num"`
	OrderPrice           float64 `json:"order_price"`
	OrderType            int     `json:"order_type"`
	OrderState           string  `json:"order_state"`
	CreateTime           string  `json:"create_time"`
	FinishTime           string  `json:"finish_time"`
	OperatorSerialNumber string  `json:"operator_serial_number"`
}

// CreateMobileOrder 创建话费订单
func (c *Client) CreateMobileOrder(ctx context.Context, params CreateMobileOrderBizContent) (*MobileOrderResult, error) {
	var result MobileOrderResult
	err := c.Request(ctx, MethodCreateMobileOrder, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CardItem 卡密商品
type CardItem struct {
	CardType     int    `json:"card_type"`
	CardNumber   string `json:"card_number"`
	CardPwd      string `json:"card_pwd"`
	CardDeadline string `json:"card_deadline"`
}

// Order 订单信息
type Order struct {
	OrderID              string     `json:"order_id"`
	CustomerOrderNO      string     `json:"customer_order_no"`
	ProductID            int64      `json:"product_id"`
	ProductName          string     `json:"product_name"`
	ChargeAccount        string     `json:"charge_account"`
	BuyNum               int        `json:"buy_num"`
	OrderPrice           float64    `json:"order_price"`
	OrderType            int        `json:"order_type"`
	OrderState           string     `json:"order_state"`
	CreateTime           string     `json:"create_time"`
	FinishTime           string     `json:"finish_time"`
	Area                 string     `json:"area"`
	Server               string     `json:"server"`
	Type                 string     `json:"type"`
	Cards                []CardItem `json:"cards"`
	OperatorSerialNumber string     `json:"operator_serial_number"`
}

// QueryOrder 订单查询
func (c *Client) QueryOrder(ctx context.Context, customerOrderNO string) (*Order, error) {
	var result Order
	var params = map[string]string{
		"customer_order_no": customerOrderNO,
	}
	err := c.Request(ctx, MethodQueryOrder, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type OrderExtendContent struct {
	ExpressNumber       string `json:"express_number"`
	RechargeDescription string `json:"recharge_description"`
	ExternalBizID       string `json:"external_biz_id"`
}

type OrderExtendTemp struct {
	OrderID            string `json:"order_id"`
	CustomerOrderNO    string `json:"customer_order_no"`
	OrderExtendContent string `json:"order_extend_content"`
}

type OrderExtend struct {
	OrderID            string              `json:"order_id"`
	CustomerOrderNO    string              `json:"customer_order_no"`
	OrderExtendContent *OrderExtendContent `json:"order_extend_content"`
}

// QueryOrderExtend 订单扩展信息查询
func (c *Client) QueryOrderExtend(ctx context.Context, customerOrderNO string) (*OrderExtend, error) {
	var result OrderExtendTemp
	var params = map[string]string{
		"customer_order_no": customerOrderNO,
	}
	err := c.Request(ctx, MethodQueryOrderExtend, params, &result)
	if err != nil {
		return nil, err
	}
	var content OrderExtendContent
	err = jsoniter.Unmarshal([]byte(result.OrderExtendContent), &content)
	if err != nil {
		return nil, err
	}
	return &OrderExtend{
		OrderID:            result.OrderID,
		CustomerOrderNO:    result.CustomerOrderNO,
		OrderExtendContent: &content,
	}, nil
}

// GetReconciliation 对账单申请
func (c *Client) GetReconciliation(ctx context.Context) {

}
