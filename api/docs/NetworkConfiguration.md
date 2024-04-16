# NetworkConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireHttps** | Pointer to **bool** | Gets or sets a value indicating whether the server should force connections over HTTPS. | [optional] 
**CertificatePath** | Pointer to **string** | Gets or sets the filesystem path of an X.509 certificate to use for SSL. | [optional] 
**CertificatePassword** | Pointer to **string** | Gets or sets the password required to access the X.509 certificate data in the file specified by Jellyfin.Networking.Configuration.NetworkConfiguration.CertificatePath. | [optional] 
**BaseUrl** | Pointer to **string** | Gets or sets a value used to specify the URL prefix that your Jellyfin instance can be accessed at. | [optional] 
**PublicHttpsPort** | Pointer to **int32** | Gets or sets the public HTTPS port. | [optional] 
**HttpServerPortNumber** | Pointer to **int32** | Gets or sets the HTTP server port number. | [optional] 
**HttpsPortNumber** | Pointer to **int32** | Gets or sets the HTTPS server port number. | [optional] 
**EnableHttps** | Pointer to **bool** | Gets or sets a value indicating whether to use HTTPS. | [optional] 
**PublicPort** | Pointer to **int32** | Gets or sets the public mapped port. | [optional] 
**UPnPCreateHttpPortMap** | Pointer to **bool** | Gets or sets a value indicating whether the http port should be mapped as part of UPnP automatic port forwarding. | [optional] 
**UDPPortRange** | Pointer to **string** | Gets or sets the UDPPortRange. | [optional] 
**EnableIPV6** | Pointer to **bool** | Gets or sets a value indicating whether gets or sets IPV6 capability. | [optional] 
**EnableIPV4** | Pointer to **bool** | Gets or sets a value indicating whether gets or sets IPV4 capability. | [optional] 
**EnableSSDPTracing** | Pointer to **bool** | Gets or sets a value indicating whether detailed SSDP logs are sent to the console/log.  \&quot;Emby.Dlna\&quot;: \&quot;Debug\&quot; must be set in logging.default.json for this property to have any effect. | [optional] 
**SSDPTracingFilter** | Pointer to **string** | Gets or sets the SSDPTracingFilter  Gets or sets a value indicating whether an IP address is to be used to filter the detailed ssdp logs that are being sent to the console/log.  If the setting \&quot;Emby.Dlna\&quot;: \&quot;Debug\&quot; msut be set in logging.default.json for this property to work. | [optional] 
**UDPSendCount** | Pointer to **int32** | Gets or sets the number of times SSDP UDP messages are sent. | [optional] 
**UDPSendDelay** | Pointer to **int32** | Gets or sets the delay between each groups of SSDP messages (in ms). | [optional] 
**IgnoreVirtualInterfaces** | Pointer to **bool** | Gets or sets a value indicating whether address names that match Jellyfin.Networking.Configuration.NetworkConfiguration.VirtualInterfaceNames should be Ignore for the purposes of binding. | [optional] 
**VirtualInterfaceNames** | Pointer to **string** | Gets or sets a value indicating the interfaces that should be ignored. The list can be comma separated. &lt;seealso cref&#x3D;\&quot;P:Jellyfin.Networking.Configuration.NetworkConfiguration.IgnoreVirtualInterfaces\&quot; /&gt;. | [optional] 
**GatewayMonitorPeriod** | Pointer to **int32** | Gets or sets the time (in seconds) between the pings of SSDP gateway monitor. | [optional] 
**EnableMultiSocketBinding** | Pointer to **bool** | Gets a value indicating whether multi-socket binding is available. | [optional] [readonly] 
**TrustAllIP6Interfaces** | Pointer to **bool** | Gets or sets a value indicating whether all IPv6 interfaces should be treated as on the internal network.  Depending on the address range implemented ULA ranges might not be used. | [optional] 
**HDHomerunPortRange** | Pointer to **string** | Gets or sets the ports that HDHomerun uses. | [optional] 
**PublishedServerUriBySubnet** | Pointer to **[]string** | Gets or sets the PublishedServerUriBySubnet  Gets or sets PublishedServerUri to advertise for specific subnets. | [optional] 
**AutoDiscoveryTracing** | Pointer to **bool** | Gets or sets a value indicating whether Autodiscovery tracing is enabled. | [optional] 
**AutoDiscovery** | Pointer to **bool** | Gets or sets a value indicating whether Autodiscovery is enabled. | [optional] 
**RemoteIPFilter** | Pointer to **[]string** | Gets or sets the filter for remote IP connectivity. Used in conjuntion with &lt;seealso cref&#x3D;\&quot;P:Jellyfin.Networking.Configuration.NetworkConfiguration.IsRemoteIPFilterBlacklist\&quot; /&gt;. | [optional] 
**IsRemoteIPFilterBlacklist** | Pointer to **bool** | Gets or sets a value indicating whether &lt;seealso cref&#x3D;\&quot;P:Jellyfin.Networking.Configuration.NetworkConfiguration.RemoteIPFilter\&quot; /&gt; contains a blacklist or a whitelist. Default is a whitelist. | [optional] 
**EnableUPnP** | Pointer to **bool** | Gets or sets a value indicating whether to enable automatic port forwarding. | [optional] 
**EnableRemoteAccess** | Pointer to **bool** | Gets or sets a value indicating whether access outside of the LAN is permitted. | [optional] 
**LocalNetworkSubnets** | Pointer to **[]string** | Gets or sets the subnets that are deemed to make up the LAN. | [optional] 
**LocalNetworkAddresses** | Pointer to **[]string** | Gets or sets the interface addresses which Jellyfin will bind to. If empty, all interfaces will be used. | [optional] 
**KnownProxies** | Pointer to **[]string** | Gets or sets the known proxies. If the proxy is a network, it&#39;s added to the KnownNetworks. | [optional] 
**EnablePublishedServerUriByRequest** | Pointer to **bool** | Gets or sets a value indicating whether the published server uri is based on information in HTTP requests. | [optional] 

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

### GetHttpServerPortNumber

`func (o *NetworkConfiguration) GetHttpServerPortNumber() int32`

GetHttpServerPortNumber returns the HttpServerPortNumber field if non-nil, zero value otherwise.

### GetHttpServerPortNumberOk

`func (o *NetworkConfiguration) GetHttpServerPortNumberOk() (*int32, bool)`

GetHttpServerPortNumberOk returns a tuple with the HttpServerPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServerPortNumber

`func (o *NetworkConfiguration) SetHttpServerPortNumber(v int32)`

SetHttpServerPortNumber sets HttpServerPortNumber field to given value.

### HasHttpServerPortNumber

`func (o *NetworkConfiguration) HasHttpServerPortNumber() bool`

HasHttpServerPortNumber returns a boolean if a field has been set.

### GetHttpsPortNumber

`func (o *NetworkConfiguration) GetHttpsPortNumber() int32`

GetHttpsPortNumber returns the HttpsPortNumber field if non-nil, zero value otherwise.

### GetHttpsPortNumberOk

`func (o *NetworkConfiguration) GetHttpsPortNumberOk() (*int32, bool)`

GetHttpsPortNumberOk returns a tuple with the HttpsPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPortNumber

`func (o *NetworkConfiguration) SetHttpsPortNumber(v int32)`

SetHttpsPortNumber sets HttpsPortNumber field to given value.

### HasHttpsPortNumber

`func (o *NetworkConfiguration) HasHttpsPortNumber() bool`

HasHttpsPortNumber returns a boolean if a field has been set.

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

### GetPublicPort

`func (o *NetworkConfiguration) GetPublicPort() int32`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *NetworkConfiguration) GetPublicPortOk() (*int32, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *NetworkConfiguration) SetPublicPort(v int32)`

SetPublicPort sets PublicPort field to given value.

### HasPublicPort

`func (o *NetworkConfiguration) HasPublicPort() bool`

HasPublicPort returns a boolean if a field has been set.

### GetUPnPCreateHttpPortMap

`func (o *NetworkConfiguration) GetUPnPCreateHttpPortMap() bool`

GetUPnPCreateHttpPortMap returns the UPnPCreateHttpPortMap field if non-nil, zero value otherwise.

### GetUPnPCreateHttpPortMapOk

`func (o *NetworkConfiguration) GetUPnPCreateHttpPortMapOk() (*bool, bool)`

GetUPnPCreateHttpPortMapOk returns a tuple with the UPnPCreateHttpPortMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPnPCreateHttpPortMap

`func (o *NetworkConfiguration) SetUPnPCreateHttpPortMap(v bool)`

SetUPnPCreateHttpPortMap sets UPnPCreateHttpPortMap field to given value.

### HasUPnPCreateHttpPortMap

`func (o *NetworkConfiguration) HasUPnPCreateHttpPortMap() bool`

HasUPnPCreateHttpPortMap returns a boolean if a field has been set.

### GetUDPPortRange

`func (o *NetworkConfiguration) GetUDPPortRange() string`

GetUDPPortRange returns the UDPPortRange field if non-nil, zero value otherwise.

### GetUDPPortRangeOk

`func (o *NetworkConfiguration) GetUDPPortRangeOk() (*string, bool)`

GetUDPPortRangeOk returns a tuple with the UDPPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUDPPortRange

`func (o *NetworkConfiguration) SetUDPPortRange(v string)`

SetUDPPortRange sets UDPPortRange field to given value.

### HasUDPPortRange

`func (o *NetworkConfiguration) HasUDPPortRange() bool`

HasUDPPortRange returns a boolean if a field has been set.

### GetEnableIPV6

`func (o *NetworkConfiguration) GetEnableIPV6() bool`

GetEnableIPV6 returns the EnableIPV6 field if non-nil, zero value otherwise.

### GetEnableIPV6Ok

`func (o *NetworkConfiguration) GetEnableIPV6Ok() (*bool, bool)`

GetEnableIPV6Ok returns a tuple with the EnableIPV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIPV6

`func (o *NetworkConfiguration) SetEnableIPV6(v bool)`

SetEnableIPV6 sets EnableIPV6 field to given value.

### HasEnableIPV6

`func (o *NetworkConfiguration) HasEnableIPV6() bool`

HasEnableIPV6 returns a boolean if a field has been set.

### GetEnableIPV4

`func (o *NetworkConfiguration) GetEnableIPV4() bool`

GetEnableIPV4 returns the EnableIPV4 field if non-nil, zero value otherwise.

### GetEnableIPV4Ok

`func (o *NetworkConfiguration) GetEnableIPV4Ok() (*bool, bool)`

GetEnableIPV4Ok returns a tuple with the EnableIPV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIPV4

`func (o *NetworkConfiguration) SetEnableIPV4(v bool)`

SetEnableIPV4 sets EnableIPV4 field to given value.

### HasEnableIPV4

`func (o *NetworkConfiguration) HasEnableIPV4() bool`

HasEnableIPV4 returns a boolean if a field has been set.

### GetEnableSSDPTracing

`func (o *NetworkConfiguration) GetEnableSSDPTracing() bool`

GetEnableSSDPTracing returns the EnableSSDPTracing field if non-nil, zero value otherwise.

### GetEnableSSDPTracingOk

`func (o *NetworkConfiguration) GetEnableSSDPTracingOk() (*bool, bool)`

GetEnableSSDPTracingOk returns a tuple with the EnableSSDPTracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSSDPTracing

`func (o *NetworkConfiguration) SetEnableSSDPTracing(v bool)`

SetEnableSSDPTracing sets EnableSSDPTracing field to given value.

### HasEnableSSDPTracing

`func (o *NetworkConfiguration) HasEnableSSDPTracing() bool`

HasEnableSSDPTracing returns a boolean if a field has been set.

### GetSSDPTracingFilter

`func (o *NetworkConfiguration) GetSSDPTracingFilter() string`

GetSSDPTracingFilter returns the SSDPTracingFilter field if non-nil, zero value otherwise.

### GetSSDPTracingFilterOk

`func (o *NetworkConfiguration) GetSSDPTracingFilterOk() (*string, bool)`

GetSSDPTracingFilterOk returns a tuple with the SSDPTracingFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSSDPTracingFilter

`func (o *NetworkConfiguration) SetSSDPTracingFilter(v string)`

SetSSDPTracingFilter sets SSDPTracingFilter field to given value.

### HasSSDPTracingFilter

`func (o *NetworkConfiguration) HasSSDPTracingFilter() bool`

HasSSDPTracingFilter returns a boolean if a field has been set.

### GetUDPSendCount

`func (o *NetworkConfiguration) GetUDPSendCount() int32`

GetUDPSendCount returns the UDPSendCount field if non-nil, zero value otherwise.

### GetUDPSendCountOk

`func (o *NetworkConfiguration) GetUDPSendCountOk() (*int32, bool)`

GetUDPSendCountOk returns a tuple with the UDPSendCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUDPSendCount

`func (o *NetworkConfiguration) SetUDPSendCount(v int32)`

SetUDPSendCount sets UDPSendCount field to given value.

### HasUDPSendCount

`func (o *NetworkConfiguration) HasUDPSendCount() bool`

HasUDPSendCount returns a boolean if a field has been set.

### GetUDPSendDelay

`func (o *NetworkConfiguration) GetUDPSendDelay() int32`

GetUDPSendDelay returns the UDPSendDelay field if non-nil, zero value otherwise.

### GetUDPSendDelayOk

`func (o *NetworkConfiguration) GetUDPSendDelayOk() (*int32, bool)`

GetUDPSendDelayOk returns a tuple with the UDPSendDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUDPSendDelay

`func (o *NetworkConfiguration) SetUDPSendDelay(v int32)`

SetUDPSendDelay sets UDPSendDelay field to given value.

### HasUDPSendDelay

`func (o *NetworkConfiguration) HasUDPSendDelay() bool`

HasUDPSendDelay returns a boolean if a field has been set.

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

`func (o *NetworkConfiguration) GetVirtualInterfaceNames() string`

GetVirtualInterfaceNames returns the VirtualInterfaceNames field if non-nil, zero value otherwise.

### GetVirtualInterfaceNamesOk

`func (o *NetworkConfiguration) GetVirtualInterfaceNamesOk() (*string, bool)`

GetVirtualInterfaceNamesOk returns a tuple with the VirtualInterfaceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualInterfaceNames

`func (o *NetworkConfiguration) SetVirtualInterfaceNames(v string)`

SetVirtualInterfaceNames sets VirtualInterfaceNames field to given value.

### HasVirtualInterfaceNames

`func (o *NetworkConfiguration) HasVirtualInterfaceNames() bool`

HasVirtualInterfaceNames returns a boolean if a field has been set.

### GetGatewayMonitorPeriod

`func (o *NetworkConfiguration) GetGatewayMonitorPeriod() int32`

GetGatewayMonitorPeriod returns the GatewayMonitorPeriod field if non-nil, zero value otherwise.

### GetGatewayMonitorPeriodOk

`func (o *NetworkConfiguration) GetGatewayMonitorPeriodOk() (*int32, bool)`

GetGatewayMonitorPeriodOk returns a tuple with the GatewayMonitorPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayMonitorPeriod

`func (o *NetworkConfiguration) SetGatewayMonitorPeriod(v int32)`

SetGatewayMonitorPeriod sets GatewayMonitorPeriod field to given value.

### HasGatewayMonitorPeriod

`func (o *NetworkConfiguration) HasGatewayMonitorPeriod() bool`

HasGatewayMonitorPeriod returns a boolean if a field has been set.

### GetEnableMultiSocketBinding

`func (o *NetworkConfiguration) GetEnableMultiSocketBinding() bool`

GetEnableMultiSocketBinding returns the EnableMultiSocketBinding field if non-nil, zero value otherwise.

### GetEnableMultiSocketBindingOk

`func (o *NetworkConfiguration) GetEnableMultiSocketBindingOk() (*bool, bool)`

GetEnableMultiSocketBindingOk returns a tuple with the EnableMultiSocketBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMultiSocketBinding

`func (o *NetworkConfiguration) SetEnableMultiSocketBinding(v bool)`

SetEnableMultiSocketBinding sets EnableMultiSocketBinding field to given value.

### HasEnableMultiSocketBinding

`func (o *NetworkConfiguration) HasEnableMultiSocketBinding() bool`

HasEnableMultiSocketBinding returns a boolean if a field has been set.

### GetTrustAllIP6Interfaces

`func (o *NetworkConfiguration) GetTrustAllIP6Interfaces() bool`

GetTrustAllIP6Interfaces returns the TrustAllIP6Interfaces field if non-nil, zero value otherwise.

### GetTrustAllIP6InterfacesOk

`func (o *NetworkConfiguration) GetTrustAllIP6InterfacesOk() (*bool, bool)`

GetTrustAllIP6InterfacesOk returns a tuple with the TrustAllIP6Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAllIP6Interfaces

`func (o *NetworkConfiguration) SetTrustAllIP6Interfaces(v bool)`

SetTrustAllIP6Interfaces sets TrustAllIP6Interfaces field to given value.

### HasTrustAllIP6Interfaces

`func (o *NetworkConfiguration) HasTrustAllIP6Interfaces() bool`

HasTrustAllIP6Interfaces returns a boolean if a field has been set.

### GetHDHomerunPortRange

`func (o *NetworkConfiguration) GetHDHomerunPortRange() string`

GetHDHomerunPortRange returns the HDHomerunPortRange field if non-nil, zero value otherwise.

### GetHDHomerunPortRangeOk

`func (o *NetworkConfiguration) GetHDHomerunPortRangeOk() (*string, bool)`

GetHDHomerunPortRangeOk returns a tuple with the HDHomerunPortRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHDHomerunPortRange

`func (o *NetworkConfiguration) SetHDHomerunPortRange(v string)`

SetHDHomerunPortRange sets HDHomerunPortRange field to given value.

### HasHDHomerunPortRange

`func (o *NetworkConfiguration) HasHDHomerunPortRange() bool`

HasHDHomerunPortRange returns a boolean if a field has been set.

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

### GetAutoDiscoveryTracing

`func (o *NetworkConfiguration) GetAutoDiscoveryTracing() bool`

GetAutoDiscoveryTracing returns the AutoDiscoveryTracing field if non-nil, zero value otherwise.

### GetAutoDiscoveryTracingOk

`func (o *NetworkConfiguration) GetAutoDiscoveryTracingOk() (*bool, bool)`

GetAutoDiscoveryTracingOk returns a tuple with the AutoDiscoveryTracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDiscoveryTracing

`func (o *NetworkConfiguration) SetAutoDiscoveryTracing(v bool)`

SetAutoDiscoveryTracing sets AutoDiscoveryTracing field to given value.

### HasAutoDiscoveryTracing

`func (o *NetworkConfiguration) HasAutoDiscoveryTracing() bool`

HasAutoDiscoveryTracing returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


