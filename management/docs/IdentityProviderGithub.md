# IdentityProviderGithub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the IdP. | [optional] 
**Enabled** | [**EnumEnabledStatus**](EnumEnabledStatus.md) |  | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Icon** | Pointer to [**IdentityProviderCommonIcon**](IdentityProviderCommonIcon.md) |  | [optional] 
**Id** | Pointer to **string** | The resource ID. | [optional] [readonly] 
**LoginButtonIcon** | Pointer to [**IdentityProviderCommonLoginButtonIcon**](IdentityProviderCommonLoginButtonIcon.md) |  | [optional] 
**Name** | **string** | The name of the IdP. | 
**Registration** | Pointer to [**IdentityProviderCommonRegistration**](IdentityProviderCommonRegistration.md) |  | [optional] 
**Type** | [**EnumIdentityProviderExt**](EnumIdentityProviderExt.md) |  | 
**ClientId** | **string** | A string that specifies the application ID from Github. This is a required property. | 
**ClientSecret** | **string** | A string that specifies the application secret from Github. This is a required property. | 

## Methods

### NewIdentityProviderGithub

`func NewIdentityProviderGithub(enabled EnumEnabledStatus, name string, type_ EnumIdentityProviderExt, clientId string, clientSecret string, ) *IdentityProviderGithub`

NewIdentityProviderGithub instantiates a new IdentityProviderGithub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderGithubWithDefaults

`func NewIdentityProviderGithubWithDefaults() *IdentityProviderGithub`

NewIdentityProviderGithubWithDefaults instantiates a new IdentityProviderGithub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IdentityProviderGithub) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityProviderGithub) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityProviderGithub) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityProviderGithub) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *IdentityProviderGithub) GetEnabled() EnumEnabledStatus`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityProviderGithub) GetEnabledOk() (*EnumEnabledStatus, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityProviderGithub) SetEnabled(v EnumEnabledStatus)`

SetEnabled sets Enabled field to given value.


### GetEnvironment

`func (o *IdentityProviderGithub) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *IdentityProviderGithub) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *IdentityProviderGithub) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *IdentityProviderGithub) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIcon

`func (o *IdentityProviderGithub) GetIcon() IdentityProviderCommonIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *IdentityProviderGithub) GetIconOk() (*IdentityProviderCommonIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *IdentityProviderGithub) SetIcon(v IdentityProviderCommonIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *IdentityProviderGithub) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *IdentityProviderGithub) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProviderGithub) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProviderGithub) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProviderGithub) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoginButtonIcon

`func (o *IdentityProviderGithub) GetLoginButtonIcon() IdentityProviderCommonLoginButtonIcon`

GetLoginButtonIcon returns the LoginButtonIcon field if non-nil, zero value otherwise.

### GetLoginButtonIconOk

`func (o *IdentityProviderGithub) GetLoginButtonIconOk() (*IdentityProviderCommonLoginButtonIcon, bool)`

GetLoginButtonIconOk returns a tuple with the LoginButtonIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginButtonIcon

`func (o *IdentityProviderGithub) SetLoginButtonIcon(v IdentityProviderCommonLoginButtonIcon)`

SetLoginButtonIcon sets LoginButtonIcon field to given value.

### HasLoginButtonIcon

`func (o *IdentityProviderGithub) HasLoginButtonIcon() bool`

HasLoginButtonIcon returns a boolean if a field has been set.

### GetName

`func (o *IdentityProviderGithub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityProviderGithub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityProviderGithub) SetName(v string)`

SetName sets Name field to given value.


### GetRegistration

`func (o *IdentityProviderGithub) GetRegistration() IdentityProviderCommonRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *IdentityProviderGithub) GetRegistrationOk() (*IdentityProviderCommonRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *IdentityProviderGithub) SetRegistration(v IdentityProviderCommonRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *IdentityProviderGithub) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetType

`func (o *IdentityProviderGithub) GetType() EnumIdentityProviderExt`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityProviderGithub) GetTypeOk() (*EnumIdentityProviderExt, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityProviderGithub) SetType(v EnumIdentityProviderExt)`

SetType sets Type field to given value.


### GetClientId

`func (o *IdentityProviderGithub) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityProviderGithub) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityProviderGithub) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *IdentityProviderGithub) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IdentityProviderGithub) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IdentityProviderGithub) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

