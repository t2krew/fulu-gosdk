package fulu_gosdk

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Method string

const (
	MethodGetProductList          = Method("fulu.goods.list.get")
	MethodGetProductInfo          = Method("fulu.goods.info.get")
	MethodGetProductTemplate      = Method("fulu.goods.template.get")
	MethodCheckProductStock       = Method("fulu.goods.stock.check")
	MethodGetAccountInfo          = Method("fulu.user.info.get")
	MethodGetQQNickname           = Method("fulu.market.qqnickname.get")
	MethodGetMobileInfo           = Method("fulu.mobile.info.get")
	MethodGetMobileMaintainStatus = Method("fulu.mobile.maintain.check")
	MethodCreateDirectOrder       = Method("fulu.order.direct.add")
	MethodCreateCardOrder         = Method("fulu.order.card.add")
	MethodCreateMobileOrder       = Method("fulu.order.mobile.add")
	MethodQueryOrder              = Method("fulu.order.info.get")
	MethodQueryOrderExtend        = Method("fulu.order.extend.get")
)

const (
	TimestampFormat = "2006-01-02 15:04:05"
)

type ReqParams struct {
	AppKey       string `json:"app_key"`
	Method       Method `json:"method"`
	Timestamp    string `json:"timestamp"`
	Version      string `json:"version"`
	Format       string `json:"format"`
	Charset      string `json:"charset"`
	SignType     string `json:"sign_type"`
	Sign         string `json:"sign"`
	AppAuthToken string `json:"app_auth_token"`
	BizContent   string `json:"biz_content"`
}

type RespData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  string `json:"result"`
	Sign    string `json:"sign"`
}

type Config struct {
	Debug        bool   `json:"debug" yaml:"debug"`
	Endpoint     string `json:"endpoint" yaml:"endpoint"`
	AppKey       string `json:"app_key" yaml:"app_key"`
	AppSecret    string `json:"app_secret" yaml:"app_secret"`
	Format       string `json:"format" yaml:"format"`
	Version      string `json:"version" yaml:"version"`
	Charset      string `json:"charset" yaml:"charset"`
	SignType     string `json:"sign_type" yaml:"sign_type"`
	AppAuthToken string `json:"app_auth_token" yaml:"app_auth_token"`
}

type Client struct {
	cfg     Config
	debug   bool
	appkey  string
	httpCli *resty.Client
}

// New 初始化福禄sdk实例
func New(cfg Config) (*Client, error) {
	return newclient(cfg, resty.New())
}

// NewWithClient 初始化自定义http.Client的福禄sdk实例
func NewWithClient(cfg Config, httpClient *http.Client) (*Client, error) {
	return newclient(cfg, resty.NewWithClient(httpClient))
}

// Request 发起接口请求
func (c *Client) Request(ctx context.Context, method Method, bizContent interface{}, result interface{}) error {
	rawContent, err := jsoniter.MarshalToString(bizContent)
	if err != nil {
		return err
	}
	var params = c.newParams(method, rawContent)

	sign, signStr, err := c.getSign(params)
	if err != nil {
		return err
	}
	if c.debug {
		log.Printf("[fulu-sdk] [%s] sign_str: %s, sign: %s", method, signStr, sign)
	}

	params.Sign = sign

	resp, err := c.httpCli.SetDebug(c.debug).R().SetContext(ctx).SetBody(params).Post(c.cfg.Endpoint)
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("api method [%s] call failed, status code is %d, %s", method, resp.StatusCode(), resp.Status())
	}

	var respdata RespData
	err = jsoniter.Unmarshal(resp.Body(), &respdata)
	if err != nil {
		return err
	}

	switch respdata.Code {
	case 0:
		err := jsoniter.Unmarshal([]byte(respdata.Result), &result)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("errno: %d, errmsg:%s", respdata.Code, respdata.Message)
	}
}

var defaultConfig = Config{
	Debug:        false,
	Format:       "json",
	Version:      "2.0",
	Charset:      "utf-8",
	SignType:     "md5",
	AppAuthToken: "",
}

func newclient(config Config, cli *resty.Client) (*Client, error) {
	var cfg = defaultConfig
	if config.AppKey != "" {
		cfg.AppKey = config.AppKey
	} else {
		return nil, errors.New("appkey is empty")
	}

	if config.AppSecret != "" {
		cfg.AppSecret = config.AppSecret
	} else {
		return nil, errors.New("appsecret is empty")
	}

	if config.Endpoint != "" {
		cfg.Endpoint = config.Endpoint
	} else {
		return nil, errors.New("endpoint is empty")
	}

	cfg.Debug = config.Debug

	if config.Format != "" {
		cfg.Format = config.Format
	}
	if config.Version != "" {
		cfg.Version = config.Version
	}
	if config.Charset != "" {
		cfg.Charset = config.Charset
	}
	if config.SignType != "" {
		cfg.SignType = config.SignType
	}
	if config.AppAuthToken != "" {
		cfg.AppAuthToken = config.AppAuthToken
	}

	return &Client{
		cfg:     cfg,
		debug:   cfg.Debug,
		appkey:  cfg.AppKey,
		httpCli: cli,
	}, nil
}

func (c *Client) getSign(params *ReqParams) (sign string, signStr string, err error) {
	return getSignWithSecret(params, c.cfg.AppSecret)
}

func (c *Client) newParams(method Method, bizContent string) *ReqParams {
	return &ReqParams{
		AppKey:       c.appkey,
		Method:       method,
		Timestamp:    time.Now().Format(TimestampFormat),
		Version:      c.cfg.Version,
		Format:       c.cfg.Format,
		Charset:      c.cfg.Charset,
		SignType:     c.cfg.SignType,
		BizContent:   bizContent,
		AppAuthToken: c.cfg.AppAuthToken,
	}
}

func getSignWithSecret(params *ReqParams, secret string) (sign string, signStr string, err error) {
	var signdata map[string]string

	raw, err := jsoniter.Marshal(params)
	if err != nil {
		return "", "", err
	}

	err = jsoniter.Unmarshal(raw, &signdata)
	if err != nil {
		return "", "", err
	}
	delete(signdata, "sign")

	serializeSigndata, err := jsoniter.MarshalToString(signdata)
	if err != nil {
		return "", "", err
	}
	var (
		runeArray = []rune(serializeSigndata)
		charArray = make([]string, 0, len(runeArray))
	)
	for _, char := range runeArray {
		charArray = append(charArray, string(char))
	}
	sort.Strings(charArray)

	signStr = strings.Join(charArray, "") + secret

	return strings.ToLower(MD5(signStr)), signStr, nil
}

func MD5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}
