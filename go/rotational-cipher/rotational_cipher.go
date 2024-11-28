package rotationalcipher

import "strings"

func RotationalCipher(plain string, shiftKey int) string {
	var out strings.Builder

	for i := range plain {
		char := plain[i]
		switch {

		case char >= 'a' && char <= 'z':
			out.WriteByte((char-'a'+byte(shiftKey))%26 + 'a')
		case char >= 'A' && char <= 'Z':
			out.WriteByte((char-'A'+byte(shiftKey))%26 + 'A')
		default:
			out.WriteByte(char)
		}
	}

	return out.String()
}
