package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	rev := ""

	for _, r := range word {
		rev = string(r) + rev
	}

	return rev
}
