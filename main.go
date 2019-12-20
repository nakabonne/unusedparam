package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nakabonne/unusedparam/pkg/unusedparam"
)

var (
	flagSet = flag.NewFlagSet("unusedparam", flag.ContinueOnError)
	verbose = flagSet.Bool("v", false, "verbose output")
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
	for _, issue := range issues {
		fmt.Println(issue)
	}
}
