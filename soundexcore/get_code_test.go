package soundexcore

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGetCode will test GetCode which implements get soundex code
func TestGetCode(t *testing.T) {
	type testSuite struct {
		tag          string
		givenS       string
		expectedCode string
		expectedErr  error
	}

	tests := []testSuite{
		{
			tag:          "tag 0: empty string",
			givenS:       "",
			expectedCode: "",
			expectedErr:  errEmptyString,
		},
		{
			tag:          "tag 1: non string",
			givenS:       "1",
			expectedCode: "1000",
			expectedErr:  nil,
		},
		{
			tag:          "tag 2: non string",
			givenS:       "a",
			expectedCode: "A000",
			expectedErr:  nil,
		},
		{
			tag:          "tag 3: string",
			givenS:       "d",
			expectedCode: "D000",
			expectedErr:  nil,
		},
		{
			tag:          "tag 4: string",
			givenS:       "Robert",
			expectedCode: "R163",
			expectedErr:  nil,
		},
		{
			tag:          "tag 5: string",
			givenS:       "Rupert",
			expectedCode: "R163",
			expectedErr:  nil,
		},
		{
			tag:          "tag 6: string",
			givenS:       "Tymczak",
			expectedCode: "T522",
			expectedErr:  nil,
		},
		{
			tag:          "tag 7: string",
			givenS:       "Honeyman",
			expectedCode: "H555",
			expectedErr:  nil,
		},
	}

	for _, test := range tests {
		t.Logf("running: %s\n", test.tag)
		outputCode, outputErr := GetCode(test.givenS)
		require.Equal(t, test.expectedCode, outputCode, test.tag)
		require.Equal(t, test.expectedErr, outputErr, test.tag)
	}
}

func Test_removeDuplicates(t *testing.T) {
	type testSuite struct {
		tag      string
		given    []byte
		expected []byte
	}
	tests := []testSuite{
		{
			tag:      "tag 0",
			given:    []byte{},
			expected: []byte{},
		},
		{
			tag:      "tag 1",
			given:    []byte{'a'},
			expected: []byte{'a'},
		},
		{
			tag:      "tag 2",
			given:    []byte{'a', 'b'},
			expected: []byte{'a', 'b'},
		},
		{
			tag:      "tag 3",
			given:    []byte{'a', 'a'},
			expected: []byte{'a'},
		},
		{
			tag:      "tag 4",
			given:    []byte{'a', 'a', 'b'},
			expected: []byte{'a', 'b'},
		},
		{
			tag:      "tag 5",
			given:    []byte{'a', 'a', 'b', 'a', 'a'},
			expected: []byte{'a', 'b', 'a'},
		},
		{
			tag:      "tag 6",
			given:    []byte{'a', 'a', 'b', 'b', 'c'},
			expected: []byte{'a', 'b', 'c'},
		},
		{
			tag:      "tag 7",
			given:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected: []byte{'a', 'b', 'c'},
		},
	}
	for _, test := range tests {
		output := removeDuplicates(test.given)
		require.Equal(t, test.expected, output, test.tag)
	}
}

func Test_removeCode(t *testing.T) {
	type testSuite struct {
		tag      string
		given    []byte
		expected []byte
	}
	tests := []testSuite{
		{
			tag:      "tag 0",
			given:    []byte{},
			expected: []byte{},
		},
		{
			tag:      "tag 1",
			given:    []byte{'0', '0'},
			expected: []byte{},
		},
		{
			tag:      "tag 2",
			given:    []byte{'a', '0', 'b'},
			expected: []byte{'a', 'b'},
		},
		{
			tag:      "tag 3",
			given:    []byte{'a', '0', 'b', '0', 'c'},
			expected: []byte{'a', 'b', 'c'},
		},
		{
			tag:      "tag 4",
			given:    []byte{'a', 'a', '0', 'b'},
			expected: []byte{'a', 'a', 'b'},
		},
	}
	for _, test := range tests {
		output := removeCode(test.given, '0')
		require.Equal(t, test.expected, output, test.tag)
	}
}
