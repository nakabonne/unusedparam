package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/nakabonne/unusedparam/pkg/unusedparam"
)

var (
	flagSet = flag.NewFlagSet("unusedparam", flag.ContinueOnError)
	verbose = flagSet.Bool("v", false, "Verbose output")
	outJSON = flagSet.Bool("json", false, "Emit json format")
)

func main() {
	flagSet.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: unusedparam [flags] [files ...]")
		flagSet.PrintDefaults()
	}
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		if err != flag.ErrHelp {
			fmt.Fprintln(os.Stderr, err)
		}
		return
	}

	var issues []*unusedparam.Issue
	for _, path := range flagSet.Args() {
		i, err := unusedparam.Check(path)
		if err != nil {
			// TODO: Use debug mode.
			if *verbose {
				fmt.Println(err)
			}
		}
		issues = append(issues, i...)
	}

	if *outJSON {
		js, err := json.Marshal(issues)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(js))
		return
	}
	for _, issue := range issues {
		fmt.Println(issue)
	}
}
