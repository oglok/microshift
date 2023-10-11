// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	v1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// ImageRegistrySpecApplyConfiguration represents an declarative configuration of the ImageRegistrySpec type for use
// with apply.
type ImageRegistrySpecApplyConfiguration struct {
	v1.OperatorSpecApplyConfiguration `json:",inline"`
	HTTPSecret                        *string                                        `json:"httpSecret,omitempty"`
	Proxy                             *ImageRegistryConfigProxyApplyConfiguration    `json:"proxy,omitempty"`
	Storage                           *ImageRegistryConfigStorageApplyConfiguration  `json:"storage,omitempty"`
	ReadOnly                          *bool                                          `json:"readOnly,omitempty"`
	DisableRedirect                   *bool                                          `json:"disableRedirect,omitempty"`
	Requests                          *ImageRegistryConfigRequestsApplyConfiguration `json:"requests,omitempty"`
	DefaultRoute                      *bool                                          `json:"defaultRoute,omitempty"`
	Routes                            []ImageRegistryConfigRouteApplyConfiguration   `json:"routes,omitempty"`
	Replicas                          *int32                                         `json:"replicas,omitempty"`
	Logging                           *int64                                         `json:"logging,omitempty"`
	Resources                         *corev1.ResourceRequirements                   `json:"resources,omitempty"`
	NodeSelector                      map[string]string                              `json:"nodeSelector,omitempty"`
	Tolerations                       []corev1.Toleration                            `json:"tolerations,omitempty"`
	RolloutStrategy                   *string                                        `json:"rolloutStrategy,omitempty"`
	Affinity                          *corev1.Affinity                               `json:"affinity,omitempty"`
	TopologySpreadConstraints         []corev1.TopologySpreadConstraint              `json:"topologySpreadConstraints,omitempty"`
}

// ImageRegistrySpecApplyConfiguration constructs an declarative configuration of the ImageRegistrySpec type for use with
// apply.
func ImageRegistrySpec() *ImageRegistrySpecApplyConfiguration {
	return &ImageRegistrySpecApplyConfiguration{}
}

// WithManagementState sets the ManagementState field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ManagementState field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithManagementState(value operatorv1.ManagementState) *ImageRegistrySpecApplyConfiguration {
	b.ManagementState = &value
	return b
}

// WithLogLevel sets the LogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogLevel field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithLogLevel(value operatorv1.LogLevel) *ImageRegistrySpecApplyConfiguration {
	b.LogLevel = &value
	return b
}

// WithOperatorLogLevel sets the OperatorLogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OperatorLogLevel field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithOperatorLogLevel(value operatorv1.LogLevel) *ImageRegistrySpecApplyConfiguration {
	b.OperatorLogLevel = &value
	return b
}

// WithUnsupportedConfigOverrides sets the UnsupportedConfigOverrides field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UnsupportedConfigOverrides field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithUnsupportedConfigOverrides(value runtime.RawExtension) *ImageRegistrySpecApplyConfiguration {
	b.UnsupportedConfigOverrides = &value
	return b
}

// WithObservedConfig sets the ObservedConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedConfig field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithObservedConfig(value runtime.RawExtension) *ImageRegistrySpecApplyConfiguration {
	b.ObservedConfig = &value
	return b
}

// WithHTTPSecret sets the HTTPSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPSecret field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithHTTPSecret(value string) *ImageRegistrySpecApplyConfiguration {
	b.HTTPSecret = &value
	return b
}

// WithProxy sets the Proxy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Proxy field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithProxy(value *ImageRegistryConfigProxyApplyConfiguration) *ImageRegistrySpecApplyConfiguration {
	b.Proxy = value
	return b
}

// WithStorage sets the Storage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Storage field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithStorage(value *ImageRegistryConfigStorageApplyConfiguration) *ImageRegistrySpecApplyConfiguration {
	b.Storage = value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithReadOnly(value bool) *ImageRegistrySpecApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// WithDisableRedirect sets the DisableRedirect field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisableRedirect field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithDisableRedirect(value bool) *ImageRegistrySpecApplyConfiguration {
	b.DisableRedirect = &value
	return b
}

// WithRequests sets the Requests field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Requests field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithRequests(value *ImageRegistryConfigRequestsApplyConfiguration) *ImageRegistrySpecApplyConfiguration {
	b.Requests = value
	return b
}

// WithDefaultRoute sets the DefaultRoute field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DefaultRoute field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithDefaultRoute(value bool) *ImageRegistrySpecApplyConfiguration {
	b.DefaultRoute = &value
	return b
}

// WithRoutes adds the given value to the Routes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Routes field.
func (b *ImageRegistrySpecApplyConfiguration) WithRoutes(values ...*ImageRegistryConfigRouteApplyConfiguration) *ImageRegistrySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRoutes")
		}
		b.Routes = append(b.Routes, *values[i])
	}
	return b
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithReplicas(value int32) *ImageRegistrySpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithLogging sets the Logging field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Logging field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithLogging(value int64) *ImageRegistrySpecApplyConfiguration {
	b.Logging = &value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithResources(value corev1.ResourceRequirements) *ImageRegistrySpecApplyConfiguration {
	b.Resources = &value
	return b
}

// WithNodeSelector puts the entries into the NodeSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NodeSelector field,
// overwriting an existing map entries in NodeSelector field with the same key.
func (b *ImageRegistrySpecApplyConfiguration) WithNodeSelector(entries map[string]string) *ImageRegistrySpecApplyConfiguration {
	if b.NodeSelector == nil && len(entries) > 0 {
		b.NodeSelector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.NodeSelector[k] = v
	}
	return b
}

// WithTolerations adds the given value to the Tolerations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tolerations field.
func (b *ImageRegistrySpecApplyConfiguration) WithTolerations(values ...corev1.Toleration) *ImageRegistrySpecApplyConfiguration {
	for i := range values {
		b.Tolerations = append(b.Tolerations, values[i])
	}
	return b
}

// WithRolloutStrategy sets the RolloutStrategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RolloutStrategy field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithRolloutStrategy(value string) *ImageRegistrySpecApplyConfiguration {
	b.RolloutStrategy = &value
	return b
}

// WithAffinity sets the Affinity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Affinity field is set to the value of the last call.
func (b *ImageRegistrySpecApplyConfiguration) WithAffinity(value corev1.Affinity) *ImageRegistrySpecApplyConfiguration {
	b.Affinity = &value
	return b
}

// WithTopologySpreadConstraints adds the given value to the TopologySpreadConstraints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TopologySpreadConstraints field.
func (b *ImageRegistrySpecApplyConfiguration) WithTopologySpreadConstraints(values ...corev1.TopologySpreadConstraint) *ImageRegistrySpecApplyConfiguration {
	for i := range values {
		b.TopologySpreadConstraints = append(b.TopologySpreadConstraints, values[i])
	}
	return b
}
