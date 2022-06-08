package main

import "fmt"

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type Fraction struct {
	a, b int
}

func (f Fraction) simplify() Fraction {
	t := gcd(f.a, f.b)
	f.a /= t
	f.b /= t
	if f.b < 0 {
		f.b *= -1
		f.a *= -1
	}
	return f
}

func fraction(a, b int) Fraction {
	ans := Fraction{a, b}
	ans = ans.simplify()
	return ans
}

func sum(a, b Fraction) Fraction {
	a1 := a.a
	a2 := a.b
	b1 := b.a
	b2 := b.b
	y := abs(a2*b2) / gcd(abs(a2), abs(b2))
	a1 *= y / a2
	b1 *= y / b2
	x := a1 + b1
	ans := Fraction{x, y}
	ans = ans.simplify()
	return ans
}

func diff(a, b Fraction) Fraction {
	a1 := a.a
	a2 := a.b
	b1 := b.a
	b2 := b.b
	y := abs(a2*b2) / gcd(abs(a2), abs(b2))
	a1 *= y / a2
	b1 *= y / b2
	x := a1 - b1
	ans := Fraction{x, y}
	ans = ans.simplify()
	return ans
}

func mult(a, b Fraction) Fraction {
	a.a *= b.a
	a.b *= b.b
	a = a.simplify()
	return a
}

func division(a, b Fraction) Fraction {
	if b.a > 0 {
		return mult(a, Fraction{b.b, b.a})
	}
	return mult(a, Fraction{-b.b, -b.a})
}

func fabs(a Fraction) Fraction {
	return Fraction{abs(a.a), a.b}
}

func gaus(a [][]Fraction, b []Fraction, n int) []Fraction {
	//fmt.Println(b)
	x := make([]Fraction, n, n)
	k := 0
	var index int
	for k < n {
		mux := fabs(a[k][k])
		index = k
		for i := k + 1; i < n; i++ {
			if diff(fabs(a[i][k]), mux).a == 0 {
				mux = fabs(a[i][k])
				index = i
			}
		}
		if mux.a == 0 {
			return x
		}
		for i := 0; i < n; i++ {
			a[k][i], a[index][i] = a[index][i], a[k][i]
		}
		b[k], b[index] = b[index], b[k]
		for i := k; i < n; i++ {
			t := a[i][k]
			if t.a == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				a[i][j] = division(a[i][j], t)
			}
			b[i] = division(b[i], t)
			if i == k {
				continue
			}
			for j := 0; j < n; j++ {
				a[i][j] = diff(a[i][j], a[k][j])
			}
			b[i] = diff(b[i], b[k])
		}
		k++
	}
	//fmt.Println(b)
	for k := n - 1; k >= 0; k-- {
		x[k] = b[k]
		for i := 0; i < k; i++ {
			b[i] = diff(b[i], mult(a[i][k], x[k]))
		}
	}
	return x
}

func main() {
	var n int
	fmt.Scan(&n)
	b := make([]Fraction, n, n)
	a := make([][]Fraction, n, n)
	for i := range a {
		a[i] = make([]Fraction, n, n)
	}
	x := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&x)
			a[i][j] = Fraction{x, 1}
		}
		fmt.Scan(&x)
		b[i] = Fraction{x, 1}
	}
	ans := gaus(a, b, n)
	if ans[0].a == 0 && ans[0].b == 0 {
		fmt.Println("No solution")
	} else {
		for i := 0; i < n; i++ {
			fmt.Printf("%d/%d\n", ans[i].a, ans[i].b)
		}
	}
}
