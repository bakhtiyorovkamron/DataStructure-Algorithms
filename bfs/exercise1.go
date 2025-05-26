package bfs

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func LevelOrder(root *Tree) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*Tree{root} // очередь для BFS

	for len(queue) > 0 {
		levelSize := len(queue)
		var level []int

		for i := 0; i < levelSize; i++ {
			node := queue[0]  // достаём первый элемент из очереди
			queue = queue[1:] // удаляем его из очереди
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left) // добавляем левого потомка
			}
			if node.Right != nil {
				queue = append(queue, node.Right) // добавляем правого потомка
			}
		}

		result = append(result, level)
	}

	return result
}
