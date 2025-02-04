# ManagedFieldsEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the version of this resource that this field set applies to. The format is \&quot;group/version\&quot; just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted. | [optional] 
**FieldsType** | Pointer to **string** | FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: \&quot;FieldsV1\&quot; | [optional] 
**Manager** | Pointer to **string** | Manager is an identifier of the workflow managing these fields. | [optional] 
**Operation** | Pointer to **string** | Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are &#39;Apply&#39; and &#39;Update&#39;. | [optional] 
**Subresource** | Pointer to **string** | Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource. | [optional] 
**Time** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. | [optional] 

## Methods

### NewManagedFieldsEntry

`func NewManagedFieldsEntry() *ManagedFieldsEntry`

NewManagedFieldsEntry instantiates a new ManagedFieldsEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedFieldsEntryWithDefaults

`func NewManagedFieldsEntryWithDefaults() *ManagedFieldsEntry`

NewManagedFieldsEntryWithDefaults instantiates a new ManagedFieldsEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ManagedFieldsEntry) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ManagedFieldsEntry) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ManagedFieldsEntry) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ManagedFieldsEntry) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetFieldsType

`func (o *ManagedFieldsEntry) GetFieldsType() string`

GetFieldsType returns the FieldsType field if non-nil, zero value otherwise.

### GetFieldsTypeOk

`func (o *ManagedFieldsEntry) GetFieldsTypeOk() (*string, bool)`

GetFieldsTypeOk returns a tuple with the FieldsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsType

`func (o *ManagedFieldsEntry) SetFieldsType(v string)`

SetFieldsType sets FieldsType field to given value.

### HasFieldsType

`func (o *ManagedFieldsEntry) HasFieldsType() bool`

HasFieldsType returns a boolean if a field has been set.

### GetManager

`func (o *ManagedFieldsEntry) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *ManagedFieldsEntry) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *ManagedFieldsEntry) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *ManagedFieldsEntry) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetOperation

`func (o *ManagedFieldsEntry) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ManagedFieldsEntry) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ManagedFieldsEntry) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ManagedFieldsEntry) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetSubresource

`func (o *ManagedFieldsEntry) GetSubresource() string`

GetSubresource returns the Subresource field if non-nil, zero value otherwise.

### GetSubresourceOk

`func (o *ManagedFieldsEntry) GetSubresourceOk() (*string, bool)`

GetSubresourceOk returns a tuple with the Subresource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresource

`func (o *ManagedFieldsEntry) SetSubresource(v string)`

SetSubresource sets Subresource field to given value.

### HasSubresource

`func (o *ManagedFieldsEntry) HasSubresource() bool`

HasSubresource returns a boolean if a field has been set.

### GetTime

`func (o *ManagedFieldsEntry) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ManagedFieldsEntry) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ManagedFieldsEntry) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ManagedFieldsEntry) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


