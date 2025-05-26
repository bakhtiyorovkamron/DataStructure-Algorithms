package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func AverageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	queue := []*TreeNode{root}
	result := []float64{}
	for len(queue) > 0 {
		levelSize := len(queue)
		var sum float64
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += float64(node.Val)
			left := node.Left
			right := node.Right
			if left != nil {
				queue = append(queue, left)
			}
			if right != nil {
				queue = append(queue, right)
			}
		}
		result = append(result, sum/float64(levelSize))
	}
	return result
}
