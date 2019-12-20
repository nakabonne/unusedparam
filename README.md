# unusedparam

`unusedparam` inspects just go source files and detect unused function parameters. All you need is files, no need to make preparations anything such as code generation, downloading modules. It means it's easy to run in CI environments.  
This is a Go static analysis tool inspired by [mvdan/unparam](https://github.com/mvdan/unparam).

# Installation
```
go get -u github.com/nakabonne/unusedparam
```

# Usage

```
usage: unusedparam [flags] [files ...]
  -v    verbose output
```

You can see how it does by using files underneath `/testdata`.
```
unusedparam ./testdata/*
```

  
    

With `Check()`, you can use in your code.

```go
import (
	"github.com/k0kubun/pp"
	"github.com/nakabonne/unusedparam/pkg/unusedparam"
)

func main() {
	issue, _ := unusedparam.Check("./testdata/assign_stmt.go")
	pp.Println(issue)
}
```

structured issue:

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
