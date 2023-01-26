func min(x, y, z int) int {
	if x <= y && x <= z {
		return x
	} else if y <= x && y <= z {
		return y
	} else {
		return z
	}
}

func minDistance(s1, s2 string) int {
	m := len(s1)
	n := len(s2)
	d := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		d[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		d[i][0] = i
	}

	for j := 0; j <= n; j++ {
		d[0][j] = j
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if s1[i-1] == s2[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = min(d[i-1][j], d[i][j-1], d[i-1][j-1]) + 1
			}
		}
	}
	return d[m][n]
}
