/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// UserAddress struct for UserAddress
type UserAddress struct {
	// A string that specifies the country name component. When specified, the value must be in ISO 3166-1 `alpha-2` code format [ISO3166]. For example, the country codes for the United States and Sweden are `US` and \"SE\", respectively. Valid characters consist of two upper-case letters (regex `[A-Z]{2}`).
	CountryCode *string `json:"countryCode,omitempty"`
	// A string that specifies the city or locality component of the address. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex `^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$`). It can have a length of no more than 256 characters (min/max=1/256).
	Locality *string `json:"locality,omitempty"`
	// A string that specifies the zip code or postal code component of the address. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex `^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$`). It can have a length of no more than 40 characters (min/max=1/40).
	PostalCode *string `json:"postalCode,omitempty"`
	// A string that specifies the state or region component of the address. The string can contain any letters, numbers, combining characters, math and currency symbols, dingbats and drawing characters, and invisible whitespace (regex `^[\\p{L}\\p{M}\\p{Zs}\\p{S}\\p{N}\\p{P}]*$`). It can have a length of no more than 256 characters (min/max=1/256).
	Region *string `json:"region,omitempty"`
	// A string that specifies the full street address component, which may include house number, street name, P.O. box, and multi-line extended street address information. This attribute may contain newlines (regex `^[\\p{L}\\p{M}\\p{N}\\p{Zs}\\p{P}\\n\\r]*$`). It can have a length of no more than 256 characters (min/max=1/256).
	StreetAddress *string `json:"streetAddress,omitempty"`
}

// NewUserAddress instantiates a new UserAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAddress() *UserAddress {
	this := UserAddress{}
	return &this
}

// NewUserAddressWithDefaults instantiates a new UserAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAddressWithDefaults() *UserAddress {
	this := UserAddress{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *UserAddress) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UserAddress) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *UserAddress) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *UserAddress) GetLocality() string {
	if o == nil || o.Locality == nil {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAddress) GetLocalityOk() (*string, bool) {
	if o == nil || o.Locality == nil {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *UserAddress) HasLocality() bool {
	if o != nil && o.Locality != nil {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *UserAddress) SetLocality(v string) {
	o.Locality = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *UserAddress) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *UserAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *UserAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *UserAddress) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAddress) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *UserAddress) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *UserAddress) SetRegion(v string) {
	o.Region = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *UserAddress) GetStreetAddress() string {
	if o == nil || o.StreetAddress == nil {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAddress) GetStreetAddressOk() (*string, bool) {
	if o == nil || o.StreetAddress == nil {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *UserAddress) HasStreetAddress() bool {
	if o != nil && o.StreetAddress != nil {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *UserAddress) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

func (o UserAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.Locality != nil {
		toSerialize["locality"] = o.Locality
	}
	if o.PostalCode != nil {
		toSerialize["postalCode"] = o.PostalCode
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.StreetAddress != nil {
		toSerialize["streetAddress"] = o.StreetAddress
	}
	return json.Marshal(toSerialize)
}

type NullableUserAddress struct {
	value *UserAddress
	isSet bool
}

func (v NullableUserAddress) Get() *UserAddress {
	return v.value
}

func (v *NullableUserAddress) Set(val *UserAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAddress(val *UserAddress) *NullableUserAddress {
	return &NullableUserAddress{value: val, isSet: true}
}

func (v NullableUserAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


