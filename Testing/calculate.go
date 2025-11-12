package Testing

func Calculate(input int) string {
	if input%3 == 0 && input != 0 {
		return "Foo"
	}
	return string(rune('0' + input))
}
