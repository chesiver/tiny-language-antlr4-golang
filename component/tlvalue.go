package component

import "reflect"

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
