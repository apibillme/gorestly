# stubby - stubbing for Go (golang)

[![Go Report](https://goreportcard.com/badge/github.com/apibillme/stubby)](https://goreportcard.com/report/github.com/apibillme/stubby) [![GolangCI](https://golangci.com/badges/github.com/apibillme/stubby.svg)](https://golangci.com/r/github.com/apibillme/stubby) [![Travis](https://travis-ci.org/apibillme/stubby.svg?branch=master)](https://travis-ci.org/apibillme/stubby#) [![codecov](https://codecov.io/gh/apibillme/stubby/branch/master/graph/badge.svg)](https://codecov.io/gh/apibillme/stubby) ![License](https://img.shields.io/github/license/apibillme/stubby.svg) ![Maintenance](https://img.shields.io/maintenance/yes/2018.svg) [![GoDoc](https://godoc.org/github.com/apibillme/stubby?status.svg)](https://godoc.org/github.com/apibillme/stubby)


```bash
go get github.com/apibillme/stubby
```

```go
var timeNow = time.Now

func GetDate() int {
	return timeNow().Day()
}

stubs := stubby.Stub(&timeNow, func() time.Time {
  return time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC)
})
defer stubs.Reset()
```

```go
// production code:
var osHostname = osHostname

// Test code:
stubs := stubby.StubFunc(&osHostname, "fakehost", nil)
defer stubs.Reset()
```
