package requestly

import (
	"testing"

	"github.com/tidwall/gjson"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	Convey("JSON", t, func() {
		req := New()
		body := `{"key":"value"}`

		Convey("GetJSON", func() {
			Convey("Success", func() {
				res, err := GetJSON(req, "https://mockbin.com/request")
				So(err, ShouldBeNil)
				So(res.Get("method").String(), ShouldResemble, "GET")
			})
			Convey("Failure", func() {
				_, err := GetJSON(req, "https://xxxxxx.org/get")
				So(err, ShouldBeError)
			})
		})

		Convey("DeleteJSON", func() {
			res, err := DeleteJSON(req, "https://mockbin.com/request")
			So(err, ShouldBeNil)
			So(res.Get("method").String(), ShouldResemble, "DELETE")
		})

		Convey("PutJSON", func() {
			res, err := PutJSON(req, "https://mockbin.com/request", body)
			So(err, ShouldBeNil)
			So(res.Get("method").String(), ShouldResemble, "PUT")
			extraParsing := gjson.Parse(res.Get("postData.text").String())
			So(extraParsing.Get("key").String(), ShouldResemble, "value")
		})

		Convey("PostJSON", func() {
			res, err := PostJSON(req, "https://mockbin.com/request", body)
			So(err, ShouldBeNil)
			So(res.Get("method").String(), ShouldResemble, "POST")
			extraParsing := gjson.Parse(res.Get("postData.text").String())
			So(extraParsing.Get("key").String(), ShouldResemble, "value")
		})

		Convey("PatchJSON", func() {
			res, err := PatchJSON(req, "https://mockbin.com/request", body)
			So(err, ShouldBeNil)
			So(res.Get("method").String(), ShouldResemble, "PATCH")
			extraParsing := gjson.Parse(res.Get("postData.text").String())
			So(extraParsing.Get("key").String(), ShouldResemble, "value")
		})
	})
}
