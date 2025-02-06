# ObjectMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | Unique identifier for this object. Populated by the system. | 
**Name** | **string** | Unique name within a namespace. Required for resource creation. | 
**Annotations** | Pointer to **map[string]string** | Annotations is an unstructured key-value map stored with a resource that may be set by external tools. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations | [optional] 
**CreationTimestamp** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. | [optional] 
**DeletionGracePeriodSeconds** | Pointer to **int64** | Number of seconds allowed for this object to gracefully terminate before removal. | [optional] 
**DeletionTimestamp** | Pointer to **time.Time** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON. | [optional] 
**Finalizers** | Pointer to **[]string** | Must be empty before deletion. | [optional] 
**GenerateName** | Pointer to **string** | Optional prefix for generating a unique name. | [optional] 
**Generation** | Pointer to **int64** | A sequence number representing a specific generation of the desired state. | [optional] 
**Labels** | Pointer to **map[string]string** | Map of string keys and values used to organize and categorize objects. | [optional] 
**ManagedFields** | Pointer to [**[]ManagedFieldsEntry**](ManagedFieldsEntry.md) | Maps workflow-id and version to the set of fields managed by that workflow. | [optional] 
**Namespace** | Pointer to **string** | Defines the space within which each name must be unique. | [optional] 
**OwnerReferences** | Pointer to [**[]OwnerReference**](OwnerReference.md) | List of objects depended by this object. | [optional] 
**ResourceVersion** | Pointer to **string** | Opaque value representing the internal version of this object. | [optional] 

## Methods

### NewObjectMeta

`func NewObjectMeta(uid string, name string, ) *ObjectMeta`

NewObjectMeta instantiates a new ObjectMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectMetaWithDefaults

`func NewObjectMetaWithDefaults() *ObjectMeta`

NewObjectMetaWithDefaults instantiates a new ObjectMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *ObjectMeta) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ObjectMeta) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ObjectMeta) SetUid(v string)`

SetUid sets Uid field to given value.


### GetName

`func (o *ObjectMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObjectMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObjectMeta) SetName(v string)`

SetName sets Name field to given value.


### GetAnnotations

`func (o *ObjectMeta) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ObjectMeta) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ObjectMeta) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *ObjectMeta) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *ObjectMeta) GetCreationTimestamp() time.Time`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *ObjectMeta) GetCreationTimestampOk() (*time.Time, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *ObjectMeta) SetCreationTimestamp(v time.Time)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *ObjectMeta) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetDeletionGracePeriodSeconds

`func (o *ObjectMeta) GetDeletionGracePeriodSeconds() int64`

GetDeletionGracePeriodSeconds returns the DeletionGracePeriodSeconds field if non-nil, zero value otherwise.

### GetDeletionGracePeriodSecondsOk

`func (o *ObjectMeta) GetDeletionGracePeriodSecondsOk() (*int64, bool)`

GetDeletionGracePeriodSecondsOk returns a tuple with the DeletionGracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionGracePeriodSeconds

`func (o *ObjectMeta) SetDeletionGracePeriodSeconds(v int64)`

SetDeletionGracePeriodSeconds sets DeletionGracePeriodSeconds field to given value.

### HasDeletionGracePeriodSeconds

`func (o *ObjectMeta) HasDeletionGracePeriodSeconds() bool`

HasDeletionGracePeriodSeconds returns a boolean if a field has been set.

### GetDeletionTimestamp

`func (o *ObjectMeta) GetDeletionTimestamp() time.Time`

GetDeletionTimestamp returns the DeletionTimestamp field if non-nil, zero value otherwise.

### GetDeletionTimestampOk

`func (o *ObjectMeta) GetDeletionTimestampOk() (*time.Time, bool)`

GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTimestamp

`func (o *ObjectMeta) SetDeletionTimestamp(v time.Time)`

SetDeletionTimestamp sets DeletionTimestamp field to given value.

### HasDeletionTimestamp

`func (o *ObjectMeta) HasDeletionTimestamp() bool`

HasDeletionTimestamp returns a boolean if a field has been set.

### GetFinalizers

`func (o *ObjectMeta) GetFinalizers() []string`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *ObjectMeta) GetFinalizersOk() (*[]string, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *ObjectMeta) SetFinalizers(v []string)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *ObjectMeta) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetGenerateName

`func (o *ObjectMeta) GetGenerateName() string`

GetGenerateName returns the GenerateName field if non-nil, zero value otherwise.

### GetGenerateNameOk

`func (o *ObjectMeta) GetGenerateNameOk() (*string, bool)`

GetGenerateNameOk returns a tuple with the GenerateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateName

`func (o *ObjectMeta) SetGenerateName(v string)`

SetGenerateName sets GenerateName field to given value.

### HasGenerateName

`func (o *ObjectMeta) HasGenerateName() bool`

HasGenerateName returns a boolean if a field has been set.

### GetGeneration

`func (o *ObjectMeta) GetGeneration() int64`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *ObjectMeta) GetGenerationOk() (*int64, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *ObjectMeta) SetGeneration(v int64)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *ObjectMeta) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetLabels

`func (o *ObjectMeta) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ObjectMeta) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ObjectMeta) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ObjectMeta) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetManagedFields

`func (o *ObjectMeta) GetManagedFields() []ManagedFieldsEntry`

GetManagedFields returns the ManagedFields field if non-nil, zero value otherwise.

### GetManagedFieldsOk

`func (o *ObjectMeta) GetManagedFieldsOk() (*[]ManagedFieldsEntry, bool)`

GetManagedFieldsOk returns a tuple with the ManagedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedFields

`func (o *ObjectMeta) SetManagedFields(v []ManagedFieldsEntry)`

SetManagedFields sets ManagedFields field to given value.

### HasManagedFields

`func (o *ObjectMeta) HasManagedFields() bool`

HasManagedFields returns a boolean if a field has been set.

### GetNamespace

`func (o *ObjectMeta) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ObjectMeta) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ObjectMeta) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ObjectMeta) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOwnerReferences

`func (o *ObjectMeta) GetOwnerReferences() []OwnerReference`

GetOwnerReferences returns the OwnerReferences field if non-nil, zero value otherwise.

### GetOwnerReferencesOk

`func (o *ObjectMeta) GetOwnerReferencesOk() (*[]OwnerReference, bool)`

GetOwnerReferencesOk returns a tuple with the OwnerReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerReferences

`func (o *ObjectMeta) SetOwnerReferences(v []OwnerReference)`

SetOwnerReferences sets OwnerReferences field to given value.

### HasOwnerReferences

`func (o *ObjectMeta) HasOwnerReferences() bool`

HasOwnerReferences returns a boolean if a field has been set.

### GetResourceVersion

`func (o *ObjectMeta) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *ObjectMeta) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *ObjectMeta) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *ObjectMeta) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


