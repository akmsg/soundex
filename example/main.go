package main

import (
	"fmt"
	"github.com/akmsg/soundex/soundexcore"
)

func main() {
	fmt.Println(soundexcore.GetCode("Robert"))
	fmt.Println(soundexcore.GetCode("Rupert"))
}
