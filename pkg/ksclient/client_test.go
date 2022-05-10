package ksclient

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestNewKApisClient(t *testing.T) {
	ksClient := NewKApisClient(
		WithBaseTokenPath("http://192.168.100.6:30880", "/login"),
		WithUsername("admin"),
		WithPwd("P@88w0rd"))
	ksClient.RestyClient.OnBeforeRequest(ksClient.TokenBeforeReq)

	//res, err := ksClient.RestyClient.R().Get(ksClient.BaseURL + "/kapis/resources.kubesphere.io/v1alpha3/namespaces/dapr-system/deployments")
	//if err != nil {
	//	t.Log(err)
	//} else {
	//	t.Log("=========1===========\n", string(res.Body()))
	//	t.Log(string(res.Body()))
	//}

	//res, err = ksClient.RestyClient.R().Get(ksClient.BaseURL + "/kapis/resources.kubesphere.io/v1alpha3/namespaces/dapr-system/deployments")
	//if err != nil {
	//	t.Log(err)
	//} else {
	//	t.Log("=========2===========\n", string(res.Body()))
	//	t.Log(string(res.Body()))
	//}
	//q := make(map[string]string)
	//q["ownerKind"] = "ReplicaSet"
	//q["sortBy"] = "startTime"
	//q["labelSelector"] = base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("app=%s", "rudder")))
	//res, err := ksClient.RestyClient.R().SetQueryParams(q).Get(ksClient.BaseURL + "/kapis/resources.kubesphere.io/v1alpha3/namespaces/dapr-system/pods")

	now := time.Now().Unix()
	q := make(map[string]string)
	q["cluster"] = "default"
	q["labelSelector"] = fmt.Sprintf("app=%s", "tkeel-monitor")
	q["resources_filter"] = "tkeel-monitor-7b7c965d9f-cxk4p$"
	q["ownerKind"] = "ReplicaSet"
	q["metrics_filter"] = "pod_cpu_usage|pod_memory_usage_wo_cache$"
	q["start"] = strconv.FormatInt(now-(60*30), 10)
	q["end"] = strconv.FormatInt(now, 10)
	q["step"] = "60s"
	q["times=30"] = "30"
	time.Now().Unix()
	res, err := ksClient.RestyClient.R().SetQueryParams(q).Get(ksClient.BaseURL + "/kapis/monitoring.kubesphere.io/v1alpha3/namespaces/dapr-system/pods")
	t.Log(res.RawResponse.Request.URL)
	if err != nil {
		t.Log(err)
	} else {
		t.Log("=========2===========\n", string(res.Body()))
		t.Log(string(res.Body()))
	}
}
