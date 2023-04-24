package component

import "reflect"

var VOID = &TLValue{}
var NULL = &TLValue{}

type TLValue struct {
	value interface{}
}

func (v *TLValue) isDouble() bool {
	return reflect.TypeOf(v.value).Name() == "float64"
}

func (v *TLValue) asDouble() float64 {
	return v.value.(float64)
}

func (v *TLValue) isString() bool {
	return reflect.TypeOf(v.value).Name() == "string"
}

func (v *TLValue) asString() string {
	return v.value.(string)
}

func (v *TLValue) isBool() bool {
	return reflect.TypeOf(v.value).Name() == "bool"
}

func (v *TLValue) asBool() bool {
	return v.value.(bool)
}
