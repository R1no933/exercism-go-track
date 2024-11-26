package cipher

import (
	"strings"
)

type shift struct {
	dstnc int
}
type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return NewShift(3)
}
func NewShift(distance int) Cipher {
	if distance < -25 || distance > 25 || distance == 0 {
		return nil
	}
	return shift{
		dstnc: distance,
	}
}
func shiftRune(r rune, dist int) rune {
	if dist < 0 {
		dist += 26
	}
	num := (int(r-'a') + dist) % 26
	return rune('a' + num)
}
func (c shift) Encode(input string) string {
	input = strings.ToLower(input)
	var result string
	for _, r := range input {
		if 'a' <= r && r <= 'z' {
			result += string(shiftRune(r, c.dstnc))
		}
	}
	return result
}
func (c shift) Decode(input string) string {
	var result string
	for _, r := range input {
		result += string(shiftRune(r, -c.dstnc))
	}
	return result
}
func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}
	valid := false
	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}
		// key should contain letters other than a
		if r != 'a' {
			valid = true
		}
	}
	if !valid {
		return nil
	}
	return vigenere{
		key: key,
	}
}
func (v vigenere) Encode(input string) string {
	input = strings.ToLower(input)
	var result string
	count := 0
	for _, r := range input {
		dist := int(v.key[count%len(v.key)] - 'a')
		if 'a' <= r && r <= 'z' {
			result += string(shiftRune(r, dist))
			count++
		}
	}
	return result
}
func (v vigenere) Decode(input string) string {
	input = strings.ToLower(input)
	var result string
	for i, r := range input {
		dist := int(v.key[i%len(v.key)] - 'a')
		if 'a' <= r && r <= 'z' {
			result += string(shiftRune(r, -dist))
		}
	}
	return result
}
