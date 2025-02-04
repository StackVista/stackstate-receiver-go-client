# LabelSelectorRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | key is the label key that the selector applies to. | 
**Operator** | **string** | operator represents a key&#39;s relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist. | 
**Values** | Pointer to **[]string** | values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch. | [optional] 

## Methods

### NewLabelSelectorRequirement

`func NewLabelSelectorRequirement(key string, operator string, ) *LabelSelectorRequirement`

NewLabelSelectorRequirement instantiates a new LabelSelectorRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelSelectorRequirementWithDefaults

`func NewLabelSelectorRequirementWithDefaults() *LabelSelectorRequirement`

NewLabelSelectorRequirementWithDefaults instantiates a new LabelSelectorRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *LabelSelectorRequirement) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LabelSelectorRequirement) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LabelSelectorRequirement) SetKey(v string)`

SetKey sets Key field to given value.


### GetOperator

`func (o *LabelSelectorRequirement) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LabelSelectorRequirement) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LabelSelectorRequirement) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValues

`func (o *LabelSelectorRequirement) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *LabelSelectorRequirement) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *LabelSelectorRequirement) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *LabelSelectorRequirement) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


