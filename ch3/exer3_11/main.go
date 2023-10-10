package main

import "strings"

func comma(s string) string {
	sign, dec, frac := split(s)

	if len(dec) <= 3 {
		return sign + s + frac
	}

	s = dec
	dec = s[len(s)-3:]
	s = s[:len(s)-3]
	for i := len(s); i > 0; i -= 3 {
		if i-3 < 0 {
			dec = s[:i] + "," + dec
			break
		}
		dec = s[i-3:i] + "," + dec
	}
	return sign + dec + frac
}

func split(s string) (sign, dec, frac string) {
	dec = s
	if s[0] == '+' || s[0] == '-' {
		sign = string(s[0])
		dec = s[1:]
	}

	i := strings.LastIndex(dec, ".")
	if i != -1 {
		frac = dec[i:]
		dec = dec[:i]
	}
	return
}
