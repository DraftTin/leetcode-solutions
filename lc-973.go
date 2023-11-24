func kClosest(points [][]int, k int) [][]int {
    distances := make([][2]int, len(points))
    for i := 0; i < len(distances); i++ {
        distances[i][0] = points[i][0] * points[i][0] + points[i][1] * points[i][1]
        distances[i][1] = i
    }
    sort.Slice(distances, func(i, j int) bool {
        return distances[i][0] > distances[j][0]
    })
    ans := make([][]int, k)
    for i := 0; i < k; i++ {
        index := distances[i][1]
        ans[i] = []int{points[index]...}
    }
    return ans
}
