package lab2

import (
    "fmt"
    "math"
    "github/NM/cmd/lab1"
)

func Task2(eps float64) {
    fmt.Println(Newton_for_systems(eps))
    fmt.Println(Simple_iterations_for_systems(eps))
}

func f1(x1 float64, x2 float64) float64{
    _, c2 := math.Sincos(x2)
    return x1 - c2 - 3
}

func f2(x1 float64, x2 float64) float64{
    s1, _ := math.Sincos(x1)
    return x2 - s1 - 3
}

func df1_dx1(x1 float64, x2 float64) float64{
    return 1.
}

func df1_dx2(x1 float64, x2 float64) float64{
    s2, _ := math.Sincos(x2)
    return s2
}

func df2_dx1(x1 float64, x2 float64) float64{
    _, c1 := math.Sincos(x1)
    return - c1
}
func df2_dx2(x1 float64, x2 float64) float64{

    return 1.
}

func Jacobi_Matrix(x1 float64, x2 float64) lab1.Massiv{
    j := lab1.Create_mass(2)
    j[0][0] = df1_dx1(x1, x2)
    j[0][1] = df1_dx2(x1, x2)
    j[1][0] = df2_dx1(x1, x2)
    j[1][1] = df2_dx2(x1, x2)

    return j
}

func a_1 (x1 float64, x2 float64) lab1.Massiv{
    a1 := lab1.Create_mass(2)
    a1[0][0] = f1(x1, x2)
    a1[0][1] = df1_dx2(x1, x2)
    a1[1][0] = f2(x1, x2)
    a1[1][1] = df2_dx2(x1, x2)
    return a1
}

func a_2 (x1 float64, x2 float64) lab1.Massiv{
    a2 := lab1.Create_mass(2)
    a2[0][0] = df1_dx1(x1, x2)
    a2[0][1] = f1(x1, x2)
    a2[1][0] = df2_dx1(x1, x2)
    a2[1][1] = f2(x1, x2)
    return a2
}

func Square_Determinant(m lab1.Massiv) float64{
    return m[0][0] * m[1][1] - m[0][1] * m[1][0];
}

func Newton_for_systems(eps float64) lab1.Massiv{
    x_prev := lab1.Undef_mass(2, 1)
    x := lab1.Undef_mass(2, 1)
    x[0][0] = 2.5
    x[1][0] = -1.7

    count := 0

    for {
        for i := 0; i < 2; i++ {
            x_prev[i][0] = x[i][0]
        }
        count++
        j := Jacobi_Matrix(x_prev[0][0], x_prev[1][0])
        a1 := a_1(x_prev[0][0], x_prev[1][0])
        a2 := a_2(x_prev[0][0], x_prev[1][0])


        det_j := Square_Determinant(j)
        det_a1 := Square_Determinant(a1)
        det_a2 := Square_Determinant(a2)
        x[0][0] = x_prev[0][0] - det_a1 / det_j
        x[1][0] = x_prev[1][0] - det_a2 / det_j
        fmt.Printf("Newton Iteration %d:\n x_1 = %f\n x_2 = %f\n", count, x[0][0], x[1][0])
        if (norm(lab1.Subt(x, x_prev)) <= eps) {
            break
        }
    }
    return x
}


func norm(a lab1.Massiv) float64{
    res := -1e10
    for i := 0; i < len(a); i++ {
        if a[i][0] > res {
            res = a[i][0]
        }
    }
    return res
}


func phi1(x1 float64 , x2 float64) float64{
    _, c2 := math.Sincos(x2)
    return 3 - c2
}

func phi2(x1 float64, x2 float64) float64{
    s1, _ := math.Sincos(x1)
    return 3 - s1
}

func dphi1_dx1 (x1 float64, x2 float64) float64{
    return 0.
}

func dphi1_dx2 (x1 float64, x2 float64) float64{
    s2, _ := math.Sincos(x2)
    return s2
}

func dphi2_dx1 (x1 float64, x2 float64) float64{
    _, c1 := math.Sincos(x1)
    return -c1
}

func dphi2_dx2 (x1 float64, x2 float64) float64{
    return 0.
}


func Simple_iterations_for_systems(eps float64) lab1.Massiv{
    q := 0.5;
    x_prev := lab1.Undef_mass(2, 1)
    x := lab1.Undef_mass(2, 1)
    x[0][0] = 2.5
    x[1][0] = -1.7

    count := 0
    for
    {
        for i := 0; i < 2; i++ {
            x_prev[i][0] = x[i][0]
        }
        x[0][0] = phi1(x_prev[0][0], x_prev[1][0])
        x[1][0] = phi2(x_prev[0][0], x_prev[1][0])
        count++

        fmt.Printf("Simple Iteration %d:\n x_1 = %f\n x_2 = %f\n", count, x[0][0], x[1][0])
        if (q / (1-q) * norm(lab1.Subt(x, x_prev)) <= eps) {
            break
        }
    }

    return x

}
