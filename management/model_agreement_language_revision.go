/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"
)

// checks if the AgreementLanguageRevision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgreementLanguageRevision{}

// AgreementLanguageRevision struct for AgreementLanguageRevision
type AgreementLanguageRevision struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
	Agreement *AgreementLanguageAgreement `json:"agreement,omitempty"`
	ContentType EnumAgreementRevisionContentType `json:"contentType"`
	// A date that specifies the start date that the revision is presented to users. This property value can be modified only if the current value is a date that has not already passed. The effective date must be unique for each language agreement, and the property value can be the present date or a future date only.
	EffectiveAt time.Time `json:"effectiveAt"`
	// A read-only string that specifies the revision ID.
	Id *string `json:"id,omitempty"`
	Language *AgreementLanguageRevisionLanguage `json:"language,omitempty"`
	// A date that specifies whether the revision is still valid in the context of all revisions for a language. This property is calculated dynamically at read time, taking into consideration the agreement language, the language enabled property, and the agreement enabled property. When a new revision is added, the notValidAfter property values for all other previous revisions might be impacted. For example, if a new revision becomes effective and it forces reconsent, then all older revisions are no longer valid.
	NotValidAfter *time.Time `json:"notValidAfter,omitempty"`
	// A boolean that specifies whether the user is required to provide consent to the language revision after it becomes effective.
	RequireReconsent bool `json:"requireReconsent"`
	// An immutable string that specifies text or HTML for the revision. This attribute is supported in POST requests only. For more information, see contentType.
	Text string `json:"text"`
}

// NewAgreementLanguageRevision instantiates a new AgreementLanguageRevision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgreementLanguageRevision(contentType EnumAgreementRevisionContentType, effectiveAt time.Time, requireReconsent bool, text string) *AgreementLanguageRevision {
	this := AgreementLanguageRevision{}
	this.ContentType = contentType
	this.EffectiveAt = effectiveAt
	this.RequireReconsent = requireReconsent
	this.Text = text
	return &this
}

// NewAgreementLanguageRevisionWithDefaults instantiates a new AgreementLanguageRevision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgreementLanguageRevisionWithDefaults() *AgreementLanguageRevision {
	this := AgreementLanguageRevision{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AgreementLanguageRevision) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AgreementLanguageRevision) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *AgreementLanguageRevision) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetAgreement returns the Agreement field value if set, zero value otherwise.
func (o *AgreementLanguageRevision) GetAgreement() AgreementLanguageAgreement {
	if o == nil || IsNil(o.Agreement) {
		var ret AgreementLanguageAgreement
		return ret
	}
	return *o.Agreement
}

// GetAgreementOk returns a tuple with the Agreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetAgreementOk() (*AgreementLanguageAgreement, bool) {
	if o == nil || IsNil(o.Agreement) {
		return nil, false
	}
	return o.Agreement, true
}

// HasAgreement returns a boolean if a field has been set.
func (o *AgreementLanguageRevision) HasAgreement() bool {
	if o != nil && !IsNil(o.Agreement) {
		return true
	}

	return false
}

// SetAgreement gets a reference to the given AgreementLanguageAgreement and assigns it to the Agreement field.
func (o *AgreementLanguageRevision) SetAgreement(v AgreementLanguageAgreement) {
	o.Agreement = &v
}

// GetContentType returns the ContentType field value
func (o *AgreementLanguageRevision) GetContentType() EnumAgreementRevisionContentType {
	if o == nil {
		var ret EnumAgreementRevisionContentType
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetContentTypeOk() (*EnumAgreementRevisionContentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *AgreementLanguageRevision) SetContentType(v EnumAgreementRevisionContentType) {
	o.ContentType = v
}

// GetEffectiveAt returns the EffectiveAt field value
func (o *AgreementLanguageRevision) GetEffectiveAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EffectiveAt
}

// GetEffectiveAtOk returns a tuple with the EffectiveAt field value
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetEffectiveAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectiveAt, true
}

// SetEffectiveAt sets field value
func (o *AgreementLanguageRevision) SetEffectiveAt(v time.Time) {
	o.EffectiveAt = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgreementLanguageRevision) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgreementLanguageRevision) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgreementLanguageRevision) SetId(v string) {
	o.Id = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AgreementLanguageRevision) GetLanguage() AgreementLanguageRevisionLanguage {
	if o == nil || IsNil(o.Language) {
		var ret AgreementLanguageRevisionLanguage
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetLanguageOk() (*AgreementLanguageRevisionLanguage, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AgreementLanguageRevision) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given AgreementLanguageRevisionLanguage and assigns it to the Language field.
func (o *AgreementLanguageRevision) SetLanguage(v AgreementLanguageRevisionLanguage) {
	o.Language = &v
}

// GetNotValidAfter returns the NotValidAfter field value if set, zero value otherwise.
func (o *AgreementLanguageRevision) GetNotValidAfter() time.Time {
	if o == nil || IsNil(o.NotValidAfter) {
		var ret time.Time
		return ret
	}
	return *o.NotValidAfter
}

// GetNotValidAfterOk returns a tuple with the NotValidAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetNotValidAfterOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NotValidAfter) {
		return nil, false
	}
	return o.NotValidAfter, true
}

// HasNotValidAfter returns a boolean if a field has been set.
func (o *AgreementLanguageRevision) HasNotValidAfter() bool {
	if o != nil && !IsNil(o.NotValidAfter) {
		return true
	}

	return false
}

// SetNotValidAfter gets a reference to the given time.Time and assigns it to the NotValidAfter field.
func (o *AgreementLanguageRevision) SetNotValidAfter(v time.Time) {
	o.NotValidAfter = &v
}

// GetRequireReconsent returns the RequireReconsent field value
func (o *AgreementLanguageRevision) GetRequireReconsent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireReconsent
}

// GetRequireReconsentOk returns a tuple with the RequireReconsent field value
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetRequireReconsentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireReconsent, true
}

// SetRequireReconsent sets field value
func (o *AgreementLanguageRevision) SetRequireReconsent(v bool) {
	o.RequireReconsent = v
}

// GetText returns the Text field value
func (o *AgreementLanguageRevision) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *AgreementLanguageRevision) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *AgreementLanguageRevision) SetText(v string) {
	o.Text = v
}

func (o AgreementLanguageRevision) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgreementLanguageRevision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Agreement) {
		toSerialize["agreement"] = o.Agreement
	}
	toSerialize["contentType"] = o.ContentType
	toSerialize["effectiveAt"] = o.EffectiveAt
	// skip: id is readOnly
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	// skip: notValidAfter is readOnly
	toSerialize["requireReconsent"] = o.RequireReconsent
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

type NullableAgreementLanguageRevision struct {
	value *AgreementLanguageRevision
	isSet bool
}

func (v NullableAgreementLanguageRevision) Get() *AgreementLanguageRevision {
	return v.value
}

func (v *NullableAgreementLanguageRevision) Set(val *AgreementLanguageRevision) {
	v.value = val
	v.isSet = true
}

func (v NullableAgreementLanguageRevision) IsSet() bool {
	return v.isSet
}

func (v *NullableAgreementLanguageRevision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgreementLanguageRevision(val *AgreementLanguageRevision) *NullableAgreementLanguageRevision {
	return &NullableAgreementLanguageRevision{value: val, isSet: true}
}

func (v NullableAgreementLanguageRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgreementLanguageRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


