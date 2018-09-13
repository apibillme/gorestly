package restly

import (
	"github.com/tidwall/gjson"

	"github.com/valyala/fasthttp"

	"github.com/beevik/etree"

	uri "net/url"
)

// for stubbing
var uriParse = uri.Parse

// New - create fasthttp request
func New() *fasthttp.Request {
	return &fasthttp.Request{}
}

func requestXML(req *fasthttp.Request, url string, query string) (*etree.Document, error) {
	urlQ, err := uriParse(url + query)
	if err != nil {
		return nil, err
	}
	q := urlQ.Query().Encode()
	req.SetRequestURI(url + `?` + q)
	res := &fasthttp.Response{}
	err = fasthttp.Do(req, res)
	doc := etree.NewDocument()
	if err != nil {
		return doc, err
	}
	err = doc.ReadFromBytes(res.Body())
	if err != nil {
		return doc, err
	}
	return doc, nil
}

func requestJSON(req *fasthttp.Request, url string, query string) (gjson.Result, error) {
	urlQ, err := uriParse(url + query)
	if err != nil {
		return gjson.Parse(""), err
	}
	q := urlQ.Query().Encode()
	req.SetRequestURI(url + `?` + q)
	res := &fasthttp.Response{}
	err = fasthttp.Do(req, res)
	if err != nil {
		return gjson.Parse(""), err
	}
	return gjson.ParseBytes(res.Body()), nil
}

func setJSONRequest(req *fasthttp.Request, method string, body string) *fasthttp.Request {
	req.Header.SetCanonical([]byte("Content-Type"), []byte("application/problem+json"))
	req.Header.Set("Accept", "application/json")
	req.Header.SetMethodBytes([]byte(method))
	req.SetBodyString(body)
	return req
}

func setXMLRequest(req *fasthttp.Request, method string, body string) *fasthttp.Request {
	req.Header.SetCanonical([]byte("Content-Type"), []byte("application/problem+xml"))
	req.Header.Set("Accept", "application/xml")
	req.Header.SetMethodBytes([]byte(method))
	req.SetBodyString(body)
	return req
}

// GetJSON - make get JSON and return searchable JSON
func GetJSON(req *fasthttp.Request, url string, query string) (gjson.Result, error) {
	req = setJSONRequest(req, "GET", "")
	return requestJSON(req, url, query)
}

// DeleteJSON - make delete JSON and return searchable JSON
func DeleteJSON(req *fasthttp.Request, url string, query string) (gjson.Result, error) {
	req = setJSONRequest(req, "DELETE", "")
	return requestJSON(req, url, query)
}

// PutJSON - make put JSON and return searchable JSON
func PutJSON(req *fasthttp.Request, url string, body string, query string) (gjson.Result, error) {
	req = setJSONRequest(req, "PUT", body)
	return requestJSON(req, url, query)
}

// PostJSON - make post JSON and return searchable JSON
func PostJSON(req *fasthttp.Request, url string, body string, query string) (gjson.Result, error) {
	req = setJSONRequest(req, "POST", body)
	return requestJSON(req, url, query)
}

// PatchJSON - make patch JSON and return searchable JSON
func PatchJSON(req *fasthttp.Request, url string, body string, query string) (gjson.Result, error) {
	req = setJSONRequest(req, "PATCH", body)
	return requestJSON(req, url, query)
}

// GetXML - make get XML and return searchable XML
func GetXML(req *fasthttp.Request, url string, query string) (*etree.Document, error) {
	req = setXMLRequest(req, "GET", "")
	return requestXML(req, url, query)
}

// DeleteXML - make delete XML and return searchable XML
func DeleteXML(req *fasthttp.Request, url string, query string) (*etree.Document, error) {
	req = setXMLRequest(req, "DELETE", "")
	return requestXML(req, url, query)
}

// PutXML - make put XML and return searchable XML
func PutXML(req *fasthttp.Request, url string, body string, query string) (*etree.Document, error) {
	req = setXMLRequest(req, "PUT", body)
	return requestXML(req, url, query)
}

// PostXML - make post XML and return searchable XML
func PostXML(req *fasthttp.Request, url string, body string, query string) (*etree.Document, error) {
	req = setXMLRequest(req, "POST", body)
	return requestXML(req, url, query)
}

// PatchXML - make patch XML and return searchable XML
func PatchXML(req *fasthttp.Request, url string, body string, query string) (*etree.Document, error) {
	req = setXMLRequest(req, "PATCH", body)
	return requestXML(req, url, query)
}
