package iteration

const repeatCount = 5

// Repeat takes in a character and repeats it x times
func Repeat(character string, times int) string {
	var repeated string

	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
