package component

type TLValue struct {
	value interface{}
}

func (v *TLValue) asDouble() float64 {
	return v.value.(float64)
}
