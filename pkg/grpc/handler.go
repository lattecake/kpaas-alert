package grpc

import (
	"context"
	"errors"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/lattecake/kpaas-alert/pkg/endpoint"
	"github.com/lattecake/kpaas-alert/pkg/grpc/pb"
	context1 "golang.org/x/net/context"
)

// makeAlertHandler creates the handler logic
func makeAlertHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AlertEndpoint, decodeAlertRequest, encodeAlertResponse, options...)
}

// decodeAlertResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Alert request.
// TODO implement the decoder
func decodeAlertRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'KpaasAlert' Decoder is not impelemented")
}

// encodeAlertResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeAlertResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'KpaasAlert' Encoder is not impelemented")
}
func (g *grpcServer) Alert(ctx context1.Context, req *pb.AlertRequest) (*pb.AlertReply, error) {
	_, rep, err := g.alert.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AlertReply), nil
}

// makeGetAlertHandler creates the handler logic
func makeGetAlertHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetAlertEndpoint, decodeGetAlertRequest, encodeGetAlertResponse, options...)
}

// decodeGetAlertResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetAlert request.
// TODO implement the decoder
func decodeGetAlertRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'KpaasAlert' Decoder is not impelemented")
}

// encodeGetAlertResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetAlertResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'KpaasAlert' Encoder is not impelemented")
}
func (g *grpcServer) GetAlert(ctx context1.Context, req *pb.GetAlertRequest) (*pb.GetAlertReply, error) {
	_, rep, err := g.getAlert.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetAlertReply), nil
}
