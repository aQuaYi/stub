package stub

import (
	"reflect"
)

// Var replaces the *original with fake variable.
func Var(original interface{}, fake interface{}) Stubber {
	return newStubs().Var(original, fake)
}

// Var replaces the value stored at varToStub with stubVal.
// varToStub must be a pointer to the variable. stubVal should have a type
// that is assignable to the variable.
func (s *stubs) Var(original, fake interface{}) Stubber {
	o := reflect.ValueOf(original)
	f := reflect.ValueOf(fake)

	// Ensure original is a pointer to the variable.
	if o.Type().Kind() != reflect.Ptr {
		panic("original parameter should be a pointer")
	}

	if _, ok := s.vars[o]; !ok {
		// Store the original value if this is the first time varPtr is being stubbed.
		// TODO: 我知道 Interface 与 ValueOf 是互逆的操作。也知道这样会把 original 的 Value 值重新复制一个。但是，我不知道通过其他的操作，达到和这个一样的效果。
		s.vars[o] = reflect.ValueOf(o.Elem().Interface())
	}

	o.Elem().Set(f)
	return s
}
