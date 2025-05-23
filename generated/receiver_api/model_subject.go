/*
SUSE Observability Receiver API

This API documentation page describes the SUSE Observability receiver API.

API version: 5.2.0
Contact: info@stackstate.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package receiver_api

import (
	"encoding/json"
)

// Subject Subject contains a reference to the object or user identities a role binding applies to.
type Subject struct {
	// APIGroup holds the API group of the referenced subject.
	ApiGroup *string `json:"apiGroup,omitempty"`
	// Kind of object being referenced.
	Kind string `json:"kind"`
	// Name of the object being referenced.
	Name string `json:"name"`
	// Namespace of the referenced object.
	Namespace *string `json:"namespace,omitempty"`
}

// NewSubject instantiates a new Subject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubject(kind string, name string) *Subject {
	this := Subject{}
	this.Kind = kind
	this.Name = name
	return &this
}

// NewSubjectWithDefaults instantiates a new Subject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectWithDefaults() *Subject {
	this := Subject{}
	return &this
}

// GetApiGroup returns the ApiGroup field value if set, zero value otherwise.
func (o *Subject) GetApiGroup() string {
	if o == nil || o.ApiGroup == nil {
		var ret string
		return ret
	}
	return *o.ApiGroup
}

// GetApiGroupOk returns a tuple with the ApiGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subject) GetApiGroupOk() (*string, bool) {
	if o == nil || o.ApiGroup == nil {
		return nil, false
	}
	return o.ApiGroup, true
}

// HasApiGroup returns a boolean if a field has been set.
func (o *Subject) HasApiGroup() bool {
	if o != nil && o.ApiGroup != nil {
		return true
	}

	return false
}

// SetApiGroup gets a reference to the given string and assigns it to the ApiGroup field.
func (o *Subject) SetApiGroup(v string) {
	o.ApiGroup = &v
}

// GetKind returns the Kind field value
func (o *Subject) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Subject) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *Subject) SetKind(v string) {
	o.Kind = v
}

// GetName returns the Name field value
func (o *Subject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Subject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Subject) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *Subject) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subject) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Subject) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *Subject) SetNamespace(v string) {
	o.Namespace = &v
}

func (o Subject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiGroup != nil {
		toSerialize["apiGroup"] = o.ApiGroup
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableSubject struct {
	value *Subject
	isSet bool
}

func (v NullableSubject) Get() *Subject {
	return v.value
}

func (v *NullableSubject) Set(val *Subject) {
	v.value = val
	v.isSet = true
}

func (v NullableSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubject(val *Subject) *NullableSubject {
	return &NullableSubject{value: val, isSet: true}
}

func (v NullableSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
