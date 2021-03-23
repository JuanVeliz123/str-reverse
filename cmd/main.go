package main

import (
	"fmt"
	"github.com/str-reverse/internal"
)

var CASES_CSV_PATH = "test/cases.csv"

func main() {
	cases := internal.ReadCases(&CASES_CSV_PATH)

	for _, c := range cases {
		result := internal.ReverseParagraph(&c)
		fmt.Printf(fmt.Sprintf("Case:\n%s\n\nResult:\n%s\n\n", c, *result))
	}
}
