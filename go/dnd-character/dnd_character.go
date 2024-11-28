package dndcharacter

import (
	"math"
	"math/rand/v2"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	return rand.IntN(18-3) + 3
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	str := Ability()
	dex := Ability()
	con := Ability()
	inte := Ability()
	wis := Ability()
	cha := Ability()
	hit := 10 + Modifier(con)

	return Character{str, dex, con, inte, wis, cha, hit}
}
