package recursion

// 排位问题
// f(n) = f(n-1) + 1
func num(n int) int {
	if n == 1 {
		return 1
	}

	return num(n-1) + 1
}

// N 台阶问题, 第N步可以从 n-1 来或者 n-2来
// f(n) = f(n-1)+f(n-2)
//
