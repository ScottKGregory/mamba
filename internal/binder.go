package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Binder struct {
	opts *Options
}

// Bind binds the config tags from the structs and binds flags to the cobra command.
func Bind(obj any, cmd *cobra.Command, options ...*Options) error {
	b := &Binder{
		opts: &Options{
			Separator:      ".",
			Persistent:     false,
			PrefixEmbedded: true,
		},
	}
	if len(options) == 1 {
		b.opts = options[0]
		if b.opts.Separator == "" {
			b.opts.Separator = "."
		}
	}

	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return b.processFields("", t, cmd)
}

func (b *Binder) processFields(prefix string, t reflect.Type, cmd *cobra.Command) error {
	for i := 0; i < t.NumField(); i++ {
		err := b.processField(prefix, t.Field(i), cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Binder) processField(prefix string, field reflect.StructField, cmd *cobra.Command) error {
	n := strings.ToLower(field.Name)
	if prefix != "" {
		n = fmt.Sprintf("%s%s%s", prefix, b.opts.Separator, strings.ToLower(field.Name))
	}

	if string(field.Name[0]) != strings.ToUpper(string(field.Name[0])) {
		return nil
	}

	k := field.Type.Kind()
	tag, present := field.Tag.Lookup("config")

	if !present {
		return nil
	}

	t, err := Parse(tag)
	if err != nil {
		return NewTagParseError(tag, k, n, err)
	}

	f := b.flags(cmd, t)
	switch k {
	case reflect.Int:
		var i int
		if t.Default != "" {
			i, err = strconv.Atoi(t.Default)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}

		if t.Shorthand != "" {
			f.IntP(n, t.Shorthand, i, t.Description)
		} else {
			f.Int(n, i, t.Description)
		}
	case reflect.String:
		if t.Shorthand != "" {
			f.StringP(n, t.Shorthand, t.Default, t.Description)
		} else {
			f.String(n, t.Default, t.Description)
		}
	case reflect.Float64:
		var i float64
		if t.Default != "" {
			i, err = strconv.ParseFloat(t.Default, 64)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}

		if t.Shorthand != "" {
			f.Float64P(n, t.Shorthand, i, t.Description)
		} else {
			f.Float64(n, i, t.Description)
		}
	case reflect.Float32:
		var i float64
		if t.Default != "" {
			i, err = strconv.ParseFloat(t.Default, 32)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}

		if t.Shorthand != "" {
			f.Float32P(n, t.Shorthand, float32(i), t.Description)
		} else {
			f.Float32(n, float32(i), t.Description)
		}
	case reflect.Int8:
		var i int64
		if t.Default != "" {
			i, err = strconv.ParseInt(t.Default, 0, 8)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}

		if t.Shorthand != "" {
			f.Int8P(n, t.Shorthand, int8(i), t.Description)
		} else {
			f.Int8(n, int8(i), t.Description)
		}
	case reflect.Int16:
		var i int64
		if t.Default != "" {
			i, err = strconv.ParseInt(t.Default, 0, 16)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}

		if t.Shorthand != "" {
			f.Int16P(n, t.Shorthand, int16(i), t.Description)
		} else {
			f.Int16(n, int16(i), t.Description)
		}
	case reflect.Int32:
		var i int64
		if t.Default != "" {
			i, err = strconv.ParseInt(t.Default, 0, 32)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}

		if t.Shorthand != "" {
			f.Int32P(n, t.Shorthand, int32(i), t.Description)
		} else {
			f.Int32(n, int32(i), t.Description)
		}
	case reflect.Int64:
		var i int64
		if t.Default != "" {
			i, err = strconv.ParseInt(t.Default, 0, 64)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}

		if t.Shorthand != "" {
			f.Int64P(n, t.Shorthand, int64(i), t.Description)
		} else {
			f.Int64(n, int64(i), t.Description)
		}
	case reflect.Bool:
		var i bool
		if t.Default != "" {
			i, err = strconv.ParseBool(t.Default)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}

		if t.Shorthand != "" {
			f.BoolP(n, t.Shorthand, i, t.Description)
		} else {
			f.Bool(n, i, t.Description)
		}
	case reflect.Array, reflect.Slice:
		err := b.processSlice(n, t, field, cmd)
		if err != nil && errors.Is(err, InvalidTypeError) {
			return nil
		} else if err != nil {
			return NewParseError(t.Default, k, n, err)
		}
	case reflect.Struct:
		return b.processFields(n, field.Type, cmd)
	case reflect.Ptr:
		if field.Type.Elem().Kind() == reflect.Struct && b.opts.PrefixEmbedded {
			n = prefix
		}
		return b.processFields(n, field.Type.Elem(), cmd)
	default:
		return NewInvalidTypeError(field.Type.Kind(), n)
	}

	err = viper.BindPFlag(n, f.Lookup(n))
	if err != nil {
		return NewBindError(field.Type.Kind(), n, err)
	}

	return nil
}

type jsonStruct struct {
	StringArray  []string  `json:"stringarray,omitempty"`
	IntArray     []int     `json:"intarray,omitempty"`
	Float64Array []float64 `json:"float64array,omitempty"`
	Float32Array []float32 `json:"float32array,omitempty"`
	Int32Array   []int32   `json:"int32array,omitempty"`
	Int64Array   []int64   `json:"int64array,omitempty"`
	BoolArray    []bool    `json:"boolarray,omitempty"`
}

func (b *Binder) processSlice(n string, t *Tag, field reflect.StructField, cmd *cobra.Command) (err error) {
	s := &jsonStruct{}
	k := field.Type.Kind()
	f := b.flags(cmd, t)
	switch field.Type.Elem().Kind() {
	case reflect.Int:
		if t.Default != "" {
			def := fmt.Sprintf("{\"intarray\":%s}", t.Default)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}

		if t.Shorthand != "" {
			f.IntSliceP(n, t.Shorthand, s.IntArray, t.Description)
		} else {
			f.IntSlice(n, s.IntArray, t.Description)
		}
	case reflect.String:
		if t.Default != "" {
			def := fmt.Sprintf("{\"stringarray\":%s}", t.Default)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}

		if t.Shorthand != "" {
			f.StringSliceP(n, t.Shorthand, s.StringArray, t.Description)
		} else {
			f.StringSlice(n, s.StringArray, t.Description)
		}
	case reflect.Float64:
		if t.Default != "" {
			def := fmt.Sprintf("{\"float64array\":%s}", t.Default)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}

		if t.Shorthand != "" {
			f.Float64SliceP(n, t.Shorthand, s.Float64Array, t.Description)
		} else {
			f.Float64Slice(n, s.Float64Array, t.Description)
		}
	case reflect.Float32:
		if t.Default != "" {
			def := fmt.Sprintf("{\"float32array\":%s}", t.Default)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}
		if t.Shorthand != "" {
			f.Float32SliceP(n, t.Shorthand, s.Float32Array, t.Description)
		} else {
			f.Float32Slice(n, s.Float32Array, t.Description)
		}
	case reflect.Int32:
		if t.Default != "" {
			def := fmt.Sprintf("{\"int32array\":%s}", t.Default)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}
		if t.Shorthand != "" {
			f.Int32SliceP(n, t.Shorthand, s.Int32Array, t.Description)
		} else {
			f.Int32Slice(n, s.Int32Array, t.Description)
		}
	case reflect.Int64:
		if t.Default != "" {
			def := fmt.Sprintf("{\"int64array\":%s}", t.Default)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}
		if t.Shorthand != "" {
			f.Int64SliceP(n, t.Shorthand, s.Int64Array, t.Description)
		} else {
			f.Int64Slice(n, s.Int64Array, t.Description)
		}
	case reflect.Bool:
		if t.Default != "" {
			def := fmt.Sprintf("{\"boolarray\":%s}", t.Default)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return NewParseError(t.Default, k, n, err)
			}
		}
		if t.Shorthand != "" {
			f.BoolSliceP(n, t.Shorthand, s.BoolArray, t.Description)
		} else {
			f.BoolSlice(n, s.BoolArray, t.Description)
		}
	default:
		return NewInvalidTypeError(k, n)
	}

	err = viper.BindPFlag(n, f.Lookup(n))
	if err != nil {
		return NewBindError(k, n, err)
	}

	return err
}

func (b *Binder) flags(cmd *cobra.Command, t *Tag) *pflag.FlagSet {
	f := cmd.Flags()
	if t.Persistent || b.opts.Persistent {
		f = cmd.PersistentFlags()
	}
	return f
}
