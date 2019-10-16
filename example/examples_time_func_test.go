package stub_test

import (
	"fmt"
	"time"

	. "github.com/aQuaYi/stub"
)

// Production code
var timeNow = time.Now

func GetDay() int {
	return timeNow().Day()
}

// Test code
func Example_stubTimeWithFunction() {
	var day = 2
	stubs := Var(&timeNow, func() time.Time {
		return time.Date(2015, 07, day, 0, 0, 0, 0, time.UTC)
	})
	defer stubs.Restore()

	firstDay := GetDay()

	day = 3
	secondDay := GetDay()

	fmt.Printf("First day: %v, second day: %v\n", firstDay, secondDay)
	// Output:
	// First day: 2, second day: 3
}

// Test code
func Example_stubTimeWithConstant() {
	stubs := Func(&timeNow, time.Date(2015, 07, 2, 0, 0, 0, 0, time.UTC))
	defer stubs.Restore()

	fmt.Println("Day:", GetDay())
	// Output:
	// Day: 2
}
