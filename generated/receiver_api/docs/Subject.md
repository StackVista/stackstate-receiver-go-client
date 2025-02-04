# Subject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroup** | Pointer to **string** | APIGroup holds the API group of the referenced subject. | [optional] 
**Kind** | **string** | Kind of object being referenced. | 
**Name** | **string** | Name of the object being referenced. | 
**Namespace** | Pointer to **string** | Namespace of the referenced object. | [optional] 

## Methods

### NewSubject

`func NewSubject(kind string, name string, ) *Subject`

NewSubject instantiates a new Subject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectWithDefaults

`func NewSubjectWithDefaults() *Subject`

NewSubjectWithDefaults instantiates a new Subject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiGroup

`func (o *Subject) GetApiGroup() string`

GetApiGroup returns the ApiGroup field if non-nil, zero value otherwise.

### GetApiGroupOk

`func (o *Subject) GetApiGroupOk() (*string, bool)`

GetApiGroupOk returns a tuple with the ApiGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGroup

`func (o *Subject) SetApiGroup(v string)`

SetApiGroup sets ApiGroup field to given value.

### HasApiGroup

`func (o *Subject) HasApiGroup() bool`

HasApiGroup returns a boolean if a field has been set.

### GetKind

`func (o *Subject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Subject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Subject) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *Subject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subject) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *Subject) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Subject) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Subject) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Subject) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


