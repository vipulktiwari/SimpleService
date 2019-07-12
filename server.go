package main

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
)

// Our basic service, which will have all the api as interface funtions
type BasicService interface {
	Ping(context.Context) (string, error)
}

type basicService struct{}

func (b basicService) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makePingEndpoint(svc BasicService) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
		resp, err = svc.Ping(ctx)
		return resp, err
	}
}

func decodePingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	// as ping is a get request we will do nothing here
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	svc := basicService{}
	pinghandler := httptransport.NewServer(
		makePingEndpoint(svc),
		decodePingRequest,
		encodeResponse,
	)

	log.Println("server started, listening on port 8080")
	http.Handle("/ping", pinghandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
