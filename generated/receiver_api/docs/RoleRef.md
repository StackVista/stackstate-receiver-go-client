# RoleRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroup** | **string** | APIGroup is the group for the resource being referenced | 
**Kind** | **string** | Kind is the type of resource being referenced | 
**Name** | **string** | Name is the name of resource being referenced | 

## Methods

### NewRoleRef

`func NewRoleRef(apiGroup string, kind string, name string, ) *RoleRef`

NewRoleRef instantiates a new RoleRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleRefWithDefaults

`func NewRoleRefWithDefaults() *RoleRef`

NewRoleRefWithDefaults instantiates a new RoleRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiGroup

`func (o *RoleRef) GetApiGroup() string`

GetApiGroup returns the ApiGroup field if non-nil, zero value otherwise.

### GetApiGroupOk

`func (o *RoleRef) GetApiGroupOk() (*string, bool)`

GetApiGroupOk returns a tuple with the ApiGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGroup

`func (o *RoleRef) SetApiGroup(v string)`

SetApiGroup sets ApiGroup field to given value.


### GetKind

`func (o *RoleRef) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RoleRef) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RoleRef) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *RoleRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleRef) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


