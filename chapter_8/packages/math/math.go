package math

func Average(xs []float64) float64 { // notice that this is named Average and not average. Starting the name with a capital letter allows us to expose the function to other programs.
	total := float64(0)
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
