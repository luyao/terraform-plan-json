package jsonplan

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// InstanceKey represents the key of an instance within an object that
// contains multiple instances due to using "count" or "for_each" arguments
// in configuration.
//
// IntKey and StringKey are the two implementations of this type. No other
// implementations are allowed. The single instance of an object that _isn't_
// using "count" or "for_each" is represented by NoKey, which is a nil
// InstanceKey.
type InstanceKey struct {
	intKey    *IntKey
	stringKey *StringKey
}

// String is override to generate string index, no matter it is a IntKey
// or StringKey.
func (k *InstanceKey) String() string {
	if k.intKey != nil {
		return k.intKey.String()
	}

	if k.stringKey != nil {
		return k.stringKey.String()
	}

	return ""
}

// UnmarshalJSON is a customized json unmarshaler, it will
// unmarshal json object in string or float64 format into
// InstanceKey.
func (k *InstanceKey) UnmarshalJSON(data []byte) error {
	var value interface{}
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	items := reflect.ValueOf(value)
	switch items.Kind() {
	case reflect.String:
		k.stringKey = new(StringKey)
		*k.stringKey = StringKey(items.String())
		return nil
	case reflect.Float64:
		k.intKey = new(IntKey)
		*k.intKey = IntKey(int64(items.Float()))
		return nil
	default:
		return fmt.Errorf("can not unmarshal Json to instance key, %v", items)
	}
}

// IntKey is the InstanceKey representation representing integer indices, as
// used when the "count" argument is specified or if for_each is used with
// a sequence type.
type IntKey int64

func (k IntKey) String() string {
	return fmt.Sprintf("[%d]", int(k))
}

// StringKey is the InstanceKey representation representing string indices, as
// used when the "for_each" argument is specified with a map or object type.
type StringKey string

func (k StringKey) String() string {
	// FIXME: This isn't _quite_ right because Go's quoted string syntax is
	// slightly different than HCL's, but we'll accept it for now.
	return fmt.Sprintf("[%q]", string(k))
}
