package stub

import (
	"reflect"
)

// Var replaces the *original with fake variable.
func Var(original interface{}, fake interface{}) Stubber {
	return newStubber().Var(original, fake)
}

// Var replaces the value stored at varToStub with stubVal.
// varToStub must be a pointer to the variable. stubVal should have a type
// that is assignable to the variable.
func (s *stubs) Var(original, fake interface{}) Stubber {
	o := reflect.ValueOf(original)
	f := reflect.ValueOf(fake)

	// Ensure original is a pointer to the variable.
	if o.Type().Kind() != reflect.Ptr {
		panic("variable to stub is expected to be a pointer")
	}

	if _, ok := s.vars[o]; !ok {
		// Store the original value if this is the first time varPtr is being stubbed.
		s.vars[o] = reflect.ValueOf(o.Elem().Interface())
	}

	// *varToStub = stubVal
	o.Elem().Set(f)
	return s
}

func (s *stubs) restoreVars() {
	for v, originalVal := range s.vars {
		v.Elem().Set(originalVal)
	}
}
