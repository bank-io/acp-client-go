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

// OpenbankingUKDynamicClientRegistrationRequest OpenbankingUKDynamicClientRegistrationRequest represents decoded jwt of dcr obuk request
// Reference: https://openbankinguk.github.io/dcr-docs-pub/v3.2/dynamic-client-registration.html
//
// swagger:model OpenbankingUKDynamicClientRegistrationRequest
type OpenbankingUKDynamicClientRegistrationRequest struct {

	// String represented type of a client application
	//
	// Client applications can be either of a `web` or `native` types.
	//
	// Web applications include clients like server web applications or service apps.
	//
	// Native applications include single-page applications (SPAs) and mobile or desktop
	// applications.
	//
	// Depending on the type of your application remember to choose appropiate security measures.
	// Example: web
	ApplicationType string `json:"application_type,omitempty"`

	// An array of dynamically calculated application types that can be used for filtering
	// Example: ["single_page","server_web","mobile_desktop","service","legacy","dcr"]
	// Read Only: true
	ApplicationTypes []string `json:"application_types"`

	// The audience for the request. This should be the unique identifier
	// for the ASPSP issued by the issuer of the software statement.
	// An ASPSP processing the software statement may validate the value
	// of the claim and reject software statements for which the ASPSP
	// is not the audience.The value must be a Base62 encoded GUID.
	// Pattern: ^[0-9a-zA-Z]{1,18}$
	Aud string `json:"aud,omitempty"`

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

	// The `client_secret_expires_at` holds an integer that defines the time at which the client
	// secret expires
	//
	// If the client secret does not expire, the value should be set to `0`.
	ClientSecretExpiresAt int64 `json:"client_secret_expires_at,omitempty"`

	// URI of a client application
	ClientURI string `json:"client_uri,omitempty"`

	// Description of a client application
	Description string `json:"description,omitempty"`

	// The time at which the request expires expressed as seconds since
	// the epoch. An ASPSP processing the request must reject requests
	// where the current time is greater than the time specified in the claim.
	Exp int64 `json:"exp,omitempty"`

	// An array of allowed OAuth client grant types
	//
	// The `grantTypes` array stores OAuth flows that are allowed for a given client application.
	//
	// To know more about OAuth grant flows, see the
	// [ACP grant flows documentation](https://docs.authorization.cloudentity.com/features/oauth/grant_flows/).
	// Example: ["password","refresh_token","client_credentials","implicit","authorization_code"]
	GrantTypes []string `json:"grant_types"`

	// The time at which the request was issued by the TPP expressed
	// as "seconds since the epoch"
	Iat int64 `json:"iat,omitempty"`

	// Algorithm for signing ID tokens issued for a client application
	//
	// The default value depends on authorization server configuration.
	// Example: ES256
	// Enum: [RS256 ES256 PS256]
	IDTokenSignedResponseAlg string `json:"id_token_signed_response_alg,omitempty"`

	// Identifier for the TPP. This value must be unique for each TPP
	// registered by the issuer of the SSA.The value must be a Base62
	// encoded GUID. For SSAs issued by the OB Directory,
	// this must be the software_id.
	// Pattern: ^[0-9a-zA-Z]{1,22}$
	Iss string `json:"iss,omitempty"`

	// A unique identifier for the JWT. The value must be a UUIDv4 GUID.
	// Pattern: ^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$
	Jti string `json:"jti,omitempty"`

	// jwks
	Jwks *ClientJWKs `json:"jwks,omitempty"`

	// URL of JSON Web Key Set containing the public keys used by a client application to authenticate itself
	// with ACP
	JwksURI string `json:"jwks_uri,omitempty"`

	// Logo URI
	LogoURI string `json:"logo_uri,omitempty"`

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

	// A unique identifier string (e.g., a Universally Unique Identifier
	// (UUID)) assigned by the client developer or software publisher
	// The "software_id" SHOULD remain the same across
	// multiple updates or versions of the same piece of software.  The
	// value of this field is not intended to be human readable and is
	// usually opaque to the client and authorization server.
	SoftwareID string `json:"software_id,omitempty"`

	// A digitally signed or MACed JSON Web Token (JWT) [RFC7519] that
	// asserts metadata values about the client software.  In some cases,
	// a software statement is issued directly by the client
	// developer.  In other cases, a software statement is issued by
	// a third-party organization for use by the client developer.
	//
	// In both cases, the trust relationship the authorization server has
	// with the issuer of the software statement is intended to be used
	// as an input to the evaluation of whether the registration request
	// is accepted.
	//
	// A software statement can be presented to an
	// authorization server as part of the client registration request.
	SoftwareStatement string `json:"software_statement,omitempty"`

	// A version identifier string for the client software identified by
	// "software_id".  The value of the "software_version" SHOULD change
	// on any update to the client software identified by the same
	// "software_id".
	SoftwareVersion string `json:"software_version,omitempty"`

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

// Validate validates this openbanking u k dynamic client registration request
func (m *OpenbankingUKDynamicClientRegistrationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAud(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrantTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDTokenSignedResponseAlg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIss(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJti(formats); err != nil {
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

	if err := m.validateUserinfoSignedResponseAlg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var openbankingUKDynamicClientRegistrationRequestApplicationTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["single_page","server_web","mobile_desktop","service","legacy","dcr"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingUKDynamicClientRegistrationRequestApplicationTypesItemsEnum = append(openbankingUKDynamicClientRegistrationRequestApplicationTypesItemsEnum, v)
	}
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateApplicationTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openbankingUKDynamicClientRegistrationRequestApplicationTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateApplicationTypes(formats strfmt.Registry) error {
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

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateAud(formats strfmt.Registry) error {
	if swag.IsZero(m.Aud) { // not required
		return nil
	}

	if err := validate.Pattern("aud", "body", m.Aud, `^[0-9a-zA-Z]{1,18}$`); err != nil {
		return err
	}

	return nil
}

var openbankingUKDynamicClientRegistrationRequestGrantTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["authorization_code","implicit","client_credentials","refresh_token","password","urn:ietf:params:oauth:grant-type:jwt-bearer","urn:openid:params:grant-type:ciba"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingUKDynamicClientRegistrationRequestGrantTypesItemsEnum = append(openbankingUKDynamicClientRegistrationRequestGrantTypesItemsEnum, v)
	}
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateGrantTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openbankingUKDynamicClientRegistrationRequestGrantTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateGrantTypes(formats strfmt.Registry) error {
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

var openbankingUKDynamicClientRegistrationRequestTypeIDTokenSignedResponseAlgPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RS256","ES256","PS256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingUKDynamicClientRegistrationRequestTypeIDTokenSignedResponseAlgPropEnum = append(openbankingUKDynamicClientRegistrationRequestTypeIDTokenSignedResponseAlgPropEnum, v)
	}
}

const (

	// OpenbankingUKDynamicClientRegistrationRequestIDTokenSignedResponseAlgRS256 captures enum value "RS256"
	OpenbankingUKDynamicClientRegistrationRequestIDTokenSignedResponseAlgRS256 string = "RS256"

	// OpenbankingUKDynamicClientRegistrationRequestIDTokenSignedResponseAlgES256 captures enum value "ES256"
	OpenbankingUKDynamicClientRegistrationRequestIDTokenSignedResponseAlgES256 string = "ES256"

	// OpenbankingUKDynamicClientRegistrationRequestIDTokenSignedResponseAlgPS256 captures enum value "PS256"
	OpenbankingUKDynamicClientRegistrationRequestIDTokenSignedResponseAlgPS256 string = "PS256"
)

// prop value enum
func (m *OpenbankingUKDynamicClientRegistrationRequest) validateIDTokenSignedResponseAlgEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openbankingUKDynamicClientRegistrationRequestTypeIDTokenSignedResponseAlgPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateIDTokenSignedResponseAlg(formats strfmt.Registry) error {
	if swag.IsZero(m.IDTokenSignedResponseAlg) { // not required
		return nil
	}

	// value enum
	if err := m.validateIDTokenSignedResponseAlgEnum("id_token_signed_response_alg", "body", m.IDTokenSignedResponseAlg); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateIss(formats strfmt.Registry) error {
	if swag.IsZero(m.Iss) { // not required
		return nil
	}

	if err := validate.Pattern("iss", "body", m.Iss, `^[0-9a-zA-Z]{1,22}$`); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateJti(formats strfmt.Registry) error {
	if swag.IsZero(m.Jti) { // not required
		return nil
	}

	if err := validate.Pattern("jti", "body", m.Jti, `^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$`); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateJwks(formats strfmt.Registry) error {
	if swag.IsZero(m.Jwks) { // not required
		return nil
	}

	if m.Jwks != nil {
		if err := m.Jwks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jwks")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validatePrivacy(formats strfmt.Registry) error {
	if swag.IsZero(m.Privacy) { // not required
		return nil
	}

	if m.Privacy != nil {
		if err := m.Privacy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privacy")
			}
			return err
		}
	}

	return nil
}

var openbankingUKDynamicClientRegistrationRequestTypeRequestObjectSigningAlgPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["any","none","RS256","ES256","PS256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingUKDynamicClientRegistrationRequestTypeRequestObjectSigningAlgPropEnum = append(openbankingUKDynamicClientRegistrationRequestTypeRequestObjectSigningAlgPropEnum, v)
	}
}

const (

	// OpenbankingUKDynamicClientRegistrationRequestRequestObjectSigningAlgAny captures enum value "any"
	OpenbankingUKDynamicClientRegistrationRequestRequestObjectSigningAlgAny string = "any"

	// OpenbankingUKDynamicClientRegistrationRequestRequestObjectSigningAlgNone captures enum value "none"
	OpenbankingUKDynamicClientRegistrationRequestRequestObjectSigningAlgNone string = "none"

	// OpenbankingUKDynamicClientRegistrationRequestRequestObjectSigningAlgRS256 captures enum value "RS256"
	OpenbankingUKDynamicClientRegistrationRequestRequestObjectSigningAlgRS256 string = "RS256"

	// OpenbankingUKDynamicClientRegistrationRequestRequestObjectSigningAlgES256 captures enum value "ES256"
	OpenbankingUKDynamicClientRegistrationRequestRequestObjectSigningAlgES256 string = "ES256"

	// OpenbankingUKDynamicClientRegistrationRequestRequestObjectSigningAlgPS256 captures enum value "PS256"
	OpenbankingUKDynamicClientRegistrationRequestRequestObjectSigningAlgPS256 string = "PS256"
)

// prop value enum
func (m *OpenbankingUKDynamicClientRegistrationRequest) validateRequestObjectSigningAlgEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openbankingUKDynamicClientRegistrationRequestTypeRequestObjectSigningAlgPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateRequestObjectSigningAlg(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestObjectSigningAlg) { // not required
		return nil
	}

	// value enum
	if err := m.validateRequestObjectSigningAlgEnum("request_object_signing_alg", "body", m.RequestObjectSigningAlg); err != nil {
		return err
	}

	return nil
}

var openbankingUKDynamicClientRegistrationRequestResponseTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["token","id_token","code","code id_token","token id_token","token code","token id_token code"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingUKDynamicClientRegistrationRequestResponseTypesItemsEnum = append(openbankingUKDynamicClientRegistrationRequestResponseTypesItemsEnum, v)
	}
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateResponseTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openbankingUKDynamicClientRegistrationRequestResponseTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateResponseTypes(formats strfmt.Registry) error {
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

var openbankingUKDynamicClientRegistrationRequestTypeSubjectTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","pairwise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingUKDynamicClientRegistrationRequestTypeSubjectTypePropEnum = append(openbankingUKDynamicClientRegistrationRequestTypeSubjectTypePropEnum, v)
	}
}

const (

	// OpenbankingUKDynamicClientRegistrationRequestSubjectTypePublic captures enum value "public"
	OpenbankingUKDynamicClientRegistrationRequestSubjectTypePublic string = "public"

	// OpenbankingUKDynamicClientRegistrationRequestSubjectTypePairwise captures enum value "pairwise"
	OpenbankingUKDynamicClientRegistrationRequestSubjectTypePairwise string = "pairwise"
)

// prop value enum
func (m *OpenbankingUKDynamicClientRegistrationRequest) validateSubjectTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openbankingUKDynamicClientRegistrationRequestTypeSubjectTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateSubjectType(formats strfmt.Registry) error {
	if swag.IsZero(m.SubjectType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSubjectTypeEnum("subject_type", "body", m.SubjectType); err != nil {
		return err
	}

	return nil
}

var openbankingUKDynamicClientRegistrationRequestTypeTokenEndpointAuthMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["client_secret_basic","client_secret_post","client_secret_jwt","private_key_jwt","self_signed_tls_client_auth","tls_client_auth","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingUKDynamicClientRegistrationRequestTypeTokenEndpointAuthMethodPropEnum = append(openbankingUKDynamicClientRegistrationRequestTypeTokenEndpointAuthMethodPropEnum, v)
	}
}

const (

	// OpenbankingUKDynamicClientRegistrationRequestTokenEndpointAuthMethodClientSecretBasic captures enum value "client_secret_basic"
	OpenbankingUKDynamicClientRegistrationRequestTokenEndpointAuthMethodClientSecretBasic string = "client_secret_basic"

	// OpenbankingUKDynamicClientRegistrationRequestTokenEndpointAuthMethodClientSecretPost captures enum value "client_secret_post"
	OpenbankingUKDynamicClientRegistrationRequestTokenEndpointAuthMethodClientSecretPost string = "client_secret_post"

	// OpenbankingUKDynamicClientRegistrationRequestTokenEndpointAuthMethodClientSecretJwt captures enum value "client_secret_jwt"
	OpenbankingUKDynamicClientRegistrationRequestTokenEndpointAuthMethodClientSecretJwt string = "client_secret_jwt"

	// OpenbankingUKDynamicClientRegistrationRequestTokenEndpointAuthMethodPrivateKeyJwt captures enum value "private_key_jwt"
	OpenbankingUKDynamicClientRegistrationRequestTokenEndpointAuthMethodPrivateKeyJwt string = "private_key_jwt"

	// OpenbankingUKDynamicClientRegistrationRequestTokenEndpointAuthMethodSelfSignedTLSClientAuth captures enum value "self_signed_tls_client_auth"
	OpenbankingUKDynamicClientRegistrationRequestTokenEndpointAuthMethodSelfSignedTLSClientAuth string = "self_signed_tls_client_auth"

	// OpenbankingUKDynamicClientRegistrationRequestTokenEndpointAuthMethodTLSClientAuth captures enum value "tls_client_auth"
	OpenbankingUKDynamicClientRegistrationRequestTokenEndpointAuthMethodTLSClientAuth string = "tls_client_auth"

	// OpenbankingUKDynamicClientRegistrationRequestTokenEndpointAuthMethodNone captures enum value "none"
	OpenbankingUKDynamicClientRegistrationRequestTokenEndpointAuthMethodNone string = "none"
)

// prop value enum
func (m *OpenbankingUKDynamicClientRegistrationRequest) validateTokenEndpointAuthMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openbankingUKDynamicClientRegistrationRequestTypeTokenEndpointAuthMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateTokenEndpointAuthMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.TokenEndpointAuthMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateTokenEndpointAuthMethodEnum("token_endpoint_auth_method", "body", m.TokenEndpointAuthMethod); err != nil {
		return err
	}

	return nil
}

var openbankingUKDynamicClientRegistrationRequestTypeUserinfoSignedResponseAlgPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","RS256","ES256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openbankingUKDynamicClientRegistrationRequestTypeUserinfoSignedResponseAlgPropEnum = append(openbankingUKDynamicClientRegistrationRequestTypeUserinfoSignedResponseAlgPropEnum, v)
	}
}

const (

	// OpenbankingUKDynamicClientRegistrationRequestUserinfoSignedResponseAlgNone captures enum value "none"
	OpenbankingUKDynamicClientRegistrationRequestUserinfoSignedResponseAlgNone string = "none"

	// OpenbankingUKDynamicClientRegistrationRequestUserinfoSignedResponseAlgRS256 captures enum value "RS256"
	OpenbankingUKDynamicClientRegistrationRequestUserinfoSignedResponseAlgRS256 string = "RS256"

	// OpenbankingUKDynamicClientRegistrationRequestUserinfoSignedResponseAlgES256 captures enum value "ES256"
	OpenbankingUKDynamicClientRegistrationRequestUserinfoSignedResponseAlgES256 string = "ES256"
)

// prop value enum
func (m *OpenbankingUKDynamicClientRegistrationRequest) validateUserinfoSignedResponseAlgEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openbankingUKDynamicClientRegistrationRequestTypeUserinfoSignedResponseAlgPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) validateUserinfoSignedResponseAlg(formats strfmt.Registry) error {
	if swag.IsZero(m.UserinfoSignedResponseAlg) { // not required
		return nil
	}

	// value enum
	if err := m.validateUserinfoSignedResponseAlgEnum("userinfo_signed_response_alg", "body", m.UserinfoSignedResponseAlg); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this openbanking u k dynamic client registration request based on the context it is used
func (m *OpenbankingUKDynamicClientRegistrationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *OpenbankingUKDynamicClientRegistrationRequest) contextValidateApplicationTypes(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "application_types", "body", []string(m.ApplicationTypes)); err != nil {
		return err
	}

	return nil
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) contextValidateJwks(ctx context.Context, formats strfmt.Registry) error {

	if m.Jwks != nil {
		if err := m.Jwks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jwks")
			}
			return err
		}
	}

	return nil
}

func (m *OpenbankingUKDynamicClientRegistrationRequest) contextValidatePrivacy(ctx context.Context, formats strfmt.Registry) error {

	if m.Privacy != nil {
		if err := m.Privacy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privacy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenbankingUKDynamicClientRegistrationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenbankingUKDynamicClientRegistrationRequest) UnmarshalBinary(b []byte) error {
	var res OpenbankingUKDynamicClientRegistrationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
