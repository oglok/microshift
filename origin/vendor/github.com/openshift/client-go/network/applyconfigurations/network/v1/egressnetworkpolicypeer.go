// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// EgressNetworkPolicyPeerApplyConfiguration represents an declarative configuration of the EgressNetworkPolicyPeer type for use
// with apply.
type EgressNetworkPolicyPeerApplyConfiguration struct {
	CIDRSelector *string `json:"cidrSelector,omitempty"`
	DNSName      *string `json:"dnsName,omitempty"`
}

// EgressNetworkPolicyPeerApplyConfiguration constructs an declarative configuration of the EgressNetworkPolicyPeer type for use with
// apply.
func EgressNetworkPolicyPeer() *EgressNetworkPolicyPeerApplyConfiguration {
	return &EgressNetworkPolicyPeerApplyConfiguration{}
}

// WithCIDRSelector sets the CIDRSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CIDRSelector field is set to the value of the last call.
func (b *EgressNetworkPolicyPeerApplyConfiguration) WithCIDRSelector(value string) *EgressNetworkPolicyPeerApplyConfiguration {
	b.CIDRSelector = &value
	return b
}

// WithDNSName sets the DNSName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DNSName field is set to the value of the last call.
func (b *EgressNetworkPolicyPeerApplyConfiguration) WithDNSName(value string) *EgressNetworkPolicyPeerApplyConfiguration {
	b.DNSName = &value
	return b
}
