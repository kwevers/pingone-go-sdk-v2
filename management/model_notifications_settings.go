/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"
)

// NotificationsSettings struct for NotificationsSettings
type NotificationsSettings struct {
	// The time the resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Restrictions *NotificationsSettingsRestrictions `json:"restrictions,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// A list which represents the execution order of different SMS/Voice providers configured for the environment. The providers and their accounts’ configurations are represented in the list by the ID of the corresponding `PhoneDeliverySettings` resource. The only provider which is not represented by a `phoneDeliverySettingsID` is the PingOne Twilio provider. The PingOne Twilio provider is represented by the `PINGONE_TWILIO` string. If the `smsProvidersFallbackChain` list is empty, an SMS or voice message will be sent using the default Ping Twilio account. Otherwise, an SMS or voice message will be sent using the first provider in the list. If the server fails to queue the message using that provider, it will use the next provider in the list to try to send the message. This process will go on until there are no more providers in the list. If the server failed to send the message using all providers, the notification status is set to `FAILED`.
	SmsProvidersFallbackChain []string `json:"smsProvidersFallbackChain,omitempty"`
	From *NotificationsSettingsFrom `json:"from,omitempty"`
	ReplyTo *NotificationsSettingsReplyTo `json:"replyTo,omitempty"`
}

// NewNotificationsSettings instantiates a new NotificationsSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettings() *NotificationsSettings {
	this := NotificationsSettings{}
	return &this
}

// NewNotificationsSettingsWithDefaults instantiates a new NotificationsSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsWithDefaults() *NotificationsSettings {
	this := NotificationsSettings{}
	return &this
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NotificationsSettings) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NotificationsSettings) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NotificationsSettings) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *NotificationsSettings) GetRestrictions() NotificationsSettingsRestrictions {
	if o == nil || isNil(o.Restrictions) {
		var ret NotificationsSettingsRestrictions
		return ret
	}
	return *o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetRestrictionsOk() (*NotificationsSettingsRestrictions, bool) {
	if o == nil || isNil(o.Restrictions) {
    return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *NotificationsSettings) HasRestrictions() bool {
	if o != nil && !isNil(o.Restrictions) {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given NotificationsSettingsRestrictions and assigns it to the Restrictions field.
func (o *NotificationsSettings) SetRestrictions(v NotificationsSettingsRestrictions) {
	o.Restrictions = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NotificationsSettings) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NotificationsSettings) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NotificationsSettings) SetId(v string) {
	o.Id = &v
}

// GetSmsProvidersFallbackChain returns the SmsProvidersFallbackChain field value if set, zero value otherwise.
func (o *NotificationsSettings) GetSmsProvidersFallbackChain() []string {
	if o == nil || isNil(o.SmsProvidersFallbackChain) {
		var ret []string
		return ret
	}
	return o.SmsProvidersFallbackChain
}

// GetSmsProvidersFallbackChainOk returns a tuple with the SmsProvidersFallbackChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetSmsProvidersFallbackChainOk() ([]string, bool) {
	if o == nil || isNil(o.SmsProvidersFallbackChain) {
    return nil, false
	}
	return o.SmsProvidersFallbackChain, true
}

// HasSmsProvidersFallbackChain returns a boolean if a field has been set.
func (o *NotificationsSettings) HasSmsProvidersFallbackChain() bool {
	if o != nil && !isNil(o.SmsProvidersFallbackChain) {
		return true
	}

	return false
}

// SetSmsProvidersFallbackChain gets a reference to the given []string and assigns it to the SmsProvidersFallbackChain field.
func (o *NotificationsSettings) SetSmsProvidersFallbackChain(v []string) {
	o.SmsProvidersFallbackChain = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *NotificationsSettings) GetFrom() NotificationsSettingsFrom {
	if o == nil || isNil(o.From) {
		var ret NotificationsSettingsFrom
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetFromOk() (*NotificationsSettingsFrom, bool) {
	if o == nil || isNil(o.From) {
    return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *NotificationsSettings) HasFrom() bool {
	if o != nil && !isNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given NotificationsSettingsFrom and assigns it to the From field.
func (o *NotificationsSettings) SetFrom(v NotificationsSettingsFrom) {
	o.From = &v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *NotificationsSettings) GetReplyTo() NotificationsSettingsReplyTo {
	if o == nil || isNil(o.ReplyTo) {
		var ret NotificationsSettingsReplyTo
		return ret
	}
	return *o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetReplyToOk() (*NotificationsSettingsReplyTo, bool) {
	if o == nil || isNil(o.ReplyTo) {
    return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *NotificationsSettings) HasReplyTo() bool {
	if o != nil && !isNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given NotificationsSettingsReplyTo and assigns it to the ReplyTo field.
func (o *NotificationsSettings) SetReplyTo(v NotificationsSettingsReplyTo) {
	o.ReplyTo = &v
}

func (o NotificationsSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.Restrictions) {
		toSerialize["restrictions"] = o.Restrictions
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.SmsProvidersFallbackChain) {
		toSerialize["smsProvidersFallbackChain"] = o.SmsProvidersFallbackChain
	}
	if !isNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !isNil(o.ReplyTo) {
		toSerialize["replyTo"] = o.ReplyTo
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationsSettings struct {
	value *NotificationsSettings
	isSet bool
}

func (v NullableNotificationsSettings) Get() *NotificationsSettings {
	return v.value
}

func (v *NullableNotificationsSettings) Set(val *NotificationsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettings(val *NotificationsSettings) *NullableNotificationsSettings {
	return &NullableNotificationsSettings{value: val, isSet: true}
}

func (v NullableNotificationsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

