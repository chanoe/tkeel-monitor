package prometheus

import (
	"strings"

	"github.com/pkg/errors"
)

var MetricsMap = map[string]string{
	"sum_tkapi_request_7d":         "sum(increase(tkapi_request_total{$}[7d]))",
	"sum_tkapi_request_24h":        "sum(increase(tkapi_request_total{$}[24h]))",
	"avg_tkapi_request_latency_7d": "sum(increase(tkapi_request_duration_seconds_sum{$}[7d])) /sum(increase(tkapi_request_duration_seconds_count{$}[7d]))",

	"p90_tkapi_request_latency": "histogram_quantile(0.90, sum by (le) (rate(tkapi_request_duration_seconds_bucket{$}[1d])))",
	"p95_tkapi_request_latency": "histogram_quantile(0.95, sum by (le) (rate(tkapi_request_duration_seconds_bucket{$}[1d])))",
	"p99_tkapi_request_latency": "histogram_quantile(0.99, sum by (le) (rate(tkapi_request_duration_seconds_bucket{$}[1d])))",

	"upstream_msg":   "sum(iothub_msg_total{direction='upstream',$})",   //上行消息数量
	"downstream_msg": "sum(iothub_msg_total{direction='downstream',$})", //下行消息数量

	"subscribe_num":          "(sum(subscribe_num{$})) / (count (sum by (pod) (subscribe_num)))",                                                                        // 订阅数
	"subscribe_entities_num": "sum(subscribe_entities_num{$}) / (count (sum by (pod) (subscribe_entities_num)))",                                                        // 订阅的实体数量
	"rule_num":               "sum(rule_num{$})  / (count (sum by (pod) (rule_num)))",                                                                                   // 路由数量
	"rule_execute_num_24h":   "sum(increase(rule_execute_num{$}[24h])) / (count (sum by (pod) (rule_num)))",                                                             //24小时 规则执行数
	"rate_rule_failure_24h":  "ceil(( (sum( increase(rule_execute_num{status='failure',$}[24h])))  /  ((sum( increase(rule_execute_num{$}[24h]))) == bool 0 ) ) * 100)", // 24小时 规则执行数量

	"sum_device_num":   "sum(device_num_total{$}) / (count (sum by (pod) (device_num_total)))",                         // 设备数量
	"sum_template_num": "sum(device_template_total{$}) / (count (sum by (pod) (device_template_total)))",               //模板数量
	"rate_online":      "ceil((sum(iothub_connected_total{tenant_id='_tKeel_system'}) / sum(device_num_total)) * 100)", // 在线率
}

func ExpressFromMetricsMap(name, label string) (string, error) {
	if v, ok := MetricsMap[name]; ok {
		expr := strings.Replace(v, "$", label, -1)
		return expr, nil
	}
	return "", errors.New("metrics name not existed")
}
