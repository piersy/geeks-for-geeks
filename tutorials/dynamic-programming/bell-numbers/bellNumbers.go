package main

func main() {
	set := make([]int, 2)
	println(partitions(set))
	set = make([]int, 3)
	println(partitions(set))
	set = make([]int, 4)
	println(partitions(set))
	set = make([]int, 5)
	println(partitions(set))
	set = make([]int, 6)
	println(partitions(set))
}

func partitions(set []int) int {
	// count := 1
	// for i := 1; i < len(set); i++ {
	// 	count += partitions(set[:i]) * partitions(set[i:])
	// }
	// return count
	count := 0
	for i := 1; i < len(set); i++ {
		count += 1 + partitions(append(set[:i], set[i+1:]...))

	}
	return count
}
