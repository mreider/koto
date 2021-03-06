// Code generated by protoc-gen-twirp v7.1.0, DO NOT EDIT.
// source: user.proto

package rpc

import bytes "bytes"
import strings "strings"
import context "context"
import fmt "fmt"
import ioutil "io/ioutil"
import http "net/http"
import strconv "strconv"

import jsonpb "github.com/golang/protobuf/jsonpb"
import proto "github.com/golang/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the twirp package used in your project.
// A compilation error at this line likely means your copy of the
// twirp package needs to be updated.
const _ = twirp.TwirpPackageIsVersion7

// =====================
// UserService Interface
// =====================

type UserService interface {
	BlockUser(context.Context, *UserBlockRequest) (*Empty, error)
}

// ===========================
// UserService Protobuf Client
// ===========================

type userServiceProtobufClient struct {
	client      HTTPClient
	urls        [1]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewUserServiceProtobufClient creates a Protobuf client that implements the UserService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewUserServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) UserService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(clientOpts.PathPrefix(), "rpc", "UserService")
	urls := [1]string{
		serviceURL + "BlockUser",
	}

	return &userServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *userServiceProtobufClient) BlockUser(ctx context.Context, in *UserBlockRequest) (*Empty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "BlockUser")
	caller := c.callBlockUser
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *UserBlockRequest) (*Empty, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UserBlockRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UserBlockRequest) when calling interceptor")
					}
					return c.callBlockUser(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *userServiceProtobufClient) callBlockUser(ctx context.Context, in *UserBlockRequest) (*Empty, error) {
	out := new(Empty)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// =======================
// UserService JSON Client
// =======================

type userServiceJSONClient struct {
	client      HTTPClient
	urls        [1]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewUserServiceJSONClient creates a JSON client that implements the UserService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewUserServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) UserService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(clientOpts.PathPrefix(), "rpc", "UserService")
	urls := [1]string{
		serviceURL + "BlockUser",
	}

	return &userServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *userServiceJSONClient) BlockUser(ctx context.Context, in *UserBlockRequest) (*Empty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "BlockUser")
	caller := c.callBlockUser
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *UserBlockRequest) (*Empty, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UserBlockRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UserBlockRequest) when calling interceptor")
					}
					return c.callBlockUser(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *userServiceJSONClient) callBlockUser(ctx context.Context, in *UserBlockRequest) (*Empty, error) {
	out := new(Empty)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ==========================
// UserService Server Handler
// ==========================

type userServiceServer struct {
	UserService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
}

// NewUserServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewUserServiceServer(svc UserService, opts ...interface{}) TwirpServer {
	serverOpts := twirp.ServerOptions{}
	for _, opt := range opts {
		switch o := opt.(type) {
		case twirp.ServerOption:
			o(&serverOpts)
		case *twirp.ServerHooks: // backwards compatibility, allow to specify hooks as an argument
			twirp.WithServerHooks(o)(&serverOpts)
		case nil: // backwards compatibility, allow nil value for the argument
			continue
		default:
			panic(fmt.Sprintf("Invalid option type %T on NewUserServiceServer", o))
		}
	}

	return &userServiceServer{
		UserService:      svc,
		pathPrefix:       serverOpts.PathPrefix(),
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		hooks:            serverOpts.Hooks,
		jsonSkipDefaults: serverOpts.JSONSkipDefaults,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *userServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// UserServicePathPrefix is a convenience constant that could used to identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// that add a "/twirp" prefix by default, and use CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const UserServicePathPrefix = "/twirp/rpc.UserService/"

func (s *userServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Verify path format: [<prefix>]/<package>.<Service>/<Method>
	prefix, pkgService, method := parseTwirpPath(req.URL.Path)
	if pkgService != "rpc.UserService" {
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
	if prefix != s.pathPrefix {
		msg := fmt.Sprintf("invalid path prefix %q, expected %q, on path %q", prefix, s.pathPrefix, req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	switch method {
	case "BlockUser":
		s.serveBlockUser(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *userServiceServer) serveBlockUser(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveBlockUserJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveBlockUserProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *userServiceServer) serveBlockUserJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "BlockUser")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(UserBlockRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	handler := s.UserService.BlockUser
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *UserBlockRequest) (*Empty, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UserBlockRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UserBlockRequest) when calling interceptor")
					}
					return s.UserService.BlockUser(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *Empty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Empty and nil error while calling BlockUser. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true, EmitDefaults: !s.jsonSkipDefaults}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) serveBlockUserProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "BlockUser")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(UserBlockRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.UserService.BlockUser
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *UserBlockRequest) (*Empty, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UserBlockRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UserBlockRequest) when calling interceptor")
					}
					return s.UserService.BlockUser(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *Empty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Empty and nil error while calling BlockUser. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor5, 0
}

func (s *userServiceServer) ProtocGenTwirpVersion() string {
	return "v7.1.0"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *userServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "rpc", "UserService")
}

var twirpFileDescriptor5 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x96, 0xe2, 0xce, 0xcd, 0x4f,
	0x49, 0xcd, 0x81, 0x88, 0x28, 0x69, 0x73, 0x09, 0x84, 0x16, 0xa7, 0x16, 0x39, 0xe5, 0xe4, 0x27,
	0x67, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x89, 0x73, 0xb1, 0x83, 0xf4, 0xc4, 0x67,
	0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xb1, 0x81, 0xb8, 0x9e, 0x29, 0x46, 0xb6, 0x5c,
	0xdc, 0x20, 0xc5, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x7a, 0x5c, 0x9c, 0x60, 0x7d,
	0x20, 0x31, 0x21, 0x51, 0xbd, 0xa2, 0x82, 0x64, 0x3d, 0x74, 0xb3, 0xa4, 0xb8, 0xc0, 0xc2, 0xae,
	0xb9, 0x05, 0x25, 0x95, 0x4e, 0x1c, 0x51, 0x6c, 0x7a, 0x7a, 0xfa, 0x45, 0x05, 0xc9, 0x49, 0x6c,
	0x60, 0xcb, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x10, 0x52, 0xc9, 0x77, 0x9c, 0x00, 0x00,
	0x00,
}
