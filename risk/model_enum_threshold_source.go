/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
	"fmt"
)

// EnumThresholdSource An enum indicating the source used to calculate the threshold. This can be `MIN_NOT_REACHED`. If the measure is less than `every.minSample`, the threshold isn't calculated. Instead, a value of `LOW` is automatically assigned. If less than five IPs were used by the user during the past hour, `MIN_NOT_REACHED` is set. `CALCULATED`. Indicates the threshold is guaranteed to be calculated. `ENVIRONMENT_FALLBACK`. Indicates a global threshold calculated for the entire environment is used. The global threshold is used when the `ipVelocityByUser.threshold` couldn't be calculated, generally due to a lack of past transactions for the risk predictor to use for the threshold calculation. `DEFAULT_FALLBACK`. Indicates the default threshold defined for the predictor (in `threshold.medium` or `threshold.high`) is used. The default threshold is used when `ENVIRONMENT_FALLBACK` (the global threshold) couldn't be calculated, generally due to a lack of past transactions for the risk predictor to use for the environment threshold calculation.
type EnumThresholdSource string

// List of EnumThresholdSource
const (
	ENUMTHRESHOLDSOURCE_MIN_NOT_REACHED EnumThresholdSource = "MIN_NOT_REACHED"
	ENUMTHRESHOLDSOURCE_CALCULATED EnumThresholdSource = "CALCULATED"
	ENUMTHRESHOLDSOURCE_ENVIRONMENT_FALLBACK EnumThresholdSource = "ENVIRONMENT_FALLBACK"
	ENUMTHRESHOLDSOURCE_DEFAULT_FALLBACK EnumThresholdSource = "DEFAULT_FALLBACK"
)

// All allowed values of EnumThresholdSource enum
var AllowedEnumThresholdSourceEnumValues = []EnumThresholdSource{
	"MIN_NOT_REACHED",
	"CALCULATED",
	"ENVIRONMENT_FALLBACK",
	"DEFAULT_FALLBACK",
}

func (v *EnumThresholdSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumThresholdSource(value)
	for _, existing := range AllowedEnumThresholdSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumThresholdSource", value)
}

// NewEnumThresholdSourceFromValue returns a pointer to a valid EnumThresholdSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumThresholdSourceFromValue(v string) (*EnumThresholdSource, error) {
	ev := EnumThresholdSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumThresholdSource: valid values are %v", v, AllowedEnumThresholdSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumThresholdSource) IsValid() bool {
	for _, existing := range AllowedEnumThresholdSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumThresholdSource value
func (v EnumThresholdSource) Ptr() *EnumThresholdSource {
	return &v
}

type NullableEnumThresholdSource struct {
	value *EnumThresholdSource
	isSet bool
}

func (v NullableEnumThresholdSource) Get() *EnumThresholdSource {
	return v.value
}

func (v *NullableEnumThresholdSource) Set(val *EnumThresholdSource) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumThresholdSource) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumThresholdSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumThresholdSource(val *EnumThresholdSource) *NullableEnumThresholdSource {
	return &NullableEnumThresholdSource{value: val, isSet: true}
}

func (v NullableEnumThresholdSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumThresholdSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

