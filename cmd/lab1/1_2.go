package lab1

import (
    "fmt"
)

func Task2(a Massiv, b []float64) {
    a.print_mass()
    fmt.Println("B =", b)
    fmt.Println("Result:")
    fmt.Println(tridiag_solver(a, b))
}

func tridiag_solver(a Massiv, b []float64) []float64 {
    x := make([]float64, len(a))
    z := make([]float64, len(a))
    y := a[0][0]
    x[0] = (-a[0][1]) / y
    z[0] = b[0] / y
    n1 := len(a) - 1
    for i := 1; i < n1; i++ {
      y = a[i][i] + a[i][i - 1] * x[i - 1]
      x[i] = -a[i][i + 1] / y
      z[i] = (b[i] - a[i][i - 1] * z[i - 1]) / y
    }
    res := make([]float64, len(a))
    res[n1] = (b[n1] - a[n1][n1 - 1] * z[n1 - 1]) / (a[n1][n1] + a[n1][n1 - 1] * x[n1 - 1])
    for i := n1 - 1; i >= 0; i-- {
      res[i] = x[i] * res[i + 1] + z[i]
    }
    return res
}
