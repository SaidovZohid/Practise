package masala2

func FindMinimalElements(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}
	var count int
	for i := 0; i < len(arr); i++ {
		if arr[i] == min {
			count++
		}
	}
	return count
}