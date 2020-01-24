package main

/*
Gold Mine Problem

Given a gold mine of n*m dimensions. Each field in this mine contains a
positive integer which is the amount of gold in tons. Initially the miner is at
first column but can be at any row. He can move only (right->,right up /,right
down\) that is from a given cell, the miner can move to the cell diagonally up
towards the right or right or diagonally down towards the right.  Find out
maximum amount of gold he can collect.

Examples:

Input : mat[][] = {{1, 3, 3},
                   {2, 1, 4},
                  {0, 6, 4}};
Output : 12
{(1,0)->(2,1)->(2,2)}

Input: mat[][] = { {1, 3, 1, 5},
                   {2, 2, 4, 1},
                   {5, 0, 2, 3},
                   {0, 6, 1, 2}};
Output : 16
(2,0) -> (1,1) -> (1,2) -> (0,3) OR
(2,0) -> (3,1) -> (2,2) -> (2,3)

Input : mat[][] = {{10, 33, 13, 15},
                  {22, 21, 04, 1},
                  {5, 0, 2, 3},
                  {0, 6, 14, 2}};
Output : 83

*/

func main() {
	mine := [][]int{{1, 3, 3},
		{2, 1, 4},
		{0, 6, 4}}
	println(maxGold(mine))

	mine = [][]int{{1, 3, 1, 5},
		{2, 2, 4, 1},
		{5, 0, 2, 3},
		{0, 6, 1, 2}}
	println(maxGold(mine))

	mine = [][]int{{10, 33, 13, 15},
		{22, 21, 04, 1},
		{5, 0, 2, 3},
		{0, 6, 14, 2}}
	println(maxGold(mine))
}

func maxGold(mine [][]int) int {
	maFromPoint := makeMultiSlice(len(mine), len(mine[0]))
	// Initialise last column
	for i := 0; i < len(mine); i++ {
		maxFromPoint[i][len(mine[0])-1] = mine[i][len(mine[0])-1]
	}

	for x := len(mine) - 2; x >= 0; x-- {
		for y := 0; y < len(mine[0]); y++ {
			maxFromPoint[y][x] = mine[y][x] + tripleMax(
				getGold(maxFromPoint, x+1, y-1),
				getGold(maxFromPoint, x+1, y),
				getGold(maxFromPoint, x+1, y+1),
			)
		}
	}

	return tripleMax(
		maxFromPoint[0][0],
		maxFromPoint[1][0],
		maxFromPoint[2][0],
	)
}

func getGold(mine [][]int, x, y int) int {
	if y >= len(mine) || y < 0 || x >= len(mine[0]) || x < 0 {
		return 0
	}
	return mine[y][x]
}

func makeMultiSlice(h, w int) [][]int {
	slice := make([][]int, h)
	for i := range slice {
		slice[i] = make([]int, w)
	}
	return slice
}

func tripleMax(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}
