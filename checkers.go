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

func checkRotate(length int) float64 {
	if length == 54 || length == 72 || length == 110 {
		return 2 * 0.0277778
	}
	return 0
}
