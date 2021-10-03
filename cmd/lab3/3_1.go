package lab3

import (
    "fmt"
    "math"
)

func Task1() {
    a := []float64{math.Pi/8, 2*math.Pi/8, 3*math.Pi/8, math.Pi/2}
    b := []float64{math.Pi/8, math.Pi/3, 3*math.Pi/8, math.Pi/2}
    x := 3 * math.Pi/16
    polynom, res := LagrangePoly(a, x)
    fmt.Println("CASE A:\n Lagrange\nresult is ", res)
    fmt.Println(polynom)
    fmt.Println("Absolute error is", math.Abs(foo(x) - res))
    res, polynom = NewtonPoly(a, x)
    fmt.Println("\nNewton\nresult is ", res)
    fmt.Println(polynom)
    polynom, res = LagrangePoly(b, x)
    fmt.Println("CASE B:\n Lagrange\nresult is ", res)
    fmt.Println(polynom)
    fmt.Println("Absolute error is", math.Abs(foo(x) - res))
    res, polynom = NewtonPoly(b, x)
    fmt.Println("\nNewton\nresult is ", res)
    fmt.Println(polynom)
}

func foo(x float64) float64{
    a, b := math.Sincos(x)
    return b/a + x
}

func omega(a []float64, x float64, j int) float64{
    val := 1.
    for i, v := range a {
        if i != j {
            val *= x - v
        }
    }
    return val
}

func omegaString(a []float64, j int) string{
    res := ""
    for i, _ := range a {
        if i != j {
            res += fmt.Sprintf("(x - %.4f)", a[i])
        }
    }
    return res
}


func sign(f float64) string{
    if f > 0 {
        return " + "
    } else {
        return ""
    }
}

func LagrangePoly(x []float64, exactX float64) (string, float64){
    var res, f_i float64
    polynom := ""
    for i, _ := range x {
        f_i = foo(x[i]) / omega(x, x[i], i)
        res += f_i * omega(x, exactX, i)
        polynom += sign(f_i) + fmt.Sprintf(" %.4f", f_i) + omegaString(x, i)
    }
    return polynom, res
}

func makePoly(x []float64, i int) string{
    res := ""
    for j := 0; j < i; j++ {
        res += fmt.Sprintf("(x - %.4f)", x[j])
    }
    return res
}

func NewtonPoly(x []float64, exactX float64) (float64, string) {
    y := make([]float64, len(x))
    polynom := ""
    for i, _ := range y {
        y[i] = foo(x[i])
    }
    table := makeTable(x, y)
    res := 0.
    tmp := 1.
    for i, _ := range table {
        res += table[i] * tmp
        tmp *= exactX - x[i]
    }
    for i, _ := range(table) {
        if table[i] > 0 {
            polynom += " + "
        }
        polynom += fmt.Sprintf("%f", table[i]) + makePoly(x, i)
    }
    return res, polynom
}

func makeTable(x []float64, y []float64) []float64{
    table := y
    for j := 1; j < len(x); j++ {
        for i := len(x) - 1; i > j - 1; i-- {
            table[i] = (table[i] - table[i - 1]) / (x[i] - x[i - j])
        }
    }
    return table
}
