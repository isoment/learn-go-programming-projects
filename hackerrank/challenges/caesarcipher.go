package challenges

func CaesarCipher(s string, k byte) string {
	// Capitals 65-90
	// Lowercase 97-122

	byteString := make([]byte, len(s))

	for i, v := range s {
		// Ascii value of the character
		ascii := byte(v)
		// Shift the character by k % 26 since there are 26 chars in the ascii alphabet
		shift := ascii + byte(k%26)
		// The new character
		var n byte

		if ascii >= 65 && ascii <= 90 {
			// Capital
			if shift > 90 {
				n = shift - 90 + 65 - 1
			} else {
				n = shift
			}
			byteString[i] = n
		} else if ascii >= 97 && ascii <= 122 {
			// Lowercase
			if shift > 122 {
				n = shift - 122 + 97 - 1
			} else {
				n = shift
			}
			byteString[i] = n
		} else {
			// Special character
			byteString[i] = ascii
		}
	}

	return string(byteString)
}
