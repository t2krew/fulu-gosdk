package fulu_gosdk

import "context"

type MobileSpType string

const (
	SpChinaMobile  = MobileSpType("1") // 中国移动
	SpChinaTelecom = MobileSpType("2") // 中国电信
	SpChinaUnicom  = MobileSpType("3") // 中国联通
)

type MobileMaintainState string

const (
	MaintainStateOK    = MobileMaintainState("正常")
	MaintainStateNotOK = MobileMaintainState("维护")
)

type GetQQNicknameResult struct {
	Nickname string `json:"nickname"`
	Photo    string `json:"photo"`
}

// GetQQNickname 获取qq昵称
func (c *Client) GetQQNickname(ctx context.Context, qqNumber string) (*GetQQNicknameResult, error) {
	var (
		result GetQQNicknameResult
		params = map[string]string{
			"qq": qqNumber,
		}
	)
	err := c.Request(ctx, MethodGetQQNickname, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetMobileInfoReqParams struct {
	Phone     string  `json:"phone"`
	FaceValue float64 `json:"face_value,omitempty"`
}

type GetMobileInfoResult struct {
	SP        string    `json:"sp"`         // 运营商名称
	CityCode  string    `json:"city_code"`  // 城市编码
	FaceValue []float64 `json:"face_value"` // 可充值面值
	City      string    `json:"city"`       // 城市名称
	Province  string    `json:"province"`   // 省份名称
	SpType    string    `json:"sp_type"`    // 运营商类型 1:移动 2:电信 3:联通
}

// GetMobileInfo 获取手机归属地
func (c *Client) GetMobileInfo(ctx context.Context, mobileNO string, faceValue ...float64) (*GetMobileInfoResult, error) {
	var (
		result GetMobileInfoResult
		params = GetMobileInfoReqParams{
			Phone: mobileNO,
		}
	)
	if len(faceValue) > 0 {
		params.FaceValue = faceValue[0]
	}
	err := c.Request(ctx, MethodGetMobileInfo, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMobileMaintainStatusReqParams 话费维护状态检查请求参数
type GetMobileMaintainStatusReqParams struct {
	Mobile    string `json:"mobile"`
	FaceValue int    `json:"face_value"`
}

// GetMobileMaintainStatusResult 话费维护状态检查结果
type GetMobileMaintainStatusResult struct {
	Province           string  `json:"province"`
	City               string  `json:"city"`
	Sp                 string  `json:"sp"`
	SpType             string  `json:"sp_type"`
	MaintainState      string  `json:"maintain_state"`
	CurrentSuccessRate float64 `json:"current_success_rate"`
}

// GetMobileMaintainStatus 话费维护状态检查
func (c *Client) GetMobileMaintainStatus(ctx context.Context, mobileNO string, faceValue int) (*GetMobileMaintainStatusResult, error) {
	var (
		result GetMobileMaintainStatusResult
		params = GetMobileMaintainStatusReqParams{
			Mobile:    mobileNO,
			FaceValue: faceValue,
		}
	)
	err := c.Request(ctx, MethodGetMobileMaintainStatus, params, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
