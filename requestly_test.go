package requestly

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	Convey("JSON", t, func() {
		req := New()
		body := `{"key":"value"}`

		Convey("GetJSON", func() {
			Convey("Success", func() {
				res, err := GetJSON(req, "https://httpbin.org/get")
				So(err, ShouldBeNil)
				So(res.Get("url").String(), ShouldResemble, "https://httpbin.org/get")
			})
			Convey("Failure", func() {
				_, err := GetJSON(req, "https://xxxxxx.org/get")
				So(err, ShouldBeError)
			})
		})

		Convey("DeleteJSON", func() {
			res, err := DeleteJSON(req, "https://httpbin.org/delete")
			So(err, ShouldBeNil)
			So(res.Get("url").String(), ShouldResemble, "https://httpbin.org/delete")
		})

		Convey("PutJSON", func() {
			res, err := PutJSON(req, "https://httpbin.org/put", body)
			So(err, ShouldBeNil)
			So(res.Get("url").String(), ShouldResemble, "https://httpbin.org/put")
			So(res.Get("json.key").String(), ShouldResemble, "value")
		})

		Convey("PostJSON", func() {
			res, err := PostJSON(req, "https://httpbin.org/post", body)
			So(err, ShouldBeNil)
			So(res.Get("url").String(), ShouldResemble, "https://httpbin.org/post")
			So(res.Get("json.key").String(), ShouldResemble, "value")
		})

		Convey("PatchJSON", func() {
			res, err := PatchJSON(req, "https://httpbin.org/patch", body)
			So(err, ShouldBeNil)
			So(res.Get("url").String(), ShouldResemble, "https://httpbin.org/patch")
			So(res.Get("json.key").String(), ShouldResemble, "value")
		})
	})
}
