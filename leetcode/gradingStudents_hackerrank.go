package leetcode

func GradingStudents(grades []int32) []int32 {
	// Write your code here
	success := []int32{}
	for i := range grades {
		if grades[i] >= 38 {
			if (grades[i]+2)%5 == 0 {
				success = append(success, grades[i]+2)
			} else {
				success = append(success, grades[i]+2)
			}
		} else {
			success = append(success, grades[i])
		}
	}

	return success
}
