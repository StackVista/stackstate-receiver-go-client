# ClusterRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. | 
**Metadata** | [**ObjectMeta**](ObjectMeta.md) |  | 
**AggregationRule** | Pointer to [**AggregationRule**](AggregationRule.md) |  | [optional] 
**Rules** | Pointer to [**[]PolicyRule**](PolicyRule.md) | Rules holds all the PolicyRules for this ClusterRole. | [optional] 

## Methods

### NewClusterRole

`func NewClusterRole(kind string, metadata ObjectMeta, ) *ClusterRole`

NewClusterRole instantiates a new ClusterRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRoleWithDefaults

`func NewClusterRoleWithDefaults() *ClusterRole`

NewClusterRoleWithDefaults instantiates a new ClusterRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ClusterRole) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterRole) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterRole) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *ClusterRole) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterRole) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterRole) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.


### GetAggregationRule

`func (o *ClusterRole) GetAggregationRule() AggregationRule`

GetAggregationRule returns the AggregationRule field if non-nil, zero value otherwise.

### GetAggregationRuleOk

`func (o *ClusterRole) GetAggregationRuleOk() (*AggregationRule, bool)`

GetAggregationRuleOk returns a tuple with the AggregationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationRule

`func (o *ClusterRole) SetAggregationRule(v AggregationRule)`

SetAggregationRule sets AggregationRule field to given value.

### HasAggregationRule

`func (o *ClusterRole) HasAggregationRule() bool`

HasAggregationRule returns a boolean if a field has been set.

### GetRules

`func (o *ClusterRole) GetRules() []PolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ClusterRole) GetRulesOk() (*[]PolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ClusterRole) SetRules(v []PolicyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *ClusterRole) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


