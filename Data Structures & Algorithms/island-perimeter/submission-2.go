//func islandPerimeter(grid [][]int) int {
 //  rows, cols := len(grid), len(grid[0])
//  visited := make(map[[2]int]bool)
//   directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

//    var dfs func(i, j int) int 
//    dfs = func(i, j int) int {
// 	if i<0 || j<0|| i>= rows|| j>=cols|| grid[i][j] == 0 {
// 		return 1
// 	}
// 	if visited[[2]int{i, j}] {
// 		return 0
// 	}

// 	visited[[2]int{i, j}]  =true
// 	return dfs(i, j+1)+dfs(i, j-1)+dfs(i+1,j)+dfs(i-1,j)
//    }

//    for i:=0;i< rows; i++{
// 	for j:=0; j<cols;j++ {
// 		if grid[i][j] ==1 {
// 			return dfs(i, j)
// 		}
// 	}
//    }
//    return 0
  
   

//}

func islandPerimeter(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    visited := make(map[[2]int]bool)
    directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

    bfs := func(r, c int) int {
        q := [][2]int{{r, c}}  
        visited[[2]int{r, c}] = true
        per := 0

        for len(q) > 0 {  
            cell := q[0]
            q = q[1:]
            x, y := cell[0], cell[1]

            for _, dir := range directions {
                nx, ny := x+dir[0], y+dir[1]  
                if nx < 0 || ny < 0 || nx >= rows || ny >= cols || grid[nx][ny] == 0 {  
                    per++
                } else if !visited[[2]int{nx, ny}] {
                    visited[[2]int{nx, ny}] = true
                    q = append(q, [2]int{nx, ny})
                }
            }
        }
        return per
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ { 
            if grid[i][j] == 1 {  
                return bfs(i, j)
            }
        }
    }
    return 0
}  

