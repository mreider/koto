// Code generated by protoc-gen-twirp v5.12.0, DO NOT EDIT.
// source: blob.proto

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
// BlobService Interface
// =====================

type BlobService interface {
	UploadLink(context.Context, *BlobUploadLinkRequest) (*BlobUploadLinkResponse, error)
}

// ===========================
// BlobService Protobuf Client
// ===========================

type blobServiceProtobufClient struct {
	client HTTPClient
	urls   [1]string
	opts   twirp.ClientOptions
}

// NewBlobServiceProtobufClient creates a Protobuf client that implements the BlobService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewBlobServiceProtobufClient(addr string, client HTTPClient, opts ...twirp.ClientOption) BlobService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + BlobServicePathPrefix
	urls := [1]string{
		prefix + "UploadLink",
	}

	return &blobServiceProtobufClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *blobServiceProtobufClient) UploadLink(ctx context.Context, in *BlobUploadLinkRequest) (*BlobUploadLinkResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "BlobService")
	ctx = ctxsetters.WithMethodName(ctx, "UploadLink")
	out := new(BlobUploadLinkResponse)
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

// =======================
// BlobService JSON Client
// =======================

type blobServiceJSONClient struct {
	client HTTPClient
	urls   [1]string
	opts   twirp.ClientOptions
}

// NewBlobServiceJSONClient creates a JSON client that implements the BlobService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewBlobServiceJSONClient(addr string, client HTTPClient, opts ...twirp.ClientOption) BlobService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + BlobServicePathPrefix
	urls := [1]string{
		prefix + "UploadLink",
	}

	return &blobServiceJSONClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *blobServiceJSONClient) UploadLink(ctx context.Context, in *BlobUploadLinkRequest) (*BlobUploadLinkResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "BlobService")
	ctx = ctxsetters.WithMethodName(ctx, "UploadLink")
	out := new(BlobUploadLinkResponse)
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

// ==========================
// BlobService Server Handler
// ==========================

type blobServiceServer struct {
	BlobService
	hooks *twirp.ServerHooks
}

func NewBlobServiceServer(svc BlobService, hooks *twirp.ServerHooks) TwirpServer {
	return &blobServiceServer{
		BlobService: svc,
		hooks:       hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *blobServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// BlobServicePathPrefix is used for all URL paths on a twirp BlobService server.
// Requests are always: POST BlobServicePathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const BlobServicePathPrefix = "/rpc.BlobService/"

func (s *blobServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "BlobService")
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
	case "/rpc.BlobService/UploadLink":
		s.serveUploadLink(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *blobServiceServer) serveUploadLink(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveUploadLinkJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveUploadLinkProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *blobServiceServer) serveUploadLinkJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "UploadLink")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(BlobUploadLinkRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *BlobUploadLinkResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.BlobService.UploadLink(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *BlobUploadLinkResponse and nil error while calling UploadLink. nil responses are not supported"))
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

func (s *blobServiceServer) serveUploadLinkProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "UploadLink")
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
	reqContent := new(BlobUploadLinkRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *BlobUploadLinkResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.BlobService.UploadLink(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *BlobUploadLinkResponse and nil error while calling UploadLink. nil responses are not supported"))
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

func (s *blobServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor1, 0
}

func (s *blobServiceServer) ProtocGenTwirpVersion() string {
	return "v5.12.0"
}

func (s *blobServiceServer) PathPrefix() string {
	return BlobServicePathPrefix
}

var twirpFileDescriptor1 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x49, 0xf3, 0x7d, 0xb1, 0xbd, 0x51, 0x90, 0x8b, 0x7f, 0x42, 0xdc, 0xd4, 0xae, 0xea,
	0x66, 0x84, 0xba, 0x91, 0xba, 0x2b, 0xb6, 0x20, 0xb8, 0x8a, 0x75, 0xe3, 0x26, 0x4c, 0x26, 0xb7,
	0x10, 0x32, 0x9d, 0x19, 0x27, 0xd3, 0x42, 0x5e, 0xd1, 0xa7, 0x92, 0xc4, 0x01, 0x51, 0x8a, 0xbb,
	0x7b, 0xce, 0x70, 0xe6, 0x77, 0x38, 0x00, 0x85, 0xd4, 0x05, 0x33, 0x56, 0x3b, 0x8d, 0xa1, 0x35,
	0x62, 0x32, 0x87, 0xf3, 0x85, 0xd4, 0xc5, 0xab, 0x91, 0x9a, 0x97, 0xcf, 0x95, 0xaa, 0x33, 0x7a,
	0xdf, 0x51, 0xe3, 0xf0, 0x1a, 0x8e, 0x85, 0x56, 0x8e, 0x94, 0xcb, 0x5d, 0x6b, 0x28, 0x09, 0xc6,
	0xc1, 0x74, 0x94, 0xc5, 0xde, 0x5b, 0xb7, 0x86, 0x26, 0x1f, 0x01, 0x5c, 0xfc, 0x0e, 0x37, 0x46,
	0xab, 0x86, 0xf0, 0x12, 0x8e, 0x3a, 0x52, 0x5e, 0x95, 0x3e, 0x18, 0x75, 0xf2, 0xa9, 0x44, 0x84,
	0x7f, 0xb2, 0x52, 0x75, 0x32, 0xe8, 0xdd, 0xfe, 0xc6, 0x15, 0x8c, 0x36, 0xda, 0x6e, 0xf3, 0x92,
	0x3b, 0x9e, 0x84, 0xe3, 0x70, 0x1a, 0xcf, 0x6e, 0x98, 0x35, 0x82, 0x1d, 0xfe, 0x9c, 0xad, 0xb4,
	0xdd, 0x3e, 0x72, 0xc7, 0x97, 0xca, 0xd9, 0x36, 0x1b, 0x6e, 0xbc, 0x4c, 0x1f, 0xe0, 0xe4, 0xc7,
	0x13, 0x9e, 0x42, 0x58, 0x53, 0xeb, 0x1b, 0x74, 0x27, 0x9e, 0xc1, 0xff, 0x3d, 0x97, 0x3b, 0xf2,
	0xfc, 0x2f, 0x31, 0x1f, 0xdc, 0x07, 0xb3, 0x35, 0xc4, 0x1d, 0xee, 0x85, 0xec, 0xbe, 0x12, 0x84,
	0x4b, 0x80, 0x6f, 0x32, 0xa6, 0x07, 0xeb, 0xf4, 0x43, 0xa5, 0x57, 0x7f, 0x54, 0x5d, 0x0c, 0xdf,
	0x22, 0xc6, 0x6e, 0xad, 0x11, 0x45, 0xd4, 0x8f, 0x7e, 0xf7, 0x19, 0x00, 0x00, 0xff, 0xff, 0xe3,
	0x31, 0xad, 0xda, 0x82, 0x01, 0x00, 0x00,
}