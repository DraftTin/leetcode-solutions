class Solution {
public:
    int maxAreaOfIsland(vector<vector<int>>& grid) {
        int ans = 0;
        vector<vector<bool>> visited(grid.size(), vector<bool>(grid[0].size(), false));
        for (int i = 0; i < grid.size(); i++) {
            for(int j = 0; j < grid[0].size(); j++) {
                ans = max(ans, getArea(grid, i, j, visited));
            }
        }
        return ans;
    }
    
    int getArea(vector<vector<int>>& grid, int row, int col, vector<vector<bool>>& visited) {
        if (row < 0 || row > grid.size() - 1 || col < 0 || col > grid[0].size() - 1 || visited[row][col] || grid[row][col] == 0) {
            return 0;
        }
        visited[row][col] = true;
        return 1 + getArea(grid, row + 1, col, visited) + getArea(grid, row - 1, col, visited) 
                + getArea(grid, row, col + 1, visited) + getArea(grid, row, col - 1, visited);
    }
};
