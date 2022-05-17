package v1

import "time"

type DeploymentPlugins struct {
	Items      []PluginItems `json:"items"`
	TotalItems int           `json:"totalItems"`
}
type PluginMetadata struct {
	Name string `json:"name"`
	UID  string `json:"uid"`
}
type Conditions struct {
	Type               string    `json:"type"`
	Status             string    `json:"status"`
	LastUpdateTime     time.Time `json:"lastUpdateTime"`
	LastTransitionTime time.Time `json:"lastTransitionTime"`
	Reason             string    `json:"reason"`
	Message            string    `json:"message"`
}
type PluginStatus struct {
	Replicas          int          `json:"replicas"`
	UpdatedReplicas   int          `json:"updatedReplicas"`
	ReadyReplicas     int          `json:"readyReplicas"`
	AvailableReplicas int          `json:"availableReplicas"`
	Conditions        []Conditions `json:"conditions"`
	Status            string       `json:"status"`
	UpdateTime        int64        `json:"updateTime"`
}
type PluginItems struct {
	Metadata PluginMetadata `json:"metadata"`
	Status   PluginStatus   `json:"status"`
}

type PluginPods struct {
	Items      []PodResult `json:"items"`
	TotalItems int         `json:"totalItems"`
}

type PodResult struct {
	Metadata PodsMetadata `json:"metadata"`
	Status   PodsStatus   `json:"status"`
}
type PodsMetadata struct {
	Name string `json:"name"`
	UID  string `json:"uid"`
}
type PodsStatus struct {
	Phase     string    `json:"phase"`
	HostIP    string    `json:"hostIP"`
	PodIP     string    `json:"podIP"`
	StartTime time.Time `json:"startTime"`
}
