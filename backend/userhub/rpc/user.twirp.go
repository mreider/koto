// Code generated by protoc-gen-twirp v5.12.0, DO NOT EDIT.
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

// =====================
// UserService Interface
// =====================

type UserService interface {
	Friends(context.Context, *Empty) (*UserFriendsResponse, error)

	FriendsOfFriends(context.Context, *Empty) (*UserFriendsOfFriendsResponse, error)

	Me(context.Context, *Empty) (*UserMeResponse, error)

	EditProfile(context.Context, *UserEditProfileRequest) (*Empty, error)

	Users(context.Context, *UserUsersRequest) (*UserUsersResponse, error)
}

// ===========================
// UserService Protobuf Client
// ===========================

type userServiceProtobufClient struct {
	client HTTPClient
	urls   [5]string
	opts   twirp.ClientOptions
}

// NewUserServiceProtobufClient creates a Protobuf client that implements the UserService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewUserServiceProtobufClient(addr string, client HTTPClient, opts ...twirp.ClientOption) UserService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + UserServicePathPrefix
	urls := [5]string{
		prefix + "Friends",
		prefix + "FriendsOfFriends",
		prefix + "Me",
		prefix + "EditProfile",
		prefix + "Users",
	}

	return &userServiceProtobufClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *userServiceProtobufClient) Friends(ctx context.Context, in *Empty) (*UserFriendsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "Friends")
	out := new(UserFriendsResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
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

func (c *userServiceProtobufClient) FriendsOfFriends(ctx context.Context, in *Empty) (*UserFriendsOfFriendsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "FriendsOfFriends")
	out := new(UserFriendsOfFriendsResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
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

func (c *userServiceProtobufClient) Me(ctx context.Context, in *Empty) (*UserMeResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "Me")
	out := new(UserMeResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
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

func (c *userServiceProtobufClient) EditProfile(ctx context.Context, in *UserEditProfileRequest) (*Empty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "EditProfile")
	out := new(Empty)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[3], in, out)
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

func (c *userServiceProtobufClient) Users(ctx context.Context, in *UserUsersRequest) (*UserUsersResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "Users")
	out := new(UserUsersResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[4], in, out)
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
	client HTTPClient
	urls   [5]string
	opts   twirp.ClientOptions
}

// NewUserServiceJSONClient creates a JSON client that implements the UserService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewUserServiceJSONClient(addr string, client HTTPClient, opts ...twirp.ClientOption) UserService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + UserServicePathPrefix
	urls := [5]string{
		prefix + "Friends",
		prefix + "FriendsOfFriends",
		prefix + "Me",
		prefix + "EditProfile",
		prefix + "Users",
	}

	return &userServiceJSONClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *userServiceJSONClient) Friends(ctx context.Context, in *Empty) (*UserFriendsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "Friends")
	out := new(UserFriendsResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
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

func (c *userServiceJSONClient) FriendsOfFriends(ctx context.Context, in *Empty) (*UserFriendsOfFriendsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "FriendsOfFriends")
	out := new(UserFriendsOfFriendsResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
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

func (c *userServiceJSONClient) Me(ctx context.Context, in *Empty) (*UserMeResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "Me")
	out := new(UserMeResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
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

func (c *userServiceJSONClient) EditProfile(ctx context.Context, in *UserEditProfileRequest) (*Empty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "EditProfile")
	out := new(Empty)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[3], in, out)
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

func (c *userServiceJSONClient) Users(ctx context.Context, in *UserUsersRequest) (*UserUsersResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "Users")
	out := new(UserUsersResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[4], in, out)
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
	hooks *twirp.ServerHooks
}

func NewUserServiceServer(svc UserService, hooks *twirp.ServerHooks) TwirpServer {
	return &userServiceServer{
		UserService: svc,
		hooks:       hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *userServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// UserServicePathPrefix is used for all URL paths on a twirp UserService server.
// Requests are always: POST UserServicePathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const UserServicePathPrefix = "/rpc.UserService/"

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
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}

	switch req.URL.Path {
	case "/rpc.UserService/Friends":
		s.serveFriends(ctx, resp, req)
		return
	case "/rpc.UserService/FriendsOfFriends":
		s.serveFriendsOfFriends(ctx, resp, req)
		return
	case "/rpc.UserService/Me":
		s.serveMe(ctx, resp, req)
		return
	case "/rpc.UserService/EditProfile":
		s.serveEditProfile(ctx, resp, req)
		return
	case "/rpc.UserService/Users":
		s.serveUsers(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *userServiceServer) serveFriends(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveFriendsJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveFriendsProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *userServiceServer) serveFriendsJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Friends")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(Empty)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *UserFriendsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.Friends(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserFriendsResponse and nil error while calling Friends. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
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
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) serveFriendsProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Friends")
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
	reqContent := new(Empty)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *UserFriendsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.Friends(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserFriendsResponse and nil error while calling Friends. nil responses are not supported"))
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
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) serveFriendsOfFriends(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveFriendsOfFriendsJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveFriendsOfFriendsProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *userServiceServer) serveFriendsOfFriendsJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "FriendsOfFriends")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(Empty)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *UserFriendsOfFriendsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.FriendsOfFriends(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserFriendsOfFriendsResponse and nil error while calling FriendsOfFriends. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
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
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) serveFriendsOfFriendsProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "FriendsOfFriends")
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
	reqContent := new(Empty)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *UserFriendsOfFriendsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.FriendsOfFriends(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserFriendsOfFriendsResponse and nil error while calling FriendsOfFriends. nil responses are not supported"))
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
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) serveMe(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveMeJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveMeProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *userServiceServer) serveMeJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Me")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(Empty)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *UserMeResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.Me(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserMeResponse and nil error while calling Me. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
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
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) serveMeProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Me")
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
	reqContent := new(Empty)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *UserMeResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.Me(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserMeResponse and nil error while calling Me. nil responses are not supported"))
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
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) serveEditProfile(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveEditProfileJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveEditProfileProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *userServiceServer) serveEditProfileJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "EditProfile")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(UserEditProfileRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Empty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.EditProfile(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Empty and nil error while calling EditProfile. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
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
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) serveEditProfileProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "EditProfile")
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
	reqContent := new(UserEditProfileRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Empty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.EditProfile(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Empty and nil error while calling EditProfile. nil responses are not supported"))
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
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) serveUsers(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveUsersJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveUsersProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *userServiceServer) serveUsersJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Users")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(UserUsersRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *UserUsersResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.Users(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserUsersResponse and nil error while calling Users. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
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
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) serveUsersProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Users")
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
	reqContent := new(UserUsersRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *UserUsersResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.Users(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserUsersResponse and nil error while calling Users. nil responses are not supported"))
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
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *userServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor7, 0
}

func (s *userServiceServer) ProtocGenTwirpVersion() string {
	return "v5.12.0"
}

func (s *userServiceServer) PathPrefix() string {
	return UserServicePathPrefix
}

var twirpFileDescriptor7 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0x95, 0x76, 0x6d, 0xd3, 0x27, 0xdb, 0x28, 0x1e, 0x54, 0x69, 0xc6, 0x4b, 0xc9, 0x34,
	0x51, 0x0e, 0x64, 0xa8, 0xa0, 0x89, 0xeb, 0x40, 0x03, 0x0d, 0x69, 0x02, 0x65, 0xe2, 0xc2, 0x25,
	0x32, 0xb1, 0x0b, 0x96, 0xda, 0x24, 0xd8, 0x69, 0x11, 0x1f, 0x82, 0x3b, 0x9f, 0x8e, 0xcf, 0x82,
	0x6c, 0x27, 0xae, 0x97, 0x15, 0x7a, 0xe1, 0x50, 0x45, 0xcf, 0x3f, 0xff, 0xe7, 0xe7, 0xe7, 0xc5,
	0x0d, 0xc0, 0x52, 0x50, 0x1e, 0x15, 0x3c, 0x2f, 0x73, 0xd4, 0xe6, 0x45, 0x1a, 0x78, 0x8b, 0x9c,
	0xd0, 0xb9, 0x56, 0xc2, 0x04, 0x46, 0x1f, 0x05, 0xe5, 0x6f, 0x38, 0xa3, 0x19, 0x11, 0xfa, 0xf1,
	0x7e, 0xa6, 0x9f, 0xe8, 0x3e, 0xec, 0xc8, 0x64, 0xdf, 0x19, 0x3b, 0x13, 0x6f, 0xda, 0x8f, 0x78,
	0x91, 0x46, 0xd2, 0x1d, 0x2b, 0x19, 0x1d, 0xc1, 0x1e, 0xcb, 0x56, 0xac, 0xa4, 0x89, 0x28, 0x71,
	0xb9, 0x14, 0x7e, 0x6b, 0xec, 0x4c, 0xfa, 0xf1, 0xae, 0x16, 0xaf, 0x94, 0x16, 0xce, 0xe1, 0xf6,
	0x8d, 0x03, 0xb6, 0x81, 0x5f, 0x42, 0x6f, 0xa6, 0xfd, 0x7e, 0x6b, 0xdc, 0x9e, 0x78, 0xd3, 0x07,
	0xc6, 0xb1, 0xb1, 0xd0, 0xb8, 0xb6, 0x87, 0x6f, 0xe1, 0xc0, 0x72, 0xc5, 0x54, 0x14, 0x79, 0x26,
	0x28, 0x7a, 0xb6, 0x06, 0x3a, 0x0a, 0x38, 0xdc, 0x0c, 0x5c, 0x83, 0x7e, 0x3a, 0x10, 0x5a, 0xaf,
	0xeb, 0x93, 0x0c, 0xf2, 0xff, 0x4d, 0x08, 0x1d, 0xad, 0x8b, 0x6b, 0xab, 0xe2, 0x2c, 0x8c, 0xa9,
	0x07, 0xc3, 0xbd, 0x7f, 0x95, 0x83, 0xce, 0x9a, 0x1d, 0x3e, 0x6e, 0x76, 0xf8, 0x97, 0x16, 0xd6,
	0x47, 0xbc, 0x83, 0x7d, 0x69, 0xbf, 0xa4, 0x06, 0xba, 0xa5, 0xbb, 0x11, 0xb8, 0x4c, 0x24, 0x98,
	0x2c, 0x58, 0xa6, 0x1a, 0x73, 0xe3, 0x1e, 0x13, 0x67, 0x32, 0x0c, 0x7f, 0x3b, 0x30, 0x94, 0xce,
	0x73, 0xc2, 0xca, 0x0f, 0x3c, 0x9f, 0xb1, 0x39, 0x8d, 0xe9, 0xb7, 0x25, 0x15, 0xa5, 0x9c, 0x09,
	0x5d, 0x60, 0x36, 0x4f, 0xd2, 0xaf, 0x38, 0xfb, 0x42, 0x89, 0xa2, 0xbb, 0xf1, 0xae, 0x12, 0x5f,
	0x6b, 0x0d, 0xdd, 0x81, 0x8e, 0x8a, 0xab, 0x81, 0xe9, 0x00, 0x3d, 0x81, 0x41, 0x81, 0x85, 0xf8,
	0x9e, 0x73, 0x62, 0xb2, 0xdb, 0x2a, 0xfb, 0x56, 0xad, 0xd7, 0x80, 0x00, 0xdc, 0x5a, 0xf2, 0x77,
	0x14, 0xc3, 0xc4, 0xe8, 0x18, 0xf6, 0xf1, 0x0a, 0x97, 0x98, 0x1b, 0x48, 0x47, 0x41, 0xf6, 0xb4,
	0x5a, 0x23, 0x0e, 0xa1, 0x5f, 0xd9, 0x18, 0xf1, 0xbb, 0x9a, 0xa1, 0x85, 0x0b, 0x12, 0x3e, 0x85,
	0x81, 0xec, 0x4f, 0xfe, 0x44, 0xdd, 0xd9, 0x08, 0x5c, 0x39, 0x97, 0x84, 0x55, 0x4b, 0xe8, 0xc7,
	0x3d, 0x19, 0x5f, 0x10, 0x11, 0xbe, 0xd0, 0xff, 0x82, 0xca, 0x5e, 0x8d, 0xf7, 0x21, 0x74, 0xe4,
	0xfb, 0x7a, 0x63, 0xd6, 0x7c, 0xb5, 0x3e, 0xfd, 0xd5, 0x02, 0x4f, 0xc6, 0x57, 0x94, 0xaf, 0x58,
	0x4a, 0xd1, 0x09, 0xf4, 0xaa, 0x1d, 0x22, 0x50, 0xe6, 0xf3, 0x45, 0x51, 0xfe, 0x08, 0xfc, 0xe6,
	0xaa, 0xad, 0x5b, 0x31, 0x68, 0x6e, 0xff, 0x5a, 0xe6, 0xa3, 0xad, 0x97, 0x04, 0x1d, 0x43, 0xeb,
	0x92, 0x5e, 0x4b, 0x3a, 0x30, 0x49, 0xd6, 0x55, 0x39, 0x05, 0xcf, 0xda, 0x35, 0x3a, 0x34, 0x9e,
	0x9b, 0x37, 0x20, 0xb0, 0x60, 0xe8, 0x14, 0x3a, 0x6a, 0x28, 0xe8, 0xae, 0xc9, 0xb0, 0x67, 0x1a,
	0x0c, 0x9b, 0xb2, 0x3e, 0xef, 0x95, 0xfb, 0xa9, 0x1b, 0x45, 0x27, 0xbc, 0x48, 0x3f, 0x77, 0xd5,
	0x87, 0xec, 0xf9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0xbc, 0x2f, 0x89, 0xe8, 0x04, 0x00,
	0x00,
}