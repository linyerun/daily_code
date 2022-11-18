package main

func allPathsSourceTarget(g [][]int) [][]int {
	var ans [][]int
	var dfs func(i int, path []int)
	dfs = func(i int, path []int) {
		if i == len(g)-1 {
			t := make([]int, len(path))
			copy(t, path)
			ans = append(ans, t)
			return
		}
		if len(g[i]) == 0 {
			return
		}
		for _, node := range g[i] {
			path = append(path, node)
			dfs(node, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{0})
	return ans
}

func main() {
	allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}})
}
