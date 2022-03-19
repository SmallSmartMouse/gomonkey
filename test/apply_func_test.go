package test

import (
	"encoding/json"
	"testing"

	. "github.com/SmallSmartMouse/gomonkey/v2"
	"github.com/SmallSmartMouse/gomonkey/v2/test/fake"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	outputExpect = "xxx-vethName100-yyy"
)

func TestApplyFunc(t *testing.T) {
	Convey("TestApplyFunc", t, func() {

		{
			patches := ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
				return outputExpect, nil
			})
			defer patches.Reset()
			output, err := fake.Exec("", "")
			So(err, ShouldEqual, nil)
			So(output, ShouldEqual, outputExpect)
		}

		{
			patches := ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
				return "", fake.ErrActual
			})
			defer patches.Reset()
			output, err := fake.Exec("", "")
			So(err, ShouldEqual, fake.ErrActual)
			So(output, ShouldEqual, "")
		}

		{
			patches := ApplyFunc(fake.Exec, func(_ string, _ ...string) (string, error) {
				return outputExpect, nil
			})
			defer patches.Reset()
			patches.ApplyFunc(fake.Belong, func(_ string, _ []string) bool {
				return true
			})
			output, err := fake.Exec("", "")
			So(err, ShouldEqual, nil)
			So(output, ShouldEqual, outputExpect)
			flag := fake.Belong("", nil)
			So(flag, ShouldBeTrue)
		}

		{
			patches := ApplyFunc(json.Unmarshal, func(data []byte, v interface{}) error {
				if data == nil {
					panic("input param is nil!")
				}
				p := v.(*map[int]int)
				*p = make(map[int]int)
				(*p)[1] = 2
				(*p)[2] = 4
				return nil
			})
			defer patches.Reset()
			var m map[int]int
			err := json.Unmarshal([]byte("123"), &m)
			So(err, ShouldEqual, nil)
			So(m[1], ShouldEqual, 2)
			So(m[2], ShouldEqual, 4)
		}
	})
}
