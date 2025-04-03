package internal

import (
	"errors"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

// Int.
type BindIntSetsDefault struct {
	IntTest int `config:"12,The int to test"`
}

func TestBindIntSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindIntSetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetInt("inttest")

	assertEqual(t, 12, i)
	assertNil(t, err)
}

type BindIntInvalidDefaultReturnsError struct {
	IntTest int `config:"FOO,The int to test"`
}

func TestBindIntInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindIntInvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// String.
type BindStringSetsDefault struct {
	StringTest string `config:"test string,The String to test"`
}

func TestBindStringSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindStringSetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetString("stringtest")

	assertEqual(t, "test string", i)
	assertNil(t, err)
}

// Float64.
type BindFloat64SetsDefault struct {
	Float64Test float64 `config:"24,The Float64 to test"`
}

func TestBindFloat64SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat64SetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetFloat64("float64test")

	assertEqual(t, float64(24), i)
	assertNil(t, err)
}

type BindFloat64InvalidDefaultReturnsError struct {
	Float64Test float64 `config:"FOO,The Float64 to test"`
}

func TestBindFloat64InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat64InvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// Float32.
type BindFloat32SetsDefault struct {
	Float32Test float32 `config:"36,The Float32 to test"`
}

func TestBindFloat32SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat32SetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetFloat32("float32test")

	assertEqual(t, float32(36), i)
	assertNil(t, err)
}

type BindFloat32InvalidDefaultReturnsError struct {
	Float32Test float32 `config:"FOO,The Float32 to test"`
}

func TestBindFloat32InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat32InvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// Int8.
type BindInt8SetsDefault struct {
	Int8Test int8 `config:"14,The int8 to test"`
}

func TestBindInt8SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt8SetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetInt8("int8test")

	assertEqual(t, int8(14), i)
	assertNil(t, err)
}

type BindInt8InvalidDefaultReturnsError struct {
	Int8Test int8 `config:"FOO,The Int8 to test"`
}

func TestBindInt8InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt8InvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// Int16.
type BindInt16SetsDefault struct {
	Int16Test int16 `config:"16,The int16 to test"`
}

func TestBindInt16SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt16SetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetInt16("int16test")

	assertEqual(t, int16(16), i)
	assertNil(t, err)
}

type BindInt16InvalidDefaultReturnsError struct {
	Int16Test int16 `config:"FOO,The Int16 to test"`
}

func TestBindInt16InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt16InvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// Int32.
type BindInt32SetsDefault struct {
	Int32Test int32 `config:"32,The int32 to test"`
}

func TestBindInt32SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt32SetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetInt32("int32test")

	assertEqual(t, int32(32), i)
	assertNil(t, err)
}

type BindInt32InvalidDefaultReturnsError struct {
	Int32Test int32 `config:"FOO,The Int32 to test"`
}

func TestBindInt32InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt32InvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// Int64.
type BindInt64SetsDefault struct {
	Int64Test int64 `config:"64,The int64 to test"`
}

func TestBindInt64SetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt64SetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetInt64("int64test")

	assertEqual(t, int64(64), i)
	assertNil(t, err)
}

type BindInt64InvalidDefaultReturnsError struct {
	Int64Test int64 `config:"FOO,The Int64 to test"`
}

func TestBindInt64InvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt64InvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// Bool.
type BindBoolSetsDefault struct {
	BoolTest bool `config:"true,The Bool to test"`
}

func TestBindBoolSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindBoolSetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetBool("booltest")

	assertEqual(t, true, i)
	assertNil(t, err)
}

type BindBoolInvalidDefaultReturnsError struct {
	BoolTest bool `config:"FOO,The Bool to test"`
}

func TestBindBoolInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindBoolInvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// IntSlice.
type BindIntSliceSetsDefault struct {
	IntSliceTest []int `config:"\"[12,13,14,15]\",The IntSlice to test"`
}

func TestBindIntSliceSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindIntSliceSetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetIntSlice("intslicetest")

	assertSliceEqual(t, []int{12, 13, 14, 15}, i)
	assertNil(t, err)
}

type BindIntSliceInvalidDefaultReturnsError struct {
	IntSliceTest []int `config:"FOO,The IntSlice to test"`
}

func TestBindIntSliceInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindIntSliceInvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// StringSlice.
type BindStringSliceSetsDefault struct {
	StringSliceTest []string `config:"\"[\"\"Value1\"\",\"\"Value2\"\"]\",The StringSlice to test"`
}

func TestBindStringSliceSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindStringSliceSetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetStringSlice("stringslicetest")

	assertSliceEqual(t, []string{"Value1", "Value2"}, i)
	assertNil(t, err)
}

type BindStringSliceInvalidDefaultReturnsError struct {
	StringSliceTest []string `config:"FOO,The StringSlice to test"`
}

func TestBindStringSliceInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindStringSliceInvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// Float64Slice.
type BindFloat64SliceSetsDefault struct {
	Float64SliceTest []float64 `config:"\"[12,13,14,15]\",The Float64Slice to test"`
}

func TestBindFloat64SliceSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat64SliceSetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetFloat64Slice("float64slicetest")

	assertSliceEqual(t, []float64{12, 13, 14, 15}, i)
	assertNil(t, err)
}

type BindFloat64SliceInvalidDefaultReturnsError struct {
	Float64SliceTest []float64 `config:"FOO,The Float64Slice to test"`
}

func TestBindFloat64SliceInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat64SliceInvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// Float32Slice.
type BindFloat32SliceSetsDefault struct {
	Float32SliceTest []float32 `config:"\"[12,13,14,15]\",The Float32Slice to test"`
}

func TestBindFloat32SliceSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat32SliceSetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetFloat32Slice("float32slicetest")

	assertSliceEqual(t, []float32{12, 13, 14, 15}, i)
	assertNil(t, err)
}

type BindFloat32SliceInvalidDefaultReturnsError struct {
	Float32SliceTest []float32 `config:"FOO,The Float32Slice to test"`
}

func TestBindFloat32SliceInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindFloat32SliceInvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// Int32Slice.
type BindInt32SliceSetsDefault struct {
	Int32SliceTest []int32 `config:"\"[12,13,14,15]\",The Int32Slice to test"`
}

func TestBindInt32SliceSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt32SliceSetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetInt32Slice("int32slicetest")

	assertSliceEqual(t, []int32{12, 13, 14, 15}, i)
	assertNil(t, err)
}

type BindInt32SliceInvalidDefaultReturnsError struct {
	Int32SliceTest []int32 `config:"FOO,The Int32Slice to test"`
}

func TestBindInt32SliceInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt32SliceInvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// Int64Slice.
type BindInt64SliceSetsDefault struct {
	Int64SliceTest []int64 `config:"\"[12,13,14,15]\",The Int64Slice to test"`
}

func TestBindInt64SliceSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt64SliceSetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetInt64Slice("int64slicetest")

	assertSliceEqual(t, []int64{12, 13, 14, 15}, i)
	assertNil(t, err)
}

type BindInt64SliceInvalidDefaultReturnsError struct {
	Int64SliceTest []int64 `config:"FOO,The Int64Slice to test"`
}

func TestBindInt64SliceInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInt64SliceInvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// BoolSlice.
type BindBoolSliceSetsDefault struct {
	BoolSliceTest []bool `config:"\"[true,true,false,true]\",The BoolSlice to test"`
}

func TestBindBoolSliceSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindBoolSliceSetsDefault{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetBoolSlice("boolslicetest")

	assertSliceEqual(t, []bool{true, true, false, true}, i)
	assertNil(t, err)
}

type BindBoolSliceInvalidDefaultReturnsError struct {
	BoolSliceTest []bool `config:"FOO,The BoolSlice to test"`
}

func TestBindBoolSliceInvalidDefaultReturnsError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindBoolSliceInvalidDefaultReturnsError{}, cmd)

	assertError(t, err)
}

// NestedStruct.
type BindNestedStructSetsDefaults struct {
	BoolSlice   []bool                            `config:"\"[true,true,false,true]\",The NestedStruct to test"`
	InnerStruct BindNestedStructInnerSetsDefaults `config:""`
}

type BindNestedStructInnerSetsDefaults struct {
	String string `config:"String Test Yo!,The NestedStruct to test"`
}

func TestBindNestedStructSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindNestedStructSetsDefaults{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetString("innerstruct.string")
	assertEqual(t, "String Test Yo!", i)
	assertNil(t, err)

	j, err := cmd.Flags().GetBoolSlice("boolslice")
	assertSliceEqual(t, []bool{true, true, false, true}, j)
	assertNil(t, err)
}

// Ignore unexpected

//nolint:structcheck
type BindIgnoresUnexportedFields struct {
	BoolSlice []bool `config:"\"[true,true,false,true]\",A slice full of bools"`
	//nolint:unused
	unexported string `config:",Unexported"`
}

func TestBindIgnoresUnexportedFields(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindIgnoresUnexportedFields{}, cmd)
	assertNil(t, err)

	j, err := cmd.Flags().GetBoolSlice("boolslice")
	assertSliceEqual(t, []bool{true, true, false, true}, j)
	assertNil(t, err)

	s, err := cmd.Flags().GetString("unexported")
	assertEqual(t, "", s)
	assertError(t, err)
}

// Error on invalid type

type BindInvalidTypeError struct {
	Map map[string]string `config:",The map"`
}

func TestBindInvalidTypeError(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindInvalidTypeError{}, cmd)
	assertError(t, err)
}

// EmbeddedStruct.
type BindEmbeddedStructSetsDefaults struct {
	*BindEmbeddedStructInnerSetsDefaults `config:""`
}

type BindEmbeddedStructInnerSetsDefaults struct {
	String string `config:"String Test Yo!,The EmbeddedStruct to test"`
}

func TestBindEmbeddedStructSetsDefault(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindEmbeddedStructSetsDefaults{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetString("string")
	assertEqual(t, "String Test Yo!", i)
	assertNil(t, err)
}

// Optional separator.
type BindNestedStructUsesOptions struct {
	BoolSlice   []bool                           `config:"\"[true,true,false,true]\",The NestedStruct to test"`
	InnerStruct BindNestedStructInnerUsesOptions `config:""`
}

type BindNestedStructInnerUsesOptions struct {
	String string `config:"String Test Yo!,The NestedStruct to test"`
}

func TestBindNestedStructUsesOptions(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindNestedStructUsesOptions{}, cmd, &Options{Separator: ">>"})
	assertNil(t, err)

	i, err := cmd.Flags().GetString("innerstruct>>string")
	assertEqual(t, "String Test Yo!", i)
	assertNil(t, err)

	j, err := cmd.Flags().GetBoolSlice("boolslice")
	assertSliceEqual(t, []bool{true, true, false, true}, j)
	assertNil(t, err)
}

func TestBindNestedStructDefaultsToFullStop(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindNestedStructUsesOptions{}, cmd, &Options{Separator: ""})
	assertNil(t, err)

	i, err := cmd.Flags().GetString("innerstruct.string")
	assertEqual(t, "String Test Yo!", i)
	assertNil(t, err)

	j, err := cmd.Flags().GetBoolSlice("boolslice")
	assertSliceEqual(t, []bool{true, true, false, true}, j)
	assertNil(t, err)
}

func TestBindDereferencesPointer(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(&BindNestedStructUsesOptions{}, cmd)
	assertNil(t, err)

	i, err := cmd.Flags().GetString("innerstruct.string")
	assertEqual(t, "String Test Yo!", i)
	assertNil(t, err)

	j, err := cmd.Flags().GetBoolSlice("boolslice")
	assertSliceEqual(t, []bool{true, true, false, true}, j)
	assertNil(t, err)
}

// Tags.
type BindSkipsUntagged struct {
	BoolSlice  []bool `config:"\"[true,true,false,true]\",The NestedStruct to test"`
	TestString string
}

func TestBindSkipsUntagged(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindSkipsUntagged{}, cmd)
	assertNil(t, err)

	j, err := cmd.Flags().GetBoolSlice("boolslice")
	assertSliceEqual(t, []bool{true, true, false, true}, j)
	assertNil(t, err)

	s, err := cmd.Flags().GetString("teststring")
	assertEqual(t, "", s)
	assertError(t, err)
}

type BindErrorsOnInvalidTag struct {
	TestString string `config:"defualt,asdfg,asg,safg,asfg,afg"`
}

func TestBindErrorsOnInvalidTag(t *testing.T) {
	cmd := &cobra.Command{}
	err := Bind(BindErrorsOnInvalidTag{}, cmd)
	assertErrorIs(t, err, &tagParseError{})
}

func assertNil(t *testing.T, val any) {
	if val != nil {
		fmt.Printf("expected nil but got %v\n", val)
		t.Fail()
	}
}

func assertEqual[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		fmt.Printf("expected '%v' but got '%v'\n", expected, actual)
		t.Fail()
	}
}

func assertSliceEqual[T comparable](t *testing.T, expected, actual []T) {
	for i, e := range expected {
		assertEqual(t, e, actual[i])
	}
}

func assertError(t *testing.T, err error) {
	if err == nil {
		fmt.Printf("expected error but got '%v'\n", err)
		t.Fail()
	}
}

func assertErrorIs(t *testing.T, actual, expected error) {
	if !errors.Is(expected, actual) {
		fmt.Printf("expected error to be '%v' but got '%v'\n", expected, actual)
		t.Fail()
	}
}
