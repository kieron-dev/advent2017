package main

import (
	"fmt"
)

func main() {
	prog := compute()
	opt(prog)
	force(prog)
	opt(prog)
	dump(prog)
	fmt.Println("max", maxModel(prog))
	fmt.Println("min", minModel(prog))
}

type val struct {
	t        int
	op       string
	n        int
	l, r     *val
	min, max int
}

func maxModel(prog []*val) string {
	m := make([]byte, 14)
	for i := range m {
		m[i] = '?'
	}

	for _, v := range prog {
		var a, b, i, j int

		switch {
		case bin(bin(bin(inp(&i), "+", num(&a)), "+", num(&b)), "force", inp(&j))(v):
			a += b
			if a < 0 {
				m[i] = '9'
				m[j] = byte(9 + a + '0')
			} else {
				m[j] = '9'
				m[i] = byte(9 - a + '0')
			}
		}

	}

	return string(m)
}

func minModel(prog []*val) string {
	m := make([]byte, 14)
	for i := range m {
		m[i] = '?'
	}

	for _, v := range prog {
		var a, b, i, j int

		switch {
		case bin(bin(bin(inp(&i), "+", num(&a)), "+", num(&b)), "force", inp(&j))(v):
			a += b
			if a < 0 {
				m[j] = '1'
				m[i] = byte(1 - a + '0')
			} else {
				m[i] = '1'
				m[j] = byte(1 + a + '0')
			}
		}

	}

	return string(m)
}

func force(prog []*val) {
	max := make(map[*val]int)
	max[prog[len(prog)-1]] = 0

	updateMax := func(v *val, m int) {
		if old, ok := max[v]; ok && old < m {
			return
		}
		max[v] = m
	}

	for i := len(prog) - 1; i >= 0; i-- {
		v := prog[i]
		m, ok := max[v]
		if !ok {
			continue
		}
		if m > v.max {
			continue
		}
		var x, y *val
		var a int
		switch {
		default:
			// panic("force " + v.op)
		case num(&a)(v):
			if a > m {
				panic("force impossible")
			}
		case bin(any(&x), "+", any(&y))(v):
			updateMax(x, m-y.min)
			updateMax(y, m-x.min)
		case bin(any(&x), "*", any(&y))(v):
			if y.min > 0 {
				updateMax(x, m/y.min)
			}
			if x.min > 0 {
				updateMax(y, m/x.min)
			}
		case bin(any(&x), "%", any(&y))(v):

		case bin(any(&x), "/", num(&a))(v):
			updateMax(x, m*a+a-1)
		case bin(bin(any(&x), "==", any(&y)), "==", con(0))(v):
			v.op = "force"
			v.l = x
			v.r = y
			v.min = 0
			v.max = 0
			updateMax(x, y.max)
			updateMax(y, x.max)
		}
	}
	_ = max
}

func opt(prog []*val) {
	for _, v := range prog {
		var a, b int
		var x, y *val
		switch {
		case bin(num(&a), "+", num(&b))(v):
			setval(v, a+b)
		case bin(num(&a), "*", num(&b))(v):
			setval(v, a*b)
		case bin(num(&a), "/", num(&b))(v):
			setval(v, a/b)
		case bin(num(&a), "%", num(&b))(v):
			setval(v, a%b)
		case bin(con(0), "*", any(&x))(v),
			bin(any(&x), "*", con(0))(v):
			setval(v, 0)
		case bin(con(0), "+", any(&x))(v),
			bin(any(&x), "+", con(0))(v),
			bin(any(&x), "/", con(1))(v),
			bin(any(&x), "*", con(1))(v):
			*v = *x
		case bin(any(&x), "==", any(&y))(v) &&
			(x.min > y.max || y.min > x.max):
			setval(v, 0)
		case bin(num(&a), "==", num(&b))(v) && a == b:
			setval(v, 1)
		case bin(bin(bin(any(&y), "*", num(&b)), "+", any(&x)), "%", num(&a))(v) && a == b && x.max < a:
			*v = *x
		case bin(bin(bin(any(&y), "*", num(&b)), "+", any(&x)), "/", num(&a))(v) && a == b && x.max < a:
			*v = *y
		case bin(bin(any(&x), "+", any(&y)), "%", num(&a))(v) && x.max+y.max < a:
			*v = val{
				op: "+",
				l:  x,
				r:  y,
			}
		}

		switch v.op {
		default:
			panic("min/max " + v.op)
		case "force":
			v.min = 0
			v.max = 0
		case "num":
			v.min = v.n
			v.max = v.n
		case "inp":
			v.min = 1
			v.max = 9
		case "==":
			v.min = 0
			v.max = 1
		case "*":
			if v.l.min < 0 || v.r.min < 0 {
				panic("min/max neg *")
			}
			v.min = v.l.min * v.r.min
			v.max = v.l.max * v.r.max
		case "+":
			v.min = v.l.min + v.r.min
			v.max = v.l.max + v.r.max
		case "%":
			if v.r.op != "num" {
				panic("min/max % non-constant")
			}
			v.min = 0
			v.max = v.r.n - 1
		case "/":
			if v.r.op != "num" {
				panic("min/max / non-constant")
			}
			v.min = v.l.min / v.r.n
			v.max = v.l.max / v.r.n

		}
	}
}

func setval(v *val, n int) {
	*v = val{op: "num", n: n}
}

type matcher func(v *val) bool

func inp(n *int) matcher {
	return func(v *val) bool {
		if v.op == "inp" {
			*n = v.n
			return true
		}
		return false
	}
}

func any(p **val) matcher {
	return func(v *val) bool {
		*p = v
		return true
	}
}

func con(n int) matcher {
	return func(v *val) bool {
		return v.op == "num" && v.n == n || v.min == n && v.max == n
	}
}

func num(n *int) matcher {
	return func(v *val) bool {
		if v.op == "num" {
			*n = v.n
			return true
		}
		if v.min == v.max {
			*n = v.min
			return true
		}
		return false
	}
}

func bin(l matcher, op string, r matcher) matcher {
	return func(v *val) bool {
		return v.op == op && l(v.l) && r(v.r)
	}
}

func (v *val) Name() string {
	return fmt.Sprint("t", v.t)
}

func (v *val) Init() string {
	switch v.op {
	case "num":
		return fmt.Sprint(v.n)
	case "inp":
		return fmt.Sprint("m", v.n)
	default:
		return fmt.Sprintf("(%v %v %v)", v.l.Name(), v.op, v.r.Name())
	}
}

func (v *val) String() string {
	return fmt.Sprintf("%v = %v", v.Name(), v.Init())
}

func dump(prog []*val) {
	count := make(map[*val]int)
	for i := len(prog) - 1; i >= 0; i-- {
		v := prog[i]
		if count[v] == 0 && i != len(prog)-1 {
			continue
		}
		count[v.l]++
		count[v.r]++
	}

	str := make(map[*val]string)
	for _, v := range prog {
		var x string
		switch v.op {
		case "inp", "num":
			x = v.Init()
		default:
			x = fmt.Sprintf("(%v %v %v)", str[v.l], v.op, str[v.r])
			if count[v] > 1 || v.op == "force" {
				fmt.Printf("%v = %v // [%d,%d]\n", v.Name(), x, v.min, v.max)
				x = v.Name()
			}
		}
		str[v] = x
	}
	fmt.Println(str[prog[len(prog)-1]])
}

func compute() []*val {
	var prog []*val

	t := 0
	emit := func(v *val) *val {
		t++
		v.t = t
		prog = append(prog, v)
		return v
	}

	i := 0
	inp := func() *val {
		v := emit(&val{
			op: "inp",
			n:  i,
		})
		i++
		return v
	}
	bin := func(l *val, op string, r *val) *val {
		return emit(&val{
			op: op,
			l:  l,
			r:  r,
		})
	}
	add := func(a, b *val) *val { return bin(a, "+", b) }
	mul := func(a, b *val) *val { return bin(a, "*", b) }
	div := func(a, b *val) *val { return bin(a, "/", b) }
	mod := func(a, b *val) *val { return bin(a, "%", b) }
	eql := func(a, b *val) *val { return bin(a, "==", b) }
	num := func(n int) *val { return emit(&val{op: "num", n: n}) }

	w, x, y, z := num(0), num(0), num(0), num(0)

	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(12))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(9))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(12))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(4))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(12))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(2))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-9))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(5))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-9))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(1))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(14))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(6))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(14))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(11))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-10))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(15))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(15))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(7))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-2))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(12))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(1))
	x = add(x, num(11))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(15))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-15))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(9))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-9))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(12))
	y = mul(y, x)
	z = add(z, y)
	w = inp()
	x = mul(x, num(0))
	x = add(x, z)
	x = mod(x, num(26))
	z = div(z, num(26))
	x = add(x, num(-3))
	x = eql(x, w)
	x = eql(x, num(0))
	y = mul(y, num(0))
	y = add(y, num(25))
	y = mul(y, x)
	y = add(y, num(1))
	z = mul(z, y)
	y = mul(y, num(0))
	y = add(y, w)
	y = add(y, num(12))
	y = mul(y, x)
	z = add(z, y)

	return prog
}
