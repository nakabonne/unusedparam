# unusedparam

[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/nakabonne/unusedparam/pkg/unusedparam)

A Go static analysis tool to inspect go source files and detect unused function parameters.  

It doesn't require to preload Go packages, all you need is files and no need to make preparations anything such as code generation, downloading modules. It allows you to run in CI environments easily even if your project depends on a complicated build system.


## Installation
```
go get -u github.com/nakabonne/unusedparam
```

## Usage

```
usage: unusedparam [flags] [files ...]
  -json Emit json format
  -v	Verbose output
```

You can see how it does by using files underneath `/testdata`.
```
unusedparam ./testdata/*
```

### Using as a library

With `Check()`, you can handle the structured issues in your code.

```go
import "github.com/nakabonne/unusedparam/pkg/unusedparam"

func main() {
	issues, _ := unusedparam.Check("./testdata/assign_stmt.go")
}
```

structured issues:

```go
[]*unusedparam.Issue{
  &unusedparam.Issue{
    Text: "m is unused in _",
    Pos:  token.Position{
      Filename: "./testdata/assign_stmt.go",
      Offset:   176,
      Line:     18,
      Column:   11,
    },
  },
  &unusedparam.Issue{
    Text: "n is unused in _",
    Pos:  token.Position{
      Filename: "./testdata/assign_stmt.go",
      Offset:   173,
      Line:     18,
      Column:   8,
    },
  },
}
```

## Inspired by

- [mvdan/unparam](https://github.com/mvdan/unparam)
