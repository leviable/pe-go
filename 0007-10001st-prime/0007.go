package main

func IsPrime(candidate int) bool {
	switch candidate {
	case 1:
		return false
	case 2:
		return true
	case 3:
		return true
	case 5:
		return true
	default:
		for i := 2; i < candidate; i++ {
			if candidate%i == 0 {
				return false
			}
		}
	}
	return true
}
