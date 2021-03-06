// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// adder service
//
// Command:
// $ goa gen goa.design/plugins/security/examples/calc/adder/design

package addersvc

import (
	"context"
)

// The adder service exposes an add method secured via API keys.
type Service interface {
	// This action returns the sum of two integers and is secured with the API key
	// scheme
	Add(context.Context, *AddPayload) (int, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "adder"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"add"}

// AddPayload is the payload type of the adder service add method.
type AddPayload struct {
	// API key
	Key string
	// Left operand
	A int
	// Right operand
	B int
}

type Unauthorized string

type InvalidScopes string

// Error returns "unauthorized".
func (e Unauthorized) Error() string {
	return "unauthorized"
}

// Error returns "invalid-scopes".
func (e InvalidScopes) Error() string {
	return "invalid-scopes"
}
