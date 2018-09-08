package restly

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/tidwall/gjson"
)

func TestSpec(t *testing.T) {

	Convey("New", t, func() {
		req := New()

		Convey("JSON", func() {
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

		Convey("XML", func() {
			body := `<?xml version="1.0" encoding="UTF-8"?><People><Person name="Jon"/><Person name="Sally"/></People></xml>`

			Convey("GetXML", func() {
				Convey("Success", func() {
					res, err := GetXML(req, "https://mockbin.com/request")
					So(err, ShouldBeNil)
					method := res.FindElement("/response/method").Text()
					So(method, ShouldResemble, "GET")
				})
				Convey("Failure Host", func() {
					_, err := GetXML(req, "https://xxxxxx.org/get")
					So(err, ShouldBeError)
				})
				Convey("Failure Stream", func() {
					_, err := GetXML(req, "https://github.com/apibillme/requestly")
					So(err, ShouldBeError)
				})
			})

			Convey("DeleteXML", func() {
				res, err := DeleteXML(req, "https://mockbin.com/request")
				So(err, ShouldBeNil)
				method := res.FindElement("/response/method").Text()
				So(method, ShouldResemble, "DELETE")
			})

			Convey("PutXML", func() {
				res, err := PutXML(req, "https://mockbin.com/request", body)
				So(err, ShouldBeNil)
				method := res.FindElement("/response/method").Text()
				So(method, ShouldResemble, "PUT")
				xml := res.FindElement("/response/postData/text").Text()
				So(xml, ShouldResemble, body)
			})

			Convey("PostXML", func() {
				res, err := PostXML(req, "https://mockbin.com/request", body)
				So(err, ShouldBeNil)
				method := res.FindElement("/response/method").Text()
				So(method, ShouldResemble, "POST")
				xml := res.FindElement("/response/postData/text").Text()
				So(xml, ShouldResemble, body)
			})

			Convey("PatchXML", func() {
				res, err := PatchXML(req, "https://mockbin.com/request", body)
				So(err, ShouldBeNil)
				method := res.FindElement("/response/method").Text()
				So(method, ShouldResemble, "PATCH")
				xml := res.FindElement("/response/postData/text").Text()
				So(xml, ShouldResemble, body)
			})
		})
	})
}
