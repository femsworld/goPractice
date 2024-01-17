package sprint

func ShiftBy(r rune, step int) rune {
	letterIndex := int(r - 'a')
	shiftedIndex := (letterIndex + step) % 26
	shiftedLetter := rune('a' + shiftedIndex)
	return shiftedLetter
}