package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pentops/log.go/log"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

// AuthHeaders translates a request into headers to pass on to the remote server
// Errors which implement gRPC status will be returned to the client as HTTP
// errors, otherwise 500 with a log line
type AuthHeaders interface {
	AuthHeaders(context.Context, *http.Request) (map[string]string, error)
}

type AuthHeadersFunc func(context.Context, *http.Request) (map[string]string, error)

func (f AuthHeadersFunc) AuthHeaders(ctx context.Context, r *http.Request) (map[string]string, error) {
	return f(ctx, r)
}

type Invoker interface {
	// Invoke is desined for gRPC ClientConn.Invoke, the two interfaces should
	// be protos...
	Invoke(context.Context, string, interface{}, interface{}, ...grpc.CallOption) error
}

type Codec interface {
	ToProto(body []byte, msg protoreflect.Message) error
	FromProto(msg protoreflect.Message) ([]byte, error)
}

type Router struct {
	router                 *mux.Router
	ForwardResponseHeaders map[string]bool
	ForwardRequestHeaders  map[string]bool
	Codec                  Codec

	middleware []func(http.Handler) http.Handler
}

func NewRouter(codec Codec) *Router {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"not found"}` + "\n")) // nolint: errcheck
		log.WithField(r.Context(), "httpURL", r.URL.String()).Debug("Not found")
	})

	return &Router{
		router: router,
		ForwardResponseHeaders: map[string]bool{
			"set-cookie": true,
			"x-version":  true,
		},
		ForwardRequestHeaders: map[string]bool{
			"cookie": true,
			"origin": true,
		},
		Codec: codec,
	}
}

func (rr *Router) Use(middleware func(http.Handler) http.Handler) {
	rr.middleware = append(rr.middleware, middleware)
}

func (rr *Router) SetNotFoundHandler(handler http.Handler) {
	rr.router.NotFoundHandler = handler
}

func (rr *Router) HealthCheck(path string, callback func() error) {
	rr.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if err := callback(); err != nil {
			doError(r.Context(), w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}` + "\n")) // nolint: errcheck
	})
}

func (rr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler = rr.router
	for _, mw := range rr.middleware {
		handler = mw(handler)
	}
	handler.ServeHTTP(w, r)
}

func (rr *Router) StaticJSON(path string, document interface{}) error {
	jb, err := json.Marshal(document)
	if err != nil {
		return err
	}

	rr.router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jb) // nolint: errcheck
	})
	return nil
}

// RegisterService calls RegisterGRPCMethod on all methods of the service with
// default config.
func (rr *Router) RegisterService(ctx context.Context, ss protoreflect.ServiceDescriptor, conn Invoker) error {
	methods := ss.Methods()
	for ii := 0; ii < methods.Len(); ii++ {
		method := methods.Get(ii)
		if err := rr.RegisterGRPCMethod(ctx, GRPCMethodConfig{
			Method:  method,
			Invoker: conn,
		}); err != nil {
			return err
		}
	}
	return nil
}

type GRPCMethodConfig struct {
	AuthHeaders AuthHeaders
	Invoker     Invoker
	Method      protoreflect.MethodDescriptor
}

func (rr *Router) RegisterGRPCMethod(ctx context.Context, config GRPCMethodConfig) error {
	handler, err := rr.buildMethod(config)
	if err != nil {
		return err
	}
	rr.router.Methods(handler.HTTPMethod).Path(handler.HTTPPath).Handler(handler)
	log.WithFields(ctx, map[string]interface{}{
		"method": handler.HTTPMethod,
		"path":   handler.HTTPPath,
		"grpc":   handler.FullName,
	}).Info("Registered HTTP Method")
	return nil
}

func (rr *Router) buildMethod(config GRPCMethodConfig) (*grpcMethod, error) {
	serviceName := config.Method.Parent().(protoreflect.ServiceDescriptor).FullName()
	methodOptions := config.Method.Options().(*descriptorpb.MethodOptions)
	httpOpt := proto.GetExtension(methodOptions, annotations.E_Http).(*annotations.HttpRule)

	var httpMethod string
	var httpPath string

	switch pt := httpOpt.Pattern.(type) {
	case *annotations.HttpRule_Get:
		httpMethod = "GET"
		httpPath = pt.Get
	case *annotations.HttpRule_Post:
		httpMethod = "POST"
		httpPath = pt.Post
	case *annotations.HttpRule_Put:
		httpMethod = "PUT"
		httpPath = pt.Put
	case *annotations.HttpRule_Delete:
		httpMethod = "DELETE"
		httpPath = pt.Delete
	case *annotations.HttpRule_Patch:
		httpMethod = "PATCH"
		httpPath = pt.Patch

	default:
		return nil, fmt.Errorf("unsupported http method %T", pt)
	}

	handler := &grpcMethod{
		// the 'FullName' method of MethodDescriptor returns this in the wrong format, i.e. all dots.
		FullName:               fmt.Sprintf("/%s/%s", serviceName, config.Method.Name()),
		Input:                  config.Method.Input(),
		Output:                 config.Method.Output(),
		Invoker:                config.Invoker,
		HTTPMethod:             httpMethod,
		HTTPPath:               httpPath,
		ForwardResponseHeaders: rr.ForwardResponseHeaders,
		ForwardRequestHeaders:  rr.ForwardRequestHeaders,
		Codec:                  rr.Codec,
		authHeaders:            config.AuthHeaders,
	}

	return handler, nil

}

type grpcMethod struct {
	FullName               string
	Input                  protoreflect.MessageDescriptor
	Output                 protoreflect.MessageDescriptor
	Invoker                Invoker
	HTTPMethod             string
	HTTPPath               string
	ForwardResponseHeaders map[string]bool
	ForwardRequestHeaders  map[string]bool
	Codec                  Codec
	authHeaders            AuthHeaders
}

func (mm *grpcMethod) mapRequest(r *http.Request) (protoreflect.Message, error) {
	inputMessage := dynamicpb.NewMessage(mm.Input)
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if len(reqBody) > 0 {
		if err := mm.Codec.ToProto(reqBody, inputMessage); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	reqVars := mux.Vars(r)
	for key, provided := range reqVars {
		fd := mm.Input.Fields().ByName(protoreflect.Name(key))
		if err := setFieldFromString(mm.Codec, inputMessage, fd, provided); err != nil {
			return nil, err
		}
	}

	query := r.URL.Query()
	for key, values := range query {
		if err := setFieldFromStrings(mm.Codec, inputMessage, key, values); err != nil {
			return nil, err
		}
	}

	return inputMessage, nil
}

func (mm *grpcMethod) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = log.WithFields(ctx, map[string]interface{}{
		"httpMethod": r.Method,
		"httpURL":    r.URL.String(),
		"gRPCMethod": mm.FullName,
	})

	inputMessage, err := mm.mapRequest(r)
	if err != nil {
		doUserError(ctx, w, err)
		return
	}

	outputMessage := dynamicpb.NewMessage(mm.Output)

	md := map[string]string{}
	for key, v := range r.Header {
		key = strings.ToLower(key)
		if !mm.ForwardRequestHeaders[key] {
			continue
		}
		md[key] = v[0] // only one value in gRPC
	}
	ctx = log.WithField(ctx, "passthroughHeaders", md)

	if mm.authHeaders != nil {
		authHeaders, err := mm.authHeaders.AuthHeaders(ctx, r)
		if err != nil {
			doUserError(ctx, w, err)
			return
		}
		for key, val := range authHeaders {
			md[key] = val

		}
		ctx = log.WithField(ctx, "authHeaders", authHeaders)
	}

	// Send request header
	ctx = metadata.NewOutgoingContext(ctx, metadata.New(md))

	// Receive response header
	var responseHeader metadata.MD

	err = mm.Invoker.Invoke(ctx, mm.FullName, inputMessage, outputMessage, grpc.Header(&responseHeader))
	if err != nil {
		doUserError(ctx, w, err)
		return
	}

	bytesOut, err := mm.Codec.FromProto(outputMessage)
	if err != nil {
		doError(ctx, w, err)
		return
	}

	headerOut := w.Header()
	headerOut.Set("Content-Type", "application/json")

	for key, vals := range responseHeader {
		key = strings.ToLower(key)
		if !mm.ForwardResponseHeaders[key] {
			continue
		}
		for _, val := range vals {
			w.Header().Add(key, val)
		}
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(bytesOut); err != nil {
		log.WithError(ctx, err).Error("Failed to write response")
		return
	}
	log.Info(ctx, "Request completed")
}

func doUserError(ctx context.Context, w http.ResponseWriter, err error) {
	// TODO: Handle specific gRPC trailer type errors
	if statusError, isStatusError := status.FromError(err); isStatusError {
		log.WithField(ctx, "httpError", statusError).Info("User error")
		doStatusError(ctx, w, statusError)
		return
	}
	doError(ctx, w, err)
}

func doError(ctx context.Context, w http.ResponseWriter, err error) {
	log.WithError(ctx, err).Error("Error handling request")
	body := map[string]string{
		"error": err.Error(),
	}
	bytesOut, err := json.Marshal(body)
	if err != nil {
		log.WithError(ctx, err).Error("Failed to marshal error response")
		http.Error(w, `{"error":"meta error marshalling error"}`, http.StatusInternalServerError)
		return
	}
	headerOut := w.Header()
	headerOut.Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := w.Write(bytesOut); err != nil {
		log.WithError(ctx, err).Error("Failed to write error response")
		return
	}
}

func doStatusError(ctx context.Context, w http.ResponseWriter, statusError *status.Status) {
	bytesOut, err := json.Marshal(map[string]string{
		"error": statusError.Message(),
	})

	if err != nil {
		log.WithError(ctx, err).Error("Failed to marshal error response")
		http.Error(w, `{"error":"meta error marshalling error"}`, http.StatusInternalServerError)
		return
	}

	headerOut := w.Header()
	headerOut.Set("Content-Type", "application/json")

	httpStatus, ok := statusToHTTPCode[statusError.Code()]
	if !ok {
		httpStatus = http.StatusInternalServerError
	}
	w.WriteHeader(httpStatus)

	if _, err := w.Write(bytesOut); err != nil {
		log.WithError(ctx, err).Error("Failed to write error response")
		return
	}
}

var statusToHTTPCode = map[codes.Code]int{
	// TODO: These were autocompleted by AI, check if they are correct
	codes.OK:                 http.StatusOK,
	codes.Canceled:           http.StatusRequestTimeout,
	codes.Unknown:            http.StatusInternalServerError,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.DeadlineExceeded:   http.StatusGatewayTimeout,
	codes.NotFound:           http.StatusNotFound,
	codes.AlreadyExists:      http.StatusConflict,
	codes.PermissionDenied:   http.StatusForbidden,
	codes.ResourceExhausted:  http.StatusTooManyRequests,
	codes.FailedPrecondition: http.StatusPreconditionFailed,
	codes.Aborted:            http.StatusConflict,
	codes.OutOfRange:         http.StatusBadRequest,
	codes.Unimplemented:      http.StatusNotImplemented,
	codes.Internal:           http.StatusInternalServerError,
	codes.Unavailable:        http.StatusServiceUnavailable,
	codes.DataLoss:           http.StatusInternalServerError,
	codes.Unauthenticated:    http.StatusUnauthorized,
}
