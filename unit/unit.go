package unit

func Percentage(p float32) float32 {
	if p > 1 {
		p /= 100
	}
	return p
}
