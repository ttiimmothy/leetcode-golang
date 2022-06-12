package main

import "fmt"

func depthFirstSearch(grid [][]byte, row int, col int, visited [][]bool) {
	var gridRow = len(grid)
	var gridCol = len(grid[0])
	if row < 0 || row >= gridRow || col < 0 || col >= gridCol || grid[row][col] == '0' || visited[row][col] {
		return
	}

	visited[row][col] = true
	var rowNeighbor = []int{0, 0, -1, 1}
	var colNeighbor = []int{1, -1, 0, 0}
	for i := 0; i < len(rowNeighbor); i++ {
		depthFirstSearch(grid, row+rowNeighbor[i], col+colNeighbor[i], visited)
	}
}

func numIslands(grid [][]byte) int {
	var count = 0
	var visited = make([][]bool, len(grid))

	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
		for j := 0; j < len(grid[i]); j++ {
			visited[i][j] = false
		}
	}

	for i := 0; i < len(visited); i++ {
		for j := 0; j < len(visited[i]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				depthFirstSearch(grid, i, j, visited)
				count++
			}
		}
	}
	return count
}

func main() {
	var grid = [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(grid))
	var visited = make([][]bool, len(grid))

	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
		for j := 0; j < len(grid[i]); j++ {
			visited[i][j] = false
		}
	}
	for i := 0; i < len(visited); i++ {
		for j := 0; j < len(visited[i]); j++ {
			if !visited[i][j] {
				visited[i][j] = true
			}
		}
	}
	fmt.Println(visited)
	fmt.Println(visited[0][1])
	fmt.Println(len(grid[0]))
}
