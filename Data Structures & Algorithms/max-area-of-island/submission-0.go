func maxAreaOfIsland(grid [][]int) int {
    if len(grid) == 0 {
        return 0
    }
    
    rows, cols := len(grid), len(grid[0])
    visited := make([][]bool, rows)
    for i := range visited {
        visited[i] = make([]bool, cols)
    }
    
    directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    var bfs func(sr, sc, mx int) int
    bfs = func(sr, sc, mx int) int {
        queue := [][2]int{{sr, sc}}
        visited[sr][sc] = true
        curNum := 0
        for len(queue) > 0 {
            r, c := queue[0][0], queue[0][1]
            queue = queue[1:]
            curNum++
            for _, dir := range directions {
                nr, nc := r+dir[0], c+dir[1]
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols && 
                   grid[nr][nc] == 1 && !visited[nr][nc] {
                    visited[nr][nc] = true
                    queue = append(queue, [2]int{nr, nc})
                }
            }
        }
		return curNum
    }
    mx := 0
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 1 && !visited[i][j] {
                num := bfs(i, j, mx)
				if num > mx {
					mx = num
				}
            }
        }
    }
    
    return mx
}
