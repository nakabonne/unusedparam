package main

import (
	"fmt"

	"github.com/nakabonne/unusedparam/pkg/unusedparam"
)

func main() {
	issues, err := unusedparam.Check("./testdata/1.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, issue := range issues {
		fmt.Println(issue)
	}
}
