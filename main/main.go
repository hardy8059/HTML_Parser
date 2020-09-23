package main

import (
	"fmt"
	project "github.com/hardy8059/HTML_Parser/src"
)

func main() {
	filename := "test2.html"
	o := project.ParserPipeline(filename)
	fmt.Printf("Output: %+v\n", o)
}
