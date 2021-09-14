package lab1

import (
    "fmt"
)

func Task3(a Massiv, b []float64, eps float64) {
    a.print_mass()
    fmt.Println("B =", b)
    alpha := Create_mass(len(a))
    betta := make([]float64, len(a))
    to_equiv(a, b, betta, alpha)
    fmt.Println(Jacobi(alpha, betta, eps))
    fmt.Println(Seidel(alpha, betta, eps))
}

func to_equiv(a Massiv, b []float64, betta []float64, alpha Massiv) {
    for i := 0; i < len(a); i++ {
        betta[i] = b[i] / a[i][i]
        for j := 0; j < len(a); j++ {
            if i == j {
                alpha[i][j] = 0
            } else {
                alpha[i][j] = -a[i][j] / a[i][i]
            }
        }
    }
}


func Jacobi(alpha Massiv, betta []float64, eps float64) []float64 {
    x_prev := betta
    count := 0
    for {
        count++
        x := sum(multy(alpha, x_prev), betta)
        fmt.Printf("Jacobi Iteration %d: x = %f\n",count ,x)
        if Abs(norma(minus(x, x_prev))) <= eps {
            fmt.Println(count)
            return x
        }
        x_prev = x
    }
}

func Seidel (alpha Massiv, betta []float64, eps float64) []float64 {
    x_prev := betta
    count := 0
    for {
        count++
        x := make([]float64, len(alpha))
        for i := 0; i < len(x); i++ {
            x[i] += betta[i]
            for j := 0; j < len(x); j++ {
                x[i] += alpha[i][j] * x_prev[j]
            }
        }
        fmt.Printf("Seidel Iteration %d: x = %f\n",count ,x)
        if Abs(norma(minus(x, x_prev))) <= eps {
            return x
        }
        x_prev = x
    }
}

func sum(a []float64, b []float64) []float64{
    c := make([]float64, len(a))
    for i := 0; i < len(a);i++ {
        c[i] = a[i] + b[i]
    }
    return c
}

func minus(a []float64, b []float64) []float64{
    c := make([]float64, len(a))
    for i := 0; i < len(a);i++ {
        c[i] = a[i] - b[i]
    }
    return c
}

func norma(a []float64) float64{
    res := float64(0)
    for i := 0; i < len(a); i++ {
        res += a[i]
    }
    return res
}

func mass_sum(a Massiv, b Massiv) Massiv{
    c := Create_mass(len(a))
    for i := 0; i < len(a);i++ {
        for j := 0; j < len(b); j++ {
            c[i][j] = a[i][j] + b[i][j]
        }
    }
    return c
}

func multy(a Massiv, b []float64) []float64 {
    res := make([]float64, len(b))
    for i := 0; i < len(b); i++ {
        for j := 0; j < len(b); j++ {
            res[i] += a[i][j] * b[j]
        }
    }
    return res
}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
