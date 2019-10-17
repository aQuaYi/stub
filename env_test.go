package stub

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestStubEnv(t *testing.T) {
	os.Setenv("GOSTUB_T1", "V1")
	os.Setenv("GOSTUB_T2", "V2")
	os.Unsetenv("GOSTUB_NONE")

	stubs := newStubs()

	stubs.Env("GOSTUB_NONE", "a")
	stubs.Env("GOSTUB_T1", "1")
	stubs.Env("GOSTUB_T1", "2")
	stubs.Env("GOSTUB_T1", "3")
	stubs.Env("GOSTUB_T2", "4")
	// stubs.UnsetEnv("GOSTUB_T2")

	assert.Equal(t, "3", os.Getenv("GOSTUB_T1"), "Wrong value for T1")
	// assert.Equal(t, "", os.Getenv("GOSTUB_T2"), "Wrong value for T2")
	assert.Equal(t, "a", os.Getenv("GOSTUB_NONE"), "Wrong value for NONE")
	stubs.Restore()

	_, ok := os.LookupEnv("GOSTUB_NONE")
	assert.False(t, ok, "NONE should be unset")

	assert.Equal(t, "V1", os.Getenv("GOSTUB_T1"), "Wrong reset value for T1")
	assert.Equal(t, "V2", os.Getenv("GOSTUB_T2"), "Wrong reset value for T2")
}

func TestEnv(t *testing.T) {
	k, v := "TestEnv", "true"
	Convey("如果想要打桩了环境变量 \"TestEnv\" 为 \"true\"", t, func() {
		Convey("在打桩前，该环境变量应该不存在", func() {
			_, ok := os.LookupEnv(k)
			So(ok, ShouldBeFalse)
		})
		stubs := Env(k, v)
		Convey("打桩后，可以查询到该环境变量", func() {
			actual, ok := os.LookupEnv(k)
			So(ok, ShouldBeTrue)
			So(actual, ShouldEqual, v)
		})
		stubs.Restore()
		Convey("恢复后，该环境变量被删除", func() {
			_, ok := os.LookupEnv(k)
			So(ok, ShouldBeFalse)
		})
	})
}
