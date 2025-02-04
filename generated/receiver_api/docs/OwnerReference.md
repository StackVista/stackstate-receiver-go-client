# OwnerReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | UID of the referent. | 
**Name** | **string** | Name of the referent. | 
**Kind** | **string** | Kind of the referent. | 
**BlockOwnerDeletion** | Pointer to **bool** | If true, AND if the owner has the &#39;foregroundDeletion&#39; finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. | [optional] 
**Controller** | Pointer to **bool** | If true, this reference points to the managing controller. | [optional] 

## Methods

### NewOwnerReference

`func NewOwnerReference(uid string, name string, kind string, ) *OwnerReference`

NewOwnerReference instantiates a new OwnerReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerReferenceWithDefaults

`func NewOwnerReferenceWithDefaults() *OwnerReference`

NewOwnerReferenceWithDefaults instantiates a new OwnerReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *OwnerReference) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *OwnerReference) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *OwnerReference) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *OwnerReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OwnerReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OwnerReference) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *OwnerReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OwnerReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OwnerReference) SetKind(v string)`

SetKind sets Kind field to given value.


### GetBlockOwnerDeletion

`func (o *OwnerReference) GetBlockOwnerDeletion() bool`

GetBlockOwnerDeletion returns the BlockOwnerDeletion field if non-nil, zero value otherwise.

### GetBlockOwnerDeletionOk

`func (o *OwnerReference) GetBlockOwnerDeletionOk() (*bool, bool)`

GetBlockOwnerDeletionOk returns a tuple with the BlockOwnerDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockOwnerDeletion

`func (o *OwnerReference) SetBlockOwnerDeletion(v bool)`

SetBlockOwnerDeletion sets BlockOwnerDeletion field to given value.

### HasBlockOwnerDeletion

`func (o *OwnerReference) HasBlockOwnerDeletion() bool`

HasBlockOwnerDeletion returns a boolean if a field has been set.

### GetController

`func (o *OwnerReference) GetController() bool`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *OwnerReference) GetControllerOk() (*bool, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *OwnerReference) SetController(v bool)`

SetController sets Controller field to given value.

### HasController

`func (o *OwnerReference) HasController() bool`

HasController returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


