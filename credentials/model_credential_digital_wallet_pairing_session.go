/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
	"time"
)

// checks if the CredentialDigitalWalletPairingSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialDigitalWalletPairingSession{}

// CredentialDigitalWalletPairingSession struct for CredentialDigitalWalletPairingSession
type CredentialDigitalWalletPairingSession struct {
	Id *string `json:"id,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	User *CredentialDigitalWalletNotificationResultsInnerNotification `json:"user,omitempty"`
	DigitalWallet *CredentialDigitalWalletNotificationResultsInnerNotification `json:"digitalWallet,omitempty"`
	Challenge *string `json:"challenge,omitempty"`
	QrUrl *string `json:"qrUrl,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewCredentialDigitalWalletPairingSession instantiates a new CredentialDigitalWalletPairingSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialDigitalWalletPairingSession() *CredentialDigitalWalletPairingSession {
	this := CredentialDigitalWalletPairingSession{}
	return &this
}

// NewCredentialDigitalWalletPairingSessionWithDefaults instantiates a new CredentialDigitalWalletPairingSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialDigitalWalletPairingSessionWithDefaults() *CredentialDigitalWalletPairingSession {
	this := CredentialDigitalWalletPairingSession{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CredentialDigitalWalletPairingSession) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWalletPairingSession) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CredentialDigitalWalletPairingSession) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CredentialDigitalWalletPairingSession) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CredentialDigitalWalletPairingSession) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWalletPairingSession) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CredentialDigitalWalletPairingSession) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CredentialDigitalWalletPairingSession) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CredentialDigitalWalletPairingSession) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWalletPairingSession) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CredentialDigitalWalletPairingSession) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CredentialDigitalWalletPairingSession) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *CredentialDigitalWalletPairingSession) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWalletPairingSession) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CredentialDigitalWalletPairingSession) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *CredentialDigitalWalletPairingSession) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CredentialDigitalWalletPairingSession) GetUser() CredentialDigitalWalletNotificationResultsInnerNotification {
	if o == nil || IsNil(o.User) {
		var ret CredentialDigitalWalletNotificationResultsInnerNotification
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWalletPairingSession) GetUserOk() (*CredentialDigitalWalletNotificationResultsInnerNotification, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CredentialDigitalWalletPairingSession) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given CredentialDigitalWalletNotificationResultsInnerNotification and assigns it to the User field.
func (o *CredentialDigitalWalletPairingSession) SetUser(v CredentialDigitalWalletNotificationResultsInnerNotification) {
	o.User = &v
}

// GetDigitalWallet returns the DigitalWallet field value if set, zero value otherwise.
func (o *CredentialDigitalWalletPairingSession) GetDigitalWallet() CredentialDigitalWalletNotificationResultsInnerNotification {
	if o == nil || IsNil(o.DigitalWallet) {
		var ret CredentialDigitalWalletNotificationResultsInnerNotification
		return ret
	}
	return *o.DigitalWallet
}

// GetDigitalWalletOk returns a tuple with the DigitalWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWalletPairingSession) GetDigitalWalletOk() (*CredentialDigitalWalletNotificationResultsInnerNotification, bool) {
	if o == nil || IsNil(o.DigitalWallet) {
		return nil, false
	}
	return o.DigitalWallet, true
}

// HasDigitalWallet returns a boolean if a field has been set.
func (o *CredentialDigitalWalletPairingSession) HasDigitalWallet() bool {
	if o != nil && !IsNil(o.DigitalWallet) {
		return true
	}

	return false
}

// SetDigitalWallet gets a reference to the given CredentialDigitalWalletNotificationResultsInnerNotification and assigns it to the DigitalWallet field.
func (o *CredentialDigitalWalletPairingSession) SetDigitalWallet(v CredentialDigitalWalletNotificationResultsInnerNotification) {
	o.DigitalWallet = &v
}

// GetChallenge returns the Challenge field value if set, zero value otherwise.
func (o *CredentialDigitalWalletPairingSession) GetChallenge() string {
	if o == nil || IsNil(o.Challenge) {
		var ret string
		return ret
	}
	return *o.Challenge
}

// GetChallengeOk returns a tuple with the Challenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWalletPairingSession) GetChallengeOk() (*string, bool) {
	if o == nil || IsNil(o.Challenge) {
		return nil, false
	}
	return o.Challenge, true
}

// HasChallenge returns a boolean if a field has been set.
func (o *CredentialDigitalWalletPairingSession) HasChallenge() bool {
	if o != nil && !IsNil(o.Challenge) {
		return true
	}

	return false
}

// SetChallenge gets a reference to the given string and assigns it to the Challenge field.
func (o *CredentialDigitalWalletPairingSession) SetChallenge(v string) {
	o.Challenge = &v
}

// GetQrUrl returns the QrUrl field value if set, zero value otherwise.
func (o *CredentialDigitalWalletPairingSession) GetQrUrl() string {
	if o == nil || IsNil(o.QrUrl) {
		var ret string
		return ret
	}
	return *o.QrUrl
}

// GetQrUrlOk returns a tuple with the QrUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWalletPairingSession) GetQrUrlOk() (*string, bool) {
	if o == nil || IsNil(o.QrUrl) {
		return nil, false
	}
	return o.QrUrl, true
}

// HasQrUrl returns a boolean if a field has been set.
func (o *CredentialDigitalWalletPairingSession) HasQrUrl() bool {
	if o != nil && !IsNil(o.QrUrl) {
		return true
	}

	return false
}

// SetQrUrl gets a reference to the given string and assigns it to the QrUrl field.
func (o *CredentialDigitalWalletPairingSession) SetQrUrl(v string) {
	o.QrUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CredentialDigitalWalletPairingSession) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialDigitalWalletPairingSession) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CredentialDigitalWalletPairingSession) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CredentialDigitalWalletPairingSession) SetStatus(v string) {
	o.Status = &v
}

func (o CredentialDigitalWalletPairingSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialDigitalWalletPairingSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.DigitalWallet) {
		toSerialize["digitalWallet"] = o.DigitalWallet
	}
	if !IsNil(o.Challenge) {
		toSerialize["challenge"] = o.Challenge
	}
	if !IsNil(o.QrUrl) {
		toSerialize["qrUrl"] = o.QrUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableCredentialDigitalWalletPairingSession struct {
	value *CredentialDigitalWalletPairingSession
	isSet bool
}

func (v NullableCredentialDigitalWalletPairingSession) Get() *CredentialDigitalWalletPairingSession {
	return v.value
}

func (v *NullableCredentialDigitalWalletPairingSession) Set(val *CredentialDigitalWalletPairingSession) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialDigitalWalletPairingSession) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialDigitalWalletPairingSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialDigitalWalletPairingSession(val *CredentialDigitalWalletPairingSession) *NullableCredentialDigitalWalletPairingSession {
	return &NullableCredentialDigitalWalletPairingSession{value: val, isSet: true}
}

func (v NullableCredentialDigitalWalletPairingSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialDigitalWalletPairingSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


