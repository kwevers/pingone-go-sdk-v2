/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// SignOnPolicyActionMFA struct for SignOnPolicyActionMFA
type SignOnPolicyActionMFA struct {
	Links map[string]interface{} `json:"_links,omitempty"`
	Condition *SignOnPolicyActionCommonConditionOrOrInner `json:"condition,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the sign-on policy assignment resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// An integer that specifies the order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first. This is a required property.
	Priority int32 `json:"priority"`
	SignOnPolicy *SignOnPolicyActionCommonSignOnPolicy `json:"signOnPolicy,omitempty"`
	Type EnumSignOnPolicyType `json:"type"`
	// Deprecated
	Authenticator *SignOnPolicyActionMFAAllOfAuthenticator `json:"authenticator,omitempty"`
	// Deprecated
	BoundBiometrics *SignOnPolicyActionMFAAllOfBoundBiometrics `json:"boundBiometrics,omitempty"`
	// Deprecated
	Email *SignOnPolicyActionMFAAllOfEmail `json:"email,omitempty"`
	// Deprecated
	SecurityKey *SignOnPolicyActionMFAAllOfSecurityKey `json:"securityKey,omitempty"`
	// Deprecated
	Sms *SignOnPolicyActionMFAAllOfSms `json:"sms,omitempty"`
	// Deprecated
	Voice *SignOnPolicyActionMFAAllOfVoice `json:"voice,omitempty"`
	// The applications collection specifies all the native native applications that are allowed in the sign-on policy action.  If the applications collection is empty, a push notification is not allowed for the action.
	// Deprecated
	Applications []SignOnPolicyActionMFAAllOfApplications `json:"applications,omitempty"`
	DeviceAuthenticationPolicy *SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy `json:"deviceAuthenticationPolicy,omitempty"`
	NoDevicesMode *EnumSignOnPolicyNoDeviceMode `json:"noDevicesMode,omitempty"`
}

// NewSignOnPolicyActionMFA instantiates a new SignOnPolicyActionMFA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionMFA(priority int32, type_ EnumSignOnPolicyType) *SignOnPolicyActionMFA {
	this := SignOnPolicyActionMFA{}
	this.Priority = priority
	this.Type = type_
	return &this
}

// NewSignOnPolicyActionMFAWithDefaults instantiates a new SignOnPolicyActionMFA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionMFAWithDefaults() *SignOnPolicyActionMFA {
	this := SignOnPolicyActionMFA{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFA) GetLinks() map[string]interface{} {
	if o == nil || o.Links == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFA) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFA) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]interface{} and assigns it to the Links field.
func (o *SignOnPolicyActionMFA) SetLinks(v map[string]interface{}) {
	o.Links = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFA) GetCondition() SignOnPolicyActionCommonConditionOrOrInner {
	if o == nil || o.Condition == nil {
		var ret SignOnPolicyActionCommonConditionOrOrInner
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFA) GetConditionOk() (*SignOnPolicyActionCommonConditionOrOrInner, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFA) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given SignOnPolicyActionCommonConditionOrOrInner and assigns it to the Condition field.
func (o *SignOnPolicyActionMFA) SetCondition(v SignOnPolicyActionCommonConditionOrOrInner) {
	o.Condition = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFA) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFA) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFA) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *SignOnPolicyActionMFA) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFA) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFA) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFA) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SignOnPolicyActionMFA) SetId(v string) {
	o.Id = &v
}

// GetPriority returns the Priority field value
func (o *SignOnPolicyActionMFA) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFA) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *SignOnPolicyActionMFA) SetPriority(v int32) {
	o.Priority = v
}

// GetSignOnPolicy returns the SignOnPolicy field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFA) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy {
	if o == nil || o.SignOnPolicy == nil {
		var ret SignOnPolicyActionCommonSignOnPolicy
		return ret
	}
	return *o.SignOnPolicy
}

// GetSignOnPolicyOk returns a tuple with the SignOnPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFA) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool) {
	if o == nil || o.SignOnPolicy == nil {
		return nil, false
	}
	return o.SignOnPolicy, true
}

// HasSignOnPolicy returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFA) HasSignOnPolicy() bool {
	if o != nil && o.SignOnPolicy != nil {
		return true
	}

	return false
}

// SetSignOnPolicy gets a reference to the given SignOnPolicyActionCommonSignOnPolicy and assigns it to the SignOnPolicy field.
func (o *SignOnPolicyActionMFA) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy) {
	o.SignOnPolicy = &v
}

// GetType returns the Type field value
func (o *SignOnPolicyActionMFA) GetType() EnumSignOnPolicyType {
	if o == nil {
		var ret EnumSignOnPolicyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFA) GetTypeOk() (*EnumSignOnPolicyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SignOnPolicyActionMFA) SetType(v EnumSignOnPolicyType) {
	o.Type = v
}

// GetAuthenticator returns the Authenticator field value if set, zero value otherwise.
// Deprecated
func (o *SignOnPolicyActionMFA) GetAuthenticator() SignOnPolicyActionMFAAllOfAuthenticator {
	if o == nil || o.Authenticator == nil {
		var ret SignOnPolicyActionMFAAllOfAuthenticator
		return ret
	}
	return *o.Authenticator
}

// GetAuthenticatorOk returns a tuple with the Authenticator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SignOnPolicyActionMFA) GetAuthenticatorOk() (*SignOnPolicyActionMFAAllOfAuthenticator, bool) {
	if o == nil || o.Authenticator == nil {
		return nil, false
	}
	return o.Authenticator, true
}

// HasAuthenticator returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFA) HasAuthenticator() bool {
	if o != nil && o.Authenticator != nil {
		return true
	}

	return false
}

// SetAuthenticator gets a reference to the given SignOnPolicyActionMFAAllOfAuthenticator and assigns it to the Authenticator field.
// Deprecated
func (o *SignOnPolicyActionMFA) SetAuthenticator(v SignOnPolicyActionMFAAllOfAuthenticator) {
	o.Authenticator = &v
}

// GetBoundBiometrics returns the BoundBiometrics field value if set, zero value otherwise.
// Deprecated
func (o *SignOnPolicyActionMFA) GetBoundBiometrics() SignOnPolicyActionMFAAllOfBoundBiometrics {
	if o == nil || o.BoundBiometrics == nil {
		var ret SignOnPolicyActionMFAAllOfBoundBiometrics
		return ret
	}
	return *o.BoundBiometrics
}

// GetBoundBiometricsOk returns a tuple with the BoundBiometrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SignOnPolicyActionMFA) GetBoundBiometricsOk() (*SignOnPolicyActionMFAAllOfBoundBiometrics, bool) {
	if o == nil || o.BoundBiometrics == nil {
		return nil, false
	}
	return o.BoundBiometrics, true
}

// HasBoundBiometrics returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFA) HasBoundBiometrics() bool {
	if o != nil && o.BoundBiometrics != nil {
		return true
	}

	return false
}

// SetBoundBiometrics gets a reference to the given SignOnPolicyActionMFAAllOfBoundBiometrics and assigns it to the BoundBiometrics field.
// Deprecated
func (o *SignOnPolicyActionMFA) SetBoundBiometrics(v SignOnPolicyActionMFAAllOfBoundBiometrics) {
	o.BoundBiometrics = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
// Deprecated
func (o *SignOnPolicyActionMFA) GetEmail() SignOnPolicyActionMFAAllOfEmail {
	if o == nil || o.Email == nil {
		var ret SignOnPolicyActionMFAAllOfEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SignOnPolicyActionMFA) GetEmailOk() (*SignOnPolicyActionMFAAllOfEmail, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFA) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given SignOnPolicyActionMFAAllOfEmail and assigns it to the Email field.
// Deprecated
func (o *SignOnPolicyActionMFA) SetEmail(v SignOnPolicyActionMFAAllOfEmail) {
	o.Email = &v
}

// GetSecurityKey returns the SecurityKey field value if set, zero value otherwise.
// Deprecated
func (o *SignOnPolicyActionMFA) GetSecurityKey() SignOnPolicyActionMFAAllOfSecurityKey {
	if o == nil || o.SecurityKey == nil {
		var ret SignOnPolicyActionMFAAllOfSecurityKey
		return ret
	}
	return *o.SecurityKey
}

// GetSecurityKeyOk returns a tuple with the SecurityKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SignOnPolicyActionMFA) GetSecurityKeyOk() (*SignOnPolicyActionMFAAllOfSecurityKey, bool) {
	if o == nil || o.SecurityKey == nil {
		return nil, false
	}
	return o.SecurityKey, true
}

// HasSecurityKey returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFA) HasSecurityKey() bool {
	if o != nil && o.SecurityKey != nil {
		return true
	}

	return false
}

// SetSecurityKey gets a reference to the given SignOnPolicyActionMFAAllOfSecurityKey and assigns it to the SecurityKey field.
// Deprecated
func (o *SignOnPolicyActionMFA) SetSecurityKey(v SignOnPolicyActionMFAAllOfSecurityKey) {
	o.SecurityKey = &v
}

// GetSms returns the Sms field value if set, zero value otherwise.
// Deprecated
func (o *SignOnPolicyActionMFA) GetSms() SignOnPolicyActionMFAAllOfSms {
	if o == nil || o.Sms == nil {
		var ret SignOnPolicyActionMFAAllOfSms
		return ret
	}
	return *o.Sms
}

// GetSmsOk returns a tuple with the Sms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SignOnPolicyActionMFA) GetSmsOk() (*SignOnPolicyActionMFAAllOfSms, bool) {
	if o == nil || o.Sms == nil {
		return nil, false
	}
	return o.Sms, true
}

// HasSms returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFA) HasSms() bool {
	if o != nil && o.Sms != nil {
		return true
	}

	return false
}

// SetSms gets a reference to the given SignOnPolicyActionMFAAllOfSms and assigns it to the Sms field.
// Deprecated
func (o *SignOnPolicyActionMFA) SetSms(v SignOnPolicyActionMFAAllOfSms) {
	o.Sms = &v
}

// GetVoice returns the Voice field value if set, zero value otherwise.
// Deprecated
func (o *SignOnPolicyActionMFA) GetVoice() SignOnPolicyActionMFAAllOfVoice {
	if o == nil || o.Voice == nil {
		var ret SignOnPolicyActionMFAAllOfVoice
		return ret
	}
	return *o.Voice
}

// GetVoiceOk returns a tuple with the Voice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SignOnPolicyActionMFA) GetVoiceOk() (*SignOnPolicyActionMFAAllOfVoice, bool) {
	if o == nil || o.Voice == nil {
		return nil, false
	}
	return o.Voice, true
}

// HasVoice returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFA) HasVoice() bool {
	if o != nil && o.Voice != nil {
		return true
	}

	return false
}

// SetVoice gets a reference to the given SignOnPolicyActionMFAAllOfVoice and assigns it to the Voice field.
// Deprecated
func (o *SignOnPolicyActionMFA) SetVoice(v SignOnPolicyActionMFAAllOfVoice) {
	o.Voice = &v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
// Deprecated
func (o *SignOnPolicyActionMFA) GetApplications() []SignOnPolicyActionMFAAllOfApplications {
	if o == nil || o.Applications == nil {
		var ret []SignOnPolicyActionMFAAllOfApplications
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SignOnPolicyActionMFA) GetApplicationsOk() ([]SignOnPolicyActionMFAAllOfApplications, bool) {
	if o == nil || o.Applications == nil {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFA) HasApplications() bool {
	if o != nil && o.Applications != nil {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []SignOnPolicyActionMFAAllOfApplications and assigns it to the Applications field.
// Deprecated
func (o *SignOnPolicyActionMFA) SetApplications(v []SignOnPolicyActionMFAAllOfApplications) {
	o.Applications = v
}

// GetDeviceAuthenticationPolicy returns the DeviceAuthenticationPolicy field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFA) GetDeviceAuthenticationPolicy() SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy {
	if o == nil || o.DeviceAuthenticationPolicy == nil {
		var ret SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy
		return ret
	}
	return *o.DeviceAuthenticationPolicy
}

// GetDeviceAuthenticationPolicyOk returns a tuple with the DeviceAuthenticationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFA) GetDeviceAuthenticationPolicyOk() (*SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy, bool) {
	if o == nil || o.DeviceAuthenticationPolicy == nil {
		return nil, false
	}
	return o.DeviceAuthenticationPolicy, true
}

// HasDeviceAuthenticationPolicy returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFA) HasDeviceAuthenticationPolicy() bool {
	if o != nil && o.DeviceAuthenticationPolicy != nil {
		return true
	}

	return false
}

// SetDeviceAuthenticationPolicy gets a reference to the given SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy and assigns it to the DeviceAuthenticationPolicy field.
func (o *SignOnPolicyActionMFA) SetDeviceAuthenticationPolicy(v SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy) {
	o.DeviceAuthenticationPolicy = &v
}

// GetNoDevicesMode returns the NoDevicesMode field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFA) GetNoDevicesMode() EnumSignOnPolicyNoDeviceMode {
	if o == nil || o.NoDevicesMode == nil {
		var ret EnumSignOnPolicyNoDeviceMode
		return ret
	}
	return *o.NoDevicesMode
}

// GetNoDevicesModeOk returns a tuple with the NoDevicesMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFA) GetNoDevicesModeOk() (*EnumSignOnPolicyNoDeviceMode, bool) {
	if o == nil || o.NoDevicesMode == nil {
		return nil, false
	}
	return o.NoDevicesMode, true
}

// HasNoDevicesMode returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFA) HasNoDevicesMode() bool {
	if o != nil && o.NoDevicesMode != nil {
		return true
	}

	return false
}

// SetNoDevicesMode gets a reference to the given EnumSignOnPolicyNoDeviceMode and assigns it to the NoDevicesMode field.
func (o *SignOnPolicyActionMFA) SetNoDevicesMode(v EnumSignOnPolicyNoDeviceMode) {
	o.NoDevicesMode = &v
}

func (o SignOnPolicyActionMFA) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["priority"] = o.Priority
	}
	if o.SignOnPolicy != nil {
		toSerialize["signOnPolicy"] = o.SignOnPolicy
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Authenticator != nil {
		toSerialize["authenticator"] = o.Authenticator
	}
	if o.BoundBiometrics != nil {
		toSerialize["boundBiometrics"] = o.BoundBiometrics
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.SecurityKey != nil {
		toSerialize["securityKey"] = o.SecurityKey
	}
	if o.Sms != nil {
		toSerialize["sms"] = o.Sms
	}
	if o.Voice != nil {
		toSerialize["voice"] = o.Voice
	}
	if o.Applications != nil {
		toSerialize["applications"] = o.Applications
	}
	if o.DeviceAuthenticationPolicy != nil {
		toSerialize["deviceAuthenticationPolicy"] = o.DeviceAuthenticationPolicy
	}
	if o.NoDevicesMode != nil {
		toSerialize["noDevicesMode"] = o.NoDevicesMode
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionMFA struct {
	value *SignOnPolicyActionMFA
	isSet bool
}

func (v NullableSignOnPolicyActionMFA) Get() *SignOnPolicyActionMFA {
	return v.value
}

func (v *NullableSignOnPolicyActionMFA) Set(val *SignOnPolicyActionMFA) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionMFA) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionMFA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionMFA(val *SignOnPolicyActionMFA) *NullableSignOnPolicyActionMFA {
	return &NullableSignOnPolicyActionMFA{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionMFA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionMFA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


