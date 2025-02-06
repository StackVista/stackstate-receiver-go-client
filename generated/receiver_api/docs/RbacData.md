# RbacData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. | 
**Metadata** | [**ObjectMeta**](ObjectMeta.md) |  | 
**AggregationRule** | Pointer to [**AggregationRule**](AggregationRule.md) |  | [optional] 
**Rules** | Pointer to [**[]PolicyRule**](PolicyRule.md) | Rules holds all the PolicyRules for this Role. | [optional] 
**RoleRef** | [**RoleRef**](RoleRef.md) |  | 
**Subjects** | Pointer to [**[]Subject**](Subject.md) | Subjects holds references to the objects the role applies to. | [optional] 

## Methods

### NewRbacData

`func NewRbacData(kind string, metadata ObjectMeta, roleRef RoleRef, ) *RbacData`

NewRbacData instantiates a new RbacData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacDataWithDefaults

`func NewRbacDataWithDefaults() *RbacData`

NewRbacDataWithDefaults instantiates a new RbacData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *RbacData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RbacData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RbacData) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *RbacData) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RbacData) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RbacData) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.


### GetAggregationRule

`func (o *RbacData) GetAggregationRule() AggregationRule`

GetAggregationRule returns the AggregationRule field if non-nil, zero value otherwise.

### GetAggregationRuleOk

`func (o *RbacData) GetAggregationRuleOk() (*AggregationRule, bool)`

GetAggregationRuleOk returns a tuple with the AggregationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationRule

`func (o *RbacData) SetAggregationRule(v AggregationRule)`

SetAggregationRule sets AggregationRule field to given value.

### HasAggregationRule

`func (o *RbacData) HasAggregationRule() bool`

HasAggregationRule returns a boolean if a field has been set.

### GetRules

`func (o *RbacData) GetRules() []PolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *RbacData) GetRulesOk() (*[]PolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *RbacData) SetRules(v []PolicyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *RbacData) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetRoleRef

`func (o *RbacData) GetRoleRef() RoleRef`

GetRoleRef returns the RoleRef field if non-nil, zero value otherwise.

### GetRoleRefOk

`func (o *RbacData) GetRoleRefOk() (*RoleRef, bool)`

GetRoleRefOk returns a tuple with the RoleRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleRef

`func (o *RbacData) SetRoleRef(v RoleRef)`

SetRoleRef sets RoleRef field to given value.


### GetSubjects

`func (o *RbacData) GetSubjects() []Subject`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *RbacData) GetSubjectsOk() (*[]Subject, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *RbacData) SetSubjects(v []Subject)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *RbacData) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


