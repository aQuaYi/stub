# stub

[![Build Status](https://travis-ci.org/aQuaYi/stub.svg?branch=master)](https://travis-ci.org/aQuaYi/stub)
[![codecov](https://codecov.io/gh/aQuaYi/stub/branch/master/graph/badge.svg)](https://codecov.io/gh/aQuaYi/stub)
[![GoDoc](https://godoc.org/github.com/aQuaYi/stub?status.svg)](https://godoc.org/github.com/aQuaYi/stub)
[![Go Report Card](https://goreportcard.com/badge/github.com/aQuaYi/stub)](https://goreportcard.com/report/github.com/aQuaYi/stub)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.13+-blue.svg)](https://golang.google.cn)

stub 让你在 Go 语言单元测试中轻松打桩。

- [安装](#%e5%ae%89%e8%a3%85)
- [导入](#%e5%af%bc%e5%85%a5)
- [使用方法](#%e4%bd%bf%e7%94%a8%e6%96%b9%e6%b3%95)
  - [对变量打桩](#%e5%af%b9%e5%8f%98%e9%87%8f%e6%89%93%e6%a1%a9)
  - [对函数打桩](#%e5%af%b9%e5%87%bd%e6%95%b0%e6%89%93%e6%a1%a9)
  - [对环境变量打桩](#%e5%af%b9%e7%8e%af%e5%a2%83%e5%8f%98%e9%87%8f%e6%89%93%e6%a1%a9)
  - [小技巧](#%e5%b0%8f%e6%8a%80%e5%b7%a7)
- [注意事项](#%e6%b3%a8%e6%84%8f%e4%ba%8b%e9%a1%b9)

## 安装

```shell
$ go get -u -v github.com/aQuaYi/stub
$
```

## 导入

```go
import (
	"github.com/aQuaYi/stub"
)
```

## 使用方法

打桩（stub）和模拟（mock）是单元测试中常用的两种技术。[Martin Fowler](https://martinfowler.com/) 写了 [Mocks Aren't Stubs](https://martinfowler.com/articles/mocksArentStubs.html) 来详细说明两者的定义和区别。以下是文章的两个版本的翻译：

- [Mock 不是 Stub - 众成翻译](https://www.zcfy.cc/article/mocks-arent-stubs)
- [Mock并非Stub（翻译）](http://www.predatorray.me/Mock%E5%B9%B6%E9%9D%9EStub-%E7%BF%BB%E8%AF%91/)

### 对变量打桩

```go
var configFile = "config.json"

func GetConfig() ([]byte, error) {
    return ioutil.ReadFile(configFile)
}

// Test code
stubs := stub.Var(&configFile, "/tmp/test.config")
data, err := GetConfig()
// data will now return contents of the /tmp/test.config file
```

### 对函数打桩

函数在 Go 语言中也是一种变量。因此，可以对全局函数进行打桩。

```go
var timeNow = time.Now

func GetDate() int {
    return timeNow().Day()
}
```

通过 stub 替换 `timeNow` 函数。

```go
stubs := stub.Var(&timeNow, func() time.Time {
    return time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC)
})
defer stubs.Restore()

// Test can check that GetDate returns 6
```

如果替换后的函数，每次都返回固定的结果。还可以使用 `Func` 替换 `Var`。

```go
stubs := stub.Func(&timeNow, time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC))
defer stubs.Restore()
```

`Func` 同样也可以对多返回值的函数打桩。

```go
var osHostname = osHostname
// [...] production code using osHostname to call it.

// Test code:
stubs := stub.Func(&osHostname, "fakehost", nil)
defer stubs.Restore()
```

### 对环境变量打桩

`Env` 可以对环境变量打桩，并在事后轻易地恢复。

```go
stubs := stub.Env("GOSTUB_VAR", "test_value")
defer stubs.Restore()
```

### 小技巧

`Stubs` 对象能够被反复使用，并用 `Restore` 方法一次恢复全部桩。

```go
stubs := stub.Var(&v1, 1)
stubs.Var(&v2, 2)
defer stubs.Restore()
```

对于简单的打桩，还可以在一行内完成打桩和恢复工作。

```go
defer stub.Var(&v1, 1).Var(&v2, 2).Restore()
```

同一个 `stubs` 对象，可以进行多次打桩，并只需要一次 `Restore`。

```go
stubs := stub.Var(&v1, 1)
defer stubs.Restore()

// Do some testing
stubs.Var(&v1, 5)

// More testing
stubs.Var(&b2, 6)
```

## 注意事项

打桩时会记录原始值，并在 `Restore` 时将其恢复。所以，如果在打桩期间，对变量进行的修改，会丢失。
