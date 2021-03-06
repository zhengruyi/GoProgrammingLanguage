package transport

import (
	"GoProgrammingLanguage/Go高并发与微服务实战/discovery/endpoint"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

var ErrorBadRequest = errors.New("invalid request parameter")

func MakeHttpHandler(ctx context.Context, endpoints endpoint.DiscoveryEndpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []kithttp.ServerOption {
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}
	r.Methods("GET").Path("/say-hello").Handler(kithttp.NewServer(
		endpoints.SayHelloEndpoint,
		decodeSayHelloRequest,
		encodeJsonResponse,
		options...,
		))

	r.Methods("GET").Path("/discovery").Handler(kithttp.NewServer(
		endpoints.DiscoveryEndpoint,
		decodeDiscoveryRequest,
		encodeJsonResponse,
		options...,
	))

	r.Methods("GET").Path("/health").Handler(kithttp.NewServer(
		endpoints.HealthCheckEndpoint,
		decodeHealthCheckRequest,
		encodeJsonResponse,
		options...,
	))
	return r
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type","application/json;charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error" : err.Error(),
	})
}

func decodeHealthCheckRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	return endpoint.HealthRequest{},nil
}

func decodeDiscoveryRequest(ctx context.Context, request2 *http.Request) (request interface{}, err error) {
	serviceName := request2.URL.Query().Get("serviceName")
	if serviceName == "" {
		return nil, ErrorBadRequest
	}
	return endpoint.DiscoveryRequest{
		ServiceName : serviceName,
	},nil
}

func encodeJsonResponse(ctx context.Context, writer http.ResponseWriter, i interface{}) error {
	writer.Header().Set("Content-Type","application/json;charset=utf-8")
	return json.NewEncoder(writer).Encode(i)
}

func decodeSayHelloRequest(ctx context.Context, request2 *http.Request) (request interface{}, err error) {
	return endpoint.SayHelloRequest{},nil
}


