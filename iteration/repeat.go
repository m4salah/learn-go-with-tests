package iteration

func Repeat(character string, times uint) string {
	repeated := ""
	for i := 0; i < int(times); i++ {
		repeated += character
	}
	return repeated
}
