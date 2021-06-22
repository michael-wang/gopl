package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToK(c Celsius) Kelvin     { return Kelvin(c) }
func KToC(k Kelvin) Celsius     { return Celsius(k) }
func FToK(f Fahrenheit) Kelvin  { return Kelvin((f - 32) * 5 / 9) }
func KToF(k Kelvin) Fahrenheit  { return Fahrenheit(k*9/5 + 32) }
