package main


func main() {
	// entry point
}
func Sum(numbers []int) (result int ) {
	result = 0
	for _, number := range numbers {
		result+=number
	}
	return  
}
func SumAll(numbersToSum ...[]int) []int {
/* 
	sums := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		sums[i]=Sum(numbers)
	}

	 */

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return  sums
}

func SumAllTails(numbersToSum ...[]int) []int {

	var sums []int
	for _, numbers := range numbersToSum {
		if(len(numbers) == 0) {
			sums = append(sums, 0)
		}else {

			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return  sums
}
