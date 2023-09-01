/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.37
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the V0037NodeAllocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0037NodeAllocation{}

// V0037NodeAllocation struct for V0037NodeAllocation
type V0037NodeAllocation struct {
	// amount of assigned job memory
	Memory *int32 `json:"memory,omitempty"`
	// amount of assigned job CPUs
	Cpus map[string]interface{} `json:"cpus,omitempty"`
	// assignment status of each socket by socket id
	Sockets map[string]interface{} `json:"sockets,omitempty"`
	// assignment status of each core by core id
	Cores map[string]interface{} `json:"cores,omitempty"`
}

// NewV0037NodeAllocation instantiates a new V0037NodeAllocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0037NodeAllocation() *V0037NodeAllocation {
	this := V0037NodeAllocation{}
	return &this
}

// NewV0037NodeAllocationWithDefaults instantiates a new V0037NodeAllocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0037NodeAllocationWithDefaults() *V0037NodeAllocation {
	this := V0037NodeAllocation{}
	return &this
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *V0037NodeAllocation) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037NodeAllocation) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *V0037NodeAllocation) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *V0037NodeAllocation) SetMemory(v int32) {
	o.Memory = &v
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *V0037NodeAllocation) GetCpus() map[string]interface{} {
	if o == nil || IsNil(o.Cpus) {
		var ret map[string]interface{}
		return ret
	}
	return o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037NodeAllocation) GetCpusOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Cpus) {
		return map[string]interface{}{}, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *V0037NodeAllocation) HasCpus() bool {
	if o != nil && !IsNil(o.Cpus) {
		return true
	}

	return false
}

// SetCpus gets a reference to the given map[string]interface{} and assigns it to the Cpus field.
func (o *V0037NodeAllocation) SetCpus(v map[string]interface{}) {
	o.Cpus = v
}

// GetSockets returns the Sockets field value if set, zero value otherwise.
func (o *V0037NodeAllocation) GetSockets() map[string]interface{} {
	if o == nil || IsNil(o.Sockets) {
		var ret map[string]interface{}
		return ret
	}
	return o.Sockets
}

// GetSocketsOk returns a tuple with the Sockets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037NodeAllocation) GetSocketsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Sockets) {
		return map[string]interface{}{}, false
	}
	return o.Sockets, true
}

// HasSockets returns a boolean if a field has been set.
func (o *V0037NodeAllocation) HasSockets() bool {
	if o != nil && !IsNil(o.Sockets) {
		return true
	}

	return false
}

// SetSockets gets a reference to the given map[string]interface{} and assigns it to the Sockets field.
func (o *V0037NodeAllocation) SetSockets(v map[string]interface{}) {
	o.Sockets = v
}

// GetCores returns the Cores field value if set, zero value otherwise.
func (o *V0037NodeAllocation) GetCores() map[string]interface{} {
	if o == nil || IsNil(o.Cores) {
		var ret map[string]interface{}
		return ret
	}
	return o.Cores
}

// GetCoresOk returns a tuple with the Cores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037NodeAllocation) GetCoresOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Cores) {
		return map[string]interface{}{}, false
	}
	return o.Cores, true
}

// HasCores returns a boolean if a field has been set.
func (o *V0037NodeAllocation) HasCores() bool {
	if o != nil && !IsNil(o.Cores) {
		return true
	}

	return false
}

// SetCores gets a reference to the given map[string]interface{} and assigns it to the Cores field.
func (o *V0037NodeAllocation) SetCores(v map[string]interface{}) {
	o.Cores = v
}

func (o V0037NodeAllocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0037NodeAllocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.Cpus) {
		toSerialize["cpus"] = o.Cpus
	}
	if !IsNil(o.Sockets) {
		toSerialize["sockets"] = o.Sockets
	}
	if !IsNil(o.Cores) {
		toSerialize["cores"] = o.Cores
	}
	return toSerialize, nil
}

type NullableV0037NodeAllocation struct {
	value *V0037NodeAllocation
	isSet bool
}

func (v NullableV0037NodeAllocation) Get() *V0037NodeAllocation {
	return v.value
}

func (v *NullableV0037NodeAllocation) Set(val *V0037NodeAllocation) {
	v.value = val
	v.isSet = true
}

func (v NullableV0037NodeAllocation) IsSet() bool {
	return v.isSet
}

func (v *NullableV0037NodeAllocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0037NodeAllocation(val *V0037NodeAllocation) *NullableV0037NodeAllocation {
	return &NullableV0037NodeAllocation{value: val, isSet: true}
}

func (v NullableV0037NodeAllocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0037NodeAllocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

