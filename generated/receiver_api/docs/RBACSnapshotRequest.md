# RBACSnapshotRequest

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

## Methods

### NewRBACSnapshotRequest

`func NewRBACSnapshotRequest(type_ string, collectionTimestamp int64, sequence int32, cluster string, rbacData []RbacData, ) *RBACSnapshotRequest`

NewRBACSnapshotRequest instantiates a new RBACSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRBACSnapshotRequestWithDefaults

`func NewRBACSnapshotRequestWithDefaults() *RBACSnapshotRequest`

NewRBACSnapshotRequestWithDefaults instantiates a new RBACSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RBACSnapshotRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RBACSnapshotRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RBACSnapshotRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCollectionTimestamp

`func (o *RBACSnapshotRequest) GetCollectionTimestamp() int64`

GetCollectionTimestamp returns the CollectionTimestamp field if non-nil, zero value otherwise.

### GetCollectionTimestampOk

`func (o *RBACSnapshotRequest) GetCollectionTimestampOk() (*int64, bool)`

GetCollectionTimestampOk returns a tuple with the CollectionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTimestamp

`func (o *RBACSnapshotRequest) SetCollectionTimestamp(v int64)`

SetCollectionTimestamp sets CollectionTimestamp field to given value.


### GetSequence

`func (o *RBACSnapshotRequest) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *RBACSnapshotRequest) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *RBACSnapshotRequest) SetSequence(v int32)`

SetSequence sets Sequence field to given value.


### GetCluster

`func (o *RBACSnapshotRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *RBACSnapshotRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *RBACSnapshotRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetStartSnapshot

`func (o *RBACSnapshotRequest) GetStartSnapshot() StartSnapshot`

GetStartSnapshot returns the StartSnapshot field if non-nil, zero value otherwise.

### GetStartSnapshotOk

`func (o *RBACSnapshotRequest) GetStartSnapshotOk() (*StartSnapshot, bool)`

GetStartSnapshotOk returns a tuple with the StartSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSnapshot

`func (o *RBACSnapshotRequest) SetStartSnapshot(v StartSnapshot)`

SetStartSnapshot sets StartSnapshot field to given value.

### HasStartSnapshot

`func (o *RBACSnapshotRequest) HasStartSnapshot() bool`

HasStartSnapshot returns a boolean if a field has been set.

### GetStopSnapshot

`func (o *RBACSnapshotRequest) GetStopSnapshot() map[string]interface{}`

GetStopSnapshot returns the StopSnapshot field if non-nil, zero value otherwise.

### GetStopSnapshotOk

`func (o *RBACSnapshotRequest) GetStopSnapshotOk() (*map[string]interface{}, bool)`

GetStopSnapshotOk returns a tuple with the StopSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopSnapshot

`func (o *RBACSnapshotRequest) SetStopSnapshot(v map[string]interface{})`

SetStopSnapshot sets StopSnapshot field to given value.

### HasStopSnapshot

`func (o *RBACSnapshotRequest) HasStopSnapshot() bool`

HasStopSnapshot returns a boolean if a field has been set.

### GetRbacData

`func (o *RBACSnapshotRequest) GetRbacData() []RbacData`

GetRbacData returns the RbacData field if non-nil, zero value otherwise.

### GetRbacDataOk

`func (o *RBACSnapshotRequest) GetRbacDataOk() (*[]RbacData, bool)`

GetRbacDataOk returns a tuple with the RbacData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbacData

`func (o *RBACSnapshotRequest) SetRbacData(v []RbacData)`

SetRbacData sets RbacData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


