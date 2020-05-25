package soundexcore

import (
	"bytes"
	"fmt"
	"strings"
)

var errEmptyString = fmt.Errorf("cannot compute soundex code for empty string")

// GetCode computes the American Soundex code for given string input s
// Soundex is a phonetic algorithm for indexing names by sound, as pronounced in English.
func GetCode(s string) (string, error) {
	b := bytes.Buffer{}
	b.Grow(len(s) + 1)

	r := bytes.Buffer{}
	r.Grow(1)

	if 1 > len(s) {
		return "", errEmptyString
	}

	// normalise the case for input s
	s = strings.ToUpper(s)

	// 1. retain first char of the s as first char of code
	r.WriteByte(s[0])

	s = s[1:]

	// 2. encode all the chars as per the code in the mapper
	for _, c := range s {
		code, ok := mapper[byte(c)]
		if !ok {
			continue
		}
		b.WriteByte(code)
	}

	// 3. eradicate all the same adjacent digits and drop all the 0s
	t := removeDuplicates(b.Bytes())
	b.Reset()
	b.Write(t)

	t = removeCode(b.Bytes(), '0')
	b.Reset()
	b.Write(t)

	// 4. create a new buffer with first char and add padding
	r.Grow(b.Len() + lenSoundexCode)
	r.Write(b.Bytes())

	// 5. add padding
	for i := 0; i < lenSoundexCode; i++ {
		r.WriteByte('0')
	}

	// 6. return first n chars
	return r.String()[:lenSoundexCode], nil
}

func removeDuplicates(b []byte) []byte {
	r := bytes.Buffer{}
	r.Grow(len(b) + 1)

	if 1 > len(b) {
		return r.Bytes()
	}

	itr1 := 0
	itr2 := 1

	r.WriteByte(b[itr1])
	for itr2 < len(b) {
		if b[itr1] != b[itr2] {
			r.WriteByte(b[itr2])
		}
		itr1++
		itr2++
	}
	return r.Bytes()
}

func removeCode(b []byte, code byte) []byte {
	r := bytes.Buffer{}
	r.Grow(len(b) + 1)

	if 1 > len(b) {
		return r.Bytes()
	}

	i := 0

	for i < len(b) {
		if b[i] != code {
			r.WriteByte(b[i])
		}
		i++
	}
	return r.Bytes()
}
