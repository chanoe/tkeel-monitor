package v1

import (
	"encoding/base64"
	go_restful "github.com/emicklei/go-restful"
	"strings"
)

func RegisterKSMetricsHTTPServer(container *go_restful.Container) {
	var ws *go_restful.WebService
	for _, v := range container.RegisteredWebServices() {
		if v.RootPath() == "/v1" {
			ws = v
			break
		}
	}
	if ws == nil {
		ws = new(go_restful.WebService)
		ws.ApiVersion("/v1")
		ws.Path("/v1").Produces(go_restful.MIME_JSON)
		container.Add(ws)
	}

	ws.Route(ws.GET("/ksmetrics"))
}

func EncryptPwd(salt, str string) string {
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
