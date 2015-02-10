package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Usage: yamlcheck <file>")
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", flag.Arg(0), err)
		os.Exit(2)
	}

	var out interface{}
	err = yaml.Unmarshal(bytes, &out)
	if err != nil {
		fmt.Printf("YAML check failed: %v\n", err)
		os.Exit(3)
	}

	fmt.Println("YAML check ok")
}
