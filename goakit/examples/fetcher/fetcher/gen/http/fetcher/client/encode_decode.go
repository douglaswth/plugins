// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// fetcher HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/fetcher/fetcher/design

package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/http"
	fetchersvc "goa.design/plugins/goakit/examples/fetcher/fetcher/gen/fetcher"
)

// BuildFetchRequest instantiates a HTTP request object with method and path
// set to call the "fetcher" service "fetch" endpoint
func (c *Client) BuildFetchRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		url_ string
	)
	{
		p, ok := v.(*fetchersvc.FetchPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("fetcher", "fetch", "*fetchersvc.FetchPayload", v)
		}
		url_ = p.URL
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: FetchFetcherPath(url_)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("fetcher", "fetch", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeFetchResponse returns a decoder for responses returned by the fetcher
// fetch endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeFetchResponse may return the following error types:
//	- *fetchersvc.Error: http.StatusBadRequest
//	- *fetchersvc.Error: http.StatusInternalServerError
//	- error: generic transport error.
func DecodeFetchResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body FetchResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("fetcher", "fetch", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, fmt.Errorf("invalid response: %s", err)
			}

			return NewFetchFetchMediaOK(&body), nil
		case http.StatusBadRequest:
			var (
				body FetchBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("fetcher", "fetch", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, fmt.Errorf("invalid response: %s", err)
			}

			return nil, NewFetchBadRequest(&body)
		case http.StatusInternalServerError:
			var (
				body FetchInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("fetcher", "fetch", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, fmt.Errorf("invalid response: %s", err)
			}

			return nil, NewFetchInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("account", "create", resp.StatusCode, string(body))
		}
	}
}
