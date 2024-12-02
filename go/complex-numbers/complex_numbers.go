package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	real      float64
	imaginary float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imaginary
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
		real:      n1.real + n2.real,
		imaginary: n1.imaginary + n2.imaginary,
	}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		real:      n1.real - n2.real,
		imaginary: n1.imaginary - n2.imaginary,
	}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		real:      n1.real*n2.real - n1.imaginary*n2.imaginary,
		imaginary: n1.imaginary*n2.real + n1.real*n2.imaginary,
	}
}

func (n Number) Times(factor float64) Number {
	return Number{
		real:      factor * n.real,
		imaginary: factor * n.imaginary,
	}
}

func (n1 Number) Divide(n2 Number) Number {
	div := n2.real*n2.real + n2.imaginary*n2.imaginary
	return Number{
		real:      (n1.real*n2.real + n1.imaginary*n2.imaginary) / div,
		imaginary: (n1.imaginary*n2.real - n1.real*n2.imaginary) / div,
	}
}

func (n Number) Conjugate() Number {
	return Number{
		real:      n.real,
		imaginary: -n.imaginary,
	}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.real*n.real + n.imaginary*n.imaginary)
}

func (n Number) Exp() Number {
	fct := math.Exp(n.real)
	return Number{
		real:      fct * math.Cos(n.imaginary),
		imaginary: fct * math.Sin(n.imaginary),
	}
}
