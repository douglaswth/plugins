// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// adder endpoints
//
// Command:
// $ goa gen goa.design/plugins/security/examples/calc/adder/design

package addersvc

import (
	"context"

	goa "goa.design/goa"
)

// Endpoints wraps the "adder" service endpoints.
type Endpoints struct {
	Add goa.Endpoint
}

// NewEndpoints wraps the methods of the "adder" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Add: NewAddEndpoint(s),
	}
}

// Use applies the given middleware to all the "adder" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Add = m(e.Add)
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "adder".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AddPayload)
		return s.Add(ctx, p)
	}
}
