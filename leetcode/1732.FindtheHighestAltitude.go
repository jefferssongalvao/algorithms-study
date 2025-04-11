package main

func largestAltitude(gain []int) int {
	altitude, maxAltitude := 0, 0
	for _, n := range gain {
		altitude += n
		maxAltitude = max(maxAltitude, altitude)
	}
	return maxAltitude
}
