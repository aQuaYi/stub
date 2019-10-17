package stub

import (
	"reflect"
)

// Stubber 包含了打桩时所需的方法，以及恢复的方法。
type Stubber interface {
	// Stub 会把 *originalPtr 变量替换成 fake 变量的内容。
	// NOTICE: 函数在 Go 语言中，也是一种变量。
	Var(originalPtr, fake interface{}) Stubber

	// StubFunc 会把原先 *funcPtr 替换成另一个函数，其具有固定的返回值 returns。
	Func(funcPtr interface{}, returns ...interface{}) Stubber

	// StubEnv 会更改环境变量的值。
	Env(key, value string) Stubber

	// Restore 会把 Stubber 替换过的所有值，全部还原。
	Restore()
}

// stubs represents a set of stubbed variables that can be reset.
type stubs struct {
	// vars is a map from the variable pointer (being stubbed) to the original value.
	vars map[reflect.Value]reflect.Value
	envs map[string]env
}

type env struct {
	val      string
	existing bool
}

// newStubs returns Stubs that can be used to stub out variables.
func newStubs() *stubs {
	return &stubs{
		// in vars:
		// 	key is
		// 	val is
		vars: make(map[reflect.Value]reflect.Value),

		// in envs:
		// 	key is
		// 	val is
		envs: make(map[string]env),
	}
}
