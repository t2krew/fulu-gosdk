package fulu_gosdk

import "context"

type AccountInfo struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
	IsOpen  int     `json:"is_open"`
}

// GetAccountInfo 获取用户信息
func (c *Client) GetAccountInfo(ctx context.Context) (*AccountInfo, error) {
	var result AccountInfo
	err := c.Request(ctx, MethodGetAccountInfo, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
