# RBACIncrementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**CollectionTimestamp** | **int64** | Timestamp where the data was collected by the RBAC Agent | 
**Cluster** | **string** | Cluster name which identifies the scope of the RBAC data | 
**Changes** | [**[]RbacDataChanges**](RbacDataChanges.md) |  | 

## Methods

### NewRBACIncrementRequest

`func NewRBACIncrementRequest(type_ string, collectionTimestamp int64, cluster string, changes []RbacDataChanges, ) *RBACIncrementRequest`

NewRBACIncrementRequest instantiates a new RBACIncrementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRBACIncrementRequestWithDefaults

`func NewRBACIncrementRequestWithDefaults() *RBACIncrementRequest`

NewRBACIncrementRequestWithDefaults instantiates a new RBACIncrementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RBACIncrementRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RBACIncrementRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RBACIncrementRequest) SetType(v string)`

SetType sets Type field to given value.


### GetCollectionTimestamp

`func (o *RBACIncrementRequest) GetCollectionTimestamp() int64`

GetCollectionTimestamp returns the CollectionTimestamp field if non-nil, zero value otherwise.

### GetCollectionTimestampOk

`func (o *RBACIncrementRequest) GetCollectionTimestampOk() (*int64, bool)`

GetCollectionTimestampOk returns a tuple with the CollectionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTimestamp

`func (o *RBACIncrementRequest) SetCollectionTimestamp(v int64)`

SetCollectionTimestamp sets CollectionTimestamp field to given value.


### GetCluster

`func (o *RBACIncrementRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *RBACIncrementRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *RBACIncrementRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetChanges

`func (o *RBACIncrementRequest) GetChanges() []RbacDataChanges`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *RBACIncrementRequest) GetChangesOk() (*[]RbacDataChanges, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *RBACIncrementRequest) SetChanges(v []RbacDataChanges)`

SetChanges sets Changes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


