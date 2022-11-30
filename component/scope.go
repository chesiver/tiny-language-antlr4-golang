package component

type Scope struct {
	parent     *Scope
	variables  map[string]*TLValue
	isFunction bool
}

func NewScope() *Scope {
	return &Scope{nil, make(map[string]*TLValue), false}
}

func (s *Scope) assignParam(name string, val *TLValue) {
	s.variables[name] = val
}

func (s *Scope) resolve(name string) *TLValue {
	val, existed := s.variables[name]
	if existed {
		return val
	} else if !s.isFunction && s.parent != nil {
		return s.parent.resolve(name)
	} else {
		return nil
	}
}

func (s *Scope) reAssign(name string, val *TLValue) {
	_, existed := s.variables[name]
	if existed {
		s.variables[name] = val
	} else if s.parent != nil {
		s.parent.reAssign(name, val)
	}
}

func (s *Scope) Assign(name string, val *TLValue) {
	if !s.isFunction && s.resolve(name) != nil {
		s.reAssign(name, val)
	} else {
		s.variables[name] = val
	}
}
