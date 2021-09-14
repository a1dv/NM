package lab2

import (
    "fmt"
    "math"
    "github/NM/cmd/lab1"
)

func Task1(eps float64) {
    a := float64(-1)
    b := float64(2)
    fmt.Println("X_1 is")
    fmt.Println(newton(a, b, eps))
    fmt.Println("X_2 is")
    fmt.Println(simple_iteration(a, b, eps))
}

func simple_iteration(a float64, b float64, eps float64) float64{
    x_prev := (a + b) / 2
    q := float64(2)
    x := float64(0)
    count := 0
    for {
        x = phi(x_prev)
        fmt.Printf("Simple Iteration %d: \n x = %f,\n x_prev= %f\n", count, x, x_prev)
        count++
        if lab1.Abs(q / (1 - q)) * lab1.Abs(x - x_prev) < eps {
            break
        }
        x_prev = x

    }
    return x
}

func newton(a float64, b float64, eps float64) float64 {
    x_prev := float64(-1)
    x := float64(0)
    if foo(a) * second_proisv(a) > 0 {
        x_prev = a
    } else {
        x_prev = b
    }
    count := 0
    for {
        x = x_prev - foo(x_prev)/proisv(x_prev)
        if lab1.Abs(x - x_prev) < eps {
            break
        }
        x_prev = x
        fmt.Printf("Newton Iteration %d:\n x = %f,\n f = %f,\n f'= %f\n", count, x, foo(x), proisv(x))
        count++
    }
    return x
}

func phi(x float64) float64 {
    s, _ := math.Sincos(x)
    return math.Sqrt(s + 1)
}

func foo(x float64) float64{
    s, _ := math.Sincos(x)
    return s - x*x + 1
}

func proisv(x float64) float64{
    _, c := math.Sincos(x)
    return c - 2*x
}

func second_proisv(x float64) float64 {
    s, _ := math.Sincos(x)
    return - s - 2
}
