package stub_test

import (
	"fmt"
	"os"

	. "github.com/aQuaYi/stub"
)

func ExampleFunc() {
	var osHostname = os.Hostname

	defer Func(&osHostname, "fakehost", nil).Restore()
	host, err := osHostname()

	fmt.Println("Host:", host, "err:", err)
	// Output:
	// Host: fakehost err: <nil>
}

func ExampleVar() {
	var counter = 100

	defer Var(&counter, 200).Restore()
	fmt.Println("Counter:", counter)
	// Output:
	// Counter: 200
}
