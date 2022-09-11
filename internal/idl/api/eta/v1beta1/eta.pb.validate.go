// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/eta/v1beta1/eta.proto

package v1beta1

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

// define the regex for a UUID once up-front
var _eta_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetEstimatedJobsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetEstimatedJobsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetEstimatedJobsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetEstimatedJobsRequestMultiError, or nil if none found.
func (m *GetEstimatedJobsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetEstimatedJobsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetAreaKey()) < 1 {
		err := GetEstimatedJobsRequestValidationError{
			field:  "AreaKey",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateUuid(m.GetWorkerId()); err != nil {
		err = GetEstimatedJobsRequestValidationError{
			field:  "WorkerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetRadiusMeters() <= 0 {
		err := GetEstimatedJobsRequestValidationError{
			field:  "RadiusMeters",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetLimit() > 25 {
		err := GetEstimatedJobsRequestValidationError{
			field:  "Limit",
			reason: "value must be less than or equal to 25",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := Type_name[int32(m.GetEta())]; !ok {
		err := GetEstimatedJobsRequestValidationError{
			field:  "Eta",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetEstimatedJobsRequestMultiError(errors)
	}

	return nil
}

func (m *GetEstimatedJobsRequest) _validateUuid(uuid string) error {
	if matched := _eta_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetEstimatedJobsRequestMultiError is an error wrapping multiple validation
// errors returned by GetEstimatedJobsRequest.ValidateAll() if the designated
// constraints aren't met.
type GetEstimatedJobsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetEstimatedJobsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetEstimatedJobsRequestMultiError) AllErrors() []error { return m }

// GetEstimatedJobsRequestValidationError is the validation error returned by
// GetEstimatedJobsRequest.Validate if the designated constraints aren't met.
type GetEstimatedJobsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEstimatedJobsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEstimatedJobsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEstimatedJobsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEstimatedJobsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEstimatedJobsRequestValidationError) ErrorName() string {
	return "GetEstimatedJobsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetEstimatedJobsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEstimatedJobsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEstimatedJobsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEstimatedJobsRequestValidationError{}

// Validate checks the field values on Location with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Location) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Location with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LocationMultiError, or nil
// if none found.
func (m *Location) ValidateAll() error {
	return m.validate(true)
}

func (m *Location) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Lat

	// no validation rules for Lng

	// no validation rules for Address

	if len(errors) > 0 {
		return LocationMultiError(errors)
	}

	return nil
}

// LocationMultiError is an error wrapping multiple validation errors returned
// by Location.ValidateAll() if the designated constraints aren't met.
type LocationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocationMultiError) AllErrors() []error { return m }

// LocationValidationError is the validation error returned by
// Location.Validate if the designated constraints aren't met.
type LocationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocationValidationError) ErrorName() string { return "LocationValidationError" }

// Error satisfies the builtin error interface
func (e LocationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocationValidationError{}

// Validate checks the field values on Estimate with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Estimate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Estimate with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EstimateMultiError, or nil
// if none found.
func (m *Estimate) ValidateAll() error {
	return m.validate(true)
}

func (m *Estimate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DistanceMeters

	if all {
		switch v := interface{}(m.GetDuration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EstimateValidationError{
					field:  "Duration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EstimateValidationError{
					field:  "Duration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EstimateValidationError{
				field:  "Duration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return EstimateMultiError(errors)
	}

	return nil
}

// EstimateMultiError is an error wrapping multiple validation errors returned
// by Estimate.ValidateAll() if the designated constraints aren't met.
type EstimateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EstimateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EstimateMultiError) AllErrors() []error { return m }

// EstimateValidationError is the validation error returned by
// Estimate.Validate if the designated constraints aren't met.
type EstimateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EstimateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EstimateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EstimateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EstimateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EstimateValidationError) ErrorName() string { return "EstimateValidationError" }

// Error satisfies the builtin error interface
func (e EstimateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEstimate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EstimateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EstimateValidationError{}

// Validate checks the field values on EstimatedJob with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EstimatedJob) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EstimatedJob with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EstimatedJobMultiError, or
// nil if none found.
func (m *EstimatedJob) ValidateAll() error {
	return m.validate(true)
}

func (m *EstimatedJob) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetWorkerToPickupEstimate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EstimatedJobValidationError{
					field:  "WorkerToPickupEstimate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EstimatedJobValidationError{
					field:  "WorkerToPickupEstimate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWorkerToPickupEstimate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EstimatedJobValidationError{
				field:  "WorkerToPickupEstimate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPickupToDropOffEstimate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EstimatedJobValidationError{
					field:  "PickupToDropOffEstimate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EstimatedJobValidationError{
					field:  "PickupToDropOffEstimate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPickupToDropOffEstimate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EstimatedJobValidationError{
				field:  "PickupToDropOffEstimate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetWorkerLocation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EstimatedJobValidationError{
					field:  "WorkerLocation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EstimatedJobValidationError{
					field:  "WorkerLocation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWorkerLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EstimatedJobValidationError{
				field:  "WorkerLocation",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPickupLocation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EstimatedJobValidationError{
					field:  "PickupLocation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EstimatedJobValidationError{
					field:  "PickupLocation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPickupLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EstimatedJobValidationError{
				field:  "PickupLocation",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDropOffLocation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EstimatedJobValidationError{
					field:  "DropOffLocation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EstimatedJobValidationError{
					field:  "DropOffLocation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDropOffLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EstimatedJobValidationError{
				field:  "DropOffLocation",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return EstimatedJobMultiError(errors)
	}

	return nil
}

// EstimatedJobMultiError is an error wrapping multiple validation errors
// returned by EstimatedJob.ValidateAll() if the designated constraints aren't met.
type EstimatedJobMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EstimatedJobMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EstimatedJobMultiError) AllErrors() []error { return m }

// EstimatedJobValidationError is the validation error returned by
// EstimatedJob.Validate if the designated constraints aren't met.
type EstimatedJobValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EstimatedJobValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EstimatedJobValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EstimatedJobValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EstimatedJobValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EstimatedJobValidationError) ErrorName() string { return "EstimatedJobValidationError" }

// Error satisfies the builtin error interface
func (e EstimatedJobValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEstimatedJob.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EstimatedJobValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EstimatedJobValidationError{}

// Validate checks the field values on GetEstimatedJobsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetEstimatedJobsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetEstimatedJobsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetEstimatedJobsResponseMultiError, or nil if none found.
func (m *GetEstimatedJobsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetEstimatedJobsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetJobs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetEstimatedJobsResponseValidationError{
						field:  fmt.Sprintf("Jobs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetEstimatedJobsResponseValidationError{
						field:  fmt.Sprintf("Jobs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetEstimatedJobsResponseValidationError{
					field:  fmt.Sprintf("Jobs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetEstimatedJobsResponseMultiError(errors)
	}

	return nil
}

// GetEstimatedJobsResponseMultiError is an error wrapping multiple validation
// errors returned by GetEstimatedJobsResponse.ValidateAll() if the designated
// constraints aren't met.
type GetEstimatedJobsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetEstimatedJobsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetEstimatedJobsResponseMultiError) AllErrors() []error { return m }

// GetEstimatedJobsResponseValidationError is the validation error returned by
// GetEstimatedJobsResponse.Validate if the designated constraints aren't met.
type GetEstimatedJobsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEstimatedJobsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEstimatedJobsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEstimatedJobsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEstimatedJobsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEstimatedJobsResponseValidationError) ErrorName() string {
	return "GetEstimatedJobsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetEstimatedJobsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEstimatedJobsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEstimatedJobsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEstimatedJobsResponseValidationError{}