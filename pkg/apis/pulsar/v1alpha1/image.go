package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

// ContainerImage defines the fields needed for a Docker repository image. The
// format here matches the predominant format used in Helm charts.
type ContainerImage struct {
	Repository string            `json:"repository"`
	Tag        string            `json:"tag"`
	PullPolicy corev1.PullPolicy `json:"pullPolicy"`
}

func (c *ContainerImage) SetDefault(cluster *PulsarCluster, podType string) bool {
	changed := false
	switch podType {
	case ZookeeperPodType:
		if cluster.Spec.ZookeeperSpec.Image.Repository == "" {
			cluster.Spec.ZookeeperSpec.Image.Repository = DefaultContainerRepository
			changed = true
		}
		if cluster.Spec.ZookeeperSpec.Image.Tag == "" {
			cluster.Spec.ZookeeperSpec.Image.Tag = DefaultContainerVersion
			changed = true
		}
		if cluster.Spec.ZookeeperSpec.Image.PullPolicy == "" {
			cluster.Spec.ZookeeperSpec.Image.PullPolicy = DefaultContainerPolicy
			changed = true
		}

	case BrokerPodType:
		if cluster.Spec.BrokerSpec.Image.Repository == "" {
			cluster.Spec.BrokerSpec.Image.Repository = DefaultContainerRepository
			changed = true
		}
		if cluster.Spec.BrokerSpec.Image.Tag == "" {
			cluster.Spec.BrokerSpec.Image.Tag = DefaultContainerVersion
			changed = true
		}
		if cluster.Spec.BrokerSpec.Image.PullPolicy == "" {
			cluster.Spec.BrokerSpec.Image.PullPolicy = DefaultContainerPolicy
			changed = true
		}

	case BookiePodType:
		if cluster.Spec.BookieSpec.Image.Repository == "" {
			cluster.Spec.BookieSpec.Image.Repository = DefaultContainerRepository
			changed = true
		}
		if cluster.Spec.BookieSpec.Image.Tag == "" {
			cluster.Spec.BookieSpec.Image.Tag = DefaultContainerVersion
			changed = true
		}
		if cluster.Spec.BookieSpec.Image.PullPolicy == "" {
			cluster.Spec.BookieSpec.Image.PullPolicy = DefaultContainerPolicy
			changed = true
		}
	}
	return changed
}