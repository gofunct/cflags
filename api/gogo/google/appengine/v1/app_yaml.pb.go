// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/appengine/v1/app_yaml.proto

package google_appengine_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import google_protobuf1 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Actions to take when the user is not logged in.
type AuthFailAction int32

const (
	// Not specified. `AUTH_FAIL_ACTION_REDIRECT` is assumed.
	AuthFailAction_AUTH_FAIL_ACTION_UNSPECIFIED AuthFailAction = 0
	// Redirects user to "accounts.google.com". The user is redirected back to the
	// application URL after signing in or creating an account.
	AuthFailAction_AUTH_FAIL_ACTION_REDIRECT AuthFailAction = 1
	// Rejects request with a `401` HTTP status code and an error
	// message.
	AuthFailAction_AUTH_FAIL_ACTION_UNAUTHORIZED AuthFailAction = 2
)

var AuthFailAction_name = map[int32]string{
	0: "AUTH_FAIL_ACTION_UNSPECIFIED",
	1: "AUTH_FAIL_ACTION_REDIRECT",
	2: "AUTH_FAIL_ACTION_UNAUTHORIZED",
}
var AuthFailAction_value = map[string]int32{
	"AUTH_FAIL_ACTION_UNSPECIFIED":  0,
	"AUTH_FAIL_ACTION_REDIRECT":     1,
	"AUTH_FAIL_ACTION_UNAUTHORIZED": 2,
}

func (x AuthFailAction) String() string {
	return proto.EnumName(AuthFailAction_name, int32(x))
}
func (AuthFailAction) EnumDescriptor() ([]byte, []int) { return fileDescriptorAppYaml, []int{0} }

// Methods to restrict access to a URL based on login status.
type LoginRequirement int32

const (
	// Not specified. `LOGIN_OPTIONAL` is assumed.
	LoginRequirement_LOGIN_UNSPECIFIED LoginRequirement = 0
	// Does not require that the user is signed in.
	LoginRequirement_LOGIN_OPTIONAL LoginRequirement = 1
	// If the user is not signed in, the `auth_fail_action` is taken.
	// In addition, if the user is not an administrator for the
	// application, they are given an error message regardless of
	// `auth_fail_action`. If the user is an administrator, the handler
	// proceeds.
	LoginRequirement_LOGIN_ADMIN LoginRequirement = 2
	// If the user has signed in, the handler proceeds normally. Otherwise, the
	// auth_fail_action is taken.
	LoginRequirement_LOGIN_REQUIRED LoginRequirement = 3
)

var LoginRequirement_name = map[int32]string{
	0: "LOGIN_UNSPECIFIED",
	1: "LOGIN_OPTIONAL",
	2: "LOGIN_ADMIN",
	3: "LOGIN_REQUIRED",
}
var LoginRequirement_value = map[string]int32{
	"LOGIN_UNSPECIFIED": 0,
	"LOGIN_OPTIONAL":    1,
	"LOGIN_ADMIN":       2,
	"LOGIN_REQUIRED":    3,
}

func (x LoginRequirement) String() string {
	return proto.EnumName(LoginRequirement_name, int32(x))
}
func (LoginRequirement) EnumDescriptor() ([]byte, []int) { return fileDescriptorAppYaml, []int{1} }

// Methods to enforce security (HTTPS) on a URL.
type SecurityLevel int32

const (
	// Not specified.
	SecurityLevel_SECURE_UNSPECIFIED SecurityLevel = 0
	// Both HTTP and HTTPS requests with URLs that match the handler succeed
	// without redirects. The application can examine the request to determine
	// which protocol was used, and respond accordingly.
	SecurityLevel_SECURE_DEFAULT SecurityLevel = 0
	// Requests for a URL that match this handler that use HTTPS are automatically
	// redirected to the HTTP equivalent URL.
	SecurityLevel_SECURE_NEVER SecurityLevel = 1
	// Both HTTP and HTTPS requests with URLs that match the handler succeed
	// without redirects. The application can examine the request to determine
	// which protocol was used and respond accordingly.
	SecurityLevel_SECURE_OPTIONAL SecurityLevel = 2
	// Requests for a URL that match this handler that do not use HTTPS are
	// automatically redirected to the HTTPS URL with the same path. Query
	// parameters are reserved for the redirect.
	SecurityLevel_SECURE_ALWAYS SecurityLevel = 3
)

var SecurityLevel_name = map[int32]string{
	0: "SECURE_UNSPECIFIED",
	// Duplicate value: 0: "SECURE_DEFAULT",
	1: "SECURE_NEVER",
	2: "SECURE_OPTIONAL",
	3: "SECURE_ALWAYS",
}
var SecurityLevel_value = map[string]int32{
	"SECURE_UNSPECIFIED": 0,
	"SECURE_DEFAULT":     0,
	"SECURE_NEVER":       1,
	"SECURE_OPTIONAL":    2,
	"SECURE_ALWAYS":      3,
}

func (x SecurityLevel) String() string {
	return proto.EnumName(SecurityLevel_name, int32(x))
}
func (SecurityLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptorAppYaml, []int{2} }

// Error codes.
type ErrorHandler_ErrorCode int32

const (
	// Not specified. ERROR_CODE_DEFAULT is assumed.
	ErrorHandler_ERROR_CODE_UNSPECIFIED ErrorHandler_ErrorCode = 0
	// All other error types.
	ErrorHandler_ERROR_CODE_DEFAULT ErrorHandler_ErrorCode = 0
	// Application has exceeded a resource quota.
	ErrorHandler_ERROR_CODE_OVER_QUOTA ErrorHandler_ErrorCode = 1
	// Client blocked by the application's Denial of Service protection
	// configuration.
	ErrorHandler_ERROR_CODE_DOS_API_DENIAL ErrorHandler_ErrorCode = 2
	// Deadline reached before the application responds.
	ErrorHandler_ERROR_CODE_TIMEOUT ErrorHandler_ErrorCode = 3
)

var ErrorHandler_ErrorCode_name = map[int32]string{
	0: "ERROR_CODE_UNSPECIFIED",
	// Duplicate value: 0: "ERROR_CODE_DEFAULT",
	1: "ERROR_CODE_OVER_QUOTA",
	2: "ERROR_CODE_DOS_API_DENIAL",
	3: "ERROR_CODE_TIMEOUT",
}
var ErrorHandler_ErrorCode_value = map[string]int32{
	"ERROR_CODE_UNSPECIFIED":    0,
	"ERROR_CODE_DEFAULT":        0,
	"ERROR_CODE_OVER_QUOTA":     1,
	"ERROR_CODE_DOS_API_DENIAL": 2,
	"ERROR_CODE_TIMEOUT":        3,
}

func (x ErrorHandler_ErrorCode) String() string {
	return proto.EnumName(ErrorHandler_ErrorCode_name, int32(x))
}
func (ErrorHandler_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorAppYaml, []int{1, 0}
}

// Redirect codes.
type UrlMap_RedirectHttpResponseCode int32

const (
	// Not specified. `302` is assumed.
	UrlMap_REDIRECT_HTTP_RESPONSE_CODE_UNSPECIFIED UrlMap_RedirectHttpResponseCode = 0
	// `301 Moved Permanently` code.
	UrlMap_REDIRECT_HTTP_RESPONSE_CODE_301 UrlMap_RedirectHttpResponseCode = 1
	// `302 Moved Temporarily` code.
	UrlMap_REDIRECT_HTTP_RESPONSE_CODE_302 UrlMap_RedirectHttpResponseCode = 2
	// `303 See Other` code.
	UrlMap_REDIRECT_HTTP_RESPONSE_CODE_303 UrlMap_RedirectHttpResponseCode = 3
	// `307 Temporary Redirect` code.
	UrlMap_REDIRECT_HTTP_RESPONSE_CODE_307 UrlMap_RedirectHttpResponseCode = 4
)

var UrlMap_RedirectHttpResponseCode_name = map[int32]string{
	0: "REDIRECT_HTTP_RESPONSE_CODE_UNSPECIFIED",
	1: "REDIRECT_HTTP_RESPONSE_CODE_301",
	2: "REDIRECT_HTTP_RESPONSE_CODE_302",
	3: "REDIRECT_HTTP_RESPONSE_CODE_303",
	4: "REDIRECT_HTTP_RESPONSE_CODE_307",
}
var UrlMap_RedirectHttpResponseCode_value = map[string]int32{
	"REDIRECT_HTTP_RESPONSE_CODE_UNSPECIFIED": 0,
	"REDIRECT_HTTP_RESPONSE_CODE_301":         1,
	"REDIRECT_HTTP_RESPONSE_CODE_302":         2,
	"REDIRECT_HTTP_RESPONSE_CODE_303":         3,
	"REDIRECT_HTTP_RESPONSE_CODE_307":         4,
}

func (x UrlMap_RedirectHttpResponseCode) String() string {
	return proto.EnumName(UrlMap_RedirectHttpResponseCode_name, int32(x))
}
func (UrlMap_RedirectHttpResponseCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorAppYaml, []int{2, 0}
}

// [Google Cloud Endpoints](https://cloud.google.com/appengine/docs/python/endpoints/)
// configuration for API handlers.
type ApiConfigHandler struct {
	// Action to take when users access resources that require
	// authentication. Defaults to `redirect`.
	AuthFailAction AuthFailAction `protobuf:"varint,1,opt,name=auth_fail_action,json=authFailAction,proto3,enum=google.appengine.v1.AuthFailAction" json:"auth_fail_action,omitempty"`
	// Level of login required to access this resource. Defaults to
	// `optional`.
	Login LoginRequirement `protobuf:"varint,2,opt,name=login,proto3,enum=google.appengine.v1.LoginRequirement" json:"login,omitempty"`
	// Path to the script from the application root directory.
	Script string `protobuf:"bytes,3,opt,name=script,proto3" json:"script,omitempty"`
	// Security (HTTPS) enforcement for this URL.
	SecurityLevel SecurityLevel `protobuf:"varint,4,opt,name=security_level,json=securityLevel,proto3,enum=google.appengine.v1.SecurityLevel" json:"security_level,omitempty"`
	// URL to serve the endpoint at.
	Url string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *ApiConfigHandler) Reset()                    { *m = ApiConfigHandler{} }
func (m *ApiConfigHandler) String() string            { return proto.CompactTextString(m) }
func (*ApiConfigHandler) ProtoMessage()               {}
func (*ApiConfigHandler) Descriptor() ([]byte, []int) { return fileDescriptorAppYaml, []int{0} }

func (m *ApiConfigHandler) GetAuthFailAction() AuthFailAction {
	if m != nil {
		return m.AuthFailAction
	}
	return AuthFailAction_AUTH_FAIL_ACTION_UNSPECIFIED
}

func (m *ApiConfigHandler) GetLogin() LoginRequirement {
	if m != nil {
		return m.Login
	}
	return LoginRequirement_LOGIN_UNSPECIFIED
}

func (m *ApiConfigHandler) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func (m *ApiConfigHandler) GetSecurityLevel() SecurityLevel {
	if m != nil {
		return m.SecurityLevel
	}
	return SecurityLevel_SECURE_UNSPECIFIED
}

func (m *ApiConfigHandler) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// Custom static error page to be served when an error occurs.
type ErrorHandler struct {
	// Error condition this handler applies to.
	ErrorCode ErrorHandler_ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=google.appengine.v1.ErrorHandler_ErrorCode" json:"error_code,omitempty"`
	// Static file content to be served for this error.
	StaticFile string `protobuf:"bytes,2,opt,name=static_file,json=staticFile,proto3" json:"static_file,omitempty"`
	// MIME type of file. Defaults to `text/html`.
	MimeType string `protobuf:"bytes,3,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (m *ErrorHandler) Reset()                    { *m = ErrorHandler{} }
func (m *ErrorHandler) String() string            { return proto.CompactTextString(m) }
func (*ErrorHandler) ProtoMessage()               {}
func (*ErrorHandler) Descriptor() ([]byte, []int) { return fileDescriptorAppYaml, []int{1} }

func (m *ErrorHandler) GetErrorCode() ErrorHandler_ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return ErrorHandler_ERROR_CODE_UNSPECIFIED
}

func (m *ErrorHandler) GetStaticFile() string {
	if m != nil {
		return m.StaticFile
	}
	return ""
}

func (m *ErrorHandler) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

// URL pattern and description of how the URL should be handled. App Engine can
// handle URLs by executing application code or by serving static files
// uploaded with the version, such as images, CSS, or JavaScript.
type UrlMap struct {
	// URL prefix. Uses regular expression syntax, which means regexp
	// special characters must be escaped, but should not contain groupings.
	// All URLs that begin with this prefix are handled by this handler, using the
	// portion of the URL after the prefix as part of the file path.
	UrlRegex string `protobuf:"bytes,1,opt,name=url_regex,json=urlRegex,proto3" json:"url_regex,omitempty"`
	// Type of handler for this URL pattern.
	//
	// Types that are valid to be assigned to HandlerType:
	//	*UrlMap_StaticFiles
	//	*UrlMap_Script
	//	*UrlMap_ApiEndpoint
	HandlerType isUrlMap_HandlerType `protobuf_oneof:"handler_type"`
	// Security (HTTPS) enforcement for this URL.
	SecurityLevel SecurityLevel `protobuf:"varint,5,opt,name=security_level,json=securityLevel,proto3,enum=google.appengine.v1.SecurityLevel" json:"security_level,omitempty"`
	// Level of login required to access this resource.
	Login LoginRequirement `protobuf:"varint,6,opt,name=login,proto3,enum=google.appengine.v1.LoginRequirement" json:"login,omitempty"`
	// Action to take when users access resources that require
	// authentication. Defaults to `redirect`.
	AuthFailAction AuthFailAction `protobuf:"varint,7,opt,name=auth_fail_action,json=authFailAction,proto3,enum=google.appengine.v1.AuthFailAction" json:"auth_fail_action,omitempty"`
	// `30x` code to use when performing redirects for the `secure` field.
	// Defaults to `302`.
	RedirectHttpResponseCode UrlMap_RedirectHttpResponseCode `protobuf:"varint,8,opt,name=redirect_http_response_code,json=redirectHttpResponseCode,proto3,enum=google.appengine.v1.UrlMap_RedirectHttpResponseCode" json:"redirect_http_response_code,omitempty"`
}

func (m *UrlMap) Reset()                    { *m = UrlMap{} }
func (m *UrlMap) String() string            { return proto.CompactTextString(m) }
func (*UrlMap) ProtoMessage()               {}
func (*UrlMap) Descriptor() ([]byte, []int) { return fileDescriptorAppYaml, []int{2} }

type isUrlMap_HandlerType interface {
	isUrlMap_HandlerType()
}

type UrlMap_StaticFiles struct {
	StaticFiles *StaticFilesHandler `protobuf:"bytes,2,opt,name=static_files,json=staticFiles,oneof"`
}
type UrlMap_Script struct {
	Script *ScriptHandler `protobuf:"bytes,3,opt,name=script,oneof"`
}
type UrlMap_ApiEndpoint struct {
	ApiEndpoint *ApiEndpointHandler `protobuf:"bytes,4,opt,name=api_endpoint,json=apiEndpoint,oneof"`
}

func (*UrlMap_StaticFiles) isUrlMap_HandlerType() {}
func (*UrlMap_Script) isUrlMap_HandlerType()      {}
func (*UrlMap_ApiEndpoint) isUrlMap_HandlerType() {}

func (m *UrlMap) GetHandlerType() isUrlMap_HandlerType {
	if m != nil {
		return m.HandlerType
	}
	return nil
}

func (m *UrlMap) GetUrlRegex() string {
	if m != nil {
		return m.UrlRegex
	}
	return ""
}

func (m *UrlMap) GetStaticFiles() *StaticFilesHandler {
	if x, ok := m.GetHandlerType().(*UrlMap_StaticFiles); ok {
		return x.StaticFiles
	}
	return nil
}

func (m *UrlMap) GetScript() *ScriptHandler {
	if x, ok := m.GetHandlerType().(*UrlMap_Script); ok {
		return x.Script
	}
	return nil
}

func (m *UrlMap) GetApiEndpoint() *ApiEndpointHandler {
	if x, ok := m.GetHandlerType().(*UrlMap_ApiEndpoint); ok {
		return x.ApiEndpoint
	}
	return nil
}

func (m *UrlMap) GetSecurityLevel() SecurityLevel {
	if m != nil {
		return m.SecurityLevel
	}
	return SecurityLevel_SECURE_UNSPECIFIED
}

func (m *UrlMap) GetLogin() LoginRequirement {
	if m != nil {
		return m.Login
	}
	return LoginRequirement_LOGIN_UNSPECIFIED
}

func (m *UrlMap) GetAuthFailAction() AuthFailAction {
	if m != nil {
		return m.AuthFailAction
	}
	return AuthFailAction_AUTH_FAIL_ACTION_UNSPECIFIED
}

func (m *UrlMap) GetRedirectHttpResponseCode() UrlMap_RedirectHttpResponseCode {
	if m != nil {
		return m.RedirectHttpResponseCode
	}
	return UrlMap_REDIRECT_HTTP_RESPONSE_CODE_UNSPECIFIED
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*UrlMap) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _UrlMap_OneofMarshaler, _UrlMap_OneofUnmarshaler, _UrlMap_OneofSizer, []interface{}{
		(*UrlMap_StaticFiles)(nil),
		(*UrlMap_Script)(nil),
		(*UrlMap_ApiEndpoint)(nil),
	}
}

func _UrlMap_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*UrlMap)
	// handler_type
	switch x := m.HandlerType.(type) {
	case *UrlMap_StaticFiles:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StaticFiles); err != nil {
			return err
		}
	case *UrlMap_Script:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Script); err != nil {
			return err
		}
	case *UrlMap_ApiEndpoint:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ApiEndpoint); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("UrlMap.HandlerType has unexpected type %T", x)
	}
	return nil
}

func _UrlMap_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*UrlMap)
	switch tag {
	case 2: // handler_type.static_files
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StaticFilesHandler)
		err := b.DecodeMessage(msg)
		m.HandlerType = &UrlMap_StaticFiles{msg}
		return true, err
	case 3: // handler_type.script
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ScriptHandler)
		err := b.DecodeMessage(msg)
		m.HandlerType = &UrlMap_Script{msg}
		return true, err
	case 4: // handler_type.api_endpoint
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ApiEndpointHandler)
		err := b.DecodeMessage(msg)
		m.HandlerType = &UrlMap_ApiEndpoint{msg}
		return true, err
	default:
		return false, nil
	}
}

func _UrlMap_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*UrlMap)
	// handler_type
	switch x := m.HandlerType.(type) {
	case *UrlMap_StaticFiles:
		s := proto.Size(x.StaticFiles)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UrlMap_Script:
		s := proto.Size(x.Script)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *UrlMap_ApiEndpoint:
		s := proto.Size(x.ApiEndpoint)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Files served directly to the user for a given URL, such as images, CSS
// stylesheets, or JavaScript source files. Static file handlers describe which
// files in the application directory are static files, and which URLs serve
// them.
type StaticFilesHandler struct {
	// Path to the static files matched by the URL pattern, from the
	// application root directory. The path can refer to text matched in groupings
	// in the URL pattern.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Regular expression that matches the file paths for all files that should be
	// referenced by this handler.
	UploadPathRegex string `protobuf:"bytes,2,opt,name=upload_path_regex,json=uploadPathRegex,proto3" json:"upload_path_regex,omitempty"`
	// HTTP headers to use for all responses from these URLs.
	HttpHeaders map[string]string `protobuf:"bytes,3,rep,name=http_headers,json=httpHeaders" json:"http_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// MIME type used to serve all files served by this handler.
	//
	// Defaults to file-specific MIME types, which are derived from each file's
	// filename extension.
	MimeType string `protobuf:"bytes,4,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Time a static file served by this handler should be cached
	// by web proxies and browsers.
	Expiration *google_protobuf1.Duration `protobuf:"bytes,5,opt,name=expiration" json:"expiration,omitempty"`
	// Whether this handler should match the request if the file
	// referenced by the handler does not exist.
	RequireMatchingFile bool `protobuf:"varint,6,opt,name=require_matching_file,json=requireMatchingFile,proto3" json:"require_matching_file,omitempty"`
	// Whether files should also be uploaded as code data. By default, files
	// declared in static file handlers are uploaded as static
	// data and are only served to end users; they cannot be read by the
	// application. If enabled, uploads are charged against both your code and
	// static data storage resource quotas.
	ApplicationReadable bool `protobuf:"varint,7,opt,name=application_readable,json=applicationReadable,proto3" json:"application_readable,omitempty"`
}

func (m *StaticFilesHandler) Reset()                    { *m = StaticFilesHandler{} }
func (m *StaticFilesHandler) String() string            { return proto.CompactTextString(m) }
func (*StaticFilesHandler) ProtoMessage()               {}
func (*StaticFilesHandler) Descriptor() ([]byte, []int) { return fileDescriptorAppYaml, []int{3} }

func (m *StaticFilesHandler) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *StaticFilesHandler) GetUploadPathRegex() string {
	if m != nil {
		return m.UploadPathRegex
	}
	return ""
}

func (m *StaticFilesHandler) GetHttpHeaders() map[string]string {
	if m != nil {
		return m.HttpHeaders
	}
	return nil
}

func (m *StaticFilesHandler) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *StaticFilesHandler) GetExpiration() *google_protobuf1.Duration {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *StaticFilesHandler) GetRequireMatchingFile() bool {
	if m != nil {
		return m.RequireMatchingFile
	}
	return false
}

func (m *StaticFilesHandler) GetApplicationReadable() bool {
	if m != nil {
		return m.ApplicationReadable
	}
	return false
}

// Executes a script to handle the request that matches the URL pattern.
type ScriptHandler struct {
	// Path to the script from the application root directory.
	ScriptPath string `protobuf:"bytes,1,opt,name=script_path,json=scriptPath,proto3" json:"script_path,omitempty"`
}

func (m *ScriptHandler) Reset()                    { *m = ScriptHandler{} }
func (m *ScriptHandler) String() string            { return proto.CompactTextString(m) }
func (*ScriptHandler) ProtoMessage()               {}
func (*ScriptHandler) Descriptor() ([]byte, []int) { return fileDescriptorAppYaml, []int{4} }

func (m *ScriptHandler) GetScriptPath() string {
	if m != nil {
		return m.ScriptPath
	}
	return ""
}

// Uses Google Cloud Endpoints to handle requests.
type ApiEndpointHandler struct {
	// Path to the script from the application root directory.
	ScriptPath string `protobuf:"bytes,1,opt,name=script_path,json=scriptPath,proto3" json:"script_path,omitempty"`
}

func (m *ApiEndpointHandler) Reset()                    { *m = ApiEndpointHandler{} }
func (m *ApiEndpointHandler) String() string            { return proto.CompactTextString(m) }
func (*ApiEndpointHandler) ProtoMessage()               {}
func (*ApiEndpointHandler) Descriptor() ([]byte, []int) { return fileDescriptorAppYaml, []int{5} }

func (m *ApiEndpointHandler) GetScriptPath() string {
	if m != nil {
		return m.ScriptPath
	}
	return ""
}

// Health checking configuration for VM instances. Unhealthy instances
// are killed and replaced with new instances. Only applicable for
// instances in App Engine flexible environment.
type HealthCheck struct {
	// Whether to explicitly disable health checks for this instance.
	DisableHealthCheck bool `protobuf:"varint,1,opt,name=disable_health_check,json=disableHealthCheck,proto3" json:"disable_health_check,omitempty"`
	// Host header to send when performing an HTTP health check.
	// Example: "myapp.appspot.com"
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// Number of consecutive successful health checks required before receiving
	// traffic.
	HealthyThreshold uint32 `protobuf:"varint,3,opt,name=healthy_threshold,json=healthyThreshold,proto3" json:"healthy_threshold,omitempty"`
	// Number of consecutive failed health checks required before removing
	// traffic.
	UnhealthyThreshold uint32 `protobuf:"varint,4,opt,name=unhealthy_threshold,json=unhealthyThreshold,proto3" json:"unhealthy_threshold,omitempty"`
	// Number of consecutive failed health checks required before an instance is
	// restarted.
	RestartThreshold uint32 `protobuf:"varint,5,opt,name=restart_threshold,json=restartThreshold,proto3" json:"restart_threshold,omitempty"`
	// Interval between health checks.
	CheckInterval *google_protobuf1.Duration `protobuf:"bytes,6,opt,name=check_interval,json=checkInterval" json:"check_interval,omitempty"`
	// Time before the health check is considered failed.
	Timeout *google_protobuf1.Duration `protobuf:"bytes,7,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *HealthCheck) Reset()                    { *m = HealthCheck{} }
func (m *HealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()               {}
func (*HealthCheck) Descriptor() ([]byte, []int) { return fileDescriptorAppYaml, []int{6} }

func (m *HealthCheck) GetDisableHealthCheck() bool {
	if m != nil {
		return m.DisableHealthCheck
	}
	return false
}

func (m *HealthCheck) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HealthCheck) GetHealthyThreshold() uint32 {
	if m != nil {
		return m.HealthyThreshold
	}
	return 0
}

func (m *HealthCheck) GetUnhealthyThreshold() uint32 {
	if m != nil {
		return m.UnhealthyThreshold
	}
	return 0
}

func (m *HealthCheck) GetRestartThreshold() uint32 {
	if m != nil {
		return m.RestartThreshold
	}
	return 0
}

func (m *HealthCheck) GetCheckInterval() *google_protobuf1.Duration {
	if m != nil {
		return m.CheckInterval
	}
	return nil
}

func (m *HealthCheck) GetTimeout() *google_protobuf1.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

// Third-party Python runtime library that is required by the application.
type Library struct {
	// Name of the library. Example: "django".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Version of the library to select, or "latest".
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *Library) Reset()                    { *m = Library{} }
func (m *Library) String() string            { return proto.CompactTextString(m) }
func (*Library) ProtoMessage()               {}
func (*Library) Descriptor() ([]byte, []int) { return fileDescriptorAppYaml, []int{7} }

func (m *Library) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Library) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*ApiConfigHandler)(nil), "google.appengine.v1.ApiConfigHandler")
	proto.RegisterType((*ErrorHandler)(nil), "google.appengine.v1.ErrorHandler")
	proto.RegisterType((*UrlMap)(nil), "google.appengine.v1.UrlMap")
	proto.RegisterType((*StaticFilesHandler)(nil), "google.appengine.v1.StaticFilesHandler")
	proto.RegisterType((*ScriptHandler)(nil), "google.appengine.v1.ScriptHandler")
	proto.RegisterType((*ApiEndpointHandler)(nil), "google.appengine.v1.ApiEndpointHandler")
	proto.RegisterType((*HealthCheck)(nil), "google.appengine.v1.HealthCheck")
	proto.RegisterType((*Library)(nil), "google.appengine.v1.Library")
	proto.RegisterEnum("google.appengine.v1.AuthFailAction", AuthFailAction_name, AuthFailAction_value)
	proto.RegisterEnum("google.appengine.v1.LoginRequirement", LoginRequirement_name, LoginRequirement_value)
	proto.RegisterEnum("google.appengine.v1.SecurityLevel", SecurityLevel_name, SecurityLevel_value)
	proto.RegisterEnum("google.appengine.v1.ErrorHandler_ErrorCode", ErrorHandler_ErrorCode_name, ErrorHandler_ErrorCode_value)
	proto.RegisterEnum("google.appengine.v1.UrlMap_RedirectHttpResponseCode", UrlMap_RedirectHttpResponseCode_name, UrlMap_RedirectHttpResponseCode_value)
}

func init() { proto.RegisterFile("google/appengine/v1/app_yaml.proto", fileDescriptorAppYaml) }

var fileDescriptorAppYaml = []byte{
	// 1205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xae, 0xec, 0xfc, 0x1e, 0x3b, 0xae, 0xc2, 0xb4, 0x9d, 0x93, 0xb6, 0x6b, 0xe6, 0x62, 0x68,
	0x97, 0x00, 0x4e, 0x93, 0x6c, 0x68, 0xf7, 0x83, 0x61, 0xaa, 0xad, 0xcc, 0x1a, 0x9c, 0xd8, 0xa5,
	0xed, 0x0e, 0xdd, 0x2e, 0x08, 0xc6, 0x66, 0x2c, 0xa1, 0xb2, 0xa4, 0x51, 0x94, 0x51, 0x3f, 0xc7,
	0xb0, 0xf7, 0xe8, 0xed, 0x1e, 0x64, 0x57, 0x7b, 0x99, 0x81, 0x14, 0xfd, 0x97, 0x38, 0xcd, 0x90,
	0x3b, 0x9e, 0x73, 0xbe, 0xef, 0x88, 0xe7, 0x97, 0x82, 0x52, 0x3f, 0x0c, 0xfb, 0x3e, 0x3b, 0xa0,
	0x51, 0xc4, 0x82, 0xbe, 0x17, 0xb0, 0x83, 0xe1, 0xa1, 0x14, 0xc8, 0x88, 0x0e, 0xfc, 0x72, 0xc4,
	0x43, 0x11, 0xa2, 0xad, 0x14, 0x53, 0x9e, 0x60, 0xca, 0xc3, 0xc3, 0x9d, 0x47, 0x13, 0xa2, 0x77,
	0x40, 0x83, 0x20, 0x14, 0x54, 0x78, 0x61, 0x10, 0xa7, 0x94, 0x9d, 0xcf, 0xb5, 0x55, 0x49, 0xe7,
	0xc9, 0xc5, 0x41, 0x2f, 0xe1, 0x0a, 0x90, 0xda, 0x4b, 0x7f, 0x66, 0xc0, 0xb4, 0x22, 0xaf, 0x12,
	0x06, 0x17, 0x5e, 0xbf, 0x46, 0x83, 0x9e, 0xcf, 0x38, 0x3a, 0x05, 0x93, 0x26, 0xc2, 0x25, 0x17,
	0xd4, 0xf3, 0x09, 0xed, 0x4a, 0x78, 0xd1, 0xd8, 0x35, 0x9e, 0x17, 0x8e, 0x9e, 0x96, 0x17, 0x5c,
	0xa1, 0x6c, 0x25, 0xc2, 0x3d, 0xa1, 0x9e, 0x6f, 0x29, 0x28, 0x2e, 0xd0, 0x39, 0x19, 0x7d, 0x0f,
	0xcb, 0x7e, 0xd8, 0xf7, 0x82, 0x62, 0x46, 0xf9, 0xf8, 0x72, 0xa1, 0x8f, 0xba, 0x44, 0x60, 0xf6,
	0x47, 0xe2, 0x71, 0x36, 0x60, 0x81, 0xc0, 0x29, 0x07, 0x3d, 0x80, 0x95, 0xb8, 0xcb, 0xbd, 0x48,
	0x14, 0xb3, 0xbb, 0xc6, 0xf3, 0x75, 0xac, 0x25, 0xe4, 0x40, 0x21, 0x66, 0xdd, 0x84, 0x7b, 0x62,
	0x44, 0x7c, 0x36, 0x64, 0x7e, 0x71, 0x49, 0x79, 0x2f, 0x2d, 0xf4, 0xde, 0xd2, 0xd0, 0xba, 0x44,
	0xe2, 0x8d, 0x78, 0x56, 0x44, 0x26, 0x64, 0x13, 0xee, 0x17, 0x97, 0x95, 0x7f, 0x79, 0x2c, 0x7d,
	0xcc, 0x40, 0xde, 0xe6, 0x3c, 0xe4, 0xe3, 0x8c, 0xfc, 0x02, 0xc0, 0xa4, 0x4c, 0xba, 0x61, 0x8f,
	0xe9, 0x5c, 0xec, 0x2f, 0xfc, 0xd2, 0x2c, 0x2d, 0x15, 0x2a, 0x61, 0x8f, 0xe1, 0x75, 0x36, 0x3e,
	0xa2, 0x27, 0x90, 0x8b, 0x65, 0x91, 0xba, 0xe4, 0xc2, 0xf3, 0x99, 0x4a, 0xca, 0x3a, 0x86, 0x54,
	0x75, 0xe2, 0xf9, 0x0c, 0x3d, 0x84, 0xf5, 0x81, 0x37, 0x60, 0x44, 0x8c, 0x22, 0xa6, 0xa3, 0x5e,
	0x93, 0x8a, 0xf6, 0x28, 0x62, 0xa5, 0xbf, 0x0c, 0x58, 0x9f, 0xb8, 0x45, 0x3b, 0xf0, 0xc0, 0xc6,
	0xb8, 0x81, 0x49, 0xa5, 0x51, 0xb5, 0x49, 0xe7, 0xac, 0xd5, 0xb4, 0x2b, 0xce, 0x89, 0x63, 0x57,
	0xcd, 0x3b, 0xe8, 0x01, 0xa0, 0x19, 0x5b, 0xd5, 0x3e, 0xb1, 0x3a, 0xf5, 0xb6, 0x79, 0x07, 0x6d,
	0xc3, 0xfd, 0x19, 0x7d, 0xe3, 0xad, 0x8d, 0xc9, 0x9b, 0x4e, 0xa3, 0x6d, 0x99, 0x06, 0x7a, 0x0c,
	0xdb, 0xb3, 0x94, 0x46, 0x8b, 0x58, 0x4d, 0x87, 0x54, 0xed, 0x33, 0xc7, 0xaa, 0x9b, 0x99, 0x4b,
	0x1e, 0xdb, 0xce, 0xa9, 0xdd, 0xe8, 0xb4, 0xcd, 0xec, 0x4e, 0xc6, 0x34, 0x4a, 0x7f, 0xaf, 0xc0,
	0x4a, 0x87, 0xfb, 0xa7, 0x34, 0x92, 0xf7, 0x4f, 0xb8, 0x4f, 0x38, 0xeb, 0xb3, 0x0f, 0x2a, 0x57,
	0xeb, 0x78, 0x2d, 0xe1, 0x3e, 0x96, 0x32, 0xaa, 0x43, 0x7e, 0x26, 0xfa, 0x58, 0x85, 0x9f, 0x3b,
	0x7a, 0xb6, 0xb8, 0x6a, 0x93, 0x9c, 0xc4, 0x3a, 0xa3, 0xb5, 0x3b, 0x38, 0x37, 0xcd, 0x54, 0x8c,
	0x7e, 0x98, 0xeb, 0x8e, 0xdc, 0x75, 0xd5, 0x57, 0x90, 0xa9, 0x8b, 0x71, 0x0f, 0xd5, 0x21, 0x4f,
	0x23, 0x8f, 0xb0, 0xa0, 0x17, 0x85, 0x5e, 0x20, 0x54, 0x07, 0x5d, 0x77, 0x17, 0x2b, 0xf2, 0x6c,
	0x8d, 0x9b, 0xb9, 0x0b, 0x9d, 0x6a, 0x17, 0x74, 0xe4, 0xf2, 0x6d, 0x3b, 0x72, 0x32, 0x31, 0x2b,
	0xb7, 0x98, 0x98, 0x45, 0xd3, 0xbb, 0x7a, 0xfb, 0xe9, 0x8d, 0xe1, 0x21, 0x67, 0x3d, 0x8f, 0xb3,
	0xae, 0x20, 0xae, 0x10, 0x11, 0xe1, 0x2c, 0x8e, 0xc2, 0x20, 0x66, 0xe9, 0x2c, 0xac, 0x29, 0xcf,
	0x5f, 0x2f, 0xf4, 0x9c, 0xf6, 0x43, 0x19, 0x6b, 0x7a, 0x4d, 0x88, 0x08, 0x6b, 0xb2, 0x1a, 0x8a,
	0x22, 0xbf, 0xc6, 0x52, 0xfa, 0xd7, 0x80, 0xe2, 0x75, 0x34, 0xb4, 0x0f, 0xcf, 0xb0, 0x5d, 0x75,
	0xb0, 0x5d, 0x69, 0x93, 0x5a, 0xbb, 0xdd, 0x24, 0xd8, 0x6e, 0x35, 0x1b, 0x67, 0x2d, 0x7b, 0xd1,
	0x14, 0x3c, 0x85, 0x27, 0x9f, 0x02, 0x1f, 0xbf, 0x38, 0x34, 0x8d, 0x9b, 0x41, 0x47, 0x66, 0xe6,
	0x66, 0xd0, 0xb1, 0x99, 0xbd, 0x19, 0xf4, 0xd2, 0x5c, 0x7a, 0x5d, 0x80, 0xbc, 0x9b, 0xf6, 0x90,
	0x9a, 0xf1, 0xd2, 0xc7, 0x2c, 0xa0, 0xab, 0xbd, 0x8e, 0x10, 0x2c, 0x45, 0x54, 0xb8, 0x7a, 0x84,
	0xd4, 0x19, 0xed, 0xc1, 0x66, 0x12, 0xf9, 0x21, 0xed, 0x11, 0x29, 0xea, 0x19, 0x4b, 0x57, 0xc8,
	0xdd, 0xd4, 0xd0, 0xa4, 0xc2, 0x4d, 0x47, 0xed, 0x77, 0xc8, 0xab, 0x82, 0xb9, 0x8c, 0xf6, 0x18,
	0x8f, 0x8b, 0xd9, 0xdd, 0xec, 0xf3, 0xdc, 0xd1, 0xab, 0xff, 0x39, 0x6a, 0x65, 0x99, 0xf7, 0x5a,
	0x4a, 0xb5, 0x03, 0xc1, 0x47, 0x38, 0xe7, 0x4e, 0x35, 0xf3, 0x4b, 0x6a, 0x69, 0x7e, 0x49, 0xa1,
	0x6f, 0x01, 0xd8, 0x87, 0xc8, 0x4b, 0x5f, 0x1a, 0x35, 0x06, 0xb9, 0xa3, 0xed, 0xf1, 0x77, 0xc7,
	0x4f, 0x51, 0xb9, 0xaa, 0x9f, 0x22, 0x3c, 0x03, 0x46, 0x47, 0x70, 0x9f, 0xa7, 0x3d, 0x4d, 0x06,
	0x54, 0x74, 0x5d, 0x2f, 0xe8, 0xa7, 0x7b, 0x52, 0x8e, 0xc2, 0x1a, 0xde, 0xd2, 0xc6, 0x53, 0x6d,
	0x53, 0x0b, 0xf3, 0x10, 0xee, 0xd1, 0x28, 0xf2, 0xbd, 0xae, 0x72, 0x41, 0x38, 0xa3, 0x3d, 0x7a,
	0xee, 0x33, 0xd5, 0xf5, 0x6b, 0x78, 0x6b, 0xc6, 0x86, 0xb5, 0x69, 0xe7, 0x47, 0x30, 0x2f, 0xc7,
	0x27, 0xdf, 0x81, 0xf7, 0x6c, 0xa4, 0xd3, 0x2d, 0x8f, 0xe8, 0x1e, 0x2c, 0x0f, 0xa9, 0x9f, 0x8c,
	0x97, 0x74, 0x2a, 0x7c, 0x97, 0x79, 0x65, 0x94, 0x5e, 0xc0, 0xc6, 0xdc, 0x56, 0x51, 0x5b, 0x5d,
	0x29, 0xc8, 0x4c, 0xcd, 0x20, 0x55, 0xc9, 0x92, 0x94, 0xbe, 0x01, 0x74, 0x75, 0x87, 0xdc, 0x4c,
	0xfb, 0x27, 0x03, 0xb9, 0x1a, 0xa3, 0xbe, 0x70, 0x2b, 0x2e, 0xeb, 0xbe, 0x47, 0x2f, 0xe0, 0x5e,
	0xcf, 0x8b, 0x65, 0x0c, 0xb2, 0xae, 0xbe, 0x70, 0x49, 0x57, 0xea, 0x15, 0x73, 0x0d, 0x23, 0x6d,
	0x9b, 0x65, 0x20, 0x58, 0x72, 0xc3, 0x58, 0xe8, 0x18, 0xd4, 0x19, 0xed, 0xc3, 0x66, 0xca, 0x1e,
	0x11, 0xe1, 0x72, 0x16, 0xbb, 0xa1, 0xdf, 0x53, 0x2b, 0x74, 0x03, 0x9b, 0xda, 0xd0, 0x1e, 0xeb,
	0xd1, 0x01, 0x6c, 0x25, 0xc1, 0x55, 0xf8, 0x92, 0x82, 0xa3, 0x89, 0x69, 0x4a, 0xd8, 0x87, 0x4d,
	0xce, 0x62, 0x41, 0xb9, 0x98, 0x81, 0x2f, 0xa7, 0xde, 0xb5, 0x61, 0x0a, 0xfe, 0x09, 0x0a, 0x2a,
	0x02, 0xe2, 0x05, 0x82, 0xf1, 0x21, 0xf5, 0x55, 0xa5, 0x3f, 0xd9, 0x2f, 0x1b, 0x8a, 0xe0, 0x68,
	0x3c, 0x3a, 0x86, 0x55, 0xe1, 0x0d, 0x58, 0x98, 0x08, 0x55, 0xf1, 0x4f, 0x52, 0xc7, 0xc8, 0xd2,
	0x4b, 0x58, 0xad, 0x7b, 0xe7, 0x9c, 0xf2, 0x91, 0x4c, 0x50, 0x40, 0x07, 0x6c, 0x3c, 0x67, 0xf2,
	0x8c, 0x8a, 0xb0, 0x3a, 0x64, 0x3c, 0x96, 0xed, 0x9b, 0xe6, 0x6d, 0x2c, 0xee, 0x09, 0x28, 0xcc,
	0x6f, 0x4c, 0xb4, 0x0b, 0x8f, 0xac, 0x4e, 0xbb, 0x46, 0x4e, 0x2c, 0xa7, 0x4e, 0xac, 0x4a, 0xdb,
	0x69, 0x9c, 0x5d, 0x5a, 0x42, 0x8f, 0x61, 0xfb, 0x0a, 0x62, 0xbc, 0x26, 0x4c, 0x03, 0x7d, 0x01,
	0x8f, 0x17, 0x38, 0x90, 0xaa, 0x06, 0x76, 0x7e, 0xb3, 0xab, 0x66, 0x66, 0xef, 0x1c, 0xcc, 0xcb,
	0xfb, 0x1e, 0xdd, 0x87, 0xcd, 0x7a, 0xe3, 0x67, 0xe7, 0xf2, 0xc7, 0x10, 0x14, 0x52, 0x75, 0xa3,
	0x29, 0x3d, 0x59, 0x75, 0xd3, 0x40, 0x77, 0x21, 0x97, 0xea, 0xac, 0xea, 0xa9, 0x73, 0x66, 0x66,
	0xa6, 0x20, 0x6c, 0xbf, 0xe9, 0x38, 0xd8, 0xae, 0x9a, 0xd9, 0xbd, 0x11, 0x6c, 0xcc, 0xbd, 0x4a,
	0xf2, 0xbd, 0x6f, 0xd9, 0x95, 0x0e, 0xb6, 0xaf, 0x7e, 0x41, 0xeb, 0xa7, 0x7f, 0x15, 0x26, 0xe4,
	0xb5, 0xee, 0xcc, 0x7e, 0x6b, 0x63, 0xd3, 0x40, 0x5b, 0x70, 0x57, 0x6b, 0x26, 0x17, 0xc9, 0xa0,
	0x4d, 0xd8, 0xd0, 0x4a, 0xab, 0xfe, 0xab, 0xf5, 0xae, 0x95, 0xfe, 0x3d, 0xbc, 0xfe, 0x0a, 0x3e,
	0xeb, 0x86, 0x83, 0x45, 0x9b, 0xe9, 0x75, 0xde, 0x8a, 0xa2, 0x77, 0x74, 0xe0, 0x37, 0x65, 0x2d,
	0x9b, 0xc6, 0xf9, 0x8a, 0x2a, 0xea, 0xf1, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xda, 0x2d, 0x20,
	0xdc, 0x31, 0x0b, 0x00, 0x00,
}
