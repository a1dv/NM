package lab3

import (
    "fmt"
    "math"
)


func Task5() {
    x0 := -1.
    xk := 1.
    h1 := 0.5
    h2 := 0.25
    simp1 := SimpsonMethod(xk, x0, h1)
    simp2 := SimpsonMethod(xk, x0, h2)
    trap1 := TrapezeMethod(xk, x0, h1)
    trap2 := TrapezeMethod(xk, x0, h2)
    rect1 := RectangleMethod(xk, x0, h1)
    rect2 := RectangleMethod(xk, x0, h2)
    fmt.Println("RECTANGLE|TRAPEZOID|SIMPSON")
    fmt.Println("for h1: ", rect1, trap1, simp1)
    fmt.Println("for h2: ", rect2, trap2, simp2)
    fmt.Println("Runge Romberg")
    fmt.Println("Rectangle: ", RungeRomberg(rect1, rect2, 1))
    fmt.Println("Trapezoid: ", RungeRomberg(trap1, trap2, 2))
    fmt.Println("Simpson: ", RungeRomberg(simp1, simp2, 3))
}


func RectangleMethod(xk float64, x0 float64, h float64) float64{
    tmp := 0.

    x := x0
    for {
        tmp += f((x + x + h) / 2.)
        x += h
        if x == xk {
            break
        }
    }
    tmp *= h
    return tmp
}

func TrapezeMethod(xk float64, x0 float64, h float64) float64{
    tmp := f(x0)/2
    x := x0+h
    for {
        tmp += f(x)
        x += h
        if x == xk {
            break
        }
    }
    tmp += f(x) / 2
    tmp *= h
    return tmp
}

func SimpsonMethod(xk float64, x0 float64, h float64) float64{
    n := xk - x0 / h
    m := int(n)
    tmp := 0.
    x := x0
    for i := 0; i != m - 1; i++ {
        tmp += (f(x) + 4 * f(x + h) + f(x + 2 * h))
        x += 2 * h
    }
    tmp *= (h / 3.)
    return tmp
}

func f(a float64) float64{
    return a / math.Pow(3*a + 4, 2)
}

func RungeRomberg(f2 float64, f1 float64, k float64) float64{
    return (f2 - f1) / (math.Pow(2, k) - 1)
}
