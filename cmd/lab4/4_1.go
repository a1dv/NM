package lab4

import (
    "math"
    "fmt"
    "github/NM/cmd/lab3"
)

func Task1() {
    x0 := 1.
    x1 := 2.
    y0 := 1.
    z0 := 1.
    h := 0.1
    _, b := EulerMethod(h, x0, x1, y0, z0)
    fmt.Println(RungeEstimateForEuler(b, h, x0, x1, y0, z0))
    _, d, _ := RungeKuttaMethod(h, x0, x1, y0, z0)
    fmt.Println(RungeEstimateForRungeKutta(d, h, x0, x1, y0, z0))
    _, g := AdamsMethod(h, x0, x1, y0, z0)
    fmt.Println(RungeEstimateForAdams(g, h, x0, x1, y0, z0))
}

func eps(y1 float64, y2 float64) float64{
    return math.Abs(y1 - y2)
}

//Task1


func y_exact1(x float64) float64{
    return 1 + math.Log(x)

}

func f1(x float64, y float64, z float64) float64{
    return z
}

func f2(x float64, y float64, z float64) float64{
    return -z / x
}



//Euler

func EulerMethod(h float64, x0 float64, x1 float64, y0 float64, z0 float64) ([]float64, []float64){
    m := (x1 - x0) / h
    n := int(m)
    x := make([]float64, n + 1)
    y := make([]float64, n + 1)
    z := make([]float64, n + 1)


    x[0] = x0
    y[0] = y0
    z[0] = z0


    for i := 0; i != n; i++ {
        x[i + 1] = x[i] + h
        y[i + 1] = y[i] + h * f1(x[i], y[i], z[i])
        z[i + 1] = z[i] + h * f2(x[i], y[i], z[i])
    }
    fmt.Println("EULER")
    fmt.Println(x)
    fmt.Println(y)

    return x, y
}


func RungeEstimateForEuler(y []float64, h float64, x0 float64, x1 float64, y0 float64, z0 float64) []float64{
    _, yHalf := EulerMethod(h / 2., x0, x1, y0, z0)
    n := len(y)
    phi := make([]float64, n)

    for i := 0; i != n; i++ {
        phi[i] = lab3.RungeRomberg(yHalf[2 * i], y[i], 1)
    }
    return phi

}


//Runge-Kutta


func RungeKuttaMethod(h float64, x0 float64, x1 float64, y0 float64, z0 float64) ([]float64, []float64, []float64){
    m := (x1 - x0) / h
    n := int(m)
    x := make([]float64, n + 1)
    y := make([]float64, n + 1)
    z := make([]float64, n + 1)


    x[0] = x0
    y[0] = y0
    z[0] = z0


    for i := 0; i != n; i++ {
        k1 := h * f1(x[i], y[i], z[i])
        l1 := h * f2(x[i], y[i], z[i])
        k2 := h * f1(x[i] + h / 2., y[i] + k1 / 2., z[i] + l1 / 2.)
        l2 := h * f2(x[i] + h / 2., y[i] + k1 / 2., z[i] + l1 / 2.)
        k3 := h * f1(x[i] + h / 2., y[i] + k2 / 2., z[i] + l2 / 2.)
        l3 := h * f2(x[i] + h / 2., y[i] + k2 / 2., z[i] + l2 / 2.)
        k4 := h * f1(x[i] + h, y[i] + k3, z[i] + l3)
        l4 := h * f2(x[i] + h, y[i] + k3, z[i] + l3)

        deltaY := (k1 + (2. * k2) + (2. * k3) + k4) / 6.
        deltaZ := (l1 + (2. * l2) + (2. * l3) + l4) / 6.

        x[i + 1] = x[i] + h
        y[i + 1] = y[i] + deltaY
        z[i + 1] = z[i] + deltaZ

    }

    fmt.Println("RUNGEKUTTA")
    fmt.Println(x)
    fmt.Println(y)

    return x, y, z
}


func RungeEstimateForRungeKutta(y []float64, h float64, x0 float64, x1 float64, y0 float64, z0 float64) []float64{
    _ , yHalf, _ := RungeKuttaMethod(h / 2., x0, x1, y0, z0);
    n := len(y)
    phi := make([]float64, n)

    for i := 0; i != n; i++ {
        phi[i] = lab3.RungeRomberg(yHalf[2 * i], y[i], 4)
    }
    return phi
}


//Adams

func AdamsMethod(h float64, x0 float64, x1 float64, y0 float64, z0 float64) ([]float64, []float64) {
    m := (x1 - x0) / h
    n := int(m)
    x_Runge, y_Runge, z_Runge := RungeKuttaMethod(h, x0, 4. * h + x0, y0, z0);

    first := make([]float64, n)
    sec := make([]float64, n)
    x := make([]float64, n + 1)
    y := make([]float64, n + 1)
    z := make([]float64, n + 1)


    for i := 0; i != len(x_Runge); i++ {
        x[i] = x_Runge[i];
        y[i] = y_Runge[i];
        z[i] = z_Runge[i];
        first[i] = f1(x[i], y[i], z[i]);
        sec[i] = f2(x[i], y[i], z[i]);
    }


    for i := 3; i != n; i++ {
        first[i] = f1(x[i], y[i], z[i]);
        sec[i] = f2(x[i], y[i], z[i]);
        x[i + 1] = x[i] + h;
        y[i + 1] = y[i] + h * (55. * first[i] - 59. * first[i - 1] + 37. * first[i - 2] - 9. * first[i - 3]) / 24.
        z[i + 1] = z[i] + h * (55. * sec[i] - 59. * sec[i - 1] + 37. * sec[i - 2] - 9. * sec[i - 3]) / 24.

    }
    fmt.Println("ADAMS ")
    fmt.Println(x)
    fmt.Println(y)

    return x,y

}

func RungeEstimateForAdams(y []float64, h float64, x0 float64, x1 float64, y0 float64, z0 float64) []float64{
    _, yHalf := AdamsMethod(h / 2., x0, x1, y0, z0)
    n := len(y)
    phi := make([]float64, n)

    for i := 0; i != n; i++ {
        phi[i] = lab3.RungeRomberg(yHalf[2 * i], y[i], 4)
    }
    return phi

}
