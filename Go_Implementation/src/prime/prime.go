package prime

//Prime : Tests whether number is prime or not.
func Prime(n int) bool {
	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}
	if n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	if n%3 == 0 {
		return false
	}
	i, w := 5, 2

	for i*i <= n {
		if n%i == 0 {
			return false
		}
		i = i + w
		w = 6 - w
	}
	return true
}
