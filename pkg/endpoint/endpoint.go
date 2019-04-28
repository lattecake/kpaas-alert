package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/lattecake/kpaas-alert/pkg/service"
)

// AlertRequest collects the request parameters for the Alert method.
type AlertRequest struct {
	Data string `json:"data"`
}

// AlertResponse collects the response parameters for the Alert method.
type AlertResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeAlertEndpoint returns an endpoint that invokes Alert on the service.
func MakeAlertEndpoint(s service.KpaasAlertService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AlertRequest)
		rs, err := s.Alert(ctx, req.Data)
		return AlertResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r AlertResponse) Failed() error {
	return r.Err
}

// GetAlertRequest collects the request parameters for the GetAlert method.
type GetAlertRequest struct {
	Id int64 `json:"id"`
}

// GetAlertResponse collects the response parameters for the GetAlert method.
type GetAlertResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeGetAlertEndpoint returns an endpoint that invokes GetAlert on the service.
func MakeGetAlertEndpoint(s service.KpaasAlertService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAlertRequest)
		rs, err := s.GetAlert(ctx, req.Id)
		return GetAlertResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAlertResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Alert implements Service. Primarily useful in a client.
func (e Endpoints) Alert(ctx context.Context, data string) (rs string, err error) {
	request := AlertRequest{Data: data}
	response, err := e.AlertEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AlertResponse).Rs, response.(AlertResponse).Err
}

// GetAlert implements Service. Primarily useful in a client.
func (e Endpoints) GetAlert(ctx context.Context, id int64) (rs string, err error) {
	request := GetAlertRequest{Id: id}
	response, err := e.GetAlertEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAlertResponse).Rs, response.(GetAlertResponse).Err
}
