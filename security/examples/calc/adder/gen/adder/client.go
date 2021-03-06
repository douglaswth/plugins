// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// adder client
//
// Command:
// $ goa gen goa.design/plugins/security/examples/calc/adder/design

package addersvc

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "adder" service client.
type Client struct {
	AddEndpoint goa.Endpoint
}

// NewClient initializes a "adder" service client given the endpoints.
func NewClient(add goa.Endpoint) *Client {
	return &Client{
		AddEndpoint: add,
	}
}

// Add calls the "add" endpoint of the "adder" service.
// Add can return the following error types:
//	- Unauthorized
//	- InvalidScopes
//	- error: generic transport error.
func (c *Client) Add(ctx context.Context, p *AddPayload) (res int, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}
