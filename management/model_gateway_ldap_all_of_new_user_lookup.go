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

// GatewayLDAPAllOfNewUserLookup The configurations for initially authenticating new users who will be migrated to PingOne. Note If there are multiple users having the same user name, only the first user processed is provisioned.
type GatewayLDAPAllOfNewUserLookup struct {
	// A list of objects supplying a mapping of PingOne attributes to external LDAP attributes. One of the entries must be a mapping for \"username`. This is required for the PingOne user schema.
	AttributeMappings []GatewayLDAPAllOfNewUserLookupAttributeMappings `json:"attributeMappings"`
	// The LDAP user search filter to use to match users against the entered user identifier at login. For example, (((uid=${identifier})(mail=${identifier})). Alternatively, this can be a search against the user directory.
	LdapFilterPattern *string `json:"ldapFilterPattern,omitempty"`
	Population *GatewayLDAPAllOfNewUserLookupPopulation `json:"population,omitempty"`
	// A map of key/value entries used to persist the external LDAP directory attributes.
	OrderedCorrelationAttributes []map[string]interface{} `json:"orderedCorrelationAttributes,omitempty"`
	PasswordAuthority *EnumGatewayPasswordAuthority `json:"passwordAuthority,omitempty"`
	// The LDAP base domain name (DN) for this user type.
	SearchBaseDn *string `json:"searchBaseDn,omitempty"`
}

// NewGatewayLDAPAllOfNewUserLookup instantiates a new GatewayLDAPAllOfNewUserLookup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayLDAPAllOfNewUserLookup(attributeMappings []GatewayLDAPAllOfNewUserLookupAttributeMappings) *GatewayLDAPAllOfNewUserLookup {
	this := GatewayLDAPAllOfNewUserLookup{}
	this.AttributeMappings = attributeMappings
	return &this
}

// NewGatewayLDAPAllOfNewUserLookupWithDefaults instantiates a new GatewayLDAPAllOfNewUserLookup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayLDAPAllOfNewUserLookupWithDefaults() *GatewayLDAPAllOfNewUserLookup {
	this := GatewayLDAPAllOfNewUserLookup{}
	return &this
}

// GetAttributeMappings returns the AttributeMappings field value
func (o *GatewayLDAPAllOfNewUserLookup) GetAttributeMappings() []GatewayLDAPAllOfNewUserLookupAttributeMappings {
	if o == nil {
		var ret []GatewayLDAPAllOfNewUserLookupAttributeMappings
		return ret
	}

	return o.AttributeMappings
}

// GetAttributeMappingsOk returns a tuple with the AttributeMappings field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfNewUserLookup) GetAttributeMappingsOk() ([]GatewayLDAPAllOfNewUserLookupAttributeMappings, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributeMappings, true
}

// SetAttributeMappings sets field value
func (o *GatewayLDAPAllOfNewUserLookup) SetAttributeMappings(v []GatewayLDAPAllOfNewUserLookupAttributeMappings) {
	o.AttributeMappings = v
}

// GetLdapFilterPattern returns the LdapFilterPattern field value if set, zero value otherwise.
func (o *GatewayLDAPAllOfNewUserLookup) GetLdapFilterPattern() string {
	if o == nil || o.LdapFilterPattern == nil {
		var ret string
		return ret
	}
	return *o.LdapFilterPattern
}

// GetLdapFilterPatternOk returns a tuple with the LdapFilterPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfNewUserLookup) GetLdapFilterPatternOk() (*string, bool) {
	if o == nil || o.LdapFilterPattern == nil {
		return nil, false
	}
	return o.LdapFilterPattern, true
}

// HasLdapFilterPattern returns a boolean if a field has been set.
func (o *GatewayLDAPAllOfNewUserLookup) HasLdapFilterPattern() bool {
	if o != nil && o.LdapFilterPattern != nil {
		return true
	}

	return false
}

// SetLdapFilterPattern gets a reference to the given string and assigns it to the LdapFilterPattern field.
func (o *GatewayLDAPAllOfNewUserLookup) SetLdapFilterPattern(v string) {
	o.LdapFilterPattern = &v
}

// GetPopulation returns the Population field value if set, zero value otherwise.
func (o *GatewayLDAPAllOfNewUserLookup) GetPopulation() GatewayLDAPAllOfNewUserLookupPopulation {
	if o == nil || o.Population == nil {
		var ret GatewayLDAPAllOfNewUserLookupPopulation
		return ret
	}
	return *o.Population
}

// GetPopulationOk returns a tuple with the Population field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfNewUserLookup) GetPopulationOk() (*GatewayLDAPAllOfNewUserLookupPopulation, bool) {
	if o == nil || o.Population == nil {
		return nil, false
	}
	return o.Population, true
}

// HasPopulation returns a boolean if a field has been set.
func (o *GatewayLDAPAllOfNewUserLookup) HasPopulation() bool {
	if o != nil && o.Population != nil {
		return true
	}

	return false
}

// SetPopulation gets a reference to the given GatewayLDAPAllOfNewUserLookupPopulation and assigns it to the Population field.
func (o *GatewayLDAPAllOfNewUserLookup) SetPopulation(v GatewayLDAPAllOfNewUserLookupPopulation) {
	o.Population = &v
}

// GetOrderedCorrelationAttributes returns the OrderedCorrelationAttributes field value if set, zero value otherwise.
func (o *GatewayLDAPAllOfNewUserLookup) GetOrderedCorrelationAttributes() []map[string]interface{} {
	if o == nil || o.OrderedCorrelationAttributes == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.OrderedCorrelationAttributes
}

// GetOrderedCorrelationAttributesOk returns a tuple with the OrderedCorrelationAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfNewUserLookup) GetOrderedCorrelationAttributesOk() ([]map[string]interface{}, bool) {
	if o == nil || o.OrderedCorrelationAttributes == nil {
		return nil, false
	}
	return o.OrderedCorrelationAttributes, true
}

// HasOrderedCorrelationAttributes returns a boolean if a field has been set.
func (o *GatewayLDAPAllOfNewUserLookup) HasOrderedCorrelationAttributes() bool {
	if o != nil && o.OrderedCorrelationAttributes != nil {
		return true
	}

	return false
}

// SetOrderedCorrelationAttributes gets a reference to the given []map[string]interface{} and assigns it to the OrderedCorrelationAttributes field.
func (o *GatewayLDAPAllOfNewUserLookup) SetOrderedCorrelationAttributes(v []map[string]interface{}) {
	o.OrderedCorrelationAttributes = v
}

// GetPasswordAuthority returns the PasswordAuthority field value if set, zero value otherwise.
func (o *GatewayLDAPAllOfNewUserLookup) GetPasswordAuthority() EnumGatewayPasswordAuthority {
	if o == nil || o.PasswordAuthority == nil {
		var ret EnumGatewayPasswordAuthority
		return ret
	}
	return *o.PasswordAuthority
}

// GetPasswordAuthorityOk returns a tuple with the PasswordAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfNewUserLookup) GetPasswordAuthorityOk() (*EnumGatewayPasswordAuthority, bool) {
	if o == nil || o.PasswordAuthority == nil {
		return nil, false
	}
	return o.PasswordAuthority, true
}

// HasPasswordAuthority returns a boolean if a field has been set.
func (o *GatewayLDAPAllOfNewUserLookup) HasPasswordAuthority() bool {
	if o != nil && o.PasswordAuthority != nil {
		return true
	}

	return false
}

// SetPasswordAuthority gets a reference to the given EnumGatewayPasswordAuthority and assigns it to the PasswordAuthority field.
func (o *GatewayLDAPAllOfNewUserLookup) SetPasswordAuthority(v EnumGatewayPasswordAuthority) {
	o.PasswordAuthority = &v
}

// GetSearchBaseDn returns the SearchBaseDn field value if set, zero value otherwise.
func (o *GatewayLDAPAllOfNewUserLookup) GetSearchBaseDn() string {
	if o == nil || o.SearchBaseDn == nil {
		var ret string
		return ret
	}
	return *o.SearchBaseDn
}

// GetSearchBaseDnOk returns a tuple with the SearchBaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAPAllOfNewUserLookup) GetSearchBaseDnOk() (*string, bool) {
	if o == nil || o.SearchBaseDn == nil {
		return nil, false
	}
	return o.SearchBaseDn, true
}

// HasSearchBaseDn returns a boolean if a field has been set.
func (o *GatewayLDAPAllOfNewUserLookup) HasSearchBaseDn() bool {
	if o != nil && o.SearchBaseDn != nil {
		return true
	}

	return false
}

// SetSearchBaseDn gets a reference to the given string and assigns it to the SearchBaseDn field.
func (o *GatewayLDAPAllOfNewUserLookup) SetSearchBaseDn(v string) {
	o.SearchBaseDn = &v
}

func (o GatewayLDAPAllOfNewUserLookup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["attributeMappings"] = o.AttributeMappings
	}
	if o.LdapFilterPattern != nil {
		toSerialize["ldapFilterPattern"] = o.LdapFilterPattern
	}
	if o.Population != nil {
		toSerialize["population"] = o.Population
	}
	if o.OrderedCorrelationAttributes != nil {
		toSerialize["orderedCorrelationAttributes"] = o.OrderedCorrelationAttributes
	}
	if o.PasswordAuthority != nil {
		toSerialize["passwordAuthority"] = o.PasswordAuthority
	}
	if o.SearchBaseDn != nil {
		toSerialize["searchBaseDn"] = o.SearchBaseDn
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayLDAPAllOfNewUserLookup struct {
	value *GatewayLDAPAllOfNewUserLookup
	isSet bool
}

func (v NullableGatewayLDAPAllOfNewUserLookup) Get() *GatewayLDAPAllOfNewUserLookup {
	return v.value
}

func (v *NullableGatewayLDAPAllOfNewUserLookup) Set(val *GatewayLDAPAllOfNewUserLookup) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayLDAPAllOfNewUserLookup) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayLDAPAllOfNewUserLookup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayLDAPAllOfNewUserLookup(val *GatewayLDAPAllOfNewUserLookup) *NullableGatewayLDAPAllOfNewUserLookup {
	return &NullableGatewayLDAPAllOfNewUserLookup{value: val, isSet: true}
}

func (v NullableGatewayLDAPAllOfNewUserLookup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayLDAPAllOfNewUserLookup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


