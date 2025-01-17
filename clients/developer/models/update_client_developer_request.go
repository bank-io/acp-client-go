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

// UpdateClientDeveloperRequest update client developer request
//
// swagger:model UpdateClientDeveloperRequest
type UpdateClientDeveloperRequest struct {

	// String represented type of a client application
	//
	// Client applications can be either of a `web` or `native` types.
	//
	// Web applications include clients like server web applications or service apps.
	//
	// Native applications include single-page applications (SPAs) and mobile or desktop
	// applications.
	//
	// Depending on the type of your application remember to choose appropriate security measures.
	// Example: web
	ApplicationType string `json:"application_type,omitempty"`

	// An array of dynamically calculated application types that can be used for filtering
	// Example: ["single_page","server_web","mobile_desktop","service","legacy","dcr"]
	// Read Only: true
	ApplicationTypes []string `json:"application_types"`

	// Identity of the intended recipients (the audience)
	//
	// Typically, the audience may be a single resources server or a list of resource servers.
	// It is considered a good practice to limit the audience of the token for security purposes.
	Audience []string `json:"audience"`

	// OPTIONAL. The JWS algorithm alg value that the Client will use for signing authentication requests.
	// When omitted, the Client will not send signed authentication requests.
	BackchannelAuthenticationRequestSigningAlg string `json:"backchannel_authentication_request_signing_alg,omitempty"`

	// REQUIRED if the token delivery mode is set to ping or push.
	// This is the endpoint to which the OP will post a notification after a successful or failed end-user authentication.
	// It MUST be an HTTPS URL.
	BackchannelClientNotificationEndpoint string `json:"backchannel_client_notification_endpoint,omitempty"`

	// REQUIRED. One of the following values: poll, ping, or push.
	BackchannelTokenDeliveryMode string `json:"backchannel_token_delivery_mode,omitempty"`

	// OPTIONAL. Boolean value specifying whether the Client supports the user_code parameter.
	// If omitted, the default value is false.
	// This parameter only applies when OP parameter backchannel_user_code_parameter_supported is true.
	BackchannelUserCodeParameter bool `json:"backchannel_user_code_parameter,omitempty"`

	// Time at which the client identifier was issued.
	//
	// The time is represented as the number of seconds from
	// 1970-01-01T00:00:00Z as measured in UTC until the date/time of issuance.
	ClientIDIssuedAt int64 `json:"client_id_issued_at,omitempty"`

	// Human readable name of a client application
	// Example: My app
	ClientName string `json:"client_name,omitempty"`

	// OAuth client secret
	//
	// If not provided, a random client secret is generated.
	// Min Length: 32
	ClientSecret string `json:"client_secret,omitempty"`

	// The `client_secret_expires_at` holds an integer that defines the time at which the client
	// secret expires
	//
	// If the client secret does not expire, the value should be set to `0`.
	ClientSecretExpiresAt int64 `json:"client_secret_expires_at,omitempty"`

	// URI of a client application
	ClientURI string `json:"client_uri,omitempty"`

	// Description of a client application
	Description string `json:"description,omitempty"`

	// An array of allowed OAuth client grant types
	//
	// The `grantTypes` array stores OAuth flows that are allowed for a given client application.
	//
	// To know more about OAuth grant flows, see the
	// [ACP grant flows documentation](https://docs.authorization.cloudentity.com/features/oauth/grant_flows/).
	// Example: ["password","refresh_token","client_credentials","implicit","authorization_code"]
	GrantTypes []string `json:"grant_types"`

	// An array of hashed rotated client secrets
	HashedRotatedSecrets []string `json:"hashed_rotated_secrets"`

	// Hashed client secret
	//
	// Hashing client secrets provides additional security for your secrets storage as it hides
	// plaintext secrets from being viewed both in the UI and the database.
	HashedSecret string `json:"hashed_secret,omitempty"`

	// JWE alg algorithm for encrypting the ID Token issued to this Client
	// Enum: [RSA-OAEP RSA-OAEP-256]
	IDTokenEncryptedResponseAlg string `json:"id_token_encrypted_response_alg,omitempty"`

	// JWE enc algorithm for encrypting the ID Token issued to this Client
	// Enum: [A256GCM A128CBC-HS256]
	IDTokenEncryptedResponseEnc string `json:"id_token_encrypted_response_enc,omitempty"`

	// Algorithm for signing ID tokens issued for a client application
	//
	// The default value depends on authorization server configuration.
	// Example: ES256
	// Enum: [RS256 ES256 PS256]
	IDTokenSignedResponseAlg string `json:"id_token_signed_response_alg,omitempty"`

	// jwks
	Jwks *ClientJWKs `json:"jwks,omitempty"`

	// URL of JSON Web Key Set containing the public keys used by a client application to authenticate itself
	// with ACP
	JwksURI string `json:"jwks_uri,omitempty"`

	// Logo URI
	LogoURI string `json:"logo_uri,omitempty"`

	// External organisation ID
	//
	// This field is used as an aud for message signing
	// Example: 5647fe90-f6bc-11eb-9a03-0242ac130003
	OrganisationID string `json:"organisation_id,omitempty"`

	// Policy URL to read about how the profile data is used
	PolicyURI string `json:"policy_uri,omitempty"`

	// privacy
	Privacy *ClientPrivacy `json:"privacy,omitempty"`

	// An array of OAuth allowed redirect URIs
	//
	// Redirect URIs are used after a user authorizes an application and ACP redirect them back to
	// the application with an authorization code or an access token included in the URL.
	// Example: ["https://example.com/callback"]
	RedirectUris []string `json:"redirect_uris"`

	// Request object signing algorithm for the token endpoint
	//
	// ACP supports signing tokens using the RS256, ES256, and PS256 algorithms. If you do not want
	// to use a signing algorithm, the value of the parameter should be set to `none`.
	// Example: none
	// Enum: [any none RS256 ES256 PS256]
	RequestObjectSigningAlg string `json:"request_object_signing_alg,omitempty"`

	// Array of absolute URIs that points to the Request Object that holds authorization request parameters.
	RequestUris []string `json:"request_uris"`

	// Boolean parameter indicating whether the only means of initiating an authorization request the client is allowed to use is PAR.
	RequirePushedAuthorizationRequests bool `json:"require_pushed_authorization_requests,omitempty"`

	// An array of OAuth client response types configured for a client application
	//
	// The array may consist of the following arguments:
	//
	// `code` - when supplied as the value for the `response_type` parameter, a successful
	// response includes an authorization code
	//
	// `code token` - when supplied as the value for the `response_type` parameter, a successful
	// response includes an access token, an access token type, and an authorization code
	//
	// `id_token token` - when supplied as the value for the `response_type` parameter, a successful
	// response includes an access token, an access token type, and an ID token
	//
	// `code id_token token` - when supplied as the value for the `response_type` parameter, a successful
	// response includes an authorization code, an ID token, an access token, and an access token
	// type.
	//
	// `token` - when supplied as the value for the `response_type` parameter, a successful
	// response includes an access token and its type. This argument is used for the implicit grant
	// flow, but is not recommended. Instead, you should use either the authorization code grant
	// flow with PKCE or client authentication set to `none` and with the use of PKCE.
	// Example: ["token","id_token","code"]
	ResponseTypes []string `json:"response_types"`

	// An array of rotated OAuth client secrets
	RotatedSecrets []string `json:"rotated_secrets"`

	// Space separated scopes for compatibility with OAuth specification
	// Example: email offline_access openid
	Scope string `json:"scope,omitempty"`

	// An array of string represented scopes assigned to a client application
	// Example: ["email","offline_access","openid"]
	Scopes []string `json:"scopes"`

	// URL using the HTTPS scheme to be used in calculating Pseudonymous Identifiers by the OpenID Provider. The URL references a
	// file with a single JSON array of redirect_uri values.
	// Example: https://api.jsonbin.io/b/5db6ef08688fed59d2841f1e
	SectorIdentifierURI string `json:"sector_identifier_uri,omitempty"`

	// Subject identifier type
	//
	// Stores information if the subject identifier is of the `public` or the `pairwise` type.
	//
	// Subject identifiers are locally unique and never reassigned identifiers within the Issuer
	// for the end-user and are inteded to be consumed by client applications. There are two types
	// of subject identifiers: `public` and `pairwise`.
	//
	// For the `public` type, the value of the `sub` (subject) token claim is the same for all clients.
	//
	// For the `pairwise` type, a different `sub` (subject) token claim is provided for each client.
	// Using the `pairwise` subject identifier makes it impossible for client applications to correlate the end-user's
	// activity without their permission.
	// Example: public
	// Enum: [public pairwise]
	SubjectType string `json:"subject_type,omitempty"`

	// A string containing the value of an expected dNSName SAN entry in the certificate.
	TLSClientAuthSanDNS string `json:"tls_client_auth_san_dns,omitempty"`

	// A string containing the value of an expected rfc822Name SAN entry in the certificate.
	TLSClientAuthSanEmail string `json:"tls_client_auth_san_email,omitempty"`

	// A string representation of an IP address in either dotted decimal notation (for IPv4) or colon-delimited hexadecimal (for IPv6, as defined in [RFC5952]) that is expected to be present as an iPAddress SAN entry in the certificate.
	TLSClientAuthSanIP string `json:"tls_client_auth_san_ip,omitempty"`

	// A string containing the value of an expected uniformResourceIdentifier SAN entry in the certificate.
	TLSClientAuthSanURI string `json:"tls_client_auth_san_uri,omitempty"`

	// An [RFC4514] string representation of the expected subject distinguished name of the certificate.
	TLSClientAuthSubjectDn string `json:"tls_client_auth_subject_dn,omitempty"`

	// Boolean value indicating server support for mutual TLS client certificate-bound access tokens. If omitted, the default value is "false".
	TLSClientCertificateBoundAccessTokens bool `json:"tls_client_certificate_bound_access_tokens,omitempty"`

	// Token endpoint authentication method configured for a client application
	//
	// ACP supports the following client authentication methods:
	// client_secret_basic, client_secret_post, client_secret_jwt, private_key_jwt,
	// self_signed_tls_client_auth, tls_client_auth, none.
	//
	// To learn more, see the [ACP client authentication documentation](https://docs.authorization.cloudentity.com/features/oauth/client_auth/)
	// Example: client_secret_basic
	// Enum: [client_secret_basic client_secret_post client_secret_jwt private_key_jwt self_signed_tls_client_auth tls_client_auth none]
	TokenEndpointAuthMethod string `json:"token_endpoint_auth_method,omitempty"`

	// Signing algorithm for the token endpoint
	//
	// ACP supports signing tokens using the RS256, ES256, PS256, and HS256 algorithms.
	//
	// If your token endpoint authentication is set to the `private_key_jwt` method, the
	// `token_endpoint_auth_signing_alg` parameter must be either RS256, ES256, or PS256.
	//
	// If your token endpoint authentication is set to the `client_secret_jwt` method,
	// the `token_endpoint_auth_signing_alg` parameter must be HS256.
	// Example: none
	// Enum: [none RS256 ES256 PS256 HS256]
	TokenEndpointAuthSigningAlg string `json:"token_endpoint_auth_signing_alg,omitempty"`

	// Terms of Service URL
	TosURI string `json:"tos_uri,omitempty"`

	// JWS alg algorithm REQUIRED for signing UserInfo Responses.
	//
	// If specified, the response is JWT
	// [JWT] serialized, and signed using JWS.
	//
	// If omitted, the default behavior is for the UserInfo Response to return the Claims
	// as an UTF-8 encoded JSON object using the application/json content-type.
	// Example: none
	// Enum: [none RS256 ES256]
	UserinfoSignedResponseAlg string `json:"userinfo_signed_response_alg,omitempty"`
}

// Validate validates this update client developer request
func (m *UpdateClientDeveloperRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrantTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDTokenEncryptedResponseAlg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDTokenEncryptedResponseEnc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDTokenSignedResponseAlg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJwks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivacy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestObjectSigningAlg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubjectType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenEndpointAuthMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenEndpointAuthSigningAlg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserinfoSignedResponseAlg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateClientDeveloperRequestApplicationTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["single_page","server_web","mobile_desktop","service","legacy","dcr"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateClientDeveloperRequestApplicationTypesItemsEnum = append(updateClientDeveloperRequestApplicationTypesItemsEnum, v)
	}
}

func (m *UpdateClientDeveloperRequest) validateApplicationTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateClientDeveloperRequestApplicationTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateClientDeveloperRequest) validateApplicationTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.ApplicationTypes); i++ {

		// value enum
		if err := m.validateApplicationTypesItemsEnum("application_types"+"."+strconv.Itoa(i), "body", m.ApplicationTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *UpdateClientDeveloperRequest) validateClientSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientSecret) { // not required
		return nil
	}

	if err := validate.MinLength("client_secret", "body", m.ClientSecret, 32); err != nil {
		return err
	}

	return nil
}

var updateClientDeveloperRequestGrantTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["authorization_code","implicit","client_credentials","refresh_token","password","urn:ietf:params:oauth:grant-type:jwt-bearer","urn:openid:params:grant-type:ciba","urn:ietf:params:oauth:grant-type:token-exchange"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateClientDeveloperRequestGrantTypesItemsEnum = append(updateClientDeveloperRequestGrantTypesItemsEnum, v)
	}
}

func (m *UpdateClientDeveloperRequest) validateGrantTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateClientDeveloperRequestGrantTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateClientDeveloperRequest) validateGrantTypes(formats strfmt.Registry) error {
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

var updateClientDeveloperRequestTypeIDTokenEncryptedResponseAlgPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RSA-OAEP","RSA-OAEP-256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateClientDeveloperRequestTypeIDTokenEncryptedResponseAlgPropEnum = append(updateClientDeveloperRequestTypeIDTokenEncryptedResponseAlgPropEnum, v)
	}
}

const (

	// UpdateClientDeveloperRequestIDTokenEncryptedResponseAlgRSADashOAEP captures enum value "RSA-OAEP"
	UpdateClientDeveloperRequestIDTokenEncryptedResponseAlgRSADashOAEP string = "RSA-OAEP"

	// UpdateClientDeveloperRequestIDTokenEncryptedResponseAlgRSADashOAEPDash256 captures enum value "RSA-OAEP-256"
	UpdateClientDeveloperRequestIDTokenEncryptedResponseAlgRSADashOAEPDash256 string = "RSA-OAEP-256"
)

// prop value enum
func (m *UpdateClientDeveloperRequest) validateIDTokenEncryptedResponseAlgEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateClientDeveloperRequestTypeIDTokenEncryptedResponseAlgPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateClientDeveloperRequest) validateIDTokenEncryptedResponseAlg(formats strfmt.Registry) error {
	if swag.IsZero(m.IDTokenEncryptedResponseAlg) { // not required
		return nil
	}

	// value enum
	if err := m.validateIDTokenEncryptedResponseAlgEnum("id_token_encrypted_response_alg", "body", m.IDTokenEncryptedResponseAlg); err != nil {
		return err
	}

	return nil
}

var updateClientDeveloperRequestTypeIDTokenEncryptedResponseEncPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A256GCM","A128CBC-HS256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateClientDeveloperRequestTypeIDTokenEncryptedResponseEncPropEnum = append(updateClientDeveloperRequestTypeIDTokenEncryptedResponseEncPropEnum, v)
	}
}

const (

	// UpdateClientDeveloperRequestIDTokenEncryptedResponseEncA256GCM captures enum value "A256GCM"
	UpdateClientDeveloperRequestIDTokenEncryptedResponseEncA256GCM string = "A256GCM"

	// UpdateClientDeveloperRequestIDTokenEncryptedResponseEncA128CBCDashHS256 captures enum value "A128CBC-HS256"
	UpdateClientDeveloperRequestIDTokenEncryptedResponseEncA128CBCDashHS256 string = "A128CBC-HS256"
)

// prop value enum
func (m *UpdateClientDeveloperRequest) validateIDTokenEncryptedResponseEncEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateClientDeveloperRequestTypeIDTokenEncryptedResponseEncPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateClientDeveloperRequest) validateIDTokenEncryptedResponseEnc(formats strfmt.Registry) error {
	if swag.IsZero(m.IDTokenEncryptedResponseEnc) { // not required
		return nil
	}

	// value enum
	if err := m.validateIDTokenEncryptedResponseEncEnum("id_token_encrypted_response_enc", "body", m.IDTokenEncryptedResponseEnc); err != nil {
		return err
	}

	return nil
}

var updateClientDeveloperRequestTypeIDTokenSignedResponseAlgPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RS256","ES256","PS256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateClientDeveloperRequestTypeIDTokenSignedResponseAlgPropEnum = append(updateClientDeveloperRequestTypeIDTokenSignedResponseAlgPropEnum, v)
	}
}

const (

	// UpdateClientDeveloperRequestIDTokenSignedResponseAlgRS256 captures enum value "RS256"
	UpdateClientDeveloperRequestIDTokenSignedResponseAlgRS256 string = "RS256"

	// UpdateClientDeveloperRequestIDTokenSignedResponseAlgES256 captures enum value "ES256"
	UpdateClientDeveloperRequestIDTokenSignedResponseAlgES256 string = "ES256"

	// UpdateClientDeveloperRequestIDTokenSignedResponseAlgPS256 captures enum value "PS256"
	UpdateClientDeveloperRequestIDTokenSignedResponseAlgPS256 string = "PS256"
)

// prop value enum
func (m *UpdateClientDeveloperRequest) validateIDTokenSignedResponseAlgEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateClientDeveloperRequestTypeIDTokenSignedResponseAlgPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateClientDeveloperRequest) validateIDTokenSignedResponseAlg(formats strfmt.Registry) error {
	if swag.IsZero(m.IDTokenSignedResponseAlg) { // not required
		return nil
	}

	// value enum
	if err := m.validateIDTokenSignedResponseAlgEnum("id_token_signed_response_alg", "body", m.IDTokenSignedResponseAlg); err != nil {
		return err
	}

	return nil
}

func (m *UpdateClientDeveloperRequest) validateJwks(formats strfmt.Registry) error {
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

func (m *UpdateClientDeveloperRequest) validatePrivacy(formats strfmt.Registry) error {
	if swag.IsZero(m.Privacy) { // not required
		return nil
	}

	if m.Privacy != nil {
		if err := m.Privacy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privacy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("privacy")
			}
			return err
		}
	}

	return nil
}

var updateClientDeveloperRequestTypeRequestObjectSigningAlgPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["any","none","RS256","ES256","PS256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateClientDeveloperRequestTypeRequestObjectSigningAlgPropEnum = append(updateClientDeveloperRequestTypeRequestObjectSigningAlgPropEnum, v)
	}
}

const (

	// UpdateClientDeveloperRequestRequestObjectSigningAlgAny captures enum value "any"
	UpdateClientDeveloperRequestRequestObjectSigningAlgAny string = "any"

	// UpdateClientDeveloperRequestRequestObjectSigningAlgNone captures enum value "none"
	UpdateClientDeveloperRequestRequestObjectSigningAlgNone string = "none"

	// UpdateClientDeveloperRequestRequestObjectSigningAlgRS256 captures enum value "RS256"
	UpdateClientDeveloperRequestRequestObjectSigningAlgRS256 string = "RS256"

	// UpdateClientDeveloperRequestRequestObjectSigningAlgES256 captures enum value "ES256"
	UpdateClientDeveloperRequestRequestObjectSigningAlgES256 string = "ES256"

	// UpdateClientDeveloperRequestRequestObjectSigningAlgPS256 captures enum value "PS256"
	UpdateClientDeveloperRequestRequestObjectSigningAlgPS256 string = "PS256"
)

// prop value enum
func (m *UpdateClientDeveloperRequest) validateRequestObjectSigningAlgEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateClientDeveloperRequestTypeRequestObjectSigningAlgPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateClientDeveloperRequest) validateRequestObjectSigningAlg(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestObjectSigningAlg) { // not required
		return nil
	}

	// value enum
	if err := m.validateRequestObjectSigningAlgEnum("request_object_signing_alg", "body", m.RequestObjectSigningAlg); err != nil {
		return err
	}

	return nil
}

var updateClientDeveloperRequestResponseTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["token","id_token","code","code id_token","token id_token","token code","token id_token code"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateClientDeveloperRequestResponseTypesItemsEnum = append(updateClientDeveloperRequestResponseTypesItemsEnum, v)
	}
}

func (m *UpdateClientDeveloperRequest) validateResponseTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateClientDeveloperRequestResponseTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateClientDeveloperRequest) validateResponseTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.ResponseTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.ResponseTypes); i++ {

		// value enum
		if err := m.validateResponseTypesItemsEnum("response_types"+"."+strconv.Itoa(i), "body", m.ResponseTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

var updateClientDeveloperRequestTypeSubjectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","pairwise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateClientDeveloperRequestTypeSubjectTypePropEnum = append(updateClientDeveloperRequestTypeSubjectTypePropEnum, v)
	}
}

const (

	// UpdateClientDeveloperRequestSubjectTypePublic captures enum value "public"
	UpdateClientDeveloperRequestSubjectTypePublic string = "public"

	// UpdateClientDeveloperRequestSubjectTypePairwise captures enum value "pairwise"
	UpdateClientDeveloperRequestSubjectTypePairwise string = "pairwise"
)

// prop value enum
func (m *UpdateClientDeveloperRequest) validateSubjectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateClientDeveloperRequestTypeSubjectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateClientDeveloperRequest) validateSubjectType(formats strfmt.Registry) error {
	if swag.IsZero(m.SubjectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSubjectTypeEnum("subject_type", "body", m.SubjectType); err != nil {
		return err
	}

	return nil
}

var updateClientDeveloperRequestTypeTokenEndpointAuthMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["client_secret_basic","client_secret_post","client_secret_jwt","private_key_jwt","self_signed_tls_client_auth","tls_client_auth","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateClientDeveloperRequestTypeTokenEndpointAuthMethodPropEnum = append(updateClientDeveloperRequestTypeTokenEndpointAuthMethodPropEnum, v)
	}
}

const (

	// UpdateClientDeveloperRequestTokenEndpointAuthMethodClientSecretBasic captures enum value "client_secret_basic"
	UpdateClientDeveloperRequestTokenEndpointAuthMethodClientSecretBasic string = "client_secret_basic"

	// UpdateClientDeveloperRequestTokenEndpointAuthMethodClientSecretPost captures enum value "client_secret_post"
	UpdateClientDeveloperRequestTokenEndpointAuthMethodClientSecretPost string = "client_secret_post"

	// UpdateClientDeveloperRequestTokenEndpointAuthMethodClientSecretJwt captures enum value "client_secret_jwt"
	UpdateClientDeveloperRequestTokenEndpointAuthMethodClientSecretJwt string = "client_secret_jwt"

	// UpdateClientDeveloperRequestTokenEndpointAuthMethodPrivateKeyJwt captures enum value "private_key_jwt"
	UpdateClientDeveloperRequestTokenEndpointAuthMethodPrivateKeyJwt string = "private_key_jwt"

	// UpdateClientDeveloperRequestTokenEndpointAuthMethodSelfSignedTLSClientAuth captures enum value "self_signed_tls_client_auth"
	UpdateClientDeveloperRequestTokenEndpointAuthMethodSelfSignedTLSClientAuth string = "self_signed_tls_client_auth"

	// UpdateClientDeveloperRequestTokenEndpointAuthMethodTLSClientAuth captures enum value "tls_client_auth"
	UpdateClientDeveloperRequestTokenEndpointAuthMethodTLSClientAuth string = "tls_client_auth"

	// UpdateClientDeveloperRequestTokenEndpointAuthMethodNone captures enum value "none"
	UpdateClientDeveloperRequestTokenEndpointAuthMethodNone string = "none"
)

// prop value enum
func (m *UpdateClientDeveloperRequest) validateTokenEndpointAuthMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateClientDeveloperRequestTypeTokenEndpointAuthMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateClientDeveloperRequest) validateTokenEndpointAuthMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.TokenEndpointAuthMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateTokenEndpointAuthMethodEnum("token_endpoint_auth_method", "body", m.TokenEndpointAuthMethod); err != nil {
		return err
	}

	return nil
}

var updateClientDeveloperRequestTypeTokenEndpointAuthSigningAlgPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","RS256","ES256","PS256","HS256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateClientDeveloperRequestTypeTokenEndpointAuthSigningAlgPropEnum = append(updateClientDeveloperRequestTypeTokenEndpointAuthSigningAlgPropEnum, v)
	}
}

const (

	// UpdateClientDeveloperRequestTokenEndpointAuthSigningAlgNone captures enum value "none"
	UpdateClientDeveloperRequestTokenEndpointAuthSigningAlgNone string = "none"

	// UpdateClientDeveloperRequestTokenEndpointAuthSigningAlgRS256 captures enum value "RS256"
	UpdateClientDeveloperRequestTokenEndpointAuthSigningAlgRS256 string = "RS256"

	// UpdateClientDeveloperRequestTokenEndpointAuthSigningAlgES256 captures enum value "ES256"
	UpdateClientDeveloperRequestTokenEndpointAuthSigningAlgES256 string = "ES256"

	// UpdateClientDeveloperRequestTokenEndpointAuthSigningAlgPS256 captures enum value "PS256"
	UpdateClientDeveloperRequestTokenEndpointAuthSigningAlgPS256 string = "PS256"

	// UpdateClientDeveloperRequestTokenEndpointAuthSigningAlgHS256 captures enum value "HS256"
	UpdateClientDeveloperRequestTokenEndpointAuthSigningAlgHS256 string = "HS256"
)

// prop value enum
func (m *UpdateClientDeveloperRequest) validateTokenEndpointAuthSigningAlgEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateClientDeveloperRequestTypeTokenEndpointAuthSigningAlgPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateClientDeveloperRequest) validateTokenEndpointAuthSigningAlg(formats strfmt.Registry) error {
	if swag.IsZero(m.TokenEndpointAuthSigningAlg) { // not required
		return nil
	}

	// value enum
	if err := m.validateTokenEndpointAuthSigningAlgEnum("token_endpoint_auth_signing_alg", "body", m.TokenEndpointAuthSigningAlg); err != nil {
		return err
	}

	return nil
}

var updateClientDeveloperRequestTypeUserinfoSignedResponseAlgPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","RS256","ES256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateClientDeveloperRequestTypeUserinfoSignedResponseAlgPropEnum = append(updateClientDeveloperRequestTypeUserinfoSignedResponseAlgPropEnum, v)
	}
}

const (

	// UpdateClientDeveloperRequestUserinfoSignedResponseAlgNone captures enum value "none"
	UpdateClientDeveloperRequestUserinfoSignedResponseAlgNone string = "none"

	// UpdateClientDeveloperRequestUserinfoSignedResponseAlgRS256 captures enum value "RS256"
	UpdateClientDeveloperRequestUserinfoSignedResponseAlgRS256 string = "RS256"

	// UpdateClientDeveloperRequestUserinfoSignedResponseAlgES256 captures enum value "ES256"
	UpdateClientDeveloperRequestUserinfoSignedResponseAlgES256 string = "ES256"
)

// prop value enum
func (m *UpdateClientDeveloperRequest) validateUserinfoSignedResponseAlgEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateClientDeveloperRequestTypeUserinfoSignedResponseAlgPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateClientDeveloperRequest) validateUserinfoSignedResponseAlg(formats strfmt.Registry) error {
	if swag.IsZero(m.UserinfoSignedResponseAlg) { // not required
		return nil
	}

	// value enum
	if err := m.validateUserinfoSignedResponseAlgEnum("userinfo_signed_response_alg", "body", m.UserinfoSignedResponseAlg); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update client developer request based on the context it is used
func (m *UpdateClientDeveloperRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplicationTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJwks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrivacy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateClientDeveloperRequest) contextValidateApplicationTypes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "application_types", "body", []string(m.ApplicationTypes)); err != nil {
		return err
	}

	return nil
}

func (m *UpdateClientDeveloperRequest) contextValidateJwks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UpdateClientDeveloperRequest) contextValidatePrivacy(ctx context.Context, formats strfmt.Registry) error {

	if m.Privacy != nil {
		if err := m.Privacy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privacy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("privacy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateClientDeveloperRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateClientDeveloperRequest) UnmarshalBinary(b []byte) error {
	var res UpdateClientDeveloperRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
