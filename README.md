# restly - Fast REST client for Go (golang)

[![Go Report](https://goreportcard.com/badge/github.com/apibillme/restly)](https://goreportcard.com/report/github.com/apibillme/restly) [![GolangCI](https://golangci.com/badges/github.com/apibillme/restly.svg)](https://golangci.com/r/github.com/apibillme/restly) [![Travis](https://travis-ci.org/apibillme/restly.svg?branch=master)](https://travis-ci.org/apibillme/restly#) [![codecov](https://codecov.io/gh/apibillme/restly/branch/master/graph/badge.svg)](https://codecov.io/gh/apibillme/restly) ![License](https://img.shields.io/github/license/apibillme/restly.svg) ![Maintenance](https://img.shields.io/maintenance/yes/2018.svg) [![GoDoc](https://godoc.org/github.com/apibillme/restly?status.svg)](https://godoc.org/github.com/apibillme/restly)


This fast REST client combines [fasthttp](https://github.com/valyala/fasthttp#readme) for performance, [gjson](https://github.com/tidwall/gjson#readme) for JSON searching, and [etree](https://github.com/beevik/etree#readme) for XML searching.

## Features:
* Fully configurable fasthttp.Request struct
* Auto encoding REST routes with client libraries ([gjson](https://github.com/tidwall/gjson#readme) & [etree](https://github.com/beevik/etree#readme)) with support for problem errors [RFC7807](https://tools.ietf.org/html/rfc7807)
* About 100 LOC
* About 10x faster than using net/http

```bash
go get github.com/apibillme/restly
```

```go
req := restly.New()

req.Header.Add("Authorization", "Bearer my_token")

jsonBody := `{"key":"value"}`

xmlBody := `<?xml version="1.0" encoding="UTF-8"?><People><Person name="Jon"/></People></xml>`

res, err := restly.GetJSON(req, "https://mockbin.com/request", "?foo=bar")
	
res, err := restly.DeleteJSON(req, "https://mockbin.com/request", "?foo=bar")
	
res, err := restly.PutJSON(req, "https://mockbin.com/request", jsonBody, "?foo=bar")
	
res, err := restly.PostJSON(req, "https://mockbin.com/request", jsonBody, "?foo=bar")
	
res, err := restly.PatchJSON(req, "https://mockbin.com/request", jsonBody, "?foo=bar")

res, err := restly.GetXML(req, "https://mockbin.com/request", "?foo=bar")
	
res, err := restly.DeleteXML(req, "https://mockbin.com/request", "?foo=bar")
	
res, err := restly.PutXML(req, "https://mockbin.com/request", xmlBody, "?foo=bar")
	
res, err := restly.PostXML(req, "https://mockbin.com/request", xmlBody, "?foo=bar")
	
res, err := restly.PatchXML(req, "https://mockbin.com/request", xmlBody, "?foo=bar")
```

## Motivation

I saw the largest problem with Go (Golang) being interacting with JSON & XML with REST clients. At the time the popular REST clients required you to strongly type out each interface you need for the request. This is painful and slow! 

I wanted a one-liner request with the ability to dynamically set, find, and extract values from JSON & XML without all the boilerplate of a [net/http](https://golang.org/pkg/net/http/) request. This library delivers you exactly these requirements!

The request body is simply a string and the find/extract interface relies on battle-tested libraries for either JSON ([gjson](https://github.com/tidwall/gjson#readme)) or XML ([etree](https://github.com/beevik/etree#readme)).

Because this library uses [fasthttp](https://github.com/valyala/fasthttp#readme) rather than [net/http](https://golang.org/pkg/net/http/) it is about 10x faster than competing libraries. It is also only about 100 LOC compared to the massive codebases of competing projects.
