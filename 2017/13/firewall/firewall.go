package firewall

func DetectorPosition(size, t int) int {
	patternLength := 2 * (size - 1)
	normalisedT := t % patternLength
	if normalisedT < size {
		return normalisedT
	}
	return patternLength - normalisedT
}

func Severity(config map[int]int, delay int) int {
	sev := 0
	for k, v := range config {
		if DetectorPosition(v, k+delay) == 0 {
			sev += k * v
		}
	}
	return sev
}

func Caught(config map[int]int, delay int) bool {
	for k, v := range config {
		if DetectorPosition(v, k+delay) == 0 {
			return true
		}
	}
	return false
}

func MinDelay(config map[int]int) int {
	d := 0
	for Caught(config, d) {
		d++
	}
	return d
}
