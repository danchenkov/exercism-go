package pascal

const testVersion = 1

func Triangle(size int) [][]int {
	triangle := [][]int{}
	for row := 1; row <= size; row++ {
		elements := []int{1}
		for i := 1; i <= row-1; i++ {
			elements = append(elements, elements[i-1]*(row-i)/i)
		}
		triangle = append(triangle, elements)
	}
	return triangle
}
