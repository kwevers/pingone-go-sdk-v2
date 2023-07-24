/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the EmailDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailDomain{}

// EmailDomain struct for EmailDomain
type EmailDomain struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
	// A string that specifies the auto-generated ID of the email domain.
	Id *string `json:"id,omitempty"`
	// A string that specifies the resource name, which must be provided and must be unique within an environment (for example, auth.shopco.com). This is a required property. Wildcards are NOT supported.
	DomainName string `json:"domainName"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
}

// NewEmailDomain instantiates a new EmailDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDomain(domainName string) *EmailDomain {
	this := EmailDomain{}
	this.DomainName = domainName
	return &this
}

// NewEmailDomainWithDefaults instantiates a new EmailDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDomainWithDefaults() *EmailDomain {
	this := EmailDomain{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EmailDomain) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomain) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EmailDomain) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *EmailDomain) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmailDomain) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomain) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmailDomain) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmailDomain) SetId(v string) {
	o.Id = &v
}

// GetDomainName returns the DomainName field value
func (o *EmailDomain) GetDomainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value
// and a boolean to check if the value has been set.
func (o *EmailDomain) GetDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainName, true
}

// SetDomainName sets field value
func (o *EmailDomain) SetDomainName(v string) {
	o.DomainName = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *EmailDomain) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomain) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *EmailDomain) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *EmailDomain) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

func (o EmailDomain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	// skip: id is readOnly
	toSerialize["domainName"] = o.DomainName
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	return toSerialize, nil
}

type NullableEmailDomain struct {
	value *EmailDomain
	isSet bool
}

func (v NullableEmailDomain) Get() *EmailDomain {
	return v.value
}

func (v *NullableEmailDomain) Set(val *EmailDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDomain(val *EmailDomain) *NullableEmailDomain {
	return &NullableEmailDomain{value: val, isSet: true}
}

func (v NullableEmailDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


