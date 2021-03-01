package mamba

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

// Int.
type BindIntSetsDefault struct {
	IntTest int `config:"12,The int to test"`
}

func TestBindIntSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindIntSetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetInt("inttest")

	assert.Equal(t, 12, i)
	assert.Nil(t, err)
}

type BindIntInvalidDefaultReturnsError struct {
	IntTest int `config:"FOO,The int to test"`
}

func TestBindIntInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindIntInvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// String.
type BindStringSetsDefault struct {
	StringTest string `config:"test string,The String to test"`
}

func TestBindStringSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindStringSetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetString("stringtest")

	assert.Equal(t, "test string", i)
	assert.Nil(t, err)
}

// Float64.
type BindFloat64SetsDefault struct {
	Float64Test float64 `config:"24,The Float64 to test"`
}

func TestBindFloat64SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat64SetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetFloat64("float64test")

	assert.Equal(t, float64(24), i)
	assert.Nil(t, err)
}

type BindFloat64InvalidDefaultReturnsError struct {
	Float64Test float64 `config:"FOO,The Float64 to test"`
}

func TestBindFloat64InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat64InvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Float32.
type BindFloat32SetsDefault struct {
	Float32Test float32 `config:"36,The Float32 to test"`
}

func TestBindFloat32SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat32SetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetFloat32("float32test")

	assert.Equal(t, float32(36), i)
	assert.Nil(t, err)
}

type BindFloat32InvalidDefaultReturnsError struct {
	Float32Test float32 `config:"FOO,The Float32 to test"`
}

func TestBindFloat32InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat32InvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Int8.
type BindInt8SetsDefault struct {
	Int8Test int8 `config:"14,The int8 to test"`
}

func TestBindInt8SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt8SetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetInt8("int8test")

	assert.Equal(t, int8(14), i)
	assert.Nil(t, err)
}

type BindInt8InvalidDefaultReturnsError struct {
	Int8Test int8 `config:"FOO,The Int8 to test"`
}

func TestBindInt8InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt8InvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Int16.
type BindInt16SetsDefault struct {
	Int16Test int16 `config:"16,The int16 to test"`
}

func TestBindInt16SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt16SetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetInt16("int16test")

	assert.Equal(t, int16(16), i)
	assert.Nil(t, err)
}

type BindInt16InvalidDefaultReturnsError struct {
	Int16Test int16 `config:"FOO,The Int16 to test"`
}

func TestBindInt16InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt16InvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Int32.
type BindInt32SetsDefault struct {
	Int32Test int32 `config:"32,The int32 to test"`
}

func TestBindInt32SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt32SetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetInt32("int32test")

	assert.Equal(t, int32(32), i)
	assert.Nil(t, err)
}

type BindInt32InvalidDefaultReturnsError struct {
	Int32Test int32 `config:"FOO,The Int32 to test"`
}

func TestBindInt32InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt32InvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Int64.
type BindInt64SetsDefault struct {
	Int64Test int64 `config:"64,The int64 to test"`
}

func TestBindInt64SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt64SetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetInt64("int64test")

	assert.Equal(t, int64(64), i)
	assert.Nil(t, err)
}

type BindInt64InvalidDefaultReturnsError struct {
	Int64Test int64 `config:"FOO,The Int64 to test"`
}

func TestBindInt64InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt64InvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Bool.
type BindBoolSetsDefault struct {
	BoolTest bool `config:"true,The Bool to test"`
}

func TestBindBoolSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindBoolSetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetBool("booltest")

	assert.Equal(t, true, i)
	assert.Nil(t, err)
}

type BindBoolInvalidDefaultReturnsError struct {
	BoolTest bool `config:"FOO,The Bool to test"`
}

func TestBindBoolInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindBoolInvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// IntSlice.
type BindIntSliceSetsDefault struct {
	IntSliceTest []int `config:"[12,13,14,15],The IntSlice to test"`
}

func TestBindIntSliceSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindIntSliceSetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetIntSlice("intslicetest")

	assert.Equal(t, []int{12, 13, 14, 15}, i)
	assert.Nil(t, err)
}

type BindIntSliceInvalidDefaultReturnsError struct {
	IntSliceTest []int `config:"FOO,The IntSlice to test"`
}

func TestBindIntSliceInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindIntSliceInvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// StringSlice.
type BindStringSliceSetsDefault struct {
	StringSliceTest []string `config:"[\"Value1\",\"Value2\"],The StringSlice to test"`
}

func TestBindStringSliceSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindStringSliceSetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetStringSlice("stringslicetest")

	assert.Equal(t, []string{"Value1", "Value2"}, i)
	assert.Nil(t, err)
}

type BindStringSliceInvalidDefaultReturnsError struct {
	StringSliceTest []string `config:"FOO,The StringSlice to test"`
}

func TestBindStringSliceInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindStringSliceInvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Float64Slice.
type BindFloat64SliceSetsDefault struct {
	Float64SliceTest []float64 `config:"[12,13,14,15],The Float64Slice to test"`
}

func TestBindFloat64SliceSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat64SliceSetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetFloat64Slice("float64slicetest")

	assert.Equal(t, []float64{12, 13, 14, 15}, i)
	assert.Nil(t, err)
}

type BindFloat64SliceInvalidDefaultReturnsError struct {
	Float64SliceTest []float64 `config:"FOO,The Float64Slice to test"`
}

func TestBindFloat64SliceInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat64SliceInvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Float32Slice.
type BindFloat32SliceSetsDefault struct {
	Float32SliceTest []float32 `config:"[12,13,14,15],The Float32Slice to test"`
}

func TestBindFloat32SliceSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat32SliceSetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetFloat32Slice("float32slicetest")

	assert.Equal(t, []float32{12, 13, 14, 15}, i)
	assert.Nil(t, err)
}

type BindFloat32SliceInvalidDefaultReturnsError struct {
	Float32SliceTest []float32 `config:"FOO,The Float32Slice to test"`
}

func TestBindFloat32SliceInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat32SliceInvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Int32Slice.
type BindInt32SliceSetsDefault struct {
	Int32SliceTest []int32 `config:"[12,13,14,15],The Int32Slice to test"`
}

func TestBindInt32SliceSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt32SliceSetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetInt32Slice("int32slicetest")

	assert.Equal(t, []int32{12, 13, 14, 15}, i)
	assert.Nil(t, err)
}

type BindInt32SliceInvalidDefaultReturnsError struct {
	Int32SliceTest []int32 `config:"FOO,The Int32Slice to test"`
}

func TestBindInt32SliceInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt32SliceInvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Int64Slice.
type BindInt64SliceSetsDefault struct {
	Int64SliceTest []int64 `config:"[12,13,14,15],The Int64Slice to test"`
}

func TestBindInt64SliceSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt64SliceSetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetInt64Slice("int64slicetest")

	assert.Equal(t, []int64{12, 13, 14, 15}, i)
	assert.Nil(t, err)
}

type BindInt64SliceInvalidDefaultReturnsError struct {
	Int64SliceTest []int64 `config:"FOO,The Int64Slice to test"`
}

func TestBindInt64SliceInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt64SliceInvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// BoolSlice.
type BindBoolSliceSetsDefault struct {
	BoolSliceTest []bool `config:"[true,true,false,true],The BoolSlice to test"`
}

func TestBindBoolSliceSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindBoolSliceSetsDefault{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetBoolSlice("boolslicetest")

	assert.Equal(t, []bool{true, true, false, true}, i)
	assert.Nil(t, err)
}

type BindBoolSliceInvalidDefaultReturnsError struct {
	BoolSliceTest []bool `config:"FOO,The BoolSlice to test"`
}

func TestBindBoolSliceInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindBoolSliceInvalidDefaultReturnsError{}, cmd)

	assert.NotNil(t, err)
}

// Tag parsing.
type InvalidTagReturnsError1 struct {
	TagTest string `config:"default"`
}

func TestInvalidTagReturnsError1(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(InvalidTagReturnsError1{}, cmd)

	assert.NotNil(t, err)
}

// NestedStruct.
type BindNestedStructSetsDefaults struct {
	BoolSlice   []bool                            `config:"[true,true,false,true],The NestedStruct to test"`
	InnerStruct BindNestedStructInnerSetsDefaults `config:","`
}

type BindNestedStructInnerSetsDefaults struct {
	String string `config:"String Test Yo!,The NestedStruct to test"`
}

func TestBindNestedStructSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindNestedStructSetsDefaults{}, cmd)
	assert.Nil(t, err)

	i, err := cmd.Flags().GetString("innerstruct.string")
	assert.Equal(t, "String Test Yo!", i)
	assert.Nil(t, err)

	j, err := cmd.Flags().GetBoolSlice("boolslice")
	assert.Equal(t, []bool{true, true, false, true}, j)
	assert.Nil(t, err)
}

// Ignore unexpected

//nolint:structcheck
type BindIgnoresUnexportedFields struct {
	BoolSlice []bool `config:"[true,true,false,true],A slice full of bools"`
	//nolint:unused
	unexported string `config:",Unexported"`
}

func TestBindIgnoresUnexportedFields(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindIgnoresUnexportedFields{}, cmd)
	assert.Nil(t, err)

	j, err := cmd.Flags().GetBoolSlice("boolslice")
	assert.Equal(t, []bool{true, true, false, true}, j)
	assert.Nil(t, err)

	s, err := cmd.Flags().GetString("unexported")
	assert.Equal(t, "", s)
	assert.NotNil(t, err)
}

// Ignore maps

type BindIgnoresMaps struct {
	Map map[string]string `config:",The map"`
}

func TestBindIgnoresMaps(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindIgnoresMaps{}, cmd)
	assert.Nil(t, err)

	j, err := cmd.Flags().GetString("map")
	assert.Equal(t, "", j)
	assert.NotNil(t, err)
}
