package main

import (
	"fmt"
	"github.com/akmsg/soundex/soundexcore"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	for _, arg := range argsWithoutProg{
		soundexCode, err := soundexcore.GetCode(arg)
		if err != nil{
			fmt.Printf("%s: error: %v\n", arg, err)
			continue
		}
		fmt.Printf("%s: %s\n", arg, soundexCode)
	}
}
