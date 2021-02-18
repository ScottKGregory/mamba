package config

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// Int
type ParseIntSetsDefault struct {
	IntTest int `config:"12,The int to test"`
}

func TestParseIntSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseIntSetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetInt("inttest")

	assert.Equal(t, 12, i)
	assert.Nil(t, err)
}

type ParseIntInvalidDefaultReturnsError struct {
	IntTest int `config:"FOO,The int to test"`
}

func TestParseIntInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseIntInvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// String
type ParseStringSetsDefault struct {
	StringTest string `config:"test string,The String to test"`
}

func TestParseStringSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseStringSetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetString("stringtest")

	assert.Equal(t, "test string", i)
	assert.Nil(t, err)
}

// Float64
type ParseFloat64SetsDefault struct {
	Float64Test float64 `config:"24,The Float64 to test"`
}

func TestParseFloat64SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseFloat64SetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetFloat64("float64test")

	assert.Equal(t, float64(24), i)
	assert.Nil(t, err)
}

type ParseFloat64InvalidDefaultReturnsError struct {
	Float64Test float64 `config:"FOO,The Float64 to test"`
}

func TestParseFloat64InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseFloat64InvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Float32
type ParseFloat32SetsDefault struct {
	Float32Test float32 `config:"36,The Float32 to test"`
}

func TestParseFloat32SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseFloat32SetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetFloat32("float32test")

	assert.Equal(t, float32(36), i)
	assert.Nil(t, err)
}

type ParseFloat32InvalidDefaultReturnsError struct {
	Float32Test float32 `config:"FOO,The Float32 to test"`
}

func TestParseFloat32InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseFloat32InvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Int8
type ParseInt8SetsDefault struct {
	Int8Test int8 `config:"14,The int8 to test"`
}

func TestParseInt8SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseInt8SetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetInt8("int8test")

	assert.Equal(t, int8(14), i)
	assert.Nil(t, err)
}

type ParseInt8InvalidDefaultReturnsError struct {
	Int8Test int8 `config:"FOO,The Int8 to test"`
}

func TestParseInt8InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseInt8InvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Int16
type ParseInt16SetsDefault struct {
	Int16Test int16 `config:"16,The int16 to test"`
}

func TestParseInt16SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseInt16SetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetInt16("int16test")

	assert.Equal(t, int16(16), i)
	assert.Nil(t, err)
}

type ParseInt16InvalidDefaultReturnsError struct {
	Int16Test int16 `config:"FOO,The Int16 to test"`
}

func TestParseInt16InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseInt16InvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Int32
type ParseInt32SetsDefault struct {
	Int32Test int32 `config:"32,The int32 to test"`
}

func TestParseInt32SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseInt32SetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetInt32("int32test")

	assert.Equal(t, int32(32), i)
	assert.Nil(t, err)
}

type ParseInt32InvalidDefaultReturnsError struct {
	Int32Test int32 `config:"FOO,The Int32 to test"`
}

func TestParseInt32InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseInt32InvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Int64
type ParseInt64SetsDefault struct {
	Int64Test int64 `config:"64,The int64 to test"`
}

func TestParseInt64SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseInt64SetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetInt64("int64test")

	assert.Equal(t, int64(64), i)
	assert.Nil(t, err)
}

type ParseInt64InvalidDefaultReturnsError struct {
	Int64Test int64 `config:"FOO,The Int64 to test"`
}

func TestParseInt64InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseInt64InvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Bool
type ParseBoolSetsDefault struct {
	BoolTest bool `config:"true,The Bool to test"`
}

func TestParseBoolSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseBoolSetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetBool("booltest")

	assert.Equal(t, true, i)
	assert.Nil(t, err)
}

type ParseBoolInvalidDefaultReturnsError struct {
	BoolTest bool `config:"FOO,The Bool to test"`
}

func TestParseBoolInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(ParseBoolInvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Tag parsing
type InvalidTagReturnsError1 struct {
	TagTest string `config:""`
}

func TestInvalidTagReturnsError1(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(InvalidTagReturnsError1{}, cmd)

	assert.NotNil(t, err)
}

type InvalidTagReturnsError2 struct {
	TagTest string `config:"default"`
}

func TestInvalidTagReturnsError2(t *testing.T) {
	cmd := &cobra.Command{}
	err := Parse(InvalidTagReturnsError2{}, cmd)

	assert.NotNil(t, err)
}
