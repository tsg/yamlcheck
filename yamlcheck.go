package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {

	quiet := flag.Bool("q", false, "Don't print anything to stderr")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: yamlcheck [options] <file>\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		if !*quiet {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", flag.Arg(0), err)
		}
		os.Exit(2)
	}

	var out interface{}
	err = yaml.Unmarshal(bytes, &out)
	if err != nil {
		if !*quiet {
			fmt.Fprintf(os.Stderr, "YAML check failed: %v\n", err)
		}
		os.Exit(3)
	}

	if !*quiet {
		fmt.Fprintf(os.Stderr, "YAML check ok\n")
	}
}
