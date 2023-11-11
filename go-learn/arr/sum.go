package arr
//import "fmt"

func Sum(numbers [5]int) int {

	var sum = 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	
	return sum
}

