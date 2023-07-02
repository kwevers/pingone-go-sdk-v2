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

// checks if the CustomDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDomain{}

// CustomDomain struct for CustomDomain
type CustomDomain struct {
	Certificate *CustomDomainCertificate `json:"certificate,omitempty"`
	// A string that specifies the domain name that should be used as the value of the CNAME record in the customer’s DNS.
	CanonicalName *string `json:"canonicalName,omitempty"`
	// A string that specifies the resource name, which must be provided and must be unique within an environment (for example, auth.shopco.com). This is a required property.
	DomainName string `json:"domainName"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	Status *EnumCustomDomainStatus `json:"status,omitempty"`
}

// NewCustomDomain instantiates a new CustomDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDomain(domainName string) *CustomDomain {
	this := CustomDomain{}
	this.DomainName = domainName
	return &this
}

// NewCustomDomainWithDefaults instantiates a new CustomDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDomainWithDefaults() *CustomDomain {
	this := CustomDomain{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *CustomDomain) GetCertificate() CustomDomainCertificate {
	if o == nil || IsNil(o.Certificate) {
		var ret CustomDomainCertificate
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetCertificateOk() (*CustomDomainCertificate, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *CustomDomain) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given CustomDomainCertificate and assigns it to the Certificate field.
func (o *CustomDomain) SetCertificate(v CustomDomainCertificate) {
	o.Certificate = &v
}

// GetCanonicalName returns the CanonicalName field value if set, zero value otherwise.
func (o *CustomDomain) GetCanonicalName() string {
	if o == nil || IsNil(o.CanonicalName) {
		var ret string
		return ret
	}
	return *o.CanonicalName
}

// GetCanonicalNameOk returns a tuple with the CanonicalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetCanonicalNameOk() (*string, bool) {
	if o == nil || IsNil(o.CanonicalName) {
		return nil, false
	}
	return o.CanonicalName, true
}

// HasCanonicalName returns a boolean if a field has been set.
func (o *CustomDomain) HasCanonicalName() bool {
	if o != nil && !IsNil(o.CanonicalName) {
		return true
	}

	return false
}

// SetCanonicalName gets a reference to the given string and assigns it to the CanonicalName field.
func (o *CustomDomain) SetCanonicalName(v string) {
	o.CanonicalName = &v
}

// GetDomainName returns the DomainName field value
func (o *CustomDomain) GetDomainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetDomainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainName, true
}

// SetDomainName sets field value
func (o *CustomDomain) SetDomainName(v string) {
	o.DomainName = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *CustomDomain) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CustomDomain) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *CustomDomain) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomDomain) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomDomain) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomDomain) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CustomDomain) GetStatus() EnumCustomDomainStatus {
	if o == nil || IsNil(o.Status) {
		var ret EnumCustomDomainStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDomain) GetStatusOk() (*EnumCustomDomainStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomDomain) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EnumCustomDomainStatus and assigns it to the Status field.
func (o *CustomDomain) SetStatus(v EnumCustomDomainStatus) {
	o.Status = &v
}

func (o CustomDomain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	// skip: canonicalName is readOnly
	toSerialize["domainName"] = o.DomainName
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	// skip: id is readOnly
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableCustomDomain struct {
	value *CustomDomain
	isSet bool
}

func (v NullableCustomDomain) Get() *CustomDomain {
	return v.value
}

func (v *NullableCustomDomain) Set(val *CustomDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDomain(val *CustomDomain) *NullableCustomDomain {
	return &NullableCustomDomain{value: val, isSet: true}
}

func (v NullableCustomDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


