package component

import (
	"fmt"
	"reflect"
)

var VOID = &TLValue{}
var NULL = &TLValue{}
var INVALID = &TLValue{}

type TLValue struct {
	value interface{}
}

func (v *TLValue) isInt() bool {
	return reflect.TypeOf(v.value).Kind() == reflect.Int
}

func (v *TLValue) asInt() int {
	if v.isDouble() {
		return int(v.asDouble())
	}
	return v.value.(int)
}

func (v *TLValue) isDouble() bool {
	return reflect.TypeOf(v.value).Kind() == reflect.Float64
}

func (v *TLValue) asDouble() float64 {
	if v.isInt() {
		return float64(v.asInt())
	}
	return v.value.(float64)
}

func (v *TLValue) isNumber() bool {
	return v.isInt() || v.isDouble()
}

func (v *TLValue) isString() bool {
	return reflect.TypeOf(v.value).Kind() == reflect.String
}

func (v *TLValue) asString() string {
	return v.value.(string)
}

func (v *TLValue) isBool() bool {
	return reflect.TypeOf(v.value).Kind() == reflect.Bool
}

func (v *TLValue) asBool() bool {
	return v.value.(bool)
}

func (v *TLValue) isList() bool {
	return reflect.TypeOf(v.value).Kind() == reflect.Slice
}

func (v *TLValue) asList() []*TLValue {
	return v.value.([]*TLValue)
}

func (v *TLValue) equals(other *TLValue) bool {
	if v == NULL {
		return other == NULL
	}
	if v == other {
		return true
	}
	if v.isInt() && other.isInt() {
		return v.asInt() == other.asInt()
	}
	if v.isDouble() && other.isDouble() {
		return v.asDouble() == other.asDouble()
	}
	if v.isString() && other.isString() {
		return v.asString() == other.asString()
	}
	if v.isBool() && other.isBool() {
		return v.asBool() == other.asBool()
	}
	return false
}

func (v *TLValue) String() string {
	if v.isList() {
		arr := v.asList()
		res := "["
		for i := 0; i < len(arr)-1; i += 1 {
			res += arr[i].String() + ", "
		}
		res += arr[len(arr)-1].String() + "]"
		return res
	} else {
		return fmt.Sprintf("%v", v.value)
	}
}
