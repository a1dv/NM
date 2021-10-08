package lab3

import (
    "fmt"
)

func Task2() {
    x := []float64{0.,1.,2.,3.,4.}
    f := []float64{0., 1.8415, 2.9093, 3.1411, 3.2432}
    exactX := 1.5
    h := make([]float64, len(x))
    h[0] = 0
    for i := 1; i < len(h); i++ {
        h[i] = x[i] - x[i-1]
    }
    fmt.Printf("F(%.4f) is %f\n", exactX, quadroSpline(x, f, h, exactX))
}

func fun(a, b, c, d, x float64) float64{
    fmt.Printf("F(x) = %.4f + %.4f * x + %.4f * x^2 + %.4f * x ^3\n", a, b, c, d)
    return a + b * x + c * (x * x) + d * (x * x * x)
}

func a(f []float64) []float64{
    res := make([]float64, 0)
    for i := 1; i < len(f); i++{
        res = append(res, f[i - 1])
    }
    return res
}

func b(f []float64, h []float64, cM []float64) []float64{
    res := make([]float64, 0)
    n := len(cM)
    for i := 1; i < n; i++ {
        res = append(res, f[i] - f[i-1]/h[i] - h[i]*(cM[i] + 2*cM[i-1])/3)
    }
    res = append(res, ((f[n] - f[n - 1]) / h[n] - 2 * h[n] * cM[n-1] / 3))
    return res
}


func c(x []float64, f []float64, h []float64) []float64{
    n := len(f)
    a := make([]float64, 0)
    a = append(a, 0)
    for i := 3; i < n; i++ {
        a = append(a, h[i-1])
    }
    b := make([]float64, 0)
    for i := 2; i < n; i++ {
        b = append(b, 2 * (h[i-1] + h[i]))
    }
    c := make([]float64, 0)
    for i := 2; i < n - 1; i++ {
        c = append(c, h[i])
    }
    c = append(c, 0)
    d := make([]float64, 0)
    for i := 2; i < n; i++ {
        d = append(d, 3 * ((f[i] - f[i - 1]) / h[i] - ((f[i - 1] - f[i - 2]) / h[i - 1])))
    }
    tri := make([][]float64, 0)
    tri = append(tri, a)
    tri = append(tri, b)
    tri = append(tri, c)
    res := make([]float64, 0)
    res = append(res, 0)
    tmp := tridiag_solver(tri, d)
    for _, v := range tmp {
        res = append(res, v)
    }
    return res
}

func d(h []float64, c []float64) []float64{
    res := make([]float64, 0)
    n := len(c) - 1
    res = append(res, (c[1] - c[0]) / 3)
    fmt.Println(c)
    for i := 1; i < n; i ++ {
        res = append(res, (c[i + 1] - c[i]) / (3 * h[i]))
    }
    res = append(res, -c[n] / (3 * h[n]))
    return res
}

func quadroSpline(x []float64, f []float64, h []float64, exactX float64) float64{
    aM := a(f)
    cM := c(x, f, h)
    bM := b(f, h, cM)
    dM := d(h, cM)
    fmt.Println("a : ", aM)
    fmt.Println("b : ", bM)
    fmt.Println("c : ", cM)
    fmt.Println("d : ", dM)
    for i := 1; i < len(x); i++ {
        if x[i] >= exactX && x[i-1] <= exactX {
            return fun(aM[i-1], bM[i-1], cM[i-1], dM[i-1], exactX - x[i-1])
        }
    }
    return -1
}


func tridiag_solver(a [][]float64, b []float64) []float64 {
    n := len(a)
    x := make([]float64, n)
    p := make([]float64, 0)
    q := make([]float64, 0)
    p = append(p,(-a[2][0] / a[1][0]))
    q = append(q,(b[0] / a[1][0]))

    for i := 1; i < n;i++ {
        if i == n - 1 {
            p = append(p, 0.)
        } else {
            p = append(p, (-a[2][i] / (a[1][i] + a[0][i] * p[i - 1])))
        }
        q = append(q, (b[i] - a[0][i] * q[i - 1]) / (a[1][i] + a[0][i] * p[i - 1]))
    }
    x[n - 1] = q[n - 1]
    for i := n - 2; i > -1; i-- {
        x[i] = p[i] * x[i + 1] + q[i]
    }
    return x
}
