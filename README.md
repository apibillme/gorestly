# Requestly - Simple JSON Request library for Go (golang)

[![Go Report](https://goreportcard.com/badge/github.com/apibillme/requestly)](https://goreportcard.com/report/github.com/apibillme/requestly) [![Travis](https://travis-ci.org/apibillme/requestly.svg?branch=master)](https://travis-ci.org/apibillme/requestly#) [![codecov](https://codecov.io/gh/apibillme/requestly/branch/master/graph/badge.svg)](https://codecov.io/gh/apibillme/requestly)


This <100 LOC library combines [fasthttp](https://github.com/valyala/fasthttp) for performance and [gjson](https://github.com/tidwall/gjson) for JSON searching.

```bash
go get github.com/apibillme/requestly
```

```go
req := requestly.New()

req.Header.Add("Authorization", "Bearer my_token")

body := `{"key":"value"}`

res, err := requestly.GetJSON(req, "https://httpbin.org/get")
	
res, err := requestly.DeleteJSON(req, "https://httpbin.org/delete")
	
res, err := requestly.PutJSON(req, "https://httpbin.org/put", body)
	
res, err := requestly.PostJSON(req, "https://httpbin.org/post", body)
	
res, err := requestly.PatchJSON(req, "https://httpbin.org/patch", body)
```
