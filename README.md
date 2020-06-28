# soundex [![Build Status](https://travis-ci.org/akmsg/soundex.svg?branch=master)](https://travis-ci.org/akmsg/soundex)
efficient go implementation of [soundex](https://en.wikipedia.org/wiki/Soundex) which is a phonetic algorithm for indexing names by sound, as pronounced in English

## Usage

### Command Line

Install the binary

    go install github.com/akmsg/soundex/soundexcmd

Run the binary and supply command line arguments for showing their soundex code

     soundexcmd robert rupert

Sample response:

     robert: R163
     rupert: R163
     
### Programmatically

Download the module

    go get github.com/akmsg/soundex/soundexcore
    
Import and use, example:

    import (
    	"fmt"
    	"github.com/akmsg/soundex/soundexcore"
    )
    
    func main() {
    	fmt.Println(soundexcore.GetCode("robert"))
    	fmt.Println(soundexcore.GetCode("rupert"))
    }
