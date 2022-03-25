package internal

import (
	"fmt"
	"reflect"
)

var InvalidTypeError *invalidTypeError = &invalidTypeError{}
var BindError *bindError = &bindError{}
var ParseError *parseError = &parseError{}
var TagParseError *tagParseError = &tagParseError{}

type genericError struct {
	Kind          reflect.Kind
	FieldName     string
	InternalError error
}

type invalidTypeError struct {
	*genericError
}

func NewInvalidTypeError(kind reflect.Kind, fieldName string, err ...error) *invalidTypeError {
	var e error
	if len(err) > 0 {
		e = err[0]
	}

	return &invalidTypeError{&genericError{kind, fieldName, e}}
}

func (e *invalidTypeError) Error() string {
	if e.InternalError != nil {
		return fmt.Sprintf("invalid/unsupported type \"%s\" for \"%s\": %v", e.Kind.String(), e.FieldName, e.InternalError)
	}

	return fmt.Sprintf("invalid/unsupported type \"%s\" for \"%s\"", e.Kind.String(), e.FieldName)
}

func (e *invalidTypeError) Is(target error) bool {
	return reflect.TypeOf(target) == reflect.TypeOf(&invalidTypeError{})
}

type bindError struct {
	*genericError
}

func NewBindError(kind reflect.Kind, fieldName string, err ...error) *bindError {
	var e error
	if len(err) > 0 {
		e = err[0]
	}

	return &bindError{&genericError{kind, fieldName, e}}
}

func (e *bindError) Error() string {
	if e.InternalError != nil {
		return fmt.Sprintf("error binding \"%s\" for \"%s\": %v", e.Kind.String(), e.FieldName, e.InternalError)
	}

	return fmt.Sprintf("error binding \"%s\" for \"%s\"", e.Kind.String(), e.FieldName)
}

func (e *bindError) Is(target error) bool {
	return reflect.TypeOf(target) == reflect.TypeOf(&bindError{})
}

type parseError struct {
	*genericError
	Default string
}

func NewParseError(def string, kind reflect.Kind, fieldName string, err ...error) *parseError {
	var e error
	if len(err) > 0 {
		e = err[0]
	}

	return &parseError{&genericError{kind, fieldName, e}, def}
}

func (e *parseError) Error() string {
	if e.InternalError != nil {
		return fmt.Sprintf("error parsing default \"%s\" to type \"%s\" for \"%s\": %v", e.Default, e.Kind.String(), e.FieldName, e.InternalError)
	}

	return fmt.Sprintf("error parsing default \"%s\" to type \"%s\" for \"%s\"", e.Default, e.Kind.String(), e.FieldName)
}

func (e *parseError) Is(target error) bool {
	return reflect.TypeOf(target) == reflect.TypeOf(&parseError{})
}

type tagParseError struct {
	*genericError
	RawTag string
}

func NewTagParseError(rawTag string, kind reflect.Kind, fieldName string, err ...error) *tagParseError {
	var e error
	if len(err) > 0 {
		e = err[0]
	}

	return &tagParseError{&genericError{kind, fieldName, e}, rawTag}
}

func (e *tagParseError) Error() string {
	if e.InternalError != nil {
		return fmt.Sprintf("error parsing tag \"%s\" for \"%s\"(%s): %v", e.RawTag, e.FieldName, e.Kind.String(), e.InternalError)
	}

	return fmt.Sprintf("error parsing tag \"%s\" for \"%s\"(%s)", e.RawTag, e.FieldName, e.Kind.String())
}

func (e *tagParseError) Is(target error) bool {
	return reflect.TypeOf(target) == reflect.TypeOf(&tagParseError{})
}
