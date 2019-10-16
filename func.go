package stub

import (
	"fmt"
	"reflect"
)

// Func replaces a function variable with a function that returns stubVal.
// funcVarToStub must be a pointer to a function variable. If the function
// returns multiple values, then multiple values should be passed to stubFunc.
// The values must match be assignable to the return values' types.
func Func(fn interface{}, returns ...interface{}) Stubber {
	return New().Func(fn, returns...)
}

// Func replaces a function variable with a function that returns stubVal.
// funcVarToStub must be a pointer to a function variable. If the function
// returns multiple values, then multiple values should be passed to stubFunc.
// The values must match be assignable to the return values' types.
func (s *stubs) Func(funcVarToStub interface{}, stubVal ...interface{}) Stubber {
	funcPtrType := reflect.TypeOf(funcVarToStub)
	if funcPtrType.Kind() != reflect.Ptr ||
		funcPtrType.Elem().Kind() != reflect.Func {
		panic("func variable to stub must be a pointer to a function")
	}
	funcType := funcPtrType.Elem()
	if funcType.NumOut() != len(stubVal) {
		panic(fmt.Sprintf("func type has %v return values, but only %v stub values provided",
			funcType.NumOut(), len(stubVal)))
	}

	return s.Var(funcVarToStub, FuncReturning(funcPtrType.Elem(), stubVal...).Interface())
}

// FuncReturning creates a new function with type funcType that returns results.
func FuncReturning(funcType reflect.Type, results ...interface{}) reflect.Value {
	var resultValues []reflect.Value
	for i, r := range results {
		var retValue reflect.Value
		if r == nil {
			// We can't use reflect.ValueOf(nil), so we need to create the zero value.
			retValue = reflect.Zero(funcType.Out(i))
		} else {
			// We cannot simply use reflect.ValueOf(r) as that does not work for
			// interface types, as reflect.ValueOf receives the dynamic type, which
			// is the underlying type. e.g. for an error, it may *errors.errorString.
			// Instead, we make the return type's expected interface value using
			// reflect.New, and set the data to the passed in value.
			tempV := reflect.New(funcType.Out(i))
			tempV.Elem().Set(reflect.ValueOf(r))
			retValue = tempV.Elem()
		}
		resultValues = append(resultValues, retValue)
	}
	return reflect.MakeFunc(funcType, func(_ []reflect.Value) []reflect.Value {
		return resultValues
	})
}
