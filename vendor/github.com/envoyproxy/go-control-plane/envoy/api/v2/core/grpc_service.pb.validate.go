// Code generated by protoc-gen-validate
// source: envoy/api/v2/core/grpc_service.proto
// DO NOT EDIT!!!

package core

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on GrpcService with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GrpcService) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrpcServiceValidationError{
				Field:  "Timeout",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetInitialMetadata() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GrpcServiceValidationError{
					Field:  fmt.Sprintf("InitialMetadata[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	switch m.TargetSpecifier.(type) {

	case *GrpcService_EnvoyGrpc_:

		if v, ok := interface{}(m.GetEnvoyGrpc()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GrpcServiceValidationError{
					Field:  "EnvoyGrpc",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *GrpcService_GoogleGrpc_:

		if v, ok := interface{}(m.GetGoogleGrpc()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GrpcServiceValidationError{
					Field:  "GoogleGrpc",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return GrpcServiceValidationError{
			Field:  "TargetSpecifier",
			Reason: "value is required",
		}

	}

	return nil
}

// GrpcServiceValidationError is the validation error returned by
// GrpcService.Validate if the designated constraints aren't met.
type GrpcServiceValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e GrpcServiceValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcService.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = GrpcServiceValidationError{}

// Validate checks the field values on GrpcService_EnvoyGrpc with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GrpcService_EnvoyGrpc) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetClusterName()) < 1 {
		return GrpcService_EnvoyGrpcValidationError{
			Field:  "ClusterName",
			Reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// GrpcService_EnvoyGrpcValidationError is the validation error returned by
// GrpcService_EnvoyGrpc.Validate if the designated constraints aren't met.
type GrpcService_EnvoyGrpcValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e GrpcService_EnvoyGrpcValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcService_EnvoyGrpc.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = GrpcService_EnvoyGrpcValidationError{}

// Validate checks the field values on GrpcService_GoogleGrpc with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GrpcService_GoogleGrpc) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetTargetUri()) < 1 {
		return GrpcService_GoogleGrpcValidationError{
			Field:  "TargetUri",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetChannelCredentials()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrpcService_GoogleGrpcValidationError{
				Field:  "ChannelCredentials",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetCallCredentials() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GrpcService_GoogleGrpcValidationError{
					Field:  fmt.Sprintf("CallCredentials[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if len(m.GetStatPrefix()) < 1 {
		return GrpcService_GoogleGrpcValidationError{
			Field:  "StatPrefix",
			Reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for CredentialsFactoryName

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrpcService_GoogleGrpcValidationError{
				Field:  "Config",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// GrpcService_GoogleGrpcValidationError is the validation error returned by
// GrpcService_GoogleGrpc.Validate if the designated constraints aren't met.
type GrpcService_GoogleGrpcValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e GrpcService_GoogleGrpcValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcService_GoogleGrpc.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = GrpcService_GoogleGrpcValidationError{}

// Validate checks the field values on GrpcService_GoogleGrpc_SslCredentials
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *GrpcService_GoogleGrpc_SslCredentials) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRootCerts()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrpcService_GoogleGrpc_SslCredentialsValidationError{
				Field:  "RootCerts",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPrivateKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrpcService_GoogleGrpc_SslCredentialsValidationError{
				Field:  "PrivateKey",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCertChain()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrpcService_GoogleGrpc_SslCredentialsValidationError{
				Field:  "CertChain",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// GrpcService_GoogleGrpc_SslCredentialsValidationError is the validation error
// returned by GrpcService_GoogleGrpc_SslCredentials.Validate if the
// designated constraints aren't met.
type GrpcService_GoogleGrpc_SslCredentialsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e GrpcService_GoogleGrpc_SslCredentialsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcService_GoogleGrpc_SslCredentials.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = GrpcService_GoogleGrpc_SslCredentialsValidationError{}

// Validate checks the field values on
// GrpcService_GoogleGrpc_ChannelCredentials with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GrpcService_GoogleGrpc_ChannelCredentials) Validate() error {
	if m == nil {
		return nil
	}

	switch m.CredentialSpecifier.(type) {

	case *GrpcService_GoogleGrpc_ChannelCredentials_SslCredentials:

		if v, ok := interface{}(m.GetSslCredentials()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GrpcService_GoogleGrpc_ChannelCredentialsValidationError{
					Field:  "SslCredentials",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *GrpcService_GoogleGrpc_ChannelCredentials_GoogleDefault:

		if v, ok := interface{}(m.GetGoogleDefault()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GrpcService_GoogleGrpc_ChannelCredentialsValidationError{
					Field:  "GoogleDefault",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return GrpcService_GoogleGrpc_ChannelCredentialsValidationError{
			Field:  "CredentialSpecifier",
			Reason: "value is required",
		}

	}

	return nil
}

// GrpcService_GoogleGrpc_ChannelCredentialsValidationError is the validation
// error returned by GrpcService_GoogleGrpc_ChannelCredentials.Validate if the
// designated constraints aren't met.
type GrpcService_GoogleGrpc_ChannelCredentialsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e GrpcService_GoogleGrpc_ChannelCredentialsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcService_GoogleGrpc_ChannelCredentials.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = GrpcService_GoogleGrpc_ChannelCredentialsValidationError{}

// Validate checks the field values on GrpcService_GoogleGrpc_CallCredentials
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *GrpcService_GoogleGrpc_CallCredentials) Validate() error {
	if m == nil {
		return nil
	}

	switch m.CredentialSpecifier.(type) {

	case *GrpcService_GoogleGrpc_CallCredentials_AccessToken:
		// no validation rules for AccessToken

	case *GrpcService_GoogleGrpc_CallCredentials_GoogleComputeEngine:

		if v, ok := interface{}(m.GetGoogleComputeEngine()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GrpcService_GoogleGrpc_CallCredentialsValidationError{
					Field:  "GoogleComputeEngine",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_GoogleRefreshToken:
		// no validation rules for GoogleRefreshToken

	case *GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJwtAccess:

		if v, ok := interface{}(m.GetServiceAccountJwtAccess()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GrpcService_GoogleGrpc_CallCredentialsValidationError{
					Field:  "ServiceAccountJwtAccess",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_GoogleIam:

		if v, ok := interface{}(m.GetGoogleIam()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GrpcService_GoogleGrpc_CallCredentialsValidationError{
					Field:  "GoogleIam",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *GrpcService_GoogleGrpc_CallCredentials_FromPlugin:

		if v, ok := interface{}(m.GetFromPlugin()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GrpcService_GoogleGrpc_CallCredentialsValidationError{
					Field:  "FromPlugin",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return GrpcService_GoogleGrpc_CallCredentialsValidationError{
			Field:  "CredentialSpecifier",
			Reason: "value is required",
		}

	}

	return nil
}

// GrpcService_GoogleGrpc_CallCredentialsValidationError is the validation
// error returned by GrpcService_GoogleGrpc_CallCredentials.Validate if the
// designated constraints aren't met.
type GrpcService_GoogleGrpc_CallCredentialsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e GrpcService_GoogleGrpc_CallCredentialsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcService_GoogleGrpc_CallCredentials.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = GrpcService_GoogleGrpc_CallCredentialsValidationError{}

// Validate checks the field values on
// GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for JsonKey

	// no validation rules for TokenLifetimeSeconds

	return nil
}

// GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentialsValidationError
// is the validation error returned by
// GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials.Validate
// if the designated constraints aren't met.
type GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentialsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentialsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentialsValidationError{}

// Validate checks the field values on
// GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AuthorizationToken

	// no validation rules for AuthoritySelector

	return nil
}

// GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentialsValidationError
// is the validation error returned by
// GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials.Validate if the
// designated constraints aren't met.
type GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentialsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentialsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentials.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = GrpcService_GoogleGrpc_CallCredentials_GoogleIAMCredentialsValidationError{}

// Validate checks the field values on
// GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPluginValidationError{
				Field:  "Config",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPluginValidationError
// is the validation error returned by
// GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin.Validate
// if the designated constraints aren't met.
type GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPluginValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPluginValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPlugin.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = GrpcService_GoogleGrpc_CallCredentials_MetadataCredentialsFromPluginValidationError{}
