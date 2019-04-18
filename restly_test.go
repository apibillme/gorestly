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

		Convey("JSON", t, func() {
			body := `{"key":"value"}`

			Convey("GetJSON", t, func() {
				Convey("Success - 200", func() {
					res, code, err := GetJSON(req, "https://mockbin.com/request", `?foo=http://foobar.com&bar=baz`)
					So(err, ShouldBeNil)
					So(res.Get("method").String(), ShouldResemble, "GET")
					So(res.Get("queryString.foo").String(), ShouldResemble, `http://foobar.com`)
					So(res.Get("queryString.bar").String(), ShouldResemble, `baz`)
					So(code, ShouldResemble, 200)
				})
				Convey("Success - 404", func() {
					_, code, err := GetJSON(req, "https://apibill.me/footest", "")
					So(err, ShouldBeNil)
					So(code, ShouldResemble, 404)
				})
				Convey("Failure - invalid host", func() {
					_, _, err := GetJSON(req, "https://xxxxxx.org/get", "")
					So(err, ShouldBeError)
				})
				Convey("Failure - uri parse fails", func() {
					stub1 := stubby.StubFunc(&uriParse, nil, errors.New("foo"))
					defer stub1.Reset()
					_, _, err := GetJSON(req, "https://mockbin.com/request", "")
					So(err, ShouldBeError)
				})
			})

			Convey("DeleteJSON", t, func() {
				res, _, err := DeleteJSON(req, "https://mockbin.com/request", "")
				So(err, ShouldBeNil)
				So(res.Get("method").String(), ShouldResemble, "DELETE")
			})

			Convey("PutJSON", t, func() {
				res, _, err := PutJSON(req, "https://mockbin.com/request", body, "")
				So(err, ShouldBeNil)
				So(res.Get("method").String(), ShouldResemble, "PUT")
				extraParsing := gjson.Parse(res.Get("postData.text").String())
				So(extraParsing.Get("key").String(), ShouldResemble, "value")
			})

			Convey("PostJSON", t, func() {
				res, _, err := PostJSON(req, "https://mockbin.com/request", body, "")
				So(err, ShouldBeNil)
				So(res.Get("method").String(), ShouldResemble, "POST")
				extraParsing := gjson.Parse(res.Get("postData.text").String())
				So(extraParsing.Get("key").String(), ShouldResemble, "value")
			})

			Convey("PatchJSON", t, func() {
				res, _, err := PatchJSON(req, "https://mockbin.com/request", body, "")
				So(err, ShouldBeNil)
				So(res.Get("method").String(), ShouldResemble, "PATCH")
				extraParsing := gjson.Parse(res.Get("postData.text").String())
				So(extraParsing.Get("key").String(), ShouldResemble, "value")
			})
		})

		Convey("XML", t, func() {
			body := `<?xml version="1.0" encoding="UTF-8"?><People><Person name="Jon"/><Person name="Sally"/></People></xml>`

			Convey("GetXML", t, func() {
				Convey("Success", func() {
					res, _, err := GetXML(req, "https://mockbin.com/request", `?foo=http://foobar.com&bar=baz`)
					So(err, ShouldBeNil)
					So(res.FindElement("/response/method").Text(), ShouldResemble, "GET")
					So(res.FindElement("/response/queryString/foo").Text(), ShouldResemble, `http://foobar.com`)
					So(res.FindElement("/response/queryString/bar").Text(), ShouldResemble, `baz`)
				})
				Convey("Failure - invalid host", func() {
					_, _, err := GetXML(req, "https://xxxxxx.org/get", "")
					So(err, ShouldBeError)
				})
				Convey("Failure - stream fails as not XML", func() {
					_, _, err := GetXML(req, "https://github.com/apibillme/restly", "")
					So(err, ShouldBeError)
				})
				Convey("Failure - uri parse fails", func() {
					stub1 := stubby.StubFunc(&uriParse, nil, errors.New("foo"))
					defer stub1.Reset()
					_, _, err := GetXML(req, "https://mockbin.com/request", "")
					So(err, ShouldBeError)
				})
			})

			Convey("DeleteXML", t, func() {
				res, _, err := DeleteXML(req, "https://mockbin.com/request", "")
				So(err, ShouldBeNil)
				method := res.FindElement("/response/method").Text()
				So(method, ShouldResemble, "DELETE")
			})

			Convey("PutXML", t, func() {
				res, _, err := PutXML(req, "https://mockbin.com/request", body, "")
				So(err, ShouldBeNil)
				method := res.FindElement("/response/method").Text()
				So(method, ShouldResemble, "PUT")
				xml := res.FindElement("/response/postData/text").Text()
				So(xml, ShouldResemble, body)
			})

			Convey("PostXML", t, func() {
				res, _, err := PostXML(req, "https://mockbin.com/request", body, "")
				So(err, ShouldBeNil)
				method := res.FindElement("/response/method").Text()
				So(method, ShouldResemble, "POST")
				xml := res.FindElement("/response/postData/text").Text()
				So(xml, ShouldResemble, body)
			})

			Convey("PatchXML", t, func() {
				res, _, err := PatchXML(req, "https://mockbin.com/request", body, "")
				So(err, ShouldBeNil)
				method := res.FindElement("/response/method").Text()
				So(method, ShouldResemble, "PATCH")
				xml := res.FindElement("/response/postData/text").Text()
				So(xml, ShouldResemble, body)
			})
		})
	})
}
