func isScramble(s1 string, s2 string) bool {
    if s1 == s2 {
        return true
    }
    if len(s1) != len(s2) {
        return false
    }
    n: = len(s1)
    dp: = make([][][] bool, n)
    for i: = range dp {
        dp[i] = make([][] bool, n)
        for j: = range dp[i] {
            dp[i][j] = make([] bool, n + 1)
        }
    }
    for i: = 0;
    i < n;
    i++{
        for j: = 0;
        j < n;
        j++{
            dp[i][j][1] = s1[i] == s2[j]
        }
    }
    for l: = 2;
    l <= n;
    l++{
        for i: = 0;
        i + l <= n;
        i++{
            for j: = 0;
            j + l <= n;
            j++{
                for k: = 1;
                k < l;
                k++{
                    if (dp[i][j][k] && dp[i + k][j + k][l - k]) || (dp[i][j + l - k][k] && dp[i + k][j][l - k]) {
                        dp[i][j][l] = true
                        break
                    }
                }
            }
        }
    }
    return dp[0][0][n]
}
