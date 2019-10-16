package stub

// Func replaces a function variable with a function that returns stubVal.
// funcVarToStub must be a pointer to a function variable. If the function
// returns multiple values, then multiple values should be passed to stubFunc.
// The values must match be assignable to the return values' types.
func Func(fn interface{}, returns ...interface{}) Stubber {
	return New().Func(fn, returns...)
}
