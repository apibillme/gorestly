package requestly

import (
	"github.com/tidwall/gjson"

	"github.com/valyala/fasthttp"
)

// New - create fasthttp request
func New() *fasthttp.Request {
	return &fasthttp.Request{}
}

func request(req *fasthttp.Request, uri string) (gjson.Result, error) {
	req.SetRequestURI(uri)
	res := &fasthttp.Response{}
	err := fasthttp.Do(req, res)
	if err != nil {
		return gjson.Parse(""), err
	}
	return gjson.ParseBytes(res.Body()), nil
}

func setJSONRequest(req *fasthttp.Request, method string, body string) *fasthttp.Request {
	req.Header.SetCanonical([]byte("Content-Type"), []byte("application/json"))
	req.Header.Set("accept", "application/json")
	req.Header.SetMethodBytes([]byte(method))
	req.SetBodyString(body)
	return req
}

// GetJSON - make get JSON and return searchable JSON
func GetJSON(req *fasthttp.Request, uri string) (gjson.Result, error) {
	req = setJSONRequest(req, "GET", "")
	return request(req, uri)
}

// DeleteJSON - make delete JSON and return searchable JSON
func DeleteJSON(req *fasthttp.Request, uri string) (gjson.Result, error) {
	req = setJSONRequest(req, "DELETE", "")
	return request(req, uri)
}

// PutJSON - make put JSON and return searchable JSON
func PutJSON(req *fasthttp.Request, uri string, body string) (gjson.Result, error) {
	req = setJSONRequest(req, "PUT", body)
	return request(req, uri)
}

// PostJSON - make post JSON and return searchable JSON
func PostJSON(req *fasthttp.Request, uri string, body string) (gjson.Result, error) {
	req = setJSONRequest(req, "POST", body)
	return request(req, uri)
}

// PatchJSON - make patch JSON and return searchable JSON
func PatchJSON(req *fasthttp.Request, uri string, body string) (gjson.Result, error) {
	req = setJSONRequest(req, "PATCH", body)
	return request(req, uri)
}
