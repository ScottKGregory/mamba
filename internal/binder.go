package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Binder struct {
	opts *Options
}

// Bind binds the config tags from the structs and binds flags to the cobra command.
func Bind(obj interface{}, cmd *cobra.Command, options ...*Options) error {
	b := &Binder{}
	if len(options) == 1 {
		b.opts = options[0]
	} else {
		b.opts = &Options{}
	}

	return b.processFields("", reflect.TypeOf(obj), cmd)
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
		n = fmt.Sprintf("%s.%s", prefix, strings.ToLower(field.Name))
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
		return err
	}

	switch k {
	case reflect.Int:
		var i int
		if t.Default != "" {
			i, err = strconv.Atoi(t.Default)
			if err != nil {
				return err
			}
		}
		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().IntP(n, t.Shorthand, i, t.Description)
			} else {
				cmd.PersistentFlags().Int(n, i, t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().IntP(n, t.Shorthand, i, t.Description)
			} else {
				cmd.Flags().Int(n, i, t.Description)
			}
		}
	case reflect.String:
		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().StringP(n, t.Shorthand, t.Default, t.Description)
			} else {
				cmd.PersistentFlags().String(n, t.Default, t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().StringP(n, t.Shorthand, t.Default, t.Description)
			} else {
				cmd.Flags().String(n, t.Default, t.Description)
			}
		}
	case reflect.Float64:
		var i float64
		if t.Default != "" {
			i, err = strconv.ParseFloat(t.Default, 64)
			if err != nil {
				return err
			}
		}

		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().Float64P(n, t.Shorthand, i, t.Description)
			} else {
				cmd.PersistentFlags().Float64(n, i, t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().Float64P(n, t.Shorthand, i, t.Description)
			} else {
				cmd.Flags().Float64(n, i, t.Description)
			}
		}
	case reflect.Float32:
		var i float64
		if t.Default != "" {
			i, err = strconv.ParseFloat(t.Default, 32)
			if err != nil {
				return err
			}
		}

		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().Float32P(n, t.Shorthand, float32(i), t.Description)
			} else {
				cmd.PersistentFlags().Float32(n, float32(i), t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().Float32P(n, t.Shorthand, float32(i), t.Description)
			} else {
				cmd.Flags().Float32(n, float32(i), t.Description)
			}
		}
	case reflect.Int8:
		var i int64
		if t.Default != "" {
			i, err = strconv.ParseInt(t.Default, 0, 8)
			if err != nil {
				return err
			}
		}

		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().Int8P(n, t.Shorthand, int8(i), t.Description)
			} else {
				cmd.PersistentFlags().Int8(n, int8(i), t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().Int8P(n, t.Shorthand, int8(i), t.Description)
			} else {
				cmd.Flags().Int8(n, int8(i), t.Description)
			}
		}
	case reflect.Int16:
		var i int64
		if t.Default != "" {
			i, err = strconv.ParseInt(t.Default, 0, 16)
			if err != nil {
				return err
			}
		}

		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().Int16P(n, t.Shorthand, int16(i), t.Description)
			} else {
				cmd.PersistentFlags().Int16(n, int16(i), t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().Int16P(n, t.Shorthand, int16(i), t.Description)
			} else {
				cmd.Flags().Int16(n, int16(i), t.Description)
			}
		}
	case reflect.Int32:
		var i int64
		if t.Default != "" {
			i, err = strconv.ParseInt(t.Default, 0, 32)
			if err != nil {
				return err
			}
		}

		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().Int32P(n, t.Shorthand, int32(i), t.Description)
			} else {
				cmd.PersistentFlags().Int32(n, int32(i), t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().Int32P(n, t.Shorthand, int32(i), t.Description)
			} else {
				cmd.Flags().Int32(n, int32(i), t.Description)
			}
		}
	case reflect.Int64:
		var i int64
		if t.Default != "" {
			i, err = strconv.ParseInt(t.Default, 0, 64)
			if err != nil {
				return err
			}
		}

		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().Int64P(n, t.Shorthand, int64(i), t.Description)
			} else {
				cmd.PersistentFlags().Int64(n, int64(i), t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().Int64P(n, t.Shorthand, int64(i), t.Description)
			} else {
				cmd.Flags().Int64(n, int64(i), t.Description)
			}
		}
	case reflect.Bool:
		var i bool
		if t.Default != "" {
			i, err = strconv.ParseBool(t.Default)
			if err != nil {
				return err
			}
		}

		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().BoolP(n, t.Shorthand, i, t.Description)
			} else {
				cmd.PersistentFlags().Bool(n, i, t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().BoolP(n, t.Shorthand, i, t.Description)
			} else {
				cmd.Flags().Bool(n, i, t.Description)
			}
		}
	case reflect.Array, reflect.Slice:
		err := b.processSlice(n, t, field, cmd)
		if err != nil && strings.HasPrefix(err.Error(), "unsupported type for slice") {
			err = nil
		} else if err != nil {
			return err
		}
	case reflect.Map:
		return nil
	case reflect.Struct:
		return b.processFields(n, field.Type, cmd)
	case reflect.Ptr:
		return b.processFields(n, field.Type.Elem(), cmd)
	default:
		return errors.New("inavlid type supplied")
	}

	if t.Persistent || b.opts.Persistent {
		err = viper.BindPFlag(n, cmd.PersistentFlags().Lookup(n))
	} else {
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
	}
	if err != nil {
		return err
	}

	return nil
}

type JSONStruct struct {
	StringArray  []string  `json:"stringarray,omitempty"`
	IntArray     []int     `json:"intarray,omitempty"`
	Float64Array []float64 `json:"float64array,omitempty"`
	Float32Array []float32 `json:"float32array,omitempty"`
	Int32Array   []int32   `json:"int32array,omitempty"`
	Int64Array   []int64   `json:"int64array,omitempty"`
	BoolArray    []bool    `json:"boolarray,omitempty"`
}

func (b *Binder) processSlice(n string, t *Tag, field reflect.StructField, cmd *cobra.Command) (err error) {
	s := &JSONStruct{}
	switch field.Type.Elem().Kind() {
	case reflect.Int:
		if t.Default != "" {
			def := fmt.Sprintf("{\"intarray\":%s}", t.Default)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return err
			}
		}
		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().IntSliceP(n, t.Shorthand, s.IntArray, t.Description)
			} else {
				cmd.PersistentFlags().IntSlice(n, s.IntArray, t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().IntSliceP(n, t.Shorthand, s.IntArray, t.Description)
			} else {
				cmd.Flags().IntSlice(n, s.IntArray, t.Description)
			}
		}
	case reflect.String:
		if t.Default != "" {
			def := fmt.Sprintf("{\"stringarray\":%s}", t.Default)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return err
			}
		}
		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().StringSliceP(n, t.Shorthand, s.StringArray, t.Description)
			} else {
				cmd.PersistentFlags().StringSlice(n, s.StringArray, t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().StringSliceP(n, t.Shorthand, s.StringArray, t.Description)
			} else {
				cmd.Flags().StringSlice(n, s.StringArray, t.Description)
			}
		}
	case reflect.Float64:
		if t.Default != "" {
			def := fmt.Sprintf("{\"float64array\":%s}", t.Default)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return err
			}
		}
		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().Float64SliceP(n, t.Shorthand, s.Float64Array, t.Description)
			} else {
				cmd.PersistentFlags().Float64Slice(n, s.Float64Array, t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().Float64SliceP(n, t.Shorthand, s.Float64Array, t.Description)
			} else {
				cmd.Flags().Float64Slice(n, s.Float64Array, t.Description)
			}
		}
	case reflect.Float32:
		if t.Default != "" {
			def := fmt.Sprintf("{\"float32array\":%s}", t.Default)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return err
			}
		}
		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().Float32SliceP(n, t.Shorthand, s.Float32Array, t.Description)
			} else {
				cmd.PersistentFlags().Float32Slice(n, s.Float32Array, t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().Float32SliceP(n, t.Shorthand, s.Float32Array, t.Description)
			} else {
				cmd.Flags().Float32Slice(n, s.Float32Array, t.Description)
			}
		}
	case reflect.Int32:
		if t.Default != "" {
			def := fmt.Sprintf("{\"int32array\":%s}", t.Default)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return err
			}
		}
		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().Int32SliceP(n, t.Shorthand, s.Int32Array, t.Description)
			} else {
				cmd.PersistentFlags().Int32Slice(n, s.Int32Array, t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().Int32SliceP(n, t.Shorthand, s.Int32Array, t.Description)
			} else {
				cmd.Flags().Int32Slice(n, s.Int32Array, t.Description)
			}
		}
	case reflect.Int64:
		if t.Default != "" {
			def := fmt.Sprintf("{\"int64array\":%s}", t.Default)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return err
			}
		}
		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().Int64SliceP(n, t.Shorthand, s.Int64Array, t.Description)
			} else {
				cmd.PersistentFlags().Int64Slice(n, s.Int64Array, t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().Int64SliceP(n, t.Shorthand, s.Int64Array, t.Description)
			} else {
				cmd.Flags().Int64Slice(n, s.Int64Array, t.Description)
			}
		}
	case reflect.Bool:
		if t.Default != "" {
			def := fmt.Sprintf("{\"boolarray\":%s}", t.Default)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return err
			}
		}
		if t.Persistent || b.opts.Persistent {
			if t.Shorthand != "" {
				cmd.PersistentFlags().BoolSliceP(n, t.Shorthand, s.BoolArray, t.Description)
			} else {
				cmd.PersistentFlags().BoolSlice(n, s.BoolArray, t.Description)
			}
		} else {
			if t.Shorthand != "" {
				cmd.Flags().BoolSliceP(n, t.Shorthand, s.BoolArray, t.Description)
			} else {
				cmd.Flags().BoolSlice(n, s.BoolArray, t.Description)
			}
		}
	default:
		return fmt.Errorf("unsupported type for slice: %s", n)
	}

	if t.Persistent || b.opts.Persistent {
		err = viper.BindPFlag(n, cmd.PersistentFlags().Lookup(n))
	} else {
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
	}
	if err != nil {
		return fmt.Errorf("error binding flag: %w", err)
	}

	return err
}
