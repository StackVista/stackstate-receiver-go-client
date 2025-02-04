# StartSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepeatIntervalS** | Pointer to **int64** | Number of seconds when the RBAC Agent will send the following snapshot. Heartbeat of the Agent | [optional] 

## Methods

### NewStartSnapshot

`func NewStartSnapshot() *StartSnapshot`

NewStartSnapshot instantiates a new StartSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartSnapshotWithDefaults

`func NewStartSnapshotWithDefaults() *StartSnapshot`

NewStartSnapshotWithDefaults instantiates a new StartSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepeatIntervalS

`func (o *StartSnapshot) GetRepeatIntervalS() int64`

GetRepeatIntervalS returns the RepeatIntervalS field if non-nil, zero value otherwise.

### GetRepeatIntervalSOk

`func (o *StartSnapshot) GetRepeatIntervalSOk() (*int64, bool)`

GetRepeatIntervalSOk returns a tuple with the RepeatIntervalS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIntervalS

`func (o *StartSnapshot) SetRepeatIntervalS(v int64)`

SetRepeatIntervalS sets RepeatIntervalS field to given value.

### HasRepeatIntervalS

`func (o *StartSnapshot) HasRepeatIntervalS() bool`

HasRepeatIntervalS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


