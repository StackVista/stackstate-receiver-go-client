# RbacDataChanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Resource** | [**RbacData**](RbacData.md) |  | 
**Uid** | **string** | UID of the referent. | 

## Methods

### NewRbacDataChanges

`func NewRbacDataChanges(type_ string, resource RbacData, uid string, ) *RbacDataChanges`

NewRbacDataChanges instantiates a new RbacDataChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacDataChangesWithDefaults

`func NewRbacDataChangesWithDefaults() *RbacDataChanges`

NewRbacDataChangesWithDefaults instantiates a new RbacDataChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RbacDataChanges) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RbacDataChanges) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RbacDataChanges) SetType(v string)`

SetType sets Type field to given value.


### GetResource

`func (o *RbacDataChanges) GetResource() RbacData`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *RbacDataChanges) GetResourceOk() (*RbacData, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *RbacDataChanges) SetResource(v RbacData)`

SetResource sets Resource field to given value.


### GetUid

`func (o *RbacDataChanges) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *RbacDataChanges) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *RbacDataChanges) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


