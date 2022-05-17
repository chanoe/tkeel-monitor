package service

import (
	"fmt"

	"github.com/tkeel-io/kit/log"
	"github.com/tkeel-io/tkeel-monitor/pkg/ksclient"
)

const (
	KSPATH_V1ALPHA3_DEPLOYMENT   = "/kapis/resources.kubesphere.io/v1alpha3/namespaces/%s/deployments"
	KSPATH_V1ALPHA3_STATEFULSETS = "/kapis/resources.kubesphere.io/v1alpha3/namespaces/%s/statefulsets"
	KSPATH_V1ALPHA3_PODS         = "/kapis/resources.kubesphere.io/v1alpha3/namespaces/%s/pods"
	KSPATH_V1ALPHA3_PODS_MONITOR = "/kapis/monitoring.kubesphere.io/v1alpha3/namespaces/%s/pods"
	KSPATH_LOGIN                 = "/login"
)

type KsMetricsService struct {
	kapiCli *ksclient.KApisClient
}

func NewKsMetricsService(c *ksclient.KApisClient) KsMetricsService {
	return KsMetricsService{c}
}

func (svc KsMetricsService) MetricsPlugins(name, status, sortBy, ascending, limit, page string) ([]byte, error) {
	R := svc.kapiCli.RestyClient.R()
	R.SetQueryParams(parseDeploymentQuery(name, status, sortBy, ascending, limit, page))
	res, err := R.Get(svc.kapiCli.BaseURL + fmt.Sprintf(KSPATH_V1ALPHA3_DEPLOYMENT, svc.kapiCli.TKNameSpace))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return res.Body(), nil
}

func (svc KsMetricsService) PluginPods(plugin string) ([]byte, error) {
	res, err := svc.kapiCli.RestyClient.R().
		SetQueryParams(parsePodsQuery(plugin)).
		Get(svc.kapiCli.BaseURL + fmt.Sprintf(KSPATH_V1ALPHA3_PODS, svc.kapiCli.TKNameSpace))
	log.Debug(res.Request.URL)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return res.Body(), nil
}

func (svc KsMetricsService) PodsCpuMem(plugin, resources, start, end, step, times string) ([]byte, error) {
	res, err := svc.kapiCli.RestyClient.R().
		SetQueryParams(parsePodsMonitorQuery(svc.kapiCli.TKCluster, plugin, resources, start, end, step, times)).
		Get(svc.kapiCli.BaseURL + fmt.Sprintf(KSPATH_V1ALPHA3_PODS_MONITOR, svc.kapiCli.TKNameSpace))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return res.Body(), nil
}

func parseDeploymentQuery(name, status, sortBy, ascending, limit, page string) map[string]string {
	q := make(map[string]string)
	if name != "" {
		q["name"] = name
	}
	if status == "running" || status == "updating" || status == "stopped" {
		q["status"] = status
	}
	if limit != "" {
		q["limit"] = limit
	}
	if page != "" {
		q["page"] = page
	}
	if sortBy != "" {
		q["sortBy"] = sortBy
	}
	if ascending != "" {
		q["ascending"] = ascending
	}
	return q
}

func parsePodsQuery(plugin string) map[string]string {
	q := make(map[string]string)
	q["ownerKind"] = "ReplicaSet"
	q["sortBy"] = "startTime"
	labelSelector := ""
	if plugin == "tkeel-device" {
		labelSelector = fmt.Sprintf("app.kubernetes.io/instance=%s,app.kubernetes.io/name=%s", plugin, plugin)
	} else {
		labelSelector = fmt.Sprintf("app=%s", plugin)
	}
	q["labelSelector"] = labelSelector
	return q
}

func parsePodsMonitorQuery(cluster, plugin, resources, start, end, step, times string) map[string]string {
	q := make(map[string]string)
	q["cluster"] = cluster
	q["labelSelector"] = fmt.Sprintf("app=%s", plugin)
	q["resources_filter"] = resources
	q["ownerKind"] = "ReplicaSet"
	if start != "" {
		q["start"] = start
		q["end"] = end
		q["step"] = step
		q["times"] = times
	}
	q["metrics_filter"] = "pod_cpu_usage|pod_memory_usage_wo_cache$"
	return q
}
