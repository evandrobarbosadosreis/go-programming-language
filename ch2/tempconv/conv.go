package tempconv

func CToF(c Celcius) Fahreinheit { return Fahreinheit(c*9/5 + 32) }
func FToC(f Fahreinheit) Celcius { return Celcius((f - 32) * 5 / 9) }
