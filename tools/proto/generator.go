package main

import (
	"flag"
	"fmt"
	"log"

	macro "github.com/tdakkota/gomacro"
	"github.com/tdakkota/gomacro/derive/cursor"
	"github.com/tdakkota/gomacro/runner"
)

func main() {
	flag.Parse()
	if err := run(flag.Arg(0), flag.Arg(1)); err != nil {
		fmt.Println(err)
		return
	}
}

func run(path, output string) error {
	m, err := cursor.Create()
	if err != nil {
		log.Fatal(err)
	}

	r := runner.Runner{
		Source:     path,
		Output:     output,
		AppendMode: true,
	}

	return r.Run(macro.Macros{
		"derive_binary": m,
	})
}
