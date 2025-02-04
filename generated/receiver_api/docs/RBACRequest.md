# RBACRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**CollectionTimestamp** | **int64** | Timestamp where the data was collected by the RBAC Agent | 
**Sequence** | **int32** | Incremental number for snapshot batches. Helpful to detect incomplete snapshots that could lead to incorrect conclusions. | 
**Cluster** | **string** | Cluster name which identifies the scope of the RBAC data | 
**StartSnapshot** | Pointer to [**StartSnapshot**](StartSnapshot.md) |  | [optional] 
**StopSnapshot** | Pointer to **map[string]interface{}** | Object that signals that an open Snapshot needs to be closed after ingesting the RBAC data | [optional] 
**RbacData** | [**[]RbacData**](RbacData.md) |  | 
**Changes** | [**[]RbacDataChanges**](RbacDataChanges.md) |  | 

## Methods

### NewRBACRequest

`func NewRBACRequest(type_ string, collectionTimestamp int64, sequence int32, cluster string, rbacData []RbacData, changes []RbacDataChanges, ) *RBACRequest`

NewRBACRequest instantiates a new RBACRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRBACRequestWithDefaults

`func NewRBACRequestWithDefaults() *RBACRequest`

NewRBACRequestWithDefaults instantiates a new RBACRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RBACRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RBACRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RBACRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCollectionTimestamp

`func (o *RBACRequest) GetCollectionTimestamp() int64`

GetCollectionTimestamp returns the CollectionTimestamp field if non-nil, zero value otherwise.

### GetCollectionTimestampOk

`func (o *RBACRequest) GetCollectionTimestampOk() (*int64, bool)`

GetCollectionTimestampOk returns a tuple with the CollectionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTimestamp

`func (o *RBACRequest) SetCollectionTimestamp(v int64)`

SetCollectionTimestamp sets CollectionTimestamp field to given value.


### GetSequence

`func (o *RBACRequest) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *RBACRequest) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *RBACRequest) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetCluster

`func (o *RBACRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *RBACRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *RBACRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetStartSnapshot

`func (o *RBACRequest) GetStartSnapshot() StartSnapshot`

GetStartSnapshot returns the StartSnapshot field if non-nil, zero value otherwise.

### GetStartSnapshotOk

`func (o *RBACRequest) GetStartSnapshotOk() (*StartSnapshot, bool)`

GetStartSnapshotOk returns a tuple with the StartSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSnapshot

`func (o *RBACRequest) SetStartSnapshot(v StartSnapshot)`

SetStartSnapshot sets StartSnapshot field to given value.

### HasStartSnapshot

`func (o *RBACRequest) HasStartSnapshot() bool`

HasStartSnapshot returns a boolean if a field has been set.

### GetStopSnapshot

`func (o *RBACRequest) GetStopSnapshot() map[string]interface{}`

GetStopSnapshot returns the StopSnapshot field if non-nil, zero value otherwise.

### GetStopSnapshotOk

`func (o *RBACRequest) GetStopSnapshotOk() (*map[string]interface{}, bool)`

GetStopSnapshotOk returns a tuple with the StopSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopSnapshot

`func (o *RBACRequest) SetStopSnapshot(v map[string]interface{})`

SetStopSnapshot sets StopSnapshot field to given value.

### HasStopSnapshot

`func (o *RBACRequest) HasStopSnapshot() bool`

HasStopSnapshot returns a boolean if a field has been set.

### GetRbacData

`func (o *RBACRequest) GetRbacData() []RbacData`

GetRbacData returns the RbacData field if non-nil, zero value otherwise.

### GetRbacDataOk

`func (o *RBACRequest) GetRbacDataOk() (*[]RbacData, bool)`

GetRbacDataOk returns a tuple with the RbacData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbacData

`func (o *RBACRequest) SetRbacData(v []RbacData)`

SetRbacData sets RbacData field to given value.


### GetChanges

`func (o *RBACRequest) GetChanges() []RbacDataChanges`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *RBACRequest) GetChangesOk() (*[]RbacDataChanges, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *RBACRequest) SetChanges(v []RbacDataChanges)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


