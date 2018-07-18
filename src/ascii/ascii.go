package ascii

//Invalid checks whether a given string is made of entirely valid ASCII characters
func Invalid(str string) bool {
	for _, char := range str {

		//if it's *, or not a valid ASCII character, it isn't allowed
		if char > 126 || char == 42 || char < 32 {
			return true
		}

	}
	return false
}
