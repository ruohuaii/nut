package nut

type stringArray []string

func (s stringArray) Contains(target string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == target {
			return true
		}
	}

	return false
}
