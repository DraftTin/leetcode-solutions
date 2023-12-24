class Solution {
public:
    int minOperations(string s) {
        string s1(s.size(), '0'), s2(s.size(), '1');
        for(int i = 0; i < s.size(); i += 2) {
            s1[i] = '1';
            s2[i] = '0';
        }
        int dist1 = 0, dist2 = 0;
        for(int i = 0; i < s.size(); i++) {
            dist1 += (s1[i] - s[i]) * (s1[i] -s[i]);

            dist2 += (s2[i] - s[i]) * (s2[i] -s[i]);
        }
        return min(dist1, dist2);
    }
};
