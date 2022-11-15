/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
)

// RiskEvaluationDetailsPreviousSuccessfulTransaction struct for RiskEvaluationDetailsPreviousSuccessfulTransaction
type RiskEvaluationDetailsPreviousSuccessfulTransaction struct {
	// A boolean that specifies whether an anonymous network was detected. Information is available twenty-four hours after the last successful transaction.
	AnonymousNetworkDetected *bool `json:"anonymousNetworkDetected,omitempty"`
	// A string that specifies the country of the IP address involved in the transaction. Information is available twenty-four hours after the last successful transaction.
	Country *string `json:"country,omitempty"`
	// A string that specifies the state of the IP address involved in the transaction. Information is available twenty-four hours after the last successful transaction.
	State *string `json:"state,omitempty"`
	// A string that specifies the city of the IP address involved in the transaction. Information is available twenty-four hours after the last successful transaction.
	City *string `json:"city,omitempty"`
	// A string that specifies the IP address involved in the transaction. Information is available twenty-four hours after the last successful transaction.
	Ip *string `json:"ip,omitempty"`
	// A date that specifies the timestamp of the transaction. Information is available twenty-four hours after the last successful transaction.
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewRiskEvaluationDetailsPreviousSuccessfulTransaction instantiates a new RiskEvaluationDetailsPreviousSuccessfulTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationDetailsPreviousSuccessfulTransaction() *RiskEvaluationDetailsPreviousSuccessfulTransaction {
	this := RiskEvaluationDetailsPreviousSuccessfulTransaction{}
	return &this
}

// NewRiskEvaluationDetailsPreviousSuccessfulTransactionWithDefaults instantiates a new RiskEvaluationDetailsPreviousSuccessfulTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationDetailsPreviousSuccessfulTransactionWithDefaults() *RiskEvaluationDetailsPreviousSuccessfulTransaction {
	this := RiskEvaluationDetailsPreviousSuccessfulTransaction{}
	return &this
}

// GetAnonymousNetworkDetected returns the AnonymousNetworkDetected field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) GetAnonymousNetworkDetected() bool {
	if o == nil || isNil(o.AnonymousNetworkDetected) {
		var ret bool
		return ret
	}
	return *o.AnonymousNetworkDetected
}

// GetAnonymousNetworkDetectedOk returns a tuple with the AnonymousNetworkDetected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) GetAnonymousNetworkDetectedOk() (*bool, bool) {
	if o == nil || isNil(o.AnonymousNetworkDetected) {
    return nil, false
	}
	return o.AnonymousNetworkDetected, true
}

// HasAnonymousNetworkDetected returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) HasAnonymousNetworkDetected() bool {
	if o != nil && !isNil(o.AnonymousNetworkDetected) {
		return true
	}

	return false
}

// SetAnonymousNetworkDetected gets a reference to the given bool and assigns it to the AnonymousNetworkDetected field.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) SetAnonymousNetworkDetected(v bool) {
	o.AnonymousNetworkDetected = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) SetCountry(v string) {
	o.Country = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) SetState(v string) {
	o.State = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) GetCity() string {
	if o == nil || isNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) GetCityOk() (*string, bool) {
	if o == nil || isNil(o.City) {
    return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) HasCity() bool {
	if o != nil && !isNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) SetCity(v string) {
	o.City = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) SetIp(v string) {
	o.Ip = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) GetTimestamp() string {
	if o == nil || isNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) GetTimestampOk() (*string, bool) {
	if o == nil || isNil(o.Timestamp) {
    return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *RiskEvaluationDetailsPreviousSuccessfulTransaction) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o RiskEvaluationDetailsPreviousSuccessfulTransaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AnonymousNetworkDetected) {
		toSerialize["anonymousNetworkDetected"] = o.AnonymousNetworkDetected
	}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableRiskEvaluationDetailsPreviousSuccessfulTransaction struct {
	value *RiskEvaluationDetailsPreviousSuccessfulTransaction
	isSet bool
}

func (v NullableRiskEvaluationDetailsPreviousSuccessfulTransaction) Get() *RiskEvaluationDetailsPreviousSuccessfulTransaction {
	return v.value
}

func (v *NullableRiskEvaluationDetailsPreviousSuccessfulTransaction) Set(val *RiskEvaluationDetailsPreviousSuccessfulTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationDetailsPreviousSuccessfulTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationDetailsPreviousSuccessfulTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationDetailsPreviousSuccessfulTransaction(val *RiskEvaluationDetailsPreviousSuccessfulTransaction) *NullableRiskEvaluationDetailsPreviousSuccessfulTransaction {
	return &NullableRiskEvaluationDetailsPreviousSuccessfulTransaction{value: val, isSet: true}
}

func (v NullableRiskEvaluationDetailsPreviousSuccessfulTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationDetailsPreviousSuccessfulTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


