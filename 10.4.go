package main

//find the average of a series of numbers
func Average(numbers ...floa64) floa64 {
	var total floa64
	for _, v := range numbers {
		total = total + v
	}
	return total / floa64(len(numbers))
}