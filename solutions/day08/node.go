package day08

import "fmt"

type Node struct {
	x int
	y int
}

func (n Node) String() string {
	return fmt.Sprintf("Node(x: %d, y: %d)", n.x, n.y)
}
func (n Node) isInBound(grid [][]byte) bool {
	return n.y >= 0 && n.y < len(grid) && n.x >= 0 && n.x < len(grid[1])
}
func (n Node) getVector(other Node) [2]int {
	return [2]int{other.y - n.y, other.x - n.x}
}
