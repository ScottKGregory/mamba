package config

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Parse(obj interface{}, cmd *cobra.Command) error {
	return processFields("", reflect.TypeOf(obj), cmd)
}

func processFields(prefix string, t reflect.Type, cmd *cobra.Command) error {
	for i := 0; i < t.NumField(); i++ {
		err := processField(prefix, t.Field(i), cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

func processField(prefix string, field reflect.StructField, cmd *cobra.Command) error {
	n := strings.ToLower(field.Name)
	if prefix != "" {
		n = fmt.Sprintf("%s.%s", prefix, strings.ToLower(field.Name))
	}

	l := log.With().Str("name", n).Logger()

	k := field.Type.Kind()
	tag := field.Tag.Get("config")
	s := strings.Split(tag, ",")
	if k == reflect.Array ||
		k == reflect.Slice ||
		k == reflect.Map ||
		k == reflect.Struct ||
		k == reflect.Ptr {
		s = []string{"", ""}
	}

	if len(s) != 2 {
		l.Error().Msg("Config tag requires two values in the format \"default,description\"")
		return errors.New("invalid config tag, both default and description are required")
	}

	l = l.With().Str("def", s[0]).Str("desc", s[1]).Logger()

	v := true
	switch k {
	case reflect.Int:
		var i int
		i, err := strconv.Atoi(s[0])
		if err != nil {
			l.Error().Err(err).Msg("Invalid default value")
			return err
		}

		cmd.Flags().Int(n, i, s[1])
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.String:
		cmd.Flags().String(n, s[0], s[1])
		err := viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Float64:
		var i float64
		i, err := strconv.ParseFloat(s[0], 64)
		if err != nil {
			l.Error().Err(err).Msg("invalid default value")
			return err
		}

		cmd.Flags().Float64(n, i, s[1])
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Float32:
		var i float64
		i, err := strconv.ParseFloat(s[0], 32)
		if err != nil {
			l.Error().Err(err).Msg("invalid default value")
			return err
		}

		cmd.Flags().Float32(n, float32(i), s[1])
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Int8:
		var i int64
		i, err := strconv.ParseInt(s[0], 0, 8)
		if err != nil {
			l.Error().Err(err).Msg("invalid default value")
			return err
		}

		cmd.Flags().Int8(n, int8(i), s[1])
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Int16:
		var i int64
		i, err := strconv.ParseInt(s[0], 0, 16)
		if err != nil {
			l.Error().Err(err).Msg("invalid default value")
			return err
		}

		cmd.Flags().Int16(n, int16(i), s[1])
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Int32:
		var i int64
		i, err := strconv.ParseInt(s[0], 0, 32)
		if err != nil {
			l.Error().Err(err).Msg("invalid default value")
			return err
		}

		cmd.Flags().Int32(n, int32(i), s[1])
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Int64:
		var i int64
		i, err := strconv.ParseInt(s[0], 0, 64)
		if err != nil {
			l.Error().Err(err).Msg("invalid default value")
			return err
		}

		cmd.Flags().Int64(n, int64(i), s[1])
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Bool:
		var i bool
		i, err := strconv.ParseBool(s[0])
		if err != nil {
			l.Error().Err(err).Msg("invalid default value")
			return err
		}

		cmd.Flags().Bool(n, i, s[1])
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Array:
		err := processArray(n, s, field, cmd)
		if err != nil && strings.HasPrefix(err.Error(), "unsupported type for array") {
			l.Warn().Msg("Usupported array type detected, no flag will be bound")
			err = nil
			v = false
		}
	case reflect.Slice:
		err := processSlice(n, s, field, cmd)
		if err != nil && strings.HasPrefix(err.Error(), "unsupported type for slice") {
			l.Warn().Msg("Usupported slice type detected, no flag will be bound")
			err = nil
			v = false
		}
	case reflect.Map:
		l.Warn().Msg("Map value detected, no flag will be bound")
		v = false
	case reflect.Struct:
		return processFields(n, field.Type, cmd)
	case reflect.Ptr:
		return processFields(n, field.Type.Elem(), cmd)
	default:
		return errors.New("inavlid type supplied")
	}

	if v {
		l.Trace().Msg("Field registered")
	}

	return nil
}

func processSlice(n string, tag []string, field reflect.StructField, cmd *cobra.Command) (err error) {
	switch field.Type.Elem().Kind() {
	case reflect.Int:
		cmd.Flags().IntSlice(n, []int{}, tag[1])
	case reflect.String:
		cmd.Flags().StringSlice(n, []string{}, tag[1])
	case reflect.Float64:
		cmd.Flags().Float64Slice(n, []float64{}, tag[1])
	case reflect.Float32:
		cmd.Flags().Float32Slice(n, []float32{}, tag[1])
	case reflect.Int32:
		cmd.Flags().Int32Slice(n, []int32{}, tag[1])
	case reflect.Int64:
		cmd.Flags().Int64Slice(n, []int64{}, tag[1])
	case reflect.Bool:
		cmd.Flags().BoolSlice(n, []bool{}, tag[1])
	case reflect.Int8,
		reflect.Int16,
		reflect.Array,
		reflect.Slice,
		reflect.Map,
		reflect.Struct,
		reflect.Ptr:
		return fmt.Errorf("unsupported type for slice: %s", n)
	default:
		return err
	}

	err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
	if err != nil {
		return err
	}
	return err
}

func processArray(n string, tag []string, field reflect.StructField, cmd *cobra.Command) (err error) {
	switch field.Type.Elem().Kind() {
	case reflect.String:
		cmd.Flags().StringArray(n, []string{}, tag[1])
	case reflect.Int,
		reflect.Float64,
		reflect.Float32,
		reflect.Int32,
		reflect.Int64,
		reflect.Bool,
		reflect.Int8,
		reflect.Int16,
		reflect.Array,
		reflect.Slice,
		reflect.Map,
		reflect.Struct,
		reflect.Ptr:
		return fmt.Errorf("unsupported type for slice: %s", n)
	default:
		return err
	}

	err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
	if err != nil {
		return err
	}
	return err
}
