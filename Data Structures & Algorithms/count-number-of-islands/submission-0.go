
func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    
    rows, cols := len(grid), len(grid[0])
    visited := make([][]bool, rows)
    for i := range visited {
        visited[i] = make([]bool, cols)
    }
    
    directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    islands := 0
    
    bfs := func(sr, sc int) {
        queue := [][2]int{{sr, sc}}
        visited[sr][sc] = true
        
        for len(queue) > 0 {
            r, c := queue[0][0], queue[0][1]
            queue = queue[1:]
            
            for _, dir := range directions {
                nr, nc := r+dir[0], c+dir[1]
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols && 
                   grid[nr][nc] == '1' && !visited[nr][nc] {
                    visited[nr][nc] = true
                    queue = append(queue, [2]int{nr, nc})
                }
            }
        }
    }
    
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == '1' && !visited[i][j] {
                islands++
                bfs(i, j)
            }
        }
    }
    
    return islands
}