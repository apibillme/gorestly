# Requestly - Fast REST Client for Go (golang)

[![Go Report](https://goreportcard.com/badge/github.com/apibillme/requestly)](https://goreportcard.com/report/github.com/apibillme/requestly) [![Travis](https://travis-ci.org/apibillme/requestly.svg?branch=master)](https://travis-ci.org/apibillme/requestly#) [![codecov](https://codecov.io/gh/apibillme/requestly/branch/master/graph/badge.svg)](https://codecov.io/gh/apibillme/requestly) ![License](https://img.shields.io/github/license/mashape/apistatus.svg) ![Maintenance](https://img.shields.io/maintenance/yes/2018.svg) [![GoDoc](https://godoc.org/github.com/apibillme/requestly?status.svg)](https://godoc.org/github.com/apibillme/requestly)


This fast REST client combines [fasthttp](https://github.com/valyala/fasthttp#readme) for performance, [gjson](https://github.com/tidwall/gjson#readme) for JSON searching, and [etree](https://github.com/beevik/etree#readme) for XML searching.

```bash
go get github.com/apibillme/requestly
```

```go
req := requestly.New()

req.Header.Add("Authorization", "Bearer my_token")

jsonBody := `{"key":"value"}`

xmlBody := `<?xml version="1.0" encoding="UTF-8"?><People><Person name="Jon"/></People></xml>`

res, err := requestly.GetJSON(req, "https://mockbin.com/request")
	
res, err := requestly.DeleteJSON(req, "https://mockbin.com/request")
	
res, err := requestly.PutJSON(req, "https://mockbin.com/request", jsonBody)
	
res, err := requestly.PostJSON(req, "https://mockbin.com/request", jsonBody)
	
res, err := requestly.PatchJSON(req, "https://mockbin.com/request", jsonBody)

res, err := requestly.GetXML(req, "https://mockbin.com/request")
	
res, err := requestly.DeleteXML(req, "https://mockbin.com/request")
	
res, err := requestly.PutXML(req, "https://mockbin.com/request", xmlBody)
	
res, err := requestly.PostXML(req, "https://mockbin.com/request", xmlBody)
	
res, err := requestly.PatchXML(req, "https://mockbin.com/request", xmlBody)
```

## Motivation

I saw the largest problem with Go (Golang) being interacting with JSON & XML with REST clients. At the time the popular REST clients required you to strongly type out each interface you need for the request. This is painful and slow! 

I wanted a one-liner request with the ability to dynamically set, find, and extract values from JSON & XML without all the boilerplate of a [net/http](https://golang.org/pkg/net/http/) request. This library delivers you exactly these requirements!

The request body is simply a string and the find/extract interface relies on battle-tested libraries for either JSON ([gjson](https://github.com/tidwall/gjson#readme)) or XML ([etree](https://github.com/beevik/etree#readme)).

Because this library uses [fasthttp](https://github.com/valyala/fasthttp#readme) rather than [net/http](https://golang.org/pkg/net/http/) it is about 10x faster than competing libraries. It is also only about 100 LOC compared to the massive codebases of competing projects.

## Compare the REST client competition

* [grequests](https://github.com/levigross/grequests)
* [resty](https://github.com/go-resty/resty)
