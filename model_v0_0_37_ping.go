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

// checks if the V0037Ping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0037Ping{}

// V0037Ping struct for V0037Ping
type V0037Ping struct {
	// slurm controller hostname
	Hostname *string `json:"hostname,omitempty"`
	// slurm controller host up
	Ping *string `json:"ping,omitempty"`
	// slurm controller mode
	Mode *string `json:"mode,omitempty"`
	// slurm controller status
	Status *int32 `json:"status,omitempty"`
}

// NewV0037Ping instantiates a new V0037Ping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0037Ping() *V0037Ping {
	this := V0037Ping{}
	return &this
}

// NewV0037PingWithDefaults instantiates a new V0037Ping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0037PingWithDefaults() *V0037Ping {
	this := V0037Ping{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *V0037Ping) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Ping) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *V0037Ping) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *V0037Ping) SetHostname(v string) {
	o.Hostname = &v
}

// GetPing returns the Ping field value if set, zero value otherwise.
func (o *V0037Ping) GetPing() string {
	if o == nil || IsNil(o.Ping) {
		var ret string
		return ret
	}
	return *o.Ping
}

// GetPingOk returns a tuple with the Ping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Ping) GetPingOk() (*string, bool) {
	if o == nil || IsNil(o.Ping) {
		return nil, false
	}
	return o.Ping, true
}

// HasPing returns a boolean if a field has been set.
func (o *V0037Ping) HasPing() bool {
	if o != nil && !IsNil(o.Ping) {
		return true
	}

	return false
}

// SetPing gets a reference to the given string and assigns it to the Ping field.
func (o *V0037Ping) SetPing(v string) {
	o.Ping = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *V0037Ping) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Ping) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *V0037Ping) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *V0037Ping) SetMode(v string) {
	o.Mode = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V0037Ping) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0037Ping) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V0037Ping) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *V0037Ping) SetStatus(v int32) {
	o.Status = &v
}

func (o V0037Ping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0037Ping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Ping) {
		toSerialize["ping"] = o.Ping
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableV0037Ping struct {
	value *V0037Ping
	isSet bool
}

func (v NullableV0037Ping) Get() *V0037Ping {
	return v.value
}

func (v *NullableV0037Ping) Set(val *V0037Ping) {
	v.value = val
	v.isSet = true
}

func (v NullableV0037Ping) IsSet() bool {
	return v.isSet
}

func (v *NullableV0037Ping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0037Ping(val *V0037Ping) *NullableV0037Ping {
	return &NullableV0037Ping{value: val, isSet: true}
}

func (v NullableV0037Ping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0037Ping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


