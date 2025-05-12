package internal

import (
	"math/rand"
)

// use simple random to get the courier
func getCourierFromProbability(probability map[string]int) string {
	counter := 0
	for _, prob := range probability {
		if prob > 0 {
			counter += prob
		}
	}
	randomNum := rand.Intn(counter)
	for courier, prob := range probability {
		if prob > 0 {
			randomNum -= prob
			if randomNum < 0 {
				return courier
			}
		}
	}
	return ""
}
