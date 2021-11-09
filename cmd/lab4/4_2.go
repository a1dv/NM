package lab4

import (
    "fmt"
    "github/NM/cmd/lab1"
    "github/NM/cmd/lab3"
    "math"
)

func Task2() {
    x0, x1 := 1.1, 2.1
    a0, a1, alpha := 0., 1., 0.
    b0, b1, beta := 1., -1., 0.
    h := 0.1
    Finite_Difference_Method(h, x0, x1, a0, a1, alpha, b0, b1, beta)
    Shooting_Method(h, x0, x1, a0, a1, alpha, b0, b1, beta)
}

//Task2

func y_exact(x float64) float64{
    return 1. + x + math.Log(x)
}

func TridiagonalAlgorithm(a lab1.Massiv, d []float64) []float64{
    n := len(a)
    x := make([]float64, n)
    p := make([]float64, 0)
    q := make([]float64, 0)
    p = append(p, -a[0][1] / a[0][1])
    q = append(q, d[0] / a[0][1])
    var p_i float64
    var q_i float64

    for i := 1; i < n; i++ {
        if i == n - 1 {
            p_i = 0
        } else {
            p_i = -a[i][2] / (a[i][1] + a[i][0] * p[i - 1])
        }
        q_i = (d[i] - a[i][0] * q[i - 1]) / (a[i][1] + a[i][0] * p[i - 1])
        p = append(p, p_i)
        q = append(q, q_i)
    }

    x[n - 1] = q[n - 1]
    for i := n - 2; i > -1; i-- {
        x[i] = p[i] * x[i + 1] + q[i]
    }
    return x
}

func P(x float64) float64{
    return 1 / (x * math.Log(x))
}

func Q(x float64) float64{
    return 1 / (x * x * math.Log(x))
}

func F(x float64) float64{
    return 0.;
}




func Finite_Difference_Method(h float64, x0 float64, x1 float64, a0 float64, a1 float64, alpha float64, b0 float64, b1 float64, beta float64) ([]float64, []float64){
    m := (x1 - x0) / h
    n := int(m)
    if !(a1 == 0) {
        a0 = a0 / a1
        alpha = alpha/a1
    }
    if !(b1 == 0) {
        b0 = b0 / b1
        beta = beta / b1
    }
    x := make([]float64, n + 1)
    x[0] = x0
    for i := 0; i != n; i++ {
        x[i + 1] = x[i] + h
    }

    a := lab1.Undef_mass(n + 1, 3)
    d := make([]float64, n + 1)

    a[0][0] = 0.
    a[0][1] = -2. / (h * (2 - P(x0) * h)) + Q(x0) * h / (2. - P(x0) * h) + a0
    a[0][2] = 2. / (h * (2 - P(x0) * h))
    d[0] = -1. * h

    for i := 1; i != n; i++ {
        a[i][0] = 1. / (h * h) - P(x[i]) / (2. * h)
        a[i][1] = -2. / (h * h) + Q(x[i])
        a[i][2] = 1. / (h * h) + P(x[i]) / (2. * h)
        d[i] = F(x[i])
    }

    a[n][0] = -2. / (h * (2 + P(x[n]) * h))
    a[n][1] = 2 / (h * (2. + P(x[n]) * h)) - Q(x[n]) * h / (2. + P(x[n]) * h) + b0
    a[n][2] = 3. * h
    d[n] = beta - h * F(x[n]) / (2. + P(x[n]) * h)
    y := TridiagonalAlgorithm(a, d)

    fmt.Println(x)
    fmt.Println(y)

    return x, y


}
func Runge_Estimate_for_Finite_Defference(y []float64, h float64, x0 float64, x1 float64, a0 float64, a1 float64, alpha float64, b0 float64, b1 float64, beta float64) []float64{
    _, yHalf := Finite_Difference_Method(h / 2., x0, x1, a0, a1, alpha, b0, b1, beta)
    n := len(y)

    phi := make([]float64, n)

    for i := 0; i != n; i++ {
        phi[i] = lab3.RungeRomberg(yHalf[2 * i], y[i], 1)
    }
    return phi
}

//Shooting


func f1S(x float64, y float64, z float64) float64{
    return z
}

func f2S(x float64, y float64, z float64) float64{
    return (x * z - y) / (x * x * math.Log(x))
}


func Runge_Kutta_Method_for_Shooting(h float64, x0 float64, x1 float64, y0 float64, z0 float64) ([]float64, []float64, []float64){
    m := (x1 - x0) / h
    n := int(m)
    x := make([]float64, n + 1)
    y := make([]float64, n + 1)
    z := make([]float64, n + 1)
    x[0] = x0;
    y[0] = y0;
    z[0] = z0;


    for i := 0; i != n; i++ {
        k1 := h * f1S(x[i], y[i], z[i])
        l1 := h * f2S(x[i], y[i], z[i])
        k2 := h * f1S(x[i] + h / 2., y[i] + k1 / 2., z[i] + l1 / 2.)
        l2 := h * f2S(x[i] + h / 2., y[i] + k1 / 2., z[i] + l1 / 2.)
        k3 := h * f1S(x[i] + h / 2., y[i] + k2 / 2., z[i] + l2 / 2.)
        l3 := h * f2S(x[i] + h / 2., y[i] + k2 / 2., z[i] + l2 / 2.)
        k4 := h * f1S(x[i] + h, y[i] + k3, z[i] + l3)
        l4 := h * f2S(x[i] + h, y[i] + k3, z[i] + l3)

        deltaY := (k1 + (2. * k2) + (2. * k3) + k4) / 6.
        deltaZ := (l1 + (2. * l2) + (2. * l3) + l4) / 6.

        x[i + 1] = x[i] + h
        y[i + 1] = y[i] + deltaY
        z[i + 1] = z[i] + deltaZ

    }
    //fmt.Println(x)
    //fmt.Println(y)
    //fmt.Println(z)
    return x, y, z
}


func Phi (b0 float64, b1 float64, ys float64, zs float64, beta float64) float64{
    return b0 * ys + b1 * zs - beta;
}


func Shooting_Method(h float64, x0 float64, x1 float64, a0 float64, a1 float64, alpha float64, b0 float64, b1 float64, beta float64) ([]float64, []float64){
    x := make([]float64, 0)
    y := make([]float64, 0)

    eps := 0.1
    m := (x1 - x0) / h
    n := int(m)
    s := make([]float64, 0)

    var c0, c1 float64
    if a0 == 0. {
        c0 = -1. / a1
        c1 = 0.
    } else {
        c0 = 0.
        c1 = -1. / a0
    }

    s = append(s, (x1 + x0) / 2.)
    s = append(s, s[0] / 2.)

    ys := make([]float64, 0)
    zs := make([]float64, 0)

    y0 := a1 * s[0] - c1 * alpha;
    z0 := a0 * s[0] - c0 * alpha;
    _, Yans1, Zans1 := Runge_Kutta_Method_for_Shooting(h, x0, x1, y0, z0)
    ys = append(ys, Yans1[n])
    zs = append(zs, Zans1[n])

    y0 = a1 * s[1] - c1 * alpha;
    z0 = a0 * s[1] - c0 * alpha;
    Xans2, Yans2, Zans2 := Runge_Kutta_Method_for_Shooting(h, x0, x1, y0, z0)
    ys = append(ys, Yans2[n])
    zs = append(zs, Zans2[n])




    i := 2
    for {
        currentS := s[i - 1] - (s[i - 1] - s[i - 2]) / (Phi(b0, b1, ys[i - 1], zs[i - 1], beta) - Phi(b0, b1, ys[i - 2], zs[i - 2], beta)) * Phi(b0, b1, ys[i - 1], zs[i - 1], beta)
        s = append(s, currentS)
        y0 = a1 * s[i] - c1 * alpha;
        z0 = a0 * s[i] - c0 * alpha;
        Xans2, Yans2, _ = Runge_Kutta_Method_for_Shooting(h, x0, x1, y0, z0);
        ys = append(ys, Yans2[n])
        zs = append(zs, Zans2[n])


        if (math.Abs(Phi(b0, b1, ys[i], zs[i], beta)) < eps) {
            x = Xans2
            y = Yans2
            break
        }

        i++

    }

    fmt.Println(x)
    fmt.Println(y)

    return x, y
}




func Runge_Estimate_for_Shooting(y []float64,h float64, x0 float64, x1 float64, a0 float64, a1 float64, alpha float64, b0 float64, b1 float64, beta float64) []float64{

    _, yHalf := Shooting_Method(h / 2., x0, x1, a0, a1, alpha, b0, b1, beta);
    n := len(y)

    phi := make([]float64, n)

    for i := 0; i != n; i ++ {
        phi[i] = lab3.RungeRomberg(yHalf[2 * i], y[i], 1)
    }
    return phi
}
