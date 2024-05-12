# NetworkConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **string** | Gets or sets a value used to specify the URL prefix that your Jellyfin instance can be accessed at. | [optional] 
**EnableHttps** | Pointer to **bool** | Gets or sets a value indicating whether to use HTTPS. | [optional] 
**RequireHttps** | Pointer to **bool** | Gets or sets a value indicating whether the server should force connections over HTTPS. | [optional] 
**CertificatePath** | Pointer to **string** | Gets or sets the filesystem path of an X.509 certificate to use for SSL. | [optional] 
**CertificatePassword** | Pointer to **string** | Gets or sets the password required to access the X.509 certificate data in the file specified by MediaBrowser.Common.Net.NetworkConfiguration.CertificatePath. | [optional] 
**InternalHttpPort** | Pointer to **int32** | Gets or sets the internal HTTP server port. | [optional] 
**InternalHttpsPort** | Pointer to **int32** | Gets or sets the internal HTTPS server port. | [optional] 
**PublicHttpPort** | Pointer to **int32** | Gets or sets the public HTTP port. | [optional] 
**PublicHttpsPort** | Pointer to **int32** | Gets or sets the public HTTPS port. | [optional] 
**AutoDiscovery** | Pointer to **bool** | Gets or sets a value indicating whether Autodiscovery is enabled. | [optional] 
**EnableUPnP** | Pointer to **bool** | Gets or sets a value indicating whether to enable automatic port forwarding. | [optional] 
**EnableIPv4** | Pointer to **bool** | Gets or sets a value indicating whether IPv6 is enabled. | [optional] 
**EnableIPv6** | Pointer to **bool** | Gets or sets a value indicating whether IPv6 is enabled. | [optional] 
**EnableRemoteAccess** | Pointer to **bool** | Gets or sets a value indicating whether access from outside of the LAN is permitted. | [optional] 
**LocalNetworkSubnets** | Pointer to **[]string** | Gets or sets the subnets that are deemed to make up the LAN. | [optional] 
**LocalNetworkAddresses** | Pointer to **[]string** | Gets or sets the interface addresses which Jellyfin will bind to. If empty, all interfaces will be used. | [optional] 
**KnownProxies** | Pointer to **[]string** | Gets or sets the known proxies. | [optional] 
**IgnoreVirtualInterfaces** | Pointer to **bool** | Gets or sets a value indicating whether address names that match MediaBrowser.Common.Net.NetworkConfiguration.VirtualInterfaceNames should be ignored for the purposes of binding. | [optional] 
**VirtualInterfaceNames** | Pointer to **[]string** | Gets or sets a value indicating the interface name prefixes that should be ignored. The list can be comma separated and values are case-insensitive. &lt;seealso cref&#x3D;\&quot;P:MediaBrowser.Common.Net.NetworkConfiguration.IgnoreVirtualInterfaces\&quot; /&gt;. | [optional] 
**EnablePublishedServerUriByRequest** | Pointer to **bool** | Gets or sets a value indicating whether the published server uri is based on information in HTTP requests. | [optional] 
**PublishedServerUriBySubnet** | Pointer to **[]string** | Gets or sets the PublishedServerUriBySubnet  Gets or sets PublishedServerUri to advertise for specific subnets. | [optional] 
**RemoteIPFilter** | Pointer to **[]string** | Gets or sets the filter for remote IP connectivity. Used in conjunction with &lt;seealso cref&#x3D;\&quot;P:MediaBrowser.Common.Net.NetworkConfiguration.IsRemoteIPFilterBlacklist\&quot; /&gt;. | [optional] 
**IsRemoteIPFilterBlacklist** | Pointer to **bool** | Gets or sets a value indicating whether &lt;seealso cref&#x3D;\&quot;P:MediaBrowser.Common.Net.NetworkConfiguration.RemoteIPFilter\&quot; /&gt; contains a blacklist or a whitelist. Default is a whitelist. | [optional] 

## Methods

### NewNetworkConfiguration

`func NewNetworkConfiguration() *NetworkConfiguration`

NewNetworkConfiguration instantiates a new NetworkConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConfigurationWithDefaults

`func NewNetworkConfigurationWithDefaults() *NetworkConfiguration`

NewNetworkConfigurationWithDefaults instantiates a new NetworkConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *NetworkConfiguration) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *NetworkConfiguration) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *NetworkConfiguration) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *NetworkConfiguration) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetEnableHttps

`func (o *NetworkConfiguration) GetEnableHttps() bool`

GetEnableHttps returns the EnableHttps field if non-nil, zero value otherwise.

### GetEnableHttpsOk

`func (o *NetworkConfiguration) GetEnableHttpsOk() (*bool, bool)`

GetEnableHttpsOk returns a tuple with the EnableHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttps

`func (o *NetworkConfiguration) SetEnableHttps(v bool)`

SetEnableHttps sets EnableHttps field to given value.

### HasEnableHttps

`func (o *NetworkConfiguration) HasEnableHttps() bool`

HasEnableHttps returns a boolean if a field has been set.

### GetRequireHttps

`func (o *NetworkConfiguration) GetRequireHttps() bool`

GetRequireHttps returns the RequireHttps field if non-nil, zero value otherwise.

### GetRequireHttpsOk

`func (o *NetworkConfiguration) GetRequireHttpsOk() (*bool, bool)`

GetRequireHttpsOk returns a tuple with the RequireHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHttps

`func (o *NetworkConfiguration) SetRequireHttps(v bool)`

SetRequireHttps sets RequireHttps field to given value.

### HasRequireHttps

`func (o *NetworkConfiguration) HasRequireHttps() bool`

HasRequireHttps returns a boolean if a field has been set.

### GetCertificatePath

`func (o *NetworkConfiguration) GetCertificatePath() string`

GetCertificatePath returns the CertificatePath field if non-nil, zero value otherwise.

### GetCertificatePathOk

`func (o *NetworkConfiguration) GetCertificatePathOk() (*string, bool)`

GetCertificatePathOk returns a tuple with the CertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePath

`func (o *NetworkConfiguration) SetCertificatePath(v string)`

SetCertificatePath sets CertificatePath field to given value.

### HasCertificatePath

`func (o *NetworkConfiguration) HasCertificatePath() bool`

HasCertificatePath returns a boolean if a field has been set.

### GetCertificatePassword

`func (o *NetworkConfiguration) GetCertificatePassword() string`

GetCertificatePassword returns the CertificatePassword field if non-nil, zero value otherwise.

### GetCertificatePasswordOk

`func (o *NetworkConfiguration) GetCertificatePasswordOk() (*string, bool)`

GetCertificatePasswordOk returns a tuple with the CertificatePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePassword

`func (o *NetworkConfiguration) SetCertificatePassword(v string)`

SetCertificatePassword sets CertificatePassword field to given value.

### HasCertificatePassword

`func (o *NetworkConfiguration) HasCertificatePassword() bool`

HasCertificatePassword returns a boolean if a field has been set.

### GetInternalHttpPort

`func (o *NetworkConfiguration) GetInternalHttpPort() int32`

GetInternalHttpPort returns the InternalHttpPort field if non-nil, zero value otherwise.

### GetInternalHttpPortOk

`func (o *NetworkConfiguration) GetInternalHttpPortOk() (*int32, bool)`

GetInternalHttpPortOk returns a tuple with the InternalHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHttpPort

`func (o *NetworkConfiguration) SetInternalHttpPort(v int32)`

SetInternalHttpPort sets InternalHttpPort field to given value.

### HasInternalHttpPort

`func (o *NetworkConfiguration) HasInternalHttpPort() bool`

HasInternalHttpPort returns a boolean if a field has been set.

### GetInternalHttpsPort

`func (o *NetworkConfiguration) GetInternalHttpsPort() int32`

GetInternalHttpsPort returns the InternalHttpsPort field if non-nil, zero value otherwise.

### GetInternalHttpsPortOk

`func (o *NetworkConfiguration) GetInternalHttpsPortOk() (*int32, bool)`

GetInternalHttpsPortOk returns a tuple with the InternalHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalHttpsPort

`func (o *NetworkConfiguration) SetInternalHttpsPort(v int32)`

SetInternalHttpsPort sets InternalHttpsPort field to given value.

### HasInternalHttpsPort

`func (o *NetworkConfiguration) HasInternalHttpsPort() bool`

HasInternalHttpsPort returns a boolean if a field has been set.

### GetPublicHttpPort

`func (o *NetworkConfiguration) GetPublicHttpPort() int32`

GetPublicHttpPort returns the PublicHttpPort field if non-nil, zero value otherwise.

### GetPublicHttpPortOk

`func (o *NetworkConfiguration) GetPublicHttpPortOk() (*int32, bool)`

GetPublicHttpPortOk returns a tuple with the PublicHttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicHttpPort

`func (o *NetworkConfiguration) SetPublicHttpPort(v int32)`

SetPublicHttpPort sets PublicHttpPort field to given value.

### HasPublicHttpPort

`func (o *NetworkConfiguration) HasPublicHttpPort() bool`

HasPublicHttpPort returns a boolean if a field has been set.

### GetPublicHttpsPort

`func (o *NetworkConfiguration) GetPublicHttpsPort() int32`

GetPublicHttpsPort returns the PublicHttpsPort field if non-nil, zero value otherwise.

### GetPublicHttpsPortOk

`func (o *NetworkConfiguration) GetPublicHttpsPortOk() (*int32, bool)`

GetPublicHttpsPortOk returns a tuple with the PublicHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicHttpsPort

`func (o *NetworkConfiguration) SetPublicHttpsPort(v int32)`

SetPublicHttpsPort sets PublicHttpsPort field to given value.

### HasPublicHttpsPort

`func (o *NetworkConfiguration) HasPublicHttpsPort() bool`

HasPublicHttpsPort returns a boolean if a field has been set.

### GetAutoDiscovery

`func (o *NetworkConfiguration) GetAutoDiscovery() bool`

GetAutoDiscovery returns the AutoDiscovery field if non-nil, zero value otherwise.

### GetAutoDiscoveryOk

`func (o *NetworkConfiguration) GetAutoDiscoveryOk() (*bool, bool)`

GetAutoDiscoveryOk returns a tuple with the AutoDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDiscovery

`func (o *NetworkConfiguration) SetAutoDiscovery(v bool)`

SetAutoDiscovery sets AutoDiscovery field to given value.

### HasAutoDiscovery

`func (o *NetworkConfiguration) HasAutoDiscovery() bool`

HasAutoDiscovery returns a boolean if a field has been set.

### GetEnableUPnP

`func (o *NetworkConfiguration) GetEnableUPnP() bool`

GetEnableUPnP returns the EnableUPnP field if non-nil, zero value otherwise.

### GetEnableUPnPOk

`func (o *NetworkConfiguration) GetEnableUPnPOk() (*bool, bool)`

GetEnableUPnPOk returns a tuple with the EnableUPnP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUPnP

`func (o *NetworkConfiguration) SetEnableUPnP(v bool)`

SetEnableUPnP sets EnableUPnP field to given value.

### HasEnableUPnP

`func (o *NetworkConfiguration) HasEnableUPnP() bool`

HasEnableUPnP returns a boolean if a field has been set.

### GetEnableIPv4

`func (o *NetworkConfiguration) GetEnableIPv4() bool`

GetEnableIPv4 returns the EnableIPv4 field if non-nil, zero value otherwise.

### GetEnableIPv4Ok

`func (o *NetworkConfiguration) GetEnableIPv4Ok() (*bool, bool)`

GetEnableIPv4Ok returns a tuple with the EnableIPv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIPv4

`func (o *NetworkConfiguration) SetEnableIPv4(v bool)`

SetEnableIPv4 sets EnableIPv4 field to given value.

### HasEnableIPv4

`func (o *NetworkConfiguration) HasEnableIPv4() bool`

HasEnableIPv4 returns a boolean if a field has been set.

### GetEnableIPv6

`func (o *NetworkConfiguration) GetEnableIPv6() bool`

GetEnableIPv6 returns the EnableIPv6 field if non-nil, zero value otherwise.

### GetEnableIPv6Ok

`func (o *NetworkConfiguration) GetEnableIPv6Ok() (*bool, bool)`

GetEnableIPv6Ok returns a tuple with the EnableIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIPv6

`func (o *NetworkConfiguration) SetEnableIPv6(v bool)`

SetEnableIPv6 sets EnableIPv6 field to given value.

### HasEnableIPv6

`func (o *NetworkConfiguration) HasEnableIPv6() bool`

HasEnableIPv6 returns a boolean if a field has been set.

### GetEnableRemoteAccess

`func (o *NetworkConfiguration) GetEnableRemoteAccess() bool`

GetEnableRemoteAccess returns the EnableRemoteAccess field if non-nil, zero value otherwise.

### GetEnableRemoteAccessOk

`func (o *NetworkConfiguration) GetEnableRemoteAccessOk() (*bool, bool)`

GetEnableRemoteAccessOk returns a tuple with the EnableRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteAccess

`func (o *NetworkConfiguration) SetEnableRemoteAccess(v bool)`

SetEnableRemoteAccess sets EnableRemoteAccess field to given value.

### HasEnableRemoteAccess

`func (o *NetworkConfiguration) HasEnableRemoteAccess() bool`

HasEnableRemoteAccess returns a boolean if a field has been set.

### GetLocalNetworkSubnets

`func (o *NetworkConfiguration) GetLocalNetworkSubnets() []string`

GetLocalNetworkSubnets returns the LocalNetworkSubnets field if non-nil, zero value otherwise.

### GetLocalNetworkSubnetsOk

`func (o *NetworkConfiguration) GetLocalNetworkSubnetsOk() (*[]string, bool)`

GetLocalNetworkSubnetsOk returns a tuple with the LocalNetworkSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworkSubnets

`func (o *NetworkConfiguration) SetLocalNetworkSubnets(v []string)`

SetLocalNetworkSubnets sets LocalNetworkSubnets field to given value.

### HasLocalNetworkSubnets

`func (o *NetworkConfiguration) HasLocalNetworkSubnets() bool`

HasLocalNetworkSubnets returns a boolean if a field has been set.

### GetLocalNetworkAddresses

`func (o *NetworkConfiguration) GetLocalNetworkAddresses() []string`

GetLocalNetworkAddresses returns the LocalNetworkAddresses field if non-nil, zero value otherwise.

### GetLocalNetworkAddressesOk

`func (o *NetworkConfiguration) GetLocalNetworkAddressesOk() (*[]string, bool)`

GetLocalNetworkAddressesOk returns a tuple with the LocalNetworkAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworkAddresses

`func (o *NetworkConfiguration) SetLocalNetworkAddresses(v []string)`

SetLocalNetworkAddresses sets LocalNetworkAddresses field to given value.

### HasLocalNetworkAddresses

`func (o *NetworkConfiguration) HasLocalNetworkAddresses() bool`

HasLocalNetworkAddresses returns a boolean if a field has been set.

### GetKnownProxies

`func (o *NetworkConfiguration) GetKnownProxies() []string`

GetKnownProxies returns the KnownProxies field if non-nil, zero value otherwise.

### GetKnownProxiesOk

`func (o *NetworkConfiguration) GetKnownProxiesOk() (*[]string, bool)`

GetKnownProxiesOk returns a tuple with the KnownProxies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownProxies

`func (o *NetworkConfiguration) SetKnownProxies(v []string)`

SetKnownProxies sets KnownProxies field to given value.

### HasKnownProxies

`func (o *NetworkConfiguration) HasKnownProxies() bool`

HasKnownProxies returns a boolean if a field has been set.

### GetIgnoreVirtualInterfaces

`func (o *NetworkConfiguration) GetIgnoreVirtualInterfaces() bool`

GetIgnoreVirtualInterfaces returns the IgnoreVirtualInterfaces field if non-nil, zero value otherwise.

### GetIgnoreVirtualInterfacesOk

`func (o *NetworkConfiguration) GetIgnoreVirtualInterfacesOk() (*bool, bool)`

GetIgnoreVirtualInterfacesOk returns a tuple with the IgnoreVirtualInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreVirtualInterfaces

`func (o *NetworkConfiguration) SetIgnoreVirtualInterfaces(v bool)`

SetIgnoreVirtualInterfaces sets IgnoreVirtualInterfaces field to given value.

### HasIgnoreVirtualInterfaces

`func (o *NetworkConfiguration) HasIgnoreVirtualInterfaces() bool`

HasIgnoreVirtualInterfaces returns a boolean if a field has been set.

### GetVirtualInterfaceNames

`func (o *NetworkConfiguration) GetVirtualInterfaceNames() []string`

GetVirtualInterfaceNames returns the VirtualInterfaceNames field if non-nil, zero value otherwise.

### GetVirtualInterfaceNamesOk

`func (o *NetworkConfiguration) GetVirtualInterfaceNamesOk() (*[]string, bool)`

GetVirtualInterfaceNamesOk returns a tuple with the VirtualInterfaceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualInterfaceNames

`func (o *NetworkConfiguration) SetVirtualInterfaceNames(v []string)`

SetVirtualInterfaceNames sets VirtualInterfaceNames field to given value.

### HasVirtualInterfaceNames

`func (o *NetworkConfiguration) HasVirtualInterfaceNames() bool`

HasVirtualInterfaceNames returns a boolean if a field has been set.

### GetEnablePublishedServerUriByRequest

`func (o *NetworkConfiguration) GetEnablePublishedServerUriByRequest() bool`

GetEnablePublishedServerUriByRequest returns the EnablePublishedServerUriByRequest field if non-nil, zero value otherwise.

### GetEnablePublishedServerUriByRequestOk

`func (o *NetworkConfiguration) GetEnablePublishedServerUriByRequestOk() (*bool, bool)`

GetEnablePublishedServerUriByRequestOk returns a tuple with the EnablePublishedServerUriByRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublishedServerUriByRequest

`func (o *NetworkConfiguration) SetEnablePublishedServerUriByRequest(v bool)`

SetEnablePublishedServerUriByRequest sets EnablePublishedServerUriByRequest field to given value.

### HasEnablePublishedServerUriByRequest

`func (o *NetworkConfiguration) HasEnablePublishedServerUriByRequest() bool`

HasEnablePublishedServerUriByRequest returns a boolean if a field has been set.

### GetPublishedServerUriBySubnet

`func (o *NetworkConfiguration) GetPublishedServerUriBySubnet() []string`

GetPublishedServerUriBySubnet returns the PublishedServerUriBySubnet field if non-nil, zero value otherwise.

### GetPublishedServerUriBySubnetOk

`func (o *NetworkConfiguration) GetPublishedServerUriBySubnetOk() (*[]string, bool)`

GetPublishedServerUriBySubnetOk returns a tuple with the PublishedServerUriBySubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedServerUriBySubnet

`func (o *NetworkConfiguration) SetPublishedServerUriBySubnet(v []string)`

SetPublishedServerUriBySubnet sets PublishedServerUriBySubnet field to given value.

### HasPublishedServerUriBySubnet

`func (o *NetworkConfiguration) HasPublishedServerUriBySubnet() bool`

HasPublishedServerUriBySubnet returns a boolean if a field has been set.

### GetRemoteIPFilter

`func (o *NetworkConfiguration) GetRemoteIPFilter() []string`

GetRemoteIPFilter returns the RemoteIPFilter field if non-nil, zero value otherwise.

### GetRemoteIPFilterOk

`func (o *NetworkConfiguration) GetRemoteIPFilterOk() (*[]string, bool)`

GetRemoteIPFilterOk returns a tuple with the RemoteIPFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIPFilter

`func (o *NetworkConfiguration) SetRemoteIPFilter(v []string)`

SetRemoteIPFilter sets RemoteIPFilter field to given value.

### HasRemoteIPFilter

`func (o *NetworkConfiguration) HasRemoteIPFilter() bool`

HasRemoteIPFilter returns a boolean if a field has been set.

### GetIsRemoteIPFilterBlacklist

`func (o *NetworkConfiguration) GetIsRemoteIPFilterBlacklist() bool`

GetIsRemoteIPFilterBlacklist returns the IsRemoteIPFilterBlacklist field if non-nil, zero value otherwise.

### GetIsRemoteIPFilterBlacklistOk

`func (o *NetworkConfiguration) GetIsRemoteIPFilterBlacklistOk() (*bool, bool)`

GetIsRemoteIPFilterBlacklistOk returns a tuple with the IsRemoteIPFilterBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteIPFilterBlacklist

`func (o *NetworkConfiguration) SetIsRemoteIPFilterBlacklist(v bool)`

SetIsRemoteIPFilterBlacklist sets IsRemoteIPFilterBlacklist field to given value.

### HasIsRemoteIPFilterBlacklist

`func (o *NetworkConfiguration) HasIsRemoteIPFilterBlacklist() bool`

HasIsRemoteIPFilterBlacklist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


