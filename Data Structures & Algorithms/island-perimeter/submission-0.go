func islandPerimeter(grid [][]int) int {
   rows, cols := len(grid), len(grid[0])
   visited := make(map[[2]int]bool)

   var dfs func(i, j int) int 
   dfs = func(i, j int) int {
	if i<0 || j<0|| i>= rows|| j>=cols|| grid[i][j] == 0 {
		return 1
	}
	if visited[[2]int{i, j}] {
		return 0
	}

	visited[[2]int{i, j}]  =true
	return dfs(i, j+1)+dfs(i, j-1)+dfs(i+1,j)+dfs(i-1,j)
   }

   for i:=0;i< rows; i++{
	for j:=0; j<cols;j++ {
		if grid[i][j] ==1 {
			return dfs(i, j)
		}
	}
   }
   return 0
}
