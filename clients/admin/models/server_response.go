// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServerResponse server response
//
// swagger:model ServerResponse
type ServerResponse struct {

	// Access token strategy
	//
	// You can choose to go either with `JWT` or `opaque` tokens.
	//
	// The content of JSON Web Tokens is readable and it is possible to be decoded by anyone that
	// has a secret or a public key in their possession.
	//
	// Opaque tokens are in a proprietary form that contains an identifier to information stored on
	// the athorization server. To validate an opaque token, the recipient must call the server that
	// issued the token.
	// Example: jwt
	// Enum: [jwt opaque]
	AccessTokenStrategy string `json:"access_token_strategy,omitempty"`

	// Access token time to live
	//
	// After an access token reaches its time to live, it expires and it cannot be used to
	// authenticate the client application.
	// Example: 1h10m30s
	// Format: duration
	AccessTokenTTL strfmt.Duration `json:"access_token_ttl,omitempty"`

	// Authorization code time to live
	//
	// After an authorization code reaches its time to live, it expires and it cannot be used to
	// authorize the request to the `/token` endpoint.
	// Example: 10m0s
	// Format: duration
	AuthorizationCodeTTL strfmt.Duration `json:"authorization_code_ttl,omitempty"`

	// backchannel token delivery modes supported
	BackchannelTokenDeliveryModesSupported []string `json:"backchannel_token_delivery_modes_supported"`

	// backchannel user code parameter supported
	BackchannelUserCodeParameterSupported bool `json:"backchannel_user_code_parameter_supported,omitempty"`

	// Your server's label color in a HEX format.
	// Example: #007FFF
	Color string `json:"color,omitempty"`

	// Namespace used when creating Open Banking Brasil consent IDs
	//
	// Consent IDs will have the form of urn:<consent_id_namespace>:<uid> when this is set.
	// If this is empty, the tenant id is used by default
	ConsentIDNamespace string `json:"consent_id_namespace,omitempty"`

	// Cookie max age
	//
	// Defines how long a cookie can live until it expires.
	// Example: 1h10m30s
	// Format: duration
	CookieMaxAge strfmt.Duration `json:"cookie_max_age,omitempty"`

	// Defines a custom issuer URL that can be used as the value of the `iss` claim in an access
	// token.
	//
	// If not provided, it is built dynamically based on the server's URL.
	// Example: http://example.com/default/default
	CustomIssuerURL string `json:"custom_issuer_url,omitempty"`

	// dynamic client registration
	DynamicClientRegistration *DynamicClientRegistrationSettings `json:"dynamic_client_registration,omitempty"`

	// You can use this property to define a separator that is used for dynamic scopes.
	//
	// For example, the default separator is `.`, so the scope could look like the following:
	// `users.*`.
	//
	// For Open Banking Brazil compliant servers, the `:` separator should be used.
	DynamicScopeSeparator string `json:"dynamic_scope_separator,omitempty"`

	// When enabled, the authorization server will encrypt any id tokens it issues
	EnableIDTokenEncryption bool `json:"enable_id_token_encryption,omitempty"`

	// If enabled, IDP discovery automatically redirects the user to their own IDP and does not
	// display IDPs of other users while the users accesses the server/application.
	// Example: false
	EnableIdpDiscovery bool `json:"enable_idp_discovery,omitempty"`

	// If enabled, it is possible to manually register clients withouth the use of software
	// statements.
	//
	// This flag is enabled, when the `enable_trust_anchor` flag is set to `false`. You can disable
	// it using API, but it cannot be manually enabled.
	EnableLegacyClientsWithNoSoftwareStatement bool `json:"enable_legacy_clients_with_no_software_statement,omitempty"`

	// If enabled, the server is visible on the Quick Access tab on the login page.
	EnableQuickAccess bool `json:"enable_quick_access,omitempty"`

	// If enabled, it makes it obligatory to provide a software statement signed by a trusted certificate authority
	// when registering a client application with the use of the Dynamic Client Registration (DCR).
	//
	// In public key infrastructure (PKI), trust anchors are certification authorities. They are
	// represented by a certificate that is used to verify the signature of a certificate issued by
	// a particular trust anchor.
	EnableTrustAnchor bool `json:"enable_trust_anchor,omitempty"`

	// Define whether you want to enforce using the Proof Key of Code Exchange (PKCE) for both
	// private and public clients.
	//
	// PKCE is an OAuth security extension that prevents malicious applications or codes that
	// intercepted authorization code from exchanging it for an access token.
	// Example: false
	EnforcePkce bool `json:"enforce_pkce,omitempty"`

	// Define whether you want to enforce using the Proof Key of Code Exchange (PKCE) for
	// public clients.
	//
	// Public clients, like mobile applications or JavaScript-based applications, by their design,
	// cannot store client secrets securely. For such clients, even encrypting the client secret
	// inside the application’s code is not a reliable way of protecting secrets as the application
	// can be decompiled and the client secret can be extracted while it is decrypted in the memory
	// of the application.
	//
	// For those reasons, ACP supports the use of PKCE as an addition to the authorization code
	// grant flow to provide a secure alternative to the implicit grant flow.
	// Example: false
	EnforcePkceForPublicClients bool `json:"enforce_pkce_for_public_clients,omitempty"`

	// An array that defines which of the OAuth 2.0 grant types are enabled for the authorization server.
	// Example: ["authorization_code","implicit","refresh_token","client_credentials"]
	GrantTypes []string `json:"grant_types"`

	// Unique identifier of an authorization server (workspace)
	//
	// If not provided, a random ID is generated.
	// Example: default
	ID string `json:"id,omitempty"`

	// ID token time to live
	//
	// After an ID token reaches its time to live, it expires and it cannot be used to provide
	// user profile information to a client application.
	// Example: 1h10m30s
	// Format: duration
	IDTokenTTL strfmt.Duration `json:"id_token_ttl,omitempty"`

	// Issuer ID that will be used to set `iss` claim on signed messages
	//
	// If issuer_id is not set then default issuer_url will be used
	// Example: 5647fe90-f6bc-11eb-9a03-0242ac130003
	IssuerID string `json:"issuer_id,omitempty"`

	// issuer url
	IssuerURL string `json:"issuer_url,omitempty"`

	// jwks
	Jwks *ServerJWKs `json:"jwks,omitempty"`

	// Determines which type of asymmetric algorithms (RSA or ECDSA) is used to generate keys for signing access and
	// ID tokens.
	//
	// It is used only as an input parameter for the Create Authorization Server API.
	// Example: rsa
	// Enum: [rsa ecdsa ps]
	KeyType string `json:"key_type,omitempty"`

	// Logo URI
	LogoURI string `json:"logo_uri,omitempty"`

	// mtls issuer url
	MtlsIssuerURL string `json:"mtls_issuer_url,omitempty"`

	// Display name of your authorization server
	// Example: Sample authorization server
	Name string `json:"name,omitempty"`

	// The profile of a server
	//
	// ACP is delivered with preconfigured workspace templates that enable quick and easy setup for
	// specific configuration patterns. For example, you can instantly create an Open Banking
	// compliant workspace that has all of the required mechanisms and settings already in place.
	// Example: default
	// Enum: [default demo workforce consumer partners third_party fapi_advanced fapi_rw fapi_ro openbanking_uk_fapi_advanced openbanking_uk openbanking_br cdr_australia cdr_australia_fapi_rw]
	Profile string `json:"profile,omitempty"`

	// Custom pushed authentication request TTL
	// If not provided, TTL is set to 60 seconds.
	// Format: duration
	PushedAuthorizationRequestTTL strfmt.Duration `json:"pushed_authorization_request_ttl,omitempty"`

	// Refresh token time to live
	//
	// After a refresh token reaches its time to live, it expires and it cannot be used to obtain
	// new access tokens for a client application.
	// Example: 720h0m0s
	// Format: duration
	RefreshTokenTTL strfmt.Duration `json:"refresh_token_ttl,omitempty"`

	// Boolean parameter indicating whether the authorization server accepts authorization request data only via PAR.
	RequirePushedAuthorizationRequests bool `json:"require_pushed_authorization_requests,omitempty"`

	// You can provide root Certificate Authority (CA) certificates encoded to the Privacy-Enhanced
	// Mail (PEM) file format which are used for the `tls_client_auth` and the
	// `self_signed_tls_client_auth` client authentication methods that use the Mutual
	// Transport Layer Security (mTLS).
	//
	// If not set, the system root CA certifiates are used instead.
	RootCas string `json:"root_cas,omitempty"`

	// An array of rotated secrets that are still used to validate tokens
	// Example: ["jFpwIvuKJP46J71WqszPv1SrzoUr-cSILP9EPdlClB4"]
	RotatedSecrets []string `json:"rotated_secrets"`

	// Secret used for hashing
	//
	// It must have at least 32 characters. If not provided, it is generated.
	// Example: hW5WhKX_7w7BLwUQ6mn7Cp70_OoKI_F1y1hLS5U8lIU
	Secret string `json:"secret,omitempty"`

	// Salt used to hash `subject` when the `pairwise` subject type is used.
	//
	// Salt is a random data which is used as an additional input to one-way functions that hashes
	// data, passwords, and, in this case, subjects.
	//
	// It is recommended that your salt value is long for security reasons. Preferably, the salt
	// value should be at least the same length as the output of the hash.
	//
	// If not provided, it is generated.
	SubjectIdentifierAlgorithmSalt string `json:"subject_identifier_algorithm_salt,omitempty"`

	// An array that defines supported subject identifier types.
	//
	// Subject identifiers are locally unique and never reassigned identifiers within the Issuer
	// for the end-user and are inteded to be consumed by client applications. There are two types
	// of subject identifiers: `public` and `pairwise`.
	//
	// `public` identifiers provide the same `sub` claim value to all client applications.
	// `pairwise` identifiers provide different `sub` claim values to each client application. With
	// this approach, it makes it impossible for client applications to correlate the end-user's
	// activity without their permission.
	// Example: ["public","pairwise"]
	SubjectIdentifierTypes []string `json:"subject_identifier_types"`

	// supported application purposes
	// Example: ["single_page","server_web","mobile_desktop","service","legacy"]
	SupportedApplicationPurposes []string `json:"supported_application_purposes"`

	// ID of a tenant
	// Example: default
	// Required: true
	TenantID string `json:"tenant_id"`

	// An array that lists all of the supported token endpoint authentication methods for the
	// authorization server.
	TokenEndpointAuthMethods []string `json:"token_endpoint_auth_methods"`

	// Deprecated: Use TokenEndpointAuthMethods instead
	TokenEndpointAuthnMethods []string `json:"token_endpoint_authn_methods"`

	// trust anchor configuration
	TrustAnchorConfiguration *TrustAnchorConfiguration `json:"trust_anchor_configuration,omitempty"`

	// Server type
	//
	// It is an internal property used to recognize if the server is created for an admin portal,
	// a developer portal, or if it is a system or a regular workspace.
	// Example: regular
	// Enum: [admin developer system regular]
	Type string `json:"type,omitempty"`
}

// Validate validates this server response
func (m *ServerResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessTokenStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccessTokenTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorizationCodeTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCookieMaxAge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDynamicClientRegistration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrantTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDTokenTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJwks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePushedAuthorizationRequestTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefreshTokenTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubjectIdentifierTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportedApplicationPurposes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenEndpointAuthMethods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenEndpointAuthnMethods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrustAnchorConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serverResponseTypeAccessTokenStrategyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["jwt","opaque"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverResponseTypeAccessTokenStrategyPropEnum = append(serverResponseTypeAccessTokenStrategyPropEnum, v)
	}
}

const (

	// ServerResponseAccessTokenStrategyJwt captures enum value "jwt"
	ServerResponseAccessTokenStrategyJwt string = "jwt"

	// ServerResponseAccessTokenStrategyOpaque captures enum value "opaque"
	ServerResponseAccessTokenStrategyOpaque string = "opaque"
)

// prop value enum
func (m *ServerResponse) validateAccessTokenStrategyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serverResponseTypeAccessTokenStrategyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServerResponse) validateAccessTokenStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessTokenStrategy) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessTokenStrategyEnum("access_token_strategy", "body", m.AccessTokenStrategy); err != nil {
		return err
	}

	return nil
}

func (m *ServerResponse) validateAccessTokenTTL(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessTokenTTL) { // not required
		return nil
	}

	if err := validate.FormatOf("access_token_ttl", "body", "duration", m.AccessTokenTTL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServerResponse) validateAuthorizationCodeTTL(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthorizationCodeTTL) { // not required
		return nil
	}

	if err := validate.FormatOf("authorization_code_ttl", "body", "duration", m.AuthorizationCodeTTL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServerResponse) validateCookieMaxAge(formats strfmt.Registry) error {
	if swag.IsZero(m.CookieMaxAge) { // not required
		return nil
	}

	if err := validate.FormatOf("cookie_max_age", "body", "duration", m.CookieMaxAge.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServerResponse) validateDynamicClientRegistration(formats strfmt.Registry) error {
	if swag.IsZero(m.DynamicClientRegistration) { // not required
		return nil
	}

	if m.DynamicClientRegistration != nil {
		if err := m.DynamicClientRegistration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamic_client_registration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dynamic_client_registration")
			}
			return err
		}
	}

	return nil
}

var serverResponseGrantTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["authorization_code","implicit","client_credentials","refresh_token","password","urn:ietf:params:oauth:grant-type:jwt-bearer","urn:openid:params:grant-type:ciba","urn:ietf:params:oauth:grant-type:token-exchange"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverResponseGrantTypesItemsEnum = append(serverResponseGrantTypesItemsEnum, v)
	}
}

func (m *ServerResponse) validateGrantTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serverResponseGrantTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServerResponse) validateGrantTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.GrantTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.GrantTypes); i++ {

		// value enum
		if err := m.validateGrantTypesItemsEnum("grant_types"+"."+strconv.Itoa(i), "body", m.GrantTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ServerResponse) validateIDTokenTTL(formats strfmt.Registry) error {
	if swag.IsZero(m.IDTokenTTL) { // not required
		return nil
	}

	if err := validate.FormatOf("id_token_ttl", "body", "duration", m.IDTokenTTL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServerResponse) validateJwks(formats strfmt.Registry) error {
	if swag.IsZero(m.Jwks) { // not required
		return nil
	}

	if m.Jwks != nil {
		if err := m.Jwks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jwks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("jwks")
			}
			return err
		}
	}

	return nil
}

var serverResponseTypeKeyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rsa","ecdsa","ps"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverResponseTypeKeyTypePropEnum = append(serverResponseTypeKeyTypePropEnum, v)
	}
}

const (

	// ServerResponseKeyTypeRsa captures enum value "rsa"
	ServerResponseKeyTypeRsa string = "rsa"

	// ServerResponseKeyTypeEcdsa captures enum value "ecdsa"
	ServerResponseKeyTypeEcdsa string = "ecdsa"

	// ServerResponseKeyTypePs captures enum value "ps"
	ServerResponseKeyTypePs string = "ps"
)

// prop value enum
func (m *ServerResponse) validateKeyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serverResponseTypeKeyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServerResponse) validateKeyType(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyType) { // not required
		return nil
	}

	// value enum
	if err := m.validateKeyTypeEnum("key_type", "body", m.KeyType); err != nil {
		return err
	}

	return nil
}

var serverResponseTypeProfilePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","demo","workforce","consumer","partners","third_party","fapi_advanced","fapi_rw","fapi_ro","openbanking_uk_fapi_advanced","openbanking_uk","openbanking_br","cdr_australia","cdr_australia_fapi_rw"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverResponseTypeProfilePropEnum = append(serverResponseTypeProfilePropEnum, v)
	}
}

const (

	// ServerResponseProfileDefault captures enum value "default"
	ServerResponseProfileDefault string = "default"

	// ServerResponseProfileDemo captures enum value "demo"
	ServerResponseProfileDemo string = "demo"

	// ServerResponseProfileWorkforce captures enum value "workforce"
	ServerResponseProfileWorkforce string = "workforce"

	// ServerResponseProfileConsumer captures enum value "consumer"
	ServerResponseProfileConsumer string = "consumer"

	// ServerResponseProfilePartners captures enum value "partners"
	ServerResponseProfilePartners string = "partners"

	// ServerResponseProfileThirdParty captures enum value "third_party"
	ServerResponseProfileThirdParty string = "third_party"

	// ServerResponseProfileFapiAdvanced captures enum value "fapi_advanced"
	ServerResponseProfileFapiAdvanced string = "fapi_advanced"

	// ServerResponseProfileFapiRw captures enum value "fapi_rw"
	ServerResponseProfileFapiRw string = "fapi_rw"

	// ServerResponseProfileFapiRo captures enum value "fapi_ro"
	ServerResponseProfileFapiRo string = "fapi_ro"

	// ServerResponseProfileOpenbankingUkFapiAdvanced captures enum value "openbanking_uk_fapi_advanced"
	ServerResponseProfileOpenbankingUkFapiAdvanced string = "openbanking_uk_fapi_advanced"

	// ServerResponseProfileOpenbankingUk captures enum value "openbanking_uk"
	ServerResponseProfileOpenbankingUk string = "openbanking_uk"

	// ServerResponseProfileOpenbankingBr captures enum value "openbanking_br"
	ServerResponseProfileOpenbankingBr string = "openbanking_br"

	// ServerResponseProfileCdrAustralia captures enum value "cdr_australia"
	ServerResponseProfileCdrAustralia string = "cdr_australia"

	// ServerResponseProfileCdrAustraliaFapiRw captures enum value "cdr_australia_fapi_rw"
	ServerResponseProfileCdrAustraliaFapiRw string = "cdr_australia_fapi_rw"
)

// prop value enum
func (m *ServerResponse) validateProfileEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serverResponseTypeProfilePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServerResponse) validateProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.Profile) { // not required
		return nil
	}

	// value enum
	if err := m.validateProfileEnum("profile", "body", m.Profile); err != nil {
		return err
	}

	return nil
}

func (m *ServerResponse) validatePushedAuthorizationRequestTTL(formats strfmt.Registry) error {
	if swag.IsZero(m.PushedAuthorizationRequestTTL) { // not required
		return nil
	}

	if err := validate.FormatOf("pushed_authorization_request_ttl", "body", "duration", m.PushedAuthorizationRequestTTL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServerResponse) validateRefreshTokenTTL(formats strfmt.Registry) error {
	if swag.IsZero(m.RefreshTokenTTL) { // not required
		return nil
	}

	if err := validate.FormatOf("refresh_token_ttl", "body", "duration", m.RefreshTokenTTL.String(), formats); err != nil {
		return err
	}

	return nil
}

var serverResponseSubjectIdentifierTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","pairwise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverResponseSubjectIdentifierTypesItemsEnum = append(serverResponseSubjectIdentifierTypesItemsEnum, v)
	}
}

func (m *ServerResponse) validateSubjectIdentifierTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serverResponseSubjectIdentifierTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServerResponse) validateSubjectIdentifierTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.SubjectIdentifierTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.SubjectIdentifierTypes); i++ {

		// value enum
		if err := m.validateSubjectIdentifierTypesItemsEnum("subject_identifier_types"+"."+strconv.Itoa(i), "body", m.SubjectIdentifierTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

var serverResponseSupportedApplicationPurposesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["single_page","server_web","mobile_desktop","service","legacy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverResponseSupportedApplicationPurposesItemsEnum = append(serverResponseSupportedApplicationPurposesItemsEnum, v)
	}
}

func (m *ServerResponse) validateSupportedApplicationPurposesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serverResponseSupportedApplicationPurposesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServerResponse) validateSupportedApplicationPurposes(formats strfmt.Registry) error {
	if swag.IsZero(m.SupportedApplicationPurposes) { // not required
		return nil
	}

	for i := 0; i < len(m.SupportedApplicationPurposes); i++ {

		// value enum
		if err := m.validateSupportedApplicationPurposesItemsEnum("supported_application_purposes"+"."+strconv.Itoa(i), "body", m.SupportedApplicationPurposes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ServerResponse) validateTenantID(formats strfmt.Registry) error {

	if err := validate.RequiredString("tenant_id", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

var serverResponseTokenEndpointAuthMethodsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["client_secret_basic","client_secret_post","client_secret_jwt","private_key_jwt","self_signed_tls_client_auth","tls_client_auth","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverResponseTokenEndpointAuthMethodsItemsEnum = append(serverResponseTokenEndpointAuthMethodsItemsEnum, v)
	}
}

func (m *ServerResponse) validateTokenEndpointAuthMethodsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serverResponseTokenEndpointAuthMethodsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServerResponse) validateTokenEndpointAuthMethods(formats strfmt.Registry) error {
	if swag.IsZero(m.TokenEndpointAuthMethods) { // not required
		return nil
	}

	for i := 0; i < len(m.TokenEndpointAuthMethods); i++ {

		// value enum
		if err := m.validateTokenEndpointAuthMethodsItemsEnum("token_endpoint_auth_methods"+"."+strconv.Itoa(i), "body", m.TokenEndpointAuthMethods[i]); err != nil {
			return err
		}

	}

	return nil
}

var serverResponseTokenEndpointAuthnMethodsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["client_secret_basic","client_secret_post","client_secret_jwt","private_key_jwt","self_signed_tls_client_auth","tls_client_auth","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverResponseTokenEndpointAuthnMethodsItemsEnum = append(serverResponseTokenEndpointAuthnMethodsItemsEnum, v)
	}
}

func (m *ServerResponse) validateTokenEndpointAuthnMethodsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serverResponseTokenEndpointAuthnMethodsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServerResponse) validateTokenEndpointAuthnMethods(formats strfmt.Registry) error {
	if swag.IsZero(m.TokenEndpointAuthnMethods) { // not required
		return nil
	}

	for i := 0; i < len(m.TokenEndpointAuthnMethods); i++ {

		// value enum
		if err := m.validateTokenEndpointAuthnMethodsItemsEnum("token_endpoint_authn_methods"+"."+strconv.Itoa(i), "body", m.TokenEndpointAuthnMethods[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ServerResponse) validateTrustAnchorConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.TrustAnchorConfiguration) { // not required
		return nil
	}

	if m.TrustAnchorConfiguration != nil {
		if err := m.TrustAnchorConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trust_anchor_configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trust_anchor_configuration")
			}
			return err
		}
	}

	return nil
}

var serverResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["admin","developer","system","regular"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverResponseTypeTypePropEnum = append(serverResponseTypeTypePropEnum, v)
	}
}

const (

	// ServerResponseTypeAdmin captures enum value "admin"
	ServerResponseTypeAdmin string = "admin"

	// ServerResponseTypeDeveloper captures enum value "developer"
	ServerResponseTypeDeveloper string = "developer"

	// ServerResponseTypeSystem captures enum value "system"
	ServerResponseTypeSystem string = "system"

	// ServerResponseTypeRegular captures enum value "regular"
	ServerResponseTypeRegular string = "regular"
)

// prop value enum
func (m *ServerResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serverResponseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServerResponse) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this server response based on the context it is used
func (m *ServerResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDynamicClientRegistration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJwks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrustAnchorConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerResponse) contextValidateDynamicClientRegistration(ctx context.Context, formats strfmt.Registry) error {

	if m.DynamicClientRegistration != nil {
		if err := m.DynamicClientRegistration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dynamic_client_registration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dynamic_client_registration")
			}
			return err
		}
	}

	return nil
}

func (m *ServerResponse) contextValidateJwks(ctx context.Context, formats strfmt.Registry) error {

	if m.Jwks != nil {
		if err := m.Jwks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jwks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("jwks")
			}
			return err
		}
	}

	return nil
}

func (m *ServerResponse) contextValidateTrustAnchorConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.TrustAnchorConfiguration != nil {
		if err := m.TrustAnchorConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trust_anchor_configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trust_anchor_configuration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerResponse) UnmarshalBinary(b []byte) error {
	var res ServerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
