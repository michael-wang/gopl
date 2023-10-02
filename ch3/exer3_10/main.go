package main

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}

	out := s[len(s)-3:]
	s = s[:len(s)-3]
	for i := len(s); i > 0; i -= 3 {
		if i-3 < 0 {
			out = s[:i] + "," + out
			break
		}
		out = s[i-3:i] + "," + out
	}
	return out
}
