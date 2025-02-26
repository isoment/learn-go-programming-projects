package challenges

import (
	"testing"
)

func TestCaesarCipher(t *testing.T) {
	type Case struct {
		name     string
		s        string
		k        byte
		expected string
	}

	tests := []Case{
		{"case one", "middle-Outz", 2, "okffng-Qwvb"},
		{"case two", "Always-Look-on-the-Bright-Side-of-Life", 5, "Fqbfdx-Qttp-ts-ymj-Gwnlmy-Xnij-tk-Qnkj"},
		{"case three", "There's-a-starman-waiting-in-the-sky", 3, "Wkhuh'v-d-vwdupdq-zdlwlqj-lq-wkh-vnb"},
		{"case four", "Hello_World!", 4, "Lipps_Asvph!"},
		{"case five", "Hello_World!", 0, "Hello_World!"},
		{"case six", "Hello_World!", 70, "Zwddg_Ogjdv!"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := CaesarCipher(tc.s, tc.k)
			if actual != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, actual)
			}
		})
	}
}
