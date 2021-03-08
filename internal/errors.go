package internal

import (
	"fmt"
	"reflect"
)

type InvalidKindError struct {
	Kind      reflect.Kind
	FieldName string
}

func (e *InvalidKindError) Error() string {
	return fmt.Sprintf("invalid kind %s for %s", e.Kind.String(), e.FieldName)
}
