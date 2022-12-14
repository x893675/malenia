// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cr/cr.proto

package cr

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

// Validate checks the field values on GetRepoRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetRepoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRepoRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetRepoRequestMultiError,
// or nil if none found.
func (m *GetRepoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRepoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetName()) > 256 {
		err := GetRepoRequestValidationError{
			field:  "Name",
			reason: "value length must be at most 256 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_GetRepoRequest_Name_Pattern.MatchString(m.GetName()) {
		err := GetRepoRequestValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\\\-]*[a-zA-Z0-9])\\\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\\\-]*[A-Za-z0-9])$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetRepoRequestMultiError(errors)
	}

	return nil
}

// GetRepoRequestMultiError is an error wrapping multiple validation errors
// returned by GetRepoRequest.ValidateAll() if the designated constraints
// aren't met.
type GetRepoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRepoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRepoRequestMultiError) AllErrors() []error { return m }

// GetRepoRequestValidationError is the validation error returned by
// GetRepoRequest.Validate if the designated constraints aren't met.
type GetRepoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRepoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRepoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRepoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRepoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRepoRequestValidationError) ErrorName() string { return "GetRepoRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRepoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRepoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRepoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRepoRequestValidationError{}

var _GetRepoRequest_Name_Pattern = regexp.MustCompile("^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])$")

// Validate checks the field values on Repo with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Repo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Repo with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RepoMultiError, or nil if none found.
func (m *Repo) ValidateAll() error {
	return m.validate(true)
}

func (m *Repo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetName()) > 256 {
		err := RepoValidationError{
			field:  "Name",
			reason: "value length must be at most 256 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Repo_Name_Pattern.MatchString(m.GetName()) {
		err := RepoValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\\\-]*[a-zA-Z0-9])\\\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\\\-]*[A-Za-z0-9])$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Visibility

	if len(errors) > 0 {
		return RepoMultiError(errors)
	}

	return nil
}

// RepoMultiError is an error wrapping multiple validation errors returned by
// Repo.ValidateAll() if the designated constraints aren't met.
type RepoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RepoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RepoMultiError) AllErrors() []error { return m }

// RepoValidationError is the validation error returned by Repo.Validate if the
// designated constraints aren't met.
type RepoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RepoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RepoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RepoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RepoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RepoValidationError) ErrorName() string { return "RepoValidationError" }

// Error satisfies the builtin error interface
func (e RepoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRepo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RepoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RepoValidationError{}

var _Repo_Name_Pattern = regexp.MustCompile("^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])$")

// Validate checks the field values on CreateRepoRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateRepoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRepoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateRepoRequestMultiError, or nil if none found.
func (m *CreateRepoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRepoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRepo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateRepoRequestValidationError{
					field:  "Repo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateRepoRequestValidationError{
					field:  "Repo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRepo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRepoRequestValidationError{
				field:  "Repo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateRepoRequestMultiError(errors)
	}

	return nil
}

// CreateRepoRequestMultiError is an error wrapping multiple validation errors
// returned by CreateRepoRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateRepoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRepoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRepoRequestMultiError) AllErrors() []error { return m }

// CreateRepoRequestValidationError is the validation error returned by
// CreateRepoRequest.Validate if the designated constraints aren't met.
type CreateRepoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRepoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRepoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRepoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRepoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRepoRequestValidationError) ErrorName() string {
	return "CreateRepoRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRepoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRepoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRepoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRepoRequestValidationError{}

// Validate checks the field values on ListReposResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListReposResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListReposResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListReposResponseMultiError, or nil if none found.
func (m *ListReposResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListReposResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetRepos() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListReposResponseValidationError{
						field:  fmt.Sprintf("Repos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListReposResponseValidationError{
						field:  fmt.Sprintf("Repos[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListReposResponseValidationError{
					field:  fmt.Sprintf("Repos[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListReposResponseMultiError(errors)
	}

	return nil
}

// ListReposResponseMultiError is an error wrapping multiple validation errors
// returned by ListReposResponse.ValidateAll() if the designated constraints
// aren't met.
type ListReposResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListReposResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListReposResponseMultiError) AllErrors() []error { return m }

// ListReposResponseValidationError is the validation error returned by
// ListReposResponse.Validate if the designated constraints aren't met.
type ListReposResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListReposResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListReposResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListReposResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListReposResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListReposResponseValidationError) ErrorName() string {
	return "ListReposResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListReposResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListReposResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListReposResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListReposResponseValidationError{}
