package ksclient

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8scli "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
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
	KSNameSpace  string
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

func WithKSNameSpace(nameSpace string) Option {
	if nameSpace == "" {
		nameSpace = "kubesphere-system"
	}
	return func(c *KApisClient) {
		c.KSNameSpace = nameSpace
	}
}
func WithTKeelNS(ns string) Option {
	if ns == "" {
		ns = "keel-system"
	}
	return func(c *KApisClient) {
		c.TKNameSpace = ns
	}
}
func WithTKeelCluster(cluster string) Option {
	if cluster == "" {
		cluster = "default"
	}
	return func(c *KApisClient) {
		c.TKCluster = cluster
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

func (c *KApisClient) SecretTokenBeforeReq(client *resty.Client, request *resty.Request) error {
	if c.Token != "" {
		request.SetCookie(&http.Cookie{
			Name:  "token",
			Value: c.Token,
		})
		return nil
	}
	conf, err := rest.InClusterConfig()
	if err != nil {
		return errors.Wrap(err, "ks client SecretTokenBeforeReq cliConf")
	}
	cliSet, err := k8scli.NewForConfig(conf)
	if err != nil {
		return errors.Wrap(err, "ks client SecretTokenBeforeReq k8sSliSet")
	}
	ctx := context.TODO()
	s, err := cliSet.CoreV1().Secrets(c.KSNameSpace).Get(ctx, "kubesphere-secret", metav1.GetOptions{})
	if err != nil {
		return errors.Wrap(err, "ks client SecretTokenBeforeReq get secret")
	}
	nodes, err := cliSet.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return errors.Wrap(err, "ks nodes list")
	}
	for i := range nodes.Items {
		if _, ok := nodes.Items[i].Labels["node-role.kubernetes.io/master"]; ok {
			for addri := range nodes.Items[i].Status.Addresses {
				if nodes.Items[i].Status.Addresses[addri].Type == "InternalIP" {
					mip := nodes.Items[i].Status.Addresses[addri].Address
					c.BaseURL = "http://" + mip + ":30880"
					break
				}
			}
			break
		}
	}
	tokenBytes := s.Data["token"]
	c.Token = string(tokenBytes)
	request.SetCookie(&http.Cookie{
		Name:  "token",
		Value: c.Token,
	})
	return nil

}
