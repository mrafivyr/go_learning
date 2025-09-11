package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	dayCosts := make([]float64, 0, len(costs))
	for _, costData := range costs {
		if costData.day == day {
			dayCosts = append(dayCosts, costData.value)
		}
	}
	return dayCosts
}
