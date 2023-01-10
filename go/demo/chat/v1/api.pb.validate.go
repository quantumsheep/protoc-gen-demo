// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: demo/chat/v1/api.proto

package chatv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on SendMessageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SendMessageRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendMessageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendMessageRequestMultiError, or nil if none found.
func (m *SendMessageRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SendMessageRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return SendMessageRequestMultiError(errors)
	}

	return nil
}

// SendMessageRequestMultiError is an error wrapping multiple validation errors
// returned by SendMessageRequest.ValidateAll() if the designated constraints
// aren't met.
type SendMessageRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendMessageRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendMessageRequestMultiError) AllErrors() []error { return m }

// SendMessageRequestValidationError is the validation error returned by
// SendMessageRequest.Validate if the designated constraints aren't met.
type SendMessageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendMessageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendMessageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendMessageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendMessageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendMessageRequestValidationError) ErrorName() string {
	return "SendMessageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SendMessageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendMessageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendMessageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendMessageRequestValidationError{}

// Validate checks the field values on SendMessageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SendMessageResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendMessageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SendMessageResponseMultiError, or nil if none found.
func (m *SendMessageResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SendMessageResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SendMessageResponseMultiError(errors)
	}

	return nil
}

// SendMessageResponseMultiError is an error wrapping multiple validation
// errors returned by SendMessageResponse.ValidateAll() if the designated
// constraints aren't met.
type SendMessageResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendMessageResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendMessageResponseMultiError) AllErrors() []error { return m }

// SendMessageResponseValidationError is the validation error returned by
// SendMessageResponse.Validate if the designated constraints aren't met.
type SendMessageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendMessageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendMessageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendMessageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendMessageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendMessageResponseValidationError) ErrorName() string {
	return "SendMessageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SendMessageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendMessageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendMessageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendMessageResponseValidationError{}

// Validate checks the field values on ListMessagesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListMessagesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListMessagesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListMessagesRequestMultiError, or nil if none found.
func (m *ListMessagesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListMessagesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	if len(errors) > 0 {
		return ListMessagesRequestMultiError(errors)
	}

	return nil
}

// ListMessagesRequestMultiError is an error wrapping multiple validation
// errors returned by ListMessagesRequest.ValidateAll() if the designated
// constraints aren't met.
type ListMessagesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListMessagesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListMessagesRequestMultiError) AllErrors() []error { return m }

// ListMessagesRequestValidationError is the validation error returned by
// ListMessagesRequest.Validate if the designated constraints aren't met.
type ListMessagesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMessagesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMessagesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMessagesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMessagesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMessagesRequestValidationError) ErrorName() string {
	return "ListMessagesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListMessagesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMessagesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMessagesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMessagesRequestValidationError{}

// Validate checks the field values on ListMessagesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListMessagesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListMessagesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListMessagesResponseMultiError, or nil if none found.
func (m *ListMessagesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListMessagesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListMessagesResponseMultiError(errors)
	}

	return nil
}

// ListMessagesResponseMultiError is an error wrapping multiple validation
// errors returned by ListMessagesResponse.ValidateAll() if the designated
// constraints aren't met.
type ListMessagesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListMessagesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListMessagesResponseMultiError) AllErrors() []error { return m }

// ListMessagesResponseValidationError is the validation error returned by
// ListMessagesResponse.Validate if the designated constraints aren't met.
type ListMessagesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListMessagesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListMessagesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListMessagesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListMessagesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListMessagesResponseValidationError) ErrorName() string {
	return "ListMessagesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListMessagesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListMessagesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListMessagesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListMessagesResponseValidationError{}
