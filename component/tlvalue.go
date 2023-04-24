package component

import "reflect"

var VOID = &TLValue{}
var NULL = &TLValue{}

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
