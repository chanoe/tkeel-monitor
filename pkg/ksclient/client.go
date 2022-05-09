package ksclient

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

type KApisClient struct {
	RestyClient  *resty.Client
	BaseURL      string
	TokenPath    string
	Token        string
	Expire       int64
	RefreshToken string
	TKNameSpace  string
	TKCluster    string
	username     string
	pwd          string
}

type Option func(client *KApisClient)

func NewKApisClient(opts ...Option) *KApisClient {
	c := &KApisClient{RestyClient: resty.New()}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func WithBaseTokenPath(base, path string) Option {
	return func(c *KApisClient) {
		c.BaseURL = base
		if path == "" {
			path = "/login"
		}
		c.TokenPath = path
	}
}

func WithUsername(username string) Option {
	return func(c *KApisClient) {
		c.username = username
	}
}
func WithTkeelNS(ns string) Option {
	if ns == "" {
		ns = "keel-system"
	}
	return func(c *KApisClient) {
		c.TKNameSpace = ns
	}
}
func WithTkeelCluster(cluster string) Option {
	if cluster == "" {
		cluster = "default"
	}
	return func(c *KApisClient) {
		c.TKCluster = cluster
	}
}
func WithPwd(pwd string) Option {
	return func(c *KApisClient) {
		c.pwd = ksConsoleEncrypt("", pwd)
	}
}

func ksConsoleEncrypt(salt, str string) string {
	if salt == "" {
		salt = "kubesphere"
	}
	str = base64.StdEncoding.EncodeToString([]byte(str))
	if len(salt) < len(str) {
		salt = salt + str[0:len(str)-len(salt)]
	}
	ret := make([]string, 0, len(salt))
	prefix := []string{}
	for i := 0; i < len(salt); i++ {
		tomix := int32(64)
		if len(str) > i {
			tomix = []rune(str)[i]
		}
		sum := int32(0)
		sum = []rune(salt)[i] + tomix

		if sum%2 == 0 {
			prefix = append(prefix, "0")
		} else {
			prefix = append(prefix, "1")
		}
		ret = append(ret, string(rune(sum/2)))
	}
	return base64.StdEncoding.EncodeToString([]byte(strings.Join(prefix, ""))) + "@" + strings.Join(ret, "")
}

func (c *KApisClient) TokenBeforeReq(client *resty.Client, request *resty.Request) error {
	now := time.Now().UnixMilli() + 3000
	if c.Expire > now {
		request.SetCookie(&http.Cookie{
			Name:  "token",
			Value: c.Token,
		})
		return nil
	}
	c.Expire = now
	cookieJar, _ := cookiejar.New(&cookiejar.Options{})
	res, err := resty.New().SetRedirectPolicy(resty.FlexibleRedirectPolicy(15)).SetCookieJar(cookieJar).R().
		SetFormData(map[string]string{"username": c.username, "encrypt": c.pwd}).
		Post(c.BaseURL + c.TokenPath)
	if err != nil {
		return errors.Wrap(err, "ks login")
	}
	resMap := make(map[string]interface{})
	json.Unmarshal(res.Body(), &resMap)
	if resMap["status"] != nil {
		return errors.New(resMap["message"].(string))
	}
	for _, cookie := range cookieJar.Cookies(res.RawResponse.Request.URL) {
		switch cookie.Name {
		case "token":
			c.Token = cookie.Value
		case "expire":
			c.Expire, err = strconv.ParseInt(cookie.Value, 10, 64)
			if err != nil {
				return errors.Wrapf(err, "ks login parse expire")
			}
		case "refreshToken":
			c.RefreshToken = cookie.Value
		default:
		}
	}
	request.SetCookie(&http.Cookie{
		Name:  "token",
		Value: c.Token,
	})
	return nil
}
