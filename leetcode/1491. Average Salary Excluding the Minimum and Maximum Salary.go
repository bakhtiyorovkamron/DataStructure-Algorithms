package leetcode

func Average(salary []int) float64 {
	min, max := salary[0], salary[0]
	amount := 0
	for i := 0; i < len(salary); i++ {
		amount += salary[i]
		if min < salary[i] {
			min = salary[i]
		}
		if max > salary[i] {
			max = salary[i]
		}
	}
	return float64(float64((amount - min - max)) / float64((len(salary) - 2)))
}
