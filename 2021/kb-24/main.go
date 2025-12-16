package main

import "fmt"

func main() {
	prog := compute()
	opt(prog)
	force(prog)
	opt(prog)
	dump(prog)
}

type val struct {
	t        int
	op       string
	n        int
	l        *val
	r        *val
	min, max int
}

func force(prog []*val) {
	max := make(map[*val]int)

	max[prog[len(prog)-1]] = 0

	setmax := func(v *val, m int) {
		old, ok := max[v]
		if !ok || m < old {
			max[v] = m
		}
	}

	for i := len(prog) - 1; i >= 0; i-- {
		v := prog[i]

		m, ok := max[v]
		if !ok {
			continue
		}

		var x, y *val
		var a int

		_ = x
		_ = y
		_ = a
		_ = m
		_ = setmax

		switch {
		default:
			panic("force " + v.op)
		case bin(any(&x), "*", any(&y))(v):
			if y.min > 0 {
				setmax(x, m/y.min)
			}
			if x.min > 0 {
				setmax(y, m/x.min)
			}

		case bin(bin(any(&x), "==", any(&y)), "==", con(0))(v) && m == 0:
			v.op = "force"
			v.l = x
			v.r = y
			v.min = 0
			v.max = 0
		}
	}
}

func opt(prog []*val) {
	for _, v := range prog {

		var x, y *val
		var a, b int

		_ = x
		_ = y
		_ = a
		_ = b

		switch {
		case bin(num(&a), "+", num(&b))(v):
			setval(v, a+b)
		case bin(num(&a), "*", num(&b))(v):
			setval(v, a*b)
		case bin(num(&a), "/", num(&b))(v):
			setval(v, a/b)
		case bin(num(&a), "%", num(&b))(v):
			setval(v, a%b)
		case bin(con(1), "*", any(&x))(v),
			bin(any(&x), "*", con(1))(v),
			bin(any(&x), "+", con(0))(v),
			bin(con(0), "+", any(&x))(v),
			bin(any(&x), "/", con(1))(v):
			*v = *x
		case bin(con(0), "*", any(&x))(v),
			bin(any(&x), "*", con(0))(v),
			bin(con(0), "%", any(&x))(v),
			bin(any(&x), "==", any(&y))(v) && (x.min > y.max || y.min > x.max):
			setval(v, 0)
		case bin(num(&a), "==", num(&b))(v) && a == b:
			setval(v, 1)
		}

		switch v.op {
		default:
			panic("min/max " + v.op)
		case "inp":
			v.min = 1
			v.max = 9
		case "num":
			v.min = v.n
			v.max = v.n
		case "*":
			v.min = v.l.min * v.r.min
			v.max = v.l.max * v.r.max
		case "+":
			v.min = v.l.min + v.r.min
			v.max = v.l.max + v.r.max
		case "/":
			if v.r.op != "num" {
				panic("min/max div non-num")
			}
			v.min = v.min / v.r.n
			v.max = v.max / v.r.n
		case "%":
			v.min = 0
			v.max = v.r.max - 1
		case "==":
			v.min = 0
			v.max = 1
		case "force":

		}

	}
}

func setval(v *val, n int) {
	*v = val{op: "num", n: n}
}

type matcher func(*val) bool

func con(n int) matcher {
	return func(v *val) bool {
		return (v.op == "num" && v.n == n) ||
			(v.min == v.max && v.min == n)
	}
}

func inp(n *int) matcher {
	return func(v *val) bool {
		if v.op == "inp" {
			*n = v.n
			return true
		}
		return false
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

func any(p **val) matcher {
	return func(v *val) bool {
		*p = v
		return true
	}
}

func dump(prog []*val) {
	count := make(map[*val]int)

	for _, v := range prog {
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
			if count[v] > 1 {
				fmt.Printf("%v = %v // [%d,%d]\n", v.Name(), x, v.min, v.max)
				x = v.Name()
			}
		}
		str[v] = x
	}
	fmt.Println(str[prog[len(prog)-1]])
}

func (v *val) Name() string {
	return fmt.Sprint("t", v.t)
}

func (v *val) Init() string {
	switch v.op {
	case "inp":
		return fmt.Sprint("m", v.n)
	case "num":
		return fmt.Sprint(v.n)
	default:
		return fmt.Sprintf("(%v %v %v)", v.l, v.op, v.r)
	}
}

func (v *val) String() string {
	return fmt.Sprintf("%v = %v", v.Name(), v.Init())
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
		i++
		return emit(&val{op: "inp", n: i})
	}
	num := func(n int) *val { return emit(&val{op: "num", n: n}) }
	bin := func(l *val, op string, r *val) *val { return emit(&val{op: op, l: l, r: r}) }
	mul := func(l, r *val) *val { return bin(l, "*", r) }
	add := func(l, r *val) *val { return bin(l, "+", r) }
	div := func(l, r *val) *val { return bin(l, "/", r) }
	mod := func(l, r *val) *val { return bin(l, "%", r) }
	eql := func(l, r *val) *val { return bin(l, "==", r) }

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
