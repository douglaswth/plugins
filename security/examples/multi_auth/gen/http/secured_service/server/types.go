// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// secured_service HTTP server types
//
// Command:
// $ goa gen goa.design/plugins/security/examples/multi_auth/design

package server

import (
	securedservice "goa.design/plugins/security/examples/multi_auth/gen/secured_service"
)

// SigninRequestBody is the type of the "secured_service" service "signin"
// endpoint HTTP request body.
type SigninRequestBody struct {
	// Username used to perform signin
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	// Username used to perform signin
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// SecureRequestBody is the type of the "secured_service" service "secure"
// endpoint HTTP request body.
type SecureRequestBody struct {
	// JWT used for authentication
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// DoublySecureRequestBody is the type of the "secured_service" service
// "doubly_secure" endpoint HTTP request body.
type DoublySecureRequestBody struct {
	// JWT used for authentication
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// AlsoDoublySecureRequestBody is the type of the "secured_service" service
// "also_doubly_secure" endpoint HTTP request body.
type AlsoDoublySecureRequestBody struct {
	// Username used to perform signin
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	// Username used to perform signin
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// JWT used for authentication
	Token      *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
	OauthToken *string `form:"oauth_token,omitempty" json:"oauth_token,omitempty" xml:"oauth_token,omitempty"`
}

// SigninUnauthorizedResponseBody is the type of the "secured_service" service
// "signin" endpoint HTTP response body for the "unauthorized" error.
type SigninUnauthorizedResponseBody string

// SecureUnauthorizedResponseBody is the type of the "secured_service" service
// "secure" endpoint HTTP response body for the "unauthorized" error.
type SecureUnauthorizedResponseBody string

// DoublySecureUnauthorizedResponseBody is the type of the "secured_service"
// service "doubly_secure" endpoint HTTP response body for the "unauthorized"
// error.
type DoublySecureUnauthorizedResponseBody string

// AlsoDoublySecureUnauthorizedResponseBody is the type of the
// "secured_service" service "also_doubly_secure" endpoint HTTP response body
// for the "unauthorized" error.
type AlsoDoublySecureUnauthorizedResponseBody string

// NewSigninUnauthorizedResponseBody builds the HTTP response body from the
// result of the "signin" endpoint of the "secured_service" service.
func NewSigninUnauthorizedResponseBody(res securedservice.Unauthorized) SigninUnauthorizedResponseBody {
	body := SigninUnauthorizedResponseBody(res)
	return body
}

// NewSecureUnauthorizedResponseBody builds the HTTP response body from the
// result of the "secure" endpoint of the "secured_service" service.
func NewSecureUnauthorizedResponseBody(res securedservice.Unauthorized) SecureUnauthorizedResponseBody {
	body := SecureUnauthorizedResponseBody(res)
	return body
}

// NewDoublySecureUnauthorizedResponseBody builds the HTTP response body from
// the result of the "doubly_secure" endpoint of the "secured_service" service.
func NewDoublySecureUnauthorizedResponseBody(res securedservice.Unauthorized) DoublySecureUnauthorizedResponseBody {
	body := DoublySecureUnauthorizedResponseBody(res)
	return body
}

// NewAlsoDoublySecureUnauthorizedResponseBody builds the HTTP response body
// from the result of the "also_doubly_secure" endpoint of the
// "secured_service" service.
func NewAlsoDoublySecureUnauthorizedResponseBody(res securedservice.Unauthorized) AlsoDoublySecureUnauthorizedResponseBody {
	body := AlsoDoublySecureUnauthorizedResponseBody(res)
	return body
}

// NewSigninSigninPayload builds a secured_service service signin endpoint
// payload.
func NewSigninSigninPayload(body *SigninRequestBody) *securedservice.SigninPayload {
	v := &securedservice.SigninPayload{
		Username: body.Username,
		Password: body.Password,
	}
	return v
}

// NewSecureSecurePayload builds a secured_service service secure endpoint
// payload.
func NewSecureSecurePayload(body *SecureRequestBody, fail *bool) *securedservice.SecurePayload {
	v := &securedservice.SecurePayload{
		Token: body.Token,
	}
	v.Fail = fail
	return v
}

// NewDoublySecureDoublySecurePayload builds a secured_service service
// doubly_secure endpoint payload.
func NewDoublySecureDoublySecurePayload(body *DoublySecureRequestBody, key *string) *securedservice.DoublySecurePayload {
	v := &securedservice.DoublySecurePayload{
		Token: body.Token,
	}
	v.Key = key
	return v
}

// NewAlsoDoublySecureAlsoDoublySecurePayload builds a secured_service service
// also_doubly_secure endpoint payload.
func NewAlsoDoublySecureAlsoDoublySecurePayload(body *AlsoDoublySecureRequestBody, key *string) *securedservice.AlsoDoublySecurePayload {
	v := &securedservice.AlsoDoublySecurePayload{
		Username:   body.Username,
		Password:   body.Password,
		Token:      body.Token,
		OauthToken: body.OauthToken,
	}
	v.Key = key
	return v
}
