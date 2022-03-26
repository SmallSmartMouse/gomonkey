package dsl_test

import (
	"github.com/SmallSmartMouse/gomonkey/test/fake"
	"testing"

	. "github.com/SmallSmartMouse/gomonkey"
	. "github.com/SmallSmartMouse/gomonkey/dsl"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPbBuilderFunc(t *testing.T) {
	Convey("TestPbBuilderFunc", t, func() {

		Convey("first dsl", func() {
			patches := NewPatches()
			defer patches.Reset()
			patchBuilder := NewPatchBuilder(patches)

			patchBuilder.
				Func(fake.Belong).
				Stubs().
				With(Eq("zxl"), Any()).
				Will(Return(true)).
				Then(Repeat(Return(false), 2)).
				End()

			flag := fake.Belong("zxl", []string{})
			So(flag, ShouldBeTrue)

			defer func() {
				if p := recover(); p != nil {
					str, ok := p.(string)
					So(ok, ShouldBeTrue)
					So(str, ShouldEqual, "input paras ddd is not matched")
				}
			}()
			fake.Belong("ddd", []string{"abc"})
		})

	})
}
