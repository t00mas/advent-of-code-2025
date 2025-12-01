package main

const Size = 100

func Rotate(s int, i Instruction) int {
	if i.Direction == "L" {
		s -= i.Steps
	} else {
		s += i.Steps
	}
	return (s%Size + Size) % Size
}

func RotateCountingZeroes(s int, i Instruction) (int, int) {
	pos, zeroes, step := s, 0, 1
	if i.Direction == "L" {
		step = -1
	}

	for n := 0; n < i.Steps; n++ {
		pos = (pos + step) % Size
		if pos < 0 {
			pos += Size
		}
		if pos == 0 {
			zeroes++
		}
	}

	return pos, zeroes
}
