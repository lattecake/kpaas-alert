package http

import (
	"context"
	"encoding/json"
	"errors"
	http1 "github.com/go-kit/kit/transport/http"
	"github.com/lattecake/kpaas-alert/pkg/endpoint"
	"net/http"
)

// makeAlertHandler creates the handler logic
func makeAlertHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/alert", http1.NewServer(endpoints.AlertEndpoint, decodeAlertRequest, encodeAlertResponse, options...))
}

// decodeAlertRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAlertRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.AlertRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAlertResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAlertResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetAlertHandler creates the handler logic
func makeGetAlertHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-alert", http1.NewServer(endpoints.GetAlertEndpoint, decodeGetAlertRequest, encodeGetAlertResponse, options...))
}

// decodeGetAlertRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetAlertRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetAlertRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetAlertResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetAlertResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	_ = json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
