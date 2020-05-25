package soundexcore

const (
	lenSoundexCode = 4
)

var (

	// map letters to respective soundex codes
	mapper = map[byte]byte{

		// assign 0 to all the vowels along with H, W and Y
		'A': '0', 'E': '0', 'I': '0', 'O': '0', 'U': '0', 'H': '0', 'W': '0', 'Y': '0',

		// assign code 1 to B, F, P, and V
		'B': '1', 'F': '1', 'P': '1', 'V': '1',

		// assign code 2 to C, G, J, K, Q, S, X, and Z
		'C': '2', 'G': '2', 'J': '2', 'K': '2', 'Q': '2', 'S': '2', 'X': '2', 'Z': '2',

		// assign code 3 to D and T
		'D': 3, 'T': '3',

		// assign code 4 to L
		'L': '4',

		// assign code 5 to M and N
		'M': '5', 'N': '5',

		// assign code 6 to R
		'R': '6',
	}
)
