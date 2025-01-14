class Solution {
    public int[] xorQueries(int[] arr, int[][] queries) {
        int n = arr.length;
        int[] pre = new int[n];
        int base = 0;
        for (int i = 0; i < arr.length; i++) {
            base = base ^ arr[i];
            pre[i] = base;
        }
        int[] res = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int tmp = pre[queries[i][1]];
            if (queries[i][0] != 0) {
                tmp = tmp ^ pre[queries[i][0] - 1];
            }
            res[i] = tmp;
        }
        return res;
    }
}
