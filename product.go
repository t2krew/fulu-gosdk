package fulu_gosdk

import "context"

type ProductDetailFormat int

const (
	ProductDetailFormatPlain = 0 // 纯文本
	ProductDetailFormatJSON  = 1 // json
)

// SaleStatus 销售状态
type SaleStatus string

const (
	SaleStatusValid         = SaleStatus("上架")
	SaleStatusInvalid       = SaleStatus("下架")
	SaleStatusMaintain      = SaleStatus("维护中")
	SaleStatusStockMaintain = SaleStatus("库存维护")
)

// StockStatus 库存状态
type StockStatus string

const (
	StockStatusEnough = StockStatus("充足")
	StockStatusOut    = StockStatus("断货")
	StockStatusAlarm  = StockStatus("警报")
)

// GetProductListParams 获取商品列表请求参数
type GetProductListParams struct {
	ProductID        int64   `json:"product_id,omitempty"`
	ProductName      string  `json:"product_name,omitempty"`
	ProductType      string  `json:"product_type,omitempty"`
	FaceValue        float64 `json:"face_value,omitempty"`
	FirstCategoryID  int     `json:"first_category_id,omitempty"`
	SecondCategoryID int     `json:"second_category_id,omitempty"`
	ThirdCategoryID  int     `json:"third_category_id,omitempty"`
}

// ProductListItem 商品列表项
type ProductListItem struct {
	ProductID     int64   `json:"product_id"`
	ProductName   string  `json:"product_name"`
	ProductType   string  `json:"product_type"`
	FaceValue     float64 `json:"face_value"`
	PurchasePrice float64 `json:"purchase_price"`
	SalesStatus   string  `json:"sales_status"`
	StockStatus   string  `json:"stock_status"`
	TemplateID    string  `json:"template_id"`
	Details       string  `json:"details"`
}

// GetProductList 获取商品列表
// method: fulu.goods.list.get
func (c *Client) GetProductList(ctx context.Context, params *GetProductListParams) ([]ProductListItem, error) {
	var result []ProductListItem
	err := c.Request(ctx, MethodGetProductList, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetProductInfoParams 获取商品信息请求参数
type GetProductInfoParams struct {
	ProductID    string `json:"product_id"`
	DetailFormat int    `json:"detail_format,omitempty"`
}

// ProductInfo 商品信息
type ProductInfo struct {
	ProductID        int64   `json:"product_id"`
	ProductName      string  `json:"product_name"`
	FaceValue        float64 `json:"face_value"`
	ProductType      string  `json:"product_type"`
	PurchasePrice    float64 `json:"purchase_price"`
	TemplateID       string  `json:"template_id"`
	StockStatus      string  `json:"stock_status"`
	SalesStatus      string  `json:"sales_status"`
	Details          string  `json:"details"`
	FourCategoryIcon string  `json:"four_category_icon"`
	DetailType       int     `json:"detail_type"`
}

// GetProductInfo 获取商品信息
// method: fulu.goods.info.get
func (c *Client) GetProductInfo(ctx context.Context, productID string, format ...ProductDetailFormat) (*ProductInfo, error) {
	var params = &GetProductInfoParams{
		ProductID: productID,
	}
	if len(format) > 0 {
		params.DetailFormat = int(format[0])
	}
	var result ProductInfo
	err := c.Request(ctx, MethodGetProductInfo, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetProductTemplateParams 获取商品模板请求参数
type GetProductTemplateParams struct {
	TemplateID string `json:"template_id"` // 商品模板编号
}

// ProductTemplate 商品模板内容
type ProductTemplate struct {
	AddressID               string                            `json:"AddressId"`               // 商品模板编号
	ElementInfo             productTemplateElementInfo        `json:"ElementInfo"`             // 包括元素
	AddressName             string                            `json:"AddressName"`             // 模板名称
	IsServiceArea           bool                              `json:"IsServiceArea"`           // 是否有区服（预留字段，不用关注）
	GameTempaltePreviewList []productTemplateGameTemplateItem `json:"GameTempaltePreviewList"` // 游戏区服模板信息
}

type productTemplateElementInfo struct {
	Inputs    []productTemplateElementInput   `json:"Inputs"`
	ChargeNum productTemplateElementChargeNum `json:"ChargeNum"`
}

type productTemplateElementInput struct {
	Type   string `json:"Type"`
	ID     string `json:"Id"`
	Name   string `json:"Name"`
	SortId int    `json:"SortId"`
}

type productTemplateElementChargeNum struct {
	ID     string                              `json:"Id"`
	Name   string                              `json:"Name"`
	Value  string                              `json:"Value"`
	Unit   productTemplateElementChargeNumUnit `json:"Unit"`
	Type   string                              `json:"Type"`
	SortId int                                 `json:"SortId"`
}

type productTemplateElementChargeNumUnit struct {
	DefaultUint      interface{} `json:"defaultUint"`
	DefaultUnitAfter interface{} `json:"defalutUnitAfter"`
	DefaultUnitRate  float64     `json:"defalutUnitRatio"`
}

type productTemplateGameTemplateItem struct {
	ChargeGame string                              `json:"ChargeGame"`
	GameList   productTemplateGameTemplateGameList `json:"gameList"`
}

type productTemplateGameTemplateGameList struct {
	ChargeRegion []productTemplateGameTemplateGameChargeRegion `json:"ChargeRegion"`
}

type productTemplateGameTemplateGameChargeRegion struct {
	Name         string                                              `json:"name"`
	Code         interface{}                                         `json:"code"`
	ChargeServer []productTemplateGameTemplateGameChargeRegionServer `json:"ChargeServer"`
}

type productTemplateGameTemplateGameChargeRegionServer struct {
	Code       interface{}   `json:"code"`
	Name       string        `json:"name"`
	ChargeType []interface{} `json:"ChargeType"`
}

// GetProductTemplate 获取商品模板
func (c *Client) GetProductTemplate(ctx context.Context, templateID string) (*ProductTemplate, error) {
	var params = &GetProductTemplateParams{TemplateID: templateID}
	var result ProductTemplate

	err := c.Request(ctx, MethodGetProductTemplate, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CheckProductStockParams 校验库存请求参数
type CheckProductStockParams struct {
	BuyNum    int    `json:"buy_num"`
	ProductID string `json:"product_id"`
}

// CheckProductStockResult 校验库存结果
type CheckProductStockResult struct {
	StockStatus string `json:"stock_status"`
	ProductID   int    `json:"product_id"`
}

// CheckProductStock 校验商品库存
func (c *Client) CheckProductStock(ctx context.Context, productID string, num int) (*CheckProductStockResult, error) {
	var params = &CheckProductStockParams{
		ProductID: productID,
		BuyNum:    num,
	}
	var result CheckProductStockResult

	err := c.Request(ctx, MethodCheckProductStock, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
