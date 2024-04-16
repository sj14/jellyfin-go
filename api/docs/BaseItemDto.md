# BaseItemDto

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
**ExtraType** | Pointer to **NullableString** |  | [optional] 
**AirsBeforeSeasonNumber** | Pointer to **NullableInt32** |  | [optional] 
**AirsAfterSeasonNumber** | Pointer to **NullableInt32** |  | [optional] 
**AirsBeforeEpisodeNumber** | Pointer to **NullableInt32** |  | [optional] 
**CanDelete** | Pointer to **NullableBool** |  | [optional] 
**CanDownload** | Pointer to **NullableBool** |  | [optional] 
**HasSubtitles** | Pointer to **NullableBool** |  | [optional] 
**PreferredMetadataLanguage** | Pointer to **NullableString** |  | [optional] 
**PreferredMetadataCountryCode** | Pointer to **NullableString** |  | [optional] 
**SupportsSync** | Pointer to **NullableBool** | Gets or sets a value indicating whether [supports synchronize]. | [optional] 
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
**Type** | Pointer to [**BaseItemKind**](BaseItemKind.md) | Gets or sets the type. | [optional] 
**People** | Pointer to [**[]BaseItemPerson**](BaseItemPerson.md) | Gets or sets the people. | [optional] 
**Studios** | Pointer to [**[]NameGuidPair**](NameGuidPair.md) | Gets or sets the studios. | [optional] 
**GenreItems** | Pointer to [**[]NameGuidPair**](NameGuidPair.md) |  | [optional] 
**ParentLogoItemId** | Pointer to **NullableString** | Gets or sets wether the item has a logo, this will hold the Id of the Parent that has one. | [optional] 
**ParentBackdropItemId** | Pointer to **NullableString** | Gets or sets wether the item has any backdrops, this will hold the Id of the Parent that has one. | [optional] 
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
**CollectionType** | Pointer to **NullableString** | Gets or sets the type of the collection. | [optional] 
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
**ParentArtItemId** | Pointer to **NullableString** | Gets or sets wether the item has fan art, this will hold the Id of the Parent that has one. | [optional] 
**ParentArtImageTag** | Pointer to **NullableString** | Gets or sets the parent art image tag. | [optional] 
**SeriesThumbImageTag** | Pointer to **NullableString** | Gets or sets the series thumb image tag. | [optional] 
**ImageBlurHashes** | Pointer to [**NullableBaseItemDtoImageBlurHashes**](BaseItemDtoImageBlurHashes.md) |  | [optional] 
**SeriesStudio** | Pointer to **NullableString** | Gets or sets the series studio. | [optional] 
**ParentThumbItemId** | Pointer to **NullableString** | Gets or sets the parent thumb item id. | [optional] 
**ParentThumbImageTag** | Pointer to **NullableString** | Gets or sets the parent thumb image tag. | [optional] 
**ParentPrimaryImageItemId** | Pointer to **NullableString** | Gets or sets the parent primary image item identifier. | [optional] 
**ParentPrimaryImageTag** | Pointer to **NullableString** | Gets or sets the parent primary image tag. | [optional] 
**Chapters** | Pointer to [**[]ChapterInfo**](ChapterInfo.md) | Gets or sets the chapters. | [optional] 
**LocationType** | Pointer to [**NullableLocationType**](LocationType.md) | Gets or sets the type of the location. | [optional] 
**IsoType** | Pointer to [**NullableIsoType**](IsoType.md) | Gets or sets the type of the iso. | [optional] 
**MediaType** | Pointer to **NullableString** | Gets or sets the type of the media. | [optional] 
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
**CurrentProgram** | Pointer to [**NullableBaseItemDtoCurrentProgram**](BaseItemDtoCurrentProgram.md) |  | [optional] 

## Methods

### NewBaseItemDto

`func NewBaseItemDto() *BaseItemDto`

NewBaseItemDto instantiates a new BaseItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseItemDtoWithDefaults

`func NewBaseItemDtoWithDefaults() *BaseItemDto`

NewBaseItemDtoWithDefaults instantiates a new BaseItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BaseItemDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseItemDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseItemDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseItemDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BaseItemDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BaseItemDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOriginalTitle

`func (o *BaseItemDto) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *BaseItemDto) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *BaseItemDto) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *BaseItemDto) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### SetOriginalTitleNil

`func (o *BaseItemDto) SetOriginalTitleNil(b bool)`

 SetOriginalTitleNil sets the value for OriginalTitle to be an explicit nil

### UnsetOriginalTitle
`func (o *BaseItemDto) UnsetOriginalTitle()`

UnsetOriginalTitle ensures that no value is present for OriginalTitle, not even an explicit nil
### GetServerId

`func (o *BaseItemDto) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *BaseItemDto) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *BaseItemDto) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *BaseItemDto) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *BaseItemDto) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *BaseItemDto) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetId

`func (o *BaseItemDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseItemDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseItemDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseItemDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEtag

`func (o *BaseItemDto) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *BaseItemDto) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *BaseItemDto) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *BaseItemDto) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### SetEtagNil

`func (o *BaseItemDto) SetEtagNil(b bool)`

 SetEtagNil sets the value for Etag to be an explicit nil

### UnsetEtag
`func (o *BaseItemDto) UnsetEtag()`

UnsetEtag ensures that no value is present for Etag, not even an explicit nil
### GetSourceType

`func (o *BaseItemDto) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *BaseItemDto) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *BaseItemDto) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *BaseItemDto) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *BaseItemDto) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *BaseItemDto) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetPlaylistItemId

`func (o *BaseItemDto) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *BaseItemDto) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *BaseItemDto) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *BaseItemDto) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### SetPlaylistItemIdNil

`func (o *BaseItemDto) SetPlaylistItemIdNil(b bool)`

 SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil

### UnsetPlaylistItemId
`func (o *BaseItemDto) UnsetPlaylistItemId()`

UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil
### GetDateCreated

`func (o *BaseItemDto) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *BaseItemDto) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *BaseItemDto) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *BaseItemDto) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *BaseItemDto) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *BaseItemDto) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateLastMediaAdded

`func (o *BaseItemDto) GetDateLastMediaAdded() time.Time`

GetDateLastMediaAdded returns the DateLastMediaAdded field if non-nil, zero value otherwise.

### GetDateLastMediaAddedOk

`func (o *BaseItemDto) GetDateLastMediaAddedOk() (*time.Time, bool)`

GetDateLastMediaAddedOk returns a tuple with the DateLastMediaAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastMediaAdded

`func (o *BaseItemDto) SetDateLastMediaAdded(v time.Time)`

SetDateLastMediaAdded sets DateLastMediaAdded field to given value.

### HasDateLastMediaAdded

`func (o *BaseItemDto) HasDateLastMediaAdded() bool`

HasDateLastMediaAdded returns a boolean if a field has been set.

### SetDateLastMediaAddedNil

`func (o *BaseItemDto) SetDateLastMediaAddedNil(b bool)`

 SetDateLastMediaAddedNil sets the value for DateLastMediaAdded to be an explicit nil

### UnsetDateLastMediaAdded
`func (o *BaseItemDto) UnsetDateLastMediaAdded()`

UnsetDateLastMediaAdded ensures that no value is present for DateLastMediaAdded, not even an explicit nil
### GetExtraType

`func (o *BaseItemDto) GetExtraType() string`

GetExtraType returns the ExtraType field if non-nil, zero value otherwise.

### GetExtraTypeOk

`func (o *BaseItemDto) GetExtraTypeOk() (*string, bool)`

GetExtraTypeOk returns a tuple with the ExtraType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraType

`func (o *BaseItemDto) SetExtraType(v string)`

SetExtraType sets ExtraType field to given value.

### HasExtraType

`func (o *BaseItemDto) HasExtraType() bool`

HasExtraType returns a boolean if a field has been set.

### SetExtraTypeNil

`func (o *BaseItemDto) SetExtraTypeNil(b bool)`

 SetExtraTypeNil sets the value for ExtraType to be an explicit nil

### UnsetExtraType
`func (o *BaseItemDto) UnsetExtraType()`

UnsetExtraType ensures that no value is present for ExtraType, not even an explicit nil
### GetAirsBeforeSeasonNumber

`func (o *BaseItemDto) GetAirsBeforeSeasonNumber() int32`

GetAirsBeforeSeasonNumber returns the AirsBeforeSeasonNumber field if non-nil, zero value otherwise.

### GetAirsBeforeSeasonNumberOk

`func (o *BaseItemDto) GetAirsBeforeSeasonNumberOk() (*int32, bool)`

GetAirsBeforeSeasonNumberOk returns a tuple with the AirsBeforeSeasonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirsBeforeSeasonNumber

`func (o *BaseItemDto) SetAirsBeforeSeasonNumber(v int32)`

SetAirsBeforeSeasonNumber sets AirsBeforeSeasonNumber field to given value.

### HasAirsBeforeSeasonNumber

`func (o *BaseItemDto) HasAirsBeforeSeasonNumber() bool`

HasAirsBeforeSeasonNumber returns a boolean if a field has been set.

### SetAirsBeforeSeasonNumberNil

`func (o *BaseItemDto) SetAirsBeforeSeasonNumberNil(b bool)`

 SetAirsBeforeSeasonNumberNil sets the value for AirsBeforeSeasonNumber to be an explicit nil

### UnsetAirsBeforeSeasonNumber
`func (o *BaseItemDto) UnsetAirsBeforeSeasonNumber()`

UnsetAirsBeforeSeasonNumber ensures that no value is present for AirsBeforeSeasonNumber, not even an explicit nil
### GetAirsAfterSeasonNumber

`func (o *BaseItemDto) GetAirsAfterSeasonNumber() int32`

GetAirsAfterSeasonNumber returns the AirsAfterSeasonNumber field if non-nil, zero value otherwise.

### GetAirsAfterSeasonNumberOk

`func (o *BaseItemDto) GetAirsAfterSeasonNumberOk() (*int32, bool)`

GetAirsAfterSeasonNumberOk returns a tuple with the AirsAfterSeasonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirsAfterSeasonNumber

`func (o *BaseItemDto) SetAirsAfterSeasonNumber(v int32)`

SetAirsAfterSeasonNumber sets AirsAfterSeasonNumber field to given value.

### HasAirsAfterSeasonNumber

`func (o *BaseItemDto) HasAirsAfterSeasonNumber() bool`

HasAirsAfterSeasonNumber returns a boolean if a field has been set.

### SetAirsAfterSeasonNumberNil

`func (o *BaseItemDto) SetAirsAfterSeasonNumberNil(b bool)`

 SetAirsAfterSeasonNumberNil sets the value for AirsAfterSeasonNumber to be an explicit nil

### UnsetAirsAfterSeasonNumber
`func (o *BaseItemDto) UnsetAirsAfterSeasonNumber()`

UnsetAirsAfterSeasonNumber ensures that no value is present for AirsAfterSeasonNumber, not even an explicit nil
### GetAirsBeforeEpisodeNumber

`func (o *BaseItemDto) GetAirsBeforeEpisodeNumber() int32`

GetAirsBeforeEpisodeNumber returns the AirsBeforeEpisodeNumber field if non-nil, zero value otherwise.

### GetAirsBeforeEpisodeNumberOk

`func (o *BaseItemDto) GetAirsBeforeEpisodeNumberOk() (*int32, bool)`

GetAirsBeforeEpisodeNumberOk returns a tuple with the AirsBeforeEpisodeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirsBeforeEpisodeNumber

`func (o *BaseItemDto) SetAirsBeforeEpisodeNumber(v int32)`

SetAirsBeforeEpisodeNumber sets AirsBeforeEpisodeNumber field to given value.

### HasAirsBeforeEpisodeNumber

`func (o *BaseItemDto) HasAirsBeforeEpisodeNumber() bool`

HasAirsBeforeEpisodeNumber returns a boolean if a field has been set.

### SetAirsBeforeEpisodeNumberNil

`func (o *BaseItemDto) SetAirsBeforeEpisodeNumberNil(b bool)`

 SetAirsBeforeEpisodeNumberNil sets the value for AirsBeforeEpisodeNumber to be an explicit nil

### UnsetAirsBeforeEpisodeNumber
`func (o *BaseItemDto) UnsetAirsBeforeEpisodeNumber()`

UnsetAirsBeforeEpisodeNumber ensures that no value is present for AirsBeforeEpisodeNumber, not even an explicit nil
### GetCanDelete

`func (o *BaseItemDto) GetCanDelete() bool`

GetCanDelete returns the CanDelete field if non-nil, zero value otherwise.

### GetCanDeleteOk

`func (o *BaseItemDto) GetCanDeleteOk() (*bool, bool)`

GetCanDeleteOk returns a tuple with the CanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelete

`func (o *BaseItemDto) SetCanDelete(v bool)`

SetCanDelete sets CanDelete field to given value.

### HasCanDelete

`func (o *BaseItemDto) HasCanDelete() bool`

HasCanDelete returns a boolean if a field has been set.

### SetCanDeleteNil

`func (o *BaseItemDto) SetCanDeleteNil(b bool)`

 SetCanDeleteNil sets the value for CanDelete to be an explicit nil

### UnsetCanDelete
`func (o *BaseItemDto) UnsetCanDelete()`

UnsetCanDelete ensures that no value is present for CanDelete, not even an explicit nil
### GetCanDownload

`func (o *BaseItemDto) GetCanDownload() bool`

GetCanDownload returns the CanDownload field if non-nil, zero value otherwise.

### GetCanDownloadOk

`func (o *BaseItemDto) GetCanDownloadOk() (*bool, bool)`

GetCanDownloadOk returns a tuple with the CanDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDownload

`func (o *BaseItemDto) SetCanDownload(v bool)`

SetCanDownload sets CanDownload field to given value.

### HasCanDownload

`func (o *BaseItemDto) HasCanDownload() bool`

HasCanDownload returns a boolean if a field has been set.

### SetCanDownloadNil

`func (o *BaseItemDto) SetCanDownloadNil(b bool)`

 SetCanDownloadNil sets the value for CanDownload to be an explicit nil

### UnsetCanDownload
`func (o *BaseItemDto) UnsetCanDownload()`

UnsetCanDownload ensures that no value is present for CanDownload, not even an explicit nil
### GetHasSubtitles

`func (o *BaseItemDto) GetHasSubtitles() bool`

GetHasSubtitles returns the HasSubtitles field if non-nil, zero value otherwise.

### GetHasSubtitlesOk

`func (o *BaseItemDto) GetHasSubtitlesOk() (*bool, bool)`

GetHasSubtitlesOk returns a tuple with the HasSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSubtitles

`func (o *BaseItemDto) SetHasSubtitles(v bool)`

SetHasSubtitles sets HasSubtitles field to given value.

### HasHasSubtitles

`func (o *BaseItemDto) HasHasSubtitles() bool`

HasHasSubtitles returns a boolean if a field has been set.

### SetHasSubtitlesNil

`func (o *BaseItemDto) SetHasSubtitlesNil(b bool)`

 SetHasSubtitlesNil sets the value for HasSubtitles to be an explicit nil

### UnsetHasSubtitles
`func (o *BaseItemDto) UnsetHasSubtitles()`

UnsetHasSubtitles ensures that no value is present for HasSubtitles, not even an explicit nil
### GetPreferredMetadataLanguage

`func (o *BaseItemDto) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *BaseItemDto) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *BaseItemDto) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *BaseItemDto) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### SetPreferredMetadataLanguageNil

`func (o *BaseItemDto) SetPreferredMetadataLanguageNil(b bool)`

 SetPreferredMetadataLanguageNil sets the value for PreferredMetadataLanguage to be an explicit nil

### UnsetPreferredMetadataLanguage
`func (o *BaseItemDto) UnsetPreferredMetadataLanguage()`

UnsetPreferredMetadataLanguage ensures that no value is present for PreferredMetadataLanguage, not even an explicit nil
### GetPreferredMetadataCountryCode

`func (o *BaseItemDto) GetPreferredMetadataCountryCode() string`

GetPreferredMetadataCountryCode returns the PreferredMetadataCountryCode field if non-nil, zero value otherwise.

### GetPreferredMetadataCountryCodeOk

`func (o *BaseItemDto) GetPreferredMetadataCountryCodeOk() (*string, bool)`

GetPreferredMetadataCountryCodeOk returns a tuple with the PreferredMetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataCountryCode

`func (o *BaseItemDto) SetPreferredMetadataCountryCode(v string)`

SetPreferredMetadataCountryCode sets PreferredMetadataCountryCode field to given value.

### HasPreferredMetadataCountryCode

`func (o *BaseItemDto) HasPreferredMetadataCountryCode() bool`

HasPreferredMetadataCountryCode returns a boolean if a field has been set.

### SetPreferredMetadataCountryCodeNil

`func (o *BaseItemDto) SetPreferredMetadataCountryCodeNil(b bool)`

 SetPreferredMetadataCountryCodeNil sets the value for PreferredMetadataCountryCode to be an explicit nil

### UnsetPreferredMetadataCountryCode
`func (o *BaseItemDto) UnsetPreferredMetadataCountryCode()`

UnsetPreferredMetadataCountryCode ensures that no value is present for PreferredMetadataCountryCode, not even an explicit nil
### GetSupportsSync

`func (o *BaseItemDto) GetSupportsSync() bool`

GetSupportsSync returns the SupportsSync field if non-nil, zero value otherwise.

### GetSupportsSyncOk

`func (o *BaseItemDto) GetSupportsSyncOk() (*bool, bool)`

GetSupportsSyncOk returns a tuple with the SupportsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSync

`func (o *BaseItemDto) SetSupportsSync(v bool)`

SetSupportsSync sets SupportsSync field to given value.

### HasSupportsSync

`func (o *BaseItemDto) HasSupportsSync() bool`

HasSupportsSync returns a boolean if a field has been set.

### SetSupportsSyncNil

`func (o *BaseItemDto) SetSupportsSyncNil(b bool)`

 SetSupportsSyncNil sets the value for SupportsSync to be an explicit nil

### UnsetSupportsSync
`func (o *BaseItemDto) UnsetSupportsSync()`

UnsetSupportsSync ensures that no value is present for SupportsSync, not even an explicit nil
### GetContainer

`func (o *BaseItemDto) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *BaseItemDto) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *BaseItemDto) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *BaseItemDto) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *BaseItemDto) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *BaseItemDto) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetSortName

`func (o *BaseItemDto) GetSortName() string`

GetSortName returns the SortName field if non-nil, zero value otherwise.

### GetSortNameOk

`func (o *BaseItemDto) GetSortNameOk() (*string, bool)`

GetSortNameOk returns a tuple with the SortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortName

`func (o *BaseItemDto) SetSortName(v string)`

SetSortName sets SortName field to given value.

### HasSortName

`func (o *BaseItemDto) HasSortName() bool`

HasSortName returns a boolean if a field has been set.

### SetSortNameNil

`func (o *BaseItemDto) SetSortNameNil(b bool)`

 SetSortNameNil sets the value for SortName to be an explicit nil

### UnsetSortName
`func (o *BaseItemDto) UnsetSortName()`

UnsetSortName ensures that no value is present for SortName, not even an explicit nil
### GetForcedSortName

`func (o *BaseItemDto) GetForcedSortName() string`

GetForcedSortName returns the ForcedSortName field if non-nil, zero value otherwise.

### GetForcedSortNameOk

`func (o *BaseItemDto) GetForcedSortNameOk() (*string, bool)`

GetForcedSortNameOk returns a tuple with the ForcedSortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSortName

`func (o *BaseItemDto) SetForcedSortName(v string)`

SetForcedSortName sets ForcedSortName field to given value.

### HasForcedSortName

`func (o *BaseItemDto) HasForcedSortName() bool`

HasForcedSortName returns a boolean if a field has been set.

### SetForcedSortNameNil

`func (o *BaseItemDto) SetForcedSortNameNil(b bool)`

 SetForcedSortNameNil sets the value for ForcedSortName to be an explicit nil

### UnsetForcedSortName
`func (o *BaseItemDto) UnsetForcedSortName()`

UnsetForcedSortName ensures that no value is present for ForcedSortName, not even an explicit nil
### GetVideo3DFormat

`func (o *BaseItemDto) GetVideo3DFormat() Video3DFormat`

GetVideo3DFormat returns the Video3DFormat field if non-nil, zero value otherwise.

### GetVideo3DFormatOk

`func (o *BaseItemDto) GetVideo3DFormatOk() (*Video3DFormat, bool)`

GetVideo3DFormatOk returns a tuple with the Video3DFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo3DFormat

`func (o *BaseItemDto) SetVideo3DFormat(v Video3DFormat)`

SetVideo3DFormat sets Video3DFormat field to given value.

### HasVideo3DFormat

`func (o *BaseItemDto) HasVideo3DFormat() bool`

HasVideo3DFormat returns a boolean if a field has been set.

### SetVideo3DFormatNil

`func (o *BaseItemDto) SetVideo3DFormatNil(b bool)`

 SetVideo3DFormatNil sets the value for Video3DFormat to be an explicit nil

### UnsetVideo3DFormat
`func (o *BaseItemDto) UnsetVideo3DFormat()`

UnsetVideo3DFormat ensures that no value is present for Video3DFormat, not even an explicit nil
### GetPremiereDate

`func (o *BaseItemDto) GetPremiereDate() time.Time`

GetPremiereDate returns the PremiereDate field if non-nil, zero value otherwise.

### GetPremiereDateOk

`func (o *BaseItemDto) GetPremiereDateOk() (*time.Time, bool)`

GetPremiereDateOk returns a tuple with the PremiereDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiereDate

`func (o *BaseItemDto) SetPremiereDate(v time.Time)`

SetPremiereDate sets PremiereDate field to given value.

### HasPremiereDate

`func (o *BaseItemDto) HasPremiereDate() bool`

HasPremiereDate returns a boolean if a field has been set.

### SetPremiereDateNil

`func (o *BaseItemDto) SetPremiereDateNil(b bool)`

 SetPremiereDateNil sets the value for PremiereDate to be an explicit nil

### UnsetPremiereDate
`func (o *BaseItemDto) UnsetPremiereDate()`

UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
### GetExternalUrls

`func (o *BaseItemDto) GetExternalUrls() []ExternalUrl`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *BaseItemDto) GetExternalUrlsOk() (*[]ExternalUrl, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *BaseItemDto) SetExternalUrls(v []ExternalUrl)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *BaseItemDto) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### SetExternalUrlsNil

`func (o *BaseItemDto) SetExternalUrlsNil(b bool)`

 SetExternalUrlsNil sets the value for ExternalUrls to be an explicit nil

### UnsetExternalUrls
`func (o *BaseItemDto) UnsetExternalUrls()`

UnsetExternalUrls ensures that no value is present for ExternalUrls, not even an explicit nil
### GetMediaSources

`func (o *BaseItemDto) GetMediaSources() []MediaSourceInfo`

GetMediaSources returns the MediaSources field if non-nil, zero value otherwise.

### GetMediaSourcesOk

`func (o *BaseItemDto) GetMediaSourcesOk() (*[]MediaSourceInfo, bool)`

GetMediaSourcesOk returns a tuple with the MediaSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSources

`func (o *BaseItemDto) SetMediaSources(v []MediaSourceInfo)`

SetMediaSources sets MediaSources field to given value.

### HasMediaSources

`func (o *BaseItemDto) HasMediaSources() bool`

HasMediaSources returns a boolean if a field has been set.

### SetMediaSourcesNil

`func (o *BaseItemDto) SetMediaSourcesNil(b bool)`

 SetMediaSourcesNil sets the value for MediaSources to be an explicit nil

### UnsetMediaSources
`func (o *BaseItemDto) UnsetMediaSources()`

UnsetMediaSources ensures that no value is present for MediaSources, not even an explicit nil
### GetCriticRating

`func (o *BaseItemDto) GetCriticRating() float32`

GetCriticRating returns the CriticRating field if non-nil, zero value otherwise.

### GetCriticRatingOk

`func (o *BaseItemDto) GetCriticRatingOk() (*float32, bool)`

GetCriticRatingOk returns a tuple with the CriticRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticRating

`func (o *BaseItemDto) SetCriticRating(v float32)`

SetCriticRating sets CriticRating field to given value.

### HasCriticRating

`func (o *BaseItemDto) HasCriticRating() bool`

HasCriticRating returns a boolean if a field has been set.

### SetCriticRatingNil

`func (o *BaseItemDto) SetCriticRatingNil(b bool)`

 SetCriticRatingNil sets the value for CriticRating to be an explicit nil

### UnsetCriticRating
`func (o *BaseItemDto) UnsetCriticRating()`

UnsetCriticRating ensures that no value is present for CriticRating, not even an explicit nil
### GetProductionLocations

`func (o *BaseItemDto) GetProductionLocations() []string`

GetProductionLocations returns the ProductionLocations field if non-nil, zero value otherwise.

### GetProductionLocationsOk

`func (o *BaseItemDto) GetProductionLocationsOk() (*[]string, bool)`

GetProductionLocationsOk returns a tuple with the ProductionLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionLocations

`func (o *BaseItemDto) SetProductionLocations(v []string)`

SetProductionLocations sets ProductionLocations field to given value.

### HasProductionLocations

`func (o *BaseItemDto) HasProductionLocations() bool`

HasProductionLocations returns a boolean if a field has been set.

### SetProductionLocationsNil

`func (o *BaseItemDto) SetProductionLocationsNil(b bool)`

 SetProductionLocationsNil sets the value for ProductionLocations to be an explicit nil

### UnsetProductionLocations
`func (o *BaseItemDto) UnsetProductionLocations()`

UnsetProductionLocations ensures that no value is present for ProductionLocations, not even an explicit nil
### GetPath

`func (o *BaseItemDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BaseItemDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BaseItemDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BaseItemDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *BaseItemDto) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *BaseItemDto) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetEnableMediaSourceDisplay

`func (o *BaseItemDto) GetEnableMediaSourceDisplay() bool`

GetEnableMediaSourceDisplay returns the EnableMediaSourceDisplay field if non-nil, zero value otherwise.

### GetEnableMediaSourceDisplayOk

`func (o *BaseItemDto) GetEnableMediaSourceDisplayOk() (*bool, bool)`

GetEnableMediaSourceDisplayOk returns a tuple with the EnableMediaSourceDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMediaSourceDisplay

`func (o *BaseItemDto) SetEnableMediaSourceDisplay(v bool)`

SetEnableMediaSourceDisplay sets EnableMediaSourceDisplay field to given value.

### HasEnableMediaSourceDisplay

`func (o *BaseItemDto) HasEnableMediaSourceDisplay() bool`

HasEnableMediaSourceDisplay returns a boolean if a field has been set.

### SetEnableMediaSourceDisplayNil

`func (o *BaseItemDto) SetEnableMediaSourceDisplayNil(b bool)`

 SetEnableMediaSourceDisplayNil sets the value for EnableMediaSourceDisplay to be an explicit nil

### UnsetEnableMediaSourceDisplay
`func (o *BaseItemDto) UnsetEnableMediaSourceDisplay()`

UnsetEnableMediaSourceDisplay ensures that no value is present for EnableMediaSourceDisplay, not even an explicit nil
### GetOfficialRating

`func (o *BaseItemDto) GetOfficialRating() string`

GetOfficialRating returns the OfficialRating field if non-nil, zero value otherwise.

### GetOfficialRatingOk

`func (o *BaseItemDto) GetOfficialRatingOk() (*string, bool)`

GetOfficialRatingOk returns a tuple with the OfficialRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialRating

`func (o *BaseItemDto) SetOfficialRating(v string)`

SetOfficialRating sets OfficialRating field to given value.

### HasOfficialRating

`func (o *BaseItemDto) HasOfficialRating() bool`

HasOfficialRating returns a boolean if a field has been set.

### SetOfficialRatingNil

`func (o *BaseItemDto) SetOfficialRatingNil(b bool)`

 SetOfficialRatingNil sets the value for OfficialRating to be an explicit nil

### UnsetOfficialRating
`func (o *BaseItemDto) UnsetOfficialRating()`

UnsetOfficialRating ensures that no value is present for OfficialRating, not even an explicit nil
### GetCustomRating

`func (o *BaseItemDto) GetCustomRating() string`

GetCustomRating returns the CustomRating field if non-nil, zero value otherwise.

### GetCustomRatingOk

`func (o *BaseItemDto) GetCustomRatingOk() (*string, bool)`

GetCustomRatingOk returns a tuple with the CustomRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRating

`func (o *BaseItemDto) SetCustomRating(v string)`

SetCustomRating sets CustomRating field to given value.

### HasCustomRating

`func (o *BaseItemDto) HasCustomRating() bool`

HasCustomRating returns a boolean if a field has been set.

### SetCustomRatingNil

`func (o *BaseItemDto) SetCustomRatingNil(b bool)`

 SetCustomRatingNil sets the value for CustomRating to be an explicit nil

### UnsetCustomRating
`func (o *BaseItemDto) UnsetCustomRating()`

UnsetCustomRating ensures that no value is present for CustomRating, not even an explicit nil
### GetChannelId

`func (o *BaseItemDto) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *BaseItemDto) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *BaseItemDto) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *BaseItemDto) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *BaseItemDto) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *BaseItemDto) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetChannelName

`func (o *BaseItemDto) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *BaseItemDto) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *BaseItemDto) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *BaseItemDto) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### SetChannelNameNil

`func (o *BaseItemDto) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *BaseItemDto) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetOverview

`func (o *BaseItemDto) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *BaseItemDto) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *BaseItemDto) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *BaseItemDto) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *BaseItemDto) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *BaseItemDto) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetTaglines

`func (o *BaseItemDto) GetTaglines() []string`

GetTaglines returns the Taglines field if non-nil, zero value otherwise.

### GetTaglinesOk

`func (o *BaseItemDto) GetTaglinesOk() (*[]string, bool)`

GetTaglinesOk returns a tuple with the Taglines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaglines

`func (o *BaseItemDto) SetTaglines(v []string)`

SetTaglines sets Taglines field to given value.

### HasTaglines

`func (o *BaseItemDto) HasTaglines() bool`

HasTaglines returns a boolean if a field has been set.

### SetTaglinesNil

`func (o *BaseItemDto) SetTaglinesNil(b bool)`

 SetTaglinesNil sets the value for Taglines to be an explicit nil

### UnsetTaglines
`func (o *BaseItemDto) UnsetTaglines()`

UnsetTaglines ensures that no value is present for Taglines, not even an explicit nil
### GetGenres

`func (o *BaseItemDto) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *BaseItemDto) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *BaseItemDto) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *BaseItemDto) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *BaseItemDto) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *BaseItemDto) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetCommunityRating

`func (o *BaseItemDto) GetCommunityRating() float32`

GetCommunityRating returns the CommunityRating field if non-nil, zero value otherwise.

### GetCommunityRatingOk

`func (o *BaseItemDto) GetCommunityRatingOk() (*float32, bool)`

GetCommunityRatingOk returns a tuple with the CommunityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityRating

`func (o *BaseItemDto) SetCommunityRating(v float32)`

SetCommunityRating sets CommunityRating field to given value.

### HasCommunityRating

`func (o *BaseItemDto) HasCommunityRating() bool`

HasCommunityRating returns a boolean if a field has been set.

### SetCommunityRatingNil

`func (o *BaseItemDto) SetCommunityRatingNil(b bool)`

 SetCommunityRatingNil sets the value for CommunityRating to be an explicit nil

### UnsetCommunityRating
`func (o *BaseItemDto) UnsetCommunityRating()`

UnsetCommunityRating ensures that no value is present for CommunityRating, not even an explicit nil
### GetCumulativeRunTimeTicks

`func (o *BaseItemDto) GetCumulativeRunTimeTicks() int64`

GetCumulativeRunTimeTicks returns the CumulativeRunTimeTicks field if non-nil, zero value otherwise.

### GetCumulativeRunTimeTicksOk

`func (o *BaseItemDto) GetCumulativeRunTimeTicksOk() (*int64, bool)`

GetCumulativeRunTimeTicksOk returns a tuple with the CumulativeRunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeRunTimeTicks

`func (o *BaseItemDto) SetCumulativeRunTimeTicks(v int64)`

SetCumulativeRunTimeTicks sets CumulativeRunTimeTicks field to given value.

### HasCumulativeRunTimeTicks

`func (o *BaseItemDto) HasCumulativeRunTimeTicks() bool`

HasCumulativeRunTimeTicks returns a boolean if a field has been set.

### SetCumulativeRunTimeTicksNil

`func (o *BaseItemDto) SetCumulativeRunTimeTicksNil(b bool)`

 SetCumulativeRunTimeTicksNil sets the value for CumulativeRunTimeTicks to be an explicit nil

### UnsetCumulativeRunTimeTicks
`func (o *BaseItemDto) UnsetCumulativeRunTimeTicks()`

UnsetCumulativeRunTimeTicks ensures that no value is present for CumulativeRunTimeTicks, not even an explicit nil
### GetRunTimeTicks

`func (o *BaseItemDto) GetRunTimeTicks() int64`

GetRunTimeTicks returns the RunTimeTicks field if non-nil, zero value otherwise.

### GetRunTimeTicksOk

`func (o *BaseItemDto) GetRunTimeTicksOk() (*int64, bool)`

GetRunTimeTicksOk returns a tuple with the RunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeTicks

`func (o *BaseItemDto) SetRunTimeTicks(v int64)`

SetRunTimeTicks sets RunTimeTicks field to given value.

### HasRunTimeTicks

`func (o *BaseItemDto) HasRunTimeTicks() bool`

HasRunTimeTicks returns a boolean if a field has been set.

### SetRunTimeTicksNil

`func (o *BaseItemDto) SetRunTimeTicksNil(b bool)`

 SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil

### UnsetRunTimeTicks
`func (o *BaseItemDto) UnsetRunTimeTicks()`

UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
### GetPlayAccess

`func (o *BaseItemDto) GetPlayAccess() PlayAccess`

GetPlayAccess returns the PlayAccess field if non-nil, zero value otherwise.

### GetPlayAccessOk

`func (o *BaseItemDto) GetPlayAccessOk() (*PlayAccess, bool)`

GetPlayAccessOk returns a tuple with the PlayAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayAccess

`func (o *BaseItemDto) SetPlayAccess(v PlayAccess)`

SetPlayAccess sets PlayAccess field to given value.

### HasPlayAccess

`func (o *BaseItemDto) HasPlayAccess() bool`

HasPlayAccess returns a boolean if a field has been set.

### SetPlayAccessNil

`func (o *BaseItemDto) SetPlayAccessNil(b bool)`

 SetPlayAccessNil sets the value for PlayAccess to be an explicit nil

### UnsetPlayAccess
`func (o *BaseItemDto) UnsetPlayAccess()`

UnsetPlayAccess ensures that no value is present for PlayAccess, not even an explicit nil
### GetAspectRatio

`func (o *BaseItemDto) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *BaseItemDto) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *BaseItemDto) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *BaseItemDto) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### SetAspectRatioNil

`func (o *BaseItemDto) SetAspectRatioNil(b bool)`

 SetAspectRatioNil sets the value for AspectRatio to be an explicit nil

### UnsetAspectRatio
`func (o *BaseItemDto) UnsetAspectRatio()`

UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil
### GetProductionYear

`func (o *BaseItemDto) GetProductionYear() int32`

GetProductionYear returns the ProductionYear field if non-nil, zero value otherwise.

### GetProductionYearOk

`func (o *BaseItemDto) GetProductionYearOk() (*int32, bool)`

GetProductionYearOk returns a tuple with the ProductionYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionYear

`func (o *BaseItemDto) SetProductionYear(v int32)`

SetProductionYear sets ProductionYear field to given value.

### HasProductionYear

`func (o *BaseItemDto) HasProductionYear() bool`

HasProductionYear returns a boolean if a field has been set.

### SetProductionYearNil

`func (o *BaseItemDto) SetProductionYearNil(b bool)`

 SetProductionYearNil sets the value for ProductionYear to be an explicit nil

### UnsetProductionYear
`func (o *BaseItemDto) UnsetProductionYear()`

UnsetProductionYear ensures that no value is present for ProductionYear, not even an explicit nil
### GetIsPlaceHolder

`func (o *BaseItemDto) GetIsPlaceHolder() bool`

GetIsPlaceHolder returns the IsPlaceHolder field if non-nil, zero value otherwise.

### GetIsPlaceHolderOk

`func (o *BaseItemDto) GetIsPlaceHolderOk() (*bool, bool)`

GetIsPlaceHolderOk returns a tuple with the IsPlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlaceHolder

`func (o *BaseItemDto) SetIsPlaceHolder(v bool)`

SetIsPlaceHolder sets IsPlaceHolder field to given value.

### HasIsPlaceHolder

`func (o *BaseItemDto) HasIsPlaceHolder() bool`

HasIsPlaceHolder returns a boolean if a field has been set.

### SetIsPlaceHolderNil

`func (o *BaseItemDto) SetIsPlaceHolderNil(b bool)`

 SetIsPlaceHolderNil sets the value for IsPlaceHolder to be an explicit nil

### UnsetIsPlaceHolder
`func (o *BaseItemDto) UnsetIsPlaceHolder()`

UnsetIsPlaceHolder ensures that no value is present for IsPlaceHolder, not even an explicit nil
### GetNumber

`func (o *BaseItemDto) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *BaseItemDto) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *BaseItemDto) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *BaseItemDto) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *BaseItemDto) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *BaseItemDto) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetChannelNumber

`func (o *BaseItemDto) GetChannelNumber() string`

GetChannelNumber returns the ChannelNumber field if non-nil, zero value otherwise.

### GetChannelNumberOk

`func (o *BaseItemDto) GetChannelNumberOk() (*string, bool)`

GetChannelNumberOk returns a tuple with the ChannelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelNumber

`func (o *BaseItemDto) SetChannelNumber(v string)`

SetChannelNumber sets ChannelNumber field to given value.

### HasChannelNumber

`func (o *BaseItemDto) HasChannelNumber() bool`

HasChannelNumber returns a boolean if a field has been set.

### SetChannelNumberNil

`func (o *BaseItemDto) SetChannelNumberNil(b bool)`

 SetChannelNumberNil sets the value for ChannelNumber to be an explicit nil

### UnsetChannelNumber
`func (o *BaseItemDto) UnsetChannelNumber()`

UnsetChannelNumber ensures that no value is present for ChannelNumber, not even an explicit nil
### GetIndexNumber

`func (o *BaseItemDto) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *BaseItemDto) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *BaseItemDto) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *BaseItemDto) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *BaseItemDto) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *BaseItemDto) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetIndexNumberEnd

`func (o *BaseItemDto) GetIndexNumberEnd() int32`

GetIndexNumberEnd returns the IndexNumberEnd field if non-nil, zero value otherwise.

### GetIndexNumberEndOk

`func (o *BaseItemDto) GetIndexNumberEndOk() (*int32, bool)`

GetIndexNumberEndOk returns a tuple with the IndexNumberEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumberEnd

`func (o *BaseItemDto) SetIndexNumberEnd(v int32)`

SetIndexNumberEnd sets IndexNumberEnd field to given value.

### HasIndexNumberEnd

`func (o *BaseItemDto) HasIndexNumberEnd() bool`

HasIndexNumberEnd returns a boolean if a field has been set.

### SetIndexNumberEndNil

`func (o *BaseItemDto) SetIndexNumberEndNil(b bool)`

 SetIndexNumberEndNil sets the value for IndexNumberEnd to be an explicit nil

### UnsetIndexNumberEnd
`func (o *BaseItemDto) UnsetIndexNumberEnd()`

UnsetIndexNumberEnd ensures that no value is present for IndexNumberEnd, not even an explicit nil
### GetParentIndexNumber

`func (o *BaseItemDto) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *BaseItemDto) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *BaseItemDto) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *BaseItemDto) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *BaseItemDto) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *BaseItemDto) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetRemoteTrailers

`func (o *BaseItemDto) GetRemoteTrailers() []MediaUrl`

GetRemoteTrailers returns the RemoteTrailers field if non-nil, zero value otherwise.

### GetRemoteTrailersOk

`func (o *BaseItemDto) GetRemoteTrailersOk() (*[]MediaUrl, bool)`

GetRemoteTrailersOk returns a tuple with the RemoteTrailers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTrailers

`func (o *BaseItemDto) SetRemoteTrailers(v []MediaUrl)`

SetRemoteTrailers sets RemoteTrailers field to given value.

### HasRemoteTrailers

`func (o *BaseItemDto) HasRemoteTrailers() bool`

HasRemoteTrailers returns a boolean if a field has been set.

### SetRemoteTrailersNil

`func (o *BaseItemDto) SetRemoteTrailersNil(b bool)`

 SetRemoteTrailersNil sets the value for RemoteTrailers to be an explicit nil

### UnsetRemoteTrailers
`func (o *BaseItemDto) UnsetRemoteTrailers()`

UnsetRemoteTrailers ensures that no value is present for RemoteTrailers, not even an explicit nil
### GetProviderIds

`func (o *BaseItemDto) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *BaseItemDto) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *BaseItemDto) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *BaseItemDto) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### SetProviderIdsNil

`func (o *BaseItemDto) SetProviderIdsNil(b bool)`

 SetProviderIdsNil sets the value for ProviderIds to be an explicit nil

### UnsetProviderIds
`func (o *BaseItemDto) UnsetProviderIds()`

UnsetProviderIds ensures that no value is present for ProviderIds, not even an explicit nil
### GetIsHD

`func (o *BaseItemDto) GetIsHD() bool`

GetIsHD returns the IsHD field if non-nil, zero value otherwise.

### GetIsHDOk

`func (o *BaseItemDto) GetIsHDOk() (*bool, bool)`

GetIsHDOk returns a tuple with the IsHD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHD

`func (o *BaseItemDto) SetIsHD(v bool)`

SetIsHD sets IsHD field to given value.

### HasIsHD

`func (o *BaseItemDto) HasIsHD() bool`

HasIsHD returns a boolean if a field has been set.

### SetIsHDNil

`func (o *BaseItemDto) SetIsHDNil(b bool)`

 SetIsHDNil sets the value for IsHD to be an explicit nil

### UnsetIsHD
`func (o *BaseItemDto) UnsetIsHD()`

UnsetIsHD ensures that no value is present for IsHD, not even an explicit nil
### GetIsFolder

`func (o *BaseItemDto) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *BaseItemDto) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *BaseItemDto) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *BaseItemDto) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### SetIsFolderNil

`func (o *BaseItemDto) SetIsFolderNil(b bool)`

 SetIsFolderNil sets the value for IsFolder to be an explicit nil

### UnsetIsFolder
`func (o *BaseItemDto) UnsetIsFolder()`

UnsetIsFolder ensures that no value is present for IsFolder, not even an explicit nil
### GetParentId

`func (o *BaseItemDto) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BaseItemDto) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BaseItemDto) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BaseItemDto) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *BaseItemDto) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *BaseItemDto) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetType

`func (o *BaseItemDto) GetType() BaseItemKind`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseItemDto) GetTypeOk() (*BaseItemKind, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseItemDto) SetType(v BaseItemKind)`

SetType sets Type field to given value.

### HasType

`func (o *BaseItemDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPeople

`func (o *BaseItemDto) GetPeople() []BaseItemPerson`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *BaseItemDto) GetPeopleOk() (*[]BaseItemPerson, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *BaseItemDto) SetPeople(v []BaseItemPerson)`

SetPeople sets People field to given value.

### HasPeople

`func (o *BaseItemDto) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### SetPeopleNil

`func (o *BaseItemDto) SetPeopleNil(b bool)`

 SetPeopleNil sets the value for People to be an explicit nil

### UnsetPeople
`func (o *BaseItemDto) UnsetPeople()`

UnsetPeople ensures that no value is present for People, not even an explicit nil
### GetStudios

`func (o *BaseItemDto) GetStudios() []NameGuidPair`

GetStudios returns the Studios field if non-nil, zero value otherwise.

### GetStudiosOk

`func (o *BaseItemDto) GetStudiosOk() (*[]NameGuidPair, bool)`

GetStudiosOk returns a tuple with the Studios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudios

`func (o *BaseItemDto) SetStudios(v []NameGuidPair)`

SetStudios sets Studios field to given value.

### HasStudios

`func (o *BaseItemDto) HasStudios() bool`

HasStudios returns a boolean if a field has been set.

### SetStudiosNil

`func (o *BaseItemDto) SetStudiosNil(b bool)`

 SetStudiosNil sets the value for Studios to be an explicit nil

### UnsetStudios
`func (o *BaseItemDto) UnsetStudios()`

UnsetStudios ensures that no value is present for Studios, not even an explicit nil
### GetGenreItems

`func (o *BaseItemDto) GetGenreItems() []NameGuidPair`

GetGenreItems returns the GenreItems field if non-nil, zero value otherwise.

### GetGenreItemsOk

`func (o *BaseItemDto) GetGenreItemsOk() (*[]NameGuidPair, bool)`

GetGenreItemsOk returns a tuple with the GenreItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreItems

`func (o *BaseItemDto) SetGenreItems(v []NameGuidPair)`

SetGenreItems sets GenreItems field to given value.

### HasGenreItems

`func (o *BaseItemDto) HasGenreItems() bool`

HasGenreItems returns a boolean if a field has been set.

### SetGenreItemsNil

`func (o *BaseItemDto) SetGenreItemsNil(b bool)`

 SetGenreItemsNil sets the value for GenreItems to be an explicit nil

### UnsetGenreItems
`func (o *BaseItemDto) UnsetGenreItems()`

UnsetGenreItems ensures that no value is present for GenreItems, not even an explicit nil
### GetParentLogoItemId

`func (o *BaseItemDto) GetParentLogoItemId() string`

GetParentLogoItemId returns the ParentLogoItemId field if non-nil, zero value otherwise.

### GetParentLogoItemIdOk

`func (o *BaseItemDto) GetParentLogoItemIdOk() (*string, bool)`

GetParentLogoItemIdOk returns a tuple with the ParentLogoItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLogoItemId

`func (o *BaseItemDto) SetParentLogoItemId(v string)`

SetParentLogoItemId sets ParentLogoItemId field to given value.

### HasParentLogoItemId

`func (o *BaseItemDto) HasParentLogoItemId() bool`

HasParentLogoItemId returns a boolean if a field has been set.

### SetParentLogoItemIdNil

`func (o *BaseItemDto) SetParentLogoItemIdNil(b bool)`

 SetParentLogoItemIdNil sets the value for ParentLogoItemId to be an explicit nil

### UnsetParentLogoItemId
`func (o *BaseItemDto) UnsetParentLogoItemId()`

UnsetParentLogoItemId ensures that no value is present for ParentLogoItemId, not even an explicit nil
### GetParentBackdropItemId

`func (o *BaseItemDto) GetParentBackdropItemId() string`

GetParentBackdropItemId returns the ParentBackdropItemId field if non-nil, zero value otherwise.

### GetParentBackdropItemIdOk

`func (o *BaseItemDto) GetParentBackdropItemIdOk() (*string, bool)`

GetParentBackdropItemIdOk returns a tuple with the ParentBackdropItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropItemId

`func (o *BaseItemDto) SetParentBackdropItemId(v string)`

SetParentBackdropItemId sets ParentBackdropItemId field to given value.

### HasParentBackdropItemId

`func (o *BaseItemDto) HasParentBackdropItemId() bool`

HasParentBackdropItemId returns a boolean if a field has been set.

### SetParentBackdropItemIdNil

`func (o *BaseItemDto) SetParentBackdropItemIdNil(b bool)`

 SetParentBackdropItemIdNil sets the value for ParentBackdropItemId to be an explicit nil

### UnsetParentBackdropItemId
`func (o *BaseItemDto) UnsetParentBackdropItemId()`

UnsetParentBackdropItemId ensures that no value is present for ParentBackdropItemId, not even an explicit nil
### GetParentBackdropImageTags

`func (o *BaseItemDto) GetParentBackdropImageTags() []string`

GetParentBackdropImageTags returns the ParentBackdropImageTags field if non-nil, zero value otherwise.

### GetParentBackdropImageTagsOk

`func (o *BaseItemDto) GetParentBackdropImageTagsOk() (*[]string, bool)`

GetParentBackdropImageTagsOk returns a tuple with the ParentBackdropImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropImageTags

`func (o *BaseItemDto) SetParentBackdropImageTags(v []string)`

SetParentBackdropImageTags sets ParentBackdropImageTags field to given value.

### HasParentBackdropImageTags

`func (o *BaseItemDto) HasParentBackdropImageTags() bool`

HasParentBackdropImageTags returns a boolean if a field has been set.

### SetParentBackdropImageTagsNil

`func (o *BaseItemDto) SetParentBackdropImageTagsNil(b bool)`

 SetParentBackdropImageTagsNil sets the value for ParentBackdropImageTags to be an explicit nil

### UnsetParentBackdropImageTags
`func (o *BaseItemDto) UnsetParentBackdropImageTags()`

UnsetParentBackdropImageTags ensures that no value is present for ParentBackdropImageTags, not even an explicit nil
### GetLocalTrailerCount

`func (o *BaseItemDto) GetLocalTrailerCount() int32`

GetLocalTrailerCount returns the LocalTrailerCount field if non-nil, zero value otherwise.

### GetLocalTrailerCountOk

`func (o *BaseItemDto) GetLocalTrailerCountOk() (*int32, bool)`

GetLocalTrailerCountOk returns a tuple with the LocalTrailerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTrailerCount

`func (o *BaseItemDto) SetLocalTrailerCount(v int32)`

SetLocalTrailerCount sets LocalTrailerCount field to given value.

### HasLocalTrailerCount

`func (o *BaseItemDto) HasLocalTrailerCount() bool`

HasLocalTrailerCount returns a boolean if a field has been set.

### SetLocalTrailerCountNil

`func (o *BaseItemDto) SetLocalTrailerCountNil(b bool)`

 SetLocalTrailerCountNil sets the value for LocalTrailerCount to be an explicit nil

### UnsetLocalTrailerCount
`func (o *BaseItemDto) UnsetLocalTrailerCount()`

UnsetLocalTrailerCount ensures that no value is present for LocalTrailerCount, not even an explicit nil
### GetUserData

`func (o *BaseItemDto) GetUserData() BaseItemDtoUserData`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *BaseItemDto) GetUserDataOk() (*BaseItemDtoUserData, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *BaseItemDto) SetUserData(v BaseItemDtoUserData)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *BaseItemDto) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *BaseItemDto) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *BaseItemDto) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetRecursiveItemCount

`func (o *BaseItemDto) GetRecursiveItemCount() int32`

GetRecursiveItemCount returns the RecursiveItemCount field if non-nil, zero value otherwise.

### GetRecursiveItemCountOk

`func (o *BaseItemDto) GetRecursiveItemCountOk() (*int32, bool)`

GetRecursiveItemCountOk returns a tuple with the RecursiveItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveItemCount

`func (o *BaseItemDto) SetRecursiveItemCount(v int32)`

SetRecursiveItemCount sets RecursiveItemCount field to given value.

### HasRecursiveItemCount

`func (o *BaseItemDto) HasRecursiveItemCount() bool`

HasRecursiveItemCount returns a boolean if a field has been set.

### SetRecursiveItemCountNil

`func (o *BaseItemDto) SetRecursiveItemCountNil(b bool)`

 SetRecursiveItemCountNil sets the value for RecursiveItemCount to be an explicit nil

### UnsetRecursiveItemCount
`func (o *BaseItemDto) UnsetRecursiveItemCount()`

UnsetRecursiveItemCount ensures that no value is present for RecursiveItemCount, not even an explicit nil
### GetChildCount

`func (o *BaseItemDto) GetChildCount() int32`

GetChildCount returns the ChildCount field if non-nil, zero value otherwise.

### GetChildCountOk

`func (o *BaseItemDto) GetChildCountOk() (*int32, bool)`

GetChildCountOk returns a tuple with the ChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCount

`func (o *BaseItemDto) SetChildCount(v int32)`

SetChildCount sets ChildCount field to given value.

### HasChildCount

`func (o *BaseItemDto) HasChildCount() bool`

HasChildCount returns a boolean if a field has been set.

### SetChildCountNil

`func (o *BaseItemDto) SetChildCountNil(b bool)`

 SetChildCountNil sets the value for ChildCount to be an explicit nil

### UnsetChildCount
`func (o *BaseItemDto) UnsetChildCount()`

UnsetChildCount ensures that no value is present for ChildCount, not even an explicit nil
### GetSeriesName

`func (o *BaseItemDto) GetSeriesName() string`

GetSeriesName returns the SeriesName field if non-nil, zero value otherwise.

### GetSeriesNameOk

`func (o *BaseItemDto) GetSeriesNameOk() (*string, bool)`

GetSeriesNameOk returns a tuple with the SeriesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesName

`func (o *BaseItemDto) SetSeriesName(v string)`

SetSeriesName sets SeriesName field to given value.

### HasSeriesName

`func (o *BaseItemDto) HasSeriesName() bool`

HasSeriesName returns a boolean if a field has been set.

### SetSeriesNameNil

`func (o *BaseItemDto) SetSeriesNameNil(b bool)`

 SetSeriesNameNil sets the value for SeriesName to be an explicit nil

### UnsetSeriesName
`func (o *BaseItemDto) UnsetSeriesName()`

UnsetSeriesName ensures that no value is present for SeriesName, not even an explicit nil
### GetSeriesId

`func (o *BaseItemDto) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *BaseItemDto) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *BaseItemDto) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *BaseItemDto) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### SetSeriesIdNil

`func (o *BaseItemDto) SetSeriesIdNil(b bool)`

 SetSeriesIdNil sets the value for SeriesId to be an explicit nil

### UnsetSeriesId
`func (o *BaseItemDto) UnsetSeriesId()`

UnsetSeriesId ensures that no value is present for SeriesId, not even an explicit nil
### GetSeasonId

`func (o *BaseItemDto) GetSeasonId() string`

GetSeasonId returns the SeasonId field if non-nil, zero value otherwise.

### GetSeasonIdOk

`func (o *BaseItemDto) GetSeasonIdOk() (*string, bool)`

GetSeasonIdOk returns a tuple with the SeasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonId

`func (o *BaseItemDto) SetSeasonId(v string)`

SetSeasonId sets SeasonId field to given value.

### HasSeasonId

`func (o *BaseItemDto) HasSeasonId() bool`

HasSeasonId returns a boolean if a field has been set.

### SetSeasonIdNil

`func (o *BaseItemDto) SetSeasonIdNil(b bool)`

 SetSeasonIdNil sets the value for SeasonId to be an explicit nil

### UnsetSeasonId
`func (o *BaseItemDto) UnsetSeasonId()`

UnsetSeasonId ensures that no value is present for SeasonId, not even an explicit nil
### GetSpecialFeatureCount

`func (o *BaseItemDto) GetSpecialFeatureCount() int32`

GetSpecialFeatureCount returns the SpecialFeatureCount field if non-nil, zero value otherwise.

### GetSpecialFeatureCountOk

`func (o *BaseItemDto) GetSpecialFeatureCountOk() (*int32, bool)`

GetSpecialFeatureCountOk returns a tuple with the SpecialFeatureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialFeatureCount

`func (o *BaseItemDto) SetSpecialFeatureCount(v int32)`

SetSpecialFeatureCount sets SpecialFeatureCount field to given value.

### HasSpecialFeatureCount

`func (o *BaseItemDto) HasSpecialFeatureCount() bool`

HasSpecialFeatureCount returns a boolean if a field has been set.

### SetSpecialFeatureCountNil

`func (o *BaseItemDto) SetSpecialFeatureCountNil(b bool)`

 SetSpecialFeatureCountNil sets the value for SpecialFeatureCount to be an explicit nil

### UnsetSpecialFeatureCount
`func (o *BaseItemDto) UnsetSpecialFeatureCount()`

UnsetSpecialFeatureCount ensures that no value is present for SpecialFeatureCount, not even an explicit nil
### GetDisplayPreferencesId

`func (o *BaseItemDto) GetDisplayPreferencesId() string`

GetDisplayPreferencesId returns the DisplayPreferencesId field if non-nil, zero value otherwise.

### GetDisplayPreferencesIdOk

`func (o *BaseItemDto) GetDisplayPreferencesIdOk() (*string, bool)`

GetDisplayPreferencesIdOk returns a tuple with the DisplayPreferencesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPreferencesId

`func (o *BaseItemDto) SetDisplayPreferencesId(v string)`

SetDisplayPreferencesId sets DisplayPreferencesId field to given value.

### HasDisplayPreferencesId

`func (o *BaseItemDto) HasDisplayPreferencesId() bool`

HasDisplayPreferencesId returns a boolean if a field has been set.

### SetDisplayPreferencesIdNil

`func (o *BaseItemDto) SetDisplayPreferencesIdNil(b bool)`

 SetDisplayPreferencesIdNil sets the value for DisplayPreferencesId to be an explicit nil

### UnsetDisplayPreferencesId
`func (o *BaseItemDto) UnsetDisplayPreferencesId()`

UnsetDisplayPreferencesId ensures that no value is present for DisplayPreferencesId, not even an explicit nil
### GetStatus

`func (o *BaseItemDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseItemDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseItemDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseItemDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BaseItemDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BaseItemDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAirTime

`func (o *BaseItemDto) GetAirTime() string`

GetAirTime returns the AirTime field if non-nil, zero value otherwise.

### GetAirTimeOk

`func (o *BaseItemDto) GetAirTimeOk() (*string, bool)`

GetAirTimeOk returns a tuple with the AirTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirTime

`func (o *BaseItemDto) SetAirTime(v string)`

SetAirTime sets AirTime field to given value.

### HasAirTime

`func (o *BaseItemDto) HasAirTime() bool`

HasAirTime returns a boolean if a field has been set.

### SetAirTimeNil

`func (o *BaseItemDto) SetAirTimeNil(b bool)`

 SetAirTimeNil sets the value for AirTime to be an explicit nil

### UnsetAirTime
`func (o *BaseItemDto) UnsetAirTime()`

UnsetAirTime ensures that no value is present for AirTime, not even an explicit nil
### GetAirDays

`func (o *BaseItemDto) GetAirDays() []DayOfWeek`

GetAirDays returns the AirDays field if non-nil, zero value otherwise.

### GetAirDaysOk

`func (o *BaseItemDto) GetAirDaysOk() (*[]DayOfWeek, bool)`

GetAirDaysOk returns a tuple with the AirDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirDays

`func (o *BaseItemDto) SetAirDays(v []DayOfWeek)`

SetAirDays sets AirDays field to given value.

### HasAirDays

`func (o *BaseItemDto) HasAirDays() bool`

HasAirDays returns a boolean if a field has been set.

### SetAirDaysNil

`func (o *BaseItemDto) SetAirDaysNil(b bool)`

 SetAirDaysNil sets the value for AirDays to be an explicit nil

### UnsetAirDays
`func (o *BaseItemDto) UnsetAirDays()`

UnsetAirDays ensures that no value is present for AirDays, not even an explicit nil
### GetTags

`func (o *BaseItemDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BaseItemDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BaseItemDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BaseItemDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *BaseItemDto) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *BaseItemDto) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetPrimaryImageAspectRatio

`func (o *BaseItemDto) GetPrimaryImageAspectRatio() float64`

GetPrimaryImageAspectRatio returns the PrimaryImageAspectRatio field if non-nil, zero value otherwise.

### GetPrimaryImageAspectRatioOk

`func (o *BaseItemDto) GetPrimaryImageAspectRatioOk() (*float64, bool)`

GetPrimaryImageAspectRatioOk returns a tuple with the PrimaryImageAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageAspectRatio

`func (o *BaseItemDto) SetPrimaryImageAspectRatio(v float64)`

SetPrimaryImageAspectRatio sets PrimaryImageAspectRatio field to given value.

### HasPrimaryImageAspectRatio

`func (o *BaseItemDto) HasPrimaryImageAspectRatio() bool`

HasPrimaryImageAspectRatio returns a boolean if a field has been set.

### SetPrimaryImageAspectRatioNil

`func (o *BaseItemDto) SetPrimaryImageAspectRatioNil(b bool)`

 SetPrimaryImageAspectRatioNil sets the value for PrimaryImageAspectRatio to be an explicit nil

### UnsetPrimaryImageAspectRatio
`func (o *BaseItemDto) UnsetPrimaryImageAspectRatio()`

UnsetPrimaryImageAspectRatio ensures that no value is present for PrimaryImageAspectRatio, not even an explicit nil
### GetArtists

`func (o *BaseItemDto) GetArtists() []string`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *BaseItemDto) GetArtistsOk() (*[]string, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *BaseItemDto) SetArtists(v []string)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *BaseItemDto) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### SetArtistsNil

`func (o *BaseItemDto) SetArtistsNil(b bool)`

 SetArtistsNil sets the value for Artists to be an explicit nil

### UnsetArtists
`func (o *BaseItemDto) UnsetArtists()`

UnsetArtists ensures that no value is present for Artists, not even an explicit nil
### GetArtistItems

`func (o *BaseItemDto) GetArtistItems() []NameGuidPair`

GetArtistItems returns the ArtistItems field if non-nil, zero value otherwise.

### GetArtistItemsOk

`func (o *BaseItemDto) GetArtistItemsOk() (*[]NameGuidPair, bool)`

GetArtistItemsOk returns a tuple with the ArtistItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistItems

`func (o *BaseItemDto) SetArtistItems(v []NameGuidPair)`

SetArtistItems sets ArtistItems field to given value.

### HasArtistItems

`func (o *BaseItemDto) HasArtistItems() bool`

HasArtistItems returns a boolean if a field has been set.

### SetArtistItemsNil

`func (o *BaseItemDto) SetArtistItemsNil(b bool)`

 SetArtistItemsNil sets the value for ArtistItems to be an explicit nil

### UnsetArtistItems
`func (o *BaseItemDto) UnsetArtistItems()`

UnsetArtistItems ensures that no value is present for ArtistItems, not even an explicit nil
### GetAlbum

`func (o *BaseItemDto) GetAlbum() string`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *BaseItemDto) GetAlbumOk() (*string, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *BaseItemDto) SetAlbum(v string)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *BaseItemDto) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### SetAlbumNil

`func (o *BaseItemDto) SetAlbumNil(b bool)`

 SetAlbumNil sets the value for Album to be an explicit nil

### UnsetAlbum
`func (o *BaseItemDto) UnsetAlbum()`

UnsetAlbum ensures that no value is present for Album, not even an explicit nil
### GetCollectionType

`func (o *BaseItemDto) GetCollectionType() string`

GetCollectionType returns the CollectionType field if non-nil, zero value otherwise.

### GetCollectionTypeOk

`func (o *BaseItemDto) GetCollectionTypeOk() (*string, bool)`

GetCollectionTypeOk returns a tuple with the CollectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionType

`func (o *BaseItemDto) SetCollectionType(v string)`

SetCollectionType sets CollectionType field to given value.

### HasCollectionType

`func (o *BaseItemDto) HasCollectionType() bool`

HasCollectionType returns a boolean if a field has been set.

### SetCollectionTypeNil

`func (o *BaseItemDto) SetCollectionTypeNil(b bool)`

 SetCollectionTypeNil sets the value for CollectionType to be an explicit nil

### UnsetCollectionType
`func (o *BaseItemDto) UnsetCollectionType()`

UnsetCollectionType ensures that no value is present for CollectionType, not even an explicit nil
### GetDisplayOrder

`func (o *BaseItemDto) GetDisplayOrder() string`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *BaseItemDto) GetDisplayOrderOk() (*string, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *BaseItemDto) SetDisplayOrder(v string)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *BaseItemDto) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### SetDisplayOrderNil

`func (o *BaseItemDto) SetDisplayOrderNil(b bool)`

 SetDisplayOrderNil sets the value for DisplayOrder to be an explicit nil

### UnsetDisplayOrder
`func (o *BaseItemDto) UnsetDisplayOrder()`

UnsetDisplayOrder ensures that no value is present for DisplayOrder, not even an explicit nil
### GetAlbumId

`func (o *BaseItemDto) GetAlbumId() string`

GetAlbumId returns the AlbumId field if non-nil, zero value otherwise.

### GetAlbumIdOk

`func (o *BaseItemDto) GetAlbumIdOk() (*string, bool)`

GetAlbumIdOk returns a tuple with the AlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumId

`func (o *BaseItemDto) SetAlbumId(v string)`

SetAlbumId sets AlbumId field to given value.

### HasAlbumId

`func (o *BaseItemDto) HasAlbumId() bool`

HasAlbumId returns a boolean if a field has been set.

### SetAlbumIdNil

`func (o *BaseItemDto) SetAlbumIdNil(b bool)`

 SetAlbumIdNil sets the value for AlbumId to be an explicit nil

### UnsetAlbumId
`func (o *BaseItemDto) UnsetAlbumId()`

UnsetAlbumId ensures that no value is present for AlbumId, not even an explicit nil
### GetAlbumPrimaryImageTag

`func (o *BaseItemDto) GetAlbumPrimaryImageTag() string`

GetAlbumPrimaryImageTag returns the AlbumPrimaryImageTag field if non-nil, zero value otherwise.

### GetAlbumPrimaryImageTagOk

`func (o *BaseItemDto) GetAlbumPrimaryImageTagOk() (*string, bool)`

GetAlbumPrimaryImageTagOk returns a tuple with the AlbumPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumPrimaryImageTag

`func (o *BaseItemDto) SetAlbumPrimaryImageTag(v string)`

SetAlbumPrimaryImageTag sets AlbumPrimaryImageTag field to given value.

### HasAlbumPrimaryImageTag

`func (o *BaseItemDto) HasAlbumPrimaryImageTag() bool`

HasAlbumPrimaryImageTag returns a boolean if a field has been set.

### SetAlbumPrimaryImageTagNil

`func (o *BaseItemDto) SetAlbumPrimaryImageTagNil(b bool)`

 SetAlbumPrimaryImageTagNil sets the value for AlbumPrimaryImageTag to be an explicit nil

### UnsetAlbumPrimaryImageTag
`func (o *BaseItemDto) UnsetAlbumPrimaryImageTag()`

UnsetAlbumPrimaryImageTag ensures that no value is present for AlbumPrimaryImageTag, not even an explicit nil
### GetSeriesPrimaryImageTag

`func (o *BaseItemDto) GetSeriesPrimaryImageTag() string`

GetSeriesPrimaryImageTag returns the SeriesPrimaryImageTag field if non-nil, zero value otherwise.

### GetSeriesPrimaryImageTagOk

`func (o *BaseItemDto) GetSeriesPrimaryImageTagOk() (*string, bool)`

GetSeriesPrimaryImageTagOk returns a tuple with the SeriesPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesPrimaryImageTag

`func (o *BaseItemDto) SetSeriesPrimaryImageTag(v string)`

SetSeriesPrimaryImageTag sets SeriesPrimaryImageTag field to given value.

### HasSeriesPrimaryImageTag

`func (o *BaseItemDto) HasSeriesPrimaryImageTag() bool`

HasSeriesPrimaryImageTag returns a boolean if a field has been set.

### SetSeriesPrimaryImageTagNil

`func (o *BaseItemDto) SetSeriesPrimaryImageTagNil(b bool)`

 SetSeriesPrimaryImageTagNil sets the value for SeriesPrimaryImageTag to be an explicit nil

### UnsetSeriesPrimaryImageTag
`func (o *BaseItemDto) UnsetSeriesPrimaryImageTag()`

UnsetSeriesPrimaryImageTag ensures that no value is present for SeriesPrimaryImageTag, not even an explicit nil
### GetAlbumArtist

`func (o *BaseItemDto) GetAlbumArtist() string`

GetAlbumArtist returns the AlbumArtist field if non-nil, zero value otherwise.

### GetAlbumArtistOk

`func (o *BaseItemDto) GetAlbumArtistOk() (*string, bool)`

GetAlbumArtistOk returns a tuple with the AlbumArtist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtist

`func (o *BaseItemDto) SetAlbumArtist(v string)`

SetAlbumArtist sets AlbumArtist field to given value.

### HasAlbumArtist

`func (o *BaseItemDto) HasAlbumArtist() bool`

HasAlbumArtist returns a boolean if a field has been set.

### SetAlbumArtistNil

`func (o *BaseItemDto) SetAlbumArtistNil(b bool)`

 SetAlbumArtistNil sets the value for AlbumArtist to be an explicit nil

### UnsetAlbumArtist
`func (o *BaseItemDto) UnsetAlbumArtist()`

UnsetAlbumArtist ensures that no value is present for AlbumArtist, not even an explicit nil
### GetAlbumArtists

`func (o *BaseItemDto) GetAlbumArtists() []NameGuidPair`

GetAlbumArtists returns the AlbumArtists field if non-nil, zero value otherwise.

### GetAlbumArtistsOk

`func (o *BaseItemDto) GetAlbumArtistsOk() (*[]NameGuidPair, bool)`

GetAlbumArtistsOk returns a tuple with the AlbumArtists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtists

`func (o *BaseItemDto) SetAlbumArtists(v []NameGuidPair)`

SetAlbumArtists sets AlbumArtists field to given value.

### HasAlbumArtists

`func (o *BaseItemDto) HasAlbumArtists() bool`

HasAlbumArtists returns a boolean if a field has been set.

### SetAlbumArtistsNil

`func (o *BaseItemDto) SetAlbumArtistsNil(b bool)`

 SetAlbumArtistsNil sets the value for AlbumArtists to be an explicit nil

### UnsetAlbumArtists
`func (o *BaseItemDto) UnsetAlbumArtists()`

UnsetAlbumArtists ensures that no value is present for AlbumArtists, not even an explicit nil
### GetSeasonName

`func (o *BaseItemDto) GetSeasonName() string`

GetSeasonName returns the SeasonName field if non-nil, zero value otherwise.

### GetSeasonNameOk

`func (o *BaseItemDto) GetSeasonNameOk() (*string, bool)`

GetSeasonNameOk returns a tuple with the SeasonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonName

`func (o *BaseItemDto) SetSeasonName(v string)`

SetSeasonName sets SeasonName field to given value.

### HasSeasonName

`func (o *BaseItemDto) HasSeasonName() bool`

HasSeasonName returns a boolean if a field has been set.

### SetSeasonNameNil

`func (o *BaseItemDto) SetSeasonNameNil(b bool)`

 SetSeasonNameNil sets the value for SeasonName to be an explicit nil

### UnsetSeasonName
`func (o *BaseItemDto) UnsetSeasonName()`

UnsetSeasonName ensures that no value is present for SeasonName, not even an explicit nil
### GetMediaStreams

`func (o *BaseItemDto) GetMediaStreams() []MediaStream`

GetMediaStreams returns the MediaStreams field if non-nil, zero value otherwise.

### GetMediaStreamsOk

`func (o *BaseItemDto) GetMediaStreamsOk() (*[]MediaStream, bool)`

GetMediaStreamsOk returns a tuple with the MediaStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaStreams

`func (o *BaseItemDto) SetMediaStreams(v []MediaStream)`

SetMediaStreams sets MediaStreams field to given value.

### HasMediaStreams

`func (o *BaseItemDto) HasMediaStreams() bool`

HasMediaStreams returns a boolean if a field has been set.

### SetMediaStreamsNil

`func (o *BaseItemDto) SetMediaStreamsNil(b bool)`

 SetMediaStreamsNil sets the value for MediaStreams to be an explicit nil

### UnsetMediaStreams
`func (o *BaseItemDto) UnsetMediaStreams()`

UnsetMediaStreams ensures that no value is present for MediaStreams, not even an explicit nil
### GetVideoType

`func (o *BaseItemDto) GetVideoType() VideoType`

GetVideoType returns the VideoType field if non-nil, zero value otherwise.

### GetVideoTypeOk

`func (o *BaseItemDto) GetVideoTypeOk() (*VideoType, bool)`

GetVideoTypeOk returns a tuple with the VideoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoType

`func (o *BaseItemDto) SetVideoType(v VideoType)`

SetVideoType sets VideoType field to given value.

### HasVideoType

`func (o *BaseItemDto) HasVideoType() bool`

HasVideoType returns a boolean if a field has been set.

### SetVideoTypeNil

`func (o *BaseItemDto) SetVideoTypeNil(b bool)`

 SetVideoTypeNil sets the value for VideoType to be an explicit nil

### UnsetVideoType
`func (o *BaseItemDto) UnsetVideoType()`

UnsetVideoType ensures that no value is present for VideoType, not even an explicit nil
### GetPartCount

`func (o *BaseItemDto) GetPartCount() int32`

GetPartCount returns the PartCount field if non-nil, zero value otherwise.

### GetPartCountOk

`func (o *BaseItemDto) GetPartCountOk() (*int32, bool)`

GetPartCountOk returns a tuple with the PartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartCount

`func (o *BaseItemDto) SetPartCount(v int32)`

SetPartCount sets PartCount field to given value.

### HasPartCount

`func (o *BaseItemDto) HasPartCount() bool`

HasPartCount returns a boolean if a field has been set.

### SetPartCountNil

`func (o *BaseItemDto) SetPartCountNil(b bool)`

 SetPartCountNil sets the value for PartCount to be an explicit nil

### UnsetPartCount
`func (o *BaseItemDto) UnsetPartCount()`

UnsetPartCount ensures that no value is present for PartCount, not even an explicit nil
### GetMediaSourceCount

`func (o *BaseItemDto) GetMediaSourceCount() int32`

GetMediaSourceCount returns the MediaSourceCount field if non-nil, zero value otherwise.

### GetMediaSourceCountOk

`func (o *BaseItemDto) GetMediaSourceCountOk() (*int32, bool)`

GetMediaSourceCountOk returns a tuple with the MediaSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceCount

`func (o *BaseItemDto) SetMediaSourceCount(v int32)`

SetMediaSourceCount sets MediaSourceCount field to given value.

### HasMediaSourceCount

`func (o *BaseItemDto) HasMediaSourceCount() bool`

HasMediaSourceCount returns a boolean if a field has been set.

### SetMediaSourceCountNil

`func (o *BaseItemDto) SetMediaSourceCountNil(b bool)`

 SetMediaSourceCountNil sets the value for MediaSourceCount to be an explicit nil

### UnsetMediaSourceCount
`func (o *BaseItemDto) UnsetMediaSourceCount()`

UnsetMediaSourceCount ensures that no value is present for MediaSourceCount, not even an explicit nil
### GetImageTags

`func (o *BaseItemDto) GetImageTags() map[string]string`

GetImageTags returns the ImageTags field if non-nil, zero value otherwise.

### GetImageTagsOk

`func (o *BaseItemDto) GetImageTagsOk() (*map[string]string, bool)`

GetImageTagsOk returns a tuple with the ImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTags

`func (o *BaseItemDto) SetImageTags(v map[string]string)`

SetImageTags sets ImageTags field to given value.

### HasImageTags

`func (o *BaseItemDto) HasImageTags() bool`

HasImageTags returns a boolean if a field has been set.

### SetImageTagsNil

`func (o *BaseItemDto) SetImageTagsNil(b bool)`

 SetImageTagsNil sets the value for ImageTags to be an explicit nil

### UnsetImageTags
`func (o *BaseItemDto) UnsetImageTags()`

UnsetImageTags ensures that no value is present for ImageTags, not even an explicit nil
### GetBackdropImageTags

`func (o *BaseItemDto) GetBackdropImageTags() []string`

GetBackdropImageTags returns the BackdropImageTags field if non-nil, zero value otherwise.

### GetBackdropImageTagsOk

`func (o *BaseItemDto) GetBackdropImageTagsOk() (*[]string, bool)`

GetBackdropImageTagsOk returns a tuple with the BackdropImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdropImageTags

`func (o *BaseItemDto) SetBackdropImageTags(v []string)`

SetBackdropImageTags sets BackdropImageTags field to given value.

### HasBackdropImageTags

`func (o *BaseItemDto) HasBackdropImageTags() bool`

HasBackdropImageTags returns a boolean if a field has been set.

### SetBackdropImageTagsNil

`func (o *BaseItemDto) SetBackdropImageTagsNil(b bool)`

 SetBackdropImageTagsNil sets the value for BackdropImageTags to be an explicit nil

### UnsetBackdropImageTags
`func (o *BaseItemDto) UnsetBackdropImageTags()`

UnsetBackdropImageTags ensures that no value is present for BackdropImageTags, not even an explicit nil
### GetScreenshotImageTags

`func (o *BaseItemDto) GetScreenshotImageTags() []string`

GetScreenshotImageTags returns the ScreenshotImageTags field if non-nil, zero value otherwise.

### GetScreenshotImageTagsOk

`func (o *BaseItemDto) GetScreenshotImageTagsOk() (*[]string, bool)`

GetScreenshotImageTagsOk returns a tuple with the ScreenshotImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotImageTags

`func (o *BaseItemDto) SetScreenshotImageTags(v []string)`

SetScreenshotImageTags sets ScreenshotImageTags field to given value.

### HasScreenshotImageTags

`func (o *BaseItemDto) HasScreenshotImageTags() bool`

HasScreenshotImageTags returns a boolean if a field has been set.

### SetScreenshotImageTagsNil

`func (o *BaseItemDto) SetScreenshotImageTagsNil(b bool)`

 SetScreenshotImageTagsNil sets the value for ScreenshotImageTags to be an explicit nil

### UnsetScreenshotImageTags
`func (o *BaseItemDto) UnsetScreenshotImageTags()`

UnsetScreenshotImageTags ensures that no value is present for ScreenshotImageTags, not even an explicit nil
### GetParentLogoImageTag

`func (o *BaseItemDto) GetParentLogoImageTag() string`

GetParentLogoImageTag returns the ParentLogoImageTag field if non-nil, zero value otherwise.

### GetParentLogoImageTagOk

`func (o *BaseItemDto) GetParentLogoImageTagOk() (*string, bool)`

GetParentLogoImageTagOk returns a tuple with the ParentLogoImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLogoImageTag

`func (o *BaseItemDto) SetParentLogoImageTag(v string)`

SetParentLogoImageTag sets ParentLogoImageTag field to given value.

### HasParentLogoImageTag

`func (o *BaseItemDto) HasParentLogoImageTag() bool`

HasParentLogoImageTag returns a boolean if a field has been set.

### SetParentLogoImageTagNil

`func (o *BaseItemDto) SetParentLogoImageTagNil(b bool)`

 SetParentLogoImageTagNil sets the value for ParentLogoImageTag to be an explicit nil

### UnsetParentLogoImageTag
`func (o *BaseItemDto) UnsetParentLogoImageTag()`

UnsetParentLogoImageTag ensures that no value is present for ParentLogoImageTag, not even an explicit nil
### GetParentArtItemId

`func (o *BaseItemDto) GetParentArtItemId() string`

GetParentArtItemId returns the ParentArtItemId field if non-nil, zero value otherwise.

### GetParentArtItemIdOk

`func (o *BaseItemDto) GetParentArtItemIdOk() (*string, bool)`

GetParentArtItemIdOk returns a tuple with the ParentArtItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentArtItemId

`func (o *BaseItemDto) SetParentArtItemId(v string)`

SetParentArtItemId sets ParentArtItemId field to given value.

### HasParentArtItemId

`func (o *BaseItemDto) HasParentArtItemId() bool`

HasParentArtItemId returns a boolean if a field has been set.

### SetParentArtItemIdNil

`func (o *BaseItemDto) SetParentArtItemIdNil(b bool)`

 SetParentArtItemIdNil sets the value for ParentArtItemId to be an explicit nil

### UnsetParentArtItemId
`func (o *BaseItemDto) UnsetParentArtItemId()`

UnsetParentArtItemId ensures that no value is present for ParentArtItemId, not even an explicit nil
### GetParentArtImageTag

`func (o *BaseItemDto) GetParentArtImageTag() string`

GetParentArtImageTag returns the ParentArtImageTag field if non-nil, zero value otherwise.

### GetParentArtImageTagOk

`func (o *BaseItemDto) GetParentArtImageTagOk() (*string, bool)`

GetParentArtImageTagOk returns a tuple with the ParentArtImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentArtImageTag

`func (o *BaseItemDto) SetParentArtImageTag(v string)`

SetParentArtImageTag sets ParentArtImageTag field to given value.

### HasParentArtImageTag

`func (o *BaseItemDto) HasParentArtImageTag() bool`

HasParentArtImageTag returns a boolean if a field has been set.

### SetParentArtImageTagNil

`func (o *BaseItemDto) SetParentArtImageTagNil(b bool)`

 SetParentArtImageTagNil sets the value for ParentArtImageTag to be an explicit nil

### UnsetParentArtImageTag
`func (o *BaseItemDto) UnsetParentArtImageTag()`

UnsetParentArtImageTag ensures that no value is present for ParentArtImageTag, not even an explicit nil
### GetSeriesThumbImageTag

`func (o *BaseItemDto) GetSeriesThumbImageTag() string`

GetSeriesThumbImageTag returns the SeriesThumbImageTag field if non-nil, zero value otherwise.

### GetSeriesThumbImageTagOk

`func (o *BaseItemDto) GetSeriesThumbImageTagOk() (*string, bool)`

GetSeriesThumbImageTagOk returns a tuple with the SeriesThumbImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesThumbImageTag

`func (o *BaseItemDto) SetSeriesThumbImageTag(v string)`

SetSeriesThumbImageTag sets SeriesThumbImageTag field to given value.

### HasSeriesThumbImageTag

`func (o *BaseItemDto) HasSeriesThumbImageTag() bool`

HasSeriesThumbImageTag returns a boolean if a field has been set.

### SetSeriesThumbImageTagNil

`func (o *BaseItemDto) SetSeriesThumbImageTagNil(b bool)`

 SetSeriesThumbImageTagNil sets the value for SeriesThumbImageTag to be an explicit nil

### UnsetSeriesThumbImageTag
`func (o *BaseItemDto) UnsetSeriesThumbImageTag()`

UnsetSeriesThumbImageTag ensures that no value is present for SeriesThumbImageTag, not even an explicit nil
### GetImageBlurHashes

`func (o *BaseItemDto) GetImageBlurHashes() BaseItemDtoImageBlurHashes`

GetImageBlurHashes returns the ImageBlurHashes field if non-nil, zero value otherwise.

### GetImageBlurHashesOk

`func (o *BaseItemDto) GetImageBlurHashesOk() (*BaseItemDtoImageBlurHashes, bool)`

GetImageBlurHashesOk returns a tuple with the ImageBlurHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBlurHashes

`func (o *BaseItemDto) SetImageBlurHashes(v BaseItemDtoImageBlurHashes)`

SetImageBlurHashes sets ImageBlurHashes field to given value.

### HasImageBlurHashes

`func (o *BaseItemDto) HasImageBlurHashes() bool`

HasImageBlurHashes returns a boolean if a field has been set.

### SetImageBlurHashesNil

`func (o *BaseItemDto) SetImageBlurHashesNil(b bool)`

 SetImageBlurHashesNil sets the value for ImageBlurHashes to be an explicit nil

### UnsetImageBlurHashes
`func (o *BaseItemDto) UnsetImageBlurHashes()`

UnsetImageBlurHashes ensures that no value is present for ImageBlurHashes, not even an explicit nil
### GetSeriesStudio

`func (o *BaseItemDto) GetSeriesStudio() string`

GetSeriesStudio returns the SeriesStudio field if non-nil, zero value otherwise.

### GetSeriesStudioOk

`func (o *BaseItemDto) GetSeriesStudioOk() (*string, bool)`

GetSeriesStudioOk returns a tuple with the SeriesStudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesStudio

`func (o *BaseItemDto) SetSeriesStudio(v string)`

SetSeriesStudio sets SeriesStudio field to given value.

### HasSeriesStudio

`func (o *BaseItemDto) HasSeriesStudio() bool`

HasSeriesStudio returns a boolean if a field has been set.

### SetSeriesStudioNil

`func (o *BaseItemDto) SetSeriesStudioNil(b bool)`

 SetSeriesStudioNil sets the value for SeriesStudio to be an explicit nil

### UnsetSeriesStudio
`func (o *BaseItemDto) UnsetSeriesStudio()`

UnsetSeriesStudio ensures that no value is present for SeriesStudio, not even an explicit nil
### GetParentThumbItemId

`func (o *BaseItemDto) GetParentThumbItemId() string`

GetParentThumbItemId returns the ParentThumbItemId field if non-nil, zero value otherwise.

### GetParentThumbItemIdOk

`func (o *BaseItemDto) GetParentThumbItemIdOk() (*string, bool)`

GetParentThumbItemIdOk returns a tuple with the ParentThumbItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentThumbItemId

`func (o *BaseItemDto) SetParentThumbItemId(v string)`

SetParentThumbItemId sets ParentThumbItemId field to given value.

### HasParentThumbItemId

`func (o *BaseItemDto) HasParentThumbItemId() bool`

HasParentThumbItemId returns a boolean if a field has been set.

### SetParentThumbItemIdNil

`func (o *BaseItemDto) SetParentThumbItemIdNil(b bool)`

 SetParentThumbItemIdNil sets the value for ParentThumbItemId to be an explicit nil

### UnsetParentThumbItemId
`func (o *BaseItemDto) UnsetParentThumbItemId()`

UnsetParentThumbItemId ensures that no value is present for ParentThumbItemId, not even an explicit nil
### GetParentThumbImageTag

`func (o *BaseItemDto) GetParentThumbImageTag() string`

GetParentThumbImageTag returns the ParentThumbImageTag field if non-nil, zero value otherwise.

### GetParentThumbImageTagOk

`func (o *BaseItemDto) GetParentThumbImageTagOk() (*string, bool)`

GetParentThumbImageTagOk returns a tuple with the ParentThumbImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentThumbImageTag

`func (o *BaseItemDto) SetParentThumbImageTag(v string)`

SetParentThumbImageTag sets ParentThumbImageTag field to given value.

### HasParentThumbImageTag

`func (o *BaseItemDto) HasParentThumbImageTag() bool`

HasParentThumbImageTag returns a boolean if a field has been set.

### SetParentThumbImageTagNil

`func (o *BaseItemDto) SetParentThumbImageTagNil(b bool)`

 SetParentThumbImageTagNil sets the value for ParentThumbImageTag to be an explicit nil

### UnsetParentThumbImageTag
`func (o *BaseItemDto) UnsetParentThumbImageTag()`

UnsetParentThumbImageTag ensures that no value is present for ParentThumbImageTag, not even an explicit nil
### GetParentPrimaryImageItemId

`func (o *BaseItemDto) GetParentPrimaryImageItemId() string`

GetParentPrimaryImageItemId returns the ParentPrimaryImageItemId field if non-nil, zero value otherwise.

### GetParentPrimaryImageItemIdOk

`func (o *BaseItemDto) GetParentPrimaryImageItemIdOk() (*string, bool)`

GetParentPrimaryImageItemIdOk returns a tuple with the ParentPrimaryImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPrimaryImageItemId

`func (o *BaseItemDto) SetParentPrimaryImageItemId(v string)`

SetParentPrimaryImageItemId sets ParentPrimaryImageItemId field to given value.

### HasParentPrimaryImageItemId

`func (o *BaseItemDto) HasParentPrimaryImageItemId() bool`

HasParentPrimaryImageItemId returns a boolean if a field has been set.

### SetParentPrimaryImageItemIdNil

`func (o *BaseItemDto) SetParentPrimaryImageItemIdNil(b bool)`

 SetParentPrimaryImageItemIdNil sets the value for ParentPrimaryImageItemId to be an explicit nil

### UnsetParentPrimaryImageItemId
`func (o *BaseItemDto) UnsetParentPrimaryImageItemId()`

UnsetParentPrimaryImageItemId ensures that no value is present for ParentPrimaryImageItemId, not even an explicit nil
### GetParentPrimaryImageTag

`func (o *BaseItemDto) GetParentPrimaryImageTag() string`

GetParentPrimaryImageTag returns the ParentPrimaryImageTag field if non-nil, zero value otherwise.

### GetParentPrimaryImageTagOk

`func (o *BaseItemDto) GetParentPrimaryImageTagOk() (*string, bool)`

GetParentPrimaryImageTagOk returns a tuple with the ParentPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPrimaryImageTag

`func (o *BaseItemDto) SetParentPrimaryImageTag(v string)`

SetParentPrimaryImageTag sets ParentPrimaryImageTag field to given value.

### HasParentPrimaryImageTag

`func (o *BaseItemDto) HasParentPrimaryImageTag() bool`

HasParentPrimaryImageTag returns a boolean if a field has been set.

### SetParentPrimaryImageTagNil

`func (o *BaseItemDto) SetParentPrimaryImageTagNil(b bool)`

 SetParentPrimaryImageTagNil sets the value for ParentPrimaryImageTag to be an explicit nil

### UnsetParentPrimaryImageTag
`func (o *BaseItemDto) UnsetParentPrimaryImageTag()`

UnsetParentPrimaryImageTag ensures that no value is present for ParentPrimaryImageTag, not even an explicit nil
### GetChapters

`func (o *BaseItemDto) GetChapters() []ChapterInfo`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *BaseItemDto) GetChaptersOk() (*[]ChapterInfo, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *BaseItemDto) SetChapters(v []ChapterInfo)`

SetChapters sets Chapters field to given value.

### HasChapters

`func (o *BaseItemDto) HasChapters() bool`

HasChapters returns a boolean if a field has been set.

### SetChaptersNil

`func (o *BaseItemDto) SetChaptersNil(b bool)`

 SetChaptersNil sets the value for Chapters to be an explicit nil

### UnsetChapters
`func (o *BaseItemDto) UnsetChapters()`

UnsetChapters ensures that no value is present for Chapters, not even an explicit nil
### GetLocationType

`func (o *BaseItemDto) GetLocationType() LocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *BaseItemDto) GetLocationTypeOk() (*LocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *BaseItemDto) SetLocationType(v LocationType)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *BaseItemDto) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### SetLocationTypeNil

`func (o *BaseItemDto) SetLocationTypeNil(b bool)`

 SetLocationTypeNil sets the value for LocationType to be an explicit nil

### UnsetLocationType
`func (o *BaseItemDto) UnsetLocationType()`

UnsetLocationType ensures that no value is present for LocationType, not even an explicit nil
### GetIsoType

`func (o *BaseItemDto) GetIsoType() IsoType`

GetIsoType returns the IsoType field if non-nil, zero value otherwise.

### GetIsoTypeOk

`func (o *BaseItemDto) GetIsoTypeOk() (*IsoType, bool)`

GetIsoTypeOk returns a tuple with the IsoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoType

`func (o *BaseItemDto) SetIsoType(v IsoType)`

SetIsoType sets IsoType field to given value.

### HasIsoType

`func (o *BaseItemDto) HasIsoType() bool`

HasIsoType returns a boolean if a field has been set.

### SetIsoTypeNil

`func (o *BaseItemDto) SetIsoTypeNil(b bool)`

 SetIsoTypeNil sets the value for IsoType to be an explicit nil

### UnsetIsoType
`func (o *BaseItemDto) UnsetIsoType()`

UnsetIsoType ensures that no value is present for IsoType, not even an explicit nil
### GetMediaType

`func (o *BaseItemDto) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *BaseItemDto) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *BaseItemDto) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *BaseItemDto) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### SetMediaTypeNil

`func (o *BaseItemDto) SetMediaTypeNil(b bool)`

 SetMediaTypeNil sets the value for MediaType to be an explicit nil

### UnsetMediaType
`func (o *BaseItemDto) UnsetMediaType()`

UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
### GetEndDate

`func (o *BaseItemDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BaseItemDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BaseItemDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BaseItemDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *BaseItemDto) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *BaseItemDto) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetLockedFields

`func (o *BaseItemDto) GetLockedFields() []MetadataField`

GetLockedFields returns the LockedFields field if non-nil, zero value otherwise.

### GetLockedFieldsOk

`func (o *BaseItemDto) GetLockedFieldsOk() (*[]MetadataField, bool)`

GetLockedFieldsOk returns a tuple with the LockedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedFields

`func (o *BaseItemDto) SetLockedFields(v []MetadataField)`

SetLockedFields sets LockedFields field to given value.

### HasLockedFields

`func (o *BaseItemDto) HasLockedFields() bool`

HasLockedFields returns a boolean if a field has been set.

### SetLockedFieldsNil

`func (o *BaseItemDto) SetLockedFieldsNil(b bool)`

 SetLockedFieldsNil sets the value for LockedFields to be an explicit nil

### UnsetLockedFields
`func (o *BaseItemDto) UnsetLockedFields()`

UnsetLockedFields ensures that no value is present for LockedFields, not even an explicit nil
### GetTrailerCount

`func (o *BaseItemDto) GetTrailerCount() int32`

GetTrailerCount returns the TrailerCount field if non-nil, zero value otherwise.

### GetTrailerCountOk

`func (o *BaseItemDto) GetTrailerCountOk() (*int32, bool)`

GetTrailerCountOk returns a tuple with the TrailerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerCount

`func (o *BaseItemDto) SetTrailerCount(v int32)`

SetTrailerCount sets TrailerCount field to given value.

### HasTrailerCount

`func (o *BaseItemDto) HasTrailerCount() bool`

HasTrailerCount returns a boolean if a field has been set.

### SetTrailerCountNil

`func (o *BaseItemDto) SetTrailerCountNil(b bool)`

 SetTrailerCountNil sets the value for TrailerCount to be an explicit nil

### UnsetTrailerCount
`func (o *BaseItemDto) UnsetTrailerCount()`

UnsetTrailerCount ensures that no value is present for TrailerCount, not even an explicit nil
### GetMovieCount

`func (o *BaseItemDto) GetMovieCount() int32`

GetMovieCount returns the MovieCount field if non-nil, zero value otherwise.

### GetMovieCountOk

`func (o *BaseItemDto) GetMovieCountOk() (*int32, bool)`

GetMovieCountOk returns a tuple with the MovieCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieCount

`func (o *BaseItemDto) SetMovieCount(v int32)`

SetMovieCount sets MovieCount field to given value.

### HasMovieCount

`func (o *BaseItemDto) HasMovieCount() bool`

HasMovieCount returns a boolean if a field has been set.

### SetMovieCountNil

`func (o *BaseItemDto) SetMovieCountNil(b bool)`

 SetMovieCountNil sets the value for MovieCount to be an explicit nil

### UnsetMovieCount
`func (o *BaseItemDto) UnsetMovieCount()`

UnsetMovieCount ensures that no value is present for MovieCount, not even an explicit nil
### GetSeriesCount

`func (o *BaseItemDto) GetSeriesCount() int32`

GetSeriesCount returns the SeriesCount field if non-nil, zero value otherwise.

### GetSeriesCountOk

`func (o *BaseItemDto) GetSeriesCountOk() (*int32, bool)`

GetSeriesCountOk returns a tuple with the SeriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesCount

`func (o *BaseItemDto) SetSeriesCount(v int32)`

SetSeriesCount sets SeriesCount field to given value.

### HasSeriesCount

`func (o *BaseItemDto) HasSeriesCount() bool`

HasSeriesCount returns a boolean if a field has been set.

### SetSeriesCountNil

`func (o *BaseItemDto) SetSeriesCountNil(b bool)`

 SetSeriesCountNil sets the value for SeriesCount to be an explicit nil

### UnsetSeriesCount
`func (o *BaseItemDto) UnsetSeriesCount()`

UnsetSeriesCount ensures that no value is present for SeriesCount, not even an explicit nil
### GetProgramCount

`func (o *BaseItemDto) GetProgramCount() int32`

GetProgramCount returns the ProgramCount field if non-nil, zero value otherwise.

### GetProgramCountOk

`func (o *BaseItemDto) GetProgramCountOk() (*int32, bool)`

GetProgramCountOk returns a tuple with the ProgramCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramCount

`func (o *BaseItemDto) SetProgramCount(v int32)`

SetProgramCount sets ProgramCount field to given value.

### HasProgramCount

`func (o *BaseItemDto) HasProgramCount() bool`

HasProgramCount returns a boolean if a field has been set.

### SetProgramCountNil

`func (o *BaseItemDto) SetProgramCountNil(b bool)`

 SetProgramCountNil sets the value for ProgramCount to be an explicit nil

### UnsetProgramCount
`func (o *BaseItemDto) UnsetProgramCount()`

UnsetProgramCount ensures that no value is present for ProgramCount, not even an explicit nil
### GetEpisodeCount

`func (o *BaseItemDto) GetEpisodeCount() int32`

GetEpisodeCount returns the EpisodeCount field if non-nil, zero value otherwise.

### GetEpisodeCountOk

`func (o *BaseItemDto) GetEpisodeCountOk() (*int32, bool)`

GetEpisodeCountOk returns a tuple with the EpisodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeCount

`func (o *BaseItemDto) SetEpisodeCount(v int32)`

SetEpisodeCount sets EpisodeCount field to given value.

### HasEpisodeCount

`func (o *BaseItemDto) HasEpisodeCount() bool`

HasEpisodeCount returns a boolean if a field has been set.

### SetEpisodeCountNil

`func (o *BaseItemDto) SetEpisodeCountNil(b bool)`

 SetEpisodeCountNil sets the value for EpisodeCount to be an explicit nil

### UnsetEpisodeCount
`func (o *BaseItemDto) UnsetEpisodeCount()`

UnsetEpisodeCount ensures that no value is present for EpisodeCount, not even an explicit nil
### GetSongCount

`func (o *BaseItemDto) GetSongCount() int32`

GetSongCount returns the SongCount field if non-nil, zero value otherwise.

### GetSongCountOk

`func (o *BaseItemDto) GetSongCountOk() (*int32, bool)`

GetSongCountOk returns a tuple with the SongCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSongCount

`func (o *BaseItemDto) SetSongCount(v int32)`

SetSongCount sets SongCount field to given value.

### HasSongCount

`func (o *BaseItemDto) HasSongCount() bool`

HasSongCount returns a boolean if a field has been set.

### SetSongCountNil

`func (o *BaseItemDto) SetSongCountNil(b bool)`

 SetSongCountNil sets the value for SongCount to be an explicit nil

### UnsetSongCount
`func (o *BaseItemDto) UnsetSongCount()`

UnsetSongCount ensures that no value is present for SongCount, not even an explicit nil
### GetAlbumCount

`func (o *BaseItemDto) GetAlbumCount() int32`

GetAlbumCount returns the AlbumCount field if non-nil, zero value otherwise.

### GetAlbumCountOk

`func (o *BaseItemDto) GetAlbumCountOk() (*int32, bool)`

GetAlbumCountOk returns a tuple with the AlbumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumCount

`func (o *BaseItemDto) SetAlbumCount(v int32)`

SetAlbumCount sets AlbumCount field to given value.

### HasAlbumCount

`func (o *BaseItemDto) HasAlbumCount() bool`

HasAlbumCount returns a boolean if a field has been set.

### SetAlbumCountNil

`func (o *BaseItemDto) SetAlbumCountNil(b bool)`

 SetAlbumCountNil sets the value for AlbumCount to be an explicit nil

### UnsetAlbumCount
`func (o *BaseItemDto) UnsetAlbumCount()`

UnsetAlbumCount ensures that no value is present for AlbumCount, not even an explicit nil
### GetArtistCount

`func (o *BaseItemDto) GetArtistCount() int32`

GetArtistCount returns the ArtistCount field if non-nil, zero value otherwise.

### GetArtistCountOk

`func (o *BaseItemDto) GetArtistCountOk() (*int32, bool)`

GetArtistCountOk returns a tuple with the ArtistCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistCount

`func (o *BaseItemDto) SetArtistCount(v int32)`

SetArtistCount sets ArtistCount field to given value.

### HasArtistCount

`func (o *BaseItemDto) HasArtistCount() bool`

HasArtistCount returns a boolean if a field has been set.

### SetArtistCountNil

`func (o *BaseItemDto) SetArtistCountNil(b bool)`

 SetArtistCountNil sets the value for ArtistCount to be an explicit nil

### UnsetArtistCount
`func (o *BaseItemDto) UnsetArtistCount()`

UnsetArtistCount ensures that no value is present for ArtistCount, not even an explicit nil
### GetMusicVideoCount

`func (o *BaseItemDto) GetMusicVideoCount() int32`

GetMusicVideoCount returns the MusicVideoCount field if non-nil, zero value otherwise.

### GetMusicVideoCountOk

`func (o *BaseItemDto) GetMusicVideoCountOk() (*int32, bool)`

GetMusicVideoCountOk returns a tuple with the MusicVideoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicVideoCount

`func (o *BaseItemDto) SetMusicVideoCount(v int32)`

SetMusicVideoCount sets MusicVideoCount field to given value.

### HasMusicVideoCount

`func (o *BaseItemDto) HasMusicVideoCount() bool`

HasMusicVideoCount returns a boolean if a field has been set.

### SetMusicVideoCountNil

`func (o *BaseItemDto) SetMusicVideoCountNil(b bool)`

 SetMusicVideoCountNil sets the value for MusicVideoCount to be an explicit nil

### UnsetMusicVideoCount
`func (o *BaseItemDto) UnsetMusicVideoCount()`

UnsetMusicVideoCount ensures that no value is present for MusicVideoCount, not even an explicit nil
### GetLockData

`func (o *BaseItemDto) GetLockData() bool`

GetLockData returns the LockData field if non-nil, zero value otherwise.

### GetLockDataOk

`func (o *BaseItemDto) GetLockDataOk() (*bool, bool)`

GetLockDataOk returns a tuple with the LockData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockData

`func (o *BaseItemDto) SetLockData(v bool)`

SetLockData sets LockData field to given value.

### HasLockData

`func (o *BaseItemDto) HasLockData() bool`

HasLockData returns a boolean if a field has been set.

### SetLockDataNil

`func (o *BaseItemDto) SetLockDataNil(b bool)`

 SetLockDataNil sets the value for LockData to be an explicit nil

### UnsetLockData
`func (o *BaseItemDto) UnsetLockData()`

UnsetLockData ensures that no value is present for LockData, not even an explicit nil
### GetWidth

`func (o *BaseItemDto) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *BaseItemDto) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *BaseItemDto) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *BaseItemDto) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *BaseItemDto) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *BaseItemDto) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *BaseItemDto) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *BaseItemDto) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *BaseItemDto) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *BaseItemDto) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *BaseItemDto) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *BaseItemDto) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetCameraMake

`func (o *BaseItemDto) GetCameraMake() string`

GetCameraMake returns the CameraMake field if non-nil, zero value otherwise.

### GetCameraMakeOk

`func (o *BaseItemDto) GetCameraMakeOk() (*string, bool)`

GetCameraMakeOk returns a tuple with the CameraMake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraMake

`func (o *BaseItemDto) SetCameraMake(v string)`

SetCameraMake sets CameraMake field to given value.

### HasCameraMake

`func (o *BaseItemDto) HasCameraMake() bool`

HasCameraMake returns a boolean if a field has been set.

### SetCameraMakeNil

`func (o *BaseItemDto) SetCameraMakeNil(b bool)`

 SetCameraMakeNil sets the value for CameraMake to be an explicit nil

### UnsetCameraMake
`func (o *BaseItemDto) UnsetCameraMake()`

UnsetCameraMake ensures that no value is present for CameraMake, not even an explicit nil
### GetCameraModel

`func (o *BaseItemDto) GetCameraModel() string`

GetCameraModel returns the CameraModel field if non-nil, zero value otherwise.

### GetCameraModelOk

`func (o *BaseItemDto) GetCameraModelOk() (*string, bool)`

GetCameraModelOk returns a tuple with the CameraModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraModel

`func (o *BaseItemDto) SetCameraModel(v string)`

SetCameraModel sets CameraModel field to given value.

### HasCameraModel

`func (o *BaseItemDto) HasCameraModel() bool`

HasCameraModel returns a boolean if a field has been set.

### SetCameraModelNil

`func (o *BaseItemDto) SetCameraModelNil(b bool)`

 SetCameraModelNil sets the value for CameraModel to be an explicit nil

### UnsetCameraModel
`func (o *BaseItemDto) UnsetCameraModel()`

UnsetCameraModel ensures that no value is present for CameraModel, not even an explicit nil
### GetSoftware

`func (o *BaseItemDto) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *BaseItemDto) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *BaseItemDto) SetSoftware(v string)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *BaseItemDto) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### SetSoftwareNil

`func (o *BaseItemDto) SetSoftwareNil(b bool)`

 SetSoftwareNil sets the value for Software to be an explicit nil

### UnsetSoftware
`func (o *BaseItemDto) UnsetSoftware()`

UnsetSoftware ensures that no value is present for Software, not even an explicit nil
### GetExposureTime

`func (o *BaseItemDto) GetExposureTime() float64`

GetExposureTime returns the ExposureTime field if non-nil, zero value otherwise.

### GetExposureTimeOk

`func (o *BaseItemDto) GetExposureTimeOk() (*float64, bool)`

GetExposureTimeOk returns a tuple with the ExposureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureTime

`func (o *BaseItemDto) SetExposureTime(v float64)`

SetExposureTime sets ExposureTime field to given value.

### HasExposureTime

`func (o *BaseItemDto) HasExposureTime() bool`

HasExposureTime returns a boolean if a field has been set.

### SetExposureTimeNil

`func (o *BaseItemDto) SetExposureTimeNil(b bool)`

 SetExposureTimeNil sets the value for ExposureTime to be an explicit nil

### UnsetExposureTime
`func (o *BaseItemDto) UnsetExposureTime()`

UnsetExposureTime ensures that no value is present for ExposureTime, not even an explicit nil
### GetFocalLength

`func (o *BaseItemDto) GetFocalLength() float64`

GetFocalLength returns the FocalLength field if non-nil, zero value otherwise.

### GetFocalLengthOk

`func (o *BaseItemDto) GetFocalLengthOk() (*float64, bool)`

GetFocalLengthOk returns a tuple with the FocalLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFocalLength

`func (o *BaseItemDto) SetFocalLength(v float64)`

SetFocalLength sets FocalLength field to given value.

### HasFocalLength

`func (o *BaseItemDto) HasFocalLength() bool`

HasFocalLength returns a boolean if a field has been set.

### SetFocalLengthNil

`func (o *BaseItemDto) SetFocalLengthNil(b bool)`

 SetFocalLengthNil sets the value for FocalLength to be an explicit nil

### UnsetFocalLength
`func (o *BaseItemDto) UnsetFocalLength()`

UnsetFocalLength ensures that no value is present for FocalLength, not even an explicit nil
### GetImageOrientation

`func (o *BaseItemDto) GetImageOrientation() ImageOrientation`

GetImageOrientation returns the ImageOrientation field if non-nil, zero value otherwise.

### GetImageOrientationOk

`func (o *BaseItemDto) GetImageOrientationOk() (*ImageOrientation, bool)`

GetImageOrientationOk returns a tuple with the ImageOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOrientation

`func (o *BaseItemDto) SetImageOrientation(v ImageOrientation)`

SetImageOrientation sets ImageOrientation field to given value.

### HasImageOrientation

`func (o *BaseItemDto) HasImageOrientation() bool`

HasImageOrientation returns a boolean if a field has been set.

### SetImageOrientationNil

`func (o *BaseItemDto) SetImageOrientationNil(b bool)`

 SetImageOrientationNil sets the value for ImageOrientation to be an explicit nil

### UnsetImageOrientation
`func (o *BaseItemDto) UnsetImageOrientation()`

UnsetImageOrientation ensures that no value is present for ImageOrientation, not even an explicit nil
### GetAperture

`func (o *BaseItemDto) GetAperture() float64`

GetAperture returns the Aperture field if non-nil, zero value otherwise.

### GetApertureOk

`func (o *BaseItemDto) GetApertureOk() (*float64, bool)`

GetApertureOk returns a tuple with the Aperture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAperture

`func (o *BaseItemDto) SetAperture(v float64)`

SetAperture sets Aperture field to given value.

### HasAperture

`func (o *BaseItemDto) HasAperture() bool`

HasAperture returns a boolean if a field has been set.

### SetApertureNil

`func (o *BaseItemDto) SetApertureNil(b bool)`

 SetApertureNil sets the value for Aperture to be an explicit nil

### UnsetAperture
`func (o *BaseItemDto) UnsetAperture()`

UnsetAperture ensures that no value is present for Aperture, not even an explicit nil
### GetShutterSpeed

`func (o *BaseItemDto) GetShutterSpeed() float64`

GetShutterSpeed returns the ShutterSpeed field if non-nil, zero value otherwise.

### GetShutterSpeedOk

`func (o *BaseItemDto) GetShutterSpeedOk() (*float64, bool)`

GetShutterSpeedOk returns a tuple with the ShutterSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutterSpeed

`func (o *BaseItemDto) SetShutterSpeed(v float64)`

SetShutterSpeed sets ShutterSpeed field to given value.

### HasShutterSpeed

`func (o *BaseItemDto) HasShutterSpeed() bool`

HasShutterSpeed returns a boolean if a field has been set.

### SetShutterSpeedNil

`func (o *BaseItemDto) SetShutterSpeedNil(b bool)`

 SetShutterSpeedNil sets the value for ShutterSpeed to be an explicit nil

### UnsetShutterSpeed
`func (o *BaseItemDto) UnsetShutterSpeed()`

UnsetShutterSpeed ensures that no value is present for ShutterSpeed, not even an explicit nil
### GetLatitude

`func (o *BaseItemDto) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *BaseItemDto) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *BaseItemDto) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *BaseItemDto) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *BaseItemDto) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *BaseItemDto) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *BaseItemDto) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *BaseItemDto) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *BaseItemDto) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *BaseItemDto) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *BaseItemDto) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *BaseItemDto) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetAltitude

`func (o *BaseItemDto) GetAltitude() float64`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *BaseItemDto) GetAltitudeOk() (*float64, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *BaseItemDto) SetAltitude(v float64)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *BaseItemDto) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### SetAltitudeNil

`func (o *BaseItemDto) SetAltitudeNil(b bool)`

 SetAltitudeNil sets the value for Altitude to be an explicit nil

### UnsetAltitude
`func (o *BaseItemDto) UnsetAltitude()`

UnsetAltitude ensures that no value is present for Altitude, not even an explicit nil
### GetIsoSpeedRating

`func (o *BaseItemDto) GetIsoSpeedRating() int32`

GetIsoSpeedRating returns the IsoSpeedRating field if non-nil, zero value otherwise.

### GetIsoSpeedRatingOk

`func (o *BaseItemDto) GetIsoSpeedRatingOk() (*int32, bool)`

GetIsoSpeedRatingOk returns a tuple with the IsoSpeedRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoSpeedRating

`func (o *BaseItemDto) SetIsoSpeedRating(v int32)`

SetIsoSpeedRating sets IsoSpeedRating field to given value.

### HasIsoSpeedRating

`func (o *BaseItemDto) HasIsoSpeedRating() bool`

HasIsoSpeedRating returns a boolean if a field has been set.

### SetIsoSpeedRatingNil

`func (o *BaseItemDto) SetIsoSpeedRatingNil(b bool)`

 SetIsoSpeedRatingNil sets the value for IsoSpeedRating to be an explicit nil

### UnsetIsoSpeedRating
`func (o *BaseItemDto) UnsetIsoSpeedRating()`

UnsetIsoSpeedRating ensures that no value is present for IsoSpeedRating, not even an explicit nil
### GetSeriesTimerId

`func (o *BaseItemDto) GetSeriesTimerId() string`

GetSeriesTimerId returns the SeriesTimerId field if non-nil, zero value otherwise.

### GetSeriesTimerIdOk

`func (o *BaseItemDto) GetSeriesTimerIdOk() (*string, bool)`

GetSeriesTimerIdOk returns a tuple with the SeriesTimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTimerId

`func (o *BaseItemDto) SetSeriesTimerId(v string)`

SetSeriesTimerId sets SeriesTimerId field to given value.

### HasSeriesTimerId

`func (o *BaseItemDto) HasSeriesTimerId() bool`

HasSeriesTimerId returns a boolean if a field has been set.

### SetSeriesTimerIdNil

`func (o *BaseItemDto) SetSeriesTimerIdNil(b bool)`

 SetSeriesTimerIdNil sets the value for SeriesTimerId to be an explicit nil

### UnsetSeriesTimerId
`func (o *BaseItemDto) UnsetSeriesTimerId()`

UnsetSeriesTimerId ensures that no value is present for SeriesTimerId, not even an explicit nil
### GetProgramId

`func (o *BaseItemDto) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *BaseItemDto) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *BaseItemDto) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *BaseItemDto) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramIdNil

`func (o *BaseItemDto) SetProgramIdNil(b bool)`

 SetProgramIdNil sets the value for ProgramId to be an explicit nil

### UnsetProgramId
`func (o *BaseItemDto) UnsetProgramId()`

UnsetProgramId ensures that no value is present for ProgramId, not even an explicit nil
### GetChannelPrimaryImageTag

`func (o *BaseItemDto) GetChannelPrimaryImageTag() string`

GetChannelPrimaryImageTag returns the ChannelPrimaryImageTag field if non-nil, zero value otherwise.

### GetChannelPrimaryImageTagOk

`func (o *BaseItemDto) GetChannelPrimaryImageTagOk() (*string, bool)`

GetChannelPrimaryImageTagOk returns a tuple with the ChannelPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPrimaryImageTag

`func (o *BaseItemDto) SetChannelPrimaryImageTag(v string)`

SetChannelPrimaryImageTag sets ChannelPrimaryImageTag field to given value.

### HasChannelPrimaryImageTag

`func (o *BaseItemDto) HasChannelPrimaryImageTag() bool`

HasChannelPrimaryImageTag returns a boolean if a field has been set.

### SetChannelPrimaryImageTagNil

`func (o *BaseItemDto) SetChannelPrimaryImageTagNil(b bool)`

 SetChannelPrimaryImageTagNil sets the value for ChannelPrimaryImageTag to be an explicit nil

### UnsetChannelPrimaryImageTag
`func (o *BaseItemDto) UnsetChannelPrimaryImageTag()`

UnsetChannelPrimaryImageTag ensures that no value is present for ChannelPrimaryImageTag, not even an explicit nil
### GetStartDate

`func (o *BaseItemDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BaseItemDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BaseItemDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BaseItemDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *BaseItemDto) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *BaseItemDto) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetCompletionPercentage

`func (o *BaseItemDto) GetCompletionPercentage() float64`

GetCompletionPercentage returns the CompletionPercentage field if non-nil, zero value otherwise.

### GetCompletionPercentageOk

`func (o *BaseItemDto) GetCompletionPercentageOk() (*float64, bool)`

GetCompletionPercentageOk returns a tuple with the CompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionPercentage

`func (o *BaseItemDto) SetCompletionPercentage(v float64)`

SetCompletionPercentage sets CompletionPercentage field to given value.

### HasCompletionPercentage

`func (o *BaseItemDto) HasCompletionPercentage() bool`

HasCompletionPercentage returns a boolean if a field has been set.

### SetCompletionPercentageNil

`func (o *BaseItemDto) SetCompletionPercentageNil(b bool)`

 SetCompletionPercentageNil sets the value for CompletionPercentage to be an explicit nil

### UnsetCompletionPercentage
`func (o *BaseItemDto) UnsetCompletionPercentage()`

UnsetCompletionPercentage ensures that no value is present for CompletionPercentage, not even an explicit nil
### GetIsRepeat

`func (o *BaseItemDto) GetIsRepeat() bool`

GetIsRepeat returns the IsRepeat field if non-nil, zero value otherwise.

### GetIsRepeatOk

`func (o *BaseItemDto) GetIsRepeatOk() (*bool, bool)`

GetIsRepeatOk returns a tuple with the IsRepeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRepeat

`func (o *BaseItemDto) SetIsRepeat(v bool)`

SetIsRepeat sets IsRepeat field to given value.

### HasIsRepeat

`func (o *BaseItemDto) HasIsRepeat() bool`

HasIsRepeat returns a boolean if a field has been set.

### SetIsRepeatNil

`func (o *BaseItemDto) SetIsRepeatNil(b bool)`

 SetIsRepeatNil sets the value for IsRepeat to be an explicit nil

### UnsetIsRepeat
`func (o *BaseItemDto) UnsetIsRepeat()`

UnsetIsRepeat ensures that no value is present for IsRepeat, not even an explicit nil
### GetEpisodeTitle

`func (o *BaseItemDto) GetEpisodeTitle() string`

GetEpisodeTitle returns the EpisodeTitle field if non-nil, zero value otherwise.

### GetEpisodeTitleOk

`func (o *BaseItemDto) GetEpisodeTitleOk() (*string, bool)`

GetEpisodeTitleOk returns a tuple with the EpisodeTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeTitle

`func (o *BaseItemDto) SetEpisodeTitle(v string)`

SetEpisodeTitle sets EpisodeTitle field to given value.

### HasEpisodeTitle

`func (o *BaseItemDto) HasEpisodeTitle() bool`

HasEpisodeTitle returns a boolean if a field has been set.

### SetEpisodeTitleNil

`func (o *BaseItemDto) SetEpisodeTitleNil(b bool)`

 SetEpisodeTitleNil sets the value for EpisodeTitle to be an explicit nil

### UnsetEpisodeTitle
`func (o *BaseItemDto) UnsetEpisodeTitle()`

UnsetEpisodeTitle ensures that no value is present for EpisodeTitle, not even an explicit nil
### GetChannelType

`func (o *BaseItemDto) GetChannelType() ChannelType`

GetChannelType returns the ChannelType field if non-nil, zero value otherwise.

### GetChannelTypeOk

`func (o *BaseItemDto) GetChannelTypeOk() (*ChannelType, bool)`

GetChannelTypeOk returns a tuple with the ChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelType

`func (o *BaseItemDto) SetChannelType(v ChannelType)`

SetChannelType sets ChannelType field to given value.

### HasChannelType

`func (o *BaseItemDto) HasChannelType() bool`

HasChannelType returns a boolean if a field has been set.

### SetChannelTypeNil

`func (o *BaseItemDto) SetChannelTypeNil(b bool)`

 SetChannelTypeNil sets the value for ChannelType to be an explicit nil

### UnsetChannelType
`func (o *BaseItemDto) UnsetChannelType()`

UnsetChannelType ensures that no value is present for ChannelType, not even an explicit nil
### GetAudio

`func (o *BaseItemDto) GetAudio() ProgramAudio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *BaseItemDto) GetAudioOk() (*ProgramAudio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *BaseItemDto) SetAudio(v ProgramAudio)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *BaseItemDto) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### SetAudioNil

`func (o *BaseItemDto) SetAudioNil(b bool)`

 SetAudioNil sets the value for Audio to be an explicit nil

### UnsetAudio
`func (o *BaseItemDto) UnsetAudio()`

UnsetAudio ensures that no value is present for Audio, not even an explicit nil
### GetIsMovie

`func (o *BaseItemDto) GetIsMovie() bool`

GetIsMovie returns the IsMovie field if non-nil, zero value otherwise.

### GetIsMovieOk

`func (o *BaseItemDto) GetIsMovieOk() (*bool, bool)`

GetIsMovieOk returns a tuple with the IsMovie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMovie

`func (o *BaseItemDto) SetIsMovie(v bool)`

SetIsMovie sets IsMovie field to given value.

### HasIsMovie

`func (o *BaseItemDto) HasIsMovie() bool`

HasIsMovie returns a boolean if a field has been set.

### SetIsMovieNil

`func (o *BaseItemDto) SetIsMovieNil(b bool)`

 SetIsMovieNil sets the value for IsMovie to be an explicit nil

### UnsetIsMovie
`func (o *BaseItemDto) UnsetIsMovie()`

UnsetIsMovie ensures that no value is present for IsMovie, not even an explicit nil
### GetIsSports

`func (o *BaseItemDto) GetIsSports() bool`

GetIsSports returns the IsSports field if non-nil, zero value otherwise.

### GetIsSportsOk

`func (o *BaseItemDto) GetIsSportsOk() (*bool, bool)`

GetIsSportsOk returns a tuple with the IsSports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSports

`func (o *BaseItemDto) SetIsSports(v bool)`

SetIsSports sets IsSports field to given value.

### HasIsSports

`func (o *BaseItemDto) HasIsSports() bool`

HasIsSports returns a boolean if a field has been set.

### SetIsSportsNil

`func (o *BaseItemDto) SetIsSportsNil(b bool)`

 SetIsSportsNil sets the value for IsSports to be an explicit nil

### UnsetIsSports
`func (o *BaseItemDto) UnsetIsSports()`

UnsetIsSports ensures that no value is present for IsSports, not even an explicit nil
### GetIsSeries

`func (o *BaseItemDto) GetIsSeries() bool`

GetIsSeries returns the IsSeries field if non-nil, zero value otherwise.

### GetIsSeriesOk

`func (o *BaseItemDto) GetIsSeriesOk() (*bool, bool)`

GetIsSeriesOk returns a tuple with the IsSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeries

`func (o *BaseItemDto) SetIsSeries(v bool)`

SetIsSeries sets IsSeries field to given value.

### HasIsSeries

`func (o *BaseItemDto) HasIsSeries() bool`

HasIsSeries returns a boolean if a field has been set.

### SetIsSeriesNil

`func (o *BaseItemDto) SetIsSeriesNil(b bool)`

 SetIsSeriesNil sets the value for IsSeries to be an explicit nil

### UnsetIsSeries
`func (o *BaseItemDto) UnsetIsSeries()`

UnsetIsSeries ensures that no value is present for IsSeries, not even an explicit nil
### GetIsLive

`func (o *BaseItemDto) GetIsLive() bool`

GetIsLive returns the IsLive field if non-nil, zero value otherwise.

### GetIsLiveOk

`func (o *BaseItemDto) GetIsLiveOk() (*bool, bool)`

GetIsLiveOk returns a tuple with the IsLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLive

`func (o *BaseItemDto) SetIsLive(v bool)`

SetIsLive sets IsLive field to given value.

### HasIsLive

`func (o *BaseItemDto) HasIsLive() bool`

HasIsLive returns a boolean if a field has been set.

### SetIsLiveNil

`func (o *BaseItemDto) SetIsLiveNil(b bool)`

 SetIsLiveNil sets the value for IsLive to be an explicit nil

### UnsetIsLive
`func (o *BaseItemDto) UnsetIsLive()`

UnsetIsLive ensures that no value is present for IsLive, not even an explicit nil
### GetIsNews

`func (o *BaseItemDto) GetIsNews() bool`

GetIsNews returns the IsNews field if non-nil, zero value otherwise.

### GetIsNewsOk

`func (o *BaseItemDto) GetIsNewsOk() (*bool, bool)`

GetIsNewsOk returns a tuple with the IsNews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNews

`func (o *BaseItemDto) SetIsNews(v bool)`

SetIsNews sets IsNews field to given value.

### HasIsNews

`func (o *BaseItemDto) HasIsNews() bool`

HasIsNews returns a boolean if a field has been set.

### SetIsNewsNil

`func (o *BaseItemDto) SetIsNewsNil(b bool)`

 SetIsNewsNil sets the value for IsNews to be an explicit nil

### UnsetIsNews
`func (o *BaseItemDto) UnsetIsNews()`

UnsetIsNews ensures that no value is present for IsNews, not even an explicit nil
### GetIsKids

`func (o *BaseItemDto) GetIsKids() bool`

GetIsKids returns the IsKids field if non-nil, zero value otherwise.

### GetIsKidsOk

`func (o *BaseItemDto) GetIsKidsOk() (*bool, bool)`

GetIsKidsOk returns a tuple with the IsKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKids

`func (o *BaseItemDto) SetIsKids(v bool)`

SetIsKids sets IsKids field to given value.

### HasIsKids

`func (o *BaseItemDto) HasIsKids() bool`

HasIsKids returns a boolean if a field has been set.

### SetIsKidsNil

`func (o *BaseItemDto) SetIsKidsNil(b bool)`

 SetIsKidsNil sets the value for IsKids to be an explicit nil

### UnsetIsKids
`func (o *BaseItemDto) UnsetIsKids()`

UnsetIsKids ensures that no value is present for IsKids, not even an explicit nil
### GetIsPremiere

`func (o *BaseItemDto) GetIsPremiere() bool`

GetIsPremiere returns the IsPremiere field if non-nil, zero value otherwise.

### GetIsPremiereOk

`func (o *BaseItemDto) GetIsPremiereOk() (*bool, bool)`

GetIsPremiereOk returns a tuple with the IsPremiere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPremiere

`func (o *BaseItemDto) SetIsPremiere(v bool)`

SetIsPremiere sets IsPremiere field to given value.

### HasIsPremiere

`func (o *BaseItemDto) HasIsPremiere() bool`

HasIsPremiere returns a boolean if a field has been set.

### SetIsPremiereNil

`func (o *BaseItemDto) SetIsPremiereNil(b bool)`

 SetIsPremiereNil sets the value for IsPremiere to be an explicit nil

### UnsetIsPremiere
`func (o *BaseItemDto) UnsetIsPremiere()`

UnsetIsPremiere ensures that no value is present for IsPremiere, not even an explicit nil
### GetTimerId

`func (o *BaseItemDto) GetTimerId() string`

GetTimerId returns the TimerId field if non-nil, zero value otherwise.

### GetTimerIdOk

`func (o *BaseItemDto) GetTimerIdOk() (*string, bool)`

GetTimerIdOk returns a tuple with the TimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerId

`func (o *BaseItemDto) SetTimerId(v string)`

SetTimerId sets TimerId field to given value.

### HasTimerId

`func (o *BaseItemDto) HasTimerId() bool`

HasTimerId returns a boolean if a field has been set.

### SetTimerIdNil

`func (o *BaseItemDto) SetTimerIdNil(b bool)`

 SetTimerIdNil sets the value for TimerId to be an explicit nil

### UnsetTimerId
`func (o *BaseItemDto) UnsetTimerId()`

UnsetTimerId ensures that no value is present for TimerId, not even an explicit nil
### GetCurrentProgram

`func (o *BaseItemDto) GetCurrentProgram() BaseItemDtoCurrentProgram`

GetCurrentProgram returns the CurrentProgram field if non-nil, zero value otherwise.

### GetCurrentProgramOk

`func (o *BaseItemDto) GetCurrentProgramOk() (*BaseItemDtoCurrentProgram, bool)`

GetCurrentProgramOk returns a tuple with the CurrentProgram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProgram

`func (o *BaseItemDto) SetCurrentProgram(v BaseItemDtoCurrentProgram)`

SetCurrentProgram sets CurrentProgram field to given value.

### HasCurrentProgram

`func (o *BaseItemDto) HasCurrentProgram() bool`

HasCurrentProgram returns a boolean if a field has been set.

### SetCurrentProgramNil

`func (o *BaseItemDto) SetCurrentProgramNil(b bool)`

 SetCurrentProgramNil sets the value for CurrentProgram to be an explicit nil

### UnsetCurrentProgram
`func (o *BaseItemDto) UnsetCurrentProgram()`

UnsetCurrentProgram ensures that no value is present for CurrentProgram, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


