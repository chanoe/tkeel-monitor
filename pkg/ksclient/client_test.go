package ksclient

import (
	"testing"
)

func TestNewKApisClient(t *testing.T) {
	ksClient := NewKApisClient(
		WithBaseTokenPath("http://192.168.100.6:30880", "/login"),
		WithUsername("admin"),
		WithPwd("P@88w0rd"))
	ksClient.RestyClient.OnBeforeRequest(ksClient.TokenBeforeReq)

	res, err := ksClient.RestyClient.R().Get(ksClient.BaseURL + "/kapis/resources.kubesphere.io/v1alpha3/namespaces/dapr-system/deployments")
	if err != nil {
		t.Log(err)
	} else {
		t.Log("=========1===========\n", string(res.Body()))
		t.Log(string(res.Body()))
	}

	res, err = ksClient.RestyClient.R().Get(ksClient.BaseURL + "/kapis/resources.kubesphere.io/v1alpha3/namespaces/dapr-system/deployments")
	if err != nil {
		t.Log(err)
	} else {
		t.Log("=========2===========\n", string(res.Body()))
		t.Log(string(res.Body()))
	}
}
