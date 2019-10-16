package stub

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Variables used in stubbing.
var v1, v2, v3, v4 int

// resetVars is used to reset the variables used in stubbing tests to their default values.
func resetVars() {
	v1 = 100
	v2 = 200
	v3 = 300
	v4 = 400
}

func TestVarAndRestore(t *testing.T) {
	Convey("如果变量直接打桩，会 panic", t, func() {
		v := 1
		So(func() { Var(v, 2) }, ShouldPanicWith, "original parameter should be a pointer")
	})

	Convey("如果存在变量 v 的值为 100 ", t, func() {
		v := 100
		Convey("对其打桩成 1 后，", func() {
			s := Var(&v, 1)
			Convey("其值应该是 1", func() {
				So(v, ShouldEqual, 1)
			})
			Convey("对其还原后，应该是原先的值 100", func() {
				s.Restore()
				So(v, ShouldEqual, 100)
			})
		})
		Convey("对其连续打桩成的话，", func() {
			s := Var(&v, 1)
			Convey("其值应该是 1", func() {
				So(v, ShouldEqual, 1)
			})
			s.Var(&v, 2)
			Convey("其值应该是 2", func() {
				So(v, ShouldEqual, 2)
			})
			s.Var(&v, 3)
			Convey("其值应该是 3", func() {
				So(v, ShouldEqual, 3)
			})
			s.Var(&v, 4)
			Convey("其值应该是 4", func() {
				So(v, ShouldEqual, 4)
			})
			Convey("对其还原后，应该是原先的值 100", func() {
				s.Restore()
				So(v, ShouldEqual, 100)
			})
		})
		Convey("对其打桩成 1 后，再恢复", func() {
			Var(&v, 1).Restore()
			Convey("其值应该是原先的 100", func() {
				So(v, ShouldEqual, 100)
			})
			v = 200
			Convey("对其进行修改后，其值改变为 200", func() {
				So(v, ShouldEqual, 200)
			})
			Convey("对其打桩成 2 后，再恢复", func() {
				Var(&v, 2).Restore()
				Convey("其值应该是原先的 200", func() {
					So(v, ShouldEqual, 200)
				})
			})
		})
	})

	Convey("如果存在多个变量 a、b、c、d，", t, func() {
		a, b, c, d := 1, 2, 3, 4
		s := Var(&a, 10)
		Convey("a 被打桩成了 10，其余不变", func() {
			So(a, ShouldEqual, 10)
			So(b, ShouldEqual, 2)
			So(c, ShouldEqual, 3)
			So(d, ShouldEqual, 4)
		})
		s.Var(&b, 20)
		Convey("a 为 10，b 为 20，其余不变", func() {
			So(a, ShouldEqual, 10)
			So(b, ShouldEqual, 20)
			So(c, ShouldEqual, 3)
			So(d, ShouldEqual, 4)
		})
		s.Var(&c, 30)
		Convey("a 为 10，b 为 20，c 为 30，d 不变", func() {
			So(a, ShouldEqual, 10)
			So(b, ShouldEqual, 20)
			So(c, ShouldEqual, 30)
			So(d, ShouldEqual, 4)
		})
		s.Var(&d, 40)
		Convey("a 为 10，b 为 20，c 为 30，d 为 40", func() {
			So(a, ShouldEqual, 10)
			So(b, ShouldEqual, 20)
			So(c, ShouldEqual, 30)
			So(d, ShouldEqual, 40)
		})
		s.Restore()
		Convey("a、b、c、d 全部还原", func() {
			So(a, ShouldEqual, 1)
			So(b, ShouldEqual, 2)
			So(c, ShouldEqual, 3)
			So(d, ShouldEqual, 4)
		})
	})
}
