# TimerInfoDtoProgramInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**OriginalTitle** | Pointer to **NullableString** |  | [optional] 
**ServerId** | Pointer to **NullableString** | Gets or sets the server identifier. | [optional] 
**Id** | Pointer to **string** | Gets or sets the id. | [optional] 
**Etag** | Pointer to **NullableString** | Gets or sets the etag. | [optional] 
**SourceType** | Pointer to **NullableString** | Gets or sets the type of the source. | [optional] 
**PlaylistItemId** | Pointer to **NullableString** | Gets or sets the playlist item identifier. | [optional] 
**DateCreated** | Pointer to **NullableTime** | Gets or sets the date created. | [optional] 
**DateLastMediaAdded** | Pointer to **NullableTime** |  | [optional] 
**ExtraType** | Pointer to [**NullableExtraType**](ExtraType.md) |  | [optional] 
**AirsBeforeSeasonNumber** | Pointer to **NullableInt32** |  | [optional] 
**AirsAfterSeasonNumber** | Pointer to **NullableInt32** |  | [optional] 
**AirsBeforeEpisodeNumber** | Pointer to **NullableInt32** |  | [optional] 
**CanDelete** | Pointer to **NullableBool** |  | [optional] 
**CanDownload** | Pointer to **NullableBool** |  | [optional] 
**HasLyrics** | Pointer to **NullableBool** |  | [optional] 
**HasSubtitles** | Pointer to **NullableBool** |  | [optional] 
**PreferredMetadataLanguage** | Pointer to **NullableString** |  | [optional] 
**PreferredMetadataCountryCode** | Pointer to **NullableString** |  | [optional] 
**Container** | Pointer to **NullableString** |  | [optional] 
**SortName** | Pointer to **NullableString** | Gets or sets the name of the sort. | [optional] 
**ForcedSortName** | Pointer to **NullableString** |  | [optional] 
**Video3DFormat** | Pointer to [**NullableVideo3DFormat**](Video3DFormat.md) | Gets or sets the video3 D format. | [optional] 
**PremiereDate** | Pointer to **NullableTime** | Gets or sets the premiere date. | [optional] 
**ExternalUrls** | Pointer to [**[]ExternalUrl**](ExternalUrl.md) | Gets or sets the external urls. | [optional] 
**MediaSources** | Pointer to [**[]MediaSourceInfo**](MediaSourceInfo.md) | Gets or sets the media versions. | [optional] 
**CriticRating** | Pointer to **NullableFloat32** | Gets or sets the critic rating. | [optional] 
**ProductionLocations** | Pointer to **[]string** |  | [optional] 
**Path** | Pointer to **NullableString** | Gets or sets the path. | [optional] 
**EnableMediaSourceDisplay** | Pointer to **NullableBool** |  | [optional] 
**OfficialRating** | Pointer to **NullableString** | Gets or sets the official rating. | [optional] 
**CustomRating** | Pointer to **NullableString** | Gets or sets the custom rating. | [optional] 
**ChannelId** | Pointer to **NullableString** | Gets or sets the channel identifier. | [optional] 
**ChannelName** | Pointer to **NullableString** |  | [optional] 
**Overview** | Pointer to **NullableString** | Gets or sets the overview. | [optional] 
**Taglines** | Pointer to **[]string** | Gets or sets the taglines. | [optional] 
**Genres** | Pointer to **[]string** | Gets or sets the genres. | [optional] 
**CommunityRating** | Pointer to **NullableFloat32** | Gets or sets the community rating. | [optional] 
**CumulativeRunTimeTicks** | Pointer to **NullableInt64** | Gets or sets the cumulative run time ticks. | [optional] 
**RunTimeTicks** | Pointer to **NullableInt64** | Gets or sets the run time ticks. | [optional] 
**PlayAccess** | Pointer to [**NullablePlayAccess**](PlayAccess.md) | Gets or sets the play access. | [optional] 
**AspectRatio** | Pointer to **NullableString** | Gets or sets the aspect ratio. | [optional] 
**ProductionYear** | Pointer to **NullableInt32** | Gets or sets the production year. | [optional] 
**IsPlaceHolder** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is place holder. | [optional] 
**Number** | Pointer to **NullableString** | Gets or sets the number. | [optional] 
**ChannelNumber** | Pointer to **NullableString** |  | [optional] 
**IndexNumber** | Pointer to **NullableInt32** | Gets or sets the index number. | [optional] 
**IndexNumberEnd** | Pointer to **NullableInt32** | Gets or sets the index number end. | [optional] 
**ParentIndexNumber** | Pointer to **NullableInt32** | Gets or sets the parent index number. | [optional] 
**RemoteTrailers** | Pointer to [**[]MediaUrl**](MediaUrl.md) | Gets or sets the trailer urls. | [optional] 
**ProviderIds** | Pointer to **map[string]string** | Gets or sets the provider ids. | [optional] 
**IsHD** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is HD. | [optional] 
**IsFolder** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is folder. | [optional] 
**ParentId** | Pointer to **NullableString** | Gets or sets the parent id. | [optional] 
**Type** | Pointer to [**BaseItemKind**](BaseItemKind.md) | The base item kind. | [optional] 
**People** | Pointer to [**[]BaseItemPerson**](BaseItemPerson.md) | Gets or sets the people. | [optional] 
**Studios** | Pointer to [**[]NameGuidPair**](NameGuidPair.md) | Gets or sets the studios. | [optional] 
**GenreItems** | Pointer to [**[]NameGuidPair**](NameGuidPair.md) |  | [optional] 
**ParentLogoItemId** | Pointer to **NullableString** | Gets or sets whether the item has a logo, this will hold the Id of the Parent that has one. | [optional] 
**ParentBackdropItemId** | Pointer to **NullableString** | Gets or sets whether the item has any backdrops, this will hold the Id of the Parent that has one. | [optional] 
**ParentBackdropImageTags** | Pointer to **[]string** | Gets or sets the parent backdrop image tags. | [optional] 
**LocalTrailerCount** | Pointer to **NullableInt32** | Gets or sets the local trailer count. | [optional] 
**UserData** | Pointer to [**NullableBaseItemDtoUserData**](BaseItemDtoUserData.md) |  | [optional] 
**RecursiveItemCount** | Pointer to **NullableInt32** | Gets or sets the recursive item count. | [optional] 
**ChildCount** | Pointer to **NullableInt32** | Gets or sets the child count. | [optional] 
**SeriesName** | Pointer to **NullableString** | Gets or sets the name of the series. | [optional] 
**SeriesId** | Pointer to **NullableString** | Gets or sets the series id. | [optional] 
**SeasonId** | Pointer to **NullableString** | Gets or sets the season identifier. | [optional] 
**SpecialFeatureCount** | Pointer to **NullableInt32** | Gets or sets the special feature count. | [optional] 
**DisplayPreferencesId** | Pointer to **NullableString** | Gets or sets the display preferences id. | [optional] 
**Status** | Pointer to **NullableString** | Gets or sets the status. | [optional] 
**AirTime** | Pointer to **NullableString** | Gets or sets the air time. | [optional] 
**AirDays** | Pointer to [**[]DayOfWeek**](DayOfWeek.md) | Gets or sets the air days. | [optional] 
**Tags** | Pointer to **[]string** | Gets or sets the tags. | [optional] 
**PrimaryImageAspectRatio** | Pointer to **NullableFloat64** | Gets or sets the primary image aspect ratio, after image enhancements. | [optional] 
**Artists** | Pointer to **[]string** | Gets or sets the artists. | [optional] 
**ArtistItems** | Pointer to [**[]NameGuidPair**](NameGuidPair.md) | Gets or sets the artist items. | [optional] 
**Album** | Pointer to **NullableString** | Gets or sets the album. | [optional] 
**CollectionType** | Pointer to [**NullableCollectionType**](CollectionType.md) | Gets or sets the type of the collection. | [optional] 
**DisplayOrder** | Pointer to **NullableString** | Gets or sets the display order. | [optional] 
**AlbumId** | Pointer to **NullableString** | Gets or sets the album id. | [optional] 
**AlbumPrimaryImageTag** | Pointer to **NullableString** | Gets or sets the album image tag. | [optional] 
**SeriesPrimaryImageTag** | Pointer to **NullableString** | Gets or sets the series primary image tag. | [optional] 
**AlbumArtist** | Pointer to **NullableString** | Gets or sets the album artist. | [optional] 
**AlbumArtists** | Pointer to [**[]NameGuidPair**](NameGuidPair.md) | Gets or sets the album artists. | [optional] 
**SeasonName** | Pointer to **NullableString** | Gets or sets the name of the season. | [optional] 
**MediaStreams** | Pointer to [**[]MediaStream**](MediaStream.md) | Gets or sets the media streams. | [optional] 
**VideoType** | Pointer to [**NullableVideoType**](VideoType.md) | Gets or sets the type of the video. | [optional] 
**PartCount** | Pointer to **NullableInt32** | Gets or sets the part count. | [optional] 
**MediaSourceCount** | Pointer to **NullableInt32** |  | [optional] 
**ImageTags** | Pointer to **map[string]string** | Gets or sets the image tags. | [optional] 
**BackdropImageTags** | Pointer to **[]string** | Gets or sets the backdrop image tags. | [optional] 
**ScreenshotImageTags** | Pointer to **[]string** | Gets or sets the screenshot image tags. | [optional] 
**ParentLogoImageTag** | Pointer to **NullableString** | Gets or sets the parent logo image tag. | [optional] 
**ParentArtItemId** | Pointer to **NullableString** | Gets or sets whether the item has fan art, this will hold the Id of the Parent that has one. | [optional] 
**ParentArtImageTag** | Pointer to **NullableString** | Gets or sets the parent art image tag. | [optional] 
**SeriesThumbImageTag** | Pointer to **NullableString** | Gets or sets the series thumb image tag. | [optional] 
**ImageBlurHashes** | Pointer to [**NullableBaseItemDtoImageBlurHashes**](BaseItemDtoImageBlurHashes.md) |  | [optional] 
**SeriesStudio** | Pointer to **NullableString** | Gets or sets the series studio. | [optional] 
**ParentThumbItemId** | Pointer to **NullableString** | Gets or sets the parent thumb item id. | [optional] 
**ParentThumbImageTag** | Pointer to **NullableString** | Gets or sets the parent thumb image tag. | [optional] 
**ParentPrimaryImageItemId** | Pointer to **NullableString** | Gets or sets the parent primary image item identifier. | [optional] 
**ParentPrimaryImageTag** | Pointer to **NullableString** | Gets or sets the parent primary image tag. | [optional] 
**Chapters** | Pointer to [**[]ChapterInfo**](ChapterInfo.md) | Gets or sets the chapters. | [optional] 
**Trickplay** | Pointer to [**map[string]map[string]TrickplayInfo**](map.md) | Gets or sets the trickplay manifest. | [optional] 
**LocationType** | Pointer to [**NullableLocationType**](LocationType.md) | Gets or sets the type of the location. | [optional] 
**IsoType** | Pointer to [**NullableIsoType**](IsoType.md) | Gets or sets the type of the iso. | [optional] 
**MediaType** | Pointer to [**MediaType**](MediaType.md) | Media types. | [optional] 
**EndDate** | Pointer to **NullableTime** | Gets or sets the end date. | [optional] 
**LockedFields** | Pointer to [**[]MetadataField**](MetadataField.md) | Gets or sets the locked fields. | [optional] 
**TrailerCount** | Pointer to **NullableInt32** | Gets or sets the trailer count. | [optional] 
**MovieCount** | Pointer to **NullableInt32** | Gets or sets the movie count. | [optional] 
**SeriesCount** | Pointer to **NullableInt32** | Gets or sets the series count. | [optional] 
**ProgramCount** | Pointer to **NullableInt32** |  | [optional] 
**EpisodeCount** | Pointer to **NullableInt32** | Gets or sets the episode count. | [optional] 
**SongCount** | Pointer to **NullableInt32** | Gets or sets the song count. | [optional] 
**AlbumCount** | Pointer to **NullableInt32** | Gets or sets the album count. | [optional] 
**ArtistCount** | Pointer to **NullableInt32** |  | [optional] 
**MusicVideoCount** | Pointer to **NullableInt32** | Gets or sets the music video count. | [optional] 
**LockData** | Pointer to **NullableBool** | Gets or sets a value indicating whether [enable internet providers]. | [optional] 
**Width** | Pointer to **NullableInt32** |  | [optional] 
**Height** | Pointer to **NullableInt32** |  | [optional] 
**CameraMake** | Pointer to **NullableString** |  | [optional] 
**CameraModel** | Pointer to **NullableString** |  | [optional] 
**Software** | Pointer to **NullableString** |  | [optional] 
**ExposureTime** | Pointer to **NullableFloat64** |  | [optional] 
**FocalLength** | Pointer to **NullableFloat64** |  | [optional] 
**ImageOrientation** | Pointer to [**NullableImageOrientation**](ImageOrientation.md) |  | [optional] 
**Aperture** | Pointer to **NullableFloat64** |  | [optional] 
**ShutterSpeed** | Pointer to **NullableFloat64** |  | [optional] 
**Latitude** | Pointer to **NullableFloat64** |  | [optional] 
**Longitude** | Pointer to **NullableFloat64** |  | [optional] 
**Altitude** | Pointer to **NullableFloat64** |  | [optional] 
**IsoSpeedRating** | Pointer to **NullableInt32** |  | [optional] 
**SeriesTimerId** | Pointer to **NullableString** | Gets or sets the series timer identifier. | [optional] 
**ProgramId** | Pointer to **NullableString** | Gets or sets the program identifier. | [optional] 
**ChannelPrimaryImageTag** | Pointer to **NullableString** | Gets or sets the channel primary image tag. | [optional] 
**StartDate** | Pointer to **NullableTime** | Gets or sets the start date of the recording, in UTC. | [optional] 
**CompletionPercentage** | Pointer to **NullableFloat64** | Gets or sets the completion percentage. | [optional] 
**IsRepeat** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is repeat. | [optional] 
**EpisodeTitle** | Pointer to **NullableString** | Gets or sets the episode title. | [optional] 
**ChannelType** | Pointer to [**NullableChannelType**](ChannelType.md) | Gets or sets the type of the channel. | [optional] 
**Audio** | Pointer to [**NullableProgramAudio**](ProgramAudio.md) | Gets or sets the audio. | [optional] 
**IsMovie** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is movie. | [optional] 
**IsSports** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is sports. | [optional] 
**IsSeries** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is series. | [optional] 
**IsLive** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is live. | [optional] 
**IsNews** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is news. | [optional] 
**IsKids** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is kids. | [optional] 
**IsPremiere** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is premiere. | [optional] 
**TimerId** | Pointer to **NullableString** | Gets or sets the timer identifier. | [optional] 
**NormalizationGain** | Pointer to **NullableFloat32** | Gets or sets the gain required for audio normalization. | [optional] 
**CurrentProgram** | Pointer to [**NullableBaseItemDtoCurrentProgram**](BaseItemDtoCurrentProgram.md) |  | [optional] 

## Methods

### NewTimerInfoDtoProgramInfo

`func NewTimerInfoDtoProgramInfo() *TimerInfoDtoProgramInfo`

NewTimerInfoDtoProgramInfo instantiates a new TimerInfoDtoProgramInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimerInfoDtoProgramInfoWithDefaults

`func NewTimerInfoDtoProgramInfoWithDefaults() *TimerInfoDtoProgramInfo`

NewTimerInfoDtoProgramInfoWithDefaults instantiates a new TimerInfoDtoProgramInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TimerInfoDtoProgramInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimerInfoDtoProgramInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimerInfoDtoProgramInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TimerInfoDtoProgramInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TimerInfoDtoProgramInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TimerInfoDtoProgramInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOriginalTitle

`func (o *TimerInfoDtoProgramInfo) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *TimerInfoDtoProgramInfo) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *TimerInfoDtoProgramInfo) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *TimerInfoDtoProgramInfo) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### SetOriginalTitleNil

`func (o *TimerInfoDtoProgramInfo) SetOriginalTitleNil(b bool)`

 SetOriginalTitleNil sets the value for OriginalTitle to be an explicit nil

### UnsetOriginalTitle
`func (o *TimerInfoDtoProgramInfo) UnsetOriginalTitle()`

UnsetOriginalTitle ensures that no value is present for OriginalTitle, not even an explicit nil
### GetServerId

`func (o *TimerInfoDtoProgramInfo) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *TimerInfoDtoProgramInfo) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *TimerInfoDtoProgramInfo) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *TimerInfoDtoProgramInfo) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *TimerInfoDtoProgramInfo) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *TimerInfoDtoProgramInfo) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetId

`func (o *TimerInfoDtoProgramInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimerInfoDtoProgramInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimerInfoDtoProgramInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TimerInfoDtoProgramInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEtag

`func (o *TimerInfoDtoProgramInfo) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *TimerInfoDtoProgramInfo) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *TimerInfoDtoProgramInfo) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *TimerInfoDtoProgramInfo) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### SetEtagNil

`func (o *TimerInfoDtoProgramInfo) SetEtagNil(b bool)`

 SetEtagNil sets the value for Etag to be an explicit nil

### UnsetEtag
`func (o *TimerInfoDtoProgramInfo) UnsetEtag()`

UnsetEtag ensures that no value is present for Etag, not even an explicit nil
### GetSourceType

`func (o *TimerInfoDtoProgramInfo) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *TimerInfoDtoProgramInfo) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *TimerInfoDtoProgramInfo) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *TimerInfoDtoProgramInfo) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *TimerInfoDtoProgramInfo) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *TimerInfoDtoProgramInfo) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetPlaylistItemId

`func (o *TimerInfoDtoProgramInfo) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *TimerInfoDtoProgramInfo) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *TimerInfoDtoProgramInfo) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *TimerInfoDtoProgramInfo) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### SetPlaylistItemIdNil

`func (o *TimerInfoDtoProgramInfo) SetPlaylistItemIdNil(b bool)`

 SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil

### UnsetPlaylistItemId
`func (o *TimerInfoDtoProgramInfo) UnsetPlaylistItemId()`

UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil
### GetDateCreated

`func (o *TimerInfoDtoProgramInfo) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TimerInfoDtoProgramInfo) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TimerInfoDtoProgramInfo) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TimerInfoDtoProgramInfo) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *TimerInfoDtoProgramInfo) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *TimerInfoDtoProgramInfo) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateLastMediaAdded

`func (o *TimerInfoDtoProgramInfo) GetDateLastMediaAdded() time.Time`

GetDateLastMediaAdded returns the DateLastMediaAdded field if non-nil, zero value otherwise.

### GetDateLastMediaAddedOk

`func (o *TimerInfoDtoProgramInfo) GetDateLastMediaAddedOk() (*time.Time, bool)`

GetDateLastMediaAddedOk returns a tuple with the DateLastMediaAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastMediaAdded

`func (o *TimerInfoDtoProgramInfo) SetDateLastMediaAdded(v time.Time)`

SetDateLastMediaAdded sets DateLastMediaAdded field to given value.

### HasDateLastMediaAdded

`func (o *TimerInfoDtoProgramInfo) HasDateLastMediaAdded() bool`

HasDateLastMediaAdded returns a boolean if a field has been set.

### SetDateLastMediaAddedNil

`func (o *TimerInfoDtoProgramInfo) SetDateLastMediaAddedNil(b bool)`

 SetDateLastMediaAddedNil sets the value for DateLastMediaAdded to be an explicit nil

### UnsetDateLastMediaAdded
`func (o *TimerInfoDtoProgramInfo) UnsetDateLastMediaAdded()`

UnsetDateLastMediaAdded ensures that no value is present for DateLastMediaAdded, not even an explicit nil
### GetExtraType

`func (o *TimerInfoDtoProgramInfo) GetExtraType() ExtraType`

GetExtraType returns the ExtraType field if non-nil, zero value otherwise.

### GetExtraTypeOk

`func (o *TimerInfoDtoProgramInfo) GetExtraTypeOk() (*ExtraType, bool)`

GetExtraTypeOk returns a tuple with the ExtraType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraType

`func (o *TimerInfoDtoProgramInfo) SetExtraType(v ExtraType)`

SetExtraType sets ExtraType field to given value.

### HasExtraType

`func (o *TimerInfoDtoProgramInfo) HasExtraType() bool`

HasExtraType returns a boolean if a field has been set.

### SetExtraTypeNil

`func (o *TimerInfoDtoProgramInfo) SetExtraTypeNil(b bool)`

 SetExtraTypeNil sets the value for ExtraType to be an explicit nil

### UnsetExtraType
`func (o *TimerInfoDtoProgramInfo) UnsetExtraType()`

UnsetExtraType ensures that no value is present for ExtraType, not even an explicit nil
### GetAirsBeforeSeasonNumber

`func (o *TimerInfoDtoProgramInfo) GetAirsBeforeSeasonNumber() int32`

GetAirsBeforeSeasonNumber returns the AirsBeforeSeasonNumber field if non-nil, zero value otherwise.

### GetAirsBeforeSeasonNumberOk

`func (o *TimerInfoDtoProgramInfo) GetAirsBeforeSeasonNumberOk() (*int32, bool)`

GetAirsBeforeSeasonNumberOk returns a tuple with the AirsBeforeSeasonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirsBeforeSeasonNumber

`func (o *TimerInfoDtoProgramInfo) SetAirsBeforeSeasonNumber(v int32)`

SetAirsBeforeSeasonNumber sets AirsBeforeSeasonNumber field to given value.

### HasAirsBeforeSeasonNumber

`func (o *TimerInfoDtoProgramInfo) HasAirsBeforeSeasonNumber() bool`

HasAirsBeforeSeasonNumber returns a boolean if a field has been set.

### SetAirsBeforeSeasonNumberNil

`func (o *TimerInfoDtoProgramInfo) SetAirsBeforeSeasonNumberNil(b bool)`

 SetAirsBeforeSeasonNumberNil sets the value for AirsBeforeSeasonNumber to be an explicit nil

### UnsetAirsBeforeSeasonNumber
`func (o *TimerInfoDtoProgramInfo) UnsetAirsBeforeSeasonNumber()`

UnsetAirsBeforeSeasonNumber ensures that no value is present for AirsBeforeSeasonNumber, not even an explicit nil
### GetAirsAfterSeasonNumber

`func (o *TimerInfoDtoProgramInfo) GetAirsAfterSeasonNumber() int32`

GetAirsAfterSeasonNumber returns the AirsAfterSeasonNumber field if non-nil, zero value otherwise.

### GetAirsAfterSeasonNumberOk

`func (o *TimerInfoDtoProgramInfo) GetAirsAfterSeasonNumberOk() (*int32, bool)`

GetAirsAfterSeasonNumberOk returns a tuple with the AirsAfterSeasonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirsAfterSeasonNumber

`func (o *TimerInfoDtoProgramInfo) SetAirsAfterSeasonNumber(v int32)`

SetAirsAfterSeasonNumber sets AirsAfterSeasonNumber field to given value.

### HasAirsAfterSeasonNumber

`func (o *TimerInfoDtoProgramInfo) HasAirsAfterSeasonNumber() bool`

HasAirsAfterSeasonNumber returns a boolean if a field has been set.

### SetAirsAfterSeasonNumberNil

`func (o *TimerInfoDtoProgramInfo) SetAirsAfterSeasonNumberNil(b bool)`

 SetAirsAfterSeasonNumberNil sets the value for AirsAfterSeasonNumber to be an explicit nil

### UnsetAirsAfterSeasonNumber
`func (o *TimerInfoDtoProgramInfo) UnsetAirsAfterSeasonNumber()`

UnsetAirsAfterSeasonNumber ensures that no value is present for AirsAfterSeasonNumber, not even an explicit nil
### GetAirsBeforeEpisodeNumber

`func (o *TimerInfoDtoProgramInfo) GetAirsBeforeEpisodeNumber() int32`

GetAirsBeforeEpisodeNumber returns the AirsBeforeEpisodeNumber field if non-nil, zero value otherwise.

### GetAirsBeforeEpisodeNumberOk

`func (o *TimerInfoDtoProgramInfo) GetAirsBeforeEpisodeNumberOk() (*int32, bool)`

GetAirsBeforeEpisodeNumberOk returns a tuple with the AirsBeforeEpisodeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirsBeforeEpisodeNumber

`func (o *TimerInfoDtoProgramInfo) SetAirsBeforeEpisodeNumber(v int32)`

SetAirsBeforeEpisodeNumber sets AirsBeforeEpisodeNumber field to given value.

### HasAirsBeforeEpisodeNumber

`func (o *TimerInfoDtoProgramInfo) HasAirsBeforeEpisodeNumber() bool`

HasAirsBeforeEpisodeNumber returns a boolean if a field has been set.

### SetAirsBeforeEpisodeNumberNil

`func (o *TimerInfoDtoProgramInfo) SetAirsBeforeEpisodeNumberNil(b bool)`

 SetAirsBeforeEpisodeNumberNil sets the value for AirsBeforeEpisodeNumber to be an explicit nil

### UnsetAirsBeforeEpisodeNumber
`func (o *TimerInfoDtoProgramInfo) UnsetAirsBeforeEpisodeNumber()`

UnsetAirsBeforeEpisodeNumber ensures that no value is present for AirsBeforeEpisodeNumber, not even an explicit nil
### GetCanDelete

`func (o *TimerInfoDtoProgramInfo) GetCanDelete() bool`

GetCanDelete returns the CanDelete field if non-nil, zero value otherwise.

### GetCanDeleteOk

`func (o *TimerInfoDtoProgramInfo) GetCanDeleteOk() (*bool, bool)`

GetCanDeleteOk returns a tuple with the CanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelete

`func (o *TimerInfoDtoProgramInfo) SetCanDelete(v bool)`

SetCanDelete sets CanDelete field to given value.

### HasCanDelete

`func (o *TimerInfoDtoProgramInfo) HasCanDelete() bool`

HasCanDelete returns a boolean if a field has been set.

### SetCanDeleteNil

`func (o *TimerInfoDtoProgramInfo) SetCanDeleteNil(b bool)`

 SetCanDeleteNil sets the value for CanDelete to be an explicit nil

### UnsetCanDelete
`func (o *TimerInfoDtoProgramInfo) UnsetCanDelete()`

UnsetCanDelete ensures that no value is present for CanDelete, not even an explicit nil
### GetCanDownload

`func (o *TimerInfoDtoProgramInfo) GetCanDownload() bool`

GetCanDownload returns the CanDownload field if non-nil, zero value otherwise.

### GetCanDownloadOk

`func (o *TimerInfoDtoProgramInfo) GetCanDownloadOk() (*bool, bool)`

GetCanDownloadOk returns a tuple with the CanDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDownload

`func (o *TimerInfoDtoProgramInfo) SetCanDownload(v bool)`

SetCanDownload sets CanDownload field to given value.

### HasCanDownload

`func (o *TimerInfoDtoProgramInfo) HasCanDownload() bool`

HasCanDownload returns a boolean if a field has been set.

### SetCanDownloadNil

`func (o *TimerInfoDtoProgramInfo) SetCanDownloadNil(b bool)`

 SetCanDownloadNil sets the value for CanDownload to be an explicit nil

### UnsetCanDownload
`func (o *TimerInfoDtoProgramInfo) UnsetCanDownload()`

UnsetCanDownload ensures that no value is present for CanDownload, not even an explicit nil
### GetHasLyrics

`func (o *TimerInfoDtoProgramInfo) GetHasLyrics() bool`

GetHasLyrics returns the HasLyrics field if non-nil, zero value otherwise.

### GetHasLyricsOk

`func (o *TimerInfoDtoProgramInfo) GetHasLyricsOk() (*bool, bool)`

GetHasLyricsOk returns a tuple with the HasLyrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLyrics

`func (o *TimerInfoDtoProgramInfo) SetHasLyrics(v bool)`

SetHasLyrics sets HasLyrics field to given value.

### HasHasLyrics

`func (o *TimerInfoDtoProgramInfo) HasHasLyrics() bool`

HasHasLyrics returns a boolean if a field has been set.

### SetHasLyricsNil

`func (o *TimerInfoDtoProgramInfo) SetHasLyricsNil(b bool)`

 SetHasLyricsNil sets the value for HasLyrics to be an explicit nil

### UnsetHasLyrics
`func (o *TimerInfoDtoProgramInfo) UnsetHasLyrics()`

UnsetHasLyrics ensures that no value is present for HasLyrics, not even an explicit nil
### GetHasSubtitles

`func (o *TimerInfoDtoProgramInfo) GetHasSubtitles() bool`

GetHasSubtitles returns the HasSubtitles field if non-nil, zero value otherwise.

### GetHasSubtitlesOk

`func (o *TimerInfoDtoProgramInfo) GetHasSubtitlesOk() (*bool, bool)`

GetHasSubtitlesOk returns a tuple with the HasSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSubtitles

`func (o *TimerInfoDtoProgramInfo) SetHasSubtitles(v bool)`

SetHasSubtitles sets HasSubtitles field to given value.

### HasHasSubtitles

`func (o *TimerInfoDtoProgramInfo) HasHasSubtitles() bool`

HasHasSubtitles returns a boolean if a field has been set.

### SetHasSubtitlesNil

`func (o *TimerInfoDtoProgramInfo) SetHasSubtitlesNil(b bool)`

 SetHasSubtitlesNil sets the value for HasSubtitles to be an explicit nil

### UnsetHasSubtitles
`func (o *TimerInfoDtoProgramInfo) UnsetHasSubtitles()`

UnsetHasSubtitles ensures that no value is present for HasSubtitles, not even an explicit nil
### GetPreferredMetadataLanguage

`func (o *TimerInfoDtoProgramInfo) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *TimerInfoDtoProgramInfo) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *TimerInfoDtoProgramInfo) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *TimerInfoDtoProgramInfo) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### SetPreferredMetadataLanguageNil

`func (o *TimerInfoDtoProgramInfo) SetPreferredMetadataLanguageNil(b bool)`

 SetPreferredMetadataLanguageNil sets the value for PreferredMetadataLanguage to be an explicit nil

### UnsetPreferredMetadataLanguage
`func (o *TimerInfoDtoProgramInfo) UnsetPreferredMetadataLanguage()`

UnsetPreferredMetadataLanguage ensures that no value is present for PreferredMetadataLanguage, not even an explicit nil
### GetPreferredMetadataCountryCode

`func (o *TimerInfoDtoProgramInfo) GetPreferredMetadataCountryCode() string`

GetPreferredMetadataCountryCode returns the PreferredMetadataCountryCode field if non-nil, zero value otherwise.

### GetPreferredMetadataCountryCodeOk

`func (o *TimerInfoDtoProgramInfo) GetPreferredMetadataCountryCodeOk() (*string, bool)`

GetPreferredMetadataCountryCodeOk returns a tuple with the PreferredMetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataCountryCode

`func (o *TimerInfoDtoProgramInfo) SetPreferredMetadataCountryCode(v string)`

SetPreferredMetadataCountryCode sets PreferredMetadataCountryCode field to given value.

### HasPreferredMetadataCountryCode

`func (o *TimerInfoDtoProgramInfo) HasPreferredMetadataCountryCode() bool`

HasPreferredMetadataCountryCode returns a boolean if a field has been set.

### SetPreferredMetadataCountryCodeNil

`func (o *TimerInfoDtoProgramInfo) SetPreferredMetadataCountryCodeNil(b bool)`

 SetPreferredMetadataCountryCodeNil sets the value for PreferredMetadataCountryCode to be an explicit nil

### UnsetPreferredMetadataCountryCode
`func (o *TimerInfoDtoProgramInfo) UnsetPreferredMetadataCountryCode()`

UnsetPreferredMetadataCountryCode ensures that no value is present for PreferredMetadataCountryCode, not even an explicit nil
### GetContainer

`func (o *TimerInfoDtoProgramInfo) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *TimerInfoDtoProgramInfo) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *TimerInfoDtoProgramInfo) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *TimerInfoDtoProgramInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *TimerInfoDtoProgramInfo) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *TimerInfoDtoProgramInfo) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetSortName

`func (o *TimerInfoDtoProgramInfo) GetSortName() string`

GetSortName returns the SortName field if non-nil, zero value otherwise.

### GetSortNameOk

`func (o *TimerInfoDtoProgramInfo) GetSortNameOk() (*string, bool)`

GetSortNameOk returns a tuple with the SortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortName

`func (o *TimerInfoDtoProgramInfo) SetSortName(v string)`

SetSortName sets SortName field to given value.

### HasSortName

`func (o *TimerInfoDtoProgramInfo) HasSortName() bool`

HasSortName returns a boolean if a field has been set.

### SetSortNameNil

`func (o *TimerInfoDtoProgramInfo) SetSortNameNil(b bool)`

 SetSortNameNil sets the value for SortName to be an explicit nil

### UnsetSortName
`func (o *TimerInfoDtoProgramInfo) UnsetSortName()`

UnsetSortName ensures that no value is present for SortName, not even an explicit nil
### GetForcedSortName

`func (o *TimerInfoDtoProgramInfo) GetForcedSortName() string`

GetForcedSortName returns the ForcedSortName field if non-nil, zero value otherwise.

### GetForcedSortNameOk

`func (o *TimerInfoDtoProgramInfo) GetForcedSortNameOk() (*string, bool)`

GetForcedSortNameOk returns a tuple with the ForcedSortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSortName

`func (o *TimerInfoDtoProgramInfo) SetForcedSortName(v string)`

SetForcedSortName sets ForcedSortName field to given value.

### HasForcedSortName

`func (o *TimerInfoDtoProgramInfo) HasForcedSortName() bool`

HasForcedSortName returns a boolean if a field has been set.

### SetForcedSortNameNil

`func (o *TimerInfoDtoProgramInfo) SetForcedSortNameNil(b bool)`

 SetForcedSortNameNil sets the value for ForcedSortName to be an explicit nil

### UnsetForcedSortName
`func (o *TimerInfoDtoProgramInfo) UnsetForcedSortName()`

UnsetForcedSortName ensures that no value is present for ForcedSortName, not even an explicit nil
### GetVideo3DFormat

`func (o *TimerInfoDtoProgramInfo) GetVideo3DFormat() Video3DFormat`

GetVideo3DFormat returns the Video3DFormat field if non-nil, zero value otherwise.

### GetVideo3DFormatOk

`func (o *TimerInfoDtoProgramInfo) GetVideo3DFormatOk() (*Video3DFormat, bool)`

GetVideo3DFormatOk returns a tuple with the Video3DFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo3DFormat

`func (o *TimerInfoDtoProgramInfo) SetVideo3DFormat(v Video3DFormat)`

SetVideo3DFormat sets Video3DFormat field to given value.

### HasVideo3DFormat

`func (o *TimerInfoDtoProgramInfo) HasVideo3DFormat() bool`

HasVideo3DFormat returns a boolean if a field has been set.

### SetVideo3DFormatNil

`func (o *TimerInfoDtoProgramInfo) SetVideo3DFormatNil(b bool)`

 SetVideo3DFormatNil sets the value for Video3DFormat to be an explicit nil

### UnsetVideo3DFormat
`func (o *TimerInfoDtoProgramInfo) UnsetVideo3DFormat()`

UnsetVideo3DFormat ensures that no value is present for Video3DFormat, not even an explicit nil
### GetPremiereDate

`func (o *TimerInfoDtoProgramInfo) GetPremiereDate() time.Time`

GetPremiereDate returns the PremiereDate field if non-nil, zero value otherwise.

### GetPremiereDateOk

`func (o *TimerInfoDtoProgramInfo) GetPremiereDateOk() (*time.Time, bool)`

GetPremiereDateOk returns a tuple with the PremiereDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiereDate

`func (o *TimerInfoDtoProgramInfo) SetPremiereDate(v time.Time)`

SetPremiereDate sets PremiereDate field to given value.

### HasPremiereDate

`func (o *TimerInfoDtoProgramInfo) HasPremiereDate() bool`

HasPremiereDate returns a boolean if a field has been set.

### SetPremiereDateNil

`func (o *TimerInfoDtoProgramInfo) SetPremiereDateNil(b bool)`

 SetPremiereDateNil sets the value for PremiereDate to be an explicit nil

### UnsetPremiereDate
`func (o *TimerInfoDtoProgramInfo) UnsetPremiereDate()`

UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
### GetExternalUrls

`func (o *TimerInfoDtoProgramInfo) GetExternalUrls() []ExternalUrl`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *TimerInfoDtoProgramInfo) GetExternalUrlsOk() (*[]ExternalUrl, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *TimerInfoDtoProgramInfo) SetExternalUrls(v []ExternalUrl)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *TimerInfoDtoProgramInfo) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### SetExternalUrlsNil

`func (o *TimerInfoDtoProgramInfo) SetExternalUrlsNil(b bool)`

 SetExternalUrlsNil sets the value for ExternalUrls to be an explicit nil

### UnsetExternalUrls
`func (o *TimerInfoDtoProgramInfo) UnsetExternalUrls()`

UnsetExternalUrls ensures that no value is present for ExternalUrls, not even an explicit nil
### GetMediaSources

`func (o *TimerInfoDtoProgramInfo) GetMediaSources() []MediaSourceInfo`

GetMediaSources returns the MediaSources field if non-nil, zero value otherwise.

### GetMediaSourcesOk

`func (o *TimerInfoDtoProgramInfo) GetMediaSourcesOk() (*[]MediaSourceInfo, bool)`

GetMediaSourcesOk returns a tuple with the MediaSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSources

`func (o *TimerInfoDtoProgramInfo) SetMediaSources(v []MediaSourceInfo)`

SetMediaSources sets MediaSources field to given value.

### HasMediaSources

`func (o *TimerInfoDtoProgramInfo) HasMediaSources() bool`

HasMediaSources returns a boolean if a field has been set.

### SetMediaSourcesNil

`func (o *TimerInfoDtoProgramInfo) SetMediaSourcesNil(b bool)`

 SetMediaSourcesNil sets the value for MediaSources to be an explicit nil

### UnsetMediaSources
`func (o *TimerInfoDtoProgramInfo) UnsetMediaSources()`

UnsetMediaSources ensures that no value is present for MediaSources, not even an explicit nil
### GetCriticRating

`func (o *TimerInfoDtoProgramInfo) GetCriticRating() float32`

GetCriticRating returns the CriticRating field if non-nil, zero value otherwise.

### GetCriticRatingOk

`func (o *TimerInfoDtoProgramInfo) GetCriticRatingOk() (*float32, bool)`

GetCriticRatingOk returns a tuple with the CriticRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticRating

`func (o *TimerInfoDtoProgramInfo) SetCriticRating(v float32)`

SetCriticRating sets CriticRating field to given value.

### HasCriticRating

`func (o *TimerInfoDtoProgramInfo) HasCriticRating() bool`

HasCriticRating returns a boolean if a field has been set.

### SetCriticRatingNil

`func (o *TimerInfoDtoProgramInfo) SetCriticRatingNil(b bool)`

 SetCriticRatingNil sets the value for CriticRating to be an explicit nil

### UnsetCriticRating
`func (o *TimerInfoDtoProgramInfo) UnsetCriticRating()`

UnsetCriticRating ensures that no value is present for CriticRating, not even an explicit nil
### GetProductionLocations

`func (o *TimerInfoDtoProgramInfo) GetProductionLocations() []string`

GetProductionLocations returns the ProductionLocations field if non-nil, zero value otherwise.

### GetProductionLocationsOk

`func (o *TimerInfoDtoProgramInfo) GetProductionLocationsOk() (*[]string, bool)`

GetProductionLocationsOk returns a tuple with the ProductionLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionLocations

`func (o *TimerInfoDtoProgramInfo) SetProductionLocations(v []string)`

SetProductionLocations sets ProductionLocations field to given value.

### HasProductionLocations

`func (o *TimerInfoDtoProgramInfo) HasProductionLocations() bool`

HasProductionLocations returns a boolean if a field has been set.

### SetProductionLocationsNil

`func (o *TimerInfoDtoProgramInfo) SetProductionLocationsNil(b bool)`

 SetProductionLocationsNil sets the value for ProductionLocations to be an explicit nil

### UnsetProductionLocations
`func (o *TimerInfoDtoProgramInfo) UnsetProductionLocations()`

UnsetProductionLocations ensures that no value is present for ProductionLocations, not even an explicit nil
### GetPath

`func (o *TimerInfoDtoProgramInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TimerInfoDtoProgramInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TimerInfoDtoProgramInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TimerInfoDtoProgramInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *TimerInfoDtoProgramInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *TimerInfoDtoProgramInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetEnableMediaSourceDisplay

`func (o *TimerInfoDtoProgramInfo) GetEnableMediaSourceDisplay() bool`

GetEnableMediaSourceDisplay returns the EnableMediaSourceDisplay field if non-nil, zero value otherwise.

### GetEnableMediaSourceDisplayOk

`func (o *TimerInfoDtoProgramInfo) GetEnableMediaSourceDisplayOk() (*bool, bool)`

GetEnableMediaSourceDisplayOk returns a tuple with the EnableMediaSourceDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMediaSourceDisplay

`func (o *TimerInfoDtoProgramInfo) SetEnableMediaSourceDisplay(v bool)`

SetEnableMediaSourceDisplay sets EnableMediaSourceDisplay field to given value.

### HasEnableMediaSourceDisplay

`func (o *TimerInfoDtoProgramInfo) HasEnableMediaSourceDisplay() bool`

HasEnableMediaSourceDisplay returns a boolean if a field has been set.

### SetEnableMediaSourceDisplayNil

`func (o *TimerInfoDtoProgramInfo) SetEnableMediaSourceDisplayNil(b bool)`

 SetEnableMediaSourceDisplayNil sets the value for EnableMediaSourceDisplay to be an explicit nil

### UnsetEnableMediaSourceDisplay
`func (o *TimerInfoDtoProgramInfo) UnsetEnableMediaSourceDisplay()`

UnsetEnableMediaSourceDisplay ensures that no value is present for EnableMediaSourceDisplay, not even an explicit nil
### GetOfficialRating

`func (o *TimerInfoDtoProgramInfo) GetOfficialRating() string`

GetOfficialRating returns the OfficialRating field if non-nil, zero value otherwise.

### GetOfficialRatingOk

`func (o *TimerInfoDtoProgramInfo) GetOfficialRatingOk() (*string, bool)`

GetOfficialRatingOk returns a tuple with the OfficialRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialRating

`func (o *TimerInfoDtoProgramInfo) SetOfficialRating(v string)`

SetOfficialRating sets OfficialRating field to given value.

### HasOfficialRating

`func (o *TimerInfoDtoProgramInfo) HasOfficialRating() bool`

HasOfficialRating returns a boolean if a field has been set.

### SetOfficialRatingNil

`func (o *TimerInfoDtoProgramInfo) SetOfficialRatingNil(b bool)`

 SetOfficialRatingNil sets the value for OfficialRating to be an explicit nil

### UnsetOfficialRating
`func (o *TimerInfoDtoProgramInfo) UnsetOfficialRating()`

UnsetOfficialRating ensures that no value is present for OfficialRating, not even an explicit nil
### GetCustomRating

`func (o *TimerInfoDtoProgramInfo) GetCustomRating() string`

GetCustomRating returns the CustomRating field if non-nil, zero value otherwise.

### GetCustomRatingOk

`func (o *TimerInfoDtoProgramInfo) GetCustomRatingOk() (*string, bool)`

GetCustomRatingOk returns a tuple with the CustomRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRating

`func (o *TimerInfoDtoProgramInfo) SetCustomRating(v string)`

SetCustomRating sets CustomRating field to given value.

### HasCustomRating

`func (o *TimerInfoDtoProgramInfo) HasCustomRating() bool`

HasCustomRating returns a boolean if a field has been set.

### SetCustomRatingNil

`func (o *TimerInfoDtoProgramInfo) SetCustomRatingNil(b bool)`

 SetCustomRatingNil sets the value for CustomRating to be an explicit nil

### UnsetCustomRating
`func (o *TimerInfoDtoProgramInfo) UnsetCustomRating()`

UnsetCustomRating ensures that no value is present for CustomRating, not even an explicit nil
### GetChannelId

`func (o *TimerInfoDtoProgramInfo) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *TimerInfoDtoProgramInfo) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *TimerInfoDtoProgramInfo) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *TimerInfoDtoProgramInfo) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *TimerInfoDtoProgramInfo) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *TimerInfoDtoProgramInfo) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetChannelName

`func (o *TimerInfoDtoProgramInfo) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *TimerInfoDtoProgramInfo) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *TimerInfoDtoProgramInfo) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *TimerInfoDtoProgramInfo) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### SetChannelNameNil

`func (o *TimerInfoDtoProgramInfo) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *TimerInfoDtoProgramInfo) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetOverview

`func (o *TimerInfoDtoProgramInfo) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *TimerInfoDtoProgramInfo) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *TimerInfoDtoProgramInfo) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *TimerInfoDtoProgramInfo) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *TimerInfoDtoProgramInfo) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *TimerInfoDtoProgramInfo) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetTaglines

`func (o *TimerInfoDtoProgramInfo) GetTaglines() []string`

GetTaglines returns the Taglines field if non-nil, zero value otherwise.

### GetTaglinesOk

`func (o *TimerInfoDtoProgramInfo) GetTaglinesOk() (*[]string, bool)`

GetTaglinesOk returns a tuple with the Taglines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaglines

`func (o *TimerInfoDtoProgramInfo) SetTaglines(v []string)`

SetTaglines sets Taglines field to given value.

### HasTaglines

`func (o *TimerInfoDtoProgramInfo) HasTaglines() bool`

HasTaglines returns a boolean if a field has been set.

### SetTaglinesNil

`func (o *TimerInfoDtoProgramInfo) SetTaglinesNil(b bool)`

 SetTaglinesNil sets the value for Taglines to be an explicit nil

### UnsetTaglines
`func (o *TimerInfoDtoProgramInfo) UnsetTaglines()`

UnsetTaglines ensures that no value is present for Taglines, not even an explicit nil
### GetGenres

`func (o *TimerInfoDtoProgramInfo) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *TimerInfoDtoProgramInfo) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *TimerInfoDtoProgramInfo) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *TimerInfoDtoProgramInfo) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *TimerInfoDtoProgramInfo) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *TimerInfoDtoProgramInfo) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetCommunityRating

`func (o *TimerInfoDtoProgramInfo) GetCommunityRating() float32`

GetCommunityRating returns the CommunityRating field if non-nil, zero value otherwise.

### GetCommunityRatingOk

`func (o *TimerInfoDtoProgramInfo) GetCommunityRatingOk() (*float32, bool)`

GetCommunityRatingOk returns a tuple with the CommunityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityRating

`func (o *TimerInfoDtoProgramInfo) SetCommunityRating(v float32)`

SetCommunityRating sets CommunityRating field to given value.

### HasCommunityRating

`func (o *TimerInfoDtoProgramInfo) HasCommunityRating() bool`

HasCommunityRating returns a boolean if a field has been set.

### SetCommunityRatingNil

`func (o *TimerInfoDtoProgramInfo) SetCommunityRatingNil(b bool)`

 SetCommunityRatingNil sets the value for CommunityRating to be an explicit nil

### UnsetCommunityRating
`func (o *TimerInfoDtoProgramInfo) UnsetCommunityRating()`

UnsetCommunityRating ensures that no value is present for CommunityRating, not even an explicit nil
### GetCumulativeRunTimeTicks

`func (o *TimerInfoDtoProgramInfo) GetCumulativeRunTimeTicks() int64`

GetCumulativeRunTimeTicks returns the CumulativeRunTimeTicks field if non-nil, zero value otherwise.

### GetCumulativeRunTimeTicksOk

`func (o *TimerInfoDtoProgramInfo) GetCumulativeRunTimeTicksOk() (*int64, bool)`

GetCumulativeRunTimeTicksOk returns a tuple with the CumulativeRunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeRunTimeTicks

`func (o *TimerInfoDtoProgramInfo) SetCumulativeRunTimeTicks(v int64)`

SetCumulativeRunTimeTicks sets CumulativeRunTimeTicks field to given value.

### HasCumulativeRunTimeTicks

`func (o *TimerInfoDtoProgramInfo) HasCumulativeRunTimeTicks() bool`

HasCumulativeRunTimeTicks returns a boolean if a field has been set.

### SetCumulativeRunTimeTicksNil

`func (o *TimerInfoDtoProgramInfo) SetCumulativeRunTimeTicksNil(b bool)`

 SetCumulativeRunTimeTicksNil sets the value for CumulativeRunTimeTicks to be an explicit nil

### UnsetCumulativeRunTimeTicks
`func (o *TimerInfoDtoProgramInfo) UnsetCumulativeRunTimeTicks()`

UnsetCumulativeRunTimeTicks ensures that no value is present for CumulativeRunTimeTicks, not even an explicit nil
### GetRunTimeTicks

`func (o *TimerInfoDtoProgramInfo) GetRunTimeTicks() int64`

GetRunTimeTicks returns the RunTimeTicks field if non-nil, zero value otherwise.

### GetRunTimeTicksOk

`func (o *TimerInfoDtoProgramInfo) GetRunTimeTicksOk() (*int64, bool)`

GetRunTimeTicksOk returns a tuple with the RunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeTicks

`func (o *TimerInfoDtoProgramInfo) SetRunTimeTicks(v int64)`

SetRunTimeTicks sets RunTimeTicks field to given value.

### HasRunTimeTicks

`func (o *TimerInfoDtoProgramInfo) HasRunTimeTicks() bool`

HasRunTimeTicks returns a boolean if a field has been set.

### SetRunTimeTicksNil

`func (o *TimerInfoDtoProgramInfo) SetRunTimeTicksNil(b bool)`

 SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil

### UnsetRunTimeTicks
`func (o *TimerInfoDtoProgramInfo) UnsetRunTimeTicks()`

UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
### GetPlayAccess

`func (o *TimerInfoDtoProgramInfo) GetPlayAccess() PlayAccess`

GetPlayAccess returns the PlayAccess field if non-nil, zero value otherwise.

### GetPlayAccessOk

`func (o *TimerInfoDtoProgramInfo) GetPlayAccessOk() (*PlayAccess, bool)`

GetPlayAccessOk returns a tuple with the PlayAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayAccess

`func (o *TimerInfoDtoProgramInfo) SetPlayAccess(v PlayAccess)`

SetPlayAccess sets PlayAccess field to given value.

### HasPlayAccess

`func (o *TimerInfoDtoProgramInfo) HasPlayAccess() bool`

HasPlayAccess returns a boolean if a field has been set.

### SetPlayAccessNil

`func (o *TimerInfoDtoProgramInfo) SetPlayAccessNil(b bool)`

 SetPlayAccessNil sets the value for PlayAccess to be an explicit nil

### UnsetPlayAccess
`func (o *TimerInfoDtoProgramInfo) UnsetPlayAccess()`

UnsetPlayAccess ensures that no value is present for PlayAccess, not even an explicit nil
### GetAspectRatio

`func (o *TimerInfoDtoProgramInfo) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *TimerInfoDtoProgramInfo) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *TimerInfoDtoProgramInfo) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *TimerInfoDtoProgramInfo) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### SetAspectRatioNil

`func (o *TimerInfoDtoProgramInfo) SetAspectRatioNil(b bool)`

 SetAspectRatioNil sets the value for AspectRatio to be an explicit nil

### UnsetAspectRatio
`func (o *TimerInfoDtoProgramInfo) UnsetAspectRatio()`

UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil
### GetProductionYear

`func (o *TimerInfoDtoProgramInfo) GetProductionYear() int32`

GetProductionYear returns the ProductionYear field if non-nil, zero value otherwise.

### GetProductionYearOk

`func (o *TimerInfoDtoProgramInfo) GetProductionYearOk() (*int32, bool)`

GetProductionYearOk returns a tuple with the ProductionYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionYear

`func (o *TimerInfoDtoProgramInfo) SetProductionYear(v int32)`

SetProductionYear sets ProductionYear field to given value.

### HasProductionYear

`func (o *TimerInfoDtoProgramInfo) HasProductionYear() bool`

HasProductionYear returns a boolean if a field has been set.

### SetProductionYearNil

`func (o *TimerInfoDtoProgramInfo) SetProductionYearNil(b bool)`

 SetProductionYearNil sets the value for ProductionYear to be an explicit nil

### UnsetProductionYear
`func (o *TimerInfoDtoProgramInfo) UnsetProductionYear()`

UnsetProductionYear ensures that no value is present for ProductionYear, not even an explicit nil
### GetIsPlaceHolder

`func (o *TimerInfoDtoProgramInfo) GetIsPlaceHolder() bool`

GetIsPlaceHolder returns the IsPlaceHolder field if non-nil, zero value otherwise.

### GetIsPlaceHolderOk

`func (o *TimerInfoDtoProgramInfo) GetIsPlaceHolderOk() (*bool, bool)`

GetIsPlaceHolderOk returns a tuple with the IsPlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlaceHolder

`func (o *TimerInfoDtoProgramInfo) SetIsPlaceHolder(v bool)`

SetIsPlaceHolder sets IsPlaceHolder field to given value.

### HasIsPlaceHolder

`func (o *TimerInfoDtoProgramInfo) HasIsPlaceHolder() bool`

HasIsPlaceHolder returns a boolean if a field has been set.

### SetIsPlaceHolderNil

`func (o *TimerInfoDtoProgramInfo) SetIsPlaceHolderNil(b bool)`

 SetIsPlaceHolderNil sets the value for IsPlaceHolder to be an explicit nil

### UnsetIsPlaceHolder
`func (o *TimerInfoDtoProgramInfo) UnsetIsPlaceHolder()`

UnsetIsPlaceHolder ensures that no value is present for IsPlaceHolder, not even an explicit nil
### GetNumber

`func (o *TimerInfoDtoProgramInfo) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *TimerInfoDtoProgramInfo) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *TimerInfoDtoProgramInfo) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *TimerInfoDtoProgramInfo) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *TimerInfoDtoProgramInfo) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *TimerInfoDtoProgramInfo) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetChannelNumber

`func (o *TimerInfoDtoProgramInfo) GetChannelNumber() string`

GetChannelNumber returns the ChannelNumber field if non-nil, zero value otherwise.

### GetChannelNumberOk

`func (o *TimerInfoDtoProgramInfo) GetChannelNumberOk() (*string, bool)`

GetChannelNumberOk returns a tuple with the ChannelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelNumber

`func (o *TimerInfoDtoProgramInfo) SetChannelNumber(v string)`

SetChannelNumber sets ChannelNumber field to given value.

### HasChannelNumber

`func (o *TimerInfoDtoProgramInfo) HasChannelNumber() bool`

HasChannelNumber returns a boolean if a field has been set.

### SetChannelNumberNil

`func (o *TimerInfoDtoProgramInfo) SetChannelNumberNil(b bool)`

 SetChannelNumberNil sets the value for ChannelNumber to be an explicit nil

### UnsetChannelNumber
`func (o *TimerInfoDtoProgramInfo) UnsetChannelNumber()`

UnsetChannelNumber ensures that no value is present for ChannelNumber, not even an explicit nil
### GetIndexNumber

`func (o *TimerInfoDtoProgramInfo) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *TimerInfoDtoProgramInfo) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *TimerInfoDtoProgramInfo) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *TimerInfoDtoProgramInfo) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *TimerInfoDtoProgramInfo) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *TimerInfoDtoProgramInfo) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetIndexNumberEnd

`func (o *TimerInfoDtoProgramInfo) GetIndexNumberEnd() int32`

GetIndexNumberEnd returns the IndexNumberEnd field if non-nil, zero value otherwise.

### GetIndexNumberEndOk

`func (o *TimerInfoDtoProgramInfo) GetIndexNumberEndOk() (*int32, bool)`

GetIndexNumberEndOk returns a tuple with the IndexNumberEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumberEnd

`func (o *TimerInfoDtoProgramInfo) SetIndexNumberEnd(v int32)`

SetIndexNumberEnd sets IndexNumberEnd field to given value.

### HasIndexNumberEnd

`func (o *TimerInfoDtoProgramInfo) HasIndexNumberEnd() bool`

HasIndexNumberEnd returns a boolean if a field has been set.

### SetIndexNumberEndNil

`func (o *TimerInfoDtoProgramInfo) SetIndexNumberEndNil(b bool)`

 SetIndexNumberEndNil sets the value for IndexNumberEnd to be an explicit nil

### UnsetIndexNumberEnd
`func (o *TimerInfoDtoProgramInfo) UnsetIndexNumberEnd()`

UnsetIndexNumberEnd ensures that no value is present for IndexNumberEnd, not even an explicit nil
### GetParentIndexNumber

`func (o *TimerInfoDtoProgramInfo) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *TimerInfoDtoProgramInfo) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *TimerInfoDtoProgramInfo) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *TimerInfoDtoProgramInfo) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *TimerInfoDtoProgramInfo) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *TimerInfoDtoProgramInfo) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetRemoteTrailers

`func (o *TimerInfoDtoProgramInfo) GetRemoteTrailers() []MediaUrl`

GetRemoteTrailers returns the RemoteTrailers field if non-nil, zero value otherwise.

### GetRemoteTrailersOk

`func (o *TimerInfoDtoProgramInfo) GetRemoteTrailersOk() (*[]MediaUrl, bool)`

GetRemoteTrailersOk returns a tuple with the RemoteTrailers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTrailers

`func (o *TimerInfoDtoProgramInfo) SetRemoteTrailers(v []MediaUrl)`

SetRemoteTrailers sets RemoteTrailers field to given value.

### HasRemoteTrailers

`func (o *TimerInfoDtoProgramInfo) HasRemoteTrailers() bool`

HasRemoteTrailers returns a boolean if a field has been set.

### SetRemoteTrailersNil

`func (o *TimerInfoDtoProgramInfo) SetRemoteTrailersNil(b bool)`

 SetRemoteTrailersNil sets the value for RemoteTrailers to be an explicit nil

### UnsetRemoteTrailers
`func (o *TimerInfoDtoProgramInfo) UnsetRemoteTrailers()`

UnsetRemoteTrailers ensures that no value is present for RemoteTrailers, not even an explicit nil
### GetProviderIds

`func (o *TimerInfoDtoProgramInfo) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *TimerInfoDtoProgramInfo) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *TimerInfoDtoProgramInfo) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *TimerInfoDtoProgramInfo) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### SetProviderIdsNil

`func (o *TimerInfoDtoProgramInfo) SetProviderIdsNil(b bool)`

 SetProviderIdsNil sets the value for ProviderIds to be an explicit nil

### UnsetProviderIds
`func (o *TimerInfoDtoProgramInfo) UnsetProviderIds()`

UnsetProviderIds ensures that no value is present for ProviderIds, not even an explicit nil
### GetIsHD

`func (o *TimerInfoDtoProgramInfo) GetIsHD() bool`

GetIsHD returns the IsHD field if non-nil, zero value otherwise.

### GetIsHDOk

`func (o *TimerInfoDtoProgramInfo) GetIsHDOk() (*bool, bool)`

GetIsHDOk returns a tuple with the IsHD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHD

`func (o *TimerInfoDtoProgramInfo) SetIsHD(v bool)`

SetIsHD sets IsHD field to given value.

### HasIsHD

`func (o *TimerInfoDtoProgramInfo) HasIsHD() bool`

HasIsHD returns a boolean if a field has been set.

### SetIsHDNil

`func (o *TimerInfoDtoProgramInfo) SetIsHDNil(b bool)`

 SetIsHDNil sets the value for IsHD to be an explicit nil

### UnsetIsHD
`func (o *TimerInfoDtoProgramInfo) UnsetIsHD()`

UnsetIsHD ensures that no value is present for IsHD, not even an explicit nil
### GetIsFolder

`func (o *TimerInfoDtoProgramInfo) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *TimerInfoDtoProgramInfo) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *TimerInfoDtoProgramInfo) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *TimerInfoDtoProgramInfo) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### SetIsFolderNil

`func (o *TimerInfoDtoProgramInfo) SetIsFolderNil(b bool)`

 SetIsFolderNil sets the value for IsFolder to be an explicit nil

### UnsetIsFolder
`func (o *TimerInfoDtoProgramInfo) UnsetIsFolder()`

UnsetIsFolder ensures that no value is present for IsFolder, not even an explicit nil
### GetParentId

`func (o *TimerInfoDtoProgramInfo) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TimerInfoDtoProgramInfo) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TimerInfoDtoProgramInfo) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TimerInfoDtoProgramInfo) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *TimerInfoDtoProgramInfo) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *TimerInfoDtoProgramInfo) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetType

`func (o *TimerInfoDtoProgramInfo) GetType() BaseItemKind`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimerInfoDtoProgramInfo) GetTypeOk() (*BaseItemKind, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimerInfoDtoProgramInfo) SetType(v BaseItemKind)`

SetType sets Type field to given value.

### HasType

`func (o *TimerInfoDtoProgramInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPeople

`func (o *TimerInfoDtoProgramInfo) GetPeople() []BaseItemPerson`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *TimerInfoDtoProgramInfo) GetPeopleOk() (*[]BaseItemPerson, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *TimerInfoDtoProgramInfo) SetPeople(v []BaseItemPerson)`

SetPeople sets People field to given value.

### HasPeople

`func (o *TimerInfoDtoProgramInfo) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### SetPeopleNil

`func (o *TimerInfoDtoProgramInfo) SetPeopleNil(b bool)`

 SetPeopleNil sets the value for People to be an explicit nil

### UnsetPeople
`func (o *TimerInfoDtoProgramInfo) UnsetPeople()`

UnsetPeople ensures that no value is present for People, not even an explicit nil
### GetStudios

`func (o *TimerInfoDtoProgramInfo) GetStudios() []NameGuidPair`

GetStudios returns the Studios field if non-nil, zero value otherwise.

### GetStudiosOk

`func (o *TimerInfoDtoProgramInfo) GetStudiosOk() (*[]NameGuidPair, bool)`

GetStudiosOk returns a tuple with the Studios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudios

`func (o *TimerInfoDtoProgramInfo) SetStudios(v []NameGuidPair)`

SetStudios sets Studios field to given value.

### HasStudios

`func (o *TimerInfoDtoProgramInfo) HasStudios() bool`

HasStudios returns a boolean if a field has been set.

### SetStudiosNil

`func (o *TimerInfoDtoProgramInfo) SetStudiosNil(b bool)`

 SetStudiosNil sets the value for Studios to be an explicit nil

### UnsetStudios
`func (o *TimerInfoDtoProgramInfo) UnsetStudios()`

UnsetStudios ensures that no value is present for Studios, not even an explicit nil
### GetGenreItems

`func (o *TimerInfoDtoProgramInfo) GetGenreItems() []NameGuidPair`

GetGenreItems returns the GenreItems field if non-nil, zero value otherwise.

### GetGenreItemsOk

`func (o *TimerInfoDtoProgramInfo) GetGenreItemsOk() (*[]NameGuidPair, bool)`

GetGenreItemsOk returns a tuple with the GenreItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreItems

`func (o *TimerInfoDtoProgramInfo) SetGenreItems(v []NameGuidPair)`

SetGenreItems sets GenreItems field to given value.

### HasGenreItems

`func (o *TimerInfoDtoProgramInfo) HasGenreItems() bool`

HasGenreItems returns a boolean if a field has been set.

### SetGenreItemsNil

`func (o *TimerInfoDtoProgramInfo) SetGenreItemsNil(b bool)`

 SetGenreItemsNil sets the value for GenreItems to be an explicit nil

### UnsetGenreItems
`func (o *TimerInfoDtoProgramInfo) UnsetGenreItems()`

UnsetGenreItems ensures that no value is present for GenreItems, not even an explicit nil
### GetParentLogoItemId

`func (o *TimerInfoDtoProgramInfo) GetParentLogoItemId() string`

GetParentLogoItemId returns the ParentLogoItemId field if non-nil, zero value otherwise.

### GetParentLogoItemIdOk

`func (o *TimerInfoDtoProgramInfo) GetParentLogoItemIdOk() (*string, bool)`

GetParentLogoItemIdOk returns a tuple with the ParentLogoItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLogoItemId

`func (o *TimerInfoDtoProgramInfo) SetParentLogoItemId(v string)`

SetParentLogoItemId sets ParentLogoItemId field to given value.

### HasParentLogoItemId

`func (o *TimerInfoDtoProgramInfo) HasParentLogoItemId() bool`

HasParentLogoItemId returns a boolean if a field has been set.

### SetParentLogoItemIdNil

`func (o *TimerInfoDtoProgramInfo) SetParentLogoItemIdNil(b bool)`

 SetParentLogoItemIdNil sets the value for ParentLogoItemId to be an explicit nil

### UnsetParentLogoItemId
`func (o *TimerInfoDtoProgramInfo) UnsetParentLogoItemId()`

UnsetParentLogoItemId ensures that no value is present for ParentLogoItemId, not even an explicit nil
### GetParentBackdropItemId

`func (o *TimerInfoDtoProgramInfo) GetParentBackdropItemId() string`

GetParentBackdropItemId returns the ParentBackdropItemId field if non-nil, zero value otherwise.

### GetParentBackdropItemIdOk

`func (o *TimerInfoDtoProgramInfo) GetParentBackdropItemIdOk() (*string, bool)`

GetParentBackdropItemIdOk returns a tuple with the ParentBackdropItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropItemId

`func (o *TimerInfoDtoProgramInfo) SetParentBackdropItemId(v string)`

SetParentBackdropItemId sets ParentBackdropItemId field to given value.

### HasParentBackdropItemId

`func (o *TimerInfoDtoProgramInfo) HasParentBackdropItemId() bool`

HasParentBackdropItemId returns a boolean if a field has been set.

### SetParentBackdropItemIdNil

`func (o *TimerInfoDtoProgramInfo) SetParentBackdropItemIdNil(b bool)`

 SetParentBackdropItemIdNil sets the value for ParentBackdropItemId to be an explicit nil

### UnsetParentBackdropItemId
`func (o *TimerInfoDtoProgramInfo) UnsetParentBackdropItemId()`

UnsetParentBackdropItemId ensures that no value is present for ParentBackdropItemId, not even an explicit nil
### GetParentBackdropImageTags

`func (o *TimerInfoDtoProgramInfo) GetParentBackdropImageTags() []string`

GetParentBackdropImageTags returns the ParentBackdropImageTags field if non-nil, zero value otherwise.

### GetParentBackdropImageTagsOk

`func (o *TimerInfoDtoProgramInfo) GetParentBackdropImageTagsOk() (*[]string, bool)`

GetParentBackdropImageTagsOk returns a tuple with the ParentBackdropImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropImageTags

`func (o *TimerInfoDtoProgramInfo) SetParentBackdropImageTags(v []string)`

SetParentBackdropImageTags sets ParentBackdropImageTags field to given value.

### HasParentBackdropImageTags

`func (o *TimerInfoDtoProgramInfo) HasParentBackdropImageTags() bool`

HasParentBackdropImageTags returns a boolean if a field has been set.

### SetParentBackdropImageTagsNil

`func (o *TimerInfoDtoProgramInfo) SetParentBackdropImageTagsNil(b bool)`

 SetParentBackdropImageTagsNil sets the value for ParentBackdropImageTags to be an explicit nil

### UnsetParentBackdropImageTags
`func (o *TimerInfoDtoProgramInfo) UnsetParentBackdropImageTags()`

UnsetParentBackdropImageTags ensures that no value is present for ParentBackdropImageTags, not even an explicit nil
### GetLocalTrailerCount

`func (o *TimerInfoDtoProgramInfo) GetLocalTrailerCount() int32`

GetLocalTrailerCount returns the LocalTrailerCount field if non-nil, zero value otherwise.

### GetLocalTrailerCountOk

`func (o *TimerInfoDtoProgramInfo) GetLocalTrailerCountOk() (*int32, bool)`

GetLocalTrailerCountOk returns a tuple with the LocalTrailerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTrailerCount

`func (o *TimerInfoDtoProgramInfo) SetLocalTrailerCount(v int32)`

SetLocalTrailerCount sets LocalTrailerCount field to given value.

### HasLocalTrailerCount

`func (o *TimerInfoDtoProgramInfo) HasLocalTrailerCount() bool`

HasLocalTrailerCount returns a boolean if a field has been set.

### SetLocalTrailerCountNil

`func (o *TimerInfoDtoProgramInfo) SetLocalTrailerCountNil(b bool)`

 SetLocalTrailerCountNil sets the value for LocalTrailerCount to be an explicit nil

### UnsetLocalTrailerCount
`func (o *TimerInfoDtoProgramInfo) UnsetLocalTrailerCount()`

UnsetLocalTrailerCount ensures that no value is present for LocalTrailerCount, not even an explicit nil
### GetUserData

`func (o *TimerInfoDtoProgramInfo) GetUserData() BaseItemDtoUserData`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *TimerInfoDtoProgramInfo) GetUserDataOk() (*BaseItemDtoUserData, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *TimerInfoDtoProgramInfo) SetUserData(v BaseItemDtoUserData)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *TimerInfoDtoProgramInfo) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *TimerInfoDtoProgramInfo) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *TimerInfoDtoProgramInfo) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetRecursiveItemCount

`func (o *TimerInfoDtoProgramInfo) GetRecursiveItemCount() int32`

GetRecursiveItemCount returns the RecursiveItemCount field if non-nil, zero value otherwise.

### GetRecursiveItemCountOk

`func (o *TimerInfoDtoProgramInfo) GetRecursiveItemCountOk() (*int32, bool)`

GetRecursiveItemCountOk returns a tuple with the RecursiveItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveItemCount

`func (o *TimerInfoDtoProgramInfo) SetRecursiveItemCount(v int32)`

SetRecursiveItemCount sets RecursiveItemCount field to given value.

### HasRecursiveItemCount

`func (o *TimerInfoDtoProgramInfo) HasRecursiveItemCount() bool`

HasRecursiveItemCount returns a boolean if a field has been set.

### SetRecursiveItemCountNil

`func (o *TimerInfoDtoProgramInfo) SetRecursiveItemCountNil(b bool)`

 SetRecursiveItemCountNil sets the value for RecursiveItemCount to be an explicit nil

### UnsetRecursiveItemCount
`func (o *TimerInfoDtoProgramInfo) UnsetRecursiveItemCount()`

UnsetRecursiveItemCount ensures that no value is present for RecursiveItemCount, not even an explicit nil
### GetChildCount

`func (o *TimerInfoDtoProgramInfo) GetChildCount() int32`

GetChildCount returns the ChildCount field if non-nil, zero value otherwise.

### GetChildCountOk

`func (o *TimerInfoDtoProgramInfo) GetChildCountOk() (*int32, bool)`

GetChildCountOk returns a tuple with the ChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCount

`func (o *TimerInfoDtoProgramInfo) SetChildCount(v int32)`

SetChildCount sets ChildCount field to given value.

### HasChildCount

`func (o *TimerInfoDtoProgramInfo) HasChildCount() bool`

HasChildCount returns a boolean if a field has been set.

### SetChildCountNil

`func (o *TimerInfoDtoProgramInfo) SetChildCountNil(b bool)`

 SetChildCountNil sets the value for ChildCount to be an explicit nil

### UnsetChildCount
`func (o *TimerInfoDtoProgramInfo) UnsetChildCount()`

UnsetChildCount ensures that no value is present for ChildCount, not even an explicit nil
### GetSeriesName

`func (o *TimerInfoDtoProgramInfo) GetSeriesName() string`

GetSeriesName returns the SeriesName field if non-nil, zero value otherwise.

### GetSeriesNameOk

`func (o *TimerInfoDtoProgramInfo) GetSeriesNameOk() (*string, bool)`

GetSeriesNameOk returns a tuple with the SeriesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesName

`func (o *TimerInfoDtoProgramInfo) SetSeriesName(v string)`

SetSeriesName sets SeriesName field to given value.

### HasSeriesName

`func (o *TimerInfoDtoProgramInfo) HasSeriesName() bool`

HasSeriesName returns a boolean if a field has been set.

### SetSeriesNameNil

`func (o *TimerInfoDtoProgramInfo) SetSeriesNameNil(b bool)`

 SetSeriesNameNil sets the value for SeriesName to be an explicit nil

### UnsetSeriesName
`func (o *TimerInfoDtoProgramInfo) UnsetSeriesName()`

UnsetSeriesName ensures that no value is present for SeriesName, not even an explicit nil
### GetSeriesId

`func (o *TimerInfoDtoProgramInfo) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *TimerInfoDtoProgramInfo) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *TimerInfoDtoProgramInfo) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *TimerInfoDtoProgramInfo) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### SetSeriesIdNil

`func (o *TimerInfoDtoProgramInfo) SetSeriesIdNil(b bool)`

 SetSeriesIdNil sets the value for SeriesId to be an explicit nil

### UnsetSeriesId
`func (o *TimerInfoDtoProgramInfo) UnsetSeriesId()`

UnsetSeriesId ensures that no value is present for SeriesId, not even an explicit nil
### GetSeasonId

`func (o *TimerInfoDtoProgramInfo) GetSeasonId() string`

GetSeasonId returns the SeasonId field if non-nil, zero value otherwise.

### GetSeasonIdOk

`func (o *TimerInfoDtoProgramInfo) GetSeasonIdOk() (*string, bool)`

GetSeasonIdOk returns a tuple with the SeasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonId

`func (o *TimerInfoDtoProgramInfo) SetSeasonId(v string)`

SetSeasonId sets SeasonId field to given value.

### HasSeasonId

`func (o *TimerInfoDtoProgramInfo) HasSeasonId() bool`

HasSeasonId returns a boolean if a field has been set.

### SetSeasonIdNil

`func (o *TimerInfoDtoProgramInfo) SetSeasonIdNil(b bool)`

 SetSeasonIdNil sets the value for SeasonId to be an explicit nil

### UnsetSeasonId
`func (o *TimerInfoDtoProgramInfo) UnsetSeasonId()`

UnsetSeasonId ensures that no value is present for SeasonId, not even an explicit nil
### GetSpecialFeatureCount

`func (o *TimerInfoDtoProgramInfo) GetSpecialFeatureCount() int32`

GetSpecialFeatureCount returns the SpecialFeatureCount field if non-nil, zero value otherwise.

### GetSpecialFeatureCountOk

`func (o *TimerInfoDtoProgramInfo) GetSpecialFeatureCountOk() (*int32, bool)`

GetSpecialFeatureCountOk returns a tuple with the SpecialFeatureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialFeatureCount

`func (o *TimerInfoDtoProgramInfo) SetSpecialFeatureCount(v int32)`

SetSpecialFeatureCount sets SpecialFeatureCount field to given value.

### HasSpecialFeatureCount

`func (o *TimerInfoDtoProgramInfo) HasSpecialFeatureCount() bool`

HasSpecialFeatureCount returns a boolean if a field has been set.

### SetSpecialFeatureCountNil

`func (o *TimerInfoDtoProgramInfo) SetSpecialFeatureCountNil(b bool)`

 SetSpecialFeatureCountNil sets the value for SpecialFeatureCount to be an explicit nil

### UnsetSpecialFeatureCount
`func (o *TimerInfoDtoProgramInfo) UnsetSpecialFeatureCount()`

UnsetSpecialFeatureCount ensures that no value is present for SpecialFeatureCount, not even an explicit nil
### GetDisplayPreferencesId

`func (o *TimerInfoDtoProgramInfo) GetDisplayPreferencesId() string`

GetDisplayPreferencesId returns the DisplayPreferencesId field if non-nil, zero value otherwise.

### GetDisplayPreferencesIdOk

`func (o *TimerInfoDtoProgramInfo) GetDisplayPreferencesIdOk() (*string, bool)`

GetDisplayPreferencesIdOk returns a tuple with the DisplayPreferencesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPreferencesId

`func (o *TimerInfoDtoProgramInfo) SetDisplayPreferencesId(v string)`

SetDisplayPreferencesId sets DisplayPreferencesId field to given value.

### HasDisplayPreferencesId

`func (o *TimerInfoDtoProgramInfo) HasDisplayPreferencesId() bool`

HasDisplayPreferencesId returns a boolean if a field has been set.

### SetDisplayPreferencesIdNil

`func (o *TimerInfoDtoProgramInfo) SetDisplayPreferencesIdNil(b bool)`

 SetDisplayPreferencesIdNil sets the value for DisplayPreferencesId to be an explicit nil

### UnsetDisplayPreferencesId
`func (o *TimerInfoDtoProgramInfo) UnsetDisplayPreferencesId()`

UnsetDisplayPreferencesId ensures that no value is present for DisplayPreferencesId, not even an explicit nil
### GetStatus

`func (o *TimerInfoDtoProgramInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TimerInfoDtoProgramInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TimerInfoDtoProgramInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TimerInfoDtoProgramInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TimerInfoDtoProgramInfo) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TimerInfoDtoProgramInfo) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAirTime

`func (o *TimerInfoDtoProgramInfo) GetAirTime() string`

GetAirTime returns the AirTime field if non-nil, zero value otherwise.

### GetAirTimeOk

`func (o *TimerInfoDtoProgramInfo) GetAirTimeOk() (*string, bool)`

GetAirTimeOk returns a tuple with the AirTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirTime

`func (o *TimerInfoDtoProgramInfo) SetAirTime(v string)`

SetAirTime sets AirTime field to given value.

### HasAirTime

`func (o *TimerInfoDtoProgramInfo) HasAirTime() bool`

HasAirTime returns a boolean if a field has been set.

### SetAirTimeNil

`func (o *TimerInfoDtoProgramInfo) SetAirTimeNil(b bool)`

 SetAirTimeNil sets the value for AirTime to be an explicit nil

### UnsetAirTime
`func (o *TimerInfoDtoProgramInfo) UnsetAirTime()`

UnsetAirTime ensures that no value is present for AirTime, not even an explicit nil
### GetAirDays

`func (o *TimerInfoDtoProgramInfo) GetAirDays() []DayOfWeek`

GetAirDays returns the AirDays field if non-nil, zero value otherwise.

### GetAirDaysOk

`func (o *TimerInfoDtoProgramInfo) GetAirDaysOk() (*[]DayOfWeek, bool)`

GetAirDaysOk returns a tuple with the AirDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirDays

`func (o *TimerInfoDtoProgramInfo) SetAirDays(v []DayOfWeek)`

SetAirDays sets AirDays field to given value.

### HasAirDays

`func (o *TimerInfoDtoProgramInfo) HasAirDays() bool`

HasAirDays returns a boolean if a field has been set.

### SetAirDaysNil

`func (o *TimerInfoDtoProgramInfo) SetAirDaysNil(b bool)`

 SetAirDaysNil sets the value for AirDays to be an explicit nil

### UnsetAirDays
`func (o *TimerInfoDtoProgramInfo) UnsetAirDays()`

UnsetAirDays ensures that no value is present for AirDays, not even an explicit nil
### GetTags

`func (o *TimerInfoDtoProgramInfo) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TimerInfoDtoProgramInfo) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TimerInfoDtoProgramInfo) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TimerInfoDtoProgramInfo) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *TimerInfoDtoProgramInfo) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *TimerInfoDtoProgramInfo) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPrimaryImageAspectRatio

`func (o *TimerInfoDtoProgramInfo) GetPrimaryImageAspectRatio() float64`

GetPrimaryImageAspectRatio returns the PrimaryImageAspectRatio field if non-nil, zero value otherwise.

### GetPrimaryImageAspectRatioOk

`func (o *TimerInfoDtoProgramInfo) GetPrimaryImageAspectRatioOk() (*float64, bool)`

GetPrimaryImageAspectRatioOk returns a tuple with the PrimaryImageAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageAspectRatio

`func (o *TimerInfoDtoProgramInfo) SetPrimaryImageAspectRatio(v float64)`

SetPrimaryImageAspectRatio sets PrimaryImageAspectRatio field to given value.

### HasPrimaryImageAspectRatio

`func (o *TimerInfoDtoProgramInfo) HasPrimaryImageAspectRatio() bool`

HasPrimaryImageAspectRatio returns a boolean if a field has been set.

### SetPrimaryImageAspectRatioNil

`func (o *TimerInfoDtoProgramInfo) SetPrimaryImageAspectRatioNil(b bool)`

 SetPrimaryImageAspectRatioNil sets the value for PrimaryImageAspectRatio to be an explicit nil

### UnsetPrimaryImageAspectRatio
`func (o *TimerInfoDtoProgramInfo) UnsetPrimaryImageAspectRatio()`

UnsetPrimaryImageAspectRatio ensures that no value is present for PrimaryImageAspectRatio, not even an explicit nil
### GetArtists

`func (o *TimerInfoDtoProgramInfo) GetArtists() []string`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *TimerInfoDtoProgramInfo) GetArtistsOk() (*[]string, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *TimerInfoDtoProgramInfo) SetArtists(v []string)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *TimerInfoDtoProgramInfo) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### SetArtistsNil

`func (o *TimerInfoDtoProgramInfo) SetArtistsNil(b bool)`

 SetArtistsNil sets the value for Artists to be an explicit nil

### UnsetArtists
`func (o *TimerInfoDtoProgramInfo) UnsetArtists()`

UnsetArtists ensures that no value is present for Artists, not even an explicit nil
### GetArtistItems

`func (o *TimerInfoDtoProgramInfo) GetArtistItems() []NameGuidPair`

GetArtistItems returns the ArtistItems field if non-nil, zero value otherwise.

### GetArtistItemsOk

`func (o *TimerInfoDtoProgramInfo) GetArtistItemsOk() (*[]NameGuidPair, bool)`

GetArtistItemsOk returns a tuple with the ArtistItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistItems

`func (o *TimerInfoDtoProgramInfo) SetArtistItems(v []NameGuidPair)`

SetArtistItems sets ArtistItems field to given value.

### HasArtistItems

`func (o *TimerInfoDtoProgramInfo) HasArtistItems() bool`

HasArtistItems returns a boolean if a field has been set.

### SetArtistItemsNil

`func (o *TimerInfoDtoProgramInfo) SetArtistItemsNil(b bool)`

 SetArtistItemsNil sets the value for ArtistItems to be an explicit nil

### UnsetArtistItems
`func (o *TimerInfoDtoProgramInfo) UnsetArtistItems()`

UnsetArtistItems ensures that no value is present for ArtistItems, not even an explicit nil
### GetAlbum

`func (o *TimerInfoDtoProgramInfo) GetAlbum() string`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *TimerInfoDtoProgramInfo) GetAlbumOk() (*string, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *TimerInfoDtoProgramInfo) SetAlbum(v string)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *TimerInfoDtoProgramInfo) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### SetAlbumNil

`func (o *TimerInfoDtoProgramInfo) SetAlbumNil(b bool)`

 SetAlbumNil sets the value for Album to be an explicit nil

### UnsetAlbum
`func (o *TimerInfoDtoProgramInfo) UnsetAlbum()`

UnsetAlbum ensures that no value is present for Album, not even an explicit nil
### GetCollectionType

`func (o *TimerInfoDtoProgramInfo) GetCollectionType() CollectionType`

GetCollectionType returns the CollectionType field if non-nil, zero value otherwise.

### GetCollectionTypeOk

`func (o *TimerInfoDtoProgramInfo) GetCollectionTypeOk() (*CollectionType, bool)`

GetCollectionTypeOk returns a tuple with the CollectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionType

`func (o *TimerInfoDtoProgramInfo) SetCollectionType(v CollectionType)`

SetCollectionType sets CollectionType field to given value.

### HasCollectionType

`func (o *TimerInfoDtoProgramInfo) HasCollectionType() bool`

HasCollectionType returns a boolean if a field has been set.

### SetCollectionTypeNil

`func (o *TimerInfoDtoProgramInfo) SetCollectionTypeNil(b bool)`

 SetCollectionTypeNil sets the value for CollectionType to be an explicit nil

### UnsetCollectionType
`func (o *TimerInfoDtoProgramInfo) UnsetCollectionType()`

UnsetCollectionType ensures that no value is present for CollectionType, not even an explicit nil
### GetDisplayOrder

`func (o *TimerInfoDtoProgramInfo) GetDisplayOrder() string`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *TimerInfoDtoProgramInfo) GetDisplayOrderOk() (*string, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *TimerInfoDtoProgramInfo) SetDisplayOrder(v string)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *TimerInfoDtoProgramInfo) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### SetDisplayOrderNil

`func (o *TimerInfoDtoProgramInfo) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *TimerInfoDtoProgramInfo) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil
### GetAlbumId

`func (o *TimerInfoDtoProgramInfo) GetAlbumId() string`

GetAlbumId returns the AlbumId field if non-nil, zero value otherwise.

### GetAlbumIdOk

`func (o *TimerInfoDtoProgramInfo) GetAlbumIdOk() (*string, bool)`

GetAlbumIdOk returns a tuple with the AlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumId

`func (o *TimerInfoDtoProgramInfo) SetAlbumId(v string)`

SetAlbumId sets AlbumId field to given value.

### HasAlbumId

`func (o *TimerInfoDtoProgramInfo) HasAlbumId() bool`

HasAlbumId returns a boolean if a field has been set.

### SetAlbumIdNil

`func (o *TimerInfoDtoProgramInfo) SetAlbumIdNil(b bool)`

 SetAlbumIdNil sets the value for AlbumId to be an explicit nil

### UnsetAlbumId
`func (o *TimerInfoDtoProgramInfo) UnsetAlbumId()`

UnsetAlbumId ensures that no value is present for AlbumId, not even an explicit nil
### GetAlbumPrimaryImageTag

`func (o *TimerInfoDtoProgramInfo) GetAlbumPrimaryImageTag() string`

GetAlbumPrimaryImageTag returns the AlbumPrimaryImageTag field if non-nil, zero value otherwise.

### GetAlbumPrimaryImageTagOk

`func (o *TimerInfoDtoProgramInfo) GetAlbumPrimaryImageTagOk() (*string, bool)`

GetAlbumPrimaryImageTagOk returns a tuple with the AlbumPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumPrimaryImageTag

`func (o *TimerInfoDtoProgramInfo) SetAlbumPrimaryImageTag(v string)`

SetAlbumPrimaryImageTag sets AlbumPrimaryImageTag field to given value.

### HasAlbumPrimaryImageTag

`func (o *TimerInfoDtoProgramInfo) HasAlbumPrimaryImageTag() bool`

HasAlbumPrimaryImageTag returns a boolean if a field has been set.

### SetAlbumPrimaryImageTagNil

`func (o *TimerInfoDtoProgramInfo) SetAlbumPrimaryImageTagNil(b bool)`

 SetAlbumPrimaryImageTagNil sets the value for AlbumPrimaryImageTag to be an explicit nil

### UnsetAlbumPrimaryImageTag
`func (o *TimerInfoDtoProgramInfo) UnsetAlbumPrimaryImageTag()`

UnsetAlbumPrimaryImageTag ensures that no value is present for AlbumPrimaryImageTag, not even an explicit nil
### GetSeriesPrimaryImageTag

`func (o *TimerInfoDtoProgramInfo) GetSeriesPrimaryImageTag() string`

GetSeriesPrimaryImageTag returns the SeriesPrimaryImageTag field if non-nil, zero value otherwise.

### GetSeriesPrimaryImageTagOk

`func (o *TimerInfoDtoProgramInfo) GetSeriesPrimaryImageTagOk() (*string, bool)`

GetSeriesPrimaryImageTagOk returns a tuple with the SeriesPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesPrimaryImageTag

`func (o *TimerInfoDtoProgramInfo) SetSeriesPrimaryImageTag(v string)`

SetSeriesPrimaryImageTag sets SeriesPrimaryImageTag field to given value.

### HasSeriesPrimaryImageTag

`func (o *TimerInfoDtoProgramInfo) HasSeriesPrimaryImageTag() bool`

HasSeriesPrimaryImageTag returns a boolean if a field has been set.

### SetSeriesPrimaryImageTagNil

`func (o *TimerInfoDtoProgramInfo) SetSeriesPrimaryImageTagNil(b bool)`

 SetSeriesPrimaryImageTagNil sets the value for SeriesPrimaryImageTag to be an explicit nil

### UnsetSeriesPrimaryImageTag
`func (o *TimerInfoDtoProgramInfo) UnsetSeriesPrimaryImageTag()`

UnsetSeriesPrimaryImageTag ensures that no value is present for SeriesPrimaryImageTag, not even an explicit nil
### GetAlbumArtist

`func (o *TimerInfoDtoProgramInfo) GetAlbumArtist() string`

GetAlbumArtist returns the AlbumArtist field if non-nil, zero value otherwise.

### GetAlbumArtistOk

`func (o *TimerInfoDtoProgramInfo) GetAlbumArtistOk() (*string, bool)`

GetAlbumArtistOk returns a tuple with the AlbumArtist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtist

`func (o *TimerInfoDtoProgramInfo) SetAlbumArtist(v string)`

SetAlbumArtist sets AlbumArtist field to given value.

### HasAlbumArtist

`func (o *TimerInfoDtoProgramInfo) HasAlbumArtist() bool`

HasAlbumArtist returns a boolean if a field has been set.

### SetAlbumArtistNil

`func (o *TimerInfoDtoProgramInfo) SetAlbumArtistNil(b bool)`

 SetAlbumArtistNil sets the value for AlbumArtist to be an explicit nil

### UnsetAlbumArtist
`func (o *TimerInfoDtoProgramInfo) UnsetAlbumArtist()`

UnsetAlbumArtist ensures that no value is present for AlbumArtist, not even an explicit nil
### GetAlbumArtists

`func (o *TimerInfoDtoProgramInfo) GetAlbumArtists() []NameGuidPair`

GetAlbumArtists returns the AlbumArtists field if non-nil, zero value otherwise.

### GetAlbumArtistsOk

`func (o *TimerInfoDtoProgramInfo) GetAlbumArtistsOk() (*[]NameGuidPair, bool)`

GetAlbumArtistsOk returns a tuple with the AlbumArtists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtists

`func (o *TimerInfoDtoProgramInfo) SetAlbumArtists(v []NameGuidPair)`

SetAlbumArtists sets AlbumArtists field to given value.

### HasAlbumArtists

`func (o *TimerInfoDtoProgramInfo) HasAlbumArtists() bool`

HasAlbumArtists returns a boolean if a field has been set.

### SetAlbumArtistsNil

`func (o *TimerInfoDtoProgramInfo) SetAlbumArtistsNil(b bool)`

 SetAlbumArtistsNil sets the value for AlbumArtists to be an explicit nil

### UnsetAlbumArtists
`func (o *TimerInfoDtoProgramInfo) UnsetAlbumArtists()`

UnsetAlbumArtists ensures that no value is present for AlbumArtists, not even an explicit nil
### GetSeasonName

`func (o *TimerInfoDtoProgramInfo) GetSeasonName() string`

GetSeasonName returns the SeasonName field if non-nil, zero value otherwise.

### GetSeasonNameOk

`func (o *TimerInfoDtoProgramInfo) GetSeasonNameOk() (*string, bool)`

GetSeasonNameOk returns a tuple with the SeasonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonName

`func (o *TimerInfoDtoProgramInfo) SetSeasonName(v string)`

SetSeasonName sets SeasonName field to given value.

### HasSeasonName

`func (o *TimerInfoDtoProgramInfo) HasSeasonName() bool`

HasSeasonName returns a boolean if a field has been set.

### SetSeasonNameNil

`func (o *TimerInfoDtoProgramInfo) SetSeasonNameNil(b bool)`

 SetSeasonNameNil sets the value for SeasonName to be an explicit nil

### UnsetSeasonName
`func (o *TimerInfoDtoProgramInfo) UnsetSeasonName()`

UnsetSeasonName ensures that no value is present for SeasonName, not even an explicit nil
### GetMediaStreams

`func (o *TimerInfoDtoProgramInfo) GetMediaStreams() []MediaStream`

GetMediaStreams returns the MediaStreams field if non-nil, zero value otherwise.

### GetMediaStreamsOk

`func (o *TimerInfoDtoProgramInfo) GetMediaStreamsOk() (*[]MediaStream, bool)`

GetMediaStreamsOk returns a tuple with the MediaStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaStreams

`func (o *TimerInfoDtoProgramInfo) SetMediaStreams(v []MediaStream)`

SetMediaStreams sets MediaStreams field to given value.

### HasMediaStreams

`func (o *TimerInfoDtoProgramInfo) HasMediaStreams() bool`

HasMediaStreams returns a boolean if a field has been set.

### SetMediaStreamsNil

`func (o *TimerInfoDtoProgramInfo) SetMediaStreamsNil(b bool)`

 SetMediaStreamsNil sets the value for MediaStreams to be an explicit nil

### UnsetMediaStreams
`func (o *TimerInfoDtoProgramInfo) UnsetMediaStreams()`

UnsetMediaStreams ensures that no value is present for MediaStreams, not even an explicit nil
### GetVideoType

`func (o *TimerInfoDtoProgramInfo) GetVideoType() VideoType`

GetVideoType returns the VideoType field if non-nil, zero value otherwise.

### GetVideoTypeOk

`func (o *TimerInfoDtoProgramInfo) GetVideoTypeOk() (*VideoType, bool)`

GetVideoTypeOk returns a tuple with the VideoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoType

`func (o *TimerInfoDtoProgramInfo) SetVideoType(v VideoType)`

SetVideoType sets VideoType field to given value.

### HasVideoType

`func (o *TimerInfoDtoProgramInfo) HasVideoType() bool`

HasVideoType returns a boolean if a field has been set.

### SetVideoTypeNil

`func (o *TimerInfoDtoProgramInfo) SetVideoTypeNil(b bool)`

 SetVideoTypeNil sets the value for VideoType to be an explicit nil

### UnsetVideoType
`func (o *TimerInfoDtoProgramInfo) UnsetVideoType()`

UnsetVideoType ensures that no value is present for VideoType, not even an explicit nil
### GetPartCount

`func (o *TimerInfoDtoProgramInfo) GetPartCount() int32`

GetPartCount returns the PartCount field if non-nil, zero value otherwise.

### GetPartCountOk

`func (o *TimerInfoDtoProgramInfo) GetPartCountOk() (*int32, bool)`

GetPartCountOk returns a tuple with the PartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartCount

`func (o *TimerInfoDtoProgramInfo) SetPartCount(v int32)`

SetPartCount sets PartCount field to given value.

### HasPartCount

`func (o *TimerInfoDtoProgramInfo) HasPartCount() bool`

HasPartCount returns a boolean if a field has been set.

### SetPartCountNil

`func (o *TimerInfoDtoProgramInfo) SetPartCountNil(b bool)`

 SetPartCountNil sets the value for PartCount to be an explicit nil

### UnsetPartCount
`func (o *TimerInfoDtoProgramInfo) UnsetPartCount()`

UnsetPartCount ensures that no value is present for PartCount, not even an explicit nil
### GetMediaSourceCount

`func (o *TimerInfoDtoProgramInfo) GetMediaSourceCount() int32`

GetMediaSourceCount returns the MediaSourceCount field if non-nil, zero value otherwise.

### GetMediaSourceCountOk

`func (o *TimerInfoDtoProgramInfo) GetMediaSourceCountOk() (*int32, bool)`

GetMediaSourceCountOk returns a tuple with the MediaSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceCount

`func (o *TimerInfoDtoProgramInfo) SetMediaSourceCount(v int32)`

SetMediaSourceCount sets MediaSourceCount field to given value.

### HasMediaSourceCount

`func (o *TimerInfoDtoProgramInfo) HasMediaSourceCount() bool`

HasMediaSourceCount returns a boolean if a field has been set.

### SetMediaSourceCountNil

`func (o *TimerInfoDtoProgramInfo) SetMediaSourceCountNil(b bool)`

 SetMediaSourceCountNil sets the value for MediaSourceCount to be an explicit nil

### UnsetMediaSourceCount
`func (o *TimerInfoDtoProgramInfo) UnsetMediaSourceCount()`

UnsetMediaSourceCount ensures that no value is present for MediaSourceCount, not even an explicit nil
### GetImageTags

`func (o *TimerInfoDtoProgramInfo) GetImageTags() map[string]string`

GetImageTags returns the ImageTags field if non-nil, zero value otherwise.

### GetImageTagsOk

`func (o *TimerInfoDtoProgramInfo) GetImageTagsOk() (*map[string]string, bool)`

GetImageTagsOk returns a tuple with the ImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTags

`func (o *TimerInfoDtoProgramInfo) SetImageTags(v map[string]string)`

SetImageTags sets ImageTags field to given value.

### HasImageTags

`func (o *TimerInfoDtoProgramInfo) HasImageTags() bool`

HasImageTags returns a boolean if a field has been set.

### SetImageTagsNil

`func (o *TimerInfoDtoProgramInfo) SetImageTagsNil(b bool)`

 SetImageTagsNil sets the value for ImageTags to be an explicit nil

### UnsetImageTags
`func (o *TimerInfoDtoProgramInfo) UnsetImageTags()`

UnsetImageTags ensures that no value is present for ImageTags, not even an explicit nil
### GetBackdropImageTags

`func (o *TimerInfoDtoProgramInfo) GetBackdropImageTags() []string`

GetBackdropImageTags returns the BackdropImageTags field if non-nil, zero value otherwise.

### GetBackdropImageTagsOk

`func (o *TimerInfoDtoProgramInfo) GetBackdropImageTagsOk() (*[]string, bool)`

GetBackdropImageTagsOk returns a tuple with the BackdropImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdropImageTags

`func (o *TimerInfoDtoProgramInfo) SetBackdropImageTags(v []string)`

SetBackdropImageTags sets BackdropImageTags field to given value.

### HasBackdropImageTags

`func (o *TimerInfoDtoProgramInfo) HasBackdropImageTags() bool`

HasBackdropImageTags returns a boolean if a field has been set.

### SetBackdropImageTagsNil

`func (o *TimerInfoDtoProgramInfo) SetBackdropImageTagsNil(b bool)`

 SetBackdropImageTagsNil sets the value for BackdropImageTags to be an explicit nil

### UnsetBackdropImageTags
`func (o *TimerInfoDtoProgramInfo) UnsetBackdropImageTags()`

UnsetBackdropImageTags ensures that no value is present for BackdropImageTags, not even an explicit nil
### GetScreenshotImageTags

`func (o *TimerInfoDtoProgramInfo) GetScreenshotImageTags() []string`

GetScreenshotImageTags returns the ScreenshotImageTags field if non-nil, zero value otherwise.

### GetScreenshotImageTagsOk

`func (o *TimerInfoDtoProgramInfo) GetScreenshotImageTagsOk() (*[]string, bool)`

GetScreenshotImageTagsOk returns a tuple with the ScreenshotImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotImageTags

`func (o *TimerInfoDtoProgramInfo) SetScreenshotImageTags(v []string)`

SetScreenshotImageTags sets ScreenshotImageTags field to given value.

### HasScreenshotImageTags

`func (o *TimerInfoDtoProgramInfo) HasScreenshotImageTags() bool`

HasScreenshotImageTags returns a boolean if a field has been set.

### SetScreenshotImageTagsNil

`func (o *TimerInfoDtoProgramInfo) SetScreenshotImageTagsNil(b bool)`

 SetScreenshotImageTagsNil sets the value for ScreenshotImageTags to be an explicit nil

### UnsetScreenshotImageTags
`func (o *TimerInfoDtoProgramInfo) UnsetScreenshotImageTags()`

UnsetScreenshotImageTags ensures that no value is present for ScreenshotImageTags, not even an explicit nil
### GetParentLogoImageTag

`func (o *TimerInfoDtoProgramInfo) GetParentLogoImageTag() string`

GetParentLogoImageTag returns the ParentLogoImageTag field if non-nil, zero value otherwise.

### GetParentLogoImageTagOk

`func (o *TimerInfoDtoProgramInfo) GetParentLogoImageTagOk() (*string, bool)`

GetParentLogoImageTagOk returns a tuple with the ParentLogoImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLogoImageTag

`func (o *TimerInfoDtoProgramInfo) SetParentLogoImageTag(v string)`

SetParentLogoImageTag sets ParentLogoImageTag field to given value.

### HasParentLogoImageTag

`func (o *TimerInfoDtoProgramInfo) HasParentLogoImageTag() bool`

HasParentLogoImageTag returns a boolean if a field has been set.

### SetParentLogoImageTagNil

`func (o *TimerInfoDtoProgramInfo) SetParentLogoImageTagNil(b bool)`

 SetParentLogoImageTagNil sets the value for ParentLogoImageTag to be an explicit nil

### UnsetParentLogoImageTag
`func (o *TimerInfoDtoProgramInfo) UnsetParentLogoImageTag()`

UnsetParentLogoImageTag ensures that no value is present for ParentLogoImageTag, not even an explicit nil
### GetParentArtItemId

`func (o *TimerInfoDtoProgramInfo) GetParentArtItemId() string`

GetParentArtItemId returns the ParentArtItemId field if non-nil, zero value otherwise.

### GetParentArtItemIdOk

`func (o *TimerInfoDtoProgramInfo) GetParentArtItemIdOk() (*string, bool)`

GetParentArtItemIdOk returns a tuple with the ParentArtItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentArtItemId

`func (o *TimerInfoDtoProgramInfo) SetParentArtItemId(v string)`

SetParentArtItemId sets ParentArtItemId field to given value.

### HasParentArtItemId

`func (o *TimerInfoDtoProgramInfo) HasParentArtItemId() bool`

HasParentArtItemId returns a boolean if a field has been set.

### SetParentArtItemIdNil

`func (o *TimerInfoDtoProgramInfo) SetParentArtItemIdNil(b bool)`

 SetParentArtItemIdNil sets the value for ParentArtItemId to be an explicit nil

### UnsetParentArtItemId
`func (o *TimerInfoDtoProgramInfo) UnsetParentArtItemId()`

UnsetParentArtItemId ensures that no value is present for ParentArtItemId, not even an explicit nil
### GetParentArtImageTag

`func (o *TimerInfoDtoProgramInfo) GetParentArtImageTag() string`

GetParentArtImageTag returns the ParentArtImageTag field if non-nil, zero value otherwise.

### GetParentArtImageTagOk

`func (o *TimerInfoDtoProgramInfo) GetParentArtImageTagOk() (*string, bool)`

GetParentArtImageTagOk returns a tuple with the ParentArtImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentArtImageTag

`func (o *TimerInfoDtoProgramInfo) SetParentArtImageTag(v string)`

SetParentArtImageTag sets ParentArtImageTag field to given value.

### HasParentArtImageTag

`func (o *TimerInfoDtoProgramInfo) HasParentArtImageTag() bool`

HasParentArtImageTag returns a boolean if a field has been set.

### SetParentArtImageTagNil

`func (o *TimerInfoDtoProgramInfo) SetParentArtImageTagNil(b bool)`

 SetParentArtImageTagNil sets the value for ParentArtImageTag to be an explicit nil

### UnsetParentArtImageTag
`func (o *TimerInfoDtoProgramInfo) UnsetParentArtImageTag()`

UnsetParentArtImageTag ensures that no value is present for ParentArtImageTag, not even an explicit nil
### GetSeriesThumbImageTag

`func (o *TimerInfoDtoProgramInfo) GetSeriesThumbImageTag() string`

GetSeriesThumbImageTag returns the SeriesThumbImageTag field if non-nil, zero value otherwise.

### GetSeriesThumbImageTagOk

`func (o *TimerInfoDtoProgramInfo) GetSeriesThumbImageTagOk() (*string, bool)`

GetSeriesThumbImageTagOk returns a tuple with the SeriesThumbImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesThumbImageTag

`func (o *TimerInfoDtoProgramInfo) SetSeriesThumbImageTag(v string)`

SetSeriesThumbImageTag sets SeriesThumbImageTag field to given value.

### HasSeriesThumbImageTag

`func (o *TimerInfoDtoProgramInfo) HasSeriesThumbImageTag() bool`

HasSeriesThumbImageTag returns a boolean if a field has been set.

### SetSeriesThumbImageTagNil

`func (o *TimerInfoDtoProgramInfo) SetSeriesThumbImageTagNil(b bool)`

 SetSeriesThumbImageTagNil sets the value for SeriesThumbImageTag to be an explicit nil

### UnsetSeriesThumbImageTag
`func (o *TimerInfoDtoProgramInfo) UnsetSeriesThumbImageTag()`

UnsetSeriesThumbImageTag ensures that no value is present for SeriesThumbImageTag, not even an explicit nil
### GetImageBlurHashes

`func (o *TimerInfoDtoProgramInfo) GetImageBlurHashes() BaseItemDtoImageBlurHashes`

GetImageBlurHashes returns the ImageBlurHashes field if non-nil, zero value otherwise.

### GetImageBlurHashesOk

`func (o *TimerInfoDtoProgramInfo) GetImageBlurHashesOk() (*BaseItemDtoImageBlurHashes, bool)`

GetImageBlurHashesOk returns a tuple with the ImageBlurHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBlurHashes

`func (o *TimerInfoDtoProgramInfo) SetImageBlurHashes(v BaseItemDtoImageBlurHashes)`

SetImageBlurHashes sets ImageBlurHashes field to given value.

### HasImageBlurHashes

`func (o *TimerInfoDtoProgramInfo) HasImageBlurHashes() bool`

HasImageBlurHashes returns a boolean if a field has been set.

### SetImageBlurHashesNil

`func (o *TimerInfoDtoProgramInfo) SetImageBlurHashesNil(b bool)`

 SetImageBlurHashesNil sets the value for ImageBlurHashes to be an explicit nil

### UnsetImageBlurHashes
`func (o *TimerInfoDtoProgramInfo) UnsetImageBlurHashes()`

UnsetImageBlurHashes ensures that no value is present for ImageBlurHashes, not even an explicit nil
### GetSeriesStudio

`func (o *TimerInfoDtoProgramInfo) GetSeriesStudio() string`

GetSeriesStudio returns the SeriesStudio field if non-nil, zero value otherwise.

### GetSeriesStudioOk

`func (o *TimerInfoDtoProgramInfo) GetSeriesStudioOk() (*string, bool)`

GetSeriesStudioOk returns a tuple with the SeriesStudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesStudio

`func (o *TimerInfoDtoProgramInfo) SetSeriesStudio(v string)`

SetSeriesStudio sets SeriesStudio field to given value.

### HasSeriesStudio

`func (o *TimerInfoDtoProgramInfo) HasSeriesStudio() bool`

HasSeriesStudio returns a boolean if a field has been set.

### SetSeriesStudioNil

`func (o *TimerInfoDtoProgramInfo) SetSeriesStudioNil(b bool)`

 SetSeriesStudioNil sets the value for SeriesStudio to be an explicit nil

### UnsetSeriesStudio
`func (o *TimerInfoDtoProgramInfo) UnsetSeriesStudio()`

UnsetSeriesStudio ensures that no value is present for SeriesStudio, not even an explicit nil
### GetParentThumbItemId

`func (o *TimerInfoDtoProgramInfo) GetParentThumbItemId() string`

GetParentThumbItemId returns the ParentThumbItemId field if non-nil, zero value otherwise.

### GetParentThumbItemIdOk

`func (o *TimerInfoDtoProgramInfo) GetParentThumbItemIdOk() (*string, bool)`

GetParentThumbItemIdOk returns a tuple with the ParentThumbItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentThumbItemId

`func (o *TimerInfoDtoProgramInfo) SetParentThumbItemId(v string)`

SetParentThumbItemId sets ParentThumbItemId field to given value.

### HasParentThumbItemId

`func (o *TimerInfoDtoProgramInfo) HasParentThumbItemId() bool`

HasParentThumbItemId returns a boolean if a field has been set.

### SetParentThumbItemIdNil

`func (o *TimerInfoDtoProgramInfo) SetParentThumbItemIdNil(b bool)`

 SetParentThumbItemIdNil sets the value for ParentThumbItemId to be an explicit nil

### UnsetParentThumbItemId
`func (o *TimerInfoDtoProgramInfo) UnsetParentThumbItemId()`

UnsetParentThumbItemId ensures that no value is present for ParentThumbItemId, not even an explicit nil
### GetParentThumbImageTag

`func (o *TimerInfoDtoProgramInfo) GetParentThumbImageTag() string`

GetParentThumbImageTag returns the ParentThumbImageTag field if non-nil, zero value otherwise.

### GetParentThumbImageTagOk

`func (o *TimerInfoDtoProgramInfo) GetParentThumbImageTagOk() (*string, bool)`

GetParentThumbImageTagOk returns a tuple with the ParentThumbImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentThumbImageTag

`func (o *TimerInfoDtoProgramInfo) SetParentThumbImageTag(v string)`

SetParentThumbImageTag sets ParentThumbImageTag field to given value.

### HasParentThumbImageTag

`func (o *TimerInfoDtoProgramInfo) HasParentThumbImageTag() bool`

HasParentThumbImageTag returns a boolean if a field has been set.

### SetParentThumbImageTagNil

`func (o *TimerInfoDtoProgramInfo) SetParentThumbImageTagNil(b bool)`

 SetParentThumbImageTagNil sets the value for ParentThumbImageTag to be an explicit nil

### UnsetParentThumbImageTag
`func (o *TimerInfoDtoProgramInfo) UnsetParentThumbImageTag()`

UnsetParentThumbImageTag ensures that no value is present for ParentThumbImageTag, not even an explicit nil
### GetParentPrimaryImageItemId

`func (o *TimerInfoDtoProgramInfo) GetParentPrimaryImageItemId() string`

GetParentPrimaryImageItemId returns the ParentPrimaryImageItemId field if non-nil, zero value otherwise.

### GetParentPrimaryImageItemIdOk

`func (o *TimerInfoDtoProgramInfo) GetParentPrimaryImageItemIdOk() (*string, bool)`

GetParentPrimaryImageItemIdOk returns a tuple with the ParentPrimaryImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPrimaryImageItemId

`func (o *TimerInfoDtoProgramInfo) SetParentPrimaryImageItemId(v string)`

SetParentPrimaryImageItemId sets ParentPrimaryImageItemId field to given value.

### HasParentPrimaryImageItemId

`func (o *TimerInfoDtoProgramInfo) HasParentPrimaryImageItemId() bool`

HasParentPrimaryImageItemId returns a boolean if a field has been set.

### SetParentPrimaryImageItemIdNil

`func (o *TimerInfoDtoProgramInfo) SetParentPrimaryImageItemIdNil(b bool)`

 SetParentPrimaryImageItemIdNil sets the value for ParentPrimaryImageItemId to be an explicit nil

### UnsetParentPrimaryImageItemId
`func (o *TimerInfoDtoProgramInfo) UnsetParentPrimaryImageItemId()`

UnsetParentPrimaryImageItemId ensures that no value is present for ParentPrimaryImageItemId, not even an explicit nil
### GetParentPrimaryImageTag

`func (o *TimerInfoDtoProgramInfo) GetParentPrimaryImageTag() string`

GetParentPrimaryImageTag returns the ParentPrimaryImageTag field if non-nil, zero value otherwise.

### GetParentPrimaryImageTagOk

`func (o *TimerInfoDtoProgramInfo) GetParentPrimaryImageTagOk() (*string, bool)`

GetParentPrimaryImageTagOk returns a tuple with the ParentPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPrimaryImageTag

`func (o *TimerInfoDtoProgramInfo) SetParentPrimaryImageTag(v string)`

SetParentPrimaryImageTag sets ParentPrimaryImageTag field to given value.

### HasParentPrimaryImageTag

`func (o *TimerInfoDtoProgramInfo) HasParentPrimaryImageTag() bool`

HasParentPrimaryImageTag returns a boolean if a field has been set.

### SetParentPrimaryImageTagNil

`func (o *TimerInfoDtoProgramInfo) SetParentPrimaryImageTagNil(b bool)`

 SetParentPrimaryImageTagNil sets the value for ParentPrimaryImageTag to be an explicit nil

### UnsetParentPrimaryImageTag
`func (o *TimerInfoDtoProgramInfo) UnsetParentPrimaryImageTag()`

UnsetParentPrimaryImageTag ensures that no value is present for ParentPrimaryImageTag, not even an explicit nil
### GetChapters

`func (o *TimerInfoDtoProgramInfo) GetChapters() []ChapterInfo`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *TimerInfoDtoProgramInfo) GetChaptersOk() (*[]ChapterInfo, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *TimerInfoDtoProgramInfo) SetChapters(v []ChapterInfo)`

SetChapters sets Chapters field to given value.

### HasChapters

`func (o *TimerInfoDtoProgramInfo) HasChapters() bool`

HasChapters returns a boolean if a field has been set.

### SetChaptersNil

`func (o *TimerInfoDtoProgramInfo) SetChaptersNil(b bool)`

 SetChaptersNil sets the value for Chapters to be an explicit nil

### UnsetChapters
`func (o *TimerInfoDtoProgramInfo) UnsetChapters()`

UnsetChapters ensures that no value is present for Chapters, not even an explicit nil
### GetTrickplay

`func (o *TimerInfoDtoProgramInfo) GetTrickplay() map[string]map[string]TrickplayInfo`

GetTrickplay returns the Trickplay field if non-nil, zero value otherwise.

### GetTrickplayOk

`func (o *TimerInfoDtoProgramInfo) GetTrickplayOk() (*map[string]map[string]TrickplayInfo, bool)`

GetTrickplayOk returns a tuple with the Trickplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrickplay

`func (o *TimerInfoDtoProgramInfo) SetTrickplay(v map[string]map[string]TrickplayInfo)`

SetTrickplay sets Trickplay field to given value.

### HasTrickplay

`func (o *TimerInfoDtoProgramInfo) HasTrickplay() bool`

HasTrickplay returns a boolean if a field has been set.

### SetTrickplayNil

`func (o *TimerInfoDtoProgramInfo) SetTrickplayNil(b bool)`

 SetTrickplayNil sets the value for Trickplay to be an explicit nil

### UnsetTrickplay
`func (o *TimerInfoDtoProgramInfo) UnsetTrickplay()`

UnsetTrickplay ensures that no value is present for Trickplay, not even an explicit nil
### GetLocationType

`func (o *TimerInfoDtoProgramInfo) GetLocationType() LocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *TimerInfoDtoProgramInfo) GetLocationTypeOk() (*LocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *TimerInfoDtoProgramInfo) SetLocationType(v LocationType)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *TimerInfoDtoProgramInfo) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### SetLocationTypeNil

`func (o *TimerInfoDtoProgramInfo) SetLocationTypeNil(b bool)`

 SetLocationTypeNil sets the value for LocationType to be an explicit nil

### UnsetLocationType
`func (o *TimerInfoDtoProgramInfo) UnsetLocationType()`

UnsetLocationType ensures that no value is present for LocationType, not even an explicit nil
### GetIsoType

`func (o *TimerInfoDtoProgramInfo) GetIsoType() IsoType`

GetIsoType returns the IsoType field if non-nil, zero value otherwise.

### GetIsoTypeOk

`func (o *TimerInfoDtoProgramInfo) GetIsoTypeOk() (*IsoType, bool)`

GetIsoTypeOk returns a tuple with the IsoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoType

`func (o *TimerInfoDtoProgramInfo) SetIsoType(v IsoType)`

SetIsoType sets IsoType field to given value.

### HasIsoType

`func (o *TimerInfoDtoProgramInfo) HasIsoType() bool`

HasIsoType returns a boolean if a field has been set.

### SetIsoTypeNil

`func (o *TimerInfoDtoProgramInfo) SetIsoTypeNil(b bool)`

 SetIsoTypeNil sets the value for IsoType to be an explicit nil

### UnsetIsoType
`func (o *TimerInfoDtoProgramInfo) UnsetIsoType()`

UnsetIsoType ensures that no value is present for IsoType, not even an explicit nil
### GetMediaType

`func (o *TimerInfoDtoProgramInfo) GetMediaType() MediaType`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *TimerInfoDtoProgramInfo) GetMediaTypeOk() (*MediaType, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *TimerInfoDtoProgramInfo) SetMediaType(v MediaType)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *TimerInfoDtoProgramInfo) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetEndDate

`func (o *TimerInfoDtoProgramInfo) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TimerInfoDtoProgramInfo) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TimerInfoDtoProgramInfo) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TimerInfoDtoProgramInfo) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *TimerInfoDtoProgramInfo) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *TimerInfoDtoProgramInfo) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetLockedFields

`func (o *TimerInfoDtoProgramInfo) GetLockedFields() []MetadataField`

GetLockedFields returns the LockedFields field if non-nil, zero value otherwise.

### GetLockedFieldsOk

`func (o *TimerInfoDtoProgramInfo) GetLockedFieldsOk() (*[]MetadataField, bool)`

GetLockedFieldsOk returns a tuple with the LockedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedFields

`func (o *TimerInfoDtoProgramInfo) SetLockedFields(v []MetadataField)`

SetLockedFields sets LockedFields field to given value.

### HasLockedFields

`func (o *TimerInfoDtoProgramInfo) HasLockedFields() bool`

HasLockedFields returns a boolean if a field has been set.

### SetLockedFieldsNil

`func (o *TimerInfoDtoProgramInfo) SetLockedFieldsNil(b bool)`

 SetLockedFieldsNil sets the value for LockedFields to be an explicit nil

### UnsetLockedFields
`func (o *TimerInfoDtoProgramInfo) UnsetLockedFields()`

UnsetLockedFields ensures that no value is present for LockedFields, not even an explicit nil
### GetTrailerCount

`func (o *TimerInfoDtoProgramInfo) GetTrailerCount() int32`

GetTrailerCount returns the TrailerCount field if non-nil, zero value otherwise.

### GetTrailerCountOk

`func (o *TimerInfoDtoProgramInfo) GetTrailerCountOk() (*int32, bool)`

GetTrailerCountOk returns a tuple with the TrailerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerCount

`func (o *TimerInfoDtoProgramInfo) SetTrailerCount(v int32)`

SetTrailerCount sets TrailerCount field to given value.

### HasTrailerCount

`func (o *TimerInfoDtoProgramInfo) HasTrailerCount() bool`

HasTrailerCount returns a boolean if a field has been set.

### SetTrailerCountNil

`func (o *TimerInfoDtoProgramInfo) SetTrailerCountNil(b bool)`

 SetTrailerCountNil sets the value for TrailerCount to be an explicit nil

### UnsetTrailerCount
`func (o *TimerInfoDtoProgramInfo) UnsetTrailerCount()`

UnsetTrailerCount ensures that no value is present for TrailerCount, not even an explicit nil
### GetMovieCount

`func (o *TimerInfoDtoProgramInfo) GetMovieCount() int32`

GetMovieCount returns the MovieCount field if non-nil, zero value otherwise.

### GetMovieCountOk

`func (o *TimerInfoDtoProgramInfo) GetMovieCountOk() (*int32, bool)`

GetMovieCountOk returns a tuple with the MovieCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieCount

`func (o *TimerInfoDtoProgramInfo) SetMovieCount(v int32)`

SetMovieCount sets MovieCount field to given value.

### HasMovieCount

`func (o *TimerInfoDtoProgramInfo) HasMovieCount() bool`

HasMovieCount returns a boolean if a field has been set.

### SetMovieCountNil

`func (o *TimerInfoDtoProgramInfo) SetMovieCountNil(b bool)`

 SetMovieCountNil sets the value for MovieCount to be an explicit nil

### UnsetMovieCount
`func (o *TimerInfoDtoProgramInfo) UnsetMovieCount()`

UnsetMovieCount ensures that no value is present for MovieCount, not even an explicit nil
### GetSeriesCount

`func (o *TimerInfoDtoProgramInfo) GetSeriesCount() int32`

GetSeriesCount returns the SeriesCount field if non-nil, zero value otherwise.

### GetSeriesCountOk

`func (o *TimerInfoDtoProgramInfo) GetSeriesCountOk() (*int32, bool)`

GetSeriesCountOk returns a tuple with the SeriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesCount

`func (o *TimerInfoDtoProgramInfo) SetSeriesCount(v int32)`

SetSeriesCount sets SeriesCount field to given value.

### HasSeriesCount

`func (o *TimerInfoDtoProgramInfo) HasSeriesCount() bool`

HasSeriesCount returns a boolean if a field has been set.

### SetSeriesCountNil

`func (o *TimerInfoDtoProgramInfo) SetSeriesCountNil(b bool)`

 SetSeriesCountNil sets the value for SeriesCount to be an explicit nil

### UnsetSeriesCount
`func (o *TimerInfoDtoProgramInfo) UnsetSeriesCount()`

UnsetSeriesCount ensures that no value is present for SeriesCount, not even an explicit nil
### GetProgramCount

`func (o *TimerInfoDtoProgramInfo) GetProgramCount() int32`

GetProgramCount returns the ProgramCount field if non-nil, zero value otherwise.

### GetProgramCountOk

`func (o *TimerInfoDtoProgramInfo) GetProgramCountOk() (*int32, bool)`

GetProgramCountOk returns a tuple with the ProgramCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramCount

`func (o *TimerInfoDtoProgramInfo) SetProgramCount(v int32)`

SetProgramCount sets ProgramCount field to given value.

### HasProgramCount

`func (o *TimerInfoDtoProgramInfo) HasProgramCount() bool`

HasProgramCount returns a boolean if a field has been set.

### SetProgramCountNil

`func (o *TimerInfoDtoProgramInfo) SetProgramCountNil(b bool)`

 SetProgramCountNil sets the value for ProgramCount to be an explicit nil

### UnsetProgramCount
`func (o *TimerInfoDtoProgramInfo) UnsetProgramCount()`

UnsetProgramCount ensures that no value is present for ProgramCount, not even an explicit nil
### GetEpisodeCount

`func (o *TimerInfoDtoProgramInfo) GetEpisodeCount() int32`

GetEpisodeCount returns the EpisodeCount field if non-nil, zero value otherwise.

### GetEpisodeCountOk

`func (o *TimerInfoDtoProgramInfo) GetEpisodeCountOk() (*int32, bool)`

GetEpisodeCountOk returns a tuple with the EpisodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeCount

`func (o *TimerInfoDtoProgramInfo) SetEpisodeCount(v int32)`

SetEpisodeCount sets EpisodeCount field to given value.

### HasEpisodeCount

`func (o *TimerInfoDtoProgramInfo) HasEpisodeCount() bool`

HasEpisodeCount returns a boolean if a field has been set.

### SetEpisodeCountNil

`func (o *TimerInfoDtoProgramInfo) SetEpisodeCountNil(b bool)`

 SetEpisodeCountNil sets the value for EpisodeCount to be an explicit nil

### UnsetEpisodeCount
`func (o *TimerInfoDtoProgramInfo) UnsetEpisodeCount()`

UnsetEpisodeCount ensures that no value is present for EpisodeCount, not even an explicit nil
### GetSongCount

`func (o *TimerInfoDtoProgramInfo) GetSongCount() int32`

GetSongCount returns the SongCount field if non-nil, zero value otherwise.

### GetSongCountOk

`func (o *TimerInfoDtoProgramInfo) GetSongCountOk() (*int32, bool)`

GetSongCountOk returns a tuple with the SongCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSongCount

`func (o *TimerInfoDtoProgramInfo) SetSongCount(v int32)`

SetSongCount sets SongCount field to given value.

### HasSongCount

`func (o *TimerInfoDtoProgramInfo) HasSongCount() bool`

HasSongCount returns a boolean if a field has been set.

### SetSongCountNil

`func (o *TimerInfoDtoProgramInfo) SetSongCountNil(b bool)`

 SetSongCountNil sets the value for SongCount to be an explicit nil

### UnsetSongCount
`func (o *TimerInfoDtoProgramInfo) UnsetSongCount()`

UnsetSongCount ensures that no value is present for SongCount, not even an explicit nil
### GetAlbumCount

`func (o *TimerInfoDtoProgramInfo) GetAlbumCount() int32`

GetAlbumCount returns the AlbumCount field if non-nil, zero value otherwise.

### GetAlbumCountOk

`func (o *TimerInfoDtoProgramInfo) GetAlbumCountOk() (*int32, bool)`

GetAlbumCountOk returns a tuple with the AlbumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumCount

`func (o *TimerInfoDtoProgramInfo) SetAlbumCount(v int32)`

SetAlbumCount sets AlbumCount field to given value.

### HasAlbumCount

`func (o *TimerInfoDtoProgramInfo) HasAlbumCount() bool`

HasAlbumCount returns a boolean if a field has been set.

### SetAlbumCountNil

`func (o *TimerInfoDtoProgramInfo) SetAlbumCountNil(b bool)`

 SetAlbumCountNil sets the value for AlbumCount to be an explicit nil

### UnsetAlbumCount
`func (o *TimerInfoDtoProgramInfo) UnsetAlbumCount()`

UnsetAlbumCount ensures that no value is present for AlbumCount, not even an explicit nil
### GetArtistCount

`func (o *TimerInfoDtoProgramInfo) GetArtistCount() int32`

GetArtistCount returns the ArtistCount field if non-nil, zero value otherwise.

### GetArtistCountOk

`func (o *TimerInfoDtoProgramInfo) GetArtistCountOk() (*int32, bool)`

GetArtistCountOk returns a tuple with the ArtistCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistCount

`func (o *TimerInfoDtoProgramInfo) SetArtistCount(v int32)`

SetArtistCount sets ArtistCount field to given value.

### HasArtistCount

`func (o *TimerInfoDtoProgramInfo) HasArtistCount() bool`

HasArtistCount returns a boolean if a field has been set.

### SetArtistCountNil

`func (o *TimerInfoDtoProgramInfo) SetArtistCountNil(b bool)`

 SetArtistCountNil sets the value for ArtistCount to be an explicit nil

### UnsetArtistCount
`func (o *TimerInfoDtoProgramInfo) UnsetArtistCount()`

UnsetArtistCount ensures that no value is present for ArtistCount, not even an explicit nil
### GetMusicVideoCount

`func (o *TimerInfoDtoProgramInfo) GetMusicVideoCount() int32`

GetMusicVideoCount returns the MusicVideoCount field if non-nil, zero value otherwise.

### GetMusicVideoCountOk

`func (o *TimerInfoDtoProgramInfo) GetMusicVideoCountOk() (*int32, bool)`

GetMusicVideoCountOk returns a tuple with the MusicVideoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicVideoCount

`func (o *TimerInfoDtoProgramInfo) SetMusicVideoCount(v int32)`

SetMusicVideoCount sets MusicVideoCount field to given value.

### HasMusicVideoCount

`func (o *TimerInfoDtoProgramInfo) HasMusicVideoCount() bool`

HasMusicVideoCount returns a boolean if a field has been set.

### SetMusicVideoCountNil

`func (o *TimerInfoDtoProgramInfo) SetMusicVideoCountNil(b bool)`

 SetMusicVideoCountNil sets the value for MusicVideoCount to be an explicit nil

### UnsetMusicVideoCount
`func (o *TimerInfoDtoProgramInfo) UnsetMusicVideoCount()`

UnsetMusicVideoCount ensures that no value is present for MusicVideoCount, not even an explicit nil
### GetLockData

`func (o *TimerInfoDtoProgramInfo) GetLockData() bool`

GetLockData returns the LockData field if non-nil, zero value otherwise.

### GetLockDataOk

`func (o *TimerInfoDtoProgramInfo) GetLockDataOk() (*bool, bool)`

GetLockDataOk returns a tuple with the LockData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockData

`func (o *TimerInfoDtoProgramInfo) SetLockData(v bool)`

SetLockData sets LockData field to given value.

### HasLockData

`func (o *TimerInfoDtoProgramInfo) HasLockData() bool`

HasLockData returns a boolean if a field has been set.

### SetLockDataNil

`func (o *TimerInfoDtoProgramInfo) SetLockDataNil(b bool)`

 SetLockDataNil sets the value for LockData to be an explicit nil

### UnsetLockData
`func (o *TimerInfoDtoProgramInfo) UnsetLockData()`

UnsetLockData ensures that no value is present for LockData, not even an explicit nil
### GetWidth

`func (o *TimerInfoDtoProgramInfo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *TimerInfoDtoProgramInfo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *TimerInfoDtoProgramInfo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *TimerInfoDtoProgramInfo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *TimerInfoDtoProgramInfo) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *TimerInfoDtoProgramInfo) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *TimerInfoDtoProgramInfo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *TimerInfoDtoProgramInfo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *TimerInfoDtoProgramInfo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *TimerInfoDtoProgramInfo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *TimerInfoDtoProgramInfo) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *TimerInfoDtoProgramInfo) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetCameraMake

`func (o *TimerInfoDtoProgramInfo) GetCameraMake() string`

GetCameraMake returns the CameraMake field if non-nil, zero value otherwise.

### GetCameraMakeOk

`func (o *TimerInfoDtoProgramInfo) GetCameraMakeOk() (*string, bool)`

GetCameraMakeOk returns a tuple with the CameraMake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraMake

`func (o *TimerInfoDtoProgramInfo) SetCameraMake(v string)`

SetCameraMake sets CameraMake field to given value.

### HasCameraMake

`func (o *TimerInfoDtoProgramInfo) HasCameraMake() bool`

HasCameraMake returns a boolean if a field has been set.

### SetCameraMakeNil

`func (o *TimerInfoDtoProgramInfo) SetCameraMakeNil(b bool)`

 SetCameraMakeNil sets the value for CameraMake to be an explicit nil

### UnsetCameraMake
`func (o *TimerInfoDtoProgramInfo) UnsetCameraMake()`

UnsetCameraMake ensures that no value is present for CameraMake, not even an explicit nil
### GetCameraModel

`func (o *TimerInfoDtoProgramInfo) GetCameraModel() string`

GetCameraModel returns the CameraModel field if non-nil, zero value otherwise.

### GetCameraModelOk

`func (o *TimerInfoDtoProgramInfo) GetCameraModelOk() (*string, bool)`

GetCameraModelOk returns a tuple with the CameraModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraModel

`func (o *TimerInfoDtoProgramInfo) SetCameraModel(v string)`

SetCameraModel sets CameraModel field to given value.

### HasCameraModel

`func (o *TimerInfoDtoProgramInfo) HasCameraModel() bool`

HasCameraModel returns a boolean if a field has been set.

### SetCameraModelNil

`func (o *TimerInfoDtoProgramInfo) SetCameraModelNil(b bool)`

 SetCameraModelNil sets the value for CameraModel to be an explicit nil

### UnsetCameraModel
`func (o *TimerInfoDtoProgramInfo) UnsetCameraModel()`

UnsetCameraModel ensures that no value is present for CameraModel, not even an explicit nil
### GetSoftware

`func (o *TimerInfoDtoProgramInfo) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *TimerInfoDtoProgramInfo) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *TimerInfoDtoProgramInfo) SetSoftware(v string)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *TimerInfoDtoProgramInfo) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### SetSoftwareNil

`func (o *TimerInfoDtoProgramInfo) SetSoftwareNil(b bool)`

 SetSoftwareNil sets the value for Software to be an explicit nil

### UnsetSoftware
`func (o *TimerInfoDtoProgramInfo) UnsetSoftware()`

UnsetSoftware ensures that no value is present for Software, not even an explicit nil
### GetExposureTime

`func (o *TimerInfoDtoProgramInfo) GetExposureTime() float64`

GetExposureTime returns the ExposureTime field if non-nil, zero value otherwise.

### GetExposureTimeOk

`func (o *TimerInfoDtoProgramInfo) GetExposureTimeOk() (*float64, bool)`

GetExposureTimeOk returns a tuple with the ExposureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureTime

`func (o *TimerInfoDtoProgramInfo) SetExposureTime(v float64)`

SetExposureTime sets ExposureTime field to given value.

### HasExposureTime

`func (o *TimerInfoDtoProgramInfo) HasExposureTime() bool`

HasExposureTime returns a boolean if a field has been set.

### SetExposureTimeNil

`func (o *TimerInfoDtoProgramInfo) SetExposureTimeNil(b bool)`

 SetExposureTimeNil sets the value for ExposureTime to be an explicit nil

### UnsetExposureTime
`func (o *TimerInfoDtoProgramInfo) UnsetExposureTime()`

UnsetExposureTime ensures that no value is present for ExposureTime, not even an explicit nil
### GetFocalLength

`func (o *TimerInfoDtoProgramInfo) GetFocalLength() float64`

GetFocalLength returns the FocalLength field if non-nil, zero value otherwise.

### GetFocalLengthOk

`func (o *TimerInfoDtoProgramInfo) GetFocalLengthOk() (*float64, bool)`

GetFocalLengthOk returns a tuple with the FocalLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFocalLength

`func (o *TimerInfoDtoProgramInfo) SetFocalLength(v float64)`

SetFocalLength sets FocalLength field to given value.

### HasFocalLength

`func (o *TimerInfoDtoProgramInfo) HasFocalLength() bool`

HasFocalLength returns a boolean if a field has been set.

### SetFocalLengthNil

`func (o *TimerInfoDtoProgramInfo) SetFocalLengthNil(b bool)`

 SetFocalLengthNil sets the value for FocalLength to be an explicit nil

### UnsetFocalLength
`func (o *TimerInfoDtoProgramInfo) UnsetFocalLength()`

UnsetFocalLength ensures that no value is present for FocalLength, not even an explicit nil
### GetImageOrientation

`func (o *TimerInfoDtoProgramInfo) GetImageOrientation() ImageOrientation`

GetImageOrientation returns the ImageOrientation field if non-nil, zero value otherwise.

### GetImageOrientationOk

`func (o *TimerInfoDtoProgramInfo) GetImageOrientationOk() (*ImageOrientation, bool)`

GetImageOrientationOk returns a tuple with the ImageOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOrientation

`func (o *TimerInfoDtoProgramInfo) SetImageOrientation(v ImageOrientation)`

SetImageOrientation sets ImageOrientation field to given value.

### HasImageOrientation

`func (o *TimerInfoDtoProgramInfo) HasImageOrientation() bool`

HasImageOrientation returns a boolean if a field has been set.

### SetImageOrientationNil

`func (o *TimerInfoDtoProgramInfo) SetImageOrientationNil(b bool)`

 SetImageOrientationNil sets the value for ImageOrientation to be an explicit nil

### UnsetImageOrientation
`func (o *TimerInfoDtoProgramInfo) UnsetImageOrientation()`

UnsetImageOrientation ensures that no value is present for ImageOrientation, not even an explicit nil
### GetAperture

`func (o *TimerInfoDtoProgramInfo) GetAperture() float64`

GetAperture returns the Aperture field if non-nil, zero value otherwise.

### GetApertureOk

`func (o *TimerInfoDtoProgramInfo) GetApertureOk() (*float64, bool)`

GetApertureOk returns a tuple with the Aperture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAperture

`func (o *TimerInfoDtoProgramInfo) SetAperture(v float64)`

SetAperture sets Aperture field to given value.

### HasAperture

`func (o *TimerInfoDtoProgramInfo) HasAperture() bool`

HasAperture returns a boolean if a field has been set.

### SetApertureNil

`func (o *TimerInfoDtoProgramInfo) SetApertureNil(b bool)`

 SetApertureNil sets the value for Aperture to be an explicit nil

### UnsetAperture
`func (o *TimerInfoDtoProgramInfo) UnsetAperture()`

UnsetAperture ensures that no value is present for Aperture, not even an explicit nil
### GetShutterSpeed

`func (o *TimerInfoDtoProgramInfo) GetShutterSpeed() float64`

GetShutterSpeed returns the ShutterSpeed field if non-nil, zero value otherwise.

### GetShutterSpeedOk

`func (o *TimerInfoDtoProgramInfo) GetShutterSpeedOk() (*float64, bool)`

GetShutterSpeedOk returns a tuple with the ShutterSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutterSpeed

`func (o *TimerInfoDtoProgramInfo) SetShutterSpeed(v float64)`

SetShutterSpeed sets ShutterSpeed field to given value.

### HasShutterSpeed

`func (o *TimerInfoDtoProgramInfo) HasShutterSpeed() bool`

HasShutterSpeed returns a boolean if a field has been set.

### SetShutterSpeedNil

`func (o *TimerInfoDtoProgramInfo) SetShutterSpeedNil(b bool)`

 SetShutterSpeedNil sets the value for ShutterSpeed to be an explicit nil

### UnsetShutterSpeed
`func (o *TimerInfoDtoProgramInfo) UnsetShutterSpeed()`

UnsetShutterSpeed ensures that no value is present for ShutterSpeed, not even an explicit nil
### GetLatitude

`func (o *TimerInfoDtoProgramInfo) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *TimerInfoDtoProgramInfo) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *TimerInfoDtoProgramInfo) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *TimerInfoDtoProgramInfo) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *TimerInfoDtoProgramInfo) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *TimerInfoDtoProgramInfo) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *TimerInfoDtoProgramInfo) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *TimerInfoDtoProgramInfo) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *TimerInfoDtoProgramInfo) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *TimerInfoDtoProgramInfo) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *TimerInfoDtoProgramInfo) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *TimerInfoDtoProgramInfo) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetAltitude

`func (o *TimerInfoDtoProgramInfo) GetAltitude() float64`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *TimerInfoDtoProgramInfo) GetAltitudeOk() (*float64, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *TimerInfoDtoProgramInfo) SetAltitude(v float64)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *TimerInfoDtoProgramInfo) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### SetAltitudeNil

`func (o *TimerInfoDtoProgramInfo) SetAltitudeNil(b bool)`

 SetAltitudeNil sets the value for Altitude to be an explicit nil

### UnsetAltitude
`func (o *TimerInfoDtoProgramInfo) UnsetAltitude()`

UnsetAltitude ensures that no value is present for Altitude, not even an explicit nil
### GetIsoSpeedRating

`func (o *TimerInfoDtoProgramInfo) GetIsoSpeedRating() int32`

GetIsoSpeedRating returns the IsoSpeedRating field if non-nil, zero value otherwise.

### GetIsoSpeedRatingOk

`func (o *TimerInfoDtoProgramInfo) GetIsoSpeedRatingOk() (*int32, bool)`

GetIsoSpeedRatingOk returns a tuple with the IsoSpeedRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoSpeedRating

`func (o *TimerInfoDtoProgramInfo) SetIsoSpeedRating(v int32)`

SetIsoSpeedRating sets IsoSpeedRating field to given value.

### HasIsoSpeedRating

`func (o *TimerInfoDtoProgramInfo) HasIsoSpeedRating() bool`

HasIsoSpeedRating returns a boolean if a field has been set.

### SetIsoSpeedRatingNil

`func (o *TimerInfoDtoProgramInfo) SetIsoSpeedRatingNil(b bool)`

 SetIsoSpeedRatingNil sets the value for IsoSpeedRating to be an explicit nil

### UnsetIsoSpeedRating
`func (o *TimerInfoDtoProgramInfo) UnsetIsoSpeedRating()`

UnsetIsoSpeedRating ensures that no value is present for IsoSpeedRating, not even an explicit nil
### GetSeriesTimerId

`func (o *TimerInfoDtoProgramInfo) GetSeriesTimerId() string`

GetSeriesTimerId returns the SeriesTimerId field if non-nil, zero value otherwise.

### GetSeriesTimerIdOk

`func (o *TimerInfoDtoProgramInfo) GetSeriesTimerIdOk() (*string, bool)`

GetSeriesTimerIdOk returns a tuple with the SeriesTimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTimerId

`func (o *TimerInfoDtoProgramInfo) SetSeriesTimerId(v string)`

SetSeriesTimerId sets SeriesTimerId field to given value.

### HasSeriesTimerId

`func (o *TimerInfoDtoProgramInfo) HasSeriesTimerId() bool`

HasSeriesTimerId returns a boolean if a field has been set.

### SetSeriesTimerIdNil

`func (o *TimerInfoDtoProgramInfo) SetSeriesTimerIdNil(b bool)`

 SetSeriesTimerIdNil sets the value for SeriesTimerId to be an explicit nil

### UnsetSeriesTimerId
`func (o *TimerInfoDtoProgramInfo) UnsetSeriesTimerId()`

UnsetSeriesTimerId ensures that no value is present for SeriesTimerId, not even an explicit nil
### GetProgramId

`func (o *TimerInfoDtoProgramInfo) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *TimerInfoDtoProgramInfo) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *TimerInfoDtoProgramInfo) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *TimerInfoDtoProgramInfo) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramIdNil

`func (o *TimerInfoDtoProgramInfo) SetProgramIdNil(b bool)`

 SetProgramIdNil sets the value for ProgramId to be an explicit nil

### UnsetProgramId
`func (o *TimerInfoDtoProgramInfo) UnsetProgramId()`

UnsetProgramId ensures that no value is present for ProgramId, not even an explicit nil
### GetChannelPrimaryImageTag

`func (o *TimerInfoDtoProgramInfo) GetChannelPrimaryImageTag() string`

GetChannelPrimaryImageTag returns the ChannelPrimaryImageTag field if non-nil, zero value otherwise.

### GetChannelPrimaryImageTagOk

`func (o *TimerInfoDtoProgramInfo) GetChannelPrimaryImageTagOk() (*string, bool)`

GetChannelPrimaryImageTagOk returns a tuple with the ChannelPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPrimaryImageTag

`func (o *TimerInfoDtoProgramInfo) SetChannelPrimaryImageTag(v string)`

SetChannelPrimaryImageTag sets ChannelPrimaryImageTag field to given value.

### HasChannelPrimaryImageTag

`func (o *TimerInfoDtoProgramInfo) HasChannelPrimaryImageTag() bool`

HasChannelPrimaryImageTag returns a boolean if a field has been set.

### SetChannelPrimaryImageTagNil

`func (o *TimerInfoDtoProgramInfo) SetChannelPrimaryImageTagNil(b bool)`

 SetChannelPrimaryImageTagNil sets the value for ChannelPrimaryImageTag to be an explicit nil

### UnsetChannelPrimaryImageTag
`func (o *TimerInfoDtoProgramInfo) UnsetChannelPrimaryImageTag()`

UnsetChannelPrimaryImageTag ensures that no value is present for ChannelPrimaryImageTag, not even an explicit nil
### GetStartDate

`func (o *TimerInfoDtoProgramInfo) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TimerInfoDtoProgramInfo) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TimerInfoDtoProgramInfo) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TimerInfoDtoProgramInfo) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *TimerInfoDtoProgramInfo) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *TimerInfoDtoProgramInfo) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetCompletionPercentage

`func (o *TimerInfoDtoProgramInfo) GetCompletionPercentage() float64`

GetCompletionPercentage returns the CompletionPercentage field if non-nil, zero value otherwise.

### GetCompletionPercentageOk

`func (o *TimerInfoDtoProgramInfo) GetCompletionPercentageOk() (*float64, bool)`

GetCompletionPercentageOk returns a tuple with the CompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionPercentage

`func (o *TimerInfoDtoProgramInfo) SetCompletionPercentage(v float64)`

SetCompletionPercentage sets CompletionPercentage field to given value.

### HasCompletionPercentage

`func (o *TimerInfoDtoProgramInfo) HasCompletionPercentage() bool`

HasCompletionPercentage returns a boolean if a field has been set.

### SetCompletionPercentageNil

`func (o *TimerInfoDtoProgramInfo) SetCompletionPercentageNil(b bool)`

 SetCompletionPercentageNil sets the value for CompletionPercentage to be an explicit nil

### UnsetCompletionPercentage
`func (o *TimerInfoDtoProgramInfo) UnsetCompletionPercentage()`

UnsetCompletionPercentage ensures that no value is present for CompletionPercentage, not even an explicit nil
### GetIsRepeat

`func (o *TimerInfoDtoProgramInfo) GetIsRepeat() bool`

GetIsRepeat returns the IsRepeat field if non-nil, zero value otherwise.

### GetIsRepeatOk

`func (o *TimerInfoDtoProgramInfo) GetIsRepeatOk() (*bool, bool)`

GetIsRepeatOk returns a tuple with the IsRepeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRepeat

`func (o *TimerInfoDtoProgramInfo) SetIsRepeat(v bool)`

SetIsRepeat sets IsRepeat field to given value.

### HasIsRepeat

`func (o *TimerInfoDtoProgramInfo) HasIsRepeat() bool`

HasIsRepeat returns a boolean if a field has been set.

### SetIsRepeatNil

`func (o *TimerInfoDtoProgramInfo) SetIsRepeatNil(b bool)`

 SetIsRepeatNil sets the value for IsRepeat to be an explicit nil

### UnsetIsRepeat
`func (o *TimerInfoDtoProgramInfo) UnsetIsRepeat()`

UnsetIsRepeat ensures that no value is present for IsRepeat, not even an explicit nil
### GetEpisodeTitle

`func (o *TimerInfoDtoProgramInfo) GetEpisodeTitle() string`

GetEpisodeTitle returns the EpisodeTitle field if non-nil, zero value otherwise.

### GetEpisodeTitleOk

`func (o *TimerInfoDtoProgramInfo) GetEpisodeTitleOk() (*string, bool)`

GetEpisodeTitleOk returns a tuple with the EpisodeTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeTitle

`func (o *TimerInfoDtoProgramInfo) SetEpisodeTitle(v string)`

SetEpisodeTitle sets EpisodeTitle field to given value.

### HasEpisodeTitle

`func (o *TimerInfoDtoProgramInfo) HasEpisodeTitle() bool`

HasEpisodeTitle returns a boolean if a field has been set.

### SetEpisodeTitleNil

`func (o *TimerInfoDtoProgramInfo) SetEpisodeTitleNil(b bool)`

 SetEpisodeTitleNil sets the value for EpisodeTitle to be an explicit nil

### UnsetEpisodeTitle
`func (o *TimerInfoDtoProgramInfo) UnsetEpisodeTitle()`

UnsetEpisodeTitle ensures that no value is present for EpisodeTitle, not even an explicit nil
### GetChannelType

`func (o *TimerInfoDtoProgramInfo) GetChannelType() ChannelType`

GetChannelType returns the ChannelType field if non-nil, zero value otherwise.

### GetChannelTypeOk

`func (o *TimerInfoDtoProgramInfo) GetChannelTypeOk() (*ChannelType, bool)`

GetChannelTypeOk returns a tuple with the ChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelType

`func (o *TimerInfoDtoProgramInfo) SetChannelType(v ChannelType)`

SetChannelType sets ChannelType field to given value.

### HasChannelType

`func (o *TimerInfoDtoProgramInfo) HasChannelType() bool`

HasChannelType returns a boolean if a field has been set.

### SetChannelTypeNil

`func (o *TimerInfoDtoProgramInfo) SetChannelTypeNil(b bool)`

 SetChannelTypeNil sets the value for ChannelType to be an explicit nil

### UnsetChannelType
`func (o *TimerInfoDtoProgramInfo) UnsetChannelType()`

UnsetChannelType ensures that no value is present for ChannelType, not even an explicit nil
### GetAudio

`func (o *TimerInfoDtoProgramInfo) GetAudio() ProgramAudio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *TimerInfoDtoProgramInfo) GetAudioOk() (*ProgramAudio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *TimerInfoDtoProgramInfo) SetAudio(v ProgramAudio)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *TimerInfoDtoProgramInfo) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### SetAudioNil

`func (o *TimerInfoDtoProgramInfo) SetAudioNil(b bool)`

 SetAudioNil sets the value for Audio to be an explicit nil

### UnsetAudio
`func (o *TimerInfoDtoProgramInfo) UnsetAudio()`

UnsetAudio ensures that no value is present for Audio, not even an explicit nil
### GetIsMovie

`func (o *TimerInfoDtoProgramInfo) GetIsMovie() bool`

GetIsMovie returns the IsMovie field if non-nil, zero value otherwise.

### GetIsMovieOk

`func (o *TimerInfoDtoProgramInfo) GetIsMovieOk() (*bool, bool)`

GetIsMovieOk returns a tuple with the IsMovie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMovie

`func (o *TimerInfoDtoProgramInfo) SetIsMovie(v bool)`

SetIsMovie sets IsMovie field to given value.

### HasIsMovie

`func (o *TimerInfoDtoProgramInfo) HasIsMovie() bool`

HasIsMovie returns a boolean if a field has been set.

### SetIsMovieNil

`func (o *TimerInfoDtoProgramInfo) SetIsMovieNil(b bool)`

 SetIsMovieNil sets the value for IsMovie to be an explicit nil

### UnsetIsMovie
`func (o *TimerInfoDtoProgramInfo) UnsetIsMovie()`

UnsetIsMovie ensures that no value is present for IsMovie, not even an explicit nil
### GetIsSports

`func (o *TimerInfoDtoProgramInfo) GetIsSports() bool`

GetIsSports returns the IsSports field if non-nil, zero value otherwise.

### GetIsSportsOk

`func (o *TimerInfoDtoProgramInfo) GetIsSportsOk() (*bool, bool)`

GetIsSportsOk returns a tuple with the IsSports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSports

`func (o *TimerInfoDtoProgramInfo) SetIsSports(v bool)`

SetIsSports sets IsSports field to given value.

### HasIsSports

`func (o *TimerInfoDtoProgramInfo) HasIsSports() bool`

HasIsSports returns a boolean if a field has been set.

### SetIsSportsNil

`func (o *TimerInfoDtoProgramInfo) SetIsSportsNil(b bool)`

 SetIsSportsNil sets the value for IsSports to be an explicit nil

### UnsetIsSports
`func (o *TimerInfoDtoProgramInfo) UnsetIsSports()`

UnsetIsSports ensures that no value is present for IsSports, not even an explicit nil
### GetIsSeries

`func (o *TimerInfoDtoProgramInfo) GetIsSeries() bool`

GetIsSeries returns the IsSeries field if non-nil, zero value otherwise.

### GetIsSeriesOk

`func (o *TimerInfoDtoProgramInfo) GetIsSeriesOk() (*bool, bool)`

GetIsSeriesOk returns a tuple with the IsSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeries

`func (o *TimerInfoDtoProgramInfo) SetIsSeries(v bool)`

SetIsSeries sets IsSeries field to given value.

### HasIsSeries

`func (o *TimerInfoDtoProgramInfo) HasIsSeries() bool`

HasIsSeries returns a boolean if a field has been set.

### SetIsSeriesNil

`func (o *TimerInfoDtoProgramInfo) SetIsSeriesNil(b bool)`

 SetIsSeriesNil sets the value for IsSeries to be an explicit nil

### UnsetIsSeries
`func (o *TimerInfoDtoProgramInfo) UnsetIsSeries()`

UnsetIsSeries ensures that no value is present for IsSeries, not even an explicit nil
### GetIsLive

`func (o *TimerInfoDtoProgramInfo) GetIsLive() bool`

GetIsLive returns the IsLive field if non-nil, zero value otherwise.

### GetIsLiveOk

`func (o *TimerInfoDtoProgramInfo) GetIsLiveOk() (*bool, bool)`

GetIsLiveOk returns a tuple with the IsLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLive

`func (o *TimerInfoDtoProgramInfo) SetIsLive(v bool)`

SetIsLive sets IsLive field to given value.

### HasIsLive

`func (o *TimerInfoDtoProgramInfo) HasIsLive() bool`

HasIsLive returns a boolean if a field has been set.

### SetIsLiveNil

`func (o *TimerInfoDtoProgramInfo) SetIsLiveNil(b bool)`

 SetIsLiveNil sets the value for IsLive to be an explicit nil

### UnsetIsLive
`func (o *TimerInfoDtoProgramInfo) UnsetIsLive()`

UnsetIsLive ensures that no value is present for IsLive, not even an explicit nil
### GetIsNews

`func (o *TimerInfoDtoProgramInfo) GetIsNews() bool`

GetIsNews returns the IsNews field if non-nil, zero value otherwise.

### GetIsNewsOk

`func (o *TimerInfoDtoProgramInfo) GetIsNewsOk() (*bool, bool)`

GetIsNewsOk returns a tuple with the IsNews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNews

`func (o *TimerInfoDtoProgramInfo) SetIsNews(v bool)`

SetIsNews sets IsNews field to given value.

### HasIsNews

`func (o *TimerInfoDtoProgramInfo) HasIsNews() bool`

HasIsNews returns a boolean if a field has been set.

### SetIsNewsNil

`func (o *TimerInfoDtoProgramInfo) SetIsNewsNil(b bool)`

 SetIsNewsNil sets the value for IsNews to be an explicit nil

### UnsetIsNews
`func (o *TimerInfoDtoProgramInfo) UnsetIsNews()`

UnsetIsNews ensures that no value is present for IsNews, not even an explicit nil
### GetIsKids

`func (o *TimerInfoDtoProgramInfo) GetIsKids() bool`

GetIsKids returns the IsKids field if non-nil, zero value otherwise.

### GetIsKidsOk

`func (o *TimerInfoDtoProgramInfo) GetIsKidsOk() (*bool, bool)`

GetIsKidsOk returns a tuple with the IsKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKids

`func (o *TimerInfoDtoProgramInfo) SetIsKids(v bool)`

SetIsKids sets IsKids field to given value.

### HasIsKids

`func (o *TimerInfoDtoProgramInfo) HasIsKids() bool`

HasIsKids returns a boolean if a field has been set.

### SetIsKidsNil

`func (o *TimerInfoDtoProgramInfo) SetIsKidsNil(b bool)`

 SetIsKidsNil sets the value for IsKids to be an explicit nil

### UnsetIsKids
`func (o *TimerInfoDtoProgramInfo) UnsetIsKids()`

UnsetIsKids ensures that no value is present for IsKids, not even an explicit nil
### GetIsPremiere

`func (o *TimerInfoDtoProgramInfo) GetIsPremiere() bool`

GetIsPremiere returns the IsPremiere field if non-nil, zero value otherwise.

### GetIsPremiereOk

`func (o *TimerInfoDtoProgramInfo) GetIsPremiereOk() (*bool, bool)`

GetIsPremiereOk returns a tuple with the IsPremiere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPremiere

`func (o *TimerInfoDtoProgramInfo) SetIsPremiere(v bool)`

SetIsPremiere sets IsPremiere field to given value.

### HasIsPremiere

`func (o *TimerInfoDtoProgramInfo) HasIsPremiere() bool`

HasIsPremiere returns a boolean if a field has been set.

### SetIsPremiereNil

`func (o *TimerInfoDtoProgramInfo) SetIsPremiereNil(b bool)`

 SetIsPremiereNil sets the value for IsPremiere to be an explicit nil

### UnsetIsPremiere
`func (o *TimerInfoDtoProgramInfo) UnsetIsPremiere()`

UnsetIsPremiere ensures that no value is present for IsPremiere, not even an explicit nil
### GetTimerId

`func (o *TimerInfoDtoProgramInfo) GetTimerId() string`

GetTimerId returns the TimerId field if non-nil, zero value otherwise.

### GetTimerIdOk

`func (o *TimerInfoDtoProgramInfo) GetTimerIdOk() (*string, bool)`

GetTimerIdOk returns a tuple with the TimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerId

`func (o *TimerInfoDtoProgramInfo) SetTimerId(v string)`

SetTimerId sets TimerId field to given value.

### HasTimerId

`func (o *TimerInfoDtoProgramInfo) HasTimerId() bool`

HasTimerId returns a boolean if a field has been set.

### SetTimerIdNil

`func (o *TimerInfoDtoProgramInfo) SetTimerIdNil(b bool)`

 SetTimerIdNil sets the value for TimerId to be an explicit nil

### UnsetTimerId
`func (o *TimerInfoDtoProgramInfo) UnsetTimerId()`

UnsetTimerId ensures that no value is present for TimerId, not even an explicit nil
### GetNormalizationGain

`func (o *TimerInfoDtoProgramInfo) GetNormalizationGain() float32`

GetNormalizationGain returns the NormalizationGain field if non-nil, zero value otherwise.

### GetNormalizationGainOk

`func (o *TimerInfoDtoProgramInfo) GetNormalizationGainOk() (*float32, bool)`

GetNormalizationGainOk returns a tuple with the NormalizationGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizationGain

`func (o *TimerInfoDtoProgramInfo) SetNormalizationGain(v float32)`

SetNormalizationGain sets NormalizationGain field to given value.

### HasNormalizationGain

`func (o *TimerInfoDtoProgramInfo) HasNormalizationGain() bool`

HasNormalizationGain returns a boolean if a field has been set.

### SetNormalizationGainNil

`func (o *TimerInfoDtoProgramInfo) SetNormalizationGainNil(b bool)`

 SetNormalizationGainNil sets the value for NormalizationGain to be an explicit nil

### UnsetNormalizationGain
`func (o *TimerInfoDtoProgramInfo) UnsetNormalizationGain()`

UnsetNormalizationGain ensures that no value is present for NormalizationGain, not even an explicit nil
### GetCurrentProgram

`func (o *TimerInfoDtoProgramInfo) GetCurrentProgram() BaseItemDtoCurrentProgram`

GetCurrentProgram returns the CurrentProgram field if non-nil, zero value otherwise.

### GetCurrentProgramOk

`func (o *TimerInfoDtoProgramInfo) GetCurrentProgramOk() (*BaseItemDtoCurrentProgram, bool)`

GetCurrentProgramOk returns a tuple with the CurrentProgram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProgram

`func (o *TimerInfoDtoProgramInfo) SetCurrentProgram(v BaseItemDtoCurrentProgram)`

SetCurrentProgram sets CurrentProgram field to given value.

### HasCurrentProgram

`func (o *TimerInfoDtoProgramInfo) HasCurrentProgram() bool`

HasCurrentProgram returns a boolean if a field has been set.

### SetCurrentProgramNil

`func (o *TimerInfoDtoProgramInfo) SetCurrentProgramNil(b bool)`

 SetCurrentProgramNil sets the value for CurrentProgram to be an explicit nil

### UnsetCurrentProgram
`func (o *TimerInfoDtoProgramInfo) UnsetCurrentProgram()`

UnsetCurrentProgram ensures that no value is present for CurrentProgram, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


