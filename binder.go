package mamba

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Binder struct {
	log zerolog.Logger
}

// Bind binds the config tags from the structs and binds flags to the cobra command.
func Bind(obj interface{}, cmd *cobra.Command, options ...*Options) error {
	p := &Binder{}

	if len(options) == 1 {
		opts := options[0]
		p.log = log.Level(opts.LogLevel)
	} else {
		p.log = log.Level(zerolog.ErrorLevel)
	}
	p.log = p.log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	return p.processFields("", reflect.TypeOf(obj), cmd)
}

func (p *Binder) processFields(prefix string, t reflect.Type, cmd *cobra.Command) error {
	for i := 0; i < t.NumField(); i++ {
		err := p.processField(prefix, t.Field(i), cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Binder) processField(prefix string, field reflect.StructField, cmd *cobra.Command) error {
	n := strings.ToLower(field.Name)
	if prefix != "" {
		n = fmt.Sprintf("%s.%s", prefix, strings.ToLower(field.Name))
	}

	l := p.log.With().Str("name", n).Logger()

	if string(field.Name[0]) != strings.ToUpper(string(field.Name[0])) {
		l.Debug().Msg("Skipping unexported field")
		return nil
	}

	k := field.Type.Kind()
	tag := field.Tag.Get("config")

	if tag == "" {
		l.Trace().Msg("Skipping untagged field")
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
			l.Error().Msg("Config tag requires two values in the format \"default,description\"")
			return errors.New("invalid config tag, both default and description are required")
		}

		lastComma := strings.LastIndex(tag, ",") // Probably needs a full tokeniser
		def = tag[:lastComma]
		desc = tag[lastComma+1:]
	}

	l = l.With().Str("def", def).Str("desc", desc).Logger()

	v := true
	switch k {
	case reflect.Int:
		var i int
		i, err := strconv.Atoi(def)
		if err != nil {
			l.Error().Err(err).Msg("Invalid default value")
			return err
		}

		cmd.Flags().Int(n, i, desc)
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.String:
		cmd.Flags().String(n, def, desc)
		err := viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Float64:
		var i float64
		i, err := strconv.ParseFloat(def, 64)
		if err != nil {
			l.Error().Err(err).Msg("invalid default value")
			return err
		}

		cmd.Flags().Float64(n, i, desc)
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Float32:
		var i float64
		i, err := strconv.ParseFloat(def, 32)
		if err != nil {
			l.Error().Err(err).Msg("invalid default value")
			return err
		}

		cmd.Flags().Float32(n, float32(i), desc)
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Int8:
		var i int64
		i, err := strconv.ParseInt(def, 0, 8)
		if err != nil {
			l.Error().Err(err).Msg("invalid default value")
			return err
		}

		cmd.Flags().Int8(n, int8(i), desc)
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Int16:
		var i int64
		i, err := strconv.ParseInt(def, 0, 16)
		if err != nil {
			l.Error().Err(err).Msg("invalid default value")
			return err
		}

		cmd.Flags().Int16(n, int16(i), desc)
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Int32:
		var i int64
		i, err := strconv.ParseInt(def, 0, 32)
		if err != nil {
			l.Error().Err(err).Msg("invalid default value")
			return err
		}

		cmd.Flags().Int32(n, int32(i), desc)
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Int64:
		var i int64
		i, err := strconv.ParseInt(def, 0, 64)
		if err != nil {
			l.Error().Err(err).Msg("invalid default value")
			return err
		}

		cmd.Flags().Int64(n, int64(i), desc)
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Bool:
		var i bool
		i, err := strconv.ParseBool(def)
		if err != nil {
			l.Error().Err(err).Msg("invalid default value")
			return err
		}

		cmd.Flags().Bool(n, i, desc)
		err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
		if err != nil {
			return err
		}
	case reflect.Array, reflect.Slice:
		err := p.processSlice(n, def, desc, field, cmd)
		if err != nil && strings.HasPrefix(err.Error(), "unsupported type for slice") {
			l.Warn().Msg("Unsupported slice type detected, no flag will be bound")
			err = nil
			v = false
		} else if err != nil {
			return err
		}
	case reflect.Map:
		l.Warn().Msg("Map value detected, no flag will be bound")
		v = false
	case reflect.Struct:
		return p.processFields(n, field.Type, cmd)
	case reflect.Ptr:
		return p.processFields(n, field.Type.Elem(), cmd)
	default:
		return errors.New("inavlid type supplied")
	}

	if v {
		l.Debug().Msg("Field registered")
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

func (p *Binder) processSlice(n, def, desc string, field reflect.StructField, cmd *cobra.Command) (err error) {
	s := &JSONStruct{}
	switch field.Type.Elem().Kind() {
	case reflect.Int:
		def := fmt.Sprintf("{\"intarray\":%s}", def)
		err := json.Unmarshal([]byte(def), &s)
		if err != nil {
			p.log.Error().Err(err).Msg("Error unmarshalling json")
			return err
		}
		cmd.Flags().IntSlice(n, s.IntArray, desc)
	case reflect.String:
		def := fmt.Sprintf("{\"stringarray\":%s}", def)
		err := json.Unmarshal([]byte(def), &s)
		if err != nil {
			p.log.Error().Err(err).Msg("Error unmarshalling json")
			return err
		}
		cmd.Flags().StringSlice(n, s.StringArray, desc)
	case reflect.Float64:
		def := fmt.Sprintf("{\"float64array\":%s}", def)
		err := json.Unmarshal([]byte(def), &s)
		if err != nil {
			p.log.Error().Err(err).Msg("Error unmarshalling json")
			return err
		}
		cmd.Flags().Float64Slice(n, s.Float64Array, desc)
	case reflect.Float32:
		def := fmt.Sprintf("{\"float32array\":%s}", def)
		err := json.Unmarshal([]byte(def), &s)
		if err != nil {
			p.log.Error().Err(err).Msg("Error unmarshalling json")
			return err
		}
		cmd.Flags().Float32Slice(n, s.Float32Array, desc)
	case reflect.Int32:
		def := fmt.Sprintf("{\"int32array\":%s}", def)
		err := json.Unmarshal([]byte(def), &s)
		if err != nil {
			p.log.Error().Err(err).Msg("Error unmarshalling json")
			return err
		}
		cmd.Flags().Int32Slice(n, s.Int32Array, desc)
	case reflect.Int64:
		def := fmt.Sprintf("{\"int64array\":%s}", def)
		err := json.Unmarshal([]byte(def), &s)
		if err != nil {
			p.log.Error().Err(err).Msg("Error unmarshalling json")
			return err
		}
		cmd.Flags().Int64Slice(n, s.Int64Array, desc)
	case reflect.Bool:
		def := fmt.Sprintf("{\"boolarray\":%s}", def)
		err := json.Unmarshal([]byte(def), &s)
		if err != nil {
			p.log.Error().Err(err).Msg("Error unmarshalling json")
			return err
		}
		cmd.Flags().BoolSlice(n, s.BoolArray, desc)
	default:
		return fmt.Errorf("unsupported type for slice: %s", n)
	}

	err = viper.BindPFlag(n, cmd.Flags().Lookup(n))
	if err != nil {
		return fmt.Errorf("error binding flag: %w", err)
	}

	return err
}
