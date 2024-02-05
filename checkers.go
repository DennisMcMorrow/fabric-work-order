package main

func checkIfHemNeeded(check bool) float64 {
	if check {
		return 2 * 0.0277778
	}
	return 0
}

func checkIfSurgeNeeded(check bool) float64 {
	if check {
		return 2 * 0.0277778
	}
	return 0
}

func checkRotate(length float64, materialLength float64) bool {
	if length == materialLength {
		return true
	}
	return false
}
