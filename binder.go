package mamba

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

func MustBind(obj interface{}, cmd *cobra.Command, options ...*Options) {
	if err := Bind(obj, cmd, options...); err != nil {
		panic(err)
	}
}

// Bind binds the config tags from the structs and binds flags to the cobra command.
func Bind(obj interface{}, cmd *cobra.Command, options ...*Options) error {
	b := &Binder{}
	if len(options) == 1 {
		b.opts = options[0]
	} else {
		// Set defaults
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
	tag := field.Tag.Get("config")

	if tag == "" {
		return nil
	}

	var def, desc string
	if k == reflect.Map ||
		k == reflect.Struct ||
		k == reflect.Ptr {
		def = ""
		desc = ""
	} else {
		if len(strings.Split(tag, ",")) < 2 {
			return errors.New("invalid config tag, both default and description are required")
		}

		lastComma := strings.LastIndex(tag, ",") // Probably needs a full tokeniser
		def = tag[:lastComma]
		desc = tag[lastComma+1:]
	}

	var err error
	switch k {
	case reflect.Int:
		var i int
		if def != "" {
			i, err = strconv.Atoi(def)
			if err != nil {
				return err
			}
		}
		if b.opts.Persistent {
			cmd.PersistentFlags().Int(n, i, desc)
		} else {
			cmd.Flags().Int(n, i, desc)
		}
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.String:
		if b.opts.Persistent {
			cmd.PersistentFlags().String(n, def, desc)
		} else {
			cmd.Flags().String(n, def, desc)
		}
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Float64:
		var i float64
		if def != "" {
			i, err = strconv.ParseFloat(def, 64)
			if err != nil {
				return err
			}
		}

		if b.opts.Persistent {
			cmd.PersistentFlags().Float64(n, i, desc)
		} else {
			cmd.Flags().Float64(n, i, desc)
		}
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Float32:
		var i float64
		if def != "" {
			i, err = strconv.ParseFloat(def, 32)
			if err != nil {
				return err
			}
		}

		if b.opts.Persistent {
			cmd.PersistentFlags().Float32(n, float32(i), desc)
		} else {
			cmd.Flags().Float32(n, float32(i), desc)
		}
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Int8:
		var i int64
		if def != "" {
			i, err = strconv.ParseInt(def, 0, 8)
			if err != nil {
				return err
			}
		}

		if b.opts.Persistent {
			cmd.PersistentFlags().Int8(n, int8(i), desc)
		} else {
			cmd.Flags().Int8(n, int8(i), desc)
		}
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Int16:
		var i int64
		if def != "" {
			i, err = strconv.ParseInt(def, 0, 16)
			if err != nil {
				return err
			}
		}

		if b.opts.Persistent {
			cmd.PersistentFlags().Int16(n, int16(i), desc)
		} else {
			cmd.Flags().Int16(n, int16(i), desc)
		}
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Int32:
		var i int64
		if def != "" {
			i, err = strconv.ParseInt(def, 0, 32)
			if err != nil {
				return err
			}
		}

		if b.opts.Persistent {
			cmd.PersistentFlags().Int32(n, int32(i), desc)
		} else {
			cmd.Flags().Int32(n, int32(i), desc)
		}
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Int64:
		var i int64
		if def != "" {
			i, err = strconv.ParseInt(def, 0, 64)
			if err != nil {
				return err
			}
		}

		if b.opts.Persistent {
			cmd.PersistentFlags().Int64(n, int64(i), desc)
		} else {
			cmd.Flags().Int64(n, int64(i), desc)
		}
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Bool:
		var i bool
		if def != "" {
			i, err = strconv.ParseBool(def)
			if err != nil {
				return err
			}
		}

		if b.opts.Persistent {
			cmd.PersistentFlags().Bool(n, i, desc)
		} else {
			cmd.Flags().Bool(n, i, desc)
		}
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Array, reflect.Slice:
		err := b.processSlice(n, def, desc, field, cmd)
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

func (b *Binder) processSlice(n, def, desc string, field reflect.StructField, cmd *cobra.Command) (err error) {
	s := &JSONStruct{}
	switch field.Type.Elem().Kind() {
	case reflect.Int:
		if def != "" {
			def := fmt.Sprintf("{\"intarray\":%s}", def)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return err
			}
		}
		if b.opts.Persistent {
			cmd.PersistentFlags().IntSlice(n, s.IntArray, desc)
		} else {
			cmd.Flags().IntSlice(n, s.IntArray, desc)
		}
	case reflect.String:
		if def != "" {
			def := fmt.Sprintf("{\"stringarray\":%s}", def)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return err
			}
		}
		if b.opts.Persistent {
			cmd.PersistentFlags().StringSlice(n, s.StringArray, desc)
		} else {
			cmd.Flags().StringSlice(n, s.StringArray, desc)
		}
	case reflect.Float64:
		if def != "" {
			def := fmt.Sprintf("{\"float64array\":%s}", def)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return err
			}
		}
		if b.opts.Persistent {
			cmd.PersistentFlags().Float64Slice(n, s.Float64Array, desc)
		} else {
			cmd.Flags().Float64Slice(n, s.Float64Array, desc)
		}
	case reflect.Float32:
		if def != "" {
			def := fmt.Sprintf("{\"float32array\":%s}", def)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return err
			}
		}
		if b.opts.Persistent {
			cmd.PersistentFlags().Float32Slice(n, s.Float32Array, desc)
		} else {
			cmd.Flags().Float32Slice(n, s.Float32Array, desc)
		}
	case reflect.Int32:
		if def != "" {
			def := fmt.Sprintf("{\"int32array\":%s}", def)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return err
			}
		}
		if b.opts.Persistent {
			cmd.PersistentFlags().Int32Slice(n, s.Int32Array, desc)
		} else {
			cmd.Flags().Int32Slice(n, s.Int32Array, desc)
		}
	case reflect.Int64:
		if def != "" {
			def := fmt.Sprintf("{\"int64array\":%s}", def)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return err
			}
		}
		if b.opts.Persistent {
			cmd.PersistentFlags().Int64Slice(n, s.Int64Array, desc)
		} else {
			cmd.Flags().Int64Slice(n, s.Int64Array, desc)
		}
	case reflect.Bool:
		if def != "" {
			def := fmt.Sprintf("{\"boolarray\":%s}", def)
			err := json.Unmarshal([]byte(def), &s)
			if err != nil {
				return err
			}
		}
		if b.opts.Persistent {
			cmd.PersistentFlags().BoolSlice(n, s.BoolArray, desc)
		} else {
			cmd.Flags().BoolSlice(n, s.BoolArray, desc)
		}
	default:
		return fmt.Errorf("unsupported type for slice: %s", n)
	}

	err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
	if err != nil {
		return fmt.Errorf("error binding flag: %w", err)
	}

	return err
}
