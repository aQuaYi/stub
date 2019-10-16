package stub

import (
	"reflect"
)

// Var replaces the *original with fake variable.
func Var(original interface{}, fake interface{}) Stubber {
	return New().Var(original, fake)
}

// Var replaces the value stored at varToStub with stubVal.
// varToStub must be a pointer to the variable. stubVal should have a type
// that is assignable to the variable.
func (s *stubs) Var(varToStub interface{}, stubVal interface{}) Stubber {
	v := reflect.ValueOf(varToStub)
	stub := reflect.ValueOf(stubVal)

	// Ensure varToStub is a pointer to the variable.
	if v.Type().Kind() != reflect.Ptr {
		panic("variable to stub is expected to be a pointer")
	}

	if _, ok := s.vars[v]; !ok {
		// Store the original value if this is the first time varPtr is being stubbed.
		s.vars[v] = reflect.ValueOf(v.Elem().Interface())
	}

	// *varToStub = stubVal
	v.Elem().Set(stub)
	return s
}

// TODO:  删除此处内容
// // ResetSingle resets a single stubbed variable back to its original value.
// func (s *stubs) ResetSingle(varToStub interface{}) {
// 	v := reflect.ValueOf(varToStub)
// 	originalVal, ok := s.vars[v]
// 	if !ok {
// 		panic("cannot reset variable as it has not been stubbed yet")
// 	}
// 	v.Elem().Set(originalVal)
// }
