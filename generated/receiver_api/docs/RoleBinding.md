# RoleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. | 
**Metadata** | [**ObjectMeta**](ObjectMeta.md) |  | 
**RoleRef** | [**RoleRef**](RoleRef.md) |  | 
**Subjects** | Pointer to [**[]Subject**](Subject.md) | Subjects holds references to the objects the role applies to. | [optional] 

## Methods

### NewRoleBinding

`func NewRoleBinding(kind string, metadata ObjectMeta, roleRef RoleRef, ) *RoleBinding`

NewRoleBinding instantiates a new RoleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleBindingWithDefaults

`func NewRoleBindingWithDefaults() *RoleBinding`

NewRoleBindingWithDefaults instantiates a new RoleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *RoleBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RoleBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RoleBinding) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *RoleBinding) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RoleBinding) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RoleBinding) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.


### GetRoleRef

`func (o *RoleBinding) GetRoleRef() RoleRef`

GetRoleRef returns the RoleRef field if non-nil, zero value otherwise.

### GetRoleRefOk

`func (o *RoleBinding) GetRoleRefOk() (*RoleRef, bool)`

GetRoleRefOk returns a tuple with the RoleRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleRef

`func (o *RoleBinding) SetRoleRef(v RoleRef)`

SetRoleRef sets RoleRef field to given value.


### GetSubjects

`func (o *RoleBinding) GetSubjects() []Subject`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *RoleBinding) GetSubjectsOk() (*[]Subject, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *RoleBinding) SetSubjects(v []Subject)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *RoleBinding) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


