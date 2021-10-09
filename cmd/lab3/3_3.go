package lab3

import (
    "fmt"
    "math"
    "github/NM/cmd/lab1"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot"
)

func Task3() {
    x := []float64{0., 1.7, 3.4, 5.1, 6.8, 8.5}
    y := []float64{0., 1.3038, 1.8439, 2.2583, 2.6077, 2.9155}
    k1 := mls(1, x, y)
    print_coeffs(k1, 1)
    f_i := make([]float64,0)
    for i := range x {
        f_i = append(f_i, F(k1, float64(i)))
    }
    fmt.Println("Err: ", error(f_i, y))
    k2 := mls(2, x, y)
    print_coeffs(k2, 2)
    pts := plotter.XYs{{X: x[0], Y: f_i[0]}, {X: x[1], Y: f_i[1]}, {X: x[2], Y: f_i[2]}, {X: x[3], Y: f_i[3]}, {X: x[4], Y: f_i[4]}}
    line, err := plotter.NewLine(pts)
    if err != nil {
        panic(err)
    }
    scatter, err := plotter.NewScatter(pts)
    if err != nil {
        panic(err)
    }
    p := plot.New()
    p.Add(line, scatter)
    f_i = make([]float64,0)
    for i := range x {
        f_i = append(f_i, F(k2, float64(i)))
    }
    fmt.Println("Err: ", error(f_i, y))
    p.X.Label.Text = "X"
    p.Y.Label.Text = "Y"
    pts = plotter.XYs{{X: x[0], Y: f_i[0]}, {X: x[1], Y: f_i[1]}, {X: x[2], Y: f_i[2]}, {X: x[3], Y: f_i[3]}, {X: x[4], Y: f_i[4]}}
    line, err = plotter.NewLine(pts)
    if err != nil {
        panic(err)
    }
    scatter, err = plotter.NewScatter(pts)
    if err != nil {
        panic(err)
    }
    p.Add(line, scatter)
    if err := p.Save(400, 400, "points.png"); err != nil {
        panic(err)
    }
}


func F(coeffs []float64, x float64) float64{
    tmp := 0.
    for i, v := range coeffs {
        tmp += math.Pow(x, float64(i)) * v
    }
    return tmp
}


func error(f []float64, y []float64) float64{
    tmp := 0.
    for i := range f {
        tmp += math.Pow(f[i] - y[i], 2)
    }
    return tmp
}

func print_coeffs(coeffs []float64, j int) {
    fmt.Printf("F_%d(x) = ", j)
    for i, v := range coeffs {
        fmt.Printf("%f*x^%d ", v, i)
        if i != len(coeffs) - 1 && coeffs[i + 1] > 0 {
            fmt.Printf("+ ")
        }
    }
    fmt.Printf("\n")
}

func mls(n int, x []float64, y []float64) []float64{
    tmp := 0.
    a := lab1.Create_mass(n + 1)
    for j := 0; j < n + 1; j++ {
        for i := 0; i < n + 1; i++ {
            for _, v := range x {
                tmp += math.Pow(v, float64(i+j))
            }
            a[j][i] = tmp
            tmp = 0.
        }
    }
    a[0][0] = float64(len(x))
    b := make([]float64, n + 1)
    for i := 0; i < n + 1; i ++ {
        tmp = 0.
        for j := range x {
            tmp += math.Pow(x[j], float64(i)) * y[j]
        }
        b[i] = tmp
    }
    l := lab1.Create_mass(len(a))
    u := lab1.Create_mass(len(a))
    k := make([]float64, len(a))
    z := make([]float64, len(a))
    lab1.LU(a, l, u)
    lab1.Solve(u, a, l, z, k, b)
    return k
}
