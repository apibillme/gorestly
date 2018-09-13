package restly

import (
	"errors"
	"testing"

	"github.com/apibillme/stubby"
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
					res, err := GetJSON(req, "https://mockbin.com/request", `?foo=http://foobar.com&bar=baz`)
					So(err, ShouldBeNil)
					So(res.Get("method").String(), ShouldResemble, "GET")
					So(res.Get("queryString.foo").String(), ShouldResemble, `http://foobar.com`)
					So(res.Get("queryString.bar").String(), ShouldResemble, `baz`)
				})
				Convey("Failure - invalid host", func() {
					_, err := GetJSON(req, "https://xxxxxx.org/get", "")
					So(err, ShouldBeError)
				})
				Convey("Failure - uri parse fails", func() {
					stub1 := stubby.StubFunc(&uriParse, nil, errors.New("foo"))
					defer stub1.Reset()
					_, err := GetJSON(req, "https://mockbin.com/request", "")
					So(err, ShouldBeError)
				})
			})

			Convey("DeleteJSON", func() {
				res, err := DeleteJSON(req, "https://mockbin.com/request", "")
				So(err, ShouldBeNil)
				So(res.Get("method").String(), ShouldResemble, "DELETE")
			})

			Convey("PutJSON", func() {
				res, err := PutJSON(req, "https://mockbin.com/request", body, "")
				So(err, ShouldBeNil)
				So(res.Get("method").String(), ShouldResemble, "PUT")
				extraParsing := gjson.Parse(res.Get("postData.text").String())
				So(extraParsing.Get("key").String(), ShouldResemble, "value")
			})

			Convey("PostJSON", func() {
				res, err := PostJSON(req, "https://mockbin.com/request", body, "")
				So(err, ShouldBeNil)
				So(res.Get("method").String(), ShouldResemble, "POST")
				extraParsing := gjson.Parse(res.Get("postData.text").String())
				So(extraParsing.Get("key").String(), ShouldResemble, "value")
			})

			Convey("PatchJSON", func() {
				res, err := PatchJSON(req, "https://mockbin.com/request", body, "")
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
					res, err := GetXML(req, "https://mockbin.com/request", `?foo=http://foobar.com&bar=baz`)
					So(err, ShouldBeNil)
					So(res.FindElement("/response/method").Text(), ShouldResemble, "GET")
					So(res.FindElement("/response/queryString/foo").Text(), ShouldResemble, `http://foobar.com`)
					So(res.FindElement("/response/queryString/bar").Text(), ShouldResemble, `baz`)
				})
				Convey("Failure - invalid host", func() {
					_, err := GetXML(req, "https://xxxxxx.org/get", "")
					So(err, ShouldBeError)
				})
				Convey("Failure - stream fails as not XML", func() {
					_, err := GetXML(req, "https://github.com/apibillme/restly", "")
					So(err, ShouldBeError)
				})
				Convey("Failure - uri parse fails", func() {
					stub1 := stubby.StubFunc(&uriParse, nil, errors.New("foo"))
					defer stub1.Reset()
					_, err := GetXML(req, "https://mockbin.com/request", "")
					So(err, ShouldBeError)
				})
			})

			Convey("DeleteXML", func() {
				res, err := DeleteXML(req, "https://mockbin.com/request", "")
				So(err, ShouldBeNil)
				method := res.FindElement("/response/method").Text()
				So(method, ShouldResemble, "DELETE")
			})

			Convey("PutXML", func() {
				res, err := PutXML(req, "https://mockbin.com/request", body, "")
				So(err, ShouldBeNil)
				method := res.FindElement("/response/method").Text()
				So(method, ShouldResemble, "PUT")
				xml := res.FindElement("/response/postData/text").Text()
				So(xml, ShouldResemble, body)
			})

			Convey("PostXML", func() {
				res, err := PostXML(req, "https://mockbin.com/request", body, "")
				So(err, ShouldBeNil)
				method := res.FindElement("/response/method").Text()
				So(method, ShouldResemble, "POST")
				xml := res.FindElement("/response/postData/text").Text()
				So(xml, ShouldResemble, body)
			})

			Convey("PatchXML", func() {
				res, err := PatchXML(req, "https://mockbin.com/request", body, "")
				So(err, ShouldBeNil)
				method := res.FindElement("/response/method").Text()
				So(method, ShouldResemble, "PATCH")
				xml := res.FindElement("/response/postData/text").Text()
				So(xml, ShouldResemble, body)
			})
		})
	})
}
