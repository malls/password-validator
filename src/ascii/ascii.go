package ascii

//Invalid checks whether a given string is made of entirely valid ASCII characters
func Invalid(str string) bool {
	for _, char := range str {

		//if it's not one of the original 128 or *, it isn't allowed
		if char > 127 || char == 42 {
			return true
		}

	}
	return false
}
