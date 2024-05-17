/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the NetworkConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkConfiguration{}

// NetworkConfiguration Defines the MediaBrowser.Common.Net.NetworkConfiguration.
type NetworkConfiguration struct {
	// Gets or sets a value used to specify the URL prefix that your Jellyfin instance can be accessed at.
	BaseUrl *string `json:"BaseUrl,omitempty"`
	// Gets or sets a value indicating whether to use HTTPS.
	EnableHttps *bool `json:"EnableHttps,omitempty"`
	// Gets or sets a value indicating whether the server should force connections over HTTPS.
	RequireHttps *bool `json:"RequireHttps,omitempty"`
	// Gets or sets the filesystem path of an X.509 certificate to use for SSL.
	CertificatePath *string `json:"CertificatePath,omitempty"`
	// Gets or sets the password required to access the X.509 certificate data in the file specified by MediaBrowser.Common.Net.NetworkConfiguration.CertificatePath.
	CertificatePassword *string `json:"CertificatePassword,omitempty"`
	// Gets or sets the internal HTTP server port.
	InternalHttpPort *int32 `json:"InternalHttpPort,omitempty"`
	// Gets or sets the internal HTTPS server port.
	InternalHttpsPort *int32 `json:"InternalHttpsPort,omitempty"`
	// Gets or sets the public HTTP port.
	PublicHttpPort *int32 `json:"PublicHttpPort,omitempty"`
	// Gets or sets the public HTTPS port.
	PublicHttpsPort *int32 `json:"PublicHttpsPort,omitempty"`
	// Gets or sets a value indicating whether Autodiscovery is enabled.
	AutoDiscovery *bool `json:"AutoDiscovery,omitempty"`
	// Gets or sets a value indicating whether to enable automatic port forwarding.
	EnableUPnP *bool `json:"EnableUPnP,omitempty"`
	// Gets or sets a value indicating whether IPv6 is enabled.
	EnableIPv4 *bool `json:"EnableIPv4,omitempty"`
	// Gets or sets a value indicating whether IPv6 is enabled.
	EnableIPv6 *bool `json:"EnableIPv6,omitempty"`
	// Gets or sets a value indicating whether access from outside of the LAN is permitted.
	EnableRemoteAccess *bool `json:"EnableRemoteAccess,omitempty"`
	// Gets or sets the subnets that are deemed to make up the LAN.
	LocalNetworkSubnets []string `json:"LocalNetworkSubnets,omitempty"`
	// Gets or sets the interface addresses which Jellyfin will bind to. If empty, all interfaces will be used.
	LocalNetworkAddresses []string `json:"LocalNetworkAddresses,omitempty"`
	// Gets or sets the known proxies.
	KnownProxies []string `json:"KnownProxies,omitempty"`
	// Gets or sets a value indicating whether address names that match MediaBrowser.Common.Net.NetworkConfiguration.VirtualInterfaceNames should be ignored for the purposes of binding.
	IgnoreVirtualInterfaces *bool `json:"IgnoreVirtualInterfaces,omitempty"`
	// Gets or sets a value indicating the interface name prefixes that should be ignored. The list can be comma separated and values are case-insensitive. <seealso cref=\"P:MediaBrowser.Common.Net.NetworkConfiguration.IgnoreVirtualInterfaces\" />.
	VirtualInterfaceNames []string `json:"VirtualInterfaceNames,omitempty"`
	// Gets or sets a value indicating whether the published server uri is based on information in HTTP requests.
	EnablePublishedServerUriByRequest *bool `json:"EnablePublishedServerUriByRequest,omitempty"`
	// Gets or sets the PublishedServerUriBySubnet  Gets or sets PublishedServerUri to advertise for specific subnets.
	PublishedServerUriBySubnet []string `json:"PublishedServerUriBySubnet,omitempty"`
	// Gets or sets the filter for remote IP connectivity. Used in conjunction with <seealso cref=\"P:MediaBrowser.Common.Net.NetworkConfiguration.IsRemoteIPFilterBlacklist\" />.
	RemoteIPFilter []string `json:"RemoteIPFilter,omitempty"`
	// Gets or sets a value indicating whether <seealso cref=\"P:MediaBrowser.Common.Net.NetworkConfiguration.RemoteIPFilter\" /> contains a blacklist or a whitelist. Default is a whitelist.
	IsRemoteIPFilterBlacklist *bool `json:"IsRemoteIPFilterBlacklist,omitempty"`
}

// NewNetworkConfiguration instantiates a new NetworkConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkConfiguration() *NetworkConfiguration {
	this := NetworkConfiguration{}
	return &this
}

// NewNetworkConfigurationWithDefaults instantiates a new NetworkConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkConfigurationWithDefaults() *NetworkConfiguration {
	this := NetworkConfiguration{}
	return &this
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasBaseUrl() bool {
	if o != nil && !IsNil(o.BaseUrl) {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *NetworkConfiguration) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetEnableHttps returns the EnableHttps field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetEnableHttps() bool {
	if o == nil || IsNil(o.EnableHttps) {
		var ret bool
		return ret
	}
	return *o.EnableHttps
}

// GetEnableHttpsOk returns a tuple with the EnableHttps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetEnableHttpsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableHttps) {
		return nil, false
	}
	return o.EnableHttps, true
}

// HasEnableHttps returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasEnableHttps() bool {
	if o != nil && !IsNil(o.EnableHttps) {
		return true
	}

	return false
}

// SetEnableHttps gets a reference to the given bool and assigns it to the EnableHttps field.
func (o *NetworkConfiguration) SetEnableHttps(v bool) {
	o.EnableHttps = &v
}

// GetRequireHttps returns the RequireHttps field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetRequireHttps() bool {
	if o == nil || IsNil(o.RequireHttps) {
		var ret bool
		return ret
	}
	return *o.RequireHttps
}

// GetRequireHttpsOk returns a tuple with the RequireHttps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetRequireHttpsOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireHttps) {
		return nil, false
	}
	return o.RequireHttps, true
}

// HasRequireHttps returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasRequireHttps() bool {
	if o != nil && !IsNil(o.RequireHttps) {
		return true
	}

	return false
}

// SetRequireHttps gets a reference to the given bool and assigns it to the RequireHttps field.
func (o *NetworkConfiguration) SetRequireHttps(v bool) {
	o.RequireHttps = &v
}

// GetCertificatePath returns the CertificatePath field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetCertificatePath() string {
	if o == nil || IsNil(o.CertificatePath) {
		var ret string
		return ret
	}
	return *o.CertificatePath
}

// GetCertificatePathOk returns a tuple with the CertificatePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetCertificatePathOk() (*string, bool) {
	if o == nil || IsNil(o.CertificatePath) {
		return nil, false
	}
	return o.CertificatePath, true
}

// HasCertificatePath returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasCertificatePath() bool {
	if o != nil && !IsNil(o.CertificatePath) {
		return true
	}

	return false
}

// SetCertificatePath gets a reference to the given string and assigns it to the CertificatePath field.
func (o *NetworkConfiguration) SetCertificatePath(v string) {
	o.CertificatePath = &v
}

// GetCertificatePassword returns the CertificatePassword field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetCertificatePassword() string {
	if o == nil || IsNil(o.CertificatePassword) {
		var ret string
		return ret
	}
	return *o.CertificatePassword
}

// GetCertificatePasswordOk returns a tuple with the CertificatePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetCertificatePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.CertificatePassword) {
		return nil, false
	}
	return o.CertificatePassword, true
}

// HasCertificatePassword returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasCertificatePassword() bool {
	if o != nil && !IsNil(o.CertificatePassword) {
		return true
	}

	return false
}

// SetCertificatePassword gets a reference to the given string and assigns it to the CertificatePassword field.
func (o *NetworkConfiguration) SetCertificatePassword(v string) {
	o.CertificatePassword = &v
}

// GetInternalHttpPort returns the InternalHttpPort field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetInternalHttpPort() int32 {
	if o == nil || IsNil(o.InternalHttpPort) {
		var ret int32
		return ret
	}
	return *o.InternalHttpPort
}

// GetInternalHttpPortOk returns a tuple with the InternalHttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetInternalHttpPortOk() (*int32, bool) {
	if o == nil || IsNil(o.InternalHttpPort) {
		return nil, false
	}
	return o.InternalHttpPort, true
}

// HasInternalHttpPort returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasInternalHttpPort() bool {
	if o != nil && !IsNil(o.InternalHttpPort) {
		return true
	}

	return false
}

// SetInternalHttpPort gets a reference to the given int32 and assigns it to the InternalHttpPort field.
func (o *NetworkConfiguration) SetInternalHttpPort(v int32) {
	o.InternalHttpPort = &v
}

// GetInternalHttpsPort returns the InternalHttpsPort field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetInternalHttpsPort() int32 {
	if o == nil || IsNil(o.InternalHttpsPort) {
		var ret int32
		return ret
	}
	return *o.InternalHttpsPort
}

// GetInternalHttpsPortOk returns a tuple with the InternalHttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetInternalHttpsPortOk() (*int32, bool) {
	if o == nil || IsNil(o.InternalHttpsPort) {
		return nil, false
	}
	return o.InternalHttpsPort, true
}

// HasInternalHttpsPort returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasInternalHttpsPort() bool {
	if o != nil && !IsNil(o.InternalHttpsPort) {
		return true
	}

	return false
}

// SetInternalHttpsPort gets a reference to the given int32 and assigns it to the InternalHttpsPort field.
func (o *NetworkConfiguration) SetInternalHttpsPort(v int32) {
	o.InternalHttpsPort = &v
}

// GetPublicHttpPort returns the PublicHttpPort field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetPublicHttpPort() int32 {
	if o == nil || IsNil(o.PublicHttpPort) {
		var ret int32
		return ret
	}
	return *o.PublicHttpPort
}

// GetPublicHttpPortOk returns a tuple with the PublicHttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetPublicHttpPortOk() (*int32, bool) {
	if o == nil || IsNil(o.PublicHttpPort) {
		return nil, false
	}
	return o.PublicHttpPort, true
}

// HasPublicHttpPort returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasPublicHttpPort() bool {
	if o != nil && !IsNil(o.PublicHttpPort) {
		return true
	}

	return false
}

// SetPublicHttpPort gets a reference to the given int32 and assigns it to the PublicHttpPort field.
func (o *NetworkConfiguration) SetPublicHttpPort(v int32) {
	o.PublicHttpPort = &v
}

// GetPublicHttpsPort returns the PublicHttpsPort field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetPublicHttpsPort() int32 {
	if o == nil || IsNil(o.PublicHttpsPort) {
		var ret int32
		return ret
	}
	return *o.PublicHttpsPort
}

// GetPublicHttpsPortOk returns a tuple with the PublicHttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetPublicHttpsPortOk() (*int32, bool) {
	if o == nil || IsNil(o.PublicHttpsPort) {
		return nil, false
	}
	return o.PublicHttpsPort, true
}

// HasPublicHttpsPort returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasPublicHttpsPort() bool {
	if o != nil && !IsNil(o.PublicHttpsPort) {
		return true
	}

	return false
}

// SetPublicHttpsPort gets a reference to the given int32 and assigns it to the PublicHttpsPort field.
func (o *NetworkConfiguration) SetPublicHttpsPort(v int32) {
	o.PublicHttpsPort = &v
}

// GetAutoDiscovery returns the AutoDiscovery field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetAutoDiscovery() bool {
	if o == nil || IsNil(o.AutoDiscovery) {
		var ret bool
		return ret
	}
	return *o.AutoDiscovery
}

// GetAutoDiscoveryOk returns a tuple with the AutoDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetAutoDiscoveryOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoDiscovery) {
		return nil, false
	}
	return o.AutoDiscovery, true
}

// HasAutoDiscovery returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasAutoDiscovery() bool {
	if o != nil && !IsNil(o.AutoDiscovery) {
		return true
	}

	return false
}

// SetAutoDiscovery gets a reference to the given bool and assigns it to the AutoDiscovery field.
func (o *NetworkConfiguration) SetAutoDiscovery(v bool) {
	o.AutoDiscovery = &v
}

// GetEnableUPnP returns the EnableUPnP field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetEnableUPnP() bool {
	if o == nil || IsNil(o.EnableUPnP) {
		var ret bool
		return ret
	}
	return *o.EnableUPnP
}

// GetEnableUPnPOk returns a tuple with the EnableUPnP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetEnableUPnPOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableUPnP) {
		return nil, false
	}
	return o.EnableUPnP, true
}

// HasEnableUPnP returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasEnableUPnP() bool {
	if o != nil && !IsNil(o.EnableUPnP) {
		return true
	}

	return false
}

// SetEnableUPnP gets a reference to the given bool and assigns it to the EnableUPnP field.
func (o *NetworkConfiguration) SetEnableUPnP(v bool) {
	o.EnableUPnP = &v
}

// GetEnableIPv4 returns the EnableIPv4 field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetEnableIPv4() bool {
	if o == nil || IsNil(o.EnableIPv4) {
		var ret bool
		return ret
	}
	return *o.EnableIPv4
}

// GetEnableIPv4Ok returns a tuple with the EnableIPv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetEnableIPv4Ok() (*bool, bool) {
	if o == nil || IsNil(o.EnableIPv4) {
		return nil, false
	}
	return o.EnableIPv4, true
}

// HasEnableIPv4 returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasEnableIPv4() bool {
	if o != nil && !IsNil(o.EnableIPv4) {
		return true
	}

	return false
}

// SetEnableIPv4 gets a reference to the given bool and assigns it to the EnableIPv4 field.
func (o *NetworkConfiguration) SetEnableIPv4(v bool) {
	o.EnableIPv4 = &v
}

// GetEnableIPv6 returns the EnableIPv6 field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetEnableIPv6() bool {
	if o == nil || IsNil(o.EnableIPv6) {
		var ret bool
		return ret
	}
	return *o.EnableIPv6
}

// GetEnableIPv6Ok returns a tuple with the EnableIPv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetEnableIPv6Ok() (*bool, bool) {
	if o == nil || IsNil(o.EnableIPv6) {
		return nil, false
	}
	return o.EnableIPv6, true
}

// HasEnableIPv6 returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasEnableIPv6() bool {
	if o != nil && !IsNil(o.EnableIPv6) {
		return true
	}

	return false
}

// SetEnableIPv6 gets a reference to the given bool and assigns it to the EnableIPv6 field.
func (o *NetworkConfiguration) SetEnableIPv6(v bool) {
	o.EnableIPv6 = &v
}

// GetEnableRemoteAccess returns the EnableRemoteAccess field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetEnableRemoteAccess() bool {
	if o == nil || IsNil(o.EnableRemoteAccess) {
		var ret bool
		return ret
	}
	return *o.EnableRemoteAccess
}

// GetEnableRemoteAccessOk returns a tuple with the EnableRemoteAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetEnableRemoteAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableRemoteAccess) {
		return nil, false
	}
	return o.EnableRemoteAccess, true
}

// HasEnableRemoteAccess returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasEnableRemoteAccess() bool {
	if o != nil && !IsNil(o.EnableRemoteAccess) {
		return true
	}

	return false
}

// SetEnableRemoteAccess gets a reference to the given bool and assigns it to the EnableRemoteAccess field.
func (o *NetworkConfiguration) SetEnableRemoteAccess(v bool) {
	o.EnableRemoteAccess = &v
}

// GetLocalNetworkSubnets returns the LocalNetworkSubnets field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetLocalNetworkSubnets() []string {
	if o == nil || IsNil(o.LocalNetworkSubnets) {
		var ret []string
		return ret
	}
	return o.LocalNetworkSubnets
}

// GetLocalNetworkSubnetsOk returns a tuple with the LocalNetworkSubnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetLocalNetworkSubnetsOk() ([]string, bool) {
	if o == nil || IsNil(o.LocalNetworkSubnets) {
		return nil, false
	}
	return o.LocalNetworkSubnets, true
}

// HasLocalNetworkSubnets returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasLocalNetworkSubnets() bool {
	if o != nil && !IsNil(o.LocalNetworkSubnets) {
		return true
	}

	return false
}

// SetLocalNetworkSubnets gets a reference to the given []string and assigns it to the LocalNetworkSubnets field.
func (o *NetworkConfiguration) SetLocalNetworkSubnets(v []string) {
	o.LocalNetworkSubnets = v
}

// GetLocalNetworkAddresses returns the LocalNetworkAddresses field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetLocalNetworkAddresses() []string {
	if o == nil || IsNil(o.LocalNetworkAddresses) {
		var ret []string
		return ret
	}
	return o.LocalNetworkAddresses
}

// GetLocalNetworkAddressesOk returns a tuple with the LocalNetworkAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetLocalNetworkAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.LocalNetworkAddresses) {
		return nil, false
	}
	return o.LocalNetworkAddresses, true
}

// HasLocalNetworkAddresses returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasLocalNetworkAddresses() bool {
	if o != nil && !IsNil(o.LocalNetworkAddresses) {
		return true
	}

	return false
}

// SetLocalNetworkAddresses gets a reference to the given []string and assigns it to the LocalNetworkAddresses field.
func (o *NetworkConfiguration) SetLocalNetworkAddresses(v []string) {
	o.LocalNetworkAddresses = v
}

// GetKnownProxies returns the KnownProxies field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetKnownProxies() []string {
	if o == nil || IsNil(o.KnownProxies) {
		var ret []string
		return ret
	}
	return o.KnownProxies
}

// GetKnownProxiesOk returns a tuple with the KnownProxies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetKnownProxiesOk() ([]string, bool) {
	if o == nil || IsNil(o.KnownProxies) {
		return nil, false
	}
	return o.KnownProxies, true
}

// HasKnownProxies returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasKnownProxies() bool {
	if o != nil && !IsNil(o.KnownProxies) {
		return true
	}

	return false
}

// SetKnownProxies gets a reference to the given []string and assigns it to the KnownProxies field.
func (o *NetworkConfiguration) SetKnownProxies(v []string) {
	o.KnownProxies = v
}

// GetIgnoreVirtualInterfaces returns the IgnoreVirtualInterfaces field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetIgnoreVirtualInterfaces() bool {
	if o == nil || IsNil(o.IgnoreVirtualInterfaces) {
		var ret bool
		return ret
	}
	return *o.IgnoreVirtualInterfaces
}

// GetIgnoreVirtualInterfacesOk returns a tuple with the IgnoreVirtualInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetIgnoreVirtualInterfacesOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreVirtualInterfaces) {
		return nil, false
	}
	return o.IgnoreVirtualInterfaces, true
}

// HasIgnoreVirtualInterfaces returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasIgnoreVirtualInterfaces() bool {
	if o != nil && !IsNil(o.IgnoreVirtualInterfaces) {
		return true
	}

	return false
}

// SetIgnoreVirtualInterfaces gets a reference to the given bool and assigns it to the IgnoreVirtualInterfaces field.
func (o *NetworkConfiguration) SetIgnoreVirtualInterfaces(v bool) {
	o.IgnoreVirtualInterfaces = &v
}

// GetVirtualInterfaceNames returns the VirtualInterfaceNames field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetVirtualInterfaceNames() []string {
	if o == nil || IsNil(o.VirtualInterfaceNames) {
		var ret []string
		return ret
	}
	return o.VirtualInterfaceNames
}

// GetVirtualInterfaceNamesOk returns a tuple with the VirtualInterfaceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetVirtualInterfaceNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.VirtualInterfaceNames) {
		return nil, false
	}
	return o.VirtualInterfaceNames, true
}

// HasVirtualInterfaceNames returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasVirtualInterfaceNames() bool {
	if o != nil && !IsNil(o.VirtualInterfaceNames) {
		return true
	}

	return false
}

// SetVirtualInterfaceNames gets a reference to the given []string and assigns it to the VirtualInterfaceNames field.
func (o *NetworkConfiguration) SetVirtualInterfaceNames(v []string) {
	o.VirtualInterfaceNames = v
}

// GetEnablePublishedServerUriByRequest returns the EnablePublishedServerUriByRequest field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetEnablePublishedServerUriByRequest() bool {
	if o == nil || IsNil(o.EnablePublishedServerUriByRequest) {
		var ret bool
		return ret
	}
	return *o.EnablePublishedServerUriByRequest
}

// GetEnablePublishedServerUriByRequestOk returns a tuple with the EnablePublishedServerUriByRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetEnablePublishedServerUriByRequestOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePublishedServerUriByRequest) {
		return nil, false
	}
	return o.EnablePublishedServerUriByRequest, true
}

// HasEnablePublishedServerUriByRequest returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasEnablePublishedServerUriByRequest() bool {
	if o != nil && !IsNil(o.EnablePublishedServerUriByRequest) {
		return true
	}

	return false
}

// SetEnablePublishedServerUriByRequest gets a reference to the given bool and assigns it to the EnablePublishedServerUriByRequest field.
func (o *NetworkConfiguration) SetEnablePublishedServerUriByRequest(v bool) {
	o.EnablePublishedServerUriByRequest = &v
}

// GetPublishedServerUriBySubnet returns the PublishedServerUriBySubnet field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetPublishedServerUriBySubnet() []string {
	if o == nil || IsNil(o.PublishedServerUriBySubnet) {
		var ret []string
		return ret
	}
	return o.PublishedServerUriBySubnet
}

// GetPublishedServerUriBySubnetOk returns a tuple with the PublishedServerUriBySubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetPublishedServerUriBySubnetOk() ([]string, bool) {
	if o == nil || IsNil(o.PublishedServerUriBySubnet) {
		return nil, false
	}
	return o.PublishedServerUriBySubnet, true
}

// HasPublishedServerUriBySubnet returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasPublishedServerUriBySubnet() bool {
	if o != nil && !IsNil(o.PublishedServerUriBySubnet) {
		return true
	}

	return false
}

// SetPublishedServerUriBySubnet gets a reference to the given []string and assigns it to the PublishedServerUriBySubnet field.
func (o *NetworkConfiguration) SetPublishedServerUriBySubnet(v []string) {
	o.PublishedServerUriBySubnet = v
}

// GetRemoteIPFilter returns the RemoteIPFilter field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetRemoteIPFilter() []string {
	if o == nil || IsNil(o.RemoteIPFilter) {
		var ret []string
		return ret
	}
	return o.RemoteIPFilter
}

// GetRemoteIPFilterOk returns a tuple with the RemoteIPFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetRemoteIPFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.RemoteIPFilter) {
		return nil, false
	}
	return o.RemoteIPFilter, true
}

// HasRemoteIPFilter returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasRemoteIPFilter() bool {
	if o != nil && !IsNil(o.RemoteIPFilter) {
		return true
	}

	return false
}

// SetRemoteIPFilter gets a reference to the given []string and assigns it to the RemoteIPFilter field.
func (o *NetworkConfiguration) SetRemoteIPFilter(v []string) {
	o.RemoteIPFilter = v
}

// GetIsRemoteIPFilterBlacklist returns the IsRemoteIPFilterBlacklist field value if set, zero value otherwise.
func (o *NetworkConfiguration) GetIsRemoteIPFilterBlacklist() bool {
	if o == nil || IsNil(o.IsRemoteIPFilterBlacklist) {
		var ret bool
		return ret
	}
	return *o.IsRemoteIPFilterBlacklist
}

// GetIsRemoteIPFilterBlacklistOk returns a tuple with the IsRemoteIPFilterBlacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfiguration) GetIsRemoteIPFilterBlacklistOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRemoteIPFilterBlacklist) {
		return nil, false
	}
	return o.IsRemoteIPFilterBlacklist, true
}

// HasIsRemoteIPFilterBlacklist returns a boolean if a field has been set.
func (o *NetworkConfiguration) HasIsRemoteIPFilterBlacklist() bool {
	if o != nil && !IsNil(o.IsRemoteIPFilterBlacklist) {
		return true
	}

	return false
}

// SetIsRemoteIPFilterBlacklist gets a reference to the given bool and assigns it to the IsRemoteIPFilterBlacklist field.
func (o *NetworkConfiguration) SetIsRemoteIPFilterBlacklist(v bool) {
	o.IsRemoteIPFilterBlacklist = &v
}

func (o NetworkConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaseUrl) {
		toSerialize["BaseUrl"] = o.BaseUrl
	}
	if !IsNil(o.EnableHttps) {
		toSerialize["EnableHttps"] = o.EnableHttps
	}
	if !IsNil(o.RequireHttps) {
		toSerialize["RequireHttps"] = o.RequireHttps
	}
	if !IsNil(o.CertificatePath) {
		toSerialize["CertificatePath"] = o.CertificatePath
	}
	if !IsNil(o.CertificatePassword) {
		toSerialize["CertificatePassword"] = o.CertificatePassword
	}
	if !IsNil(o.InternalHttpPort) {
		toSerialize["InternalHttpPort"] = o.InternalHttpPort
	}
	if !IsNil(o.InternalHttpsPort) {
		toSerialize["InternalHttpsPort"] = o.InternalHttpsPort
	}
	if !IsNil(o.PublicHttpPort) {
		toSerialize["PublicHttpPort"] = o.PublicHttpPort
	}
	if !IsNil(o.PublicHttpsPort) {
		toSerialize["PublicHttpsPort"] = o.PublicHttpsPort
	}
	if !IsNil(o.AutoDiscovery) {
		toSerialize["AutoDiscovery"] = o.AutoDiscovery
	}
	if !IsNil(o.EnableUPnP) {
		toSerialize["EnableUPnP"] = o.EnableUPnP
	}
	if !IsNil(o.EnableIPv4) {
		toSerialize["EnableIPv4"] = o.EnableIPv4
	}
	if !IsNil(o.EnableIPv6) {
		toSerialize["EnableIPv6"] = o.EnableIPv6
	}
	if !IsNil(o.EnableRemoteAccess) {
		toSerialize["EnableRemoteAccess"] = o.EnableRemoteAccess
	}
	if !IsNil(o.LocalNetworkSubnets) {
		toSerialize["LocalNetworkSubnets"] = o.LocalNetworkSubnets
	}
	if !IsNil(o.LocalNetworkAddresses) {
		toSerialize["LocalNetworkAddresses"] = o.LocalNetworkAddresses
	}
	if !IsNil(o.KnownProxies) {
		toSerialize["KnownProxies"] = o.KnownProxies
	}
	if !IsNil(o.IgnoreVirtualInterfaces) {
		toSerialize["IgnoreVirtualInterfaces"] = o.IgnoreVirtualInterfaces
	}
	if !IsNil(o.VirtualInterfaceNames) {
		toSerialize["VirtualInterfaceNames"] = o.VirtualInterfaceNames
	}
	if !IsNil(o.EnablePublishedServerUriByRequest) {
		toSerialize["EnablePublishedServerUriByRequest"] = o.EnablePublishedServerUriByRequest
	}
	if !IsNil(o.PublishedServerUriBySubnet) {
		toSerialize["PublishedServerUriBySubnet"] = o.PublishedServerUriBySubnet
	}
	if !IsNil(o.RemoteIPFilter) {
		toSerialize["RemoteIPFilter"] = o.RemoteIPFilter
	}
	if !IsNil(o.IsRemoteIPFilterBlacklist) {
		toSerialize["IsRemoteIPFilterBlacklist"] = o.IsRemoteIPFilterBlacklist
	}
	return toSerialize, nil
}

type NullableNetworkConfiguration struct {
	value *NetworkConfiguration
	isSet bool
}

func (v NullableNetworkConfiguration) Get() *NetworkConfiguration {
	return v.value
}

func (v *NullableNetworkConfiguration) Set(val *NetworkConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkConfiguration(val *NetworkConfiguration) *NullableNetworkConfiguration {
	return &NullableNetworkConfiguration{value: val, isSet: true}
}

func (v NullableNetworkConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


